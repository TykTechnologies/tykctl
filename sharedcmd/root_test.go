package sharedcmd

import (
	"testing"

	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/testutil"
)

func TestNewRootCmd(t *testing.T) {
	flags := []internal.Flag{{
		Description: "Test config flag",
		Name:        "config",
		Shorthand:   "",
		Default:     "",
		Value:       "",
	}}
	cmd := NewRootCmd()
	testutil.TestFlags(t, cmd.PersistentFlags(), flags)
}
