package cloudcmd

import (
	"github.com/spf13/pflag"
)

// addOrgFlag adds the org flag to the flags
// it uses the default org set in you config file if it is not passed.
func addOrgFlag(f *pflag.FlagSet) {
	f.String(org, "", "The organization to use")
}
func addTeamFlag(f *pflag.FlagSet) {
	f.String(team, "", "The team to use")
}
func addEnvFlag(f *pflag.FlagSet) {
	f.String(env, "", "The environment to use")
}

func addOutPutFlags(f *pflag.FlagSet) {
	f.StringP(outPut, "o", "table", "Format you want to use can be table,json")
}
