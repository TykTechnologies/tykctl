/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"tykcli/swagger-gen"
)

const createTeamDesc = `
This command will create a team.

You have to pass the name you want to give the team and org in which you want to create the team.

If the org is not provided we will use the one you set in the config file.

To set a default team in the config file run:

tykctl cloud init

Sample usage for this command:

tyckctl cloud team create --name="first team" --org=<org uuid>

`

// createTeamCmd represents the createTeam command
var createTeamCmd = &cobra.Command{
	Use:     "create",
	Short:   "create a team",
	Long:    createTeamDesc,
	Example: `tyckctl cloud team create --name="first team" --org=<org uuid>`,
	Run: func(cmd *cobra.Command, args []string) {
		org := viper.GetString("org")
		if len(org) == 0 {
			cmd.Println("organization is required")
			return
		}
		teamName, err := cmd.Flags().GetString("name")
		if err != nil {

			cmd.Println(err)
			return
		}

		if len(teamName) == 0 {
			cmd.Println("Team name is required")
			return
		}
		cmd.Println(teamName)
		team := swagger.Team{
			Name: teamName,
		}
		s.Prefix = "creating teams "
		s.Start()
		teams, _, err := client.TeamsApi.CreateTeam(cmd.Context(), team, org)
		s.Stop()
		if err != nil {

			cmd.Println(err)
			return
		}
		cmd.Printf("Team %s created successfully", teams.Payload.UID)
	},
}

func init() {
	teamCmd.AddCommand(createTeamCmd)
	createTeamCmd.Flags().StringP("name", "n", "", "name to give the new team")
	createTeamCmd.MarkFlagRequired("name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createTeamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createTeamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
