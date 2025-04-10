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
// Source: clientfactory.go
//
// Generated by this command:
//
//	mockgen -copyright_file ../LICENSE -package client -source clientfactory.go -destination client_factory_mock.go
//

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"
	time "time"

	workflowservice "go.temporal.io/api/workflowservice/v1"
	adminservice "go.temporal.io/server/api/adminservice/v1"
	historyservice "go.temporal.io/server/api/historyservice/v1"
	matchingservice "go.temporal.io/server/api/matchingservice/v1"
	common "go.temporal.io/server/common"
	dynamicconfig "go.temporal.io/server/common/dynamicconfig"
	log "go.temporal.io/server/common/log"
	membership "go.temporal.io/server/common/membership"
	metrics "go.temporal.io/server/common/metrics"
	testhooks "go.temporal.io/server/common/testing/testhooks"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
	isgomock struct{}
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// NewHistoryClientWithTimeout mocks base method.
func (m *MockFactory) NewHistoryClientWithTimeout(timeout time.Duration) (historyservice.HistoryServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewHistoryClientWithTimeout", timeout)
	ret0, _ := ret[0].(historyservice.HistoryServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewHistoryClientWithTimeout indicates an expected call of NewHistoryClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewHistoryClientWithTimeout(timeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHistoryClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewHistoryClientWithTimeout), timeout)
}

// NewLocalAdminClientWithTimeout mocks base method.
func (m *MockFactory) NewLocalAdminClientWithTimeout(timeout, largeTimeout time.Duration) (adminservice.AdminServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewLocalAdminClientWithTimeout", timeout, largeTimeout)
	ret0, _ := ret[0].(adminservice.AdminServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewLocalAdminClientWithTimeout indicates an expected call of NewLocalAdminClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewLocalAdminClientWithTimeout(timeout, largeTimeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLocalAdminClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewLocalAdminClientWithTimeout), timeout, largeTimeout)
}

// NewLocalFrontendClientWithTimeout mocks base method.
func (m *MockFactory) NewLocalFrontendClientWithTimeout(timeout, longPollTimeout time.Duration) (grpc.ClientConnInterface, workflowservice.WorkflowServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewLocalFrontendClientWithTimeout", timeout, longPollTimeout)
	ret0, _ := ret[0].(grpc.ClientConnInterface)
	ret1, _ := ret[1].(workflowservice.WorkflowServiceClient)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NewLocalFrontendClientWithTimeout indicates an expected call of NewLocalFrontendClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewLocalFrontendClientWithTimeout(timeout, longPollTimeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLocalFrontendClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewLocalFrontendClientWithTimeout), timeout, longPollTimeout)
}

// NewMatchingClientWithTimeout mocks base method.
func (m *MockFactory) NewMatchingClientWithTimeout(namespaceIDToName NamespaceIDToNameFunc, timeout, longPollTimeout time.Duration) (matchingservice.MatchingServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMatchingClientWithTimeout", namespaceIDToName, timeout, longPollTimeout)
	ret0, _ := ret[0].(matchingservice.MatchingServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewMatchingClientWithTimeout indicates an expected call of NewMatchingClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewMatchingClientWithTimeout(namespaceIDToName, timeout, longPollTimeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMatchingClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewMatchingClientWithTimeout), namespaceIDToName, timeout, longPollTimeout)
}

// NewRemoteAdminClientWithTimeout mocks base method.
func (m *MockFactory) NewRemoteAdminClientWithTimeout(rpcAddress string, timeout, largeTimeout time.Duration) adminservice.AdminServiceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRemoteAdminClientWithTimeout", rpcAddress, timeout, largeTimeout)
	ret0, _ := ret[0].(adminservice.AdminServiceClient)
	return ret0
}

// NewRemoteAdminClientWithTimeout indicates an expected call of NewRemoteAdminClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewRemoteAdminClientWithTimeout(rpcAddress, timeout, largeTimeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRemoteAdminClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewRemoteAdminClientWithTimeout), rpcAddress, timeout, largeTimeout)
}

// NewRemoteFrontendClientWithTimeout mocks base method.
func (m *MockFactory) NewRemoteFrontendClientWithTimeout(rpcAddress string, timeout, longPollTimeout time.Duration) (grpc.ClientConnInterface, workflowservice.WorkflowServiceClient) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRemoteFrontendClientWithTimeout", rpcAddress, timeout, longPollTimeout)
	ret0, _ := ret[0].(grpc.ClientConnInterface)
	ret1, _ := ret[1].(workflowservice.WorkflowServiceClient)
	return ret0, ret1
}

// NewRemoteFrontendClientWithTimeout indicates an expected call of NewRemoteFrontendClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewRemoteFrontendClientWithTimeout(rpcAddress, timeout, longPollTimeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRemoteFrontendClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewRemoteFrontendClientWithTimeout), rpcAddress, timeout, longPollTimeout)
}

// MockFactoryProvider is a mock of FactoryProvider interface.
type MockFactoryProvider struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryProviderMockRecorder
	isgomock struct{}
}

// MockFactoryProviderMockRecorder is the mock recorder for MockFactoryProvider.
type MockFactoryProviderMockRecorder struct {
	mock *MockFactoryProvider
}

// NewMockFactoryProvider creates a new mock instance.
func NewMockFactoryProvider(ctrl *gomock.Controller) *MockFactoryProvider {
	mock := &MockFactoryProvider{ctrl: ctrl}
	mock.recorder = &MockFactoryProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactoryProvider) EXPECT() *MockFactoryProviderMockRecorder {
	return m.recorder
}

// NewFactory mocks base method.
func (m *MockFactoryProvider) NewFactory(rpcFactory common.RPCFactory, monitor membership.Monitor, metricsHandler metrics.Handler, dc *dynamicconfig.Collection, testHooks testhooks.TestHooks, numberOfHistoryShards int32, logger, throttledLogger log.Logger) Factory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFactory", rpcFactory, monitor, metricsHandler, dc, testHooks, numberOfHistoryShards, logger, throttledLogger)
	ret0, _ := ret[0].(Factory)
	return ret0
}

// NewFactory indicates an expected call of NewFactory.
func (mr *MockFactoryProviderMockRecorder) NewFactory(rpcFactory, monitor, metricsHandler, dc, testHooks, numberOfHistoryShards, logger, throttledLogger any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFactory", reflect.TypeOf((*MockFactoryProvider)(nil).NewFactory), rpcFactory, monitor, metricsHandler, dc, testHooks, numberOfHistoryShards, logger, throttledLogger)
}
