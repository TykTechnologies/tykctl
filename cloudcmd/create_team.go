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

const createTeamDesc = `
This command will create a team in a given organization.
You have to pass the name you want to give the team and org in which you want to create the team.
If the org is not provided we will use the one you set in the config file.
To set a default org in the config file run:
tykctl cloud init
Sample usage for this command:
tyckctl cloud teams create --name="first team" --org=<org uuid>
`

var (
	ErrorCreatingTeam = errors.New("error creating team")
	ErrorOrgRequired  = errors.New("org flag is required")
	ErrorNameRequired = errors.New("name flag is required")
)

func NewCreateTeamCmd(client internal.CloudClient) *cobra.Command {
	return internal.NewCmd(create).WithFlagAdder(false, createTeamFlags).
		WithLongDescription(createTeamDesc).
		WithDescription("creates a team in a given organization.").
		WithExample("tyckctl cloud teams create --name='first team' --org=<org uuid>").
		WithBindFlagOnPreRun([]internal.BindFlag{{Name: "org", Persistent: false}}).
		NoArgs(func(ctx context.Context, command cobra.Command) error {
			org := viper.GetString(org)
			teamName, err := command.Flags().GetString(name)
			if err != nil {
				command.Println(err)
				return err
			}
			team, err := validateFlagsAndCreateTeam(ctx, client, teamName, org)
			if err != nil {
				command.Println(err)
				return err
			}
			command.Println(fmt.Sprintf("team %s created successfully", team.UID))
			return nil
		})
}

// validateFlagsAndCreateTeam validate that org and name are not empty and send request to create a team.
func validateFlagsAndCreateTeam(ctx context.Context, client internal.CloudClient, teamName, orgId string) (*cloud.Team, error) {
	if util.StringIsEmpty(orgId) {
		return nil, ErrorOrgRequired
	}
	if util.StringIsEmpty(teamName) {
		return nil, ErrorNameRequired
	}
	team := cloud.Team{Name: teamName}
	createdTeam, err := CreateTeam(ctx, client, team, orgId)
	if err != nil {
		return nil, err
	}
	return createdTeam, nil
}

// createTeamFlags declares local flags to be added to the team command.
func createTeamFlags(f *pflag.FlagSet) {
	f.StringP(name, "n", "", "name for the team you want to create.")
}

// CreateTeam the team send a request to the cloud to create a team.
func CreateTeam(ctx context.Context, client internal.CloudClient, team cloud.Team, orgId string) (*cloud.Team, error) {
	teamResponse, resp, err := client.CreateTeam(ctx, team, orgId)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, ErrorCreatingTeam
	}
	if teamResponse.Status != statusOK {
		return nil, errors.New(teamResponse.Error_)
	}
	return teamResponse.Payload, nil
}
