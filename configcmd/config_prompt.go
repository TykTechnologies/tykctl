package configcmd

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"

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

func (i PickConfigPrompt) PickConfig(current string, availableConfigFiles []string, shouldInitialize bool) (string, error) {
	currentTrimmed := internal.ConfigFileDisplayName(current)

	firstSelection := "Re-initialize the current configuration [%s] with new settings"
	newFile := "Create a new configuration"
	existing := "Switch to and re-initialize existing configuration: %s"

	if !shouldInitialize {
		firstSelection = "Continue using the current configuration [%s]"
		newFile = "Create and switch to a new configuration file"
		existing = "Switch to existing configuration: %s"
	}

	currentSelection := fmt.Sprintf(firstSelection, currentTrimmed)

	selections := []string{currentSelection, newFile}

	indexAvailableConfigFiles := []string{currentTrimmed, "Create a new configuration"}

	for _, file := range availableConfigFiles {
		fileTrimmed := internal.ConfigFileDisplayName(file)
		if currentTrimmed == fileTrimmed {
			continue
		}

		indexAvailableConfigFiles = append(indexAvailableConfigFiles, fileTrimmed)
		otherSelection := fmt.Sprintf(existing, fileTrimmed)
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

	err = survey.AskOne(namePrompt, &newFileName, survey.WithValidator(survey.Required), survey.WithValidator(validateFileName))
	if err != nil {
		return "", err
	}

	if slices.Contains(indexAvailableConfigFiles, newFileName) {
		return "", fmt.Errorf("config file with the name %s already exist.Use a different name", newFileName)
	}

	return newFileName, err
}

func validateFileName(val interface{}) error {
	value := reflect.ValueOf(val)
	if isZero(value) {
		return errors.New("value is required")
	}

	valStr := fmt.Sprint(val)

	if !nameMatch(valStr) {
		return errors.New("invalid filename")
	}

	return nil
}

func nameMatch(file string) bool {
	re := regexp.MustCompile(`^[a-z-]+$`)
	return re.MatchString(file)
}

// isZero returns true if the passed value is the zero object.
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		return v.Len() == 0
	}

	// compare the types directly with more general coverage
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
