/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// environmentCmd represents the environment command
var environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("environment called")
	},
}

func init() {
	rootCmd.AddCommand(environmentCmd)

	environmentCmd.PersistentFlags().String("org", "", "The organization")
	environmentCmd.PersistentFlags().StringP("format", "f", "table", "Format you want to use can be table,json")
	viper.BindPFlag("org", teamCmd.PersistentFlags().Lookup("org"))
	environmentCmd.PersistentFlags().String("team", "", "The organization")
	viper.BindPFlag("team", environmentCmd.PersistentFlags().Lookup("team"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// environmentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// environmentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
