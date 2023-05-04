package cloudcmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
)

const orgDesc = `
This is the parent command for all actions that can be done within an organization.
The list of subcommand supported by this command are:

1. tykctl cloud orgs fetch - list all your organizations.

2. tykctl cloud orgs fetch [orgID] - list a single organization.
`

func NewOrgCommand(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(orgs).
		WithAliases([]string{org, "organization", "organisation"}).
		WithDescription("Parent command for all actions that can be done within an organization.").
		WithLongDescription(orgDesc).WithCommands(
		NewOrgListCommand(factory))
}
