// The MIT License
//
// Copyright (c) 2024 Temporal Technologies Inc.  All rights reserved.
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

package xdc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"sync/atomic"
	"testing"
	"time"

	"github.com/nexus-rpc/sdk-go/nexus"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/suite"
	commandpb "go.temporal.io/api/command/v1"
	commonpb "go.temporal.io/api/common/v1"
	enumspb "go.temporal.io/api/enums/v1"
	historypb "go.temporal.io/api/history/v1"
	nexuspb "go.temporal.io/api/nexus/v1"
	"go.temporal.io/api/operatorservice/v1"
	taskqueuepb "go.temporal.io/api/taskqueue/v1"
	workflowpb "go.temporal.io/api/workflow/v1"
	"go.temporal.io/api/workflowservice/v1"
	sdkclient "go.temporal.io/sdk/client"
	"go.temporal.io/server/common"
	"go.temporal.io/server/common/dynamicconfig"
	commonnexus "go.temporal.io/server/common/nexus"
	"go.temporal.io/server/common/nexus/nexustest"
	"go.temporal.io/server/common/testing/testvars"
	"go.temporal.io/server/components/callbacks"
	"go.temporal.io/server/components/nexusoperations"
	"go.temporal.io/server/tests/testcore"
	"google.golang.org/protobuf/types/known/durationpb"
)

type NexusStateReplicationSuite struct {
	xdcBaseSuite
}

func TestNexusStateReplicationTestSuite(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name                    string
		enableTransitionHistory bool
	}{
		{
			name:                    "DisableTransitionHistory",
			enableTransitionHistory: false,
		},
		{
			name:                    "EnableTransitionHistory",
			enableTransitionHistory: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := &NexusStateReplicationSuite{}
			s.enableTransitionHistory = tc.enableTransitionHistory
			suite.Run(t, s)
		})
	}
}

func (s *NexusStateReplicationSuite) SetupSuite() {
	s.dynamicConfigOverrides = map[dynamicconfig.Key]any{
		// Make sure we don't hit the rate limiter in tests
		dynamicconfig.FrontendGlobalNamespaceNamespaceReplicationInducingAPIsRPS.Key(): 1000,
		dynamicconfig.EnableNexus.Key():                  true,
		dynamicconfig.RefreshNexusEndpointsMinWait.Key(): 1 * time.Millisecond,
		callbacks.AllowedAddresses.Key():                 []any{map[string]any{"Pattern": "*", "AllowInsecure": true}},
	}
	suffix := "_" + common.GenerateRandomString(5)
	s.setupSuite([]string{"nexus_state_replication_active" + suffix, "nexus_state_replication_standby" + suffix})
}

func (s *NexusStateReplicationSuite) SetupTest() {
	s.setupTest()
}

func (s *NexusStateReplicationSuite) TearDownSuite() {
	s.tearDownSuite()
}

