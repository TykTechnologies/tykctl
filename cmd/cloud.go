/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

const CloudDesc = `
All the cloud operations use this as the parent command.

It has the subcommand to do all the cloud operations such as creating a team,login and initialize the cloud.

To do anything on the cloud you need to first login:

tykctl cloud login 

You can also set the default parameters to your config by running:

tykctl cloud init

`

func NewCloudCommand() *cobra.Command {
	return NewCmd("cloud", "cloud").
		WithDescription("All the operation for the tyk cloud").
		WithLongDescription(CloudDesc).Hidden().WithCommands(NewLoginCommand())

}

/*// cloudCmd represents the cloud command
var cloudCmd = &cobra.Command{
	Use:     "cloud",
	Short:   "All the operation for the tyk cloud",
	Example: "tykctl cloud org list",
	Long:    CloudDesc,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
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
*/
