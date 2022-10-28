package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

func NewOrgCreateCmd() *cobra.Command {
	return NewCmd("create").
		NoArgs(func(ctx context.Context, command cobra.Command) error {
			return nil
		})

}
