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

// SaveMapToConfig takes a map and write it to configuration file.
// the map keys are used as the keys in the config file.
func SaveMapToConfig(data map[string]string) error {
	for key, value := range data {
		viper.Set(key, value)
	}
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("Couldn't write config: %s\n", err.Error())
	}
	return nil
}

// SaveMapToCloudUserContext takes a map and write it to configuration file in the cloud service context
// under a certain user context.
func SaveMapToCloudUserContext(userId string, data map[string]string) error {
	for key, value := range data {
		cloudKey := fmt.Sprintf("cloud.%s.%s", userId, key)
		viper.Set(cloudKey, value)
	}
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("Couldn't write config: %s\n", err.Error())
	}
	return nil
}

// CreateKeyFromPath concatenate the paths provided to create a hierarchical config.
func CreateKeyFromPath(paths ...string) string {
	key := ""
	for i, path := range paths {
		if i != 0 {
			key = key + "."
		}
		key = key + path
	}
	return key
}
