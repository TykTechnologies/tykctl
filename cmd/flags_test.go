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
