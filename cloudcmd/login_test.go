package cloudcmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/TykTechnologies/tykctl/testutil"
	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddLoginFlags(t *testing.T) {
	cmd := internal.NewCmd("test").WithFlagAdder(false, addLoginFlags).NoArgs(nil)
	flags := []internal.Flag{{
		Description: "Test email address",
		Name:        "email",
		Shorthand:   "e",
		Default:     "",
		Value:       "",
	}, {
		Description: "Test password",
		Name:        "password",
		Shorthand:   "p",
		Value:       "",
		Default:     "",
	},
		{
			Description: "Test ba-user",
			Name:        "ba-user",
			Shorthand:   "",
			Value:       "",
			Default:     "",
		},
		{
			Description: "Test ba-pass",
			Name:        "ba-pass",
			Shorthand:   "",
			Value:       "",
			Default:     "",
		},
	}
	testutil.TestFlags(t, cmd.Flags(), flags)
}
func TestExtractToken(t *testing.T) {
	var testModelList = []ExtractTestModel{
		{
			Cookies: []*http.Cookie{{
				Name:  "cookieAuthorisation",
				Value: "hello",
			}, {
				Name:  "signature",
				Value: "there",
			},
			},
			Name:          "Test Jwt is Extracted",
			ShouldErr:     false,
			ExpectedJwt:   "hello.there",
			ExpectedError: nil,
			StatusCode:    200,
		},

		{
			Cookies: []*http.Cookie{{
				Name:  "cookieAuthorisation",
				Value: "hello",
			},
			},
			Name:          "Test empty signature cookie",
			ShouldErr:     true,
			ExpectedJwt:   "",
			ExpectedError: ErrSignatureNotFound,
			StatusCode:    200,
		},
		{
			Cookies: []*http.Cookie{{
				Name:  "signature",
				Value: "hello",
			},
			},
			Name:          "Test empty cookieAuthorisation cookie",
			ShouldErr:     true,
			ExpectedJwt:   "",
			ExpectedError: ErrTokenNotFound,
			StatusCode:    200,
		},
		{
			Name:          "Test empty Cookies",
			ShouldErr:     true,
			ExpectedJwt:   "",
			ExpectedError: ErrTokenNotFound,
			StatusCode:    200,
		},

		{
			Cookies: []*http.Cookie{{
				Name:  "cookieAuthorisation",
				Value: "hello",
			},
				{
					Name:  "signature",
					Value: "there",
				},
			},
			Name:          "Test status code 404",
			ShouldErr:     true,
			ExpectedJwt:   "",
			ExpectedError: fmt.Errorf("login failed: %s\n", string([]byte(""))),
			StatusCode:    404,
		},
	}
	for _, tt := range testModelList {
		t.Run(tt.Name, func(t *testing.T) {
			extractTokenRequest(t, tt)
		})
	}

}
func TestDashboardLoginRequest(t *testing.T) {
	///url := "https://dash.ara-staging.tyk.technology"
	var testModelList = []DashBoardTestingModel{
		{
			Description:   "Test All values Presents",
			Email:         "itachi.w@tyk.io",
			Password:      "ita.fg7574%¡",
			BasicUser:     "iNUi3OpL",
			BasicPassword: "N$.890TestThus",
		},
		{
			Description:   "Test All values present",
			Email:         "sasuke.w@tyk.io",
			Password:      "suke.8%¡",
			BasicUser:     "poy¡",
			BasicPassword: "><<>[ddd][rt]",
		},
		{
			Description:   "Test All values present",
			Email:         "[][]{?/v.w@tyk.io",
			Password:      "-[]fk•#",
			BasicUser:     "po904873",
			BasicPassword: "p#cb1djdk",
		},
		{
			Description:   "Test Basic user and auth absent",
			Email:         "[][]{?/v.w@tyk.io",
			Password:      "-[]fk•#",
			BasicUser:     "",
			BasicPassword: "",
		},
	}
	for _, tt := range testModelList {
		t.Run(tt.Description, func(t *testing.T) {
			dashboardLoginRequestTest(t, tt)
		})
	}

}

