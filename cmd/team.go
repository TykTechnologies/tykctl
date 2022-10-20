/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

const teamDesc = `
This command is the parent command to all team operations.

The supported commands are:
1. tykctl cloud team list -fetch all org teams
2. tykctl cloud team create -create a team in an org.

All subcommands require an org id.If it is not passed we use the default one in the config file.

To set the default org run :
tykctl cloud init

`

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	ValidArgs: []string{"list", "create"},
	Use:       "team",
	Short:     "Create,list teams in an organization",
	Long:      teamDesc,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	///cloudCmd.AddCommand(teamCmd)

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
