/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ara-client-sdk/swagger-gen"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		orgs, _, err := client.OrganisationsApi.GetOrgs(cmd.Context())
		if err != nil {
			cmd.Println(orgs.Error_)
			cmd.Println(orgs.Status)
			cmd.Println(orgs.Payload)
			cmd.Println(orgs.Status)
			cmd.Println(err)
			return
		}
		if len(orgs.Payload.Organisations) == 0 {
			cmd.Println("You have not created any organization yet")
			return
		}
		selectedOrg, err := organizationPrompt(orgs.Payload.Organisations)
		if err != nil {
			cmd.Println(err)
			return
		}
		cmd.Println(fmt.Sprintf("%s has been set a the default org", selectedOrg.Name))
		viper.Set("org", selectedOrg.UID)
		if len(selectedOrg.Teams) == 0 {
			cmd.Println("You have not created a team yet")
			return
		}
		selectedTeam, err := teamsPrompt(selectedOrg.Teams)
		if err != nil {
			cmd.Println(err)
			return
		}
		viper.Set("tid", selectedTeam.UID)
		cmd.Println(fmt.Sprintf("%s has been set a the default team", selectedTeam.Name))
		if err = viper.WriteConfig(); err != nil {
			log.Println("error writing to config file")
			return

		}
		//login if user has not logged in

		//log.Println(run)

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
