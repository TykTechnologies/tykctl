package cmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

func NewTeamCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(teams).
		WithFlagAdder(true, addOrgFlag).
		WithBindFlagOnPreRun([]BindFlag{{Name: "org", Persistent: true}}).
		WithCommands(NewCreateTeamCmd(client), NewFetchTeamCmd(client))
}
