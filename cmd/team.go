package cmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

const teamDesc = `
This command is the parent command to all team operations.
The supported commands are :
1. tykctl cloud teams fetch 
2. tykctl cloud teams create -create a team in an org.
All subcommands require an org id.If it is not passed we use the default one in the config file.
To set the default org run :
tykctl cloud init
`

func NewTeamCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(teams).
		WithFlagAdder(true, addOrgFlag).
		WithLongDescription(teamDesc).
		WithDescription("parent command for all action you can perform in a team.").
		WithBindFlagOnPreRun([]BindFlag{{Name: org, Persistent: true}}).
		WithCommands(NewCreateTeamCmd(client), NewFetchTeamCmd(client))
}
