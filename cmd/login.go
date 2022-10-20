package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/swagger-gen"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

const (
	dashboardUrl = "https://dash.ara-staging.tyk.technology"
)

type LoginClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewLoginCommand() *cobra.Command {
	return NewCmd("login", "login").
		WithLongDescription(loginDesc).
		WithFlagAdder(false, addLoginFlags).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			err := login(cmd)
			if err != nil {
				cmd.Println(err)
				return err
			}
			return err
			///	return loginService.dashboardLogin()
		})

}

func login(cmd cobra.Command) error {
	dashboard := viper.GetString("dashboard")
	if util.StringIsEmpty(dashboard) {
		return errors.New("dashboard url is required")
	}
	email, err := cmd.Flags().GetString("email")
	if err != nil {
		return err
	}
	err = util.ValidateEmail(email, "email address is required")
	if err != nil {
		return err
	}
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		return err
	}
	if util.StringIsEmpty(password) {
		return errors.New("password is required")
	}
	baUser := viper.GetString("ba-user")
	baPass := viper.GetString("ba-pass")
	err = getAndSaveToken(dashboard, email, password, baUser, baPass)
	if err != nil {
		return err
	}
	///dashboardLogin(dashboard)
	return err
}

// flags required by the login command.
func addLoginFlags(f *pflag.FlagSet) {
	f.StringP("email", "e", "", "email address you used to login into the dashboard")
	f.StringP("password", "p", "", "email address you used to login into the dashboard")
	f.String("ba-user", "", "Basic auth user.This should only be used for staging server")
	viper.BindPFlag("ba-user", loginCmd.Flags().Lookup("ba-user"))
	f.String("ba-pass", "", "Basic auth password")
	viper.BindPFlag("ba-pass", loginCmd.Flags().Lookup("ba-pass"))
	f.String("dashboard", dashboardUrl, "Url to connect to the dashboard(Default is the staging url)")
	viper.BindPFlag("dashboard", loginCmd.Flags().Lookup("dashboard"))
}

// /dashboardLogin send a request to ara dashboard to get a token to use for login.
func dashboardLogin(url, email, password, basicUser, basicPassword string) (*http.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	body := LoginBody{
		Email:    email,
		Password: password,
	}
	fullUrl := fmt.Sprintf("%s/api/login", url)
	req, err := internal.CreatePostRequest(fullUrl, body, headers)
	if err != nil {
		return nil, err
	}
	if len(basicUser) > 0 && len(basicPassword) > 0 {
		req.SetBasicAuth(basicUser, basicPassword)
	}
	loginClient := &http.Client{}
	response, err := loginClient.Do(req)
	if err != nil {
		return nil, err
	}
	return response, err
}

func getAndSaveToken(url, email, password, basicUser, basicPassword string) error {
	resp, err := dashboardLogin(url, email, password, basicUser, basicPassword)
	if err != nil {
		return nil
	}
	if resp.StatusCode != 200 {
		if resp.Body != nil {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				message := err.Error()
				if myerr, ok := err.(swagger.GenericSwaggerError); ok {
					message = string(myerr.Body())
					// handle myerr
				}

				return errors.New(message)
			}

			return errors.New(fmt.Sprintf("Login failed: %s\n", string(b)))

		}
	}
	var token string
	var signature string
	for _, cookie := range resp.Cookies() {
		switch cookie.Name {
		case "cookieAuthorisation":
			token = cookie.Value

		case "signature":

			signature = cookie.Value
		}

	}
	if len(token) == 0 {
		return errors.New("no token found")
	}
	if signature == "" {
		return errors.New("signature not found")
	}
	jwt := fmt.Sprintf("%s.%s", token, signature)
	return SaveToConfig("token", jwt)
}

func SaveToConfig(key, token string) error {
	viper.Set(key, token)
	err := viper.WriteConfig()
	if err != nil {
		message := fmt.Sprintf("Couldn't write config: %s\n", err.Error())
		return errors.New(message)
	}
	return err

}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}
