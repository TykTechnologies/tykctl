package templates

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
)

func NewGenerateTemplatesCmd() *cobra.Command {
	return internal.NewCmd(shared.Generate).
		WithFlagAdder(false, outPutFileFlags).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			templateDIR, err := cmd.Flags().GetString(tempDIR)
			if err != nil {
				return err
			}

			return generateTemplates(templateDIR)
		})
}

func generateTemplates(dir string) error {
	err := util.CheckDirectory(dir)
	if err != nil {
		return err
	}

	err = generateApis(dir)
	if err != nil {
		return err
	}

	return generateKeys(dir)
}

func generateApis(dir string) error {
	err := SaveTemplateToFile("lean-keyless-api", dir, createLeanKeylessAPIDefinition())
	if err != nil {
		return err
	}

	fullProtectedAPIDef, err := createFullProtectedAPIDefinition()
	if err != nil {
		return err
	}

	return SaveTemplateToFile("full-protected-api", dir, fullProtectedAPIDef)
}

func generateKeys(dir string) error {
	return SaveTemplateToFile("lean-key", dir, createLeanKeyTemplate())
}

func outPutFileFlags(f *pflag.FlagSet) {
	f.StringP(tempDIR, "d", ".", "where to store the template")
}
