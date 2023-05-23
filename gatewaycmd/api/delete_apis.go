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

func NewDeleteApisCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Delete).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())

			err := shared.AddGatewayServers(apimClient.Client.GetConfig())
			if err != nil {
				return err
			}

			return nil
		}).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			apiStatusMessage, err := deleteAPI(cmd.Context(), apimClient.Client.APIsAPI, args[0])
			if err != nil {
				return err
			}

			return internal.PrintJSON(apiStatusMessage)
		})
}

func deleteAPI(ctx context.Context, apisAPI apim.APIsAPI, apiID string) (*apim.ApiStatusMessage, error) {
	status, resp, err := apisAPI.DeleteApiExecute(apisAPI.DeleteApi(ctx, apiID))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return status, nil
}
