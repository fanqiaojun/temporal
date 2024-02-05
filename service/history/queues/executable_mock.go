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
// Source: executable.go

// Package queues is a generated GoMock package.
package queues

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1 "go.temporal.io/server/api/enums/v1"
	backoff "go.temporal.io/server/common/backoff"
	tasks "go.temporal.io/server/common/tasks"
	tasks0 "go.temporal.io/server/service/history/tasks"
)

// MockExecutable is a mock of Executable interface.
type MockExecutable struct {
	ctrl     *gomock.Controller
	recorder *MockExecutableMockRecorder
}

// MockExecutableMockRecorder is the mock recorder for MockExecutable.
type MockExecutableMockRecorder struct {
	mock *MockExecutable
}

// NewMockExecutable creates a new mock instance.
func NewMockExecutable(ctrl *gomock.Controller) *MockExecutable {
	mock := &MockExecutable{ctrl: ctrl}
	mock.recorder = &MockExecutableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutable) EXPECT() *MockExecutableMockRecorder {
	return m.recorder
}

// Abort mocks base method.
func (m *MockExecutable) Abort() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Abort")
}

// Abort indicates an expected call of Abort.
func (mr *MockExecutableMockRecorder) Abort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockExecutable)(nil).Abort))
}

// Ack mocks base method.
func (m *MockExecutable) Ack() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ack")
}

// Ack indicates an expected call of Ack.
func (mr *MockExecutableMockRecorder) Ack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockExecutable)(nil).Ack))
}

// Attempt mocks base method.
func (m *MockExecutable) Attempt() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attempt")
	ret0, _ := ret[0].(int)
	return ret0
}

// Attempt indicates an expected call of Attempt.
func (mr *MockExecutableMockRecorder) Attempt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attempt", reflect.TypeOf((*MockExecutable)(nil).Attempt))
}

// Cancel mocks base method.
func (m *MockExecutable) Cancel() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cancel")
}

// Cancel indicates an expected call of Cancel.
func (mr *MockExecutableMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockExecutable)(nil).Cancel))
}

// Execute mocks base method.
func (m *MockExecutable) Execute() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute")
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockExecutableMockRecorder) Execute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockExecutable)(nil).Execute))
}

// GetCategory mocks base method.
func (m *MockExecutable) GetCategory() tasks0.Category {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory")
	ret0, _ := ret[0].(tasks0.Category)
	return ret0
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockExecutableMockRecorder) GetCategory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockExecutable)(nil).GetCategory))
}

// GetKey mocks base method.
func (m *MockExecutable) GetKey() tasks0.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey")
	ret0, _ := ret[0].(tasks0.Key)
	return ret0
}

// GetKey indicates an expected call of GetKey.
func (mr *MockExecutableMockRecorder) GetKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockExecutable)(nil).GetKey))
}

// GetNamespaceID mocks base method.
func (m *MockExecutable) GetNamespaceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespaceID indicates an expected call of GetNamespaceID.
func (mr *MockExecutableMockRecorder) GetNamespaceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceID", reflect.TypeOf((*MockExecutable)(nil).GetNamespaceID))
}

// GetPriority mocks base method.
func (m *MockExecutable) GetPriority() tasks.Priority {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPriority")
	ret0, _ := ret[0].(tasks.Priority)
	return ret0
}

// GetPriority indicates an expected call of GetPriority.
func (mr *MockExecutableMockRecorder) GetPriority() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPriority", reflect.TypeOf((*MockExecutable)(nil).GetPriority))
}

// GetRunID mocks base method.
func (m *MockExecutable) GetRunID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRunID indicates an expected call of GetRunID.
func (mr *MockExecutableMockRecorder) GetRunID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunID", reflect.TypeOf((*MockExecutable)(nil).GetRunID))
}

