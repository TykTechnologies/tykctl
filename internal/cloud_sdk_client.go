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
