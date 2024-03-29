package cloudcmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
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
		WithBindFlagOnPreRun([]internal.BindFlag{{Name: org, Persistent: false, Type: internal.Cloud}}).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			return nil
		})
}

/*func updateTeam(ctx context.Context, client internal.CloudClient, orgID, teamID, teamName string) {
	teamPayload, response, err := client.UpdateTeam(ctx)
	if err != nil {
		return nil, err
	}
	if status {

	}
}*/
