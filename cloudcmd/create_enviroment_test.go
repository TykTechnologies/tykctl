package cloudcmd

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/TykTechnologies/tykctl/testutil"
)

func TestFlagsAreAddedToCreateEnvironment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	prompt := mock.NewMockCloudPrompt(ctrl)
	factory := internal.CloudFactory{
		Client: m,
		Prompt: prompt,
	}
	parentCmd := NewEnvironmentCmd(factory)
	cmd := NewCreateEnvironmentCmd(factory)
	parentCmd.AddCommand(cmd)

	localFlags := []internal.Flag{{
		Description: "Test team name is added.",
		Name:        "name",
		Shorthand:   "n",
		Value:       "",
		Default:     "",
	}}
	testutil.TestFlags(t, cmd.Flags(), localFlags)

	inheritedFlags := []internal.Flag{
		{
			Description: "Test team flag is added",
			Name:        "team",
			Shorthand:   "",
			Value:       "",
			Default:     "",
		},
		{
			Description: "Test org is passed from parent",
			Name:        "org",
			Shorthand:   "",
			Value:       "",
			Default:     "",
		},
	}
	testutil.TestFlags(t, cmd.InheritedFlags(), inheritedFlags)
}

func TestCreateEnvironment(t *testing.T) {
	tests := []struct {
		name                string
		mockError           error
		mockHTTPResponse    *http.Response
		mockResponse        cloud.InlineResponse2012
		expectedEnvResponse *cloud.Loadout
		ExpectedError       error
	}{
		{
			name:             "Check status ok",
			mockError:        nil,
			mockHTTPResponse: &http.Response{StatusCode: http.StatusCreated},
			mockResponse: cloud.InlineResponse2012{
				Error_:  "",
				Payload: &generateEnvs(1)[0],
				Status:  "ok",
			},
			expectedEnvResponse: &cloud.Loadout{OID: "1", UID: "1"},
			ExpectedError:       nil,
		},
		{
			name:             "Test status is not 201",
			mockError:        nil,
			mockHTTPResponse: &http.Response{StatusCode: http.StatusBadGateway},
			mockResponse: cloud.InlineResponse2012{
				Error_:  "",
				Payload: &generateEnvs(1)[0],
				Status:  "ok",
			},
			expectedEnvResponse: nil,
			ExpectedError:       ErrorCreatingEnv,
		},
		{
			name:             "Test payload status is not ok",
			mockError:        nil,
			mockHTTPResponse: &http.Response{StatusCode: http.StatusCreated},
			mockResponse: cloud.InlineResponse2012{
				Error_:  "error found here 1",
				Payload: &generateEnvs(3)[0],
				Status:  "error",
			},
			expectedEnvResponse: nil,
			ExpectedError:       errors.New("error found here 1"),
		},
		{
			name:                "Test error returned by cloud",
			mockError:           ErrorGenericError,
			mockHTTPResponse:    &http.Response{StatusCode: http.StatusCreated},
			mockResponse:        cloud.InlineResponse2012{},
			expectedEnvResponse: nil,
			ExpectedError:       ErrorGenericError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().CreateEnv(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(tt.mockResponse, tt.mockHTTPResponse, tt.mockError)
			env, err := CreateEnvironment(context.Background(), m, cloud.Loadout{}, "", "")
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.expectedEnvResponse, env)
		})
	}
}

func TestValidateFlagsAndCreateEnv(t *testing.T) {
	testsCases := []struct {
		name                 string
		ExpectedError        error
		ExpectedEnv          *cloud.Loadout
		ExpectedMockFuncCall int
		mockError            error
		mockHTTPResponse     *http.Response
		mockResponse         cloud.InlineResponse2012
		envName              string
		orgID                string
		teamID               string
	}{
		{
			name:                 "Test success",
			ExpectedError:        nil,
			ExpectedEnv:          &generateEnvs(3)[0],
			ExpectedMockFuncCall: 1,
			mockError:            nil,
			mockHTTPResponse:     &http.Response{StatusCode: http.StatusCreated},
			mockResponse: cloud.InlineResponse2012{
				Error_:  "",
				Payload: &generateEnvs(3)[0],
				Status:  "ok",
			},
			envName: "env name",
			orgID:   "orgID here",
			teamID:  "teamid here",
		},
		{
			name:                 "Error organization id is required",
			ExpectedError:        ErrorOrgRequired,
			ExpectedEnv:          nil,
			ExpectedMockFuncCall: 0,
			mockError:            nil,
			mockHTTPResponse:     nil,
			mockResponse:         cloud.InlineResponse2012{},
			envName:              "",
			orgID:                "",
			teamID:               "",
		},
		{
			name:                 "Test Team is empty",
			ExpectedError:        ErrorTeamRequired,
			ExpectedEnv:          nil,
			ExpectedMockFuncCall: 0,
			mockError:            nil,
			mockHTTPResponse:     nil,
			mockResponse:         cloud.InlineResponse2012{},
			envName:              "env here",
			orgID:                "org here",
			teamID:               "",
		},
		{
			name:                 "Test env name is empty",
			ExpectedError:        ErrorNameRequired,
			ExpectedEnv:          nil,
			ExpectedMockFuncCall: 0,
			mockError:            nil,
			mockHTTPResponse:     nil,
			mockResponse:         cloud.InlineResponse2012{},
			envName:              "",
			orgID:                "ijskhsnsn",
			teamID:               "nbxdouert",
		},
	}
	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().CreateEnv(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(tt.ExpectedMockFuncCall).Return(tt.mockResponse, tt.mockHTTPResponse, tt.mockError)
			env, err := validateFlagsAndCreateEnv(context.Background(), m, tt.envName, tt.teamID, tt.orgID, nil)
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.ExpectedEnv, env)
		})
	}
}

func generateEnvs(size int) []cloud.Loadout {
	var loadouts []cloud.Loadout
	for i := 0; i < size; i++ {
		loadouts = append(loadouts, cloud.Loadout{
			OID: strconv.Itoa(i + 1),
			UID: strconv.Itoa(i + 1),
		})
	}

	return loadouts
}
