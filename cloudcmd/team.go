package cloudcmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
)

const teamDesc = `
This command is the parent command to all team operations.

The supported commands are :

1. tykctl cloud teams fetch.

2. tykctl cloud teams create -create a team in an org.

All subcommands require an org id.If it is not passed we use the default one in the config file.

To set the default org run :

tykctl cloud init
`

func NewTeamCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(teams).
		WithAliases([]string{team}).
		WithFlagAdder(true, addOrgFlag).
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: org, Persistent: true}}).
		WithLongDescription(teamDesc).
		WithDescription("Parent command for all actions you can perform in a team.").
		WithBindFlagOnPreRun([]internal.BindFlag{{Name: org, Persistent: true}}).
		WithCommands(NewCreateTeamCmd(factory), NewFetchTeamCmd(factory))
}
