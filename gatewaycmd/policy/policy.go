package policy

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

func NewPolicyCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Policy).
		WithCommands(
			NewCreatePolicyCmd(apimClient),
		)
}
