package cloudcmd

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

const deploymentStatusDesc = `
This command will output the status of a deployment given its uuid.
You will receive an error if the deployment does not exist.

Sample usage for this command:
tykctl cloud deployments status b5c503e8-c632-4ce0-9629-b0ee3e3c2c62 
`

func NewDeploymentStatusCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(status).
		AddPreRunFuncs(NewCloudRbac(TeamMember).CloudRbac).
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		WithExample("tykctl cloud deployments status <deployment uuid>").
		WithLongDescription(deploymentStatusDesc).
		WithDescription("output the status of a deployment given its uuid.").
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			checkStatus, err := validateFlagsAndCheckStatus(cmd.Context(), factory.Client, args[0])
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			cmd.Println(checkStatus.CurrentState)
			return nil
		})

}
func validateFlagsAndCheckStatus(ctx context.Context, client internal.CloudClient, deploymentID string) (*Status, error) {
	deploymentFlags, err := validateCommonDeploymentFlags()
	if err != nil {
		return nil, err
	}
	status, err := deploymentStatus(ctx, client, deploymentFlags.OrgId, deploymentFlags.TeamId, deploymentFlags.EnvId, deploymentID)
	if err != nil {
		return nil, err
	}
	return status, err
}

func deploymentStatus(ctx context.Context, client internal.CloudClient, orgId, teamId, envId, deploymentId string) (*Status, error) {
	deploymentStatus, resp, err := client.GetDeploymentStatus(ctx, orgId, teamId, envId, deploymentId)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrorFetchingDeploymentStatus
	}
	if deploymentStatus.Status != statusOK {
		return nil, errors.New(deploymentStatus.Error_)
	}
	if deploymentStatus.Payload == nil {
		return nil, nil
	}
	b, err := json.Marshal(deploymentStatus.Payload)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(b, &status)
	return &status, err
}

type Status struct {
	// The state of the current deployment
	CurrentState State `json:"CurrentState"`
	// The timestamp of the last state
	Timestamp time.Time `json:"Timestamp"`
	// Metadata specific to the deployment, set by the control plane
	// swagger:ignore
	DriverData interface{} `secure:"default" bson:"-"`
}

type State string

const (
	Starting   State = "starting"
	Started    State = "started"
	Restarting State = "restarting"
	Stopping   State = "stopping"
	Stopped    State = "stopped"
	NoInfo     State = "no_info"
	Warning    State = "warning"
)
