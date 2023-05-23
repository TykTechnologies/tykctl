package templates

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tykctl/internal"
)

const (
	template = "templates"
)

func NewTemplateCmd() *cobra.Command {
	return internal.NewCmd(template).
		WithCommands(
			NewCreateTemplate(),
			NewGenerateTemplatesCmd(),
		)
}
