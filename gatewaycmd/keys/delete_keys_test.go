package keys

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	mock "github.com/TykTechnologies/tykctl/gatewaycmd/mocks"
	"github.com/TykTechnologies/tykctl/util"
)

func TestDeleteKeyByID(t *testing.T) {
	apiStatusMessage := apim.ApiStatusMessage{
		Message: util.GetStrPtr("ok"),
		Status:  util.GetStrPtr("success"),
	}

	tests := []struct {
		name             string
		mockError        error
		mockHTTPResponse *http.Response
		mockResponse     *apim.ApiStatusMessage
		want             *apim.ApiStatusMessage
		id               string
		ExpectedError    error
	}{
		{
			name:      "Test Success",
			mockError: nil,
			mockHTTPResponse: &http.Response{
				StatusCode: http.StatusOK,
			},
			mockResponse:  &apiStatusMessage,
			want:          &apiStatusMessage,
			id:            "89du",
			ExpectedError: nil,
		},
		{
			name:             "Test key not found",
			mockError:        errors.New("there is no such key found"),
			mockHTTPResponse: &http.Response{StatusCode: 404, Status: "not found"},
			mockResponse:     nil,
			want:             nil,
			id:               "515c1cb4b5a847139177864d7931e9de",
			ExpectedError:    errors.New("there is no such key found"),
		},
		{
			name:      "Test http error code",
			mockError: nil,
			mockHTTPResponse: &http.Response{
				StatusCode: http.StatusMethodNotAllowed,
				Status:     "method not allowed",
			},
			mockResponse:  &apiStatusMessage,
			want:          nil,
			id:            "467e",
			ExpectedError: errors.New("method not allowed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockKeysAPI(ctrl)
			m.EXPECT().DeleteKey(gomock.Any(), gomock.Any())
			m.EXPECT().DeleteKeyExecute(gomock.Any()).Return(tt.mockResponse, tt.mockHTTPResponse, tt.mockError)
			got, err := deleteKeyByID(context.Background(), m, tt.id)
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
