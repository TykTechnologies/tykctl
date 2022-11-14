package cloudcmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

func NewEnvironmentCmd(client internal.CloudClient) *cobra.Command {
	return internal.NewCmd(environments).
		WithFlagAdder(true, addOrgFlag).
		WithBindFlagOnPreRun([]internal.BindFlag{{Name: "org", Persistent: true}, {Name: "team", Persistent: true}}).
		WithFlagAdder(true, addTeamFlag).
		WithCommands(
			NewCreateEnvironmentCmd(client),
			NewFetchEnvironmentCmd(client),
		)
}