func TestNewLoginCommand(t *testing.T) {
	cmd := NewLoginCommand(internal.CloudFactory{})
	flags := []internal.Flag{{
		Description: "Test email address is passed to login command",
		Name:        "email",
		Shorthand:   "e",
		Default:     "",
		Value:       "",
	}, {
		Description: "Test password is passed to login command",
		Name:        "password",
		Shorthand:   "p",
		Value:       "",
		Default:     "",
	},
		{
			Description: "Test ba-user is passed to login command",
			Name:        "ba-user",
			Shorthand:   "",
			Value:       "",
			Default:     "",
		},
		{
			Description: "Test ba-pass is passed to login command",
			Name:        "ba-pass",
			Shorthand:   "",
			Value:       "",
			Default:     "",
		},
	}
	testutil.TestFlags(t, cmd.Flags(), flags)

}

func extractTokenRequest(t *testing.T, model ExtractTestModel) {
	t.Helper()
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for _, cookie := range model.Cookies {
			http.SetCookie(w, cookie)
		}
		w.WriteHeader(model.StatusCode)

	}))
	defer s.Close()
	response, err := mockHttp(context.Background(), s.URL)
	if err != nil {
		t.Fatal(err)
	}

	token, err := extractToken(response)
	if !model.ShouldErr {
		assert.NoError(t, err)
	}
	if model.ShouldErr && model.ExpectedError != nil {
		assert.Error(t, err, model.Name)
	}
	assert.Equal(t, model.ExpectedError, err, model.Name)
	assert.Equal(t, model.ExpectedJwt, token, "wrong token returned")
}

func dashboardLoginRequestTest(t *testing.T, model DashBoardTestingModel) {
	t.Helper()
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/login", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		username, password, _ := r.BasicAuth()
		assert.Equal(t, model.BasicUser, username)
		assert.Equal(t, model.BasicPassword, password)
		assert.NotNil(t, r.Body)
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var body LoginBody
		err = json.Unmarshal(b, &body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, model.Email, body.Email)
		assert.Equal(t, model.Password, body.Password)
		w.WriteHeader(http.StatusOK)

	}))
	defer s.Close()
	_, err := dashboardLogin(context.Background(), s.URL, model.Email, model.Password, model.BasicUser, model.BasicPassword)
	if err != nil {
		t.Fatal(err)
		///st.Expect(t, err, nil)
	}
}

type DashBoardTestingModel struct {
	Description   string
	Email         string
	Password      string
	BasicUser     string
	BasicPassword string
}

type ExtractTestModel struct {
	Cookies       []*http.Cookie
	Name          string
	ExpectedJwt   string
	ShouldErr     bool
	ExpectedError error
	StatusCode    int
}

func mockHttp(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func TestGetUserRole(t *testing.T) {
	tests := []struct {
		name          string
		roles         []internal.Role
		want          *internal.Role
		ExpectedError error
	}{
		{
			name:          "Test has team and org",
			ExpectedError: nil,
			roles: []internal.Role{{
				Role:      "billing_admin",
				OrgID:     "12fGHtmi567",
				TeamID:    "",
				OrgName:   "data test",
				TeamName:  "",
				AccountID: "",
			},
				{
					Role:      "org_admin",
					OrgID:     "24568674d",
					TeamID:    "",
					OrgName:   "dx org",
					TeamName:  "itachi team",
					AccountID: "78906756",
				},
			},
			want: &internal.Role{
				Role:      "org_admin",
				OrgID:     "24568674d",
				TeamID:    "",
				OrgName:   "dx org",
				TeamName:  "itachi team",
				AccountID: "78906756",
			},
		},
		{
			name: "Test Empty Role",
			roles: []internal.Role{{
				Role:      "",
				OrgID:     "457685098",
				TeamID:    "6547586",
				OrgName:   "My org",
				TeamName:  "Org Name",
				AccountID: "",
			}},
			want:          nil,
			ExpectedError: ErrorNoRoleFound,
		},
		{
			name: "Test has invalid role",
			roles: []internal.Role{{
				Role:      "org_admi",
				OrgID:     "45689675f",
				TeamID:    "y123465j5",
				OrgName:   "",
				TeamName:  "",
				AccountID: "",
			},
			},
			ExpectedError: ErrorNoRoleFound,
			want:          nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			role, err := getUserRole(tt.roles)
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equalf(t, tt.want, role, "getUserRole(%v)", tt.roles)
		})
	}
}

