package configcmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
)

func NewConfigCmd(prompt internal.ConfigPrompt, configEntry internal.ConfigEntry, factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(internal.Config).
		WithCommands(newInitConfigCmd(prompt, configEntry, factory))
}
