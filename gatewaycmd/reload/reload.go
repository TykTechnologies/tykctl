package reload

import (
	"context"
	"errors"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

const (
	reload = "reload"
	block  = "block"
	group  = "group"
)

func NewReloadCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(reload).
		WithFlagAdder(false, reloadFlags).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())
			shared.AddGatewayServers(apimClient.Client.GetConfig())
			return nil
		}).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			block, err := cmd.Flags().GetBool(block)
			if err != nil {
				return err
			}

			group, err := cmd.Flags().GetBool(group)
			if err != nil {
				return err
			}

			_, err = reloadGateway(cmd.Context(), apimClient.Client.HotReloadAPI, block, group)
			if err != nil {
				return err
			}
			cmd.Println("gateway reloaded successfully")

			return nil
		})
}

func reloadGateway(ctx context.Context, hotReload apim.HotReloadAPI, block, group bool) (*apim.ApiStatusMessage, error) {
	if group {
		return reloadGroup(ctx, hotReload)
	}

	return reloadSingleNode(ctx, hotReload, block)
}

func reloadSingleNode(ctx context.Context, hotReload apim.HotReloadAPI, block bool) (*apim.ApiStatusMessage, error) {
	status, resp, err := hotReload.HotReloadExecute(hotReload.HotReload(ctx).Block(block))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return status, nil
}

func reloadGroup(ctx context.Context, hotReload apim.HotReloadAPI) (*apim.ApiStatusMessage, error) {
	status, resp, err := hotReload.HotReloadGroupExecute(hotReload.HotReloadGroup(ctx))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return status, nil
}

func reloadFlags(f *pflag.FlagSet) {
	f.BoolP(block, "b", false, "Block a response until the reload is performed. This can be useful in scripting environments like CI/CD workflows.")
	f.BoolP(group, "g", false, "To reload a whole group of Tyk nodes")
}
