package gatewaycmd

import (
	"github.com/TykTechnologies/tykctl/gatewaycmd/reload"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

// NewGatewayCommand  creates the gateway service parent command.
func NewGatewayCommand(factory internal.ApimFactory) *cobra.Command {
	return internal.NewCmd(gateway).
		WithCommands(
			reload.NewReloadCmd(factory))

}
