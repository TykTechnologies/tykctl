package cmd

import (
	"context"
	"errors"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetTeamById(t *testing.T) {
	testCases := []struct {
		name             string
		mockResponse     cloud.InlineResponse2011
		mockHttpResponse *http.Response
		mockError        error
		ExpectedError    error
		ExpectedTeam     *cloud.Team
	}{
		{
			name: "Test success response",
			mockResponse: cloud.InlineResponse2011{
				Error_:  "",
				Payload: &generateTeams(1)[0],
				Status:  "ok",
			},
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			mockError:        nil,
			ExpectedError:    nil,
			ExpectedTeam:     &cloud.Team{OID: "1", UID: "1"},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().GetTeamById(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			team, err := GetTeamById(context.Background(), m, "orgId", "teamId")
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.ExpectedTeam, team)

		})
	}

}
func TestGetTeams(t *testing.T) {
	testCases := []struct {
		name              string
		mockError         error
		mockResponse      cloud.InlineResponse20017
		mockHttpResponse  *http.Response
		ExpectedOrgLength int
		ExpectedError     error
		ExpectedResponse  []cloud.Team
	}{
		{
			name:      "Check status ok",
			mockError: nil,
			mockResponse: cloud.InlineResponse20017{
				Error_:  "",
				Payload: &cloud.Teams{Teams: generateTeams(3)},
				Status:  "ok",
			},
			mockHttpResponse:  &http.Response{StatusCode: http.StatusOK},
			ExpectedOrgLength: 3,
			ExpectedError:     nil,
			ExpectedResponse:  nil,
		},
		{
			name:              "Test error returned by cloud sdk",
			mockError:         ErrorCreatingTeam,
			mockResponse:      cloud.InlineResponse20017{},
			mockHttpResponse:  &http.Response{StatusCode: http.StatusOK},
			ExpectedOrgLength: 0,
			ExpectedError:     ErrorCreatingTeam,
			ExpectedResponse:  nil,
		},
		{
			name:      "Test status code not equal 200",
			mockError: nil,
			mockResponse: cloud.InlineResponse20017{
				Error_:  "ok",
				Payload: nil,
				Status:  "",
			},
			mockHttpResponse:  &http.Response{StatusCode: http.StatusForbidden},
			ExpectedOrgLength: 0,
			ExpectedError:     ErrorFetchingTeam,
			ExpectedResponse:  nil,
		},
		{
			name:              "Test payload has error",
			mockHttpResponse:  &http.Response{StatusCode: http.StatusOK},
			ExpectedOrgLength: 0,
			mockError:         nil,
			mockResponse: cloud.InlineResponse20017{
				Error_:  "error found here",
				Payload: nil,
				Status:  "error",
			},
			ExpectedResponse: nil,
			ExpectedError:    errors.New("error found here"),
		},
		{
			name:              "Test nil payload",
			mockHttpResponse:  &http.Response{StatusCode: http.StatusOK},
			ExpectedOrgLength: 0,
			mockError:         nil,
			mockResponse: cloud.InlineResponse20017{
				Error_:  "",
				Payload: nil,
				Status:  "ok",
			},
			ExpectedResponse: nil,
			ExpectedError:    nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().GetTeams(gomock.Any(), gomock.Any()).Times(1).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			teams, err := GetTeams(context.Background(), m, "")
			assert.Equal(t, tt.ExpectedError, err)
			if tt.mockResponse.Payload != nil {
				assert.Equal(t, tt.mockResponse.Payload.Teams, teams)
			}
			assert.Equal(t, tt.ExpectedOrgLength, len(teams))
		})
	}
}

func TestCreateTeamHeadersAndRows(t *testing.T) {
	tests := []struct {
		name    string
		teams   []cloud.Team
		rows    [][]string
		headers []string
	}{
		{
			name:    "Test the correct rows are returned",
			rows:    [][]string{{"Test team name", "c9f4a54c-59bb-11ed-9b6a-0242ac120002"}},
			teams:   []cloud.Team{{Name: "Test team name", UID: "c9f4a54c-59bb-11ed-9b6a-0242ac120002"}},
			headers: []string{"Name", "UID", "Environments", "Deployments"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnedHeaders, returnedRows := CreateTeamHeadersAndRows(tt.teams)
			assert.Equal(t, tt.headers, returnedHeaders)
			assert.Equal(t, tt.rows, returnedRows)
		})
	}
}
