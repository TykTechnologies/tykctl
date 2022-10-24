package cmd

import "github.com/spf13/cobra"

const orgDesc = `
This is the parent command for all action that can be done to an organization.
The list of subcommand supported by this command are:
1. tykctl cloud org list - list all you organizations.
`

func NewOrgCommand() *cobra.Command {
	return NewCmd("org").
		WithLongDescription(orgDesc).WithCommands(NewOrgListCommand(), NewOrgCreateCmd())

}
