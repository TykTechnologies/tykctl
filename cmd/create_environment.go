/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/TykTechnologies/tykctl/swagger-gen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const createEnvDec = `
This command create an environment in a team.

You must pass the name of the environment.

You must also set the org and team you want to create this environment in.

If you don't pass the org and team we will use the one set in the config file.

Sample usage:

tyk cloud environment create --name="staging"
`

// createEnvironmentCmd represents the createEnvironment command
var createEnvironmentCmd = &cobra.Command{
	Use:     "create",
	Short:   "create an environment",
	Long:    createEnvDec,
	Example: `tyk cloud environment create --name="staging"`,
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
		s.Prefix = "creating environment "
		s.Start()
		loadout, _, err := client.LoadoutsApi.CreateLoadout(cmd.Context(), loadOut, org, team)
		s.Stop()
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
