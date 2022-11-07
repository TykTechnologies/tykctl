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
	"net/http"
	"time"
)

var (
	ErrorCreatingDeployment = errors.New("error creating deployment")
)

func NewCreateDeploymentCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(create).
		WithFlagAdder(true, addDeploymentFlag).
		WithBindFlagOnPreRun([]BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		WithCommands(NewCreateHomeDeployment(client), NewCreateEdgeDeployment(client))
}

func addDeploymentFlag(f *pflag.FlagSet) {
	f.StringP(name, "n", "", "name for the deployment you want to create.")
	f.String(zone, "", "the region you want to deploy into")
	f.String(domain, "", "custom domain for your deployment")
}

func CreateDeployment(ctx context.Context, client internal.CloudClient, deployment cloud.Deployment, orgId, teamId, envId string) (*cloud.Deployment, error) {
	deployResponse, resp, err := client.CreateDeployment(ctx, deployment, orgId, teamId, envId)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, ErrorCreatingDeployment
	}
	if deployResponse.Status != statusOK {
		return nil, errors.New(deployResponse.Error_)
	}
	return deployResponse.Payload, nil
}

func newDeployment() cloud.Deployment {
	d := cloud.Deployment{
		Created: time.Now().UTC(),
		Driver:  "K8s_sp",
		DriverMetaData: &cloud.Status{
			CurrentState: "starting",
			Timestamp:    time.Now().UTC(),
		},
		ExtraContext: &cloud.MetaDataStore{
			Data: map[string]map[string]interface{}{},
		},
		LinkedDeployments: map[string]string{},
		FriendlyNames:     make(map[string]string),
		Ingresses:         make(map[string]string),
		Kind:              gateway,
		LastUpdate:        time.Now().UTC(),
		Tags:              make([]string, 0),
		ZoneCode:          "",
	}
	return d
}

func extractCommonDeploymentFlags(f *pflag.FlagSet) (*cloud.Deployment, error) {
	orgId := viper.GetString(org)
	if util.StringIsEmpty(orgId) {
		return nil, ErrorOrgRequired
	}
	tid := viper.GetString(team)
	if util.StringIsEmpty(tid) {
		return nil, ErrorTeamRequired
	}
	envId := viper.GetString(env)
	if util.StringIsEmpty(envId) {
		return nil, ErrorEnvRequired
	}
	deploymentName, err := f.GetString(name)
	if err != nil {
		return nil, err
	}
	if util.StringIsEmpty(deploymentName) {
		return nil, ErrorNameRequired
	}
	zone, err := f.GetString(zone)
	if err != nil {
		return nil, err
	}
	d := newDeployment()
	d.ZoneCode = zone
	d.Name = deploymentName
	d.OID = orgId
	d.LID = envId
	d.TID = tid
	return &d, nil
}
