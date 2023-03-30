package certs

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

func NewCertsCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Certs).WithCommands(
		NewFetchCertsCmd(apimClient),
		NewDeleteCerts(apimClient),
	)
}
