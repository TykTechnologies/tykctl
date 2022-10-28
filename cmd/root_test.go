package cmd

import (
	"testing"
)

func TestNewRootCmd(t *testing.T) {
	flags := []Flag{{
		Description: "Test toggle flag",
		Name:        "toggle",
		Shorthand:   "t",
		Default:     "false",
		Value:       "false",
	}}

	cmd := NewRootCmd()
	testFlags(t, cmd.Flags(), flags)

}
