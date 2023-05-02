package cloudcmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
)

func NewDeleteDeployment(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(delete).
		WithFlagAdder(false, deleteDeploymentFlag).
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			deployment, err := validateFlagsAndDeleteDeployment(ctx, factory.Client, factory.Config, args[0], cmd.Flags())
			if err != nil {
				return err
			}

			cmd.Printf("deleted %s\n successfully", deployment.UID)

			return nil
		})
}

func validateFlagsAndDeleteDeployment(ctx context.Context, client internal.CloudClient, config internal.UserConfig, deploymentID string, f *pflag.FlagSet) (*cloud.Deployment, error) {
	deploymentFlags, err := validateCommonDeploymentFlags(config)
	if err != nil {
		return nil, err
	}

	purgeFlag, err := f.GetBool(purge)
	if err != nil {
		return nil, err
	}

	deleteFlag, err := f.GetBool(delete)
	if err != nil {
		return nil, err
	}

	return deleteDeployment(ctx, client, deploymentFlags.OrgID, deploymentFlags.TeamID, deploymentFlags.EnvID, deploymentID, deleteFlag, purgeFlag)
}

func deleteDeployment(ctx context.Context, client internal.CloudClient, orgID, teamID, envID, id string, delete, purge bool) (*cloud.Deployment, error) {
	localVar := cloud.DeploymentsApiDestroyDeploymentOpts{
		Delete: optional.NewBool(delete),
		Purge:  optional.NewBool(purge),
	}

	deleteResponse, resp, err := client.DestroyDeployment(ctx, orgID, teamID, envID, id, &localVar)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		return nil, fmt.Errorf("returned error code %d while deleting deployment", resp.StatusCode)
	}

	return deleteResponse.Payload, nil
}

func deleteDeploymentFlag(f *pflag.FlagSet) {
	f.String(delete, "", "mark deployment as deleted")
	f.String(purge, "", "purge deployment from storage")
}
