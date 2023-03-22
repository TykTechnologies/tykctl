package api

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

func NewCreateAPICmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Create).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())
			shared.AddGatewayServers(apimClient.Client.GetConfig())

			return nil
		}).
		NoArgs(func(ctx context.Context, command cobra.Command) error {
			return nil
		})
}
