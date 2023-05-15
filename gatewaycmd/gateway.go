package gatewaycmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/gatewaycmd/keys"
	"github.com/TykTechnologies/tykctl/gatewaycmd/reload"
	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

// NewGatewayCommand  creates the gateway service parent command.
func NewGatewayCommand(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Gateway).
		WithCommands(GatewayCommands(apimClient)...)
}

func GatewayCommands(apimClient internal.ApimClient) []*cobra.Command {
	commands := []*cobra.Command{
		reload.NewReloadCmd(apimClient),
		keys.NewKeysCmd(apimClient),
	}

	return commands
}
