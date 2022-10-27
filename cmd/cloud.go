/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

const CloudDesc = `
All the cloud operations use this as the parent command.

It has the subcommand to do all the cloud operations such as creating a team,login and initialize the cloud.

To do anything on the cloud you need to first login:

tykctl cloud login 

You can also set the default parameters to your config by running:

tykctl cloud init

`

func NewCloudCommand(client internal.CloudClient) *cobra.Command {
	return NewCmd("cloud").
		WithDescription("All the operation for the tyk cloud").
		WithLongDescription(CloudDesc).WithCommands(NewLoginCommand(), NewOrgCommand(client))
}
