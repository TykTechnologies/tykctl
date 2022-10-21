package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

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
