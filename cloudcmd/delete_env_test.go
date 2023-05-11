package cloudcmd

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
)

func TestNewDeleteEnvCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	prompt := mock.NewMockCloudPrompt(ctrl)
	config := mock.NewMockUserConfig(ctrl)
	factory := internal.CloudFactory{
		Client: m,
		Prompt: prompt,
		Config: config,
	}

	prompt.EXPECT().PerformActionPrompt(gomock.Any()).Return(true, nil)
	config.EXPECT().GetCurrentUserOrg().Return("my-org")
	config.EXPECT().GetCurrentUserTeam().Return("my-team")
	m.EXPECT().DeleteEnv(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(cloud.InlineResponse2012{
		Error_:  "",
		Payload: &cloud.Loadout{},
		Status:  "",
	}, &http.Response{StatusCode: 200}, nil)

	cmd := NewDeleteEnvCmd(factory)
	cmd.SetArgs([]string{
		"1",
		fmt.Sprintf("--org=%s", "mmmmf"),
	})

	err := cmd.Execute()
	assert.Nil(t, err)
}
