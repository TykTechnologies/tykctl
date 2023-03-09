package api

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

const (
	apis = "apis"
)

func NewCreateApiCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(apis).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())
			shared.AddGatewayServers(apimClient.Client.GetConfig())

			return nil
		}).
		NoArgs(func(ctx context.Context, command cobra.Command) error {
			return nil
		})
}

func CreateApi(ctx context.Context, aPIsAPI apim.APIsAPI) {
}
