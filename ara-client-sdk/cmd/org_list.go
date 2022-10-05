/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
	"tykcli/internal"
	"tykcli/swagger-gen"
)

// orgListCmd represents the orgList command

const orgListDesc = `
This command will list all your organizations.

Currently you can only be part of one organization hence we will return a single organization.

Sample command usage:
tykctl cloud org list --output<json/table>

You can get the output either in table or json format.The default is table format.

user the --output flag to change the format.

`

var orgListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all your organizations.",
	Long:    orgListDesc,
	Example: "tykctl cloud org list --output<json/table>",
	Run: func(cmd *cobra.Command, args []string) {
		orgs, _, err := client.OrganisationsApi.GetOrgs(cmd.Context())
		if err != nil {
			message := err.Error()
			if myerr, ok := err.(swagger.GenericSwaggerError); ok {
				message = string(myerr.Body())
				// handle myerr
			}
			cmd.Println(message)
			return
		}
		f, err := cmd.Flags().GetString("output")
		if err != nil {
			cmd.Println(err)
			cmd.Println("wrong output format sent")
			return
		}
		if f == "json" {
			marshal, err := json.Marshal(&orgs.Payload.Organisations)
			if err != nil {
				log.Println(err)
				return
			}
			internal.ShowJson(marshal)

			return
		}
		internal.PrintOrganizationInTable(orgs.Payload.Organisations)

	},
}

func init() {
	orgCmd.AddCommand(orgListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// orgListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// orgListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
