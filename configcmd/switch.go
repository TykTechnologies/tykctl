package configcmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

func newSwitchConfigCmd(prompt internal.ConfigPrompt, configEntry internal.ConfigEntry) *cobra.Command {
	return internal.NewCmd(shared.Switch).MaximumArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
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
