/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/TykTechnologies/tykctl/swagger-gen"
	"github.com/spf13/cobra"
)

// orgCmd represents the org command
const orgDesc = `
This is the parent command for all action that can be done to an organization.

The list of subcommand supported by this command are:
1. tykctl cloud org list - list all you organizations.
`

var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "Parent command for all org actions",
	Long:  orgDesc,
	Run: func(cmd *cobra.Command, args []string) {

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
