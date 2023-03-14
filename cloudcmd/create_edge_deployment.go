package cloudcmd

import (
	"context"
	"errors"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
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

var ErrorControlPlaneRequired = errors.New("error control plane to link the gateway to is required")

func NewCreateEdgeDeployment(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(edge).
		WithLongDescription(createEdgeDeploymentDesc).
		AddPreRunFuncs(NewCloudRbac(TeamAdmin, factory.Config).CloudRbac).
		WithDescription("will create the edge gateway in a given environment").
		WithExample("tykctl cloud deployments create edge --name='test deployment'").
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		WithFlagAdder(false, addEdgeDeploymentFlag).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			_, err := validateEdgeDeploymentFlagAndCreate(cmd.Context(), factory.Client, cmd.Flags(), factory.Config)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			return nil
		})
}

func addEdgeDeploymentFlag(f *pflag.FlagSet) {
	f.String(linkedControlPlane, "", "control plane to link the edge gateway to.")
}

func validateEdgeDeploymentFlagAndCreate(ctx context.Context, client internal.CloudClient, f *pflag.FlagSet, config internal.UserConfig) (*cloud.Deployment, error) {
	deployment, err := extractCommonDeploymentFlags(f, config)
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
		_, err = validateFlagsAndStartDeployment(ctx, client, config, deploymentResponse.UID)
		if err != nil {
			return nil, err
		}

		log.Println("deploying...")
	}

	return deploymentResponse, nil
}
