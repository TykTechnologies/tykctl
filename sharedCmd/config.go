package sharedCmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		coreDir, err := getCoreDir()
		cobra.CheckErr(err)

		v, err := createViper(coreDir, coreConfig)
		cobra.CheckErr(err)

		currentConf := v.GetString(currentConfig)

		cobra.CheckErr(err)
		// Search config in home directory with name ".tykctl" (without extension).
		dir, err := getDefaultConfigDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(dir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(currentConf)
	}

	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func getCoreDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, defaultConfigDir)

	return configDir, nil
}

func getDefaultConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, defaultConfigDir, config)

	return configDir, nil
}

// CreateFile creates a file in a given directory is it does not exist.
func CreateFile(dir, file string) error {
	result := filepath.Join(dir, file)

	_, err := os.Stat(result)
	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	f, err := os.Create(result)
	if err != nil {
		return err
	}

	return f.Close()
}
