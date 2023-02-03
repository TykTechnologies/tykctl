package cloudcmd

import (
	"context"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewUpdateHomeDeployment(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(home).
		AddPreRunFuncs(NewCloudRbac(TeamAdmin, factory.Config).CloudRbac).
		WithFlagAdder(false, addHomeDeploymentFlag).
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, strings []string) error {

			return nil
		})

}

func validateHomeDeploymentAndUpdate(ctx context.Context, client internal.CloudClient, f *pflag.FlagSet, config internal.UserConfig) {
	deployment, err := extractCommonDeploymentFlags(f, config)
	if err != nil {
		return nil, err
	}

}
