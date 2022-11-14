package cmd

import (
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
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
	flags := []Flag{{
		Description: "Test Organization is added",
		Name:        "org",
		Shorthand:   "",
		Value:       "",
		Default:     "",
	}}
	testFlags(t, cmd.PersistentFlags(), flags)
}
