package internal

import (
	"context"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"net/http"
)

var (
	_ CloudClient = (*cloudSdkClient)(nil)
)

// CloudSdkClient should implement CloudClient as it will be used to make request to Ara.
type cloudSdkClient struct {
	Client *cloud.APIClient
	Config *cloud.Configuration
	//beforeExecute will store function you need called before each method runs .
	beforeExecute []func(*cloud.APIClient, *cloud.Configuration) error
}

// NewCloudSdkClient creates a new cloudSdkClient to make sure that a client is never nil.
func NewCloudSdkClient(conf *cloud.Configuration) *cloudSdkClient {
	client := cloud.NewAPIClient(conf)
	return &cloudSdkClient{
		Client:        client,
		beforeExecute: nil,
		Config:        conf,
	}
}

// GetOrgs a users organizations from Ara.
func (c *cloudSdkClient) GetOrgs(ctx context.Context) (cloud.InlineResponse20014, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse20014{}, nil, err
	}
	return c.Client.OrganisationsApi.GetOrgs(ctx)
}

// GetOrgById will get a single organization using its id.
func (c *cloudSdkClient) GetOrgById(ctx context.Context, oid string) (cloud.InlineResponse2005, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2005{}, nil, err
	}
	return c.Client.OrganisationsApi.GetOrg(ctx, oid)
}

// CreateTeam create a team in an organization.
func (c *cloudSdkClient) CreateTeam(ctx context.Context, team cloud.Team, oid string) (cloud.InlineResponse2011, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2011{}, nil, err
	}
	return c.Client.TeamsApi.CreateTeam(ctx, team, oid)
}

// GetTeamById fetch a single team by its id.
func (c *cloudSdkClient) GetTeamById(ctx context.Context, oid string, tid string) (cloud.InlineResponse2011, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2011{}, nil, err
	}
	return c.Client.TeamsApi.GetTeam(ctx, oid, tid)
}

// GetTeams send request to get all teams for an organization.
func (c *cloudSdkClient) GetTeams(ctx context.Context, oid string) (cloud.InlineResponse20017, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse20017{}, nil, err
	}
	return c.Client.TeamsApi.GetTeams(ctx, oid)
}

// CreateEnv create an environment in a given team.
func (c *cloudSdkClient) CreateEnv(ctx context.Context, env cloud.Loadout, orgId string, teamId string) (cloud.InlineResponse2012, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2012{}, nil, err
	}
	return c.Client.LoadoutsApi.CreateLoadout(ctx, env, orgId, teamId)
}

// GetEnvById gets a single environment with its uuid.
func (c *cloudSdkClient) GetEnvById(ctx context.Context, orgId string, teamId string, envId string) (cloud.InlineResponse2012, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2012{}, nil, err
	}
	return c.Client.LoadoutsApi.GetLoadout(ctx, orgId, teamId, envId)
}

// GetEnvs gets all environments in a team.
func (c *cloudSdkClient) GetEnvs(ctx context.Context, orgId string, teamId string) (cloud.InlineResponse20016, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse20016{}, nil, err
	}
	return c.Client.LoadoutsApi.GetLoadouts(ctx, orgId, teamId)
}

// CreateDeployment creates a home or edge deployment.
func (c *cloudSdkClient) CreateDeployment(ctx context.Context, deployment cloud.Deployment, orgId, teamId, envId string) (cloud.InlineResponse2001, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2001{}, nil, err
	}
	return c.Client.DeploymentsApi.CreateDeployment(ctx, deployment, orgId, teamId, envId)
}

// AddBeforeExecuteFunc adds functions that should be executed before each client request
// You can for example add a function that changes the baseurl here or set new headers.
func (c *cloudSdkClient) AddBeforeExecuteFunc(beforeExecuteFunc ...func(*cloud.APIClient, *cloud.Configuration) error) {
	c.beforeExecute = append(c.beforeExecute, beforeExecuteFunc...)
}

// runBeforeExecute will call all the functions in beforeExecute and return an error if any of them fails.
func (c *cloudSdkClient) runBeforeExecute() error {
	for _, f := range c.beforeExecute {
		err := f(c.Client, c.Config)
		if err != nil {
			return err
		}
	}
	return nil
}

// ExtractErrorMessage returns the body error message from our response.
func ExtractErrorMessage(err error) string {
	if genericError, ok := err.(cloud.GenericSwaggerError); ok {
		return string(genericError.Body())
	}
	return err.Error()
}
