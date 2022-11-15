package cloudcmd

import (
	"github.com/TykTechnologies/tykctl/internal"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/TykTechnologies/tykctl/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTeamCmdPersistentFlags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	cmd := NewTeamCmd(m)
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