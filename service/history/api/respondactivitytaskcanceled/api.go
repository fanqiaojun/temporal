// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package respondactivitytaskcanceled

import (
	"context"
	"time"

	"go.temporal.io/server/api/historyservice/v1"
	"go.temporal.io/server/common"
	"go.temporal.io/server/common/definition"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
	"go.temporal.io/server/common/tasktoken"
	"go.temporal.io/server/service/history/api"
	"go.temporal.io/server/service/history/consts"
	historyi "go.temporal.io/server/service/history/interfaces"
	"go.temporal.io/server/service/history/workflow"
)

func Invoke(
	ctx context.Context,
	req *historyservice.RespondActivityTaskCanceledRequest,
	shard historyi.ShardContext,
	workflowConsistencyChecker api.WorkflowConsistencyChecker,
) (resp *historyservice.RespondActivityTaskCanceledResponse, retError error) {
	namespaceEntry, err := api.GetActiveNamespace(shard, namespace.ID(req.GetNamespaceId()))
	if err != nil {
		return nil, err
	}
	namespace := namespaceEntry.Name()

	request := req.CancelRequest
	tokenSerializer := tasktoken.NewSerializer()
	token, err0 := tokenSerializer.Deserialize(request.TaskToken)
	if err0 != nil {
		return nil, consts.ErrDeserializingToken
	}
	if err := api.SetActivityTaskRunID(ctx, token, workflowConsistencyChecker); err != nil {
		return nil, err
	}

	var activityStartedTime time.Time
	var taskQueue string
	var workflowTypeName string
	err = api.GetAndUpdateWorkflowWithNew(
		ctx,
		token.Clock,
		definition.NewWorkflowKey(
			token.NamespaceId,
			token.WorkflowId,
			token.RunId,
		),
		func(workflowLease api.WorkflowLease) (*api.UpdateWorkflowAction, error) {
			mutableState := workflowLease.GetMutableState()
			workflowTypeName = mutableState.GetWorkflowType().GetName()
			if !mutableState.IsWorkflowExecutionRunning() {
				return nil, consts.ErrWorkflowCompleted
			}

			scheduledEventID := token.GetScheduledEventId()
			if scheduledEventID == common.EmptyEventID { // client call CompleteActivityById, so get scheduledEventID by activityID
				scheduledEventID, err0 = api.GetActivityScheduledEventID(token.GetActivityId(), mutableState)
				if err0 != nil {
					return nil, err0
				}
			}
			ai, isRunning := mutableState.GetActivityInfo(scheduledEventID)

			// First check to see if cache needs to be refreshed as we could potentially have stale workflow execution in
			// some extreme cassandra failure cases.
			if !isRunning && scheduledEventID >= mutableState.GetNextEventID() {
				metrics.StaleMutableStateCounter.With(shard.GetMetricsHandler()).Record(
					1,
					metrics.OperationTag(metrics.HistoryRespondActivityTaskCanceledScope))
				return nil, consts.ErrStaleState
			}

			if !isRunning ||
				ai.StartedEventId == common.EmptyEventID ||
				(token.GetScheduledEventId() != common.EmptyEventID && token.Attempt != ai.Attempt) ||
				(token.GetVersion() != common.EmptyVersion && token.Version != ai.Version) {
				return nil, consts.ErrActivityTaskNotFound
			}

			// sanity check if activity is requested to be cancelled
			if !ai.CancelRequested {
				return nil, consts.ErrActivityTaskNotCancelRequested
			}

			if _, err := mutableState.AddActivityTaskCanceledEvent(
				scheduledEventID,
				ai.StartedEventId,
				ai.CancelRequestId,
				request.Details,
				request.Identity); err != nil {
				// Unable to add ActivityTaskCanceled event to history
				return nil, err
			}

			activityStartedTime = ai.StartedTime.AsTime()
			taskQueue = ai.TaskQueue
			return &api.UpdateWorkflowAction{
				Noop:               false,
				CreateWorkflowTask: true,
			}, nil
		},
		nil,
		shard,
		workflowConsistencyChecker,
	)

	if err == nil && !activityStartedTime.IsZero() {
		metrics.ActivityE2ELatency.With(
			workflow.GetPerTaskQueueFamilyScope(
				shard.GetMetricsHandler(), namespace, taskQueue, shard.GetConfig(),
				metrics.OperationTag(metrics.HistoryRespondActivityTaskCanceledScope),
				metrics.WorkflowTypeTag(workflowTypeName),
				metrics.ActivityTypeTag(token.ActivityType),
			),
		).Record(time.Since(activityStartedTime))
	}
	return &historyservice.RespondActivityTaskCanceledResponse{}, err
}
