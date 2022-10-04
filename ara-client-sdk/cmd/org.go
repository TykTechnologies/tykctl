/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"tykcli/swagger-gen"

	"github.com/spf13/cobra"
)

// orgCmd represents the org command
var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("org called")
	},
}

func init() {
	cloudCmd.AddCommand(orgCmd)
	orgCmd.PersistentFlags().StringP("output", "o", "table", "Format you want to use can be table,json")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// orgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// orgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func parseError(err error) string {
	message := err.Error()
	if myerr, ok := err.(swagger.GenericSwaggerError); ok {
		message = string(myerr.Body())
		// handle myerr
	}

	return message

}
