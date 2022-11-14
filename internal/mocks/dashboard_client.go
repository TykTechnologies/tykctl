// Code generated by MockGen. DO NOT EDIT.
// Source: dashboard_client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	internal "github.com/TykTechnologies/tykctl/internal"
	gomock "github.com/golang/mock/gomock"
)

// MockDashboardClient is a mock of DashboardClient interface.
type MockDashboardClient struct {
	ctrl     *gomock.Controller
	recorder *MockDashboardClientMockRecorder
}

// MockDashboardClientMockRecorder is the mock recorder for MockDashboardClient.
type MockDashboardClientMockRecorder struct {
	mock *MockDashboardClient
}

// NewMockDashboardClient creates a new mock instance.
func NewMockDashboardClient(ctrl *gomock.Controller) *MockDashboardClient {
	mock := &MockDashboardClient{ctrl: ctrl}
	mock.recorder = &MockDashboardClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDashboardClient) EXPECT() *MockDashboardClientMockRecorder {
	return m.recorder
}

// GetDeploymentZones mocks base method.
func (m *MockDashboardClient) GetDeploymentZones(ctx context.Context) (internal.ZoneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentZones", ctx)
	ret0, _ := ret[0].(internal.ZoneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentZones indicates an expected call of GetDeploymentZones.
func (mr *MockDashboardClientMockRecorder) GetDeploymentZones(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentZones", reflect.TypeOf((*MockDashboardClient)(nil).GetDeploymentZones), ctx)
}
