package internal

import (
	"context"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/TykTechnologies/cloud-sdk/cloud"
)

//go:generate mockgen -source=cloud_client.go -destination=./mocks/cloud_client.go -package=mock CloudClient
type CloudClient interface {
	GetOrgs(ctx context.Context) (cloud.InlineResponse20014, *http.Response, error)
	GetOrgByID(ctx context.Context, oid string) (cloud.InlineResponse2005, *http.Response, error)
	CreateTeam(ctx context.Context, team cloud.Team, oid string) (cloud.InlineResponse2011, *http.Response, error)
	GetTeamByID(ctx context.Context, oid, tid string) (cloud.InlineResponse2011, *http.Response, error)
	GetTeams(ctx context.Context, oid string) (cloud.InlineResponse20017, *http.Response, error)
	UpdateTeam(ctx context.Context, team cloud.Team, orgID, teamID string) (cloud.InlineResponse2011, *http.Response, error)
	CreateEnv(ctx context.Context, env cloud.Loadout, orgID, teamID string) (cloud.InlineResponse2012, *http.Response, error)
	GetEnvByID(ctx context.Context, orgID, teamID, envID string) (cloud.InlineResponse2012, *http.Response, error)
	GetEnvs(ctx context.Context, orgID, teamID string) (cloud.InlineResponse20016, *http.Response, error)
	CreateDeployment(ctx context.Context, deployment cloud.Deployment, orgID, teamID, envID string) (cloud.InlineResponse2001, *http.Response, error)
	GetEnvDeployments(ctx context.Context, oid, tid, lid string) (cloud.InlineResponse200, *http.Response, error)
	GetDeploymentByID(ctx context.Context, orgID, teamID, envID, id string, localVarOptionals *cloud.DeploymentsApiGetDeploymentOpts) (cloud.InlineResponse2001, *http.Response, error)
	StartDeployment(ctx context.Context, orgID, teamID, envID, id string) (cloud.InlineResponse2001, *http.Response, error)
	RestartDeployment(ctx context.Context, oid, tid, lid, id string) (cloud.InlineResponse2001, *http.Response, error)
	GetDeploymentStatus(ctx context.Context, orgID, teamID, envID, deploymentID string) (cloud.Payload, *http.Response, error)
	GetDeploymentZones(ctx context.Context) (*ZoneResponse, *resty.Response, error)
	GetUserInfo(ctx context.Context) (*UserInfo, *resty.Response, error)
	GetOrgInfo(ctx context.Context, orgID string) (*OrgInfo, *resty.Response, error)
	UpdateDeployment(ctx context.Context, body cloud.Deployment, orgID, teamID, envID, id string, localVarOptionals *cloud.DeploymentsApiUpdateDeploymentOpts) (cloud.InlineResponse2001, *http.Response, error)
}