// TestNexusOperationEventsReplicated tests that nexus related operation events and state updates are replicated
// across clusters and that the operation machinary functions as expected when failover happens.
// General outline:
// 1. Start two clusters, cluster1 set to active, cluster2 set to standby.
// 2. Start a workflow on cluster1.
// 3. Schedule a nexus operation on cluster1. An error is injected to fail the operation on start.
// 4. Check the operation scheduled event and state changes are replicated to cluster2.
// 5. Failover to cluster2 and unblock the operation start by removing the injected error.
// 6. Wait for the operation to be started on cluster2.
// 7. Fail back to cluster1.
// 8. Complete the operation via callback on cluster1.
// 9. Check that the operation completion triggers a workflow task when we poll on cluster1.
// 10. Complete the workflow.
func (s *NexusStateReplicationSuite) TestNexusOperationEventsReplicated() {
	var callbackToken string
	var publicCallbackUrl string

	failStartOperation := atomic.Bool{}
	failStartOperation.Store(true)
	h := nexustest.Handler{
		OnStartOperation: func(ctx context.Context, service, operation string, input *nexus.LazyValue, options nexus.StartOperationOptions) (nexus.HandlerStartOperationResult[any], error) {
			if failStartOperation.Load() {
				return nil, errors.New("injected error for failing nexus start operation")
			}

			callbackToken = options.CallbackHeader.Get(commonnexus.CallbackTokenHeader)
			publicCallbackUrl = options.CallbackURL
			return &nexus.HandlerStartOperationResultAsync{OperationID: "test"}, nil
		},
	}
	listenAddr := nexustest.AllocListenAddress()
	nexustest.NewNexusServer(s.T(), listenAddr, h)

	ctx := testcore.NewContext()
	ns := s.createGlobalNamespace()
	endpointName := testcore.RandomizedNexusEndpoint(s.T().Name())

	// Set URL template after httpAPAddress is set, see commonnexus.RouteCompletionCallback.
	for _, cluster := range []*testcore.TestCluster{s.cluster1, s.cluster2} {
		cluster.OverrideDynamicConfig(
			s.T(),
			nexusoperations.CallbackURLTemplate,
			// We'll send the callback to cluster1, when we fail back to it.
			"http://"+s.cluster1.Host().FrontendHTTPAddress()+"/namespaces/{{.NamespaceName}}/nexus/callback")
	}

	// Nexus endpoints registry isn't replicated yet, manually create the same endpoint in both clusters.
	for _, cl := range []operatorservice.OperatorServiceClient{s.cluster1.OperatorClient(), s.cluster2.OperatorClient()} {
		_, err := cl.CreateNexusEndpoint(ctx, &operatorservice.CreateNexusEndpointRequest{
			Spec: &nexuspb.EndpointSpec{
				Name: endpointName,
				Target: &nexuspb.EndpointTarget{
					Variant: &nexuspb.EndpointTarget_External_{
						External: &nexuspb.EndpointTarget_External{
							Url: "http://" + listenAddr,
						},
					},
				},
			},
		})
		s.NoError(err)
	}

	sdkClient1, err := sdkclient.Dial(sdkclient.Options{
		HostPort:  s.cluster1.Host().FrontendGRPCAddress(),
		Namespace: ns,
	})
	s.NoError(err)
	sdkClient2, err := sdkclient.Dial(sdkclient.Options{
		HostPort:  s.cluster2.Host().FrontendGRPCAddress(),
		Namespace: ns,
	})
	s.NoError(err)

	run, err := sdkClient1.ExecuteWorkflow(ctx, sdkclient.StartWorkflowOptions{
		TaskQueue: "tq",
		ID:        "test",
	}, "workflow")
	s.NoError(err)

	pollRes := s.pollWorkflowTask(ctx, s.cluster1.FrontendClient(), ns)
	_, err = s.cluster1.FrontendClient().RespondWorkflowTaskCompleted(ctx, &workflowservice.RespondWorkflowTaskCompletedRequest{
		TaskToken: pollRes.TaskToken,
		Commands: []*commandpb.Command{
			{
				CommandType: enumspb.COMMAND_TYPE_SCHEDULE_NEXUS_OPERATION,
				Attributes: &commandpb.Command_ScheduleNexusOperationCommandAttributes{
					ScheduleNexusOperationCommandAttributes: &commandpb.ScheduleNexusOperationCommandAttributes{
						Endpoint:  endpointName,
						Service:   "service",
						Operation: "operation",
					},
				},
			},
		},
	})
	s.NoError(err)

	// Ensure the scheduled event is replicated.
	s.waitEvent(ctx, sdkClient2, run, enumspb.EVENT_TYPE_NEXUS_OPERATION_SCHEDULED)

	// Check operation state changes are replicated to cluster2.
	s.waitOperationRetry(ctx, sdkClient2, run)

	// Now failover, and let cluster2 be the active.
	s.failover(ns, s.clusterNames[1], 2, s.cluster1.FrontendClient())

	s.NoError(sdkClient2.SignalWorkflow(ctx, run.GetID(), run.GetRunID(), "dont-care", nil))

	pollRes = s.pollWorkflowTask(ctx, s.cluster2.FrontendClient(), ns)

	// Unblock nexus operation start after failover.
	failStartOperation.Store(false)

	s.Eventually(func() bool {
		describeRes, err := sdkClient2.DescribeWorkflowExecution(ctx, run.GetID(), run.GetRunID())
		s.NoError(err)
		s.Equal(1, len(describeRes.PendingNexusOperations))
		op := describeRes.PendingNexusOperations[0]
		return op.State == enumspb.PENDING_NEXUS_OPERATION_STATE_STARTED
	}, time.Second*20, time.Millisecond*100)

	_, err = s.cluster2.FrontendClient().RespondWorkflowTaskCompleted(ctx, &workflowservice.RespondWorkflowTaskCompletedRequest{
		TaskToken: pollRes.TaskToken,
		Commands:  []*commandpb.Command{}, // No need to generate other commands, this "workflow" just waits for the operation to complete.
	})
	s.NoError(err)

	// Poll in cluster2 (previously standby) and verify the operation was started.
	pollRes = s.pollWorkflowTask(ctx, s.cluster2.FrontendClient(), ns)
	_, err = s.cluster2.FrontendClient().RespondWorkflowTaskCompleted(ctx, &workflowservice.RespondWorkflowTaskCompletedRequest{
		TaskToken: pollRes.TaskToken,
		Commands:  []*commandpb.Command{}, // No need to generate other commands, this "workflow" just waits for the operation to complete.
	})
	s.NoError(err)
	idx := slices.IndexFunc(pollRes.History.Events, func(ev *historypb.HistoryEvent) bool {
		return ev.EventType == enumspb.EVENT_TYPE_NEXUS_OPERATION_STARTED
	})
	s.Greater(idx, -1)

	// Ensure the started event is replicated back to cluster1.
	s.waitEvent(ctx, sdkClient1, run, enumspb.EVENT_TYPE_NEXUS_OPERATION_STARTED)

	// Fail back to cluster1.
	s.failover(ns, s.clusterNames[0], 11, s.cluster2.FrontendClient())

	s.completeNexusOperation(ctx, "result", publicCallbackUrl, callbackToken)

	// Verify completion triggers a new workflow task and that the workflow completes.
	pollRes = s.pollWorkflowTask(ctx, s.cluster1.FrontendClient(), ns)
	_, err = s.cluster1.FrontendClient().RespondWorkflowTaskCompleted(ctx, &workflowservice.RespondWorkflowTaskCompletedRequest{
		TaskToken: pollRes.TaskToken,
		Commands: []*commandpb.Command{
			{
				CommandType: enumspb.COMMAND_TYPE_COMPLETE_WORKFLOW_EXECUTION,
				Attributes: &commandpb.Command_CompleteWorkflowExecutionCommandAttributes{
					CompleteWorkflowExecutionCommandAttributes: &commandpb.CompleteWorkflowExecutionCommandAttributes{},
				},
			},
		},
	})
	s.NoError(err)
	idx = slices.IndexFunc(pollRes.History.Events, func(ev *historypb.HistoryEvent) bool {
		return ev.EventType == enumspb.EVENT_TYPE_NEXUS_OPERATION_COMPLETED
	})
	s.Greater(idx, -1)

	s.NoError(run.Get(ctx, nil))
}

