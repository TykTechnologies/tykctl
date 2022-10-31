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

// addOrgFlag adds the org flag to the flags
// it uses the default org set in you config file if it is not passed.
func addOrgFlag(f *pflag.FlagSet) {
	f.String(org, "", "The organization to use")
	viper.BindPFlag(org, f.Lookup(org))
}

func addNameFlag(f *pflag.FlagSet) {
	f.StringP(name, "n", "", "name for the object you want to create")
}
func addOutPutFlags(f *pflag.FlagSet) {
	f.StringP(outPut, "o", "table", "Format you want to use can be table,json")

}

// SaveToConfig writes data to the config file provided by --config
func SaveToConfig(key, value string) error {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("Couldn't write config: %s\n", err.Error())
	}
	return nil

}
