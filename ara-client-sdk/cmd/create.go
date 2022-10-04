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
	"tykcli/swagger-gen"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal("error reading name")
		}
		f, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println("wrong output format sent")
			return
		}
		org := swagger.Organisation{

			Name: name,
		}

		res, _, err := client.OrganisationsApi.CreateOrg(cmd.Context(), org)
		if err != nil {
			log.Fatal(err)
		}
		if f == "json" {
			marshal, err := json.Marshal(&res.Payload)
			if err != nil {
				log.Fatal(err)
			}
			internal.ShowJson(marshal)

		} else {
			orgs := []swagger.Organisation{
				*res.Payload,
			}
			internal.PrintOrganizationInTable(orgs)
		}
	},
}

func init() {
	orgCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	createCmd.Flags().StringP("name", "n", "name", "name of the organization to create")

	createCmd.Flags().StringP("output", "o", "table", "Format you want to use can be table,json")

	//err := createCmd.MarkPersistentFlagRequired("name")
	///if err != nil {
	//	return
	//}
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