func (s *NexusStateReplicationSuite) TestNexusOperationCancelationReplicated() {
	var callbackToken string
	var publicCallbackUrl string
	h := nexustest.Handler{
		OnStartOperation: func(ctx context.Context, service, operation string, input *nexus.LazyValue, options nexus.StartOperationOptions) (nexus.HandlerStartOperationResult[any], error) {
			callbackToken = options.CallbackHeader.Get(commonnexus.CallbackTokenHeader)
			publicCallbackUrl = options.CallbackURL
			return &nexus.HandlerStartOperationResultAsync{OperationID: "test"}, nil
		},
		OnCancelOperation: func(ctx context.Context, service, operation, operationID string, options nexus.CancelOperationOptions) error {
			return nil
		},
	}
	listenAddr := nexustest.AllocListenAddress()
	nexustest.NewNexusServer(s.T(), listenAddr, h)

	ctx := testcore.NewContext()
	ns := s.createGlobalNamespace()
	endpointName := testcore.RandomizedNexusEndpoint(s.T().Name())

	// Set URL template after httpAPAddress is set, see commonnexus.RouteCompletionCallback.
	// We don't actually want to deliver callbacks in this test, the config just has to be set for nexus task execution.
	for _, cluster := range []*testcore.TestCluster{s.cluster1, s.cluster2} {
		cluster.OverrideDynamicConfig(
			s.T(),
			nexusoperations.CallbackURLTemplate,
			"http://"+s.cluster1.Host().FrontendHTTPAddress()+"/namespaces/{{.NamespaceName}}/nexus/callback")
	}

	// Nexus endpoints registry isn't replicated yet, manually create the same endpoint in both clusters.
	for _, cl := range []operatorservice.OperatorServiceClient{s.cluster1.OperatorClient(), s.cluster2.OperatorClient()} {
		_, err := cl.CreateNexusEndpoint(ctx, &operatorservice.CreateNexusEndpointRequest{
			Spec: &nexuspb.EndpointSpec{
				Name: endpointName,
				Target: &nexuspb.EndpointTarget{
					Variant: &nexuspb.EndpointTarget_External_{
						External: &nexuspb.EndpointTarget_External{
							Url: "http://" + listenAddr,
						},
					},
				},
			},
		})
		s.NoError(err)
	}

	sdkClient1, err := sdkclient.Dial(sdkclient.Options{
		HostPort:  s.cluster1.Host().FrontendGRPCAddress(),
		Namespace: ns,
	})
	s.NoError(err)
	sdkClient2, err := sdkclient.Dial(sdkclient.Options{
		HostPort:  s.cluster2.Host().FrontendGRPCAddress(),
		Namespace: ns,
	})
	s.NoError(err)

	run, err := sdkClient1.ExecuteWorkflow(ctx, sdkclient.StartWorkflowOptions{
		TaskQueue: "tq",
		ID:        "test",
	}, "workflow")
	s.NoError(err)

	pollRes := s.pollWorkflowTask(ctx, s.cluster1.FrontendClient(), ns)
	_, err = s.cluster1.FrontendClient().RespondWorkflowTaskCompleted(ctx, &workflowservice.RespondWorkflowTaskCompletedRequest{
		TaskToken: pollRes.TaskToken,
		Commands: []*commandpb.Command{
			{
				CommandType: enumspb.COMMAND_TYPE_SCHEDULE_NEXUS_OPERATION,
				Attributes: &commandpb.Command_ScheduleNexusOperationCommandAttributes{
					ScheduleNexusOperationCommandAttributes: &commandpb.ScheduleNexusOperationCommandAttributes{
						Endpoint:  endpointName,
						Service:   "service",
						Operation: "operation",
					},
				},
			},
		},
	})
	s.NoError(err)

	// Ensure the scheduled event is replicated.
	scheduledEventID := s.waitEvent(ctx, sdkClient2, run, enumspb.EVENT_TYPE_NEXUS_OPERATION_SCHEDULED)

	// Verify the operation started and it is replicated in the passive cluster
	s.waitEvent(ctx, sdkClient2, run, enumspb.EVENT_TYPE_NEXUS_OPERATION_STARTED)

	pollRes = s.pollWorkflowTask(ctx, s.cluster1.FrontendClient(), ns)
	_, err = s.cluster1.FrontendClient().RespondWorkflowTaskCompleted(ctx, &workflowservice.RespondWorkflowTaskCompletedRequest{
		TaskToken: pollRes.TaskToken,
		Commands: []*commandpb.Command{
			{
				CommandType: enumspb.COMMAND_TYPE_REQUEST_CANCEL_NEXUS_OPERATION,
				Attributes: &commandpb.Command_RequestCancelNexusOperationCommandAttributes{
					RequestCancelNexusOperationCommandAttributes: &commandpb.RequestCancelNexusOperationCommandAttributes{
						ScheduledEventId: scheduledEventID,
					},
				},
			},
		},
	})
	s.NoError(err)

	s.Eventually(func() bool {
		describeRes, err := sdkClient1.DescribeWorkflowExecution(ctx, run.GetID(), run.GetRunID())
		s.NoError(err)
		s.Equal(1, len(describeRes.PendingNexusOperations))
		op := describeRes.PendingNexusOperations[0]
		fmt.Println(op.CancellationInfo)
		s.NotNil(op.CancellationInfo)
		return op.CancellationInfo.State == enumspb.NEXUS_OPERATION_CANCELLATION_STATE_SUCCEEDED
	}, time.Second*20, time.Millisecond*100)

	s.cancelNexusOperation(ctx, publicCallbackUrl, callbackToken)

	// Verify the canceled event is replicated and the passive cluster catches up.
	s.waitEvent(ctx, sdkClient2, run, enumspb.EVENT_TYPE_NEXUS_OPERATION_CANCELED)

	pollRes = s.pollWorkflowTask(ctx, s.cluster1.FrontendClient(), ns)
	_, err = s.cluster1.FrontendClient().RespondWorkflowTaskCompleted(ctx, &workflowservice.RespondWorkflowTaskCompletedRequest{
		TaskToken: pollRes.TaskToken,
		Commands: []*commandpb.Command{
			{
				CommandType: enumspb.COMMAND_TYPE_COMPLETE_WORKFLOW_EXECUTION,
				Attributes: &commandpb.Command_CompleteWorkflowExecutionCommandAttributes{
					CompleteWorkflowExecutionCommandAttributes: &commandpb.CompleteWorkflowExecutionCommandAttributes{},
				},
			},
		},
	})
	s.NoError(err)

	s.NoError(run.Get(ctx, nil))
}

