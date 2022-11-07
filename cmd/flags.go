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
type BindFlag struct {
	Name       string
	Persistent bool
}

// addOrgFlag adds the org flag to the flags
// it uses the default org set in you config file if it is not passed.
func addOrgFlag(f *pflag.FlagSet) {
	f.String(org, "", "The organization to use")
	err := viper.BindPFlag(org, f.Lookup(org))
	if err != nil {
		panic(err)
	}
}
func addTeamFlag(f *pflag.FlagSet) {
	f.String(team, "", "The team to use")
	err := viper.BindPFlag(team, f.Lookup(team))
	if err != nil {
		panic(err)
	}
}
func addEnvFlag(f *pflag.FlagSet) {
	f.String(env, "", "The environment to use")
	err := viper.BindPFlag(env, f.Lookup(env))
	if err != nil {
		panic(err)
	}
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
