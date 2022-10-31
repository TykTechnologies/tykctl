package cmd

import (
	"context"
	"errors"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestCreateTeam(t *testing.T) {
	tests := []struct {
		name                 string
		ExpectedError        error
		mockHttpResponse     *http.Response
		mockError            error
		mockResponse         cloud.InlineResponse2011
		expectedTeamResponse *cloud.Team
	}{
		{
			name:                 "Check status Ok",
			ExpectedError:        nil,
			mockHttpResponse:     &http.Response{StatusCode: http.StatusCreated},
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
			mockHttpResponse: &http.Response{StatusCode: http.StatusCreated},
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
			mockHttpResponse: &http.Response{StatusCode: http.StatusCreated},
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
			m.EXPECT().CreateTeam(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			team, err := CreateTeam(context.Background(), m, cloud.Team{}, "")
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.expectedTeamResponse, team)
		})
	}
}
func TestCreateTeamFlags(t *testing.T) {
	cmd := NewCmd("test").WithFlagAdder(false, createTeamFlags).NoArgs(nil)
	flags := []Flag{{
		Description: "Test name is added as flag",
		Name:        "name",
		Shorthand:   "n",
		Default:     "",
		Value:       "",
	}}

	testFlags(t, cmd.Flags(), flags)
}
func TestAllFlagsAreAddedToCreateTeamCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	parentCmd := NewTeamCmd(m)
	cmd := NewCreateTeamCmd(m)
	parentCmd.AddCommand(cmd)
	localFlags := []Flag{{
		Description: "Test name is added.",
		Name:        "name",
		Shorthand:   "n",
		Value:       "",
		Default:     "",
	}}
	testFlags(t, cmd.Flags(), localFlags)
	inheritedFlags := []Flag{{
		Description: "Test org is passed from parent",
		Name:        "org",
		Shorthand:   "",
		Value:       "",
		Default:     "",
	}}
	testFlags(t, cmd.InheritedFlags(), inheritedFlags)

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