// GetScheduledTime mocks base method.
func (m *MockExecutable) GetScheduledTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScheduledTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetScheduledTime indicates an expected call of GetScheduledTime.
func (mr *MockExecutableMockRecorder) GetScheduledTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduledTime", reflect.TypeOf((*MockExecutable)(nil).GetScheduledTime))
}

// GetTask mocks base method.
func (m *MockExecutable) GetTask() tasks0.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask")
	ret0, _ := ret[0].(tasks0.Task)
	return ret0
}

// GetTask indicates an expected call of GetTask.
func (mr *MockExecutableMockRecorder) GetTask() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockExecutable)(nil).GetTask))
}

// GetTaskID mocks base method.
func (m *MockExecutable) GetTaskID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetTaskID indicates an expected call of GetTaskID.
func (mr *MockExecutableMockRecorder) GetTaskID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskID", reflect.TypeOf((*MockExecutable)(nil).GetTaskID))
}

// GetType mocks base method.
func (m *MockExecutable) GetType() v1.TaskType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(v1.TaskType)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockExecutableMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockExecutable)(nil).GetType))
}

// GetVersion mocks base method.
func (m *MockExecutable) GetVersion() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockExecutableMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockExecutable)(nil).GetVersion))
}

// GetVisibilityTime mocks base method.
func (m *MockExecutable) GetVisibilityTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVisibilityTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetVisibilityTime indicates an expected call of GetVisibilityTime.
func (mr *MockExecutableMockRecorder) GetVisibilityTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVisibilityTime", reflect.TypeOf((*MockExecutable)(nil).GetVisibilityTime))
}

// GetWorkflowID mocks base method.
func (m *MockExecutable) GetWorkflowID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWorkflowID indicates an expected call of GetWorkflowID.
func (mr *MockExecutableMockRecorder) GetWorkflowID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowID", reflect.TypeOf((*MockExecutable)(nil).GetWorkflowID))
}

// HandleErr mocks base method.
func (m *MockExecutable) HandleErr(err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleErr", err)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleErr indicates an expected call of HandleErr.
func (mr *MockExecutableMockRecorder) HandleErr(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleErr", reflect.TypeOf((*MockExecutable)(nil).HandleErr), err)
}

// IsRetryableError mocks base method.
func (m *MockExecutable) IsRetryableError(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRetryableError", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRetryableError indicates an expected call of IsRetryableError.
func (mr *MockExecutableMockRecorder) IsRetryableError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRetryableError", reflect.TypeOf((*MockExecutable)(nil).IsRetryableError), err)
}

// Nack mocks base method.
func (m *MockExecutable) Nack(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Nack", err)
}

// Nack indicates an expected call of Nack.
func (mr *MockExecutableMockRecorder) Nack(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nack", reflect.TypeOf((*MockExecutable)(nil).Nack), err)
}

// Reschedule mocks base method.
func (m *MockExecutable) Reschedule() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reschedule")
}

// Reschedule indicates an expected call of Reschedule.
func (mr *MockExecutableMockRecorder) Reschedule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reschedule", reflect.TypeOf((*MockExecutable)(nil).Reschedule))
}

// RetryPolicy mocks base method.
func (m *MockExecutable) RetryPolicy() backoff.RetryPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetryPolicy")
	ret0, _ := ret[0].(backoff.RetryPolicy)
	return ret0
}

// RetryPolicy indicates an expected call of RetryPolicy.
func (mr *MockExecutableMockRecorder) RetryPolicy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryPolicy", reflect.TypeOf((*MockExecutable)(nil).RetryPolicy))
}

// SetScheduledTime mocks base method.
func (m *MockExecutable) SetScheduledTime(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetScheduledTime", arg0)
}

// SetScheduledTime indicates an expected call of SetScheduledTime.
func (mr *MockExecutableMockRecorder) SetScheduledTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetScheduledTime", reflect.TypeOf((*MockExecutable)(nil).SetScheduledTime), arg0)
}

// SetTaskID mocks base method.
func (m *MockExecutable) SetTaskID(id int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTaskID", id)
}

