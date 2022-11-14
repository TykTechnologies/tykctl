package cmd

import (
	"context"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
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

func NewInitCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(initCloud).
		WithLongDescription(initDesc).
		WithExample("tykctl cloud init").
		WithDescription("initialize the cli and set the default region and organization.").
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			err := SetupPrompt(cmd.Context(), client)
			if err != nil {
				cmd.Println(err)
				return err
			}
			return nil
		})
}
func SetupPrompt(ctx context.Context, client internal.CloudClient) error {
	zones, _, err := client.GetDeploymentZones(ctx)
	if err != nil {
		return err
	}
	regions := make([]string, 0)
	for k := range zones.Payload.Tags {
		regions = append(regions, k)
	}
	selectedRegion, err := RegionPrompt(regions)
	if err != nil {
		return err
	}
	url, err := util.GenerateUrlFromZone(selectedRegion)
	if err != nil {
		return err
	}
	err = SaveToConfig(controller, url)
	if err != nil {
		return err
	}
	orgs, err := GetOrgs(ctx, client)
	if err != nil {
		return err
	}
	selectedOrg, err := OrgPrompt(orgs)
	if err != nil {
		return err
	}
	err = SaveToConfig(org, selectedOrg.UID)
	if err != nil {
		return err
	}
	selectedTeam, err := teamPrompt(selectedOrg.Teams)
	if err != nil {
		return err
	}
	err = SaveToConfig(team, selectedTeam.UID)
	if err != nil {
		return err
	}
	selectedEnv, err := envPrompt(selectedTeam.Loadouts)
	if err != nil {
		return err
	}
	err = SaveToConfig(env, selectedEnv.UID)
	if err != nil {
		return err
	}
	return nil

}
