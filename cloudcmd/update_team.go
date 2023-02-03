package cloudcmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"net/http"
)

const updateTeamDesc = `
This command will update a team given it's uuid.
You have to pass the name you want to give the team and org the team belongs to.
If the org is not provided we will use the one you set in the config file.
To set a default org in the config file run:

tykctl cloud init

Sample usage for this command:

tyckctl cloud teams update <uuid> --name=<new name> --org=<org uuid>
`

func NewUpdateTeamCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(update).
		AddPreRunFuncs(NewCloudRbac(TeamAdmin, factory.Config).CloudRbac).
		WithFlagAdder(false, createTeamFlags).
		WithLongDescription(updateTeamDesc).
		WithDescription("update a team given it's uuid").
		WithExample("tyckctl cloud teams update <uuid> --name=<new name> --org=<org uuid>").
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: org, Persistent: false}}).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			team, err := validateFlagsAndUpdateTeam(cmd.Context(), factory.Client, factory.Config, cmd, args)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			cmd.Println(fmt.Sprintf("team %s updated successfully", team.UID))
			return nil
		})
}

// validateFlagsAndUpdateTeam extracts flags and args from cmd validate them and send the update request .
func validateFlagsAndUpdateTeam(ctx context.Context, client internal.CloudClient, config internal.UserConfig, cmd cobra.Command, args []string) (*cloud.Team, error) {
	orgId := config.GetCurrentUserOrg()
	if util.StringIsEmpty(orgId) {
		return nil, ErrorOrgRequired
	}
	teamName, err := cmd.Flags().GetString(name)
	if err != nil {
		return nil, nil
	}
	if util.StringIsEmpty(teamName) {
		return nil, ErrorNameRequired
	}
	if len(args) == 0 {
		return nil, ErrorTeamIdRequired
	}
	return updateTeam(ctx, client, teamName, args[0], orgId)
}

// updateTeam sends the update request to tyk cloud to update teams name.
func updateTeam(ctx context.Context, client internal.CloudClient, teamName, teamId, orgId string) (*cloud.Team, error) {
	team := cloud.Team{
		Name: teamName,
		UID:  teamId,
	}
	teamPayload, resp, err := client.UpdateTeam(ctx, team, orgId, team.UID)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrorUpdatingTeam
	}
	if teamPayload.Status != statusOK {
		return nil, errors.New(teamPayload.Error_)
	}
	return teamPayload.Payload, nil
}
