/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ara-client-sdk/swagger-gen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createTeamCmd represents the createTeam command
var createTeamCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		org := viper.GetString("org")
		if len(org) == 0 {
			cmd.Println("organization is required")
			return
		}
		teamName, err := cmd.Flags().GetString("name")
		if len(teamName) == 0 {
			cmd.Println("Team name is required")
			return
		}
		cmd.Println(teamName)
		team := swagger.Team{
			Name: teamName,
		}
		teams, _, err := client.TeamsApi.CreateTeam(cmd.Context(), team, org)
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
