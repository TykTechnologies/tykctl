package cloudcmd

import (
	"context"
	"errors"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
)

const fetchEnvDesc = `
This command will fetch all the environment in a team.

You must pass the --org and --team.If they are not passed we will use the default org and team set in your config file.

We support json and table as the output format.To set the output format use the --output<json/table> flag.

Sample usage of this command:

tykctl cloud environments fetch --team=<teamID> --org=<orgID> --output=<json/table>
`

var (
	ErrorFetchingEnvironment = errors.New("error fetching environments")
	ErrorEnvRequired         = errors.New("error env is required")
)

func NewFetchEnvironmentCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(fetch).
		AddPreRunFuncs(NewCloudRbac(TeamMember, factory.Config).CloudRbac).
		WithFlagAdder(false, addOutPutFlags).
		WithFlagAdder(false, getValues).
		WithLongDescription(fetchEnvDesc).
		WithDescription("fetch environments from a given team.").
		WithExample("tykctl cloud environments fetch --team=<teamID> --org=<orgID>").
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: org, Persistent: false}, {Name: team, Persistent: false}}).
		MaximumArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			outPut, err := cmd.Flags().GetString(outPut)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}

			org := factory.Config.GetCurrentUserOrg()
			team := factory.Config.GetCurrentUserTeam()
			if len(args) == 0 {
				err = validateFlagsAndPrintEnvs(cmd.Context(), factory.Client, outPut, org, team)
				if err != nil {
					cmd.PrintErrln(err)
					return err
				}

				return nil
			}

			getVal, err := cmd.Flags().GetStringSlice(get)
			if err != nil {
				return err
			}

			err = validateFlagsAndPrintEnvByID(cmd.Context(), factory.Client, outPut, org, team, args[0], getVal)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			return nil
		})
}

// validateFlagsAndPrintEnvs validate flags passed to cloudcmd are not empty and prints env in a team .
func validateFlagsAndPrintEnvs(ctx context.Context, client internal.CloudClient, output, orgID, teamID string) error {
	if output != table && output != jsonFormat {
		return ErrorOutPutFormat
	}

	if util.StringIsEmpty(orgID) {
		return ErrorOrgRequired
	}

	if util.StringIsEmpty(teamID) {
		return ErrorTeamRequired
	}

	envs, err := GetEnvs(ctx, client, orgID, teamID)
	if err != nil {
		return err
	}

	if output == table {
		internal.Printable(CreateEnvHeadersAndRows(envs))
		return nil
	}

	return internal.PrintJSON(envs)
}

func validateFlagsAndPrintEnvByID(ctx context.Context, client internal.CloudClient, output, orgID, teamID, envID string, getValues []string) error {
	if output != table && output != jsonFormat {
		return ErrorOutPutFormat
	}

	if util.StringIsEmpty(orgID) {
		return ErrorOrgRequired
	}

	if util.StringIsEmpty(teamID) {
		return ErrorTeamRequired
	}

	if util.StringIsEmpty(envID) {
		return ErrorEnvRequired
	}

	env, err := GetEnvByID(ctx, client, orgID, teamID, envID)
	if err != nil {
		return err
	}

	if len(getValues) > 0 {
		err = internal.HandleGets(env, getValues)
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

	return internal.PrintJSON(env)
}

// GetEnvByID gets an environment by its id.
func GetEnvByID(ctx context.Context, client internal.CloudClient, orgID, teamID, envID string) (*cloud.Loadout, error) {
	envResponse, resp, err := client.GetEnvByID(ctx, orgID, teamID, envID)
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
func GetEnvs(ctx context.Context, client internal.CloudClient, orgID, teamID string) ([]cloud.Loadout, error) {
	envResponse, resp, err := client.GetEnvs(ctx, orgID, teamID)
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
