// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-cni-k8s/rpc (interfaces: CNIBackendClient,NPBackendClient)

// Package mock_rpc is a generated GoMock package.
package mock_rpc

import (
	context "context"
	reflect "reflect"

	rpc "github.com/aws/amazon-vpc-cni-k8s/rpc"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockCNIBackendClient is a mock of CNIBackendClient interface.
type MockCNIBackendClient struct {
	ctrl     *gomock.Controller
	recorder *MockCNIBackendClientMockRecorder
}

// MockCNIBackendClientMockRecorder is the mock recorder for MockCNIBackendClient.
type MockCNIBackendClientMockRecorder struct {
	mock *MockCNIBackendClient
}

// NewMockCNIBackendClient creates a new mock instance.
func NewMockCNIBackendClient(ctrl *gomock.Controller) *MockCNIBackendClient {
	mock := &MockCNIBackendClient{ctrl: ctrl}
	mock.recorder = &MockCNIBackendClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCNIBackendClient) EXPECT() *MockCNIBackendClientMockRecorder {
	return m.recorder
}

// AddNetwork mocks base method.
func (m *MockCNIBackendClient) AddNetwork(arg0 context.Context, arg1 *rpc.AddNetworkRequest, arg2 ...grpc.CallOption) (*rpc.AddNetworkReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddNetwork", varargs...)
	ret0, _ := ret[0].(*rpc.AddNetworkReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNetwork indicates an expected call of AddNetwork.
func (mr *MockCNIBackendClientMockRecorder) AddNetwork(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNetwork", reflect.TypeOf((*MockCNIBackendClient)(nil).AddNetwork), varargs...)
}

// DelNetwork mocks base method.
func (m *MockCNIBackendClient) DelNetwork(arg0 context.Context, arg1 *rpc.DelNetworkRequest, arg2 ...grpc.CallOption) (*rpc.DelNetworkReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DelNetwork", varargs...)
	ret0, _ := ret[0].(*rpc.DelNetworkReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelNetwork indicates an expected call of DelNetwork.
func (mr *MockCNIBackendClientMockRecorder) DelNetwork(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelNetwork", reflect.TypeOf((*MockCNIBackendClient)(nil).DelNetwork), varargs...)
}

// MockNPBackendClient is a mock of NPBackendClient interface.
type MockNPBackendClient struct {
	ctrl     *gomock.Controller
	recorder *MockNPBackendClientMockRecorder
}

// MockNPBackendClientMockRecorder is the mock recorder for MockNPBackendClient.
type MockNPBackendClientMockRecorder struct {
	mock *MockNPBackendClient
}

// NewMockNPBackendClient creates a new mock instance.
func NewMockNPBackendClient(ctrl *gomock.Controller) *MockNPBackendClient {
	mock := &MockNPBackendClient{ctrl: ctrl}
	mock.recorder = &MockNPBackendClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNPBackendClient) EXPECT() *MockNPBackendClientMockRecorder {
	return m.recorder
}

// DeletePodNp mocks base method.
func (m *MockNPBackendClient) DeletePodNp(arg0 context.Context, arg1 *rpc.DeleteNpRequest, arg2 ...grpc.CallOption) (*rpc.DeleteNpReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePodNp", varargs...)
	ret0, _ := ret[0].(*rpc.DeleteNpReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePodNp indicates an expected call of DeletePodNp.
func (mr *MockNPBackendClientMockRecorder) DeletePodNp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePodNp", reflect.TypeOf((*MockNPBackendClient)(nil).DeletePodNp), varargs...)
}

// EnforceNpToPod mocks base method.
func (m *MockNPBackendClient) EnforceNpToPod(arg0 context.Context, arg1 *rpc.EnforceNpRequest, arg2 ...grpc.CallOption) (*rpc.EnforceNpReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnforceNpToPod", varargs...)
	ret0, _ := ret[0].(*rpc.EnforceNpReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnforceNpToPod indicates an expected call of EnforceNpToPod.
func (mr *MockNPBackendClientMockRecorder) EnforceNpToPod(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnforceNpToPod", reflect.TypeOf((*MockNPBackendClient)(nil).EnforceNpToPod), varargs...)
}
