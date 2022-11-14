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

func TestGetEnvById(t *testing.T) {
	testCases := []struct {
		name             string
		mockHttpResponse *http.Response
		ExpectedError    error
		mockError        error
		ExpectedEnv      *cloud.Loadout
		mockResponse     cloud.InlineResponse2012
	}{
		{
			name:             "Test Success response",
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			ExpectedError:    nil,
			mockError:        nil,
			ExpectedEnv:      &cloud.Loadout{OID: "2", UID: "2"},
			mockResponse: cloud.InlineResponse2012{
				Error_:  "",
				Payload: &generateEnvs(3)[1],
				Status:  "ok",
			},
		},
		{
			name:             "Test error returned by client",
			mockHttpResponse: &http.Response{StatusCode: http.StatusForbidden},
			ExpectedError:    ErrorGenericError,
			mockError:        ErrorGenericError,
			ExpectedEnv:      nil,
			mockResponse: cloud.InlineResponse2012{
				Error_:  "error",
				Payload: nil,
				Status:  "error",
			},
		},
		{
			name:             "Test response status is not 200",
			mockHttpResponse: &http.Response{StatusCode: http.StatusInternalServerError},
			ExpectedError:    ErrorFetchingEnvironment,
			mockError:        nil,
			ExpectedEnv:      nil,
			mockResponse:     cloud.InlineResponse2012{},
		},
		{
			name:             "Test payload status is not ok",
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			ExpectedError:    errors.New("error was found here"),
			mockError:        nil,
			ExpectedEnv:      nil,
			mockResponse: cloud.InlineResponse2012{
				Error_:  "error was found here",
				Payload: nil,
				Status:  "error",
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().GetEnvById(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			env, err := GetEnvById(context.Background(), m, "orgId", "teamId", "envId")
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.ExpectedEnv, env)
		})
	}
}
func TestGetEnvs(t *testing.T) {
	testCases := []struct {
		name              string
		mockHttpResponse  *http.Response
		ExpectedError     error
		mockError         error
		ExpectedEnvs      []cloud.Loadout
		mockResponse      cloud.InlineResponse20016
		ExpectedOrgLength int
	}{
		{
			name:              "Test Success",
			mockHttpResponse:  &http.Response{StatusCode: http.StatusOK},
			ExpectedError:     nil,
			mockError:         nil,
			ExpectedEnvs:      generateEnvs(3),
			ExpectedOrgLength: 3,
			mockResponse: cloud.InlineResponse20016{
				Error_:  "",
				Payload: &cloud.Loadouts{Loadouts: generateEnvs(3)},
				Status:  "ok",
			},
		},
		{
			name:             "Test error returned by client",
			mockHttpResponse: &http.Response{StatusCode: http.StatusNotFound},
			ExpectedError:    ErrorOutPutFormat,
			mockError:        ErrorOutPutFormat,
			ExpectedEnvs:     nil,
			mockResponse: cloud.InlineResponse20016{
				Error_:  "errors",
				Payload: nil,
				Status:  "error",
			},
			ExpectedOrgLength: 0,
		},
		{
			name:              "Test http response is not 200",
			mockHttpResponse:  &http.Response{StatusCode: http.StatusBadGateway},
			ExpectedError:     ErrorFetchingEnvironment,
			mockError:         nil,
			ExpectedEnvs:      nil,
			mockResponse:      cloud.InlineResponse20016{},
			ExpectedOrgLength: 0,
		},
		{
			name:             "Test Payload status is not ok",
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			ExpectedError:    errors.New("i am an error"),
			mockError:        nil,
			ExpectedEnvs:     nil,
			mockResponse: cloud.InlineResponse20016{
				Error_:  "i am an error",
				Payload: nil,
				Status:  "error",
			},
			ExpectedOrgLength: 0,
		},
		{
			name:             "Test payload is nil",
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			ExpectedError:    nil,
			mockError:        nil,
			ExpectedEnvs:     nil,
			mockResponse: cloud.InlineResponse20016{
				Error_:  "",
				Payload: nil,
				Status:  "ok",
			},
			ExpectedOrgLength: 0,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().GetEnvs(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			envs, err := GetEnvs(context.Background(), m, "orgId", "teamId")
			assert.Equal(t, tt.ExpectedError, err)
			if tt.mockResponse.Payload != nil {
				assert.Equal(t, tt.mockResponse.Payload.Loadouts, envs)
			}
			assert.Equal(t, tt.ExpectedOrgLength, len(envs))
		})
	}

}
func TestCreateEnvHeadersAndRows(t *testing.T) {
	tests := []struct {
		name    string
		Envs    []cloud.Loadout
		headers []string
		rows    [][]string
	}{
		{
			name:    "Test correct rows are returned",
			Envs:    []cloud.Loadout{{Name: "First Env", TeamName: "Staging team", UID: "d82bd628-7856-43af-972f-ef38cb80cbbd"}},
			headers: []string{"Name", "UID", "Team", "Active Deployments"},
			rows:    [][]string{{"First Env", "d82bd628-7856-43af-972f-ef38cb80cbbd", "Staging team"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnedHeaders, returnedRows := CreateEnvHeadersAndRows(tt.Envs)
			assert.Equalf(t, tt.headers, returnedHeaders, "CreateEnvHeadersAndRows(%v)", tt.Envs)
			assert.Equalf(t, tt.rows, returnedRows, "CreateEnvHeadersAndRows(%v)", tt.Envs)
		})
	}
}
