// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo-edge/projects/gloo/cli/pkg/cmd/install (interfaces: HelmReleaseListRunner)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	release "helm.sh/helm/v3/pkg/release"
)

// MockHelmReleaseListRunner is a mock of HelmReleaseListRunner interface
type MockHelmReleaseListRunner struct {
	ctrl     *gomock.Controller
	recorder *MockHelmReleaseListRunnerMockRecorder
}

// MockHelmReleaseListRunnerMockRecorder is the mock recorder for MockHelmReleaseListRunner
type MockHelmReleaseListRunnerMockRecorder struct {
	mock *MockHelmReleaseListRunner
}

// NewMockHelmReleaseListRunner creates a new mock instance
func NewMockHelmReleaseListRunner(ctrl *gomock.Controller) *MockHelmReleaseListRunner {
	mock := &MockHelmReleaseListRunner{ctrl: ctrl}
	mock.recorder = &MockHelmReleaseListRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHelmReleaseListRunner) EXPECT() *MockHelmReleaseListRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockHelmReleaseListRunner) Run() ([]*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].([]*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockHelmReleaseListRunnerMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockHelmReleaseListRunner)(nil).Run))
}

// SetFilter mocks base method
func (m *MockHelmReleaseListRunner) SetFilter(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFilter", arg0)
}

// SetFilter indicates an expected call of SetFilter
func (mr *MockHelmReleaseListRunnerMockRecorder) SetFilter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFilter", reflect.TypeOf((*MockHelmReleaseListRunner)(nil).SetFilter), arg0)
}
