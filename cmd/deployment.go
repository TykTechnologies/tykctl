package cmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

func NewDeployment(client internal.CloudClient) *cobra.Command {
	return NewCmd(deployments).
		WithFlagAdder(true, addOrgFlag).
		WithFlagAdder(true, addTeamFlag).
		WithFlagAdder(true, addEnvFlag).
		WithCommands(NewCreateDeploymentCmd(client), NewFetchDeploymentCmd(client))

}
