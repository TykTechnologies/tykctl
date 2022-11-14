package sharedCmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		// Search config in home directory with name ".tykctl" (without extension).
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

// CreateConfigFile creates a file in a given directory is it does not exist.
func CreateConfigFile(dir, file string) error {
	result := filepath.Join(dir, file)
	_, err := os.Stat(result)
	if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(result)
		if err != nil {
			return err
		}
		defer f.Close()
		return nil
	}
	return err
}
