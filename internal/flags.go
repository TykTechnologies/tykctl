package internal

import (
	"fmt"
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

// SaveToConfig writes data to the config file provided by --config
func SaveToConfig(key, value string) error {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("Couldn't write config: %s\n", err.Error())
	}
	return nil
}
