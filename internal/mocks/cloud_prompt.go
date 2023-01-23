// Code generated by MockGen. DO NOT EDIT.
// Source: cloud_prompt.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	cloud "github.com/TykTechnologies/cloud-sdk/cloud"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudPrompt is a mock of CloudPrompt interface.
type MockCloudPrompt struct {
	ctrl     *gomock.Controller
	recorder *MockCloudPromptMockRecorder
}

// MockCloudPromptMockRecorder is the mock recorder for MockCloudPrompt.
type MockCloudPromptMockRecorder struct {
	mock *MockCloudPrompt
}

// NewMockCloudPrompt creates a new mock instance.
func NewMockCloudPrompt(ctrl *gomock.Controller) *MockCloudPrompt {
	mock := &MockCloudPrompt{ctrl: ctrl}
	mock.recorder = &MockCloudPromptMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudPrompt) EXPECT() *MockCloudPromptMockRecorder {
	return m.recorder
}

// EnvPrompt mocks base method.
func (m *MockCloudPrompt) EnvPrompt(envs []cloud.Loadout) (*cloud.Loadout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvPrompt", envs)
	ret0, _ := ret[0].(*cloud.Loadout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvPrompt indicates an expected call of EnvPrompt.
func (mr *MockCloudPromptMockRecorder) EnvPrompt(envs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvPrompt", reflect.TypeOf((*MockCloudPrompt)(nil).EnvPrompt), envs)
}

// OrgPrompt mocks base method.
func (m *MockCloudPrompt) OrgPrompt(orgs []cloud.Organisation) (*cloud.Organisation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrgPrompt", orgs)
	ret0, _ := ret[0].(*cloud.Organisation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrgPrompt indicates an expected call of OrgPrompt.
func (mr *MockCloudPromptMockRecorder) OrgPrompt(orgs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrgPrompt", reflect.TypeOf((*MockCloudPrompt)(nil).OrgPrompt), orgs)
}

// RegionPrompt mocks base method.
func (m *MockCloudPrompt) RegionPrompt(regions []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegionPrompt", regions)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegionPrompt indicates an expected call of RegionPrompt.
func (mr *MockCloudPromptMockRecorder) RegionPrompt(regions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegionPrompt", reflect.TypeOf((*MockCloudPrompt)(nil).RegionPrompt), regions)
}

// TeamPrompt mocks base method.
func (m *MockCloudPrompt) TeamPrompt(teams []cloud.Team) (*cloud.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamPrompt", teams)
	ret0, _ := ret[0].(*cloud.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamPrompt indicates an expected call of TeamPrompt.
func (mr *MockCloudPromptMockRecorder) TeamPrompt(teams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamPrompt", reflect.TypeOf((*MockCloudPrompt)(nil).TeamPrompt), teams)
}
