package cloudcmd

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/tykctl/internal"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/TykTechnologies/tykctl/testutil"
)

func TestNewTeamCmdPersistentFlags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	prompt := mock.NewMockCloudPrompt(ctrl)
	factory := internal.CloudFactory{
		Client: m,
		Prompt: prompt,
	}
	cmd := NewTeamCmd(factory)
	assert.Equal(t, true, cmd.PersistentFlags().HasFlags())
	flags := []internal.Flag{{
		Description: "Test Organization is added",
		Name:        "org",
		Shorthand:   "",
		Value:       "",
		Default:     "",
	}}
	testutil.TestFlags(t, cmd.PersistentFlags(), flags)
}