// SetTaskID indicates an expected call of SetTaskID.
func (mr *MockExecutableMockRecorder) SetTaskID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTaskID", reflect.TypeOf((*MockExecutable)(nil).SetTaskID), id)
}

// SetVisibilityTime mocks base method.
func (m *MockExecutable) SetVisibilityTime(timestamp time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVisibilityTime", timestamp)
}

// SetVisibilityTime indicates an expected call of SetVisibilityTime.
func (mr *MockExecutableMockRecorder) SetVisibilityTime(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVisibilityTime", reflect.TypeOf((*MockExecutable)(nil).SetVisibilityTime), timestamp)
}

// State mocks base method.
func (m *MockExecutable) State() tasks.State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(tasks.State)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockExecutableMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockExecutable)(nil).State))
}

// MockExecutor is a mock of Executor interface.
type MockExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorMockRecorder
}

// MockExecutorMockRecorder is the mock recorder for MockExecutor.
type MockExecutorMockRecorder struct {
	mock *MockExecutor
}

// NewMockExecutor creates a new mock instance.
func NewMockExecutor(ctrl *gomock.Controller) *MockExecutor {
	mock := &MockExecutor{ctrl: ctrl}
	mock.recorder = &MockExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutor) EXPECT() *MockExecutorMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockExecutor) Execute(arg0 context.Context, arg1 Executable) ExecuteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1)
	ret0, _ := ret[0].(ExecuteResponse)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockExecutorMockRecorder) Execute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockExecutor)(nil).Execute), arg0, arg1)
}

// MockExecutorWrapper is a mock of ExecutorWrapper interface.
type MockExecutorWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorWrapperMockRecorder
}

// MockExecutorWrapperMockRecorder is the mock recorder for MockExecutorWrapper.
type MockExecutorWrapperMockRecorder struct {
	mock *MockExecutorWrapper
}

// NewMockExecutorWrapper creates a new mock instance.
func NewMockExecutorWrapper(ctrl *gomock.Controller) *MockExecutorWrapper {
	mock := &MockExecutorWrapper{ctrl: ctrl}
	mock.recorder = &MockExecutorWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutorWrapper) EXPECT() *MockExecutorWrapperMockRecorder {
	return m.recorder
}

// Wrap mocks base method.
func (m *MockExecutorWrapper) Wrap(delegate Executor) Executor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wrap", delegate)
	ret0, _ := ret[0].(Executor)
	return ret0
}

// Wrap indicates an expected call of Wrap.
func (mr *MockExecutorWrapperMockRecorder) Wrap(delegate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wrap", reflect.TypeOf((*MockExecutorWrapper)(nil).Wrap), delegate)
}

// MockTerminalTaskError is a mock of TerminalTaskError interface.
type MockTerminalTaskError struct {
	ctrl     *gomock.Controller
	recorder *MockTerminalTaskErrorMockRecorder
}

// MockTerminalTaskErrorMockRecorder is the mock recorder for MockTerminalTaskError.
type MockTerminalTaskErrorMockRecorder struct {
	mock *MockTerminalTaskError
}

// NewMockTerminalTaskError creates a new mock instance.
func NewMockTerminalTaskError(ctrl *gomock.Controller) *MockTerminalTaskError {
	mock := &MockTerminalTaskError{ctrl: ctrl}
	mock.recorder = &MockTerminalTaskErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTerminalTaskError) EXPECT() *MockTerminalTaskErrorMockRecorder {
	return m.recorder
}

// IsTerminalTaskError mocks base method.
func (m *MockTerminalTaskError) IsTerminalTaskError() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTerminalTaskError")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTerminalTaskError indicates an expected call of IsTerminalTaskError.
func (mr *MockTerminalTaskErrorMockRecorder) IsTerminalTaskError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTerminalTaskError", reflect.TypeOf((*MockTerminalTaskError)(nil).IsTerminalTaskError))
}
