package internal

import (
	"context"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/go-resty/resty/v2"
	"net/http"
)

//go:generate mockgen -source=cloud_client.go -destination=./mocks/cloud_client.go -package=mock CloudClient
type CloudClient interface {
	GetOrgs(ctx context.Context) (cloud.InlineResponse20014, *http.Response, error)
	GetOrgById(ctx context.Context, oid string) (cloud.InlineResponse2005, *http.Response, error)
	CreateTeam(ctx context.Context, team cloud.Team, oid string) (cloud.InlineResponse2011, *http.Response, error)
	GetTeamById(ctx context.Context, oid string, tid string) (cloud.InlineResponse2011, *http.Response, error)
	GetTeams(ctx context.Context, oid string) (cloud.InlineResponse20017, *http.Response, error)
	UpdateTeam(ctx context.Context, team cloud.Team, orgId string, teamId string) (cloud.InlineResponse2011, *http.Response, error)
	CreateEnv(ctx context.Context, env cloud.Loadout, orgId string, teamId string) (cloud.InlineResponse2012, *http.Response, error)
	GetEnvById(ctx context.Context, orgId string, teamId string, envId string) (cloud.InlineResponse2012, *http.Response, error)
	GetEnvs(ctx context.Context, orgId string, teamId string) (cloud.InlineResponse20016, *http.Response, error)
	CreateDeployment(ctx context.Context, deployment cloud.Deployment, orgId, teamId, envId string) (cloud.InlineResponse2001, *http.Response, error)
	GetEnvDeployments(ctx context.Context, oid string, tid string, lid string) (cloud.InlineResponse200, *http.Response, error)
	GetDeploymentById(ctx context.Context, orgId string, teamId string, envId string, id string, localVarOptionals *cloud.DeploymentsApiGetDeploymentOpts) (cloud.InlineResponse2001, *http.Response, error)
	StartDeployment(ctx context.Context, orgID, teamId, envId string, id string) (cloud.InlineResponse2001, *http.Response, error)
	RestartDeployment(ctx context.Context, oid string, tid string, lid string, id string) (cloud.InlineResponse2001, *http.Response, error)
	GetDeploymentStatus(ctx context.Context, orgId string, teamId string, envId string, deploymentId string) (cloud.Payload, *http.Response, error)
	GetDeploymentZones(ctx context.Context) (*ZoneResponse, *resty.Response, error)
	GetUserInfo(ctx context.Context) (*UserInfo, *resty.Response, error)
	GetOrgInfo(ctx context.Context, orgId string) (*OrgInfo, *resty.Response, error)
}
