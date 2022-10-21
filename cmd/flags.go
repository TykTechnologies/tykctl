package cmd

import (
	"errors"
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

func SaveToConfig(key, token string) error {
	viper.Set(key, token)
	err := viper.WriteConfig()
	if err != nil {
		message := fmt.Sprintf("Couldn't write config: %s\n", err.Error())
		return errors.New(message)
	}
	return err

}
