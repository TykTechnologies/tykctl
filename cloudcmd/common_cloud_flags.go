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

func getValues(f *pflag.FlagSet) {
	f.StringSlice(get, []string{}, "Get a value from the object using dot-notation")
}

func setValues(f *pflag.FlagSet) {
	f.StringSlice(set, []string{}, "set a value for the object using dot-notation")
}

func envValues(f *pflag.FlagSet) {
	f.StringSlice(envValue, []string{}, "change a deployment's environment variables using dot-notation")
}

func getEnvValues(f *pflag.FlagSet) {
	f.StringSlice(envValue, []string{}, "get a deployment's environment variables using dot-notation")
}

func confirmFlag(f *pflag.FlagSet) {
	f.BoolP(confirm, "c", false, "delete the deployment without a confirmation prompt")
}

func cascadeFlag(f *pflag.FlagSet) {
	f.Bool(cascade, false, "delete the deployment without a confirmation prompt")
}
