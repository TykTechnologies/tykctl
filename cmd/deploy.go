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

func NewStartDeploymentCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(deploy).
		WithBindFlagOnPreRun([]BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
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
