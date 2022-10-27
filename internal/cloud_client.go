package internal

import (
	"context"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"net/http"
)

//go:generate mockgen -source=cloud_client.go -destination=./mocks/cloud_client.go -package=mock CloudClient
type CloudClient interface {
	GetOrgs(ctx context.Context) (cloud.InlineResponse20014, *http.Response, error)
}
