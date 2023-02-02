package cloudcmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/viper"
)

func getCurrentUserOrg() string {
	return viper.GetString(internal.CreateKeyFromPath(cloudPath, viper.GetString(currentCloudUser), org))
}

func getCurrentUserTeam() string {
	return viper.GetString(internal.CreateKeyFromPath(cloudPath, viper.GetString(currentCloudUser), team))
}

func getCurrentUserEnv() string {
	return viper.GetString(internal.CreateKeyFromPath(cloudPath, viper.GetString(currentCloudUser), env))
}

func getCurrentUserRole() string {
	return viper.GetString(internal.CreateKeyFromPath(cloudPath, viper.GetString(currentCloudUser), userRole))
}
