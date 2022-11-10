package cmd

import (
	"context"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
)

func NewIntCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(initCloud).
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
