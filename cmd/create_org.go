package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
)

func NewOrgCreateCmd() *cobra.Command {
	return NewCmd("create").
		NoArgs(func(ctx context.Context, command cobra.Command) error {
			log.Println("running this command")
			createOrg(command.Context())
			return nil
		})

}

func createOrg(ctx context.Context) {

}
