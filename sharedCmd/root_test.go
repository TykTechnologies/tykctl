package sharedCmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/testutil"
	"testing"
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
