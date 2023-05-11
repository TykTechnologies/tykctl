package sharedCmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/TykTechnologies/tykctl/internal"
)

var checkoutLongDescription = `Set the current service to make sure you only see commands related to service you are using, eg. "tykctl checkout cloud" will set the current_service to cloud service`

func NewCheckoutCmd() *cobra.Command {
	return internal.NewCmd("checkout").
		WithDescription("Sets tykctl to present commands within the context of the chosen service.").
		WithLongDescription(checkoutLongDescription).
		WithExample("tykctl checkout cloud").
		WithValidArgs([]string{internal.Cloud, internal.Gateway}).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			v, err := internal.CreateCoreViper()
			if err != nil {
				return err
			}

			service := args[0]
			if !slices.Contains(internal.AllowedServices, service) {
				return fmt.Errorf("%s is not allowed as an arg.Only %s,%s and %s are allowed as args", service, internal.Cloud, internal.Gateway, internal.All)
			}

			v.Set(internal.CurrentService, service)

			err = v.WriteConfig()
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}

			message := fmt.Sprintf("you will now only see commands %s service", service)
			if service == internal.All {
				message = "you will now see commands for all services"
			}

			cmd.Println(message)
			return nil
		})
}
