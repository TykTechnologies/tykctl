package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

func NewFetchApisCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Fetch).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())
			return shared.AddGatewayServers(apimClient.Client.GetConfig())
		}).
		MaximumArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			if len(args) == 0 {
				apis, err := getApis(ctx, apimClient.Client.APIsAPI)
				if err != nil {
					return err
				}

				return internal.PrintJSON(apis)
			}

			api, err := getApisByID(ctx, apimClient.Client.APIsAPI, args[0])
			if err != nil {
				return err
			}

			return internal.PrintJSON(api)
		})
}

func getApis(ctx context.Context, apisAPI apim.APIsAPI) ([]apim.APIDefinition, error) {
	apis, resp, err := apisAPI.ListApisExecute(apisAPI.ListApis(ctx))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return apis, nil
}

func getApisByID(ctx context.Context, apisAPI apim.APIsAPI, apiID string) (*apim.APIDefinition, error) {
	api, resp, err := apisAPI.GetApiExecute(apisAPI.GetApi(ctx, apiID))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return api, nil
}
