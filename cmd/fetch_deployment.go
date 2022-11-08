package cmd

import (
	"context"
	"errors"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	ErrorFetchingDeployments = errors.New("error fetching deployments")
)

func NewFetchDeploymentCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(fetch).
		WithFlagAdder(false, addOutPutFlags).
		WithBindFlagOnPreRun([]BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		MaximumArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			if len(args) == 0 {

				err := validateAndFetchEnvDeployments(cmd.Context(), client, cmd.Flags())
				if err != nil {
					cmd.Println(err)
					return err
				}
				return nil
			}
			err := validateAndFetchDeploymentById(ctx, client, cmd.Flags(), args[0])
			if err != nil {
				cmd.Println(err)
				return err
			}
			return nil
		})
}
func validateAndFetchDeploymentById(ctx context.Context, client internal.CloudClient, f *pflag.FlagSet, id string) error {
	log.Println(id)
	deploymentFlags, err := validateCommonDeploymentFetchFlags(f)
	if err != nil {
		return err
	}
	deployment, err := GetDeploymentById(ctx, client, deploymentFlags.OrgId, deploymentFlags.TeamId, deploymentFlags.EnvId, id)
	if err != nil {
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
	return internal.PrintJson(deployment)
}
func validateAndFetchEnvDeployments(ctx context.Context, client internal.CloudClient, f *pflag.FlagSet) error {
	deploymentFlags, err := validateCommonDeploymentFetchFlags(f)
	if err != nil {
		return err
	}
	deployments, err := GetEnvDeployments(ctx, client, deploymentFlags.OrgId, deploymentFlags.TeamId, deploymentFlags.EnvId)
	if err != nil {
		return err
	}
	if deploymentFlags.OutPut == table {
		internal.Printable(CreateDeploymentHeadersAndRows(deployments))
		return nil
	}
	return internal.PrintJson(deployments)
}

func GetEnvDeployments(ctx context.Context, client internal.CloudClient, orgId, teamId, envId string) ([]cloud.Deployment, error) {
	depResponse, resp, err := client.GetEnvDeployments(ctx, orgId, teamId, envId)
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
func GetDeploymentById(ctx context.Context, client internal.CloudClient, orgId, teamId, envId, id string) (*cloud.Deployment, error) {
	depResponse, resp, err := client.GetDeploymentById(ctx, orgId, teamId, envId, id, nil)
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

func validateCommonDeploymentFetchFlags(f *pflag.FlagSet) (*DeploymentFetchFlags, error) {
	output, err := f.GetString(outPut)
	if err != nil {
		return nil, err
	}
	if output != table && output != jsonFormat {
		return nil, ErrorOutPutFormat
	}
	var deploymentFlag DeploymentFetchFlags
	deploymentFlag.OutPut = output
	deploymentFlag.OrgId = viper.GetString(org)
	if util.StringIsEmpty(deploymentFlag.OrgId) {
		return nil, ErrorOrgRequired
	}
	deploymentFlag.TeamId = viper.GetString(team)
	if util.StringIsEmpty(deploymentFlag.TeamId) {
		return nil, ErrorTeamRequired
	}
	deploymentFlag.EnvId = viper.GetString(env)
	if util.StringIsEmpty(deploymentFlag.EnvId) {
		return nil, ErrorEnvRequired
	}
	return &deploymentFlag, nil
}

type DeploymentFetchFlags struct {
	OrgId  string
	TeamId string
	EnvId  string
	OutPut string
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
