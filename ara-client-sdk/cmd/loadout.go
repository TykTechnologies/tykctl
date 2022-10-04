/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"tykcli/internal"
)

// loadoutCmd represents the loadout command
var loadoutCmd = &cobra.Command{
	Use:   "loadout",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		///fmt.Println("loadout called")
		name, err := cmd.Flags().GetString("org")
		if err != nil {
			log.Println("error reading name")
			return
		}
		if name == "" {
			log.Println("org id is required")
			return
			///return
		}
		log.Println(name)
		f, err := cmd.Flags().GetString("format")
		if err != nil {
			fmt.Println("wrong output format sent")
			return
		}

		loadOut, _, err := client.LoadoutsApi.GetOrgLoadouts(cmd.Context(), "0e7f69c6-2720-4821-9349-195a1dec7eb5")

		if err != nil {
			log.Println("error getting loadout")
			return
		}
		if len(loadOut.Payload.Loadouts) == 0 {
			log.Println("this org does not have a loud-out")
			return
		}
		if f == "json" {
			marshal, err := json.Marshal(&loadOut.Payload)
			if err != nil {
				log.Println("error converting to bytes")
				return
			}
			internal.ShowJson(marshal)

		} else {
			internal.PrintLoadOutInTable(loadOut.Payload.Loadouts)
		}

	},
}

func init() {
	orgCmd.AddCommand(loadoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	loadoutCmd.Flags().StringP("format", "f", "table", "Format you want to use can be table,json")

	loadoutCmd.Flags().StringP("org", "o", "name", "Help message for toggle")
}
