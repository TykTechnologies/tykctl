/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const deploymentDesc = `
This command has subcommand to create,deploy,restart and fetch deployments.

Supported subcommand: 
    - tykctl cloud deployment list
    - tykctl cloud deployment create

Most of the subcommand will require the orgID and teamID flags.

This can also be set in the config file by running:

tykctl cloud init.

`

// deploymentsCmd represents the deployments command
var deploymentsCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Parent command for all deployment actions",
	Long:  deploymentDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployments called")

	},
}

func init() {
	cloudCmd.AddCommand(deploymentsCmd)
	deploymentsCmd.PersistentFlags().String("org", "", "The organization")
	deploymentsCmd.PersistentFlags().StringP("output", "o", "table", "Format you want to use can be table,json")
	viper.BindPFlag("org", deploymentsCmd.PersistentFlags().Lookup("org"))
	deploymentsCmd.PersistentFlags().String("team", "", "The team")
	viper.BindPFlag("team", deploymentsCmd.PersistentFlags().Lookup("team"))
	deploymentsCmd.PersistentFlags().String("env", "", "The environment")
	viper.BindPFlag("env", environmentCmd.PersistentFlags().Lookup("env"))
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deploymentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deploymentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
