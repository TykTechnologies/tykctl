// Code generated by MockGen. DO NOT EDIT.
// Source: cloud_client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	cloud "github.com/TykTechnologies/cloud-sdk/cloud"
	internal "github.com/TykTechnologies/tykctl/internal"
	resty "github.com/go-resty/resty/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudClient is a mock of CloudClient interface.
type MockCloudClient struct {
	ctrl     *gomock.Controller
	recorder *MockCloudClientMockRecorder
}

// MockCloudClientMockRecorder is the mock recorder for MockCloudClient.
type MockCloudClientMockRecorder struct {
	mock *MockCloudClient
}

// NewMockCloudClient creates a new mock instance.
func NewMockCloudClient(ctrl *gomock.Controller) *MockCloudClient {
	mock := &MockCloudClient{ctrl: ctrl}
	mock.recorder = &MockCloudClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudClient) EXPECT() *MockCloudClientMockRecorder {
	return m.recorder
}

// CreateDeployment mocks base method.
func (m *MockCloudClient) CreateDeployment(ctx context.Context, deployment cloud.Deployment, orgId, teamId, envId string) (cloud.InlineResponse2001, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeployment", ctx, deployment, orgId, teamId, envId)
	ret0, _ := ret[0].(cloud.InlineResponse2001)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDeployment indicates an expected call of CreateDeployment.
func (mr *MockCloudClientMockRecorder) CreateDeployment(ctx, deployment, orgId, teamId, envId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeployment", reflect.TypeOf((*MockCloudClient)(nil).CreateDeployment), ctx, deployment, orgId, teamId, envId)
}

// CreateEnv mocks base method.
func (m *MockCloudClient) CreateEnv(ctx context.Context, env cloud.Loadout, orgId, teamId string) (cloud.InlineResponse2012, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEnv", ctx, env, orgId, teamId)
	ret0, _ := ret[0].(cloud.InlineResponse2012)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateEnv indicates an expected call of CreateEnv.
func (mr *MockCloudClientMockRecorder) CreateEnv(ctx, env, orgId, teamId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnv", reflect.TypeOf((*MockCloudClient)(nil).CreateEnv), ctx, env, orgId, teamId)
}

// CreateTeam mocks base method.
func (m *MockCloudClient) CreateTeam(ctx context.Context, team cloud.Team, oid string) (cloud.InlineResponse2011, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeam", ctx, team, oid)
	ret0, _ := ret[0].(cloud.InlineResponse2011)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTeam indicates an expected call of CreateTeam.
func (mr *MockCloudClientMockRecorder) CreateTeam(ctx, team, oid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*MockCloudClient)(nil).CreateTeam), ctx, team, oid)
}

