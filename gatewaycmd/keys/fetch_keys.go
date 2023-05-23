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

func NewFetchKeysCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Fetch).
		WithFlagAdder(false, shared.AddOutPutFlags).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())
			return shared.AddGatewayServers(apimClient.Client.GetConfig())
		}).
		MaximumArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			outPut, err := cmd.Flags().GetString(shared.OutPut)
			if err != nil {
				return err
			}

			if len(args) == 1 {
				sessionState, err := fetchKeyByID(ctx, apimClient.Client.KeysAPI, args[0])
				if err != nil {
					return err
				}
				return internal.PrintJSON(sessionState)
			}

			keys, err := fetchKeys(ctx, apimClient.Client.KeysAPI)
			if err != nil {
				return err
			}

			if outPut == shared.Table {
				internal.Printable(createKeysHeadersAndRows(keys.Keys))
				return nil
			}

			return internal.PrintJSON(keys)
		})
}

func fetchKeys(ctx context.Context, keysAPI apim.KeysAPI) (*apim.ApiAllKeys, error) {
	keys, resp, err := keysAPI.ListKeysExecute(keysAPI.ListKeys(ctx))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return keys, nil
}

func fetchKeyByID(ctx context.Context, keysAPI apim.KeysAPI, id string) (*apim.SessionState, error) {
	sessionState, resp, err := keysAPI.GetKeyExecute(keysAPI.GetKey(ctx, id))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return sessionState, nil
}

// createKeysHeadersAndRows create headers and rows to be used in creating keys table.
func createKeysHeadersAndRows(keys []string) ([]string, [][]string) {
	header := []string{"Key"}
	rows := make([][]string, 0)

	for _, key := range keys {
		row := []string{
			key,
		}
		rows = append(rows, row)
	}

	return header, rows
}
