package cloudcmd

import (
	"context"
	"errors"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log"
)

const createEdgeDeploymentDesc = ` 
will create an Edge Gateway.

NOTE: This does not deploy the deployment it just create it.You can use the deploy flag to deploy after create.You can also use the deploy command to deploy the created deployment.

You must pass the organization,team,zone and environment you want deploy this deployment.

If you do not pass the org,zone or environment we will use the ones on your config file.You can set the default org,team and environment by running:

		tykctl cloud init

Sample usage for this command

 		tykctl cloud deployments create edge --name="test deployment"
`

var (
	ErrorControlPlaneRequired = errors.New("error control plane to link the gateway to is required")
)

func NewCreateEdgeDeployment(client internal.CloudClient) *cobra.Command {
	return internal.NewCmd(edge).
		WithLongDescription(createEdgeDeploymentDesc).
		WithDescription("will create the edge gateway in a given environment").
		WithExample("tykctl cloud deployments create edge --name='test deployment'").
		WithBindFlagOnPreRun([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		WithFlagAdder(false, addEdgeDeploymentFlag).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			_, err := validateEdgeDeploymentFlagAndCreate(cmd.Context(), client, cmd.Flags())
			if err != nil {
				cmd.Println(err)
				return err
			}
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
	deployHome, err := f.GetBool(deploy)
	if err != nil {
		return nil, err
	}
	deploymentResponse, err := CreateDeployment(ctx, client, *deployment, deployment.OID, deployment.TID, deployment.LID)
	if err != nil {
		return nil, err
	}
	log.Printf("deployment %s created successfully", deploymentResponse.UID)
	if deployHome {
		_, err := validateFlagsAndStartDeployment(ctx, client, deploymentResponse.UID)
		if err != nil {
			return nil, err
		}
		log.Println("deploying...")
	}
	return deploymentResponse, nil
}
