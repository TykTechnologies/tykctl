package cloudcmd

import (
	"context"
	"fmt"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

var (
	ctxLongDescription = `Set config flags using arguments, eg.
		"tykctl ctx org bar"
	will set the variable "org" to "bar"`
)

func NewCtxCmd() *cobra.Command {
	return internal.NewCmd(ctxCmd).
		WithDescription("set a value in the provided config file.").
		WithLongDescription(ctxLongDescription).
		WithExample("tykctl ctx org bar").
		ExactArgs(2, func(ctx context.Context, cmd cobra.Command, args []string) error {
			err := internal.SaveToConfig(args[0], args[1])
			if err != nil {
				cmd.Println(err)
				return err
			}
			cmd.Println(fmt.Sprintf("%s set successfully", args[0]))
			return nil
		})
}
