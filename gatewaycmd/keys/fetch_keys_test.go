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

func TestCreateKeysHeadersAndRows(t *testing.T) {
	tests := []struct {
		name    string
		keys    []string
		headers []string
		rows    [][]string
	}{
		{
			name:    "Test the correct and headers are returned",
			keys:    []string{"First Key", "bb4fc08c-34c0-45b2-8780-acfdc22e68be", "1"},
			headers: []string{"Key"},
			rows:    [][]string{{"First Key"}, {"bb4fc08c-34c0-45b2-8780-acfdc22e68be"}, {"1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnedHeaders, returnedRows := createKeysHeadersAndRows(tt.keys)
			assert.Equal(t, tt.headers, returnedHeaders)
			assert.Equal(t, tt.rows, returnedRows)
		})
	}
}

func TestFetchKeys(t *testing.T) {
	tests := []struct {
		name             string
		want             *apim.ApiAllKeys
		mockError        error
		mockHTTPResponse *http.Response
		mockResponse     *apim.ApiAllKeys
		ExpectedError    error
	}{
		{
			name:      "Test gateway error",
			want:      nil,
			mockError: errors.New("attempted administrative access"),
			mockHTTPResponse: &http.Response{
				Status:     "secret is required",
				StatusCode: 403,
			},
			mockResponse:  &apim.ApiAllKeys{Keys: []string{"first-key"}},
			ExpectedError: errors.New("attempted administrative access"),
		},
		{
			name: "Test successful request",
			want: &apim.ApiAllKeys{
				Keys: []string{"8101ace4-f59b-4dcf-893d-f862f03767f6", "28ed59ca"},
			},
			mockError: nil,
			mockHTTPResponse: &http.Response{
				StatusCode: http.StatusOK,
			},
			mockResponse: &apim.ApiAllKeys{
				Keys: []string{"8101ace4-f59b-4dcf-893d-f862f03767f6", "28ed59ca"},
			},
			ExpectedError: nil,
		},
		{
			name:      "Test http error code",
			want:      nil,
			mockError: nil,
			mockHTTPResponse: &http.Response{
				Status:     "token is needed",
				StatusCode: http.StatusUnauthorized,
			},
			mockResponse:  nil,
			ExpectedError: errors.New("token is needed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockKeysAPI(ctrl)
			m.EXPECT().ListKeys(gomock.Any())
			m.EXPECT().ListKeysExecute(gomock.Any()).Return(tt.mockResponse, tt.mockHTTPResponse, tt.mockError)
			keys, err := fetchKeys(context.Background(), m)
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.want, keys)
		})
	}
}

func TestFetchKeyByID(t *testing.T) {
	sessionState := apim.SessionState{
		Tags:          []string{"da79826a-e39a-4f69-946c-0a5b4918b8f2"},
		AccessRights:  nil,
		Alias:         util.GetStrPtr("alias"),
		ApplyPolicies: []string{"34", "40"},
		Certificate:   util.GetStrPtr("my-certificate"),
		IsInactive:    util.GetBoolPtr(true),
		OrgId:         util.GetStrPtr("third-org"),
	}

	tests := []struct {
		want             *apim.SessionState
		name             string
		mockError        error
		mockHTTPResponse *http.Response
		mockResponse     *apim.SessionState
		ExpectedError    error
		id               string
	}{
		{
			name:      "Test Success",
			want:      &sessionState,
			mockError: nil,
			mockHTTPResponse: &http.Response{
				StatusCode: http.StatusOK,
			},
			mockResponse:  &sessionState,
			ExpectedError: nil,
			id:            "d5455edf-7047-41ee-87d7-5657fa972bc8",
		},
		{
			want:             nil,
			name:             "Test http error code",
			mockError:        nil,
			mockHTTPResponse: &http.Response{StatusCode: 405, Status: "method not allowed"},
			mockResponse:     nil,
			ExpectedError:    errors.New("method not allowed"),
			id:               "5678583f",
		},
		{
			want:             nil,
			name:             "Test gateway error",
			mockError:        errors.New("key not found"),
			mockHTTPResponse: &http.Response{StatusCode: 404, Status: "not found"},
			mockResponse:     nil,
			ExpectedError:    errors.New("key not found"),
			id:               "6786546",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockKeysAPI(ctrl)
			m.EXPECT().GetKey(gomock.Any(), gomock.Any())
			m.EXPECT().GetKeyExecute(gomock.Any()).Return(tt.mockResponse, tt.mockHTTPResponse, tt.mockError)
			got, err := fetchKeyByID(context.Background(), m, tt.id)
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
