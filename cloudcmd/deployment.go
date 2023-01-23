package cloudcmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const deploymentDesc = `This is the parent command to all deployment operation.Such as creating teams and fetching teams.

  Note: All subcommands for this command must pass an --env, --team  and --org flag command.

The supported commands are :
1. tykctl cloud deployments fetch 

2. tykctl cloud deployments create home -create a team in an org.

3. tykctl cloud deployments create edge -create a team in an org.

4. tykctl cloud deployments status.

5. tykctl cloud deployments deploy [deploymentID].

5. tykctl cloud deployments fetch [deploymentID].

`

func NewDeployment(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(deployments).
		WithAliases([]string{dep}).
		WithLongDescription(deploymentDesc).
		WithDescription("parent command for all action you can perform in a deployment.").
		WithFlagAdder(true, addOrgFlag).
		WithFlagAdder(true, addTeamFlag).
		WithFlagAdder(true, addEnvFlag).
		WithCommands(
			NewCreateDeploymentCmd(factory),
			NewFetchDeploymentCmd(factory),
			NewStartDeploymentCmd(factory),
			NewDeploymentStatusCmd(factory))

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
