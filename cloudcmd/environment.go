package cloudcmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
)

const environmentDesc = `This is the parent command to all environment operations.

   Note: All subcommands for this command must pass a --team  and --org flag command.
  
`

func NewEnvironmentCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(environments).
		WithAliases([]string{env}).
		WithFlagAdder(true, addOrgFlag).
		WithLongDescription(environmentDesc).
		WithDescription("parent command to all environment operations").
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: org, Persistent: true}, {Name: team, Persistent: true}}).
		WithFlagAdder(true, addTeamFlag).
		WithCommands(
			NewCreateEnvironmentCmd(factory),
			NewFetchEnvironmentCmd(factory),
			NewUpdateEnvCmd(factory),
		)
}

func validateCommonEnvFlags(Config internal.UserConfig, cmd cobra.Command) (*CommonEnvFlags, error) {
	orgId := Config.GetCurrentUserOrg()
	if util.StringIsEmpty(orgId) {
		return nil, ErrorOrgRequired
	}
	teamId := Config.GetCurrentUserTeam()
	if util.StringIsEmpty(teamId) {
		return nil, ErrorTeamRequired
	}
	envName, err := cmd.Flags().GetString(name)
	if err != nil {
		return nil, err
	}
	if util.StringIsEmpty(envName) {
		return nil, ErrorNameRequired
	}
	return &CommonEnvFlags{
		orgId:   orgId,
		teamId:  teamId,
		envName: envName,
	}, nil
}

type CommonEnvFlags struct {
	orgId   string
	teamId  string
	envName string
}
