package cloudcmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
)

const orgDesc = `
This is the parent command for all action that can be done to an organization.
The list of subcommand supported by this command are:

1. tykctl cloud orgs fetch - list all you organizations.

2. tykctl cloud orgs fetch [orgID] - list a single organizations.
`

func NewOrgCommand(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(orgs).
		WithAliases([]string{org, "organization", "organisation"}).
		WithDescription("This is the parent command for all action that can be done to an organization.").
		WithLongDescription(orgDesc).WithCommands(
		NewOrgListCommand(factory))
}
