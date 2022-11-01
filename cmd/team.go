package cmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

func NewTeamCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(teams).WithFlagAdder(true, addOrgFlag).WithCommands(NewCreateTeamCmd(client), NewFetchTeamCmd(client))
}
