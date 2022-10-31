package cmd

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Flag struct {
	Description string
	Name        string
	Shorthand   string
	Value       string
	Default     string
}

func addOutPutFlags(f *pflag.FlagSet) {
	f.StringP(outPut, "o", "table", "Format you want to use can be table,json")

}

// SaveToConfig writes data to the config file provided by --config
func SaveToConfig(key, token string) error {
	viper.Set(key, token)
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("Couldn't write config: %s\n", err.Error())
	}
	return nil

}
