package configcmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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

			viper.SetConfigName(fmt.Sprintf("config_%s", pickedConfig))
			use, err := prompt.PickServiceToUse(true)
			if err != nil {
				return err
			}

			if use == internal.Cloud || use == internal.All {
				shouldLoginCloud, err := prompt.AskCloudLogin()
				if err != nil {
					return err
				}

				if shouldLoginCloud {
					err = prompt.LoginCloud(ctx)
					if err != nil {
						return err
					}

					err = prompt.InitUserConfigFile(cmd.Context(), factory)
					if err != nil {
						return err
					}
				}
			}

			err = viper.WriteConfig()
			if err != nil {
				return err
			}

			cmd.Println(viper.ConfigFileUsed())
			return nil
		})
}
