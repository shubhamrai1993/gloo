// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo-edge/projects/gloo/pkg/api/v1 (interfaces: ArtifactClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-edge/projects/gloo/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockArtifactClient is a mock of ArtifactClient interface
type MockArtifactClient struct {
	ctrl     *gomock.Controller
	recorder *MockArtifactClientMockRecorder
}

// MockArtifactClientMockRecorder is the mock recorder for MockArtifactClient
type MockArtifactClientMockRecorder struct {
	mock *MockArtifactClient
}

// NewMockArtifactClient creates a new mock instance
func NewMockArtifactClient(ctrl *gomock.Controller) *MockArtifactClient {
	mock := &MockArtifactClient{ctrl: ctrl}
	mock.recorder = &MockArtifactClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArtifactClient) EXPECT() *MockArtifactClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockArtifactClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockArtifactClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockArtifactClient)(nil).BaseClient))
}

// Delete mocks base method
func (m *MockArtifactClient) Delete(arg0, arg1 string, arg2 clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockArtifactClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArtifactClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockArtifactClient) List(arg0 string, arg1 clients.ListOpts) (v1.ArtifactList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(v1.ArtifactList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockArtifactClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArtifactClient)(nil).List), arg0, arg1)
}

// Read mocks base method
func (m *MockArtifactClient) Read(arg0, arg1 string, arg2 clients.ReadOpts) (*v1.Artifact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Artifact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockArtifactClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockArtifactClient)(nil).Read), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockArtifactClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockArtifactClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockArtifactClient)(nil).Register))
}

// Watch mocks base method
func (m *MockArtifactClient) Watch(arg0 string, arg1 clients.WatchOpts) (<-chan v1.ArtifactList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.ArtifactList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockArtifactClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockArtifactClient)(nil).Watch), arg0, arg1)
}

// Write mocks base method
func (m *MockArtifactClient) Write(arg0 *v1.Artifact, arg1 clients.WriteOpts) (*v1.Artifact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*v1.Artifact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockArtifactClientMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockArtifactClient)(nil).Write), arg0, arg1)
}
