/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ara-client-sdk/swagger-gen"
	"context"
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// deployDeploymentCmd represents the deployDeployment command
var deployDeploymentCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		deployment, err := deploy(cmd.Context(), org, team, env, deploymentId)

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
