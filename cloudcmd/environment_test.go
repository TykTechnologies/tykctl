package cloudcmd

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/TykTechnologies/tykctl/internal"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/TykTechnologies/tykctl/testutil"
)

func TestNewEnvironmentCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	prompt := mock.NewMockCloudPrompt(ctrl)
	factory := internal.CloudFactory{
		Client: m,
		Prompt: prompt,
	}
	envCmd := NewEnvironmentCmd(factory)

	flags := []internal.Flag{
		{
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
