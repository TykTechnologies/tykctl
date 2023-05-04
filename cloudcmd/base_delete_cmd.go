package cloudcmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/tykctl/internal"
)

type DeleteInterface interface {
	GetUID() string
	Delete(ctx context.Context, client internal.CloudClient, config internal.UserConfig, id string, f *pflag.FlagSet) error
}

func NewDeleteBaseCmd(factory internal.CloudFactory, deleteInterface DeleteInterface, objectType CloudObjectType) *cobra.Command {
	builder := internal.NewCmd(delete).
		WithAliases([]string{del}).
		WithFlagAdder(false, confirmFlag).
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: org, Persistent: false}})

	if objectType == Env {
		builder.WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: team, Persistent: false}})
		builder.WithFlagAdder(false, cascadeFlag)
	}

	if objectType == Dep {
		builder.WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}})
		builder.WithFlagAdder(false, deleteDeploymentFlag)
	}

	deleteCmd := builder.ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
		delEnv, err := shouldDelete(factory.Prompt, "Environment", cmd.Flags())
		if err != nil {
			return err
		}

		if !delEnv {
			return nil
		}

		err = deleteInterface.Delete(cmd.Context(), factory.Client, factory.Config, args[0], cmd.Flags())
		if err != nil {
			return err
		}

		cmd.Printf("deleted %s\n successfully", deleteInterface.GetUID())

		return nil
	})

	return deleteCmd
}
