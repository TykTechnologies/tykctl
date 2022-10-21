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
	loginDesc    = `
        This command will login into your cloud account and set the token in your config file.
        Note: The token will only last for 30 minute you will need to login again after 30 minutes.
		You will be prompted to provide your email and  password to login.
		When using the cloud service you should always run this command first as each command will require a token.
		For the staging server you will also need to provide nginx basic auth.
		Sample usage:
		tykctl cloud login --ba-pass=<use this only is staging> --ba-pass=<use this in staging>
`
)

type LoginClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewLoginCommand() *cobra.Command {
	return NewCmd("login").
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

// flags required by the login command.
func addLoginFlags(f *pflag.FlagSet) {
	f.StringP("email", "e", "", "email address you used to login into the dashboard")
	f.StringP("password", "p", "", "password you used to login into the dashboard")
	f.String("ba-user", "", "Basic auth user.This should only be used for staging server")
	err := viper.BindPFlag("ba-user", f.Lookup("ba-user"))
	if err != nil {
		panic(err)
	}
	f.String("ba-pass", "", "Basic auth password")
	err = viper.BindPFlag("ba-pass", f.Lookup("ba-pass"))
	if err != nil {
		panic(err)
	}
	f.StringP("dashboard", "d", dashboardUrl, "Url to connect to the dashboard(Default is the staging url)")
	err = viper.BindPFlag("dashboard", f.Lookup("dashboard"))
	if err != nil {
		panic(err)
	}
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

// /get the response from ara and extract the token returned.
func extractToken(resp *http.Response) (string, error) {
	if resp.StatusCode != 200 && resp.Body != nil {
		///
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			message := err.Error()
			if myerr, ok := err.(swagger.GenericSwaggerError); ok {
				message = string(myerr.Body())
				// handle myerr
			}

			return "", errors.New(message)
		}

		return "", fmt.Errorf("login failed: %s\n", string(b))

	} else if resp.StatusCode != 200 {
		///
		return "", errors.New("login failed")
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
		return "", errors.New("no token found")
	}
	if signature == "" {
		return "", errors.New("signature not found")
	}
	jwt := fmt.Sprintf("%s.%s", token, signature)
	return jwt, nil
}

// validate cli flags and pass them to login
func login(cmd cobra.Command) error {
	dashboard := viper.GetString("dashboard")
	if util.StringIsEmpty(dashboard) {
		return errors.New("dashboard url is required")
	}
	email, err := cmd.Flags().GetString("email")
	if err != nil {
		return err
	}
	err = util.ValidateEmail(email)
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
	return err
}

// save token to configuration file
func getAndSaveToken(url, email, password, basicUser, basicPassword string) error {
	resp, err := dashboardLogin(url, email, password, basicUser, basicPassword)
	if err != nil {
		return err
	}
	token, err := extractToken(resp)
	if err != nil {
		return err
	}
	return SaveToConfig("token", token)
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
