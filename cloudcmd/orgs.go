package cloudcmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

const orgDesc = `
This is the parent command for all action that can be done to an organization.
The list of subcommand supported by this command are:

1. tykctl cloud orgs fetch - list all you organizations.

2. tykctl cloud orgs fetch [orgId] - list a single organizations.
`

func NewOrgCommand(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(orgs).
		WithDescription("This is the parent command for all action that can be done to an organization.").
		WithLongDescription(orgDesc).WithCommands(
		NewOrgListCommand(factory))
}
