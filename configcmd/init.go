package configcmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/TykTechnologies/tykctl/internal"
)

const configInitDesc = `
 This command will initialize an existing or a new configuration file.

The configuration file you initialize with this command will be made the active config and will be used as the --config flag value
if you do not pass it when running the commands.

The command will also allow you to select the service that you want to activate.

If you decide to activate all services or the cloud service you will be prompted you login to the cloud and the token will be saved in the file you have initialised.

`

func newInitConfigCmd(prompt internal.ConfigPrompt, configEntry internal.ConfigEntry, factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(internal.Init).
		WithExample("tykctl config init").
		WithLongDescription(configInitDesc).
		WithDescription("initialize a new or existing configuration file").
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			configFiles, err := configEntry.GetAllConfig(false)
			if err != nil {
				return err
			}

			currentActiveConfig, err := configEntry.GetCurrentActiveConfig()
			if err != nil {
				return err
			}

			pickedConfig, err := prompt.PickConfig(currentActiveConfig, configFiles, true)
			if err != nil {
				return err
			}

			err = configEntry.CreateConfigFile(pickedConfig, true)
			if err != nil {
				return err
			}

			viper.SetConfigName(fmt.Sprintf("config_%s", pickedConfig))

			service, err := prompt.PickServiceToUse(true)
			if err != nil {
				return err
			}

			err = AddGatewayURL(prompt, service)
			if err != nil {
				return err
			}

			return loginToCloud(ctx, prompt, factory, service)
		})
}

func AddGatewayURL(prompt internal.ConfigPrompt, service string) error {
	if service != internal.GatewayService && service != internal.All {
		return nil
	}

	shouldSetServer, err := prompt.AskGatewayURL()
	if err != nil || !shouldSetServer {
		return err
	}

	url, secret, err := prompt.SetGatewayURL()
	if err != nil {
		return err
	}

	err = internal.SaveValueToConfig("gateway.urls", []string{url})
	if err != nil {
		return err
	}

	return internal.SaveValueToConfig("gateway.secret", secret)
}

func loginToCloud(ctx context.Context, prompt internal.ConfigPrompt, factory internal.CloudFactory, service string) error {
	if service != internal.Cloud && service != internal.All {
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
