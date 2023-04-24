package cloudcmd

import (
	"context"
	"errors"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
)

func NewUpdateDeployment(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(update).
		WithFlagAdder(false, setValues).
		WithFlagAdder(false, envValues).
		WithFlagAdder(false, updateDeploymentFlag).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			deployment, err := validateDeploymentFlagsAndUpdate(ctx, factory.Client, factory.Config, cmd.Flags(), args[0])
			if err != nil {
				return err
			}

			cmd.Printf("updated %s\n", deployment.UID)

			deployAfterUpdate, err := cmd.Flags().GetBool(deploy)
			if err != nil {
				return err
			}

			if deployAfterUpdate {
				_, err = validateFlagsAndStartDeployment(ctx, factory.Client, factory.Config, deployment.UID)

				if err != nil {
					return err
				}

				log.Println("deploying...")
				return nil
			}

			return nil
		})
}

func updateDeploymentFlag(f *pflag.FlagSet) {
	f.String(zone, "", "the region you want to deploy into e.g aws-eu-west-2")
	f.String(domain, "", "custom domain for your deployment")
	f.Bool(deploy, false, "deploy the deployment after create")
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

	zone, err := f.GetString(zone)
	if err != nil {
		return nil, err
	}

	if !util.StringIsEmpty(zone) {
		dep.ZoneCode = zone
	}

	err = handleDeploymentDynamicVars(dep, f)
	if err != nil {
		return nil, err
	}

	return UpdateDeployment(ctx, client, *dep, deploymentFlags.OrgID, deploymentFlags.TeamID, deploymentFlags.EnvID, id)
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

func handleEnvVariables(deployment *cloud.Deployment, sets []string) error {
	if reflect.ValueOf(deployment).Kind() != reflect.Ptr {
		return errors.New("out put must be a pointer")
	}

	for _, set := range sets {
		var value bool
		var err error

		keyValue := strings.Split(set, "=")
		if keyValue[1] == "true" || keyValue[1] == "false" {
			value, err = strconv.ParseBool(keyValue[1])
			if err != nil {
				return err
			}

			deployment.ExtraContext.Data["EnvData"][keyValue[0]] = value
		} else {
			deployment.ExtraContext.Data["EnvData"][keyValue[0]] = keyValue[1]
		}
	}

	return nil
}

func handleDeploymentDynamicVars(deployment *cloud.Deployment, f *pflag.FlagSet) error {
	setVals, err := f.GetStringSlice(set)
	if err != nil {
		return err
	}

	err = internal.HandleSets(deployment, setVals)
	if err != nil {
		return err
	}

	envVars, err := f.GetStringSlice(envValue)
	if err != nil {
		return err
	}

	err = handleEnvVariables(deployment, envVars)
	if err != nil {
		return err
	}

	return nil
}
