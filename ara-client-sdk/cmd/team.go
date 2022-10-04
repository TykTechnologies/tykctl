/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	ValidArgs: []string{"list", "create"},
	Use:       "team",
	Short:     "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("team called")
	},
}

func init() {
	cloudCmd.AddCommand(teamCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	teamCmd.PersistentFlags().String("org", "", "The organization")
	teamCmd.PersistentFlags().StringP("output", "o", "table", "Format you want to use can be table,json")
	viper.BindPFlag("org", teamCmd.PersistentFlags().Lookup("org"))
	///teamCmd.MarkPersistentFlagRequired("org")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// teamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
