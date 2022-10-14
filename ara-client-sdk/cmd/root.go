/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/TykTechnologies/tykctl/swagger-gen"
	"github.com/briandowns/spinner"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var client *swagger.APIClient
var s *spinner.Spinner

const rootDesc = `
Tykctl is a cli that can be used to interact with all tyk components (tyk cloud,tyk gateway and tyk dashboard).

The cli is grouped into services.
For example to use the tyk cloud services you should prefix all your subcommands with:
tykcli cloud <subcommand here>

Currently we only support tyk cloud.

`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tykctl",
	Short: rootDesc,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	s = spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Color("white")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tykcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ara-client-sdk" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".tykctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	client = createClient()
}

func createClient() *swagger.APIClient {

	config := &swagger.Configuration{
		BasePath:      viper.GetString("controller"),
		DefaultHeader: map[string]string{},
	}
	///log.Printf(viper.GetString("token"))
	token := fmt.Sprintf("Bearer %s", viper.GetString("token"))
	///log.Println(token)
	config.AddDefaultHeader("Authorization", token)
	return swagger.NewAPIClient(config)

}
