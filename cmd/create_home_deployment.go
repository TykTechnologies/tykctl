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

func NewCreateHomeDeployment(client internal.CloudClient) *cobra.Command {
	return NewCmd(home).
		WithFlagAdder(false, addHomeDeploymentFlag).
		WithBindFlagOnPreRun([]BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			deployment, err := validateHomeDeploymentFlagAndCreate(cmd.Context(), client, cmd.Flags())
			if err != nil {
				cmd.Println(err)
				return err
			}
			cmd.Println(fmt.Sprintf("deployment %s created successfully", deployment.UID))
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
	return CreateDeployment(ctx, client, *deployment, deployment.OID, deployment.TID, deployment.LID)
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
