package configcmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
)

const configDesc = `This is the parent command to all environment operations.

   This commands has subcommands to create new configuration files and populate them,list all config files and switch from one config file to another.
  
`

func NewConfigCmd(prompt internal.ConfigPrompt, configEntry internal.ConfigEntry, factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(internal.Config).
		WithAliases([]string{"conf"}).
		WithDescription("Parent command for all config files operations.").
		WithLongDescription(configDesc).
		WithCommands(newInitConfigCmd(prompt, configEntry, factory),
			newConfigFetchCmd(configEntry),
			newSwitchConfigCmd(prompt, configEntry))
}
