package cloudcmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
)

const environmentDesc = `This is the parent command to all environment operations.

   Note: All subcommands for this command must pass a --team  and --org flag command.
  
`

func NewEnvironmentCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(environments).
		WithAliases([]string{env}).
		WithLongDescription(environmentDesc).
		WithDescription("Parent command for all environment operations.").
		WithBindFlagOnPreRun([]internal.BindFlag{{Name: org, Persistent: true, Type: internal.Cloud}, {Name: team, Persistent: true, Type: internal.Cloud}}).
		WithCommands(
			NewCreateEnvironmentCmd(factory),
			NewFetchEnvironmentCmd(factory),
			NewDeleteEnvCmd(factory),
		)
}

func validateCommonEnvFlags(config internal.UserConfig) (*EnvFlags, error) {
	var envFlags EnvFlags

	envFlags.OrgID = config.GetCurrentUserOrg()
	if util.StringIsEmpty(envFlags.OrgID) {
		return nil, ErrorOrgRequired
	}

	envFlags.TeamID = config.GetCurrentUserTeam()
	if util.StringIsEmpty(envFlags.TeamID) {
		return nil, ErrorTeamRequired
	}

	return &envFlags, nil
}

type EnvFlags struct {
	OrgID  string
	TeamID string
	OutPut string
}
