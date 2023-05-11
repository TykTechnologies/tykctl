package configcmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
)

func newInitConfigCmd(prompt internal.ConfigPrompt, configEntry internal.ConfigEntry, factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(internal.Init).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			configFiles, err := configEntry.GetAllConfig()
			if err != nil {
				return err
			}

			currentActiveConfig, err := configEntry.GetCurrentActiveConfig()
			if err != nil {
				return err
			}

			pickedConfig, err := prompt.PickConfig(currentActiveConfig, configFiles)
			if err != nil {
				return err
			}

			err = configEntry.CreateConfigFile(pickedConfig, true)
			if err != nil {
				return err
			}

			service, err := prompt.PickServiceToUse(true)
			if err != nil {
				return err
			}

			return loginToCloud(ctx, prompt, factory, service)
		})
}

func loginToCloud(ctx context.Context, prompt internal.ConfigPrompt, factory internal.CloudFactory, service string) error {
	if service == internal.Cloud || service == internal.All {
		return nil
	}

	shouldLoginCloud, err := prompt.AskCloudLogin()
	if err != nil || !shouldLoginCloud {
		return err
	}

	err = prompt.LoginCloud(ctx)
	if err != nil {
		return err
	}

	return prompt.InitUserConfigFile(ctx, factory)
}
