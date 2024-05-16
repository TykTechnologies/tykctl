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
func (m *MockCloudClient) CreateDeployment(ctx context.Context, deployment cloud.Deployment, orgID, teamID, envID string) (cloud.InlineResponse2001, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeployment", ctx, deployment, orgID, teamID, envID)
	ret0, _ := ret[0].(cloud.InlineResponse2001)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDeployment indicates an expected call of CreateDeployment.
func (mr *MockCloudClientMockRecorder) CreateDeployment(ctx, deployment, orgID, teamID, envID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeployment", reflect.TypeOf((*MockCloudClient)(nil).CreateDeployment), ctx, deployment, orgID, teamID, envID)
}

// CreateEnv mocks base method.
func (m *MockCloudClient) CreateEnv(ctx context.Context, env cloud.Loadout, orgID, teamID string) (cloud.InlineResponse2012, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEnv", ctx, env, orgID, teamID)
	ret0, _ := ret[0].(cloud.InlineResponse2012)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateEnv indicates an expected call of CreateEnv.
func (mr *MockCloudClientMockRecorder) CreateEnv(ctx, env, orgID, teamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnv", reflect.TypeOf((*MockCloudClient)(nil).CreateEnv), ctx, env, orgID, teamID)
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

// DeleteEnv mocks base method.
func (m *MockCloudClient) DeleteEnv(ctx context.Context, oid, tid, id string, localVarOptionals *cloud.LoadoutsApiDeleteLoadoutOpts) (cloud.InlineResponse2012, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEnv", ctx, oid, tid, id, localVarOptionals)
	ret0, _ := ret[0].(cloud.InlineResponse2012)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteEnv indicates an expected call of DeleteEnv.
func (mr *MockCloudClientMockRecorder) DeleteEnv(ctx, oid, tid, id, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnv", reflect.TypeOf((*MockCloudClient)(nil).DeleteEnv), ctx, oid, tid, id, localVarOptionals)
}

// DeleteTeam mocks base method.
func (m *MockCloudClient) DeleteTeam(ctx context.Context, oid, tid string, localVarOptionals *cloud.TeamsApiDeleteTeamOpts) (cloud.InlineResponse2011, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTeam", ctx, oid, tid, localVarOptionals)
	ret0, _ := ret[0].(cloud.InlineResponse2011)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteTeam indicates an expected call of DeleteTeam.
func (mr *MockCloudClientMockRecorder) DeleteTeam(ctx, oid, tid, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeam", reflect.TypeOf((*MockCloudClient)(nil).DeleteTeam), ctx, oid, tid, localVarOptionals)
}

// DestroyDeployment mocks base method.
func (m *MockCloudClient) DestroyDeployment(ctx context.Context, oid, tid, lid, id string, localVarOptionals *cloud.DeploymentsApiDestroyDeploymentOpts) (cloud.InlineResponse2001, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyDeployment", ctx, oid, tid, lid, id, localVarOptionals)
	ret0, _ := ret[0].(cloud.InlineResponse2001)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DestroyDeployment indicates an expected call of DestroyDeployment.
func (mr *MockCloudClientMockRecorder) DestroyDeployment(ctx, oid, tid, lid, id, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyDeployment", reflect.TypeOf((*MockCloudClient)(nil).DestroyDeployment), ctx, oid, tid, lid, id, localVarOptionals)
}

// GetDeploymentByID mocks base method.
func (m *MockCloudClient) GetDeploymentByID(ctx context.Context, orgID, teamID, envID, id string, localVarOptionals *cloud.DeploymentsApiGetDeploymentOpts) (cloud.InlineResponse2001, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentByID", ctx, orgID, teamID, envID, id, localVarOptionals)
	ret0, _ := ret[0].(cloud.InlineResponse2001)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeploymentByID indicates an expected call of GetDeploymentByID.
func (mr *MockCloudClientMockRecorder) GetDeploymentByID(ctx, orgID, teamID, envID, id, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentByID", reflect.TypeOf((*MockCloudClient)(nil).GetDeploymentByID), ctx, orgID, teamID, envID, id, localVarOptionals)
}

// GetDeploymentStatus mocks base method.
func (m *MockCloudClient) GetDeploymentStatus(ctx context.Context, orgID, teamID, envID, deploymentID string) (cloud.Payload, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentStatus", ctx, orgID, teamID, envID, deploymentID)
	ret0, _ := ret[0].(cloud.Payload)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeploymentStatus indicates an expected call of GetDeploymentStatus.
func (mr *MockCloudClientMockRecorder) GetDeploymentStatus(ctx, orgID, teamID, envID, deploymentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentStatus", reflect.TypeOf((*MockCloudClient)(nil).GetDeploymentStatus), ctx, orgID, teamID, envID, deploymentID)
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

// GetEnvByID mocks base method.
func (m *MockCloudClient) GetEnvByID(ctx context.Context, orgID, teamID, envID string) (cloud.InlineResponse2012, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvByID", ctx, orgID, teamID, envID)
	ret0, _ := ret[0].(cloud.InlineResponse2012)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEnvByID indicates an expected call of GetEnvByID.
func (mr *MockCloudClientMockRecorder) GetEnvByID(ctx, orgID, teamID, envID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvByID", reflect.TypeOf((*MockCloudClient)(nil).GetEnvByID), ctx, orgID, teamID, envID)
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
func (m *MockCloudClient) GetEnvs(ctx context.Context, orgID, teamID string) (cloud.InlineResponse20016, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvs", ctx, orgID, teamID)
	ret0, _ := ret[0].(cloud.InlineResponse20016)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEnvs indicates an expected call of GetEnvs.
func (mr *MockCloudClientMockRecorder) GetEnvs(ctx, orgID, teamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvs", reflect.TypeOf((*MockCloudClient)(nil).GetEnvs), ctx, orgID, teamID)
}

// GetOrgByID mocks base method.
func (m *MockCloudClient) GetOrgByID(ctx context.Context, oid string) (cloud.InlineResponse2005, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgByID", ctx, oid)
	ret0, _ := ret[0].(cloud.InlineResponse2005)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrgByID indicates an expected call of GetOrgByID.
func (mr *MockCloudClientMockRecorder) GetOrgByID(ctx, oid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgByID", reflect.TypeOf((*MockCloudClient)(nil).GetOrgByID), ctx, oid)
}

// GetOrgInfo mocks base method.
func (m *MockCloudClient) GetOrgInfo(ctx context.Context, orgID string) (*internal.OrgInfo, *resty.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgInfo", ctx, orgID)
	ret0, _ := ret[0].(*internal.OrgInfo)
	ret1, _ := ret[1].(*resty.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrgInfo indicates an expected call of GetOrgInfo.
func (mr *MockCloudClientMockRecorder) GetOrgInfo(ctx, orgID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgInfo", reflect.TypeOf((*MockCloudClient)(nil).GetOrgInfo), ctx, orgID)
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

// GetTeamByID mocks base method.
func (m *MockCloudClient) GetTeamByID(ctx context.Context, oid, tid string) (cloud.InlineResponse2011, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeamByID", ctx, oid, tid)
	ret0, _ := ret[0].(cloud.InlineResponse2011)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTeamByID indicates an expected call of GetTeamByID.
func (mr *MockCloudClientMockRecorder) GetTeamByID(ctx, oid, tid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeamByID", reflect.TypeOf((*MockCloudClient)(nil).GetTeamByID), ctx, oid, tid)
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

// GetUserInfo mocks base method.
func (m *MockCloudClient) GetUserInfo(ctx context.Context) (*internal.UserInfo, *resty.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", ctx)
	ret0, _ := ret[0].(*internal.UserInfo)
	ret1, _ := ret[1].(*resty.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockCloudClientMockRecorder) GetUserInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockCloudClient)(nil).GetUserInfo), ctx)
}

// RestartDeployment mocks base method.
func (m *MockCloudClient) RestartDeployment(ctx context.Context, oid, tid, lid, id string) (cloud.InlineResponse2001, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestartDeployment", ctx, oid, tid, lid, id)
	ret0, _ := ret[0].(cloud.InlineResponse2001)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RestartDeployment indicates an expected call of RestartDeployment.
func (mr *MockCloudClientMockRecorder) RestartDeployment(ctx, oid, tid, lid, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestartDeployment", reflect.TypeOf((*MockCloudClient)(nil).RestartDeployment), ctx, oid, tid, lid, id)
}

// StartDeployment mocks base method.
func (m *MockCloudClient) StartDeployment(ctx context.Context, orgID, teamID, envID, id string) (cloud.InlineResponse2001, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartDeployment", ctx, orgID, teamID, envID, id)
	ret0, _ := ret[0].(cloud.InlineResponse2001)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StartDeployment indicates an expected call of StartDeployment.
func (mr *MockCloudClientMockRecorder) StartDeployment(ctx, orgID, teamID, envID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDeployment", reflect.TypeOf((*MockCloudClient)(nil).StartDeployment), ctx, orgID, teamID, envID, id)
}

// UpdateDeployment mocks base method.
func (m *MockCloudClient) UpdateDeployment(ctx context.Context, body cloud.Deployment, orgID, teamID, envID, id string, localVarOptionals *cloud.DeploymentsApiUpdateDeploymentOpts) (cloud.InlineResponse2001, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeployment", ctx, body, orgID, teamID, envID, id, localVarOptionals)
	ret0, _ := ret[0].(cloud.InlineResponse2001)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateDeployment indicates an expected call of UpdateDeployment.
func (mr *MockCloudClientMockRecorder) UpdateDeployment(ctx, body, orgID, teamID, envID, id, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeployment", reflect.TypeOf((*MockCloudClient)(nil).UpdateDeployment), ctx, body, orgID, teamID, envID, id, localVarOptionals)
}

// UpdateTeam mocks base method.
func (m *MockCloudClient) UpdateTeam(ctx context.Context, team cloud.Team, orgID, teamID string) (cloud.InlineResponse2011, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeam", ctx, team, orgID, teamID)
	ret0, _ := ret[0].(cloud.InlineResponse2011)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateTeam indicates an expected call of UpdateTeam.
func (mr *MockCloudClientMockRecorder) UpdateTeam(ctx, team, orgID, teamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeam", reflect.TypeOf((*MockCloudClient)(nil).UpdateTeam), ctx, team, orgID, teamID)
}
