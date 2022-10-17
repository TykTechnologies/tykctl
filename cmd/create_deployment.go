/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/TykTechnologies/tykctl/swagger-gen"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const createDeploymentDesc = ` 
This command creates a Home or a Edge Gateway.

NOTE: This does not deploy the deployment it just create it.You can use the deploy command to deploy the created deployment.

You must pass the organization,team,zone and environment you want deploy this deployment.

NOTE: For the home deployment you have to select the you home zone as the deployment zone.

If you do not pass the org,zone or environment we will use the ones on your config file.You can set the default org,team and environment by running:

tykctl cloud init

Sample usage for this command

tykctl cloud deployment create --name="test deployment" --kind="Home"

`

// createDeploymentCmd represents the createDeployment command
var createDeploymentCmd = &cobra.Command{
	Use:   "create",
	Short: "create a deployment",
	Long:  createDeploymentDesc,
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
		s.Prefix = "creating deployment "
		s.Start()
		deployRsponse, _, err := client.DeploymentsApi.CreateDeployment(cmd.Context(), deployment, org, team, env)
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
		cmd.Printf("Deployment %s created successfully", deployRsponse.Payload.UID)

		shouldDeploy, err := cmd.Flags().GetBool("deploy")
		if err != nil {
			cmd.Println(err)
			return
		}
		if shouldDeploy {
			s.Prefix = "deploying deployment "
			s.Start()
			deployRes, err := deploy(cmd.Context(), org, team, env, deployRsponse.Payload.UID)
			s.Stop()
			if err != nil {
				message := parseError(err)
				cmd.Println(message)
				return
			}
			cmd.Printf("%s deployment has started", deployRes.Payload.UID)
		}
		//here deploy

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
	viper.BindPFlag("zone", createDeploymentCmd.Flags().Lookup("zone"))
	createDeploymentCmd.Flags().BoolP("deploy", "d", false, "deploy the deployment after creation")
}
