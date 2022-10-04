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

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := cmd.Flags().GetString("format")
		if err != nil {
			fmt.Println("wrong output format sent")
			return
		}

		orgs, _, err := client.OrganisationsApi.GetOrgs(cmd.Context())
		if err != nil {
			log.Println(err)
			return
		}
		if f == "json" {
			marshal, err := json.Marshal(&orgs.Payload)
			if err != nil {
				log.Println(err)
				return
			}
			internal.ShowJson(marshal)
		} else {
			internal.PrintOrganizationInTable(orgs.Payload.Organisations)
		}

		///s, _ := colorjson.Marshal(len())

	},
}

func init() {
	orgCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringP("format", "f", "table", "Format you want to use can be table,json")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
