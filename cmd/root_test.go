package cmd

import (
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestNewRootCmd(t *testing.T) {
	flags := []Flag{{
		Description: "Test toggle flag",
		Name:        "toggle",
		Shorthand:   "t",
		Default:     "false",
		Value:       "false",
	}}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	cmd := NewRootCmd(m)
	testFlags(t, cmd.Flags(), flags)

}
