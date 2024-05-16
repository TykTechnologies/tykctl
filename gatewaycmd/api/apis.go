package api

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
)

const (
	apis = "apis"
)

func NewApisCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(apis).
		WithAliases([]string{"api"}).
		WithCommands(NewFetchApisCmd(apimClient),
			NewDeleteApisCmd(apimClient),
			NewCreateAPICmd(apimClient))
}
