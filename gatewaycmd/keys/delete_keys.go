package keys

import (
	"context"
	"errors"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

func NewDeleteKeyCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Delete).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())
			shared.AddGatewayServers(apimClient.Client.GetConfig())
			return nil
		}).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			_, err := deleteKeyByID(ctx, apimClient.Client.KeysAPI, args[0])
			if err != nil {
				return err
			}

			cmd.Printf("key %s deleted successfully", args[0])

			return nil
		})
}

// deleteKeyByID will delete a key from a gateway instance given the id.
func deleteKeyByID(ctx context.Context, keysAPI apim.KeysAPI, id string) (*apim.ApiStatusMessage, error) {
	apiStatus, resp, err := keysAPI.DeleteKeyExecute(keysAPI.DeleteKey(ctx, id))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return apiStatus, nil
}
