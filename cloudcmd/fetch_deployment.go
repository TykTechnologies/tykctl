package cloudcmd

import (
	"context"
	"errors"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
)

const fetchDeploymentDesc = `
This command will fetch all the deployment belonging to a environment.

Note: You must pass an --env, --team  and --org flag command.

If you don't pass the org,team or env we will use the one set in your config file.

You can either get the data as json or in a table.
Use the --output<table,json> flag to change the format default is table.

Sample usage: 

tykctl cloud deployments fetch --org=<organization id> --output=<json/table>
`

var (
	ErrorFetchingDeployments      = errors.New("error fetching deployments")
	ErrorFetchingDeploymentStatus = errors.New("error fetching deployment status")
)

func NewFetchDeploymentCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(fetch).
		AddPreRunFuncs(NewCloudRbac(TeamMember, factory.Config).CloudRbac).
		WithFlagAdder(false, addOutPutFlags).
		WithFlagAdder(false, getValues).
		WithLongDescription(fetchDeploymentDesc).
		WithDescription("fetch deployment from an environment.").
		WithExample("tykctl cloud deployments fetch").
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		MaximumArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			if len(args) == 0 {
				err := validateAndFetchEnvDeployments(cmd.Context(), factory.Client, factory.Config, cmd.Flags())
				if err != nil {
					cmd.PrintErrln(err)
					return err
				}
				return nil
			}
			err := validateAndFetchDeploymentByID(ctx, factory.Client, factory.Config, cmd.Flags(), args[0])
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			return nil
		})
}

func validateAndFetchDeploymentByID(ctx context.Context, client internal.CloudClient, config internal.UserConfig, f *pflag.FlagSet, id string) error {
	deploymentFlags, err := validateCommonDeploymentFetchFlags(f, config)
	if err != nil {
		return err
	}

	deployment, err := GetDeploymentByID(ctx, client, deploymentFlags.OrgID, deploymentFlags.TeamID, deploymentFlags.EnvID, id)
	if err != nil {
		return err
	}

	getVals, err := f.GetStringSlice(get)
	if err != nil {
		return err
	}

	if len(getVals) > 0 {
		err := internal.HandleGets(deployment, getVals)
		return err
	}

	if deploymentFlags.OutPut == table {
		var deployments []cloud.Deployment
		if deployment != nil {
			deployments = append(deployments, *deployment)
		}

		internal.Printable(CreateDeploymentHeadersAndRows(deployments))

		return nil
	}

	return internal.PrintJSON(deployment)
}

func validateAndFetchEnvDeployments(ctx context.Context, client internal.CloudClient, config internal.UserConfig, f *pflag.FlagSet) error {
	deploymentFlags, err := validateCommonDeploymentFetchFlags(f, config)
	if err != nil {
		return err
	}

	deployments, err := GetEnvDeployments(ctx, client, deploymentFlags.OrgID, deploymentFlags.TeamID, deploymentFlags.EnvID)
	if err != nil {
		return err
	}

	if deploymentFlags.OutPut == table {
		internal.Printable(CreateDeploymentHeadersAndRows(deployments))
		return nil
	}

	return internal.PrintJSON(deployments)
}

func GetEnvDeployments(ctx context.Context, client internal.CloudClient, orgID, teamID, envID string) ([]cloud.Deployment, error) {
	depResponse, resp, err := client.GetEnvDeployments(ctx, orgID, teamID, envID)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrorFetchingDeployments
	}

	if depResponse.Status != statusOK {
		return nil, errors.New(depResponse.Error_)
	}

	if depResponse.Payload == nil {
		return nil, nil
	}

	return depResponse.Payload.Deployments, nil
}

func GetDeploymentByID(ctx context.Context, client internal.CloudClient, orgID, teamID, envID, id string) (*cloud.Deployment, error) {
	depResponse, resp, err := client.GetDeploymentByID(ctx, orgID, teamID, envID, id, nil)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrorFetchingDeployments
	}

	if depResponse.Status != statusOK {
		return nil, errors.New(depResponse.Error_)
	}

	return depResponse.Payload, nil
}

func validateCommonDeploymentFetchFlags(f *pflag.FlagSet, config internal.UserConfig) (*DeploymentFlags, error) {
	deploymentFlag, err := validateCommonDeploymentFlags(config)
	if err != nil {
		return nil, err
	}

	output, err := f.GetString(outPut)
	if err != nil {
		return nil, err
	}

	deploymentFlag.OutPut = output

	if output != table && output != jsonFormat {
		return nil, ErrorOutPutFormat
	}

	return deploymentFlag, nil
}

func CreateDeploymentHeadersAndRows(deployments []cloud.Deployment) ([]string, [][]string) {
	header := []string{"Name", "UID", "Kind", "Region", "State"}
	rows := make([][]string, 0)

	for _, deployment := range deployments {
		row := []string{
			deployment.Name, deployment.UID, deployment.Kind, deployment.ZoneCode, deployment.State,
		}
		rows = append(rows, row)
	}

	return header, rows
}
