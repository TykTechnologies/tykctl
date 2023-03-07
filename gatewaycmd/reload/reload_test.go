package reload

import (
	"context"
	"errors"
	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	mock "github.com/TykTechnologies/tykctl/gatewaycmd/mocks"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/testutil"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

var (
	ErrorSecretRequired = errors.New("a secret is required to perform this action")
)

func TestReloadFlags(t *testing.T) {
	factory := internal.ApimFactory{}
	cmd := NewReloadCmd(factory)
	localFlags := []internal.Flag{{
		Description: "Test block flag",
		Name:        "block",
		Shorthand:   "b",
		Value:       "false",
		Default:     "false",
	}, {
		Description: "Test group flag",
		Name:        "group",
		Shorthand:   "g",
		Value:       "false",
		Default:     "false",
	},
	}
	testutil.TestFlags(t, cmd.Flags(), localFlags)
}
func TestReloadSingleNode(t *testing.T) {
	message := "success"
	status := "ok"
	tests := []struct {
		name             string
		want             *apim.ApiStatusMessage
		mockError        error
		mockHttpResponse *http.Response
		mockResponse     *apim.ApiStatusMessage
		ExpectedError    error
	}{
		{
			name:      "Test gateway Error",
			want:      nil,
			mockError: ErrorSecretRequired,
			mockHttpResponse: &http.Response{
				Status:     "secret is required",
				StatusCode: 403,
			},
			mockResponse:  &apim.ApiStatusMessage{},
			ExpectedError: ErrorSecretRequired,
		},
		{
			name:      "Test http error code",
			want:      nil,
			mockError: nil,
			mockHttpResponse: &http.Response{
				Status:     "not found",
				StatusCode: 404,
			},
			mockResponse:  &apim.ApiStatusMessage{},
			ExpectedError: errors.New("not found"),
		},
		{
			name: "Test Success",
			want: &apim.ApiStatusMessage{
				Message: &message,
				Status:  &status,
			},
			mockError: nil,
			mockHttpResponse: &http.Response{
				StatusCode: http.StatusOK,
			},
			mockResponse: &apim.ApiStatusMessage{
				Message: &message,
				Status:  &status,
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockHotReloadAPI(ctrl)
			m.EXPECT().HotReload(gomock.Any())
			m.EXPECT().HotReloadExecute(gomock.Any()).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			got, err := reloadSingleNode(context.Background(), m, true)
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReloadGroup(t *testing.T) {
	message := "success"
	status := "ok"
	administrativeAccess := "attempted administrative access"
	tests := []struct {
		name             string
		want             *apim.ApiStatusMessage
		mockError        error
		mockHttpResponse *http.Response
		mockResponse     *apim.ApiStatusMessage
		ExpectedError    error
	}{
		{
			name: "Test Success",
			want: &apim.ApiStatusMessage{
				Message: &message,
				Status:  &status,
			},
			mockError: nil,
			mockHttpResponse: &http.Response{
				StatusCode: http.StatusOK,
			},
			mockResponse: &apim.ApiStatusMessage{
				Message: &message,
				Status:  &status,
			},
			ExpectedError: nil,
		},
		{
			name:      "Test http error 500",
			want:      nil,
			mockError: nil,
			mockHttpResponse: &http.Response{
				Status:     "internal server error",
				StatusCode: http.StatusInternalServerError,
			},
			mockResponse:  nil,
			ExpectedError: errors.New("internal server error"),
		},
		{
			name:      "Test error returned by gateway",
			want:      nil,
			mockError: errors.New(administrativeAccess),
			mockHttpResponse: &http.Response{
				Status:     administrativeAccess,
				StatusCode: 401,
			},
			mockResponse:  &apim.ApiStatusMessage{},
			ExpectedError: errors.New(administrativeAccess),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockHotReloadAPI(ctrl)
			m.EXPECT().HotReloadGroup(gomock.Any())
			m.EXPECT().HotReloadGroupExecute(gomock.Any()).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			got, err := reloadGroup(context.Background(), m)
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReloadGateway(t *testing.T) {
	message := "success"
	status := "ok"
	tests := []struct {
		name                      string
		block                     bool
		group                     bool
		numberOfTimesToCallGroup  int
		numberOfTimesToCallReload int
		want                      *apim.ApiStatusMessage
		mockError                 error
		mockHttpResponse          *http.Response
		mockResponse              *apim.ApiStatusMessage
		ExpectedError             error
	}{
		{
			name:                      "Test that reload group is called when group is true",
			block:                     false,
			group:                     true,
			numberOfTimesToCallGroup:  1,
			numberOfTimesToCallReload: 0,
			want: &apim.ApiStatusMessage{
				Message: &message,
				Status:  &status,
			},
			mockError:        nil,
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			mockResponse: &apim.ApiStatusMessage{
				Message: &message,
				Status:  &status,
			},
			ExpectedError: nil,
		},
		{
			name:                      "Test that single node is called when group is false",
			block:                     true,
			group:                     false,
			numberOfTimesToCallGroup:  0,
			numberOfTimesToCallReload: 1,
			want: &apim.ApiStatusMessage{
				Message: &message,
				Status:  &status,
			},
			mockError:        nil,
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			mockResponse: &apim.ApiStatusMessage{
				Message: &message,
				Status:  &status,
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockHotReloadAPI(ctrl)
			m.EXPECT().HotReloadGroup(gomock.Any()).Times(tt.numberOfTimesToCallGroup)
			m.EXPECT().HotReloadGroupExecute(gomock.Any()).Times(tt.numberOfTimesToCallGroup).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			m.EXPECT().HotReload(gomock.Any()).Times(tt.numberOfTimesToCallReload)
			m.EXPECT().HotReloadExecute(gomock.Any()).Times(tt.numberOfTimesToCallReload).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			got, err := reloadGateway(context.Background(), m, tt.block, tt.group)
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
