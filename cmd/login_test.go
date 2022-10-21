package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TykTechnologies/tykctl/testutil"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
			ExpectedError: errors.New("signature not found"),
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
			ExpectedError: errors.New("no token found"),
			StatusCode:    200,
		},
		{
			Name:          "Test empty Cookies",
			ShouldErr:     true,
			ExpectedJwt:   "",
			ExpectedError: errors.New("no token found"),
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
			ExpectedError: errors.New(fmt.Sprintf("login failed: %s\n", string([]byte("")))),
			StatusCode:    404,
		},
	}

	for _, model := range testModelList {
		extractTokenRequest(t, model)
	}

}
func TestDashboardLoginRequest(t *testing.T) {
	///url := "https://dash.ara-staging.tyk.technology"
	var testModelList = []DashBoardTestingModel{
		{
			Email:         "itachi.w@tyk.io",
			Password:      "ita.fg7574%¡",
			BasicUser:     "iNUi3OpL",
			BasicPassword: "N$.890TestThus",
		},
		{
			Email:         "sasuke.w@tyk.io",
			Password:      "suke.8%¡",
			BasicUser:     "poy¡",
			BasicPassword: "><<>[ddd][rt]",
		},
		{
			Email:         "[][]{?/v.w@tyk.io",
			Password:      "-[]fk•#",
			BasicUser:     "po904873",
			BasicPassword: "p#cb1djdk",
		},
		{
			Email:         "[][]{?/v.w@tyk.io",
			Password:      "-[]fk•#",
			BasicUser:     "",
			BasicPassword: "",
		},
	}
	for _, model := range testModelList {
		dashboardLoginRequestTest(t, model)
	}

}

func extractTokenRequest(t *testing.T, model ExtractTestModel) {
	/*resp := &http.Response{
		StatusCode: 200,
	}*/
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
	if err != nil && !model.ShouldErr {
		t.Errorf("%s,expected nil error, found %s", model.Name, err)
	}
	if !testutil.EqualError(err, model.ExpectedError) {
		t.Errorf("%s,expected %s error, found %s", model.Name, model.ExpectedError, err)
	}
	if err == nil && model.ShouldErr {
		t.Errorf("%s,expected %s error, found nil", model.Name, err)
	}
	if token != model.ExpectedJwt {
		t.Errorf("%s,expected %s token, found %s", model.Name, model.ExpectedJwt, token)
	}
}

func dashboardLoginRequestTest(t *testing.T, model DashBoardTestingModel) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/login" {
			t.Errorf("Expected to request '/api/login', got: %s", r.URL.Path)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type: application/json header, got: %s", r.Header.Get("Content-Type"))
		}
		username, password, _ := r.BasicAuth()
		if username != model.BasicUser {
			t.Errorf("Username wanted %s, got: %s", "itachi", username)
		}
		if password != model.BasicPassword {
			t.Errorf("Password wanted %s, got: %s", model.BasicPassword, password)
		}
		if r.Body == nil {
			t.Errorf("Body wanted  got nil")
		}
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var body LoginBody
		err = json.Unmarshal(b, &body)
		if err != nil {
			t.Fatal(err)
		}
		if body.Email != model.Email {
			t.Errorf("Email wanted %s, got: %s", model.Email, body.Email)
		}
		if body.Password != model.Password {
			t.Errorf("Password wanted %s, got: %s", model.Password, body.Password)
		}

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
