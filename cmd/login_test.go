package cmd

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
