package gatewaycmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/gatewaycmd/reload"
	"github.com/TykTechnologies/tykctl/internal"
)

// NewGatewayCommand  creates the gateway service parent command.
func NewGatewayCommand(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(gateway).
		WithCommands(
			reload.NewReloadCmd(apimClient))
}
