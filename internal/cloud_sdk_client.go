package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

var (
	_               CloudClient = (*CloudSdkClient)(nil)
	zonePath                    = "/api/deployments/zones"
	userInfoPath                = "/api/users/whoami"
	orgInfoPath                 = "api/organisations/"
	applicationJSON             = "application/json"
	contentType                 = "Content-Type"
)

// CloudSdkClient should implement CloudClient as it will be used to make request to Ara.
type CloudSdkClient struct {
	Client *cloud.APIClient
	Config *cloud.Configuration
	// beforeExecute will store function you need called before each method runs .
	beforeExecute      []func(*cloud.APIClient, *cloud.Configuration) error
	dashboardClient    *resty.Client
	beforeRestyExecute []func(*resty.Client) error
}

// NewCloudSdkClient creates a new CloudSdkClient to make sure that a client is never nil.
func NewCloudSdkClient(conf *cloud.Configuration) *CloudSdkClient {
	client := cloud.NewAPIClient(conf)

	return &CloudSdkClient{
		Client:          client,
		beforeExecute:   nil,
		dashboardClient: resty.New(),
		Config:          conf,
	}
}

// GetOrgs a users organizations from Ara.
func (c *CloudSdkClient) GetOrgs(ctx context.Context) (cloud.InlineResponse20014, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse20014{}, nil, err
	}

	return c.Client.OrganisationsApi.GetOrgs(ctx)
}

// GetOrgByID will get a single organization using its id.
func (c *CloudSdkClient) GetOrgByID(ctx context.Context, oid string) (cloud.InlineResponse2005, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2005{}, nil, err
	}

	return c.Client.OrganisationsApi.GetOrg(ctx, oid)
}

// CreateTeam create a team in an organization.
func (c *CloudSdkClient) CreateTeam(ctx context.Context, team cloud.Team, oid string) (cloud.InlineResponse2011, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2011{}, nil, err
	}

	return c.Client.TeamsApi.CreateTeam(ctx, team, oid)
}

// GetTeamByID fetch a single team by its id.
func (c *CloudSdkClient) GetTeamByID(ctx context.Context, oid, tid string) (cloud.InlineResponse2011, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2011{}, nil, err
	}

	return c.Client.TeamsApi.GetTeam(ctx, oid, tid)
}

// GetTeams send request to get all teams for an organization.
func (c *CloudSdkClient) GetTeams(ctx context.Context, oid string) (cloud.InlineResponse20017, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse20017{}, nil, err
	}

	return c.Client.TeamsApi.GetTeams(ctx, oid)
}

// UpdateTeam updates the name on the team that is already created.
func (c *CloudSdkClient) UpdateTeam(ctx context.Context, team cloud.Team, orgID, teamID string) (cloud.InlineResponse2011, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2011{}, nil, err
	}

	return c.Client.TeamsApi.UpdateTeam(ctx, team, orgID, teamID)
}

// CreateEnv create an environment in a given team.
func (c *CloudSdkClient) CreateEnv(ctx context.Context, env cloud.Loadout, orgID, teamID string) (cloud.InlineResponse2012, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2012{}, nil, err
	}

	return c.Client.LoadoutsApi.CreateLoadout(ctx, env, orgID, teamID)
}

// GetEnvByID gets a single environment with its uuid.
func (c *CloudSdkClient) GetEnvByID(ctx context.Context, orgID, teamID, envID string) (cloud.InlineResponse2012, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2012{}, nil, err
	}

	return c.Client.LoadoutsApi.GetLoadout(ctx, orgID, teamID, envID)
}

// GetEnvs gets all environments in a team.
func (c *CloudSdkClient) GetEnvs(ctx context.Context, orgID, teamID string) (cloud.InlineResponse20016, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse20016{}, nil, err
	}

	return c.Client.LoadoutsApi.GetLoadouts(ctx, orgID, teamID)
}

// CreateDeployment creates a home or edge deployment.
func (c *CloudSdkClient) CreateDeployment(ctx context.Context, deployment cloud.Deployment, orgID, teamID, envID string) (cloud.InlineResponse2001, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2001{}, nil, err
	}

	return c.Client.DeploymentsApi.CreateDeployment(ctx, deployment, orgID, teamID, envID)
}

func (c *CloudSdkClient) GetEnvDeployments(ctx context.Context, orgID, teamID, envID string) (cloud.InlineResponse200, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse200{}, nil, err
	}

	return c.Client.LoadoutsApi.GetDeploymentsForLoadout(ctx, orgID, teamID, envID)
}

