package cmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewDeployment(client internal.CloudClient) *cobra.Command {
	return NewCmd(deployments).
		WithFlagAdder(true, addOrgFlag).
		WithFlagAdder(true, addTeamFlag).
		WithFlagAdder(true, addEnvFlag).
		WithCommands(NewCreateDeploymentCmd(client), NewFetchDeploymentCmd(client), NewStartDeploymentCmd(client), NewDeploymentStatusCmd(client))

}

func validateCommonDeploymentFlags() (*DeploymentFlags, error) {
	var deploymentFlag DeploymentFlags
	deploymentFlag.OrgId = viper.GetString(org)
	if util.StringIsEmpty(deploymentFlag.OrgId) {
		return nil, ErrorOrgRequired
	}
	deploymentFlag.TeamId = viper.GetString(team)
	if util.StringIsEmpty(deploymentFlag.TeamId) {
		return nil, ErrorTeamRequired
	}
	deploymentFlag.EnvId = viper.GetString(env)
	if util.StringIsEmpty(deploymentFlag.EnvId) {
		return nil, ErrorEnvRequired
	}
	return &deploymentFlag, nil
}

type DeploymentFlags struct {
	OrgId  string
	TeamId string
	EnvId  string
	OutPut string
}
