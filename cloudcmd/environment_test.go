package cloudcmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/TykTechnologies/tykctl/testutil"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestNewEnvironmentCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	envCmd := NewEnvironmentCmd(m)

	flags := []internal.Flag{{
		Description: "Test org is added",
		Name:        "org",
		Shorthand:   "",
		Default:     "",
		Value:       "",
	},
		{
			Description: "Test Team flag",
			Name:        "team",
			Shorthand:   "",
			Value:       "",
			Default:     "",
		},
	}
	testutil.TestFlags(t, envCmd.PersistentFlags(), flags)

}
