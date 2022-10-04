/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"tykcli/internal"
	"tykcli/swagger-gen"
)

// enviromentListCmd represents the enviromentList command
var enviromentListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		org := viper.GetString("org")
		if len(org) == 0 {
			cmd.Println("organization is required")
			return
		}
		loadouts, _, err := client.LoadoutsApi.GetOrgLoadouts(cmd.Context(), org)
		if err != nil {
			message := err.Error()
			if myerr, ok := err.(swagger.GenericSwaggerError); ok {
				message = string(myerr.Body())
				// handle myerr
			}
			cmd.Println(message)
			return
		}
		f, err := cmd.Flags().GetString("format")
		if err != nil {
			cmd.Println(err)
			cmd.Println("wrong output format sent")
			return
		}
		if f == "json" {
			marshal, err := json.Marshal(&loadouts.Payload.Loadouts)
			if err != nil {
				log.Println(err)
				return
			}
			internal.ShowJson(marshal)

			return
		}
		internal.PrintLoadOutInTable(loadouts.Payload.Loadouts)
	},
}

func init() {
	environmentCmd.AddCommand(enviromentListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enviromentListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enviromentListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
