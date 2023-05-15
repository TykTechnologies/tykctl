package configcmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

const switchDesc = `
This command will change from one configuration file to another.

You can pass the name of the config file you want to use as the first arg ("tykctl config switch <config file name>").

Note: the config file you pass must exist otherwise you will get an error that the file does not exist.

You can also switch to a new or existing file interactively by failing to provide an arg (tykctl config switch). 

The file you switch to will be made the active config and will be used as the --config flag value
if you do not pass it when running the commands.

`

func newSwitchConfigCmd(prompt internal.ConfigPrompt, configEntry internal.ConfigEntry) *cobra.Command {
	return internal.NewCmd(shared.Switch).
		WithExample("tykctl config switch").
		WithExample("tykctl config switch <config file name>").
		WithLongDescription(switchDesc).
		WithDescription("change from one configuration file to another one").
		MaximumArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			config, err := configEntry.GetAllConfig(true)
			if err != nil {
				return err
			}

			currentActiveConfig, err := configEntry.GetCurrentActiveConfig()
			if err != nil {
				return err
			}

			var pickedConfig string
			if len(args) == 0 {
				pickedConfig, err = prompt.PickConfig(currentActiveConfig, config, false)
				if err != nil {
					return err
				}
			} else {
				if !slices.Contains(config, args[0]) {
					return fmt.Errorf("config file %s does not exist", args[0])
				}
				pickedConfig = args[0]
			}

			err = configEntry.CreateConfigFile(pickedConfig, true)
			if err != nil {
				return err
			}

			cmd.Printf("you are now using config file %s", pickedConfig)

			return nil
		})
}
