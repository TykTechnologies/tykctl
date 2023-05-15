package configcmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

func newConfigFetchCmd(configEntry internal.ConfigEntry) *cobra.Command {
	return internal.NewCmd(shared.Fetch).NoArgs(func(ctx context.Context, cmd cobra.Command) error {
		config, err := configEntry.GetAllConfig(true)
		if err != nil {
			return err
		}

		activeConfig, err := configEntry.GetCurrentActiveConfig()
		if err != nil {
			return err
		}

		internal.PrintList("", config, activeConfigIndex(config, activeConfig))
		return nil
	})
}

func activeConfigIndex(files []string, active string) []int {
	var highlight []int

	for index, file := range files {
		if file == internal.ConfigFileDisplayName(active) {
			highlight = append(highlight, index)
		}
	}

	return highlight
}