// GetDeploymentById mocks base method.
func (m *MockCloudClient) GetDeploymentById(ctx context.Context, orgId, teamId, envId, id string, localVarOptionals *cloud.DeploymentsApiGetDeploymentOpts) (cloud.InlineResponse2001, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentById", ctx, orgId, teamId, envId, id, localVarOptionals)
	ret0, _ := ret[0].(cloud.InlineResponse2001)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeploymentById indicates an expected call of GetDeploymentById.
func (mr *MockCloudClientMockRecorder) GetDeploymentById(ctx, orgId, teamId, envId, id, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentById", reflect.TypeOf((*MockCloudClient)(nil).GetDeploymentById), ctx, orgId, teamId, envId, id, localVarOptionals)
}

// GetDeploymentStatus mocks base method.
func (m *MockCloudClient) GetDeploymentStatus(ctx context.Context, orgId, teamId, envId, deploymentId string) (cloud.Payload, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentStatus", ctx, orgId, teamId, envId, deploymentId)
	ret0, _ := ret[0].(cloud.Payload)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeploymentStatus indicates an expected call of GetDeploymentStatus.
func (mr *MockCloudClientMockRecorder) GetDeploymentStatus(ctx, orgId, teamId, envId, deploymentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentStatus", reflect.TypeOf((*MockCloudClient)(nil).GetDeploymentStatus), ctx, orgId, teamId, envId, deploymentId)
}

// GetDeploymentZones mocks base method.
func (m *MockCloudClient) GetDeploymentZones(ctx context.Context) (*internal.ZoneResponse, *resty.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentZones", ctx)
	ret0, _ := ret[0].(*internal.ZoneResponse)
	ret1, _ := ret[1].(*resty.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeploymentZones indicates an expected call of GetDeploymentZones.
func (mr *MockCloudClientMockRecorder) GetDeploymentZones(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentZones", reflect.TypeOf((*MockCloudClient)(nil).GetDeploymentZones), ctx)
}

// GetEnvById mocks base method.
func (m *MockCloudClient) GetEnvById(ctx context.Context, orgId, teamId, envId string) (cloud.InlineResponse2012, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvById", ctx, orgId, teamId, envId)
	ret0, _ := ret[0].(cloud.InlineResponse2012)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEnvById indicates an expected call of GetEnvById.
func (mr *MockCloudClientMockRecorder) GetEnvById(ctx, orgId, teamId, envId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvById", reflect.TypeOf((*MockCloudClient)(nil).GetEnvById), ctx, orgId, teamId, envId)
}

// GetEnvDeployments mocks base method.
func (m *MockCloudClient) GetEnvDeployments(ctx context.Context, oid, tid, lid string) (cloud.InlineResponse200, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvDeployments", ctx, oid, tid, lid)
	ret0, _ := ret[0].(cloud.InlineResponse200)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEnvDeployments indicates an expected call of GetEnvDeployments.
func (mr *MockCloudClientMockRecorder) GetEnvDeployments(ctx, oid, tid, lid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvDeployments", reflect.TypeOf((*MockCloudClient)(nil).GetEnvDeployments), ctx, oid, tid, lid)
}

// GetEnvs mocks base method.
func (m *MockCloudClient) GetEnvs(ctx context.Context, orgId, teamId string) (cloud.InlineResponse20016, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvs", ctx, orgId, teamId)
	ret0, _ := ret[0].(cloud.InlineResponse20016)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEnvs indicates an expected call of GetEnvs.
func (mr *MockCloudClientMockRecorder) GetEnvs(ctx, orgId, teamId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvs", reflect.TypeOf((*MockCloudClient)(nil).GetEnvs), ctx, orgId, teamId)
}

// GetOrgById mocks base method.
func (m *MockCloudClient) GetOrgById(ctx context.Context, oid string) (cloud.InlineResponse2005, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgById", ctx, oid)
	ret0, _ := ret[0].(cloud.InlineResponse2005)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrgById indicates an expected call of GetOrgById.
func (mr *MockCloudClientMockRecorder) GetOrgById(ctx, oid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgById", reflect.TypeOf((*MockCloudClient)(nil).GetOrgById), ctx, oid)
}

// GetOrgs mocks base method.
func (m *MockCloudClient) GetOrgs(ctx context.Context) (cloud.InlineResponse20014, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgs", ctx)
	ret0, _ := ret[0].(cloud.InlineResponse20014)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrgs indicates an expected call of GetOrgs.
func (mr *MockCloudClientMockRecorder) GetOrgs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgs", reflect.TypeOf((*MockCloudClient)(nil).GetOrgs), ctx)
}

// GetTeamById mocks base method.
func (m *MockCloudClient) GetTeamById(ctx context.Context, oid, tid string) (cloud.InlineResponse2011, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeamById", ctx, oid, tid)
	ret0, _ := ret[0].(cloud.InlineResponse2011)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTeamById indicates an expected call of GetTeamById.
func (mr *MockCloudClientMockRecorder) GetTeamById(ctx, oid, tid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeamById", reflect.TypeOf((*MockCloudClient)(nil).GetTeamById), ctx, oid, tid)
}

// GetTeams mocks base method.
func (m *MockCloudClient) GetTeams(ctx context.Context, oid string) (cloud.InlineResponse20017, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeams", ctx, oid)
	ret0, _ := ret[0].(cloud.InlineResponse20017)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTeams indicates an expected call of GetTeams.
func (mr *MockCloudClientMockRecorder) GetTeams(ctx, oid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeams", reflect.TypeOf((*MockCloudClient)(nil).GetTeams), ctx, oid)
}

// StartDeployment mocks base method.
func (m *MockCloudClient) StartDeployment(ctx context.Context, orgID, teamId, envId, id string) (cloud.InlineResponse2001, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartDeployment", ctx, orgID, teamId, envId, id)
	ret0, _ := ret[0].(cloud.InlineResponse2001)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StartDeployment indicates an expected call of StartDeployment.
func (mr *MockCloudClientMockRecorder) StartDeployment(ctx, orgID, teamId, envId, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDeployment", reflect.TypeOf((*MockCloudClient)(nil).StartDeployment), ctx, orgID, teamId, envId, id)
}
