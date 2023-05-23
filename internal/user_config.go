package internal

import (
	"github.com/spf13/viper"
)

// UserConfig this will help us make viper more testable.
//
//go:generate mockgen -source=user_config.go -destination=./mocks/user_config.go -package=mock UserConfig
type UserConfig interface {
	GetCurrentUserOrg() string
	GetCurrentUserTeam() string
	GetCurrentUserEnv() string
	GetCurrentUserRole() string
}

var _ UserConfig = (*ViperConfig)(nil)

type ViperConfig struct{}

func (v ViperConfig) GetCurrentUserOrg() string {
	return viper.GetString(CreateKeyFromPath(cloudPath, org))
}

func (v ViperConfig) GetCurrentUserTeam() string {
	return viper.GetString(CreateKeyFromPath(cloudPath, team))
}

func (v ViperConfig) GetCurrentUserEnv() string {
	return viper.GetString(CreateKeyFromPath(cloudPath, env))
}

func (v ViperConfig) GetCurrentUserRole() string {
	return viper.GetString(CreateKeyFromPath(cloudPath, userRole))
}
