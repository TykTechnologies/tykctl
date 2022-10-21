/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
)

var cfgFile string

const rootDesc = `
Tykctl is a cli that can be used to interact with all tyk components (tyk cloud,tyk gateway and tyk dashboard).

The cli is grouped into services.
For example to use the tyk cloud services you should prefix all your subcommands with:
tykcli cloud <subcommand here>

Currently we only support tyk cloud.

`

func NewRootCmd() *cobra.Command {
	return NewCmd("tykctl").WithLongDescription(rootDesc).
		WithDescription("access all tyk service via the cli").
		WithFlagAdder(true, addGlobalPersistentFlags).
		WithFlagAdder(false, addRootLocalFlags).
		WithCommands(NewCloudCommand())
}

func Execute() {
	err := NewRootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addGlobalPersistentFlags(f *pflag.FlagSet) {
	f.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tykctl.yaml)")
}

func addRootLocalFlags(f *pflag.FlagSet) {
	f.BoolP("toggle", "t", false, "Help message for toggle")
}

func init() {
	err := createConfigFile()
	if err != nil {
		log.Fatal(err)
	}
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	//cobra.CheckErr(err)
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

}
func createConfigFile() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	result := fmt.Sprintf("%s/%s", home, ".tykctl.yaml")
	_, err = os.Stat(result)
	if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(result)
		log.Println(err)
		if err != nil {
			return err
		}
		defer f.Close()
		return nil

	}
	return err

}
