package cloudcmd

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
)

func NewUpdateDeployment(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(update).
		WithFlagAdder(false, setValues).
		WithFlagAdder(false, envValues).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			deployment, err := validateDeploymentFlagsAndUpdate(ctx, factory.Client, factory.Config, cmd.Flags(), args[0])
			if err != nil {
				return err
			}

			cmd.Printf("updated %s\n", deployment.UID)

			return nil
		})
}

func validateDeploymentFlagsAndUpdate(ctx context.Context, client internal.CloudClient, config internal.UserConfig, f *pflag.FlagSet, id string) (*cloud.Deployment, error) {
	deploymentFlags, err := validateCommonDeploymentFlags(config)
	if err != nil {
		return nil, err
	}

	dep, err := GetDeploymentByID(ctx, client, deploymentFlags.OrgID, deploymentFlags.TeamID, deploymentFlags.EnvID, id)
	if err != nil {
		return nil, err
	}

	setVals, err := f.GetStringSlice(set)
	if err != nil {
		return nil, err
	}

	err = internal.HandleSets(dep, setVals)
	if err != nil {
		return nil, err
	}

	envVars, err := f.GetStringSlice(envValue)
	if err != nil {
		return nil, err
	}

	deployment, err := handleEnvVariables(*dep, envVars)
	if err != nil {
		return nil, err
	}

	return UpdateDeployment(ctx, client, *deployment, deploymentFlags.OrgID, deploymentFlags.TeamID, deploymentFlags.EnvID, id)
}

func UpdateDeployment(ctx context.Context, client internal.CloudClient, deployment cloud.Deployment, orgID, teamID, envID, id string) (*cloud.Deployment, error) {
	deployResponse, resp, err := client.UpdateDeployment(ctx, deployment, orgID, teamID, envID, id, &cloud.DeploymentsApiUpdateDeploymentOpts{RetainSecrets: optional.NewBool(false)})
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(deployResponse.Error_)
	}

	return deployResponse.Payload, nil
}

func handleEnvVariables(deployment cloud.Deployment, sets []string) (*cloud.Deployment, error) {
	values := map[string]interface{}{}
	///deployment.ExtraContext.Data["EnvData"]
	for _, set := range sets {
		var value bool
		var err error

		keyValue := strings.Split(set, "=")
		if keyValue[1] == "true" || keyValue[1] == "false" {
			value, err = strconv.ParseBool(keyValue[1])
			if err != nil {
				return nil, err
			}

			values[keyValue[0]] = value
		} else {
			values[keyValue[0]] = keyValue[1]
		}
	}

	deployment.ExtraContext.Data["EnvData"] = values

	return &deployment, nil
}