func TestInitOrgInfo(t *testing.T) {
	testCases := []struct {
		name               string
		mockResponse       *internal.OrgInfo
		mockHttpResponse   *resty.Response
		mockError          error
		ExpectedError      error
		want               *internal.OrgInit
		orgId              string
		teamPromptResponse *cloud.Team
		teamPromptError    error
		teamPromptCalls    int
	}{
		{
			name:         "Test GetOrgInfo returns an error",
			mockResponse: nil,
			mockHttpResponse: &resty.Response{
				RawResponse: &http.Response{StatusCode: http.StatusForbidden},
			},
			mockError:          ErrorGenericError,
			ExpectedError:      ErrorGenericError,
			want:               nil,
			orgId:              "",
			teamPromptResponse: nil,
			teamPromptError:    nil,
			teamPromptCalls:    0,
		},

		{
			name: "Test Success",
			mockResponse: &internal.OrgInfo{Organisation: cloud.Organisation{
				Zone:  "aws-us-west-2",
				Teams: generateTeams(1),
			}},
			mockHttpResponse: &resty.Response{
				RawResponse: &http.Response{StatusCode: http.StatusOK},
			},
			mockError:     nil,
			ExpectedError: nil,
			want: &internal.OrgInit{
				Controller: "https://controller-aws-usw2.cloud-ara.tyk.io:37001",
				Org:        "helloOrg", Team: "654536rty56",
			},
			orgId: "helloOrg",
			teamPromptResponse: &cloud.Team{
				OID: "4598756363",
				UID: "654536rty56",
			},
			teamPromptError: nil,
			teamPromptCalls: 1,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			prompt := mock.NewMockCloudPrompt(ctrl)
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().GetOrgInfo(gomock.Any(), gomock.Any()).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			prompt.EXPECT().TeamPrompt(gomock.Any()).Return(tt.teamPromptResponse, tt.teamPromptError).Times(tt.teamPromptCalls)
			info, err := initOrgInfo(context.Background(), m, prompt, tt.orgId)
			assert.Equal(t, tt.ExpectedError, err)
			assert.Equal(t, tt.want, info)

		})
	}
}

func TestInitUserProfile(t *testing.T) {
	tests := []struct {
		name             string
		mockResponse     *internal.UserInfo
		mockHttpResponse *resty.Response
		mockError        error
		want             *internal.Role
		ExpectedError    error
	}{
		{
			name:         "Test 401 http error code",
			mockResponse: nil,
			mockHttpResponse: &resty.Response{
				RawResponse: &http.Response{StatusCode: http.StatusUnauthorized},
			},
			mockError:     ErrorGenericError,
			want:          nil,
			ExpectedError: ErrorGenericError,
		},

		{
			name: "Test Success Response",
			mockResponse: &internal.UserInfo{
				Email:     "",
				LastName:  "",
				AccountID: "",
				Roles: []internal.Role{{
					Role:   "org_admin",
					OrgID:  "986765",
					TeamID: "8908756y",
				}},
			},
			mockHttpResponse: &resty.Response{
				RawResponse: &http.Response{StatusCode: http.StatusOK},
			},
			mockError: nil,
			want: &internal.Role{
				Role:   "org_admin",
				OrgID:  "986765",
				TeamID: "8908756y",
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().GetUserInfo(gomock.Any()).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			got, err := initUserProfile(context.Background(), m)
			assert.Equalf(t, tt.want, got, "initUserProfile")
			assert.Equal(t, tt.ExpectedError, err)
		})
	}
}
