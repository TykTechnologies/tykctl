package sharedCmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/TykTechnologies/tykctl/internal"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		coreDir, err := internal.GetCoreDir()
		cobra.CheckErr(err)

		v, err := internal.CreateViper(coreDir, internal.CoreConfig)
		cobra.CheckErr(err)

		currentConf := v.GetString(internal.CurrentConfig)

		cobra.CheckErr(err)
		// Search config in home directory with name ".tykctl" (without extension).
		dir, err := internal.GetDefaultConfigDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(dir)
		viper.SetConfigType("yaml")

		viper.SetConfigName(fmt.Sprintf("config_%s", currentConf))
	}

	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
