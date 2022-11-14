package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
	"net/http"
)

const deployDesc = `
This command will deploy a Home or edge gateway given its id.
Note: You need to first create the Home or edge gateway before you can deploy it.
Use tykctl cloud deployments create to create the deployment.
The org,team,environment where the deployment was created has to be provided.
If org,team and environment are not set we will use the default set on your config file. 
You must also provide the uuid of the deployment you want to deploy.
to get the uuid run : tykctl cloud deployments fetch
Sample usage of this command:
tykctl cloud deployments deploy --org=<org here> --team=<team here> --env=<environment here> --uid=<deployment id>
`

func NewStartDeploymentCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(deploy).
		WithBindFlagOnPreRun([]BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		WithLongDescription(deployDesc).
		WithDescription("deploy a home or edge gateway deployment given its id").
		WithExample("tykctl cloud deployments deploy --org=<org here> --team=<team here> --env=<environment here> --uid=<deployment id>").
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			deployment, err := validateFlagsAndStartDeployment(ctx, client, args[0])
			if err != nil {
				cmd.Println(err)
				return err
			}
			cmd.Println(fmt.Sprintf("Deploying %s...", deployment.UID))
			return nil
		})
}

func validateFlagsAndStartDeployment(ctx context.Context, client internal.CloudClient, deploymentID string) (*cloud.Deployment, error) {
	deploymentFlags, err := validateCommonDeploymentFlags()
	if err != nil {
		return nil, err
	}
	deployment, err := StartDeployment(ctx, client, deploymentFlags.OrgId, deploymentFlags.TeamId, deploymentFlags.EnvId, deploymentID)
	if err != nil {
		return deployment, err
	}
	return deployment, nil
}

func StartDeployment(ctx context.Context, client internal.CloudClient, orgId, teamId, envId, id string) (*cloud.Deployment, error) {
	startResponse, resp, err := client.StartDeployment(ctx, orgId, teamId, envId, id)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		return nil, ErrorStartingDeployment
	}
	if startResponse.Status != statusOK {
		return nil, errors.New(startResponse.Error_)
	}
	return startResponse.Payload, nil
}
