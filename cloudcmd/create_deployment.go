package cloudcmd

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
)

const createDeploymentDesc = `This is the parent command for 
            creating either an edge or a home gateway.You must provide exactly one arg which can be either home or edge.

Note: You will need to use the deploy flag if you want the deployment to be deployed after create.
     
     Example: tykctl deployments create home
`

var (
	ErrorCreatingDeployment   = errors.New("error creating deployment")
	ErrorStartingDeployment   = errors.New("error starting deployment")
	ErrorRestartingDeployment = errors.New("error restarting deployment")
	ErrorZoneCodeIsRequired   = errors.New("error zone is required")
)

func NewCreateDeploymentCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(create).
		WithLongDescription(createDeploymentDesc).
		WithDescription("This is the parent command for creating the edge or home deployment.").
		WithFlagAdder(true, addDeploymentFlag).
		WithExample("tykctl deployments create home").
		AddPreRunFuncs(NewCloudRbac(TeamAdmin, factory.Config).CloudRbac).
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		WithCommands(
			NewCreateHomeDeployment(factory),
			NewCreateEdgeDeployment(factory),
		)
}

func addDeploymentFlag(f *pflag.FlagSet) {
	f.StringP(name, "n", "", "name for the deployment you want to create.")
	f.String(zone, "", "the region you want to deploy into e.g aws-eu-west-2")
	f.String(domain, "", "custom domain for your deployment")
	f.Bool(deploy, false, "deploy the deployment after create")
}

func CreateDeployment(ctx context.Context, client internal.CloudClient, deployment cloud.Deployment, orgID, teamID, envID string) (*cloud.Deployment, error) {
	deployResponse, resp, err := client.CreateDeployment(ctx, deployment, orgID, teamID, envID)
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

func extractCommonDeploymentFlags(f *pflag.FlagSet, config internal.UserConfig) (*cloud.Deployment, error) {
	deploymentFlags, err := validateCommonDeploymentFlags(config)
	if err != nil {
		return nil, err
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

	if util.StringIsEmpty(zone) {
		return nil, ErrorZoneCodeIsRequired
	}

	d := newDeployment()
	d.ZoneCode = zone
	d.Name = deploymentName
	d.OID = deploymentFlags.OrgID
	d.LID = deploymentFlags.EnvID
	d.TID = deploymentFlags.TeamID

	return &d, nil
}