func (c *CloudSdkClient) GetDeploymentByID(ctx context.Context, orgID, teamID, envID, id string, localVarOptionals *cloud.DeploymentsApiGetDeploymentOpts) (cloud.InlineResponse2001, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2001{}, nil, err
	}

	return c.Client.DeploymentsApi.GetDeployment(ctx, orgID, teamID, envID, id, localVarOptionals)
}

func (c *CloudSdkClient) StartDeployment(ctx context.Context, orgID, teamID, envID, id string) (cloud.InlineResponse2001, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2001{}, nil, err
	}

	return c.Client.DeploymentsApi.StartDeployment(ctx, orgID, teamID, envID, id)
}

func (c *CloudSdkClient) RestartDeployment(ctx context.Context, oid, tid, lid, id string) (cloud.InlineResponse2001, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.InlineResponse2001{}, nil, err
	}

	return c.Client.DeploymentsApi.RestartDeployment(ctx, oid, tid, lid, id)
}

func (c *CloudSdkClient) GetDeploymentStatus(ctx context.Context, orgID, teamID, envID, deploymentID string) (cloud.Payload, *http.Response, error) {
	err := c.runBeforeExecute()
	if err != nil {
		return cloud.Payload{}, nil, err
	}

	return c.Client.DeploymentsApi.GetDeploymentStatus(ctx, orgID, teamID, envID, deploymentID)
}

func (c *CloudSdkClient) GetDeploymentZones(ctx context.Context) (*ZoneResponse, *resty.Response, error) {
	err := c.runBeforeRestyExecute()
	if err != nil {
		return nil, nil, err
	}
	var zoneResponse ZoneResponse
	request := c.dashboardClient.R().SetHeader(contentType, applicationJSON).SetResult(&zoneResponse)
	request.SetContext(ctx)
	response, err := request.Get(zonePath)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode() != 200 {
		return nil, nil, NewGenericHTTPError(response.String())
	}

	return &zoneResponse, response, nil
}

// GetUserInfo will get userRole, orgId and the team the user belongs to from tyk cloud.
func (c *CloudSdkClient) GetUserInfo(ctx context.Context) (*UserInfo, *resty.Response, error) {
	err := c.runBeforeRestyExecute()
	if err != nil {
		return nil, nil, err
	}
	var userInfo UserInfo
	request := c.dashboardClient.R().SetHeader(contentType, applicationJSON).SetResult(&userInfo)
	request.SetContext(ctx)
	response, err := request.Get(userInfoPath)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode() != 200 {
		return nil, nil, NewGenericHTTPError(response.String())
	}

	return &userInfo, response, nil
}

// GetOrgInfo will fetch an organization using its id from tyk cloud.
func (c *CloudSdkClient) GetOrgInfo(ctx context.Context, orgID string) (*OrgInfo, *resty.Response, error) {
	err := c.runBeforeRestyExecute()
	if err != nil {
		return nil, nil, err
	}
	var orgInfo OrgInfo
	request := c.dashboardClient.R().SetHeader(contentType, applicationJSON).SetResult(&orgInfo)
	request.SetContext(ctx)

	path := fmt.Sprintf("%s%s", orgInfoPath, orgID)
	response, err := request.Get(path)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode() != 200 {
		return nil, nil, NewGenericHTTPError(response.String())
	}

	return &orgInfo, response, nil
}

// AddBeforeExecuteFunc adds functions that should be executed before each client request
// You can for example add a function that changes the baseurl here or set new headers.
func (c *CloudSdkClient) AddBeforeExecuteFunc(beforeExecuteFunc ...func(*cloud.APIClient, *cloud.Configuration) error) {
	c.beforeExecute = append(c.beforeExecute, beforeExecuteFunc...)
}

func (c *CloudSdkClient) AddBeforeRestyExecute(beforeExecuteRestyFunc ...func(*resty.Client) error) {
	c.beforeRestyExecute = append(c.beforeRestyExecute, beforeExecuteRestyFunc...)
}

// runBeforeExecute will call all the functions in beforeExecute and return an error if any of them fails.
func (c *CloudSdkClient) runBeforeExecute() error {
	for _, f := range c.beforeExecute {
		err := f(c.Client, c.Config)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CloudSdkClient) runBeforeRestyExecute() error {
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
	var genericOpenAPIError *apim.GenericOpenAPIError
	if errors.As(err, &genericError) {
		return string(genericError.Body())
	}

	if errors.As(err, &genericOpenAPIError) {
		return string(genericOpenAPIError.Body())
	}

	return err.Error()
}
