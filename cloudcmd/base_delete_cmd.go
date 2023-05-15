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
		WithBindFlagOnPreRun([]internal.BindFlag{{Name: org, Persistent: false, Type: internal.Cloud}})

	if objectType == Env {
		builder.WithBindFlagOnPreRun([]internal.BindFlag{{Name: team, Persistent: false, Type: internal.Cloud}})
		builder.WithFlagAdder(false, cascadeFlag)
		builder.WithFlagAdder(false, envFlags)
	}

	if objectType == Dep {
		builder.WithBindFlagOnPreRun([]internal.BindFlag{{Name: env, Persistent: false, Type: internal.Cloud}, {Name: team, Persistent: false, Type: internal.Cloud}})
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

		cmd.Printf("deleted %s successfully", deleteInterface.GetUID())

		return nil
	})

	return deleteCmd
}

func shouldDelete(prompt internal.CloudPrompt, object string, f *pflag.FlagSet) (bool, error) {
	confirmed, err := f.GetBool(confirm)
	if err != nil {
		return false, err
	}

	if confirmed {
		return true, nil
	}

	return prompt.PerformActionPrompt(object)
}

type CloudObjectType int64

const (
	Org CloudObjectType = iota
	Team
	Env
	Dep
)