// TestNexusCallbackReplicated tests that nexus callback and state updates are replicated
// across clusters and callback can work as expected when failover happens.
// General outline:
// 1. Start two clusters, cluster1 set to active, cluster2 set to standby.
// 2. Start a workflow will callback on cluster1.
// 3. Terminate the workflow to trigger the callback. An error is injected to always fail the callback.
// 4. Check the callback state changes are replicated to cluster2.
// 5. Failover to cluster2 and unblock the callback by removing the injected error.
// 6. Wait for the callback to complete on both clusters.
func (s *NexusStateReplicationSuite) TestNexusCallbackReplicated() {
	failCallback := atomic.Bool{}
	failCallback.Store(true)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if failCallback.Load() {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Injected error to trigger callback retry")
			return
		}
		fmt.Fprintln(w, "Callback succeeded")
	}))
	defer ts.Close()

	ctx := testcore.NewContext()
	ns := s.createGlobalNamespace()

	sdkClient1, err := sdkclient.Dial(sdkclient.Options{
		HostPort:  s.cluster1.Host().FrontendGRPCAddress(),
		Namespace: ns,
	})
	s.NoError(err)
	sdkClient2, err := sdkclient.Dial(sdkclient.Options{
		HostPort:  s.cluster2.Host().FrontendGRPCAddress(),
		Namespace: ns,
	})
	s.NoError(err)

	tv := testvars.New(s.T())
	startResp, err := sdkClient1.WorkflowService().StartWorkflowExecution(ctx, &workflowservice.StartWorkflowExecutionRequest{
		Namespace:    ns,
		WorkflowId:   tv.WorkflowID(),
		WorkflowType: tv.WorkflowType(),
		TaskQueue:    tv.TaskQueue(),
		RequestId:    uuid.New(),
		CompletionCallbacks: []*commonpb.Callback{
			{
				Variant: &commonpb.Callback_Nexus_{
					Nexus: &commonpb.Callback_Nexus{
						Url: ts.URL,
					},
				},
			},
		},
	})
	s.NoError(err)

	// Terminate the workflow to trigger the callback.
	err = sdkClient1.TerminateWorkflow(ctx, tv.WorkflowID(), startResp.RunId, "terminate workflow to trigger callback")
	s.NoError(err)

	// Check callback state changes are replicated to cluster2.
	s.waitCallback(ctx, sdkClient2, &commonpb.WorkflowExecution{
		WorkflowId: tv.WorkflowID(),
		RunId:      startResp.GetRunId(),
	}, func(callback *workflowpb.CallbackInfo) bool {
		return callback.Attempt > 2
	})

	// Failover to cluster2.
	s.failover(ns, s.clusterNames[1], 2, s.cluster1.FrontendClient())

	// Unblock callback after failover.
	failCallback.Store(false)

	// Check callback can complete on cluster2 after failover,
	// and succeeded state will be replicated back to cluster1.
	for _, sdkClient := range []sdkclient.Client{sdkClient1, sdkClient2} {
		s.waitCallback(ctx, sdkClient, &commonpb.WorkflowExecution{
			WorkflowId: tv.WorkflowID(),
			RunId:      startResp.GetRunId(),
		}, func(callback *workflowpb.CallbackInfo) bool {
			return callback.State == enumspb.CALLBACK_STATE_SUCCEEDED
		})
	}
}

