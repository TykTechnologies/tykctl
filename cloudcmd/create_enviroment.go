package cloudcmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"net/http"
)

const createEnvDesc = `
This command create an environment in a team.
You must pass the name of the environment.
You must also set the org and team you want to create this environment in.
If you don't pass the org and team we will use the one set in the config file.

Sample usage:

tyk cloud environments create --name="staging"
`

var (
	ErrorCreatingEnv = errors.New("error creating environment")
)

func NewCreateEnvironmentCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(create).
		AddPreRunFuncs(NewCloudRbac(TeamAdmin, factory.Config).CloudRbac).
		WithLongDescription(createEnvDesc).
		WithDescription("creates an environment in a given team.").
		WithFlagAdder(false, createEnvironment).
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: org, Persistent: false}, {Name: team, Persistent: false}}).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			env, err := validateFlagsAndCreateEnv(ctx, factory.Client, factory.Config, cmd)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			cmd.Println(fmt.Sprintf("Environment %s created successfully", env.UID))
			return nil
		})
}

// createEnvironment declares the local flags to be used to create environment.
func createEnvironment(f *pflag.FlagSet) {
	f.StringP(name, "n", "", "name for the environment you want to create.")
}

// validateFlagsAndCreateEnv validate that the cloudcmd flags are not empty and create environment in a team.
func validateFlagsAndCreateEnv(ctx context.Context, client internal.CloudClient, config internal.UserConfig, cmd cobra.Command) (*cloud.Loadout, error) {
	commonEnvFlags, err := validateCommonEnvFlags(config, cmd)
	env := cloud.Loadout{Name: commonEnvFlags.envName}
	environment, err := CreateEnvironment(ctx, client, env, commonEnvFlags.orgId, commonEnvFlags.teamId)
	if err != nil {
		return nil, err
	}
	return environment, nil
}

// CreateEnvironment creates an environment is a given team.
func CreateEnvironment(ctx context.Context, client internal.CloudClient, env cloud.Loadout, orgId, teamId string) (*cloud.Loadout, error) {
	envResponse, resp, err := client.CreateEnv(ctx, env, orgId, teamId)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, ErrorCreatingEnv
	}
	if envResponse.Status != statusOK {
		return nil, errors.New(envResponse.Error_)
	}
	return envResponse.Payload, err
}
