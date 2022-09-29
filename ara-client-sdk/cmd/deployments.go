/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// deploymentsCmd represents the deployments command
var deploymentsCmd = &cobra.Command{
	Use:   "deployment",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployments called")

	},
}

func init() {
	rootCmd.AddCommand(deploymentsCmd)
	deploymentsCmd.PersistentFlags().String("org", "", "The organization")
	deploymentsCmd.PersistentFlags().StringP("format", "f", "table", "Format you want to use can be table,json")
	viper.BindPFlag("org", deploymentsCmd.PersistentFlags().Lookup("org"))
	deploymentsCmd.PersistentFlags().String("team", "", "The team")
	viper.BindPFlag("team", deploymentsCmd.PersistentFlags().Lookup("team"))
	environmentCmd.PersistentFlags().String("env", "", "The environment")
	viper.BindPFlag("env", environmentCmd.PersistentFlags().Lookup("env"))
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deploymentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deploymentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
