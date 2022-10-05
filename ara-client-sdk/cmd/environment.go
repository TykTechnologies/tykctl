/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const enviromentDesc = `
This command is the parent command for all environment operations.

It has subcommands to:
 1. Create an environment: tykctl environment create --name="environment name"
 2. Fetch all environments in an org: tykctl environment list --org=<orgID>

`

// environmentCmd represents the environment command
var environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Create,list and delete environment",
	Long:  enviromentDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("environment called")
	},
}

func init() {
	cloudCmd.AddCommand(environmentCmd)
	environmentCmd.PersistentFlags().String("org", "", "The organization")
	environmentCmd.PersistentFlags().StringP("output", "o", "table", "Format you want to use can be table,json")
	viper.BindPFlag("org", environmentCmd.PersistentFlags().Lookup("org"))
	environmentCmd.PersistentFlags().String("team", "", "The team")
	viper.BindPFlag("team", environmentCmd.PersistentFlags().Lookup("team"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// environmentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// environmentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
