package internal

import (
	"context"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"net/http"
)

//go:generate mockgen -source=cloud_client.go -destination=./mocks/cloud_client.go -package=mock CloudClient
type CloudClient interface {
	GetOrgs(ctx context.Context) (cloud.InlineResponse20014, *http.Response, error)
	GetOrgById(ctx context.Context, oid string) (cloud.InlineResponse2005, *http.Response, error)
	CreateTeam(ctx context.Context, team cloud.Team, oid string) (cloud.InlineResponse2011, *http.Response, error)
}
