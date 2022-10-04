/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cloudCmd represents the cloud command
var cloudCmd = &cobra.Command{
	Use:   "cloud",
	Short: "All the operation for the tyk cloud",
	Long:  `This is the parent command for all the tyk cloud commands.This include teams,deployments and environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cloud called")
	},
}

func init() {
	rootCmd.AddCommand(cloudCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloudCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloudCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
