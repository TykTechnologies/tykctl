/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"tykcli/internal"
	"tykcli/swagger-gen"
)

const enviromentListDesc = `
This command will fetch all the environment in an organization.

You must pass the --org.If it is not passed we will use the default org set in your config file.

We support json and table as the output format.To set the output format use the --output<json/table> flag.

Sample usage of this command:

tckctl cloud environment list --org=<orgID> --output=<json/table>

`

var enviromentListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Fetch all environments in an organization.",
	Long:    enviromentListDesc,
	Example: `tckctl cloud environment list --org=<orgID> --output=<json/table>`,
	Run: func(cmd *cobra.Command, args []string) {
		org := viper.GetString("org")
		if len(org) == 0 {
			cmd.Println("organization is required")
			return
		}
		s.Prefix = "fetching environments "
		s.Start()
		loadouts, _, err := client.LoadoutsApi.GetOrgLoadouts(cmd.Context(), org)
		s.Stop()
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
			marshal, err := json.Marshal(&loadouts.Payload.Loadouts)
			if err != nil {
				log.Println(err)
				return
			}
			internal.ShowJson(marshal)

			return
		}
		internal.PrintLoadOutInTable(loadouts.Payload.Loadouts)
	},
}

func init() {
	environmentCmd.AddCommand(enviromentListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enviromentListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enviromentListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
