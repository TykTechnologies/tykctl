package cloudcmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

func NewCreateEnvironmentCmd(client internal.CloudClient) *cobra.Command {
	return internal.NewCmd(create).
		WithLongDescription(createEnvDesc).
		WithDescription("creates an environment in a given team.").
		WithFlagAdder(false, createEnvironment).
		WithBindFlagOnPreRun([]internal.BindFlag{{Name: "org", Persistent: false}, {Name: "team", Persistent: false}}).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			org := viper.GetString(org)
			team := viper.GetString(team)
			envName, err := cmd.Flags().GetString(name)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			env, err := validateFlagsAndCreateEnv(ctx, client, envName, team, org)
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
func validateFlagsAndCreateEnv(ctx context.Context, client internal.CloudClient, envName, teamId, orgId string) (*cloud.Loadout, error) {
	if util.StringIsEmpty(orgId) {
		return nil, ErrorOrgRequired
	}
	if util.StringIsEmpty(teamId) {
		return nil, ErrorTeamRequired
	}
	env := cloud.Loadout{Name: envName}
	if util.StringIsEmpty(env.Name) {
		return nil, ErrorNameRequired
	}
	environment, err := CreateEnvironment(ctx, client, env, orgId, teamId)
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
