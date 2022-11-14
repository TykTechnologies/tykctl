package cloudcmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/testutil"
	"testing"
)

func TestAddOutPutFlags(t *testing.T) {
	cmd := internal.NewCmd("test").WithFlagAdder(false, addOutPutFlags).NoArgs(nil)
	flags := []internal.Flag{{
		Description: "Test output format is added",
		Name:        "output",
		Shorthand:   "o",
		Default:     "table",
		Value:       "table",
	}}
	testutil.TestFlags(t, cmd.Flags(), flags)
}

func TestAddOrgFlag(t *testing.T) {
	cmd := internal.NewCmd("test").WithFlagAdder(true, addOrgFlag).NoArgs(nil)
	flags := []internal.Flag{{
		Description: "Test org is added to persistent flags",
		Name:        "org",
		Shorthand:   "",
		Value:       "",
		Default:     "",
	}}
	testutil.TestFlags(t, cmd.PersistentFlags(), flags)
}

func TestAddTeamFlag(t *testing.T) {
	cmd := internal.NewCmd("test").WithFlagAdder(true, addTeamFlag).NoArgs(nil)
	flags := []internal.Flag{{
		Description: "Test team is added to flags",
		Name:        "team",
		Shorthand:   "",
		Value:       "",
		Default:     "",
	}}
	testutil.TestFlags(t, cmd.PersistentFlags(), flags)
}
