package cmd

import (
	"testing"
)

func TestAddOutPutFlags(t *testing.T) {
	cmd := NewCmd("test").WithFlagAdder(false, addOutPutFlags).NoArgs(nil)
	flags := []Flag{{
		Description: "Test output format is added",
		Name:        "output",
		Shorthand:   "o",
		Default:     "table",
		Value:       "table",
	}}
	testFlags(t, cmd.Flags(), flags)
}

func TestAddOrgFlag(t *testing.T) {
	cmd := NewCmd("test").WithFlagAdder(true, addOrgFlag).NoArgs(nil)
	flags := []Flag{{
		Description: "Test org is added to persistent flags",
		Name:        "org",
		Shorthand:   "",
		Value:       "",
		Default:     "",
	}}
	testFlags(t, cmd.PersistentFlags(), flags)
}

func TestAddTeamFlag(t *testing.T) {
	cmd := NewCmd("test").WithFlagAdder(true, addTeamFlag).NoArgs(nil)
	flags := []Flag{{
		Description: "Test team is added to flags",
		Name:        "team",
		Shorthand:   "",
		Value:       "",
		Default:     "",
	}}
	testFlags(t, cmd.PersistentFlags(), flags)
}
