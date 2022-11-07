package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	ErrorControlPlaneRequired = errors.New("error control plane to link the gateway to is required")
)

func NewCreateEdgeDeployment(client internal.CloudClient) *cobra.Command {
	return NewCmd(edge).
		WithBindFlagOnPreRun([]BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		WithFlagAdder(false, addEdgeDeploymentFlag).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			deployment, err := validateEdgeDeploymentFlagAndCreate(cmd.Context(), client, cmd.Flags())
			if err != nil {
				cmd.Println(err)
				return err
			}
			cmd.Println(fmt.Sprintf("deployment %s created successfully", deployment.UID))
			return nil
		})
}

func addEdgeDeploymentFlag(f *pflag.FlagSet) {
	f.String(linkedControlPlane, "", "control plane to link the edge gateway to.")
}

func validateEdgeDeploymentFlagAndCreate(ctx context.Context, client internal.CloudClient, f *pflag.FlagSet) (*cloud.Deployment, error) {
	deployment, err := extractCommonDeploymentFlags(f)
	if err != nil {
		return nil, err
	}
	controlPlane, err := f.GetString(linkedControlPlane)
	if err != nil {
		return nil, err
	}
	if util.StringIsEmpty(controlPlane) {
		return nil, ErrorControlPlaneRequired
	}
	deployment.LinkedDeployments["LinkedMDCBID"] = controlPlane
	deployment.Kind = gateway
	return CreateDeployment(ctx, client, *deployment, deployment.OID, deployment.TID, deployment.LID)
}
