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

func NewDeleteDeploymentCmd(factory internal.CloudFactory) *cobra.Command {
	depObject := DepObject{}

	return NewDeleteBaseCmd(factory, &depObject, Dep)
}

type DepObject struct {
	DepResponse *cloud.Deployment
}

func (d *DepObject) Delete(ctx context.Context, client internal.CloudClient, config internal.UserConfig, id string, f *pflag.FlagSet) error {
	depResponse, err := validateFlagsAndDeleteDeployment(ctx, client, config, id, f)
	if err != nil {
		return err
	}

	d.DepResponse = depResponse

	return nil
}

func (d *DepObject) GetUID() string {
	if d.DepResponse == nil {
		return ""
	}

	return d.DepResponse.UID
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

	deleteFlag, err := f.GetBool(deleteCmd)
	if err != nil {
		return nil, err
	}

	localVar := &cloud.DeploymentsApiDestroyDeploymentOpts{
		Delete: optional.NewBool(deleteFlag),
		Purge:  optional.NewBool(purgeFlag),
	}

	return deleteDeployment(ctx, client, deploymentFlags.OrgID, deploymentFlags.TeamID, deploymentFlags.EnvID, deploymentID, localVar)
}

func deleteDeployment(ctx context.Context, client internal.CloudClient, orgID, teamID, envID, id string, deploymentQuery *cloud.DeploymentsApiDestroyDeploymentOpts) (*cloud.Deployment, error) {
	deleteResponse, resp, err := client.DestroyDeployment(ctx, orgID, teamID, envID, id, deploymentQuery)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		return nil, fmt.Errorf("returned error code %d while deleting deployment", resp.StatusCode)
	}

	return deleteResponse.Payload, nil
}

func deleteDeploymentFlag(f *pflag.FlagSet) {
	f.BoolP(deleteCmd, "d", false, "mark deployment as deleted")
	f.BoolP(purge, "p", false, "purge deployment from storage")
}