func (s *NexusStateReplicationSuite) TestNexusOperationBufferedCompletionReplicated() {
	ctx := testcore.NewContext()
	ns := s.createGlobalNamespace()
	taskQueue := "tq"

	allowCompletion := atomic.Bool{}
	attemptCount := atomic.Int32{}

	h := nexustest.Handler{
		OnStartOperation: func(
			ctx context.Context,
			service, operation string,
			input *nexus.LazyValue,
			options nexus.StartOperationOptions,
		) (nexus.HandlerStartOperationResult[any], error) {
			attemptCount.Add(1)
			if !allowCompletion.Load() {
				return nil, nexus.HandlerErrorf(nexus.HandlerErrorTypeInternal, "injected error to trigger operation retry")
			}

			return &nexus.HandlerStartOperationResultSync[any]{
				Value: "",
			}, nil
		},
	}

	listenAddr := nexustest.AllocListenAddress()
	nexustest.NewNexusServer(s.T(), listenAddr, h)

	for _, cluster := range []*testcore.TestCluster{s.cluster1, s.cluster2} {
		cluster.OverrideDynamicConfig(
			s.T(),
			nexusoperations.CallbackURLTemplate,
			"http://"+s.cluster1.Host().FrontendHTTPAddress()+"/namespaces/{{.NamespaceName}}/nexus/callback",
		)
	}

	endpointName := testcore.RandomizedNexusEndpoint(s.T().Name())
	for _, cl := range []operatorservice.OperatorServiceClient{s.cluster1.OperatorClient(), s.cluster2.OperatorClient()} {
		_, err := cl.CreateNexusEndpoint(ctx, &operatorservice.CreateNexusEndpointRequest{
			Spec: &nexuspb.EndpointSpec{
				Name: endpointName,
				Target: &nexuspb.EndpointTarget{
					Variant: &nexuspb.EndpointTarget_External_{
						External: &nexuspb.EndpointTarget_External{
							Url: "http://" + listenAddr,
						},
					},
				},
			},
		})
		s.NoError(err)
	}

	sdkClient1, err := sdkclient.Dial(sdkclient.Options{
		HostPort:  s.cluster1.Host().FrontendGRPCAddress(),
		Namespace: ns,
	})
	s.NoError(err)
	sdkClient2, err := sdkclient.Dial(sdkclient.Options{
		HostPort:  s.cluster2.Host().FrontendGRPCAddress(),
		Namespace: ns,
	})
	s.NoError(err)

	run, err := sdkClient1.ExecuteWorkflow(ctx, sdkclient.StartWorkflowOptions{
		TaskQueue: taskQueue,
	}, "workflow")
	s.NoError(err)

	pollResp := s.pollWorkflowTask(ctx, s.cluster1.FrontendClient(), ns)

	// Schedule operation and start timer to force next WFT
	_, err = s.cluster1.FrontendClient().RespondWorkflowTaskCompleted(ctx, &workflowservice.RespondWorkflowTaskCompletedRequest{
		Identity:  "test",
		TaskToken: pollResp.TaskToken,
		Commands: []*commandpb.Command{
			{
				CommandType: enumspb.COMMAND_TYPE_SCHEDULE_NEXUS_OPERATION,
				Attributes: &commandpb.Command_ScheduleNexusOperationCommandAttributes{
					ScheduleNexusOperationCommandAttributes: &commandpb.ScheduleNexusOperationCommandAttributes{
						Endpoint:  endpointName,
						Service:   "service",
						Operation: "operation",
					},
				},
			},
			{
				CommandType: enumspb.COMMAND_TYPE_START_TIMER,
				Attributes: &commandpb.Command_StartTimerCommandAttributes{
					StartTimerCommandAttributes: &commandpb.StartTimerCommandAttributes{
						TimerId:            "timer-1",
						StartToFireTimeout: durationpb.New(1 * time.Second),
					},
				},
			},
		},
	})
	s.NoError(err)

	// Poll next WFT which will be scheduled when timer fires
	secondPollResp, err := s.cluster1.FrontendClient().PollWorkflowTaskQueue(ctx, &workflowservice.PollWorkflowTaskQueueRequest{
		Namespace: ns,
		TaskQueue: &taskqueuepb.TaskQueue{Name: taskQueue, Kind: enumspb.TASK_QUEUE_KIND_NORMAL},
		Identity:  "test",
	})
	s.NoError(err)
	s.NotNil(secondPollResp)

	// Find the scheduled event ID from history
	var scheduledEventID int64
	for _, e := range secondPollResp.History.Events {
		if e.GetEventType() == enumspb.EVENT_TYPE_NEXUS_OPERATION_SCHEDULED {
			scheduledEventID = e.GetEventId()
			break
		}
	}
	s.Greater(scheduledEventID, int64(0))

	// Allow operation to complete synchronously during next retry attempt
	allowCompletion.Store(true)

	// Wait for operation completion to be recorded
	s.Eventually(func() bool {
		desc, err := sdkClient1.DescribeWorkflowExecution(ctx, run.GetID(), run.GetRunID())
		s.NoError(err)
		return len(desc.PendingNexusOperations) == 0
	}, 10*time.Second, 200*time.Millisecond)

	// Try to cancel operation - should succeed since completion is buffered
	_, err = s.cluster1.FrontendClient().RespondWorkflowTaskCompleted(ctx, &workflowservice.RespondWorkflowTaskCompletedRequest{
		Identity:  "test",
		TaskToken: secondPollResp.TaskToken,
		Commands: []*commandpb.Command{
			{
				CommandType: enumspb.COMMAND_TYPE_REQUEST_CANCEL_NEXUS_OPERATION,
				Attributes: &commandpb.Command_RequestCancelNexusOperationCommandAttributes{
					RequestCancelNexusOperationCommandAttributes: &commandpb.RequestCancelNexusOperationCommandAttributes{
						ScheduledEventId: scheduledEventID,
					},
				},
			},
		},
	})
	s.NoError(err, "Cancel request should be accepted when operation has buffered completion")

	// Ensure no pending operations in passive cluster state
	s.Eventually(func() bool {
		desc, err := sdkClient2.DescribeWorkflowExecution(ctx, run.GetID(), run.GetRunID())
		s.NoError(err)
		return len(desc.PendingNexusOperations) == 0
	}, 10*time.Second, 200*time.Millisecond)

	finalPollResp := s.pollWorkflowTask(ctx, s.cluster1.FrontendClient(), ns)
	_, err = s.cluster1.FrontendClient().RespondWorkflowTaskCompleted(ctx, &workflowservice.RespondWorkflowTaskCompletedRequest{
		Identity:  "test",
		TaskToken: finalPollResp.TaskToken,
		Commands: []*commandpb.Command{
			{
				CommandType: enumspb.COMMAND_TYPE_COMPLETE_WORKFLOW_EXECUTION,
				Attributes: &commandpb.Command_CompleteWorkflowExecutionCommandAttributes{
					CompleteWorkflowExecutionCommandAttributes: &commandpb.CompleteWorkflowExecutionCommandAttributes{},
				},
			},
		},
	})
	s.NoError(err)

	// Verify history in both clusters
	for _, sdkClient := range []sdkclient.Client{sdkClient1, sdkClient2} {
		historyIterator := sdkClient.GetWorkflowHistory(ctx, run.GetID(), run.GetRunID(), false, enumspb.HISTORY_EVENT_FILTER_TYPE_ALL_EVENT)
		var events []*historypb.HistoryEvent
		for historyIterator.HasNext() {
			e, err := historyIterator.Next()
			s.NoError(err)
			events = append(events, e)
		}
		s.ContainsHistoryEvents(`
NexusOperationCancelRequested
NexusOperationCompleted
`, events)
	}
}

