/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/TykTechnologies/tykctl/swagger-gen"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialization the cli",
	Long:  initDesc,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		if len(token) == 0 {
			cmd.Println("Please login first using the command aractl login ")
			return
		}
		err := regionPrompt()
		if err != nil {
			cmd.Println(err)
			return
		}
		///then select the default organization
		s.Prefix = "fetching organizations "
		s.Start()
		orgs, _, err := client.OrganisationsApi.GetOrgs(cmd.Context())
		if err != nil {
			cmd.Println(orgs.Error_)
			cmd.Println(orgs.Status)
			cmd.Println(orgs.Payload)
			cmd.Println(orgs.Status)
			cmd.Println(err)
			return
		}
		s.Stop()
		if len(orgs.Payload.Organisations) == 0 {
			cmd.Println("You have not created any organization yet")
			return
		}
		selectedOrg, err := organizationPrompt(orgs.Payload.Organisations)
		if err != nil {
			cmd.Println(err)
			return
		}

		viper.Set("org", selectedOrg.UID)
		if err = viper.WriteConfig(); err != nil {
			log.Println("error writing to config file")
			return

		}
		cmd.Println(fmt.Sprintf("%s has been set a the default org", selectedOrg.Name))
		if len(selectedOrg.Teams) == 0 {
			cmd.Println("You have not created a team yet")
			return
		}
		selectedTeam, err := teamsPrompt(selectedOrg.Teams)
		if err != nil {
			cmd.Println(err)
			return
		}
		viper.Set("team", selectedTeam.UID)
		if err = viper.WriteConfig(); err != nil {
			log.Println("error writing to config file")
			return

		}
		cmd.Println(fmt.Sprintf("%s has been set a the default team", selectedTeam.Name))
		if len(selectedTeam.Loadouts) == 0 {
			cmd.Println("You have not created an environment yet")
			return
		}
		selectedEnv, err := envPrompt(selectedTeam.Loadouts)
		if err != nil {
			cmd.Println(err)
			return
		}
		viper.Set("env", selectedEnv.UID)
		if err = viper.WriteConfig(); err != nil {
			log.Println("error writing to config file")
			return

		}
		cmd.Println(fmt.Sprintf("%s has been set a the default enviroment", selectedEnv.Name))
		//login if user has not logged in
		cmd.Println("The cli is now initialized you can start creating deployment")

		//log.Println(run)

	},
}

func init() {
	cloudCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func envPrompt(loadouts []swagger.Loadout) (*swagger.Loadout, error) {
	allEnv := make([]string, len(loadouts))
	for index, env := range loadouts {
		allEnv[index] = fmt.Sprintf("%s - (%d-Deployments)", env.Name, len(env.Deployments))

	}
	envPrompt := promptui.Select{
		Label: "Select your default environment",
		Items: allEnv,
	}
	index, _, err := envPrompt.Run()
	if err != nil {
		return nil, err
	}

	return &loadouts[index], err

}
func teamsPrompt(teams []swagger.Team) (*swagger.Team, error) {
	allTeams := make([]string, len(teams))
	for index, team := range teams {

		allTeams[index] = fmt.Sprintf("%s - (%d-enviroments)", team.Name, len(team.Loadouts))
	}
	teamPrompt := promptui.Select{
		Label: "Select your default Team",
		Items: allTeams,
	}

	index, _, err := teamPrompt.Run()
	if err != nil {
		return nil, err
	}
	return &teams[index], err

}

func organizationPrompt(orgs []swagger.Organisation) (*swagger.Organisation, error) {

	allOrganizations := make([]string, len(orgs))
	//i := 0
	for index, org := range orgs {
		allOrganizations[index] = fmt.Sprintf("%s -(%d-teams)", org.Name, len(org.Teams))

	}
	orgPrompt := promptui.Select{
		Label: "Select your default organization",
		Items: allOrganizations,
	}

	index, _, err := orgPrompt.Run()
	if err != nil {
		return nil, err
	}
	return &orgs[index], err

}

func regionPrompt() error {
	araRegion := make([]string, len(regions))
	i := 0
	for k := range regions {
		araRegion[i] = k
		i++
	}
	regionPrompt := promptui.Select{
		Label: "Select your home region",
		Items: araRegion,
	}
	_, s, err := regionPrompt.Run()
	if err != nil {
		log.Println("error selecting a region")
		return err
	}
	viper.Set("controller", regions[s])
	viper.Set("dashboard", "https://dash.ara-staging.tyk.technology")
	if err = viper.WriteConfig(); err != nil {
		log.Println("error setting url")
		return err

	}
	log.Println("current regions has been set to:", s)
	return err
}

///for now we will list the zones in a map but they can be downloaded

var regions = map[string]string{
	"staging-aws-eun1": "https://controller-aws-eun1.ara-staging.tyk.technology:37001",
	"staging-aws-use2": "https://controller-aws-use2.ara-staging.tyk.technology:37001",
	"staging-aws-use1": "https://controller-aws-use1.ara-staging.tyk.technology",
}

func loginPrompt() (string, error) {
	prompt := promptui.Prompt{
		Label:    "You need to first login before you continue.Would you like to login y/n?",
		Validate: validateNotEmpty,
	}
	return prompt.Run()
}
