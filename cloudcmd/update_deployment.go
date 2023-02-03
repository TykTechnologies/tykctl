package cloudcmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewUpdateDeployment(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(update).
		WithFlagAdder(true, updateDeploymentFlag).
		AddPreRunFuncs(NewCloudRbac(TeamAdmin, factory.Config).CloudRbac).
		WithBindFlagWithCurrentUserContext([]internal.BindFlag{{Name: env, Persistent: false}, {Name: team, Persistent: false}, {Name: org, Persistent: false}}).
		WithCommands(NewUpdateHomeDeployment(factory))

}

func updateDeploymentFlag(f *pflag.FlagSet) {
	f.StringP(name, "n", "", "name for the deployment you want to create.")
	f.String(domain, "", "custom domain for your deployment")
}
