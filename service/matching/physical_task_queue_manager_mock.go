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

// Code generated by MockGen. DO NOT EDIT.
// Source: physical_task_queue_manager_interface.go
//
// Generated by this command:
//
//	mockgen -copyright_file ../../LICENSE -package matching -source physical_task_queue_manager_interface.go -destination physical_task_queue_manager_mock.go
//

// Package matching is a generated GoMock package.
package matching

import (
	context "context"
	reflect "reflect"
	time "time"

	taskqueue "go.temporal.io/api/taskqueue/v1"
	matchingservice "go.temporal.io/server/api/matchingservice/v1"
	persistence "go.temporal.io/server/api/persistence/v1"
	taskqueue0 "go.temporal.io/server/api/taskqueue/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockphysicalTaskQueueManager is a mock of physicalTaskQueueManager interface.
type MockphysicalTaskQueueManager struct {
	ctrl     *gomock.Controller
	recorder *MockphysicalTaskQueueManagerMockRecorder
	isgomock struct{}
}

// MockphysicalTaskQueueManagerMockRecorder is the mock recorder for MockphysicalTaskQueueManager.
type MockphysicalTaskQueueManagerMockRecorder struct {
	mock *MockphysicalTaskQueueManager
}

// NewMockphysicalTaskQueueManager creates a new mock instance.
func NewMockphysicalTaskQueueManager(ctrl *gomock.Controller) *MockphysicalTaskQueueManager {
	mock := &MockphysicalTaskQueueManager{ctrl: ctrl}
	mock.recorder = &MockphysicalTaskQueueManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockphysicalTaskQueueManager) EXPECT() *MockphysicalTaskQueueManagerMockRecorder {
	return m.recorder
}

// AddSpooledTask mocks base method.
func (m *MockphysicalTaskQueueManager) AddSpooledTask(task *internalTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSpooledTask", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSpooledTask indicates an expected call of AddSpooledTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) AddSpooledTask(task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSpooledTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).AddSpooledTask), task)
}

// AddSpooledTaskToMatcher mocks base method.
func (m *MockphysicalTaskQueueManager) AddSpooledTaskToMatcher(task *internalTask) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSpooledTaskToMatcher", task)
}

// AddSpooledTaskToMatcher indicates an expected call of AddSpooledTaskToMatcher.
func (mr *MockphysicalTaskQueueManagerMockRecorder) AddSpooledTaskToMatcher(task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSpooledTaskToMatcher", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).AddSpooledTaskToMatcher), task)
}

// DispatchNexusTask mocks base method.
func (m *MockphysicalTaskQueueManager) DispatchNexusTask(ctx context.Context, taskId string, request *matchingservice.DispatchNexusTaskRequest) (*matchingservice.DispatchNexusTaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchNexusTask", ctx, taskId, request)
	ret0, _ := ret[0].(*matchingservice.DispatchNexusTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DispatchNexusTask indicates an expected call of DispatchNexusTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) DispatchNexusTask(ctx, taskId, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchNexusTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).DispatchNexusTask), ctx, taskId, request)
}

// DispatchQueryTask mocks base method.
func (m *MockphysicalTaskQueueManager) DispatchQueryTask(ctx context.Context, taskId string, request *matchingservice.QueryWorkflowRequest) (*matchingservice.QueryWorkflowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchQueryTask", ctx, taskId, request)
	ret0, _ := ret[0].(*matchingservice.QueryWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DispatchQueryTask indicates an expected call of DispatchQueryTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) DispatchQueryTask(ctx, taskId, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchQueryTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).DispatchQueryTask), ctx, taskId, request)
}

// DispatchSpooledTask mocks base method.
func (m *MockphysicalTaskQueueManager) DispatchSpooledTask(ctx context.Context, task *internalTask, userDataChanged <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchSpooledTask", ctx, task, userDataChanged)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchSpooledTask indicates an expected call of DispatchSpooledTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) DispatchSpooledTask(ctx, task, userDataChanged any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchSpooledTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).DispatchSpooledTask), ctx, task, userDataChanged)
}

// GetAllPollerInfo mocks base method.
func (m *MockphysicalTaskQueueManager) GetAllPollerInfo() []*taskqueue.PollerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPollerInfo")
	ret0, _ := ret[0].([]*taskqueue.PollerInfo)
	return ret0
}

// GetAllPollerInfo indicates an expected call of GetAllPollerInfo.
func (mr *MockphysicalTaskQueueManagerMockRecorder) GetAllPollerInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPollerInfo", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).GetAllPollerInfo))
}

// GetInternalTaskQueueStatus mocks base method.
func (m *MockphysicalTaskQueueManager) GetInternalTaskQueueStatus() []*taskqueue0.InternalTaskQueueStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInternalTaskQueueStatus")
	ret0, _ := ret[0].([]*taskqueue0.InternalTaskQueueStatus)
	return ret0
}

// GetInternalTaskQueueStatus indicates an expected call of GetInternalTaskQueueStatus.
func (mr *MockphysicalTaskQueueManagerMockRecorder) GetInternalTaskQueueStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInternalTaskQueueStatus", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).GetInternalTaskQueueStatus))
}

// GetStats mocks base method.
func (m *MockphysicalTaskQueueManager) GetStats() *taskqueue.TaskQueueStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats")
	ret0, _ := ret[0].(*taskqueue.TaskQueueStats)
	return ret0
}

// GetStats indicates an expected call of GetStats.
func (mr *MockphysicalTaskQueueManagerMockRecorder) GetStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).GetStats))
}

// HasPollerAfter mocks base method.
func (m *MockphysicalTaskQueueManager) HasPollerAfter(accessTime time.Time) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPollerAfter", accessTime)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPollerAfter indicates an expected call of HasPollerAfter.
func (mr *MockphysicalTaskQueueManagerMockRecorder) HasPollerAfter(accessTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPollerAfter", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).HasPollerAfter), accessTime)
}

