package cloudcmd

import (
	"context"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

const restartDepDesc = `
This command will restart a Home or edge gateway given its uuid.

The org,team,environment where the deployment was created has to be provided.

If org,team and environment are not set we will use the default set on your config file. 

Sample usage of this command:

tykctl cloud dep restart --org=<org here> --team=<team here> --env=<environment here> --uid=<deployment id>

`

func NewRestartDeploymentCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(restart).
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		WithLongDescription(restartDepDesc).
		WithDescription("restart a home or edge gateway deployment given its uuid").
		WithExample("tykctl cloud dep restart <deployment id> --org=<org here> --team=<team here> --env=<environment here> ").
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			deployment, err := validateFlagsAndRestartDeployment(ctx, factory.Client, factory.Config, args[0])
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			cmd.Println(fmt.Sprintf("Successfully restarted deployment %s...", deployment.UID))
			return nil
		})
}
func validateFlagsAndRestartDeployment(ctx context.Context, client internal.CloudClient, config internal.UserConfig, deploymentID string) (*cloud.Deployment, error) {
	deploymentFlags, err := validateCommonDeploymentFlags(config)
	if err != nil {
		return nil, err
	}
	deployment, _, err := client.RestartDeployment(ctx, deploymentFlags.OrgId, deploymentFlags.TeamId, deploymentFlags.EnvId, deploymentID)
	if err != nil {
		return nil, err
	}
	return deployment.Payload, nil
}
