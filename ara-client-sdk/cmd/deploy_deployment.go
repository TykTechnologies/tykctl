/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/TykTechnologies/tykctl/swagger-gen"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const deployDesc = `
This command will deploy a Home or edge gateway.

Note: You need to first create a the Home or edge gateway before you can deploy it.
Use tykctl cloud deployment create to create the deployment.

The org,team,environment where the deployment was created has to be provided.

If org,team and environment are not set we will use the default set on your config file. 

You must also provide the uuid of the deployment you want to deploy.
to get the uuid run : tykctl cloud deployment list

Sample usage of this command:

tykctl cloud deployment deploy --org=<org here> --team=<team here> --env=<environment here> --uid=<deployment id>

`

// deployDeploymentCmd represents the deployDeployment command
var deployDeploymentCmd = &cobra.Command{
	Use:     "deploy",
	Short:   "Deploy a created deployment ",
	Long:    deployDesc,
	Example: `tykctl cloud deployment deploy --org=<org here> --team=<team here> --env=<environment here> --uid=<deployment id>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployDeployment called")
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
		deploymentId, err := cmd.Flags().GetString("uid")
		if err != nil {
			return
		}
		s.Prefix = "deploying deployments "
		s.Start()
		deployment, err := deploy(cmd.Context(), org, team, env, deploymentId)
		s.Stop()
		if err != nil {
			cmd.Println(parseError(err))
			return
		}
		cmd.Println(deployment.Status)
		cmd.Println(deployment.Payload.ZoneCode)
		cmd.Printf("Deployment %s deployed successfully", deployment.Payload.UID)
	},
}

func deploy(ctx context.Context, org, team, env, deploymentId string) (*swagger.InlineResponse2001, error) {
	deployment, _, err := client.DeploymentsApi.StartDeployment(ctx, org, team, env, deploymentId)
	if err != nil {
		return nil, err
	}

	return &deployment, err

}

func init() {
	deploymentsCmd.AddCommand(deployDeploymentCmd)
	deployDeploymentCmd.Flags().StringP("uid", "u", "", "Help message for toggle")
	deployDeploymentCmd.MarkFlagRequired("uid")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployDeploymentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployDeploymentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
