package cloudcmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
)

const CloudDesc = `
All the cloud operations use this as the parent command.

It has the subcommand to do all the cloud operations such as creating a team,login and initialize the cloud.

To do anything on the cloud you need to first login:

tykctl cloud login 

You can also set the default parameters to your config by running:

tykctl cloud init

`

// NewCloudCommand  creates the cloud service parent command.
func NewCloudCommand(cloudFactory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(cloudPath).
		WithDescription("Manage Tyk cloud resources (API Management on the Cloud).").
		WithLongDescription(CloudDesc).
		WithCommands(CloudCommands(cloudFactory)...)
}

func CloudCommands(cloudFactory internal.CloudFactory) []*cobra.Command {
	commands := []*cobra.Command{
		NewLoginCommand(cloudFactory),
		NewDeployment(cloudFactory),
		NewOrgCommand(cloudFactory),
		NewTeamCmd(cloudFactory),
		NewEnvironmentCmd(cloudFactory),
		NewZonesCmd(cloudFactory),
		NewInitCmd(cloudFactory),
	}

	return commands
}
