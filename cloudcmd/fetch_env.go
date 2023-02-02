package cloudcmd

import (
	"context"
	"errors"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"net/http"
)

const fetchEnvDesc = `
This command will fetch all the environment in a team.

You must pass the --org and --team.If they are not passed we will use the default org and team set in your config file.

We support json and table as the output format.To set the output format use the --output<json/table> flag.

Sample usage of this command:

tykctl cloud environments fetch --team=<teamId> --org=<orgID> --output=<json/table>
`

var (
	ErrorFetchingEnvironment = errors.New("error fetching environments")
	ErrorEnvRequired         = errors.New("error env is required")
)

func NewFetchEnvironmentCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(fetch).
		AddPreRunFuncs(NewCloudRbac(TeamMember).CloudRbac).
		WithFlagAdder(false, addOutPutFlags).
		WithLongDescription(fetchEnvDesc).
		WithDescription("fetch environments from a given team.").
		WithExample("tykctl cloud environments fetch --team=<teamId> --org=<orgID>").
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: org, Persistent: false}, {Name: team, Persistent: false}}).
		MaximumArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			outPut, err := cmd.Flags().GetString(outPut)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			org := getCurrentUserOrg()
			team := getCurrentUserTeam()
			if len(args) == 0 {
				err := validateFlagsAndPrintEnvs(cmd.Context(), factory.Client, outPut, org, team)
				if err != nil {
					cmd.PrintErrln(err)
					return err
				}
				return nil
			}
			err = validateFlagsAndPrintEnvById(cmd.Context(), factory.Client, outPut, org, team, args[0])
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			return nil
		})
}

// validateFlagsAndPrintEnvs validate flags passed to cloudcmd are not empty and prints env in a team .
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
