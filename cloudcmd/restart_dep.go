package cloudcmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
)

const restartDepDesc = `
This command will restart a Home or edge gateway given its uuid.

The org,team,environment where the deployment was created has to be provided.

If org,team and environment are not set we will use the default set on your config file. 

Sample usage of this command:

tykctl cloud dep restart --org=<org here> --team=<team here> --env=<environment here> --uid=<deployment id>

`

// NewRestartDeploymentCmd will restart a home deployment or edge deployment given a uuid.
func NewRestartDeploymentCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(restart).
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{
			{Name: env, Persistent: false},
			{Name: team, Persistent: false},
			{Name: org, Persistent: false},
		}).
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
	deployment, err := restartDeployment(ctx, client, deploymentFlags.OrgID, deploymentFlags.TeamID, deploymentFlags.EnvID, deploymentID)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// restartDeployment will make a http request to tyk cloud to restart the deployment.
func restartDeployment(ctx context.Context, client internal.CloudClient, orgID, teamID, envID, id string) (*cloud.Deployment, error) {
	deployment, resp, err := client.RestartDeployment(ctx, orgID, teamID, envID, id)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		return nil, ErrorRestartingDeployment
	}
	if deployment.Status != statusOK {
		return nil, errors.New(deployment.Error_)
	}
	return deployment.Payload, nil
}
