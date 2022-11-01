package cmd

import (
	"context"
	"errors"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

var (
	ErrorFetchingTeam = errors.New("error fetching team")
	ErrorTeamRequired = errors.New("team flag is required")
)

func NewFetchTeamCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(fetch).
		WithFlagAdder(false, addOutPutFlags).
		WithBindFlagOnPreRun([]BindFlag{{Name: "org", Persistent: false}}).
		WithExample("tykctl cloud teams fetch --output<json/table>").
		MaximumArgs(1, func(ctx context.Context, command cobra.Command, args []string) error {
			outPut, err := command.Flags().GetString(outPut)
			if err != nil {
				command.Println(err)
				return err
			}
			org := viper.GetString(org)
			if len(args) == 0 {
				err := FetchAndPrintTeams(ctx, client, outPut, org)
				if err != nil {
					command.Println(err)
					return err
				}
				return nil
			}
			err = FetchAndPrintTeamById(ctx, client, outPut, org, args[0])
			if err != nil {
				command.Println(err)
				return err
			}
			return nil
		})
}

// FetchAndPrintTeams all the teams in an organization and print them as json or table.
func FetchAndPrintTeams(ctx context.Context, client internal.CloudClient, output, orgId string) error {
	if output != table && output != jsonFormat {
		return ErrorOutPutFormat
	}
	if util.StringIsEmpty(orgId) {
		return ErrorOrgRequired
	}
	teams, err := GetTeams(ctx, client, orgId)
	if err != nil {
		return err
	}
	if output == table {
		internal.Printable(CreateTeamHeadersAndRows(teams))
		return nil
	}
	return internal.PrintJson(teams)
}

// FetchAndPrintTeamById print a single team by uuid and print it as a table or json.
func FetchAndPrintTeamById(ctx context.Context, client internal.CloudClient, output, orgId, teamId string) error {
	if output != table && output != jsonFormat {
		return ErrorOutPutFormat
	}
	if util.StringIsEmpty(orgId) {
		return ErrorTeamRequired
	}
	team, err := GetTeamById(ctx, client, orgId, teamId)
	if err != nil {
		return err
	}
	if output == table {
		var teams []cloud.Team
		if team != nil {
			teams = append(teams, *team)
		}
		internal.Printable(CreateTeamHeadersAndRows(teams))
		return nil
	}
	return internal.PrintJson(team)
}

// GetTeamById fetch a single team using its uuid.
func GetTeamById(ctx context.Context, client internal.CloudClient, orgId, teamId string) (*cloud.Team, error) {
	teamResponse, resp, err := client.GetTeamById(ctx, orgId, teamId)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrorFetchingTeam
	}
	if teamResponse.Status != statusOK {
		return nil, errors.New(teamResponse.Error_)
	}
	return teamResponse.Payload, nil
}

// GetTeams will return all the teams in an organization.
func GetTeams(ctx context.Context, client internal.CloudClient, orgId string) ([]cloud.Team, error) {
	teamResponse, resp, err := client.GetTeams(ctx, orgId)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrorFetchingTeam
	}
	if teamResponse.Status != statusOK {
		return nil, errors.New(teamResponse.Error_)
	}
	if teamResponse.Payload == nil {
		return nil, nil
	}
	return teamResponse.Payload.Teams, nil
}

// CreateTeamHeadersAndRows create headers and rows to be used in creating teams table.
func CreateTeamHeadersAndRows(teams []cloud.Team) ([]string, [][]string) {
	header := []string{"Name", "UID", "Environments", "Deployments"}
	rows := make([][]string, 0)
	for _, team := range teams {
		row := []string{
			team.Name, team.UID,
		}
		rows = append(rows, row)
	}
	return header, rows
}
