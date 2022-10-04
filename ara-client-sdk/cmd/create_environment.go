/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"tykcli/swagger-gen"
)

// createEnvironmentCmd represents the createEnvironment command
var createEnvironmentCmd = &cobra.Command{
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
			cmd.Println("team to add this environment to is required")
			return
		}
		envName, err := cmd.Flags().GetString("name")
		if err != nil {

			cmd.Println(err)
			return
		}
		if len(envName) == 0 {
			cmd.Println("Environment name is required")
			return
		}
		loadOut := swagger.Loadout{
			Name: envName,
		}
		loadout, _, err := client.LoadoutsApi.CreateLoadout(cmd.Context(), loadOut, org, team)
		if err != nil {
			cmd.Println(err)
			return
		}
		cmd.Printf("Environment %s created successfully", loadout.Payload.UID)
	},
}

func init() {
	environmentCmd.AddCommand(createEnvironmentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createEnvironmentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createEnvironmentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createEnvironmentCmd.Flags().StringP("name", "n", "", "name to give the new environment")
	createEnvironmentCmd.MarkFlagRequired("name")
}
