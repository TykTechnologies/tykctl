package cloudcmd

import (
	"context"
	"errors"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const initDesc = `
This command will initialize the cli and set default in the config file.
Before using this command you will need to login with:

tykctl cloud login

Use this command to:

1. Set the default organization.

2. Set the default team

3. Set the default environment.

4. Set your zone and home region.
This command should ideally be run after the login command.
`

func NewInitCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(initCloud).
		WithLongDescription(initDesc).
		WithExample("tykctl cloud init").
		WithDescription("initialize the cli and set the default region and organization.").
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			userId := viper.GetString(currentCloudUser)
			if userId == "" {
				cmd.Println("Please login in first before running thid command")
				return errors.New("you need to login to run this command")
			}
			err := SetupPrompt(cmd.Context(), factory.Client, factory.Prompt, factory.Config.GetCurrentUserOrg(), userId)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			cmd.Println("Config file initialized successfully")
			return nil
		})
}
func SetupPrompt(ctx context.Context, client internal.CloudClient, prompt internal.CloudPrompt, orgId, userId string) error {
	info, _, err := client.GetOrgInfo(ctx, orgId)
	if err != nil {
		return err
	}
	selectedTeam, err := prompt.TeamPrompt(info.Organisation.Teams)
	if err != nil {
		return err
	}
	var orgInit internal.OrgInit
	if selectedTeam != nil {
		orgInit.Team = selectedTeam.UID
		selectedEnv, err := prompt.EnvPrompt(selectedTeam.Loadouts)
		if err != nil {
			return err
		}
		if selectedEnv != nil {
			orgInit.Env = selectedEnv.UID
		}
	}
	err = internal.SaveMapToCloudUserContext(userId, orgInit.OrgInitToMap())
	if err != nil {
		return err
	}
	return nil
}
