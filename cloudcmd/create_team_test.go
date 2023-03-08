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

func TestCreateTeam(t *testing.T) {
	tests := []struct {
		name                 string
		ExpectedError        error
		mockHTTPResponse     *http.Response
		mockError            error
		mockResponse         cloud.InlineResponse2011
		expectedTeamResponse *cloud.Team
	}{
		{
			name:                 "Check status ok",
			ExpectedError:        nil,
			mockHTTPResponse:     &http.Response{StatusCode: http.StatusCreated},
			mockError:            nil,
			expectedTeamResponse: &cloud.Team{OID: "1", UID: "1"},
			mockResponse: cloud.InlineResponse2011{
				Error_:  "",
				Payload: &generateTeams(1)[0],
				Status:  "ok",
			},
		},
		{
			name:             "Test Error returned by cloud api",
			ExpectedError:    ErrorGenericError,
			mockHTTPResponse: &http.Response{StatusCode: http.StatusCreated},
			mockError:        ErrorGenericError,
			mockResponse: cloud.InlineResponse2011{
				Error_:  "error here",
				Payload: nil,
				Status:  "error",
			},
			expectedTeamResponse: nil,
		},
		{
			name:             "Test payload status is error",
			ExpectedError:    errors.New("check error"),
			mockHTTPResponse: &http.Response{StatusCode: http.StatusCreated},
			mockError:        nil,
			mockResponse: cloud.InlineResponse2011{
				Error_:  "check error",
				Payload: &generateTeams(1)[0],
				Status:  "error",
			},
			expectedTeamResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().CreateTeam(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(tt.mockResponse, tt.mockHTTPResponse, tt.mockError)
			team, err := CreateTeam(context.Background(), m, cloud.Team{}, "")
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.expectedTeamResponse, team)
		})
	}
}

func TestValidateFlagsAndCreateTeam(t *testing.T) {
	tests := []struct {
		name                       string
		ExpectedError              error
		mockError                  error
		mockHTTPResponse           *http.Response
		mockResponse               cloud.InlineResponse2011
		numberOfTimeToCallMockFunc int
		teamName                   string
		org                        string
		expectedTeamResponse       *cloud.Team
	}{
		{
			name:             "Test Success",
			ExpectedError:    nil,
			mockError:        nil,
			mockHTTPResponse: &http.Response{StatusCode: http.StatusCreated},
			mockResponse: cloud.InlineResponse2011{
				Error_:  "",
				Payload: &generateTeams(1)[0],
				Status:  "ok",
			},
			numberOfTimeToCallMockFunc: 1,
			teamName:                   "team name",
			org:                        "u4ieu47ueu",
			expectedTeamResponse:       &cloud.Team{OID: "1", UID: "1"},
		},
		{
			name:                       "Test org is required",
			ExpectedError:              ErrorOrgRequired,
			mockError:                  nil,
			mockHTTPResponse:           nil,
			mockResponse:               cloud.InlineResponse2011{},
			numberOfTimeToCallMockFunc: 0,
			teamName:                   "",
			org:                        "",
			expectedTeamResponse:       nil,
		},
		{
			name:                       "Test team name is required",
			ExpectedError:              ErrorNameRequired,
			mockError:                  nil,
			mockHTTPResponse:           nil,
			mockResponse:               cloud.InlineResponse2011{},
			numberOfTimeToCallMockFunc: 0,
			teamName:                   "",
			org:                        "4i57564j",
			expectedTeamResponse:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().CreateTeam(gomock.Any(), gomock.Any(), gomock.Any()).Times(tt.numberOfTimeToCallMockFunc).Return(tt.mockResponse, tt.mockHTTPResponse, tt.mockError)
			team, err := validateFlagsAndCreateTeam(context.Background(), m, tt.teamName, tt.org)
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.expectedTeamResponse, team)
		})
	}
}

func TestCreateTeamFlags(t *testing.T) {
	cmd := internal.NewCmd("test").WithFlagAdder(false, createTeamFlags).NoArgs(nil)
	flags := []internal.Flag{{
		Description: "Test name is added as flag",
		Name:        "name",
		Shorthand:   "n",
		Default:     "",
		Value:       "",
	}}

	testutil.TestFlags(t, cmd.Flags(), flags)
}

func TestAllFlagsAreAddedToCreateTeamCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	prompt := mock.NewMockCloudPrompt(ctrl)
	factory := internal.CloudFactory{
		Client: m,
		Prompt: prompt,
	}
	parentCmd := NewTeamCmd(factory)
	cmd := NewCreateTeamCmd(factory)
	parentCmd.AddCommand(cmd)

	localFlags := []internal.Flag{{
		Description: "Test name is added.",
		Name:        "name",
		Shorthand:   "n",
		Value:       "",
		Default:     "",
	}}
	testutil.TestFlags(t, cmd.Flags(), localFlags)

	inheritedFlags := []internal.Flag{{
		Description: "Test org is passed from parent",
		Name:        "org",
		Shorthand:   "",
		Value:       "",
		Default:     "",
	}}
	testutil.TestFlags(t, cmd.InheritedFlags(), inheritedFlags)
}

func generateTeams(size int) []cloud.Team {
	var teams []cloud.Team
	for i := 0; i < size; i++ {
		teams = append(teams, cloud.Team{
			OID:          strconv.Itoa(i + 1),
			Organisation: nil,
			UID:          strconv.Itoa(i + 1),
		})
	}

	return teams
}
