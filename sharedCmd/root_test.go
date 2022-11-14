package sharedCmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/testutil"
	"testing"
)

func TestNewRootCmd(t *testing.T) {
	flags := []internal.Flag{{
		Description: "Test toggle flag",
		Name:        "toggle",
		Shorthand:   "t",
		Default:     "false",
		Value:       "false",
	}}
	cmd := NewRootCmd()
	testutil.TestFlags(t, cmd.Flags(), flags)
}
