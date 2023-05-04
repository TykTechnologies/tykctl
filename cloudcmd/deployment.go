package cloudcmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
)

const deploymentDesc = `This is the parent command to all deployment operation, such as creating teams and fetching teams.

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
		WithDescription("Parent command for all actions you can perform in a deployment.").
		WithFlagAdder(true, addOrgFlag).
		WithFlagAdder(true, addTeamFlag).
		WithFlagAdder(true, addEnvFlag).
		WithCommands(
			NewCreateDeploymentCmd(factory),
			NewFetchDeploymentCmd(factory),
			NewStartDeploymentCmd(factory),
			NewDeploymentStatusCmd(factory),
			NewRestartDeploymentCmd(factory),
			NewUpdateDeployment(factory),
			NewDeleteDeploymentCmd(factory),
		)
}

func validateCommonDeploymentFlags(config internal.UserConfig) (*DeploymentFlags, error) {
	envFlags, err := validateCommonEnvFlags(config)
	if err != nil {
		return nil, err
	}

	var deploymentFlag DeploymentFlags
	deploymentFlag.OrgID = envFlags.OrgID
	deploymentFlag.TeamID = envFlags.TeamID
	deploymentFlag.EnvID = config.GetCurrentUserEnv()

	if util.StringIsEmpty(deploymentFlag.EnvID) {
		return nil, ErrorEnvRequired
	}

	return &deploymentFlag, nil
}

type DeploymentFlags struct {
	EnvFlags
	EnvID string
}
