package internal

import "context"

//go:generate mockgen -source=config_prompt.go -destination=./mocks/config_prompt.go -package=mock ConfigPrompt
type ConfigPrompt interface {
	PickConfig(current string, availableConfigFiles []string) (string, error)
	PickServiceToUse(shouldSave bool) (string, error)
	AskCloudLogin() (bool, error)
	LoginCloud(ctx context.Context) error
	InitUserConfigFile(ctx context.Context, factory CloudFactory) error
}
