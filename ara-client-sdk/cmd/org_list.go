/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// orgListCmd represents the orgList command
var orgListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		orgs, _, err := client.OrganisationsApi.GetOrgs(cmd.Context())
		if err != nil {
			cmd.Println(err)
			return
		}
		if len(orgs.Payload.Organisations) > 0 {
			cmd.Println(orgs.Payload.Organisations[0].UID)

		}
		cmd.Println(len(orgs.Payload.Organisations))
	},
}

func init() {
	orgCmd.AddCommand(orgListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// orgListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// orgListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
