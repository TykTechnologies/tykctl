package cloudcmd

import (
	"encoding/json"
	"fmt"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/testutil"
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
	cmd := NewLoginCommand(nil)
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
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for _, cookie := range model.Cookies {
			http.SetCookie(w, cookie)
		}
		w.WriteHeader(model.StatusCode)

	}))
	defer s.Close()
	response, err := mockHttp(s.URL)
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
	_, err := dashboardLogin(s.URL, model.Email, model.Password, model.BasicUser, model.BasicPassword)
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

func mockHttp(url string) (*http.Response, error) {

	return http.Get(url)
}
