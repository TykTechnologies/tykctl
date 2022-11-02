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
	ErrorFetchingEnvironment = errors.New("error fetching environments")
	ErrorEnvRequired         = errors.New("error env is required")
)

func NewFetchEnvironmentCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(fetch).
		WithFlagAdder(false, addOutPutFlags).
		WithBindFlagOnPreRun([]BindFlag{{Name: "org", Persistent: false}, {Name: "team", Persistent: false}}).
		MaximumArgs(1, func(ctx context.Context, command cobra.Command, args []string) error {
			outPut, err := command.Flags().GetString(outPut)
			if err != nil {
				command.Println(err)
				return err
			}
			org := viper.GetString(org)
			team := viper.GetString(team)
			if len(args) == 0 {
				err := validateFlagsAndPrintEnvs(command.Context(), client, outPut, org, team)
				if err != nil {
					command.Println(err)
					return err
				}
				return nil
			}
			err = validateFlagsAndPrintEnvById(command.Context(), client, outPut, org, team, args[0])
			if err != nil {
				command.Println(err)
				return err
			}
			return nil
		})
}

// validateFlagsAndPrintEnvs validate flags passed to cmd are not empty and prints env in a team .
func validateFlagsAndPrintEnvs(ctx context.Context, client internal.CloudClient, output, orgId, teamId string) error {
	if output != table && output != jsonFormat {
		return ErrorOutPutFormat
	}
	if util.StringIsEmpty(orgId) {
		return ErrorOrgRequired
	}
	if util.StringIsEmpty(teamId) {
		return ErrorTeamRequired
	}
	envs, err := GetEnvs(ctx, client, orgId, teamId)
	if err != nil {
		return err
	}
	if output == table {
		internal.Printable(CreateEnvHeadersAndRows(envs))
		return nil
	}
	return internal.PrintJson(envs)
}
func validateFlagsAndPrintEnvById(ctx context.Context, client internal.CloudClient, output, orgId, teamId, envId string) error {
	if output != table && output != jsonFormat {
		return ErrorOutPutFormat
	}
	if util.StringIsEmpty(orgId) {
		return ErrorOrgRequired
	}
	if util.StringIsEmpty(teamId) {
		return ErrorTeamRequired
	}
	if util.StringIsEmpty(envId) {
		return ErrorEnvRequired
	}
	env, err := GetEnvById(ctx, client, orgId, teamId, envId)
	if err != nil {
		return err
	}
	if output == table {
		var envs []cloud.Loadout
		if env != nil {
			envs = append(envs, *env)
		}
		internal.Printable(CreateEnvHeadersAndRows(envs))
		return nil
	}
	return internal.PrintJson(env)
}

// GetEnvById gets an environment by its id.
func GetEnvById(ctx context.Context, client internal.CloudClient, orgId, teamId, envId string) (*cloud.Loadout, error) {
	envResponse, resp, err := client.GetEnvById(ctx, orgId, teamId, envId)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrorFetchingEnvironment
	}
	if envResponse.Status != statusOK {
		return nil, errors.New(envResponse.Error_)
	}
	return envResponse.Payload, nil
}

// GetEnvs gets all the environments in a team.
func GetEnvs(ctx context.Context, client internal.CloudClient, orgId string, teamId string) ([]cloud.Loadout, error) {
	envResponse, resp, err := client.GetEnvs(ctx, orgId, teamId)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrorFetchingEnvironment
	}
	if envResponse.Status != statusOK {
		return nil, errors.New(envResponse.Error_)
	}
	if envResponse.Payload == nil {
		return nil, nil
	}
	return envResponse.Payload.Loadouts, nil
}

// CreateEnvHeadersAndRows creates headers and tables to be used in creating env tables.
func CreateEnvHeadersAndRows(envs []cloud.Loadout) ([]string, [][]string) {
	header := []string{"Name", "UID", "Team", "Active Deployments"}
	rows := make([][]string, 0)
	for _, env := range envs {
		row := []string{
			env.Name, env.UID, env.TeamName,
		}
		rows = append(rows, row)
	}
	return header, rows
}
