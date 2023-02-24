package cloudcmd

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

func TestRestartDeployment(t *testing.T) {

	tests := []struct {
		name             string
		mockHttpResponse *http.Response
		mockResponse     cloud.InlineResponse2001
		mockError        error
		want             *cloud.Deployment
		ExpectedError    error
	}{
		{
			name:             "Check Successful request",
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			mockResponse: cloud.InlineResponse2001{
				Error_: "",
				Payload: &cloud.Deployment{
					OID: "5697876",
				},
				Status: "ok",
			},
			mockError: nil,
			want:      &cloud.Deployment{OID: "5697876"},

			ExpectedError: nil,
		},
		{
			name:             "Check http status 401",
			mockHttpResponse: &http.Response{StatusCode: http.StatusUnauthorized},
			mockResponse:     cloud.InlineResponse2001{},
			mockError:        nil,
			want:             nil,
			ExpectedError:    ErrorRestartingDeployment,
		},
		{
			name:             "Check Payload status field is not okay",
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			mockResponse: cloud.InlineResponse2001{
				Error_:  "an error occurred on the server",
				Payload: nil,
				Status:  "error",
			},
			mockError:     nil,
			want:          nil,
			ExpectedError: errors.New("an error occurred on the server"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().RestartDeployment(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			got, err := restartDeployment(context.Background(), m, "", "", "", "")
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.want, got)

		})
	}
}