func (s *NexusStateReplicationSuite) waitEvent(ctx context.Context, sdkClient sdkclient.Client, run sdkclient.WorkflowRun, eventType enumspb.EventType) (eventID int64) {
	s.Eventually(func() bool {
		history := sdkClient.GetWorkflowHistory(ctx, run.GetID(), run.GetRunID(), false, enumspb.HISTORY_EVENT_FILTER_TYPE_ALL_EVENT)
		for history.HasNext() {
			event, err := history.Next()
			s.NoError(err)
			if event.EventType == eventType {
				eventID = event.EventId
				return true
			}
		}
		return false
	}, time.Second*10, time.Millisecond*100)
	return
}

func (s *NexusStateReplicationSuite) waitOperationRetry(
	ctx context.Context,
	sdkClient sdkclient.Client,
	run sdkclient.WorkflowRun,
) {
	s.Eventually(func() bool {
		descResp, err := sdkClient.DescribeWorkflowExecution(ctx, run.GetID(), run.GetRunID())
		s.NoError(err)
		s.Len(descResp.GetPendingNexusOperations(), 1)
		return descResp.GetPendingNexusOperations()[0].Attempt > 2
	}, time.Second*10, time.Millisecond*100)
}

func (s *NexusStateReplicationSuite) pollWorkflowTask(ctx context.Context, client workflowservice.WorkflowServiceClient, ns string) *workflowservice.PollWorkflowTaskQueueResponse {
	pollRes, err := client.PollWorkflowTaskQueue(ctx, &workflowservice.PollWorkflowTaskQueueRequest{
		Namespace: ns,
		TaskQueue: &taskqueuepb.TaskQueue{
			Name: "tq",
			Kind: enumspb.TASK_QUEUE_KIND_NORMAL,
		},
		Identity: "test",
	})
	s.NoError(err)
	return pollRes
}

