/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"tykcli/internal"
	"tykcli/swagger-gen"
)

const listDeploymentDesc = `
This command will list all the deployment belonging to an organization

You will need to pass the org whose deployment you want to get.
If you don't pass the org we will use the one in your config file.

You can either get the data as json or in a table.

Use the --output<table,json> flag to change the format default is table.

Sample usage: 

tykctl cloud deployment list --org=<organization id> --output=<json/table>

`

// deploymentListCmd represents the deploymentList command
var deploymentListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Fetch all deployment in an organization",
	Long:    listDeploymentDesc,
	Example: `tykctl cloud deployment list --org=<orgID> --output=<json/table>`,
	Run: func(cmd *cobra.Command, args []string) {
		org := viper.GetString("org")
		if len(org) == 0 {
			cmd.Println("organization is required")
			return
		}
		s.Prefix = "fetching deployments "
		s.Start()
		deployment, _, err := client.OrganisationsApi.GetDeploymentsForOrg(cmd.Context(), org)
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
			marshal, err := json.Marshal(&deployment.Payload)
			if err != nil {
				log.Println(err)
				return
			}
			internal.ShowJson(marshal)

			return
		}
		internal.PrintDeploymentInTable(deployment.Payload.Deployments)
		///team := viper.GetString("team")
		//if len(team) == 0 {
		//	cmd.Println("team is required")
		///	return
		//}
		//env := viper.GetString("env")
		///if len(env) == 0 {
		//cmd.Println("environment is required")
		//return
		//}

	},
}

func init() {
	deploymentsCmd.AddCommand(deploymentListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deploymentListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deploymentListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//return an message and a boolean showing if it is has all fields
func validateDeploymentFields() error {
	org := viper.GetString("org")
	if len(org) == 0 {
		return errors.New("organization is required")
	}
	team := viper.GetString("team")
	if len(team) == 0 {
		return errors.New("team is required")
	}

	return nil
}
