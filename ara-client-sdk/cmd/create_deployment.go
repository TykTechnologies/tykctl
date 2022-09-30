/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ara-client-sdk/swagger-gen"
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// createDeploymentCmd represents the createDeployment command
var createDeploymentCmd = &cobra.Command{
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
		team := viper.GetString("team")
		if len(team) == 0 {
			cmd.Println("team to add this deployment to is required")
			return
		}
		env := viper.GetString("env")
		if len(team) == 0 {
			cmd.Println("environment is needed")
			return
		}
		kind, err := cmd.Flags().GetString("kind")
		if err != nil {
			cmd.Println(err)
			return
		}
		if len(kind) == 0 {
			cmd.Println("kind is required")
			return
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			cmd.Println(err)
			return
		}

		zone := viper.GetString("zone")

		fmt.Println("createDeployment called")
		deployment := swagger.Deployment{
			Kind: kind,

			Name: name,

			ZoneCode: zone,
		}

		deploy, _, err := client.DeploymentsApi.CreateDeployment(cmd.Context(), deployment, org, team, env)
		if err != nil {
			message := err.Error()
			if myerr, ok := err.(swagger.GenericSwaggerError); ok {
				message = string(myerr.Body())
				// handle myerr
			}

			cmd.Println(message)
			return

		}
		cmd.Printf("Deployment %s created successfully", deploy.Payload.UID)

	},
}

func init() {
	deploymentsCmd.AddCommand(createDeploymentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createDeploymentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createDeploymentCmd.Flags().StringP("kind", "k", "Home", "Help message for toggle")
	createDeploymentCmd.MarkFlagRequired("kind")
	createDeploymentCmd.Flags().StringP("name", "n", "", "name to give the new team")
	createDeploymentCmd.MarkFlagRequired("name")
	createDeploymentCmd.Flags().StringP("zone", "z", "", "zone you want to deploy into")
	viper.BindPFlag("zone", createDeploymentCmd.PersistentFlags().Lookup("zone"))
}