// LegacyDescribeTaskQueue mocks base method.
func (m *MockphysicalTaskQueueManager) LegacyDescribeTaskQueue(includeTaskQueueStatus bool) *matchingservice.DescribeTaskQueueResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LegacyDescribeTaskQueue", includeTaskQueueStatus)
	ret0, _ := ret[0].(*matchingservice.DescribeTaskQueueResponse)
	return ret0
}

// LegacyDescribeTaskQueue indicates an expected call of LegacyDescribeTaskQueue.
func (mr *MockphysicalTaskQueueManagerMockRecorder) LegacyDescribeTaskQueue(includeTaskQueueStatus any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LegacyDescribeTaskQueue", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).LegacyDescribeTaskQueue), includeTaskQueueStatus)
}

// MakePollerScalingDecision mocks base method.
func (m *MockphysicalTaskQueueManager) MakePollerScalingDecision(pollStartTime time.Time) *taskqueue.PollerScalingDecision {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakePollerScalingDecision", pollStartTime)
	ret0, _ := ret[0].(*taskqueue.PollerScalingDecision)
	return ret0
}

// MakePollerScalingDecision indicates an expected call of MakePollerScalingDecision.
func (mr *MockphysicalTaskQueueManagerMockRecorder) MakePollerScalingDecision(pollStartTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakePollerScalingDecision", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).MakePollerScalingDecision), pollStartTime)
}

// MarkAlive mocks base method.
func (m *MockphysicalTaskQueueManager) MarkAlive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkAlive")
}

// MarkAlive indicates an expected call of MarkAlive.
func (mr *MockphysicalTaskQueueManagerMockRecorder) MarkAlive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAlive", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).MarkAlive))
}

// PollTask mocks base method.
func (m *MockphysicalTaskQueueManager) PollTask(ctx context.Context, pollMetadata *pollMetadata) (*internalTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollTask", ctx, pollMetadata)
	ret0, _ := ret[0].(*internalTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollTask indicates an expected call of PollTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) PollTask(ctx, pollMetadata any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).PollTask), ctx, pollMetadata)
}

// ProcessSpooledTask mocks base method.
func (m *MockphysicalTaskQueueManager) ProcessSpooledTask(ctx context.Context, task *internalTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessSpooledTask", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessSpooledTask indicates an expected call of ProcessSpooledTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) ProcessSpooledTask(ctx, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSpooledTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).ProcessSpooledTask), ctx, task)
}

// QueueKey mocks base method.
func (m *MockphysicalTaskQueueManager) QueueKey() *PhysicalTaskQueueKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueKey")
	ret0, _ := ret[0].(*PhysicalTaskQueueKey)
	return ret0
}

// QueueKey indicates an expected call of QueueKey.
func (mr *MockphysicalTaskQueueManagerMockRecorder) QueueKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueKey", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).QueueKey))
}

// SpoolTask mocks base method.
func (m *MockphysicalTaskQueueManager) SpoolTask(taskInfo *persistence.TaskInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpoolTask", taskInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// SpoolTask indicates an expected call of SpoolTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) SpoolTask(taskInfo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpoolTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).SpoolTask), taskInfo)
}

// Start mocks base method.
func (m *MockphysicalTaskQueueManager) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockphysicalTaskQueueManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).Start))
}

// Stop mocks base method.
func (m *MockphysicalTaskQueueManager) Stop(arg0 unloadCause) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", arg0)
}

// Stop indicates an expected call of Stop.
func (mr *MockphysicalTaskQueueManagerMockRecorder) Stop(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).Stop), arg0)
}

// TrySyncMatch mocks base method.
func (m *MockphysicalTaskQueueManager) TrySyncMatch(ctx context.Context, task *internalTask) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrySyncMatch", ctx, task)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrySyncMatch indicates an expected call of TrySyncMatch.
func (mr *MockphysicalTaskQueueManagerMockRecorder) TrySyncMatch(ctx, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrySyncMatch", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).TrySyncMatch), ctx, task)
}

// UnloadFromPartitionManager mocks base method.
func (m *MockphysicalTaskQueueManager) UnloadFromPartitionManager(arg0 unloadCause) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnloadFromPartitionManager", arg0)
}

// UnloadFromPartitionManager indicates an expected call of UnloadFromPartitionManager.
func (mr *MockphysicalTaskQueueManagerMockRecorder) UnloadFromPartitionManager(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnloadFromPartitionManager", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).UnloadFromPartitionManager), arg0)
}

// UpdatePollerInfo mocks base method.
func (m *MockphysicalTaskQueueManager) UpdatePollerInfo(arg0 pollerIdentity, arg1 *pollMetadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePollerInfo", arg0, arg1)
}

// UpdatePollerInfo indicates an expected call of UpdatePollerInfo.
func (mr *MockphysicalTaskQueueManagerMockRecorder) UpdatePollerInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePollerInfo", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).UpdatePollerInfo), arg0, arg1)
}

// UserDataChanged mocks base method.
func (m *MockphysicalTaskQueueManager) UserDataChanged() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UserDataChanged")
}

// UserDataChanged indicates an expected call of UserDataChanged.
func (mr *MockphysicalTaskQueueManagerMockRecorder) UserDataChanged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDataChanged", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).UserDataChanged))
}

// WaitUntilInitialized mocks base method.
func (m *MockphysicalTaskQueueManager) WaitUntilInitialized(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilInitialized", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilInitialized indicates an expected call of WaitUntilInitialized.
func (mr *MockphysicalTaskQueueManagerMockRecorder) WaitUntilInitialized(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilInitialized", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).WaitUntilInitialized), arg0)
}
