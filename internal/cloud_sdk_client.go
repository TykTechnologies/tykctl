package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/go-resty/resty/v2"
	"net/http"
)

var (
	_               CloudClient = (*cloudSdkClient)(nil)
	zonePath                    = "/api/deployments/zones"
	userInfoPath                = "/api/users/whoami"
	orgInfoPath                 = "api/organisations/"
	applicationJson             = "application/json"
	contentType                 = "Content-Type"
)

// CloudSdkClient should implement CloudClient as it will be used to make request to Ara.
type cloudSdkClient struct {
	Client *cloud.APIClient
	Config *cloud.Configuration
	//beforeExecute will store function you need called before each method runs .
	beforeExecute      []func(*cloud.APIClient, *cloud.Configuration) error
	dashboardClient    *resty.Client
	beforeRestyExecute []func(*resty.Client) error
}

// NewCloudSdkClient creates a new cloudSdkClient to make sure that a client is never nil.
func NewCloudSdkClient(conf *cloud.Configuration) *cloudSdkClient {
	client := cloud.NewAPIClient(conf)
	return &cloudSdkClient{
		Client:          client,
		beforeExecute:   nil,
		dashboardClient: resty.New(),
		Config:          conf,
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

// UpdateTeam updates the name on the team that is already created.
func (c *cloudSdkClient) UpdateTeam(ctx context.Context, team cloud.Team, orgId string, teamId string) (cloud.InlineResponse2011, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2011{}, nil, err
	}
	return c.Client.TeamsApi.UpdateTeam(ctx, team, orgId, teamId)
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
func (c *cloudSdkClient) GetEnvDeployments(ctx context.Context, orgId string, teamId string, envId string) (cloud.InlineResponse200, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse200{}, nil, err
	}
	return c.Client.LoadoutsApi.GetDeploymentsForLoadout(ctx, orgId, teamId, envId)
}
func (c *cloudSdkClient) GetDeploymentById(ctx context.Context, orgId string, teamId string, envId string, id string, localVarOptionals *cloud.DeploymentsApiGetDeploymentOpts) (cloud.InlineResponse2001, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2001{}, nil, err
	}
	return c.Client.DeploymentsApi.GetDeployment(ctx, orgId, teamId, envId, id, localVarOptionals)
}
func (c *cloudSdkClient) StartDeployment(ctx context.Context, orgId, teamId, envId, id string) (cloud.InlineResponse2001, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2001{}, nil, err
	}
	return c.Client.DeploymentsApi.StartDeployment(ctx, orgId, teamId, envId, id)
}
func (c *cloudSdkClient) GetDeploymentStatus(ctx context.Context, orgId string, teamId string, envId string, deploymentId string) (cloud.Payload, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.Payload{}, nil, err
	}
	return c.Client.DeploymentsApi.GetDeploymentStatus(ctx, orgId, teamId, envId, deploymentId)
}

func (c *cloudSdkClient) GetDeploymentZones(ctx context.Context) (*ZoneResponse, *resty.Response, error) {
	err := c.runBeforeRestyExecute()
	if err != nil {
		return nil, nil, err
	}
	var zoneResponse ZoneResponse
	request := c.dashboardClient.R().SetHeader(contentType, applicationJson).SetResult(&zoneResponse)
	request.SetContext(ctx)
	response, err := request.Get(zonePath)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode() != 200 {
		return nil, nil, NewGenericHttpError(response.String())
	}
	return &zoneResponse, response, nil
}

// GetUserInfo will get userRole, orgId and the team the user belongs to from tyk cloud.
func (c *cloudSdkClient) GetUserInfo(ctx context.Context) (*UserInfo, *resty.Response, error) {
	err := c.runBeforeRestyExecute()
	if err != nil {
		return nil, nil, err
	}
	var userInfo UserInfo
	request := c.dashboardClient.R().SetHeader(contentType, applicationJson).SetResult(&userInfo)
	request.SetContext(ctx)
	response, err := request.Get(userInfoPath)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode() != 200 {
		return nil, nil, NewGenericHttpError(response.String())
	}
	return &userInfo, response, nil
}

// GetOrgInfo will fetch an organization using its id from tyk cloud.
func (c *cloudSdkClient) GetOrgInfo(ctx context.Context, orgId string) (*OrgInfo, *resty.Response, error) {
	err := c.runBeforeRestyExecute()
	if err != nil {
		return nil, nil, err
	}
	var orgInfo OrgInfo
	request := c.dashboardClient.R().SetHeader(contentType, applicationJson).SetResult(&orgInfo)
	request.SetContext(ctx)
	path := fmt.Sprintf("%s%s", orgInfoPath, orgId)
	response, err := request.Get(path)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode() != 200 {
		return nil, nil, NewGenericHttpError(response.String())
	}
	return &orgInfo, response, nil
}

// AddBeforeExecuteFunc adds functions that should be executed before each client request
// You can for example add a function that changes the baseurl here or set new headers.
func (c *cloudSdkClient) AddBeforeExecuteFunc(beforeExecuteFunc ...func(*cloud.APIClient, *cloud.Configuration) error) {
	c.beforeExecute = append(c.beforeExecute, beforeExecuteFunc...)
}
func (c *cloudSdkClient) AddBeforeRestyExecute(beforeExecuteRestyFunc ...func(*resty.Client) error) {
	c.beforeRestyExecute = append(c.beforeRestyExecute, beforeExecuteRestyFunc...)
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

func (c *cloudSdkClient) runBeforeRestyExecute() error {
	c.dashboardClient.OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
		for _, f := range c.beforeRestyExecute {
			err := f(client)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

// ExtractErrorMessage returns the body error message from our response.
func ExtractErrorMessage(err error) string {
	var genericError *cloud.GenericSwaggerError
	if errors.As(err, &genericError) {
		return string(genericError.Body())
	}
	return err.Error()
}
