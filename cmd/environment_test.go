package cmd

import (
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestNewEnvironmentCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	envCmd := NewEnvironmentCmd(m)

	flags := []Flag{{
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
	testFlags(t, envCmd.PersistentFlags(), flags)

}
