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

const createHomeDeploymentDesc = ` 
NOTE: Your home deployment zone has to be the same as your organization zone.

This command creates an Home Deployment.

NOTE: This does not deploy the deployment it just create it.You can use the deploy flag to deploy after create.You can also use the deploy command to deploy the created deployment.

You must pass the organization,team,zone and environment you want deploy this deployment.

If you do not pass the org,zone or environment we will use the ones on your config file.You can set the default org,team and environment by running:

tykctl cloud init

Sample usage for this command

tykctl cloud deployments create edge --name="test deployment"
`

func NewCreateHomeDeployment(client internal.CloudClient) *cobra.Command {
	return internal.NewCmd(home).
		WithLongDescription(createHomeDeploymentDesc).
		WithFlagAdder(false, addHomeDeploymentFlag).
		WithDescription("create a control plane in your home region.").
		WithExample("tykctl cloud deployments create home --name='home-deployment'").
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			_, err := validateHomeDeploymentFlagAndCreate(cmd.Context(), client, cmd.Flags())
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			return nil
		})
}

func addHomeDeploymentFlag(f *pflag.FlagSet) {
	f.String(awsKeyID, "", "API Key ID to use with AWS")
	f.String(awsRegion, "", "AWS region to deploy into (e.g. us-west-2)")
	f.String(awsSecret, "", "API Secret to use with AWS")
	f.Bool(enablePlugins, false, "enable plugins for the control plane")
}

func validateHomeDeploymentFlagAndCreate(ctx context.Context, client internal.CloudClient, f *pflag.FlagSet) (*cloud.Deployment, error) {
	deployment, err := extractCommonDeploymentFlags(f)
	if err != nil {
		return nil, err
	}
	plugins, err := f.GetBool(enablePlugins)
	if err != nil {
		return nil, err
	}
	deployment.HasAWSSecrets = plugins
	if plugins {
		aws, err := getAwsKeys(f)
		if err != nil {
			return nil, err
		}
		deployment.ExtraContext.Data["CFData"] = map[string]interface{}{"AWSSecret": aws.AwsSecret, "AWSKeyID": aws.AwsKeyID, "AWSRegion": aws.AwsRegion}

	}
	deployment.Kind = controlPlane
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

func getAwsKeys(f *pflag.FlagSet) (*Aws, error) {
	awsKey, err := f.GetString(awsKeyID)
	if err != nil {
		return nil, err
	}
	awsReg, err := f.GetString(awsRegion)
	if err != nil {
		return nil, err
	}
	awsSec, err := f.GetString(awsSecret)
	if err != nil {
		return nil, err
	}
	if util.StringIsEmpty(awsKey) || util.StringIsEmpty(awsReg) || util.StringIsEmpty(awsSec) {
		return nil, errors.New("--aws-key-id, --aws-secret, and --aws-region, must all be set when plugins are enabled")
	}
	awsValues := Aws{
		AwsRegion: awsReg,
		AwsSecret: awsSec,
		AwsKeyID:  awsKey,
	}

	return &awsValues, nil
}

type Aws struct {
	AwsRegion string
	AwsSecret string
	AwsKeyID  string
}
