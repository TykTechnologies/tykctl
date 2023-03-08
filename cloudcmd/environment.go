package cloudcmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
)

const environmentDesc = `This is the parent command to all environment operations.

   Note: All subcommands for this command must pass a --team  and --org flag command.
  
`

func NewEnvironmentCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(environments).
		WithAliases([]string{env}).
		WithFlagAdder(true, addOrgFlag).
		WithLongDescription(environmentDesc).
		WithDescription("parent command to all environment operations").
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: org, Persistent: true}, {Name: team, Persistent: true}}).
		WithFlagAdder(true, addTeamFlag).
		WithCommands(
			NewCreateEnvironmentCmd(factory),
			NewFetchEnvironmentCmd(factory),
		)
}
