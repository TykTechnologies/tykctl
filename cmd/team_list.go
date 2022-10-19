/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// listCmd represents the list command
const teamListDesc = `
This command will fetch and list all the teams in an organization.

You must pass the --org flag.If it is not passed we will use the default one set in the config file.

The output can be either json or table. Default is table.
To change the format use --output=<json/table> flag.

Sample usage:

tykctl team list --org=<orgID> --output=<json/table>

`

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Fetch all teams in an organization",
	Long:    teamListDesc,
	Example: `tykctl team list --org=<orgID> --output=<json/table>`,
	Run: func(cmd *cobra.Command, args []string) {
		org := viper.GetString("org")
		if len(org) == 0 {
			cmd.Println("organization is required")
			return
		}
		s.Prefix = "fetching teams "
		s.Start()

		teams, _, err := client.TeamsApi.GetTeams(cmd.Context(), org)
		s.Stop()
		if err != nil {
			cmd.Println(err)
			return
		}

		f, err := cmd.Flags().GetString("output")
		if err != nil {
			cmd.Println(err)
			cmd.Println("wrong output format sent")
			return
		}
		if f == "json" {
			marshal, err := json.Marshal(&teams.Payload)
			if err != nil {
				log.Println(err)
				return
			}
			internal.ShowJson(marshal)

			return
		}

		internal.PrintTeamInTable(teams.Payload.Teams)

	},
}

func init() {
	teamCmd.AddCommand(listCmd)
	////listCmd.MarkFlagRequired("org")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