func (s *NexusStateReplicationSuite) waitCallback(
	ctx context.Context,
	sdkClient sdkclient.Client,
	execution *commonpb.WorkflowExecution,
	condition func(callback *workflowpb.CallbackInfo) bool,
) {
	s.Eventually(func() bool {
		descResp, err := sdkClient.DescribeWorkflowExecution(ctx, execution.WorkflowId, execution.RunId)
		s.NoError(err)
		s.Len(descResp.GetCallbacks(), 1)
		return condition(descResp.GetCallbacks()[0])
	}, time.Second*20, time.Millisecond*100)
}

func (s *NexusStateReplicationSuite) completeNexusOperation(ctx context.Context, result any, callbackUrl, callbackToken string) {
	completion, err := nexus.NewOperationCompletionSuccessful(s.mustToPayload(result), nexus.OperationCompletionSuccessfulOptions{
		Serializer: commonnexus.PayloadSerializer,
	})
	s.NoError(err)
	req, err := nexus.NewCompletionHTTPRequest(ctx, callbackUrl, completion)
	s.NoError(err)
	if callbackToken != "" {
		req.Header.Add(commonnexus.CallbackTokenHeader, callbackToken)
	}

	res, err := http.DefaultClient.Do(req)
	s.NoError(err)
	defer res.Body.Close()
	_, err = io.ReadAll(res.Body)
	s.NoError(err)
	s.Equal(http.StatusOK, res.StatusCode)
}

func (s *NexusStateReplicationSuite) cancelNexusOperation(ctx context.Context, callbackUrl, callbackToken string) {
	completion, err := nexus.NewOperationCompletionUnsuccessful(
		nexus.NewCanceledOperationError(errors.New("operation canceled")),
		nexus.OperationCompletionUnsuccessfulOptions{},
	)
	s.NoError(err)
	req, err := nexus.NewCompletionHTTPRequest(ctx, callbackUrl, completion)
	s.NoError(err)
	if callbackToken != "" {
		req.Header.Add(commonnexus.CallbackTokenHeader, callbackToken)
	}

	res, err := http.DefaultClient.Do(req)
	s.NoError(err)
	defer res.Body.Close()
	_, err = io.ReadAll(res.Body)
	s.NoError(err)
	s.Equal(http.StatusOK, res.StatusCode)
}