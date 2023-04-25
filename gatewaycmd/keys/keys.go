package keys

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

func NewKeysCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Keys).
		WithCommands(
			NewFetchKeysCmd(apimClient),
		)
}
