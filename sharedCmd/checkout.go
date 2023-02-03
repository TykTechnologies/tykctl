package sharedCmd

import (
	"context"
	"fmt"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

var (
	checkoutLongDescription = `Set the current service to make sure you only see commands related to service you are using, eg.
		
     "tykctl checkout cloud"
	
     will set the current_service to cloud service`
)

const (
	CurrentService = "current_service"
	Cloud          = "cloud"
	Gateway        = "gateway"
)

func NewCheckoutCmd() *cobra.Command {
	return internal.NewCmd("checkout").
		WithDescription(checkoutLongDescription).
		WithLongDescription(checkoutLongDescription).
		WithExample("tykctl checkout cloud").
		WithValidArgs([]string{Cloud, Gateway}).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			err := internal.SaveToConfig(CurrentService, args[0])
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			cmd.Println(fmt.Sprintf("you will now only see commands %s service", args[0]))
			return nil
		})
}
