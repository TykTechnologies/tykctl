package cloudcmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

const environmentDesc = `This is the parent command to all environment operations.

   Note: All subcommands for this command must pass a --team  and --org flag command.
  
`

func NewEnvironmentCmd(client internal.CloudClient) *cobra.Command {
	return internal.NewCmd(environments).
		WithFlagAdder(true, addOrgFlag).
		WithLongDescription(environmentDesc).
		WithDescription("parent command to all environment operations").
		WithBindFlagOnPreRun([]internal.BindFlag{{Name: "org", Persistent: true}, {Name: "team", Persistent: true}}).
		WithFlagAdder(true, addTeamFlag).
		WithCommands(
			NewCreateEnvironmentCmd(client),
			NewFetchEnvironmentCmd(client),
		)
}
