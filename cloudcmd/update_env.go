package cloudcmd

import (
	"context"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
	"net/http"
)

const updateEnvDesc = `
This command will update a environment given it's uuid.
You have to pass the new name you want to give the environment, the org and the team this environment belongs to.

If the org or team is not provided we will use the one you set in the config file.

To set a default org and team in the config file run:

tykctl cloud init

Sample usage for this command:

tyckctl cloud env update <uuid> --name=<new name> --org=<org uuid> --team=<team uuid>
`

func NewUpdateEnvCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(update).
		AddPreRunFuncs(NewCloudRbac(TeamAdmin, factory.Config).CloudRbac).
		WithFlagAdder(false, createEnvironment).
		WithLongDescription(updateEnvDesc).
		WithExample("tyckctl cloud env update <uuid> --name=<new name> --org=<org uuid> --team=<team uuid>").
		WithDescription("update a environment given its uuid").
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: org, Persistent: false}, {Name: team, Persistent: false}}).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			env, err := validateFlagsAndUpdateEnv(ctx, factory.Client, factory.Config, cmd, args)
			if err != nil {
				return err
			}
			cmd.Println(fmt.Sprintf("enviroment %s updated successfully", env.UID))
			return nil
		})
}

func validateFlagsAndUpdateEnv(ctx context.Context, client internal.CloudClient, config internal.UserConfig, cmd cobra.Command, args []string) (*cloud.Loadout, error) {
	commonEnvFlags, err := validateCommonEnvFlags(config, cmd)
	if err != nil {
		return nil, nil
	}
	if len(args) == 0 {
		return nil, ErrorEnvRequired
	}
	return updateEnv(ctx, client, commonEnvFlags.orgId, commonEnvFlags.teamId, commonEnvFlags.envName, args[0])
}

func updateEnv(ctx context.Context, client internal.CloudClient, orgId, teamId, envName, envId string) (*cloud.Loadout, error) {
	env := cloud.Loadout{
		Name: envName,
		UID:  envId,
	}
	envPayload, resp, err := client.UpdateEnv(ctx, env, orgId, teamId, envId)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrorUpdatingTeam
	}
	return envPayload.Payload, nil
}
