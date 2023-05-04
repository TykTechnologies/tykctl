package cloudcmd

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/testutil"
)

func TestNewDeleteDeploymentFlags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	factory := internal.CloudFactory{
		Client: nil,
		Prompt: nil,
		Config: nil,
	}

	cmd := NewDeleteDeploymentCmd(factory)

	flags := []internal.Flag{
		{
			Description: "Test delete flag",
			Name:        "delete",
			Shorthand:   "d",
			Default:     "false",
			Value:       "false",
		}, {
			Description: "Test purge flag",
			Name:        "purge",
			Shorthand:   "p",
			Value:       "false",
			Default:     "false",
		}, {
			Description: "Test confirm flag",
			Name:        "confirm",
			Shorthand:   "c",
			Value:       "false",
			Default:     "false",
		},
	}

	testutil.TestFlags(t, cmd.Flags(), flags)
}
