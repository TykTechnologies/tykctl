package configcmd

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"golang.org/x/exp/slices"

	"github.com/TykTechnologies/tykctl/cloudcmd"
	"github.com/TykTechnologies/tykctl/internal"
)

var _ internal.ConfigPrompt = (*PickConfigPrompt)(nil)

type PickConfigPrompt struct{}

func (i PickConfigPrompt) InitUserConfigFile(ctx context.Context, factory internal.CloudFactory) error {
	return cloudcmd.InitUserConfigFile(ctx, factory)
}

func (i PickConfigPrompt) AskCloudLogin() (bool, error) {
	value := false
	prompt := &survey.Confirm{
		Message: "Would you like to login to tyk cloud?",
	}

	err := survey.AskOne(prompt, &value)

	return value, err
}

func (i PickConfigPrompt) LoginCloud(ctx context.Context) error {
	loginBody, err := cloudcmd.LoginInteractive()
	if err != nil {
		return err
	}

	return cloudcmd.ValidateAndLogin(ctx, loginBody)
}

func (i PickConfigPrompt) PickServiceToUse(shouldSave bool) (string, error) {
	prompt := &survey.Select{
		Message: "Which service do you want to enable:",
		Options: internal.AllowedServices,
	}

	var selected string

	err := survey.AskOne(prompt, &selected, survey.WithValidator(survey.Required))
	if err != nil {
		return "", err
	}

	if !shouldSave {
		return selected, nil
	}

	v, err := internal.CreateCoreViper()
	if err != nil {
		return "", err
	}

	v.Set(internal.CurrentService, selected)

	err = v.WriteConfig()
	if err != nil {
		return "", err
	}

	return selected, err
}

func (i PickConfigPrompt) PickConfig(current string, availableConfigFiles []string) (string, error) {
	currentTrimmed := internal.ConfigFileDisplayName(current)
	currentSelection := fmt.Sprintf("Re-initialize the current configuration [%s] with new settings", currentTrimmed)
	selections := []string{currentSelection, "Create a new configuration"}

	indexAvailableConfigFiles := []string{currentTrimmed, "Create a new configuration"}

	for _, file := range availableConfigFiles {
		fileTrimmed := internal.ConfigFileDisplayName(file)
		if currentTrimmed == fileTrimmed {
			continue
		}

		indexAvailableConfigFiles = append(indexAvailableConfigFiles, fileTrimmed)
		otherSelection := fmt.Sprintf("Switch to and re-initialize existing configuration: %s", fileTrimmed)
		selections = append(selections, otherSelection)
	}

	prompt := &survey.Select{
		Message: "Pick configuration to use:",
		Options: selections,
	}

	var selectedIndex int

	err := survey.AskOne(prompt, &selectedIndex, survey.WithValidator(survey.Required))
	if err != nil {
		return "", err
	}

	if selectedIndex != 1 {
		return indexAvailableConfigFiles[selectedIndex], nil
	}

	newFileName := ""
	namePrompt := &survey.Input{
		Message: "Enter configuration name. Names start with a lower case letter and contain only \nlower case letters a-z, digits 0-9, and hyphens '-':",
	}

	err = survey.AskOne(namePrompt, &newFileName, survey.WithValidator(survey.Required))
	if err != nil {
		return "", err
	}

	if slices.Contains(indexAvailableConfigFiles, newFileName) {
		return "", fmt.Errorf("config file with the name %s already exist.Use a different name", newFileName)
	}

	return newFileName, err
}
