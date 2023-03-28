package gatewaycmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/gatewaycmd/api"
	"github.com/TykTechnologies/tykctl/gatewaycmd/reload"
	"github.com/TykTechnologies/tykctl/gatewaycmd/templates"

	"github.com/TykTechnologies/tykctl/gatewaycmd/keys"
	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"

	"github.com/TykTechnologies/tykctl/internal"
)

// NewGatewayCommand  creates the gateway service parent command.
func NewGatewayCommand(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Gateway).
		WithCommands(
			reload.NewReloadCmd(apimClient),
			api.NewApisCmd(apimClient),
			templates.NewTemplateCmd(),
			keys.NewKeysCmd(apimClient),
		)
}
