package cloudcmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	loginDesc = `
        This command will login into your cloud account and set the token in your config file.
       
        Note: The token will only last for 30 minute you will need to login again after 30 minutes.
		
        You will be prompted to provide your email and  password to login if you use the interactive mode.
		When using the cloud service you should always run this command first as each command will require a token.
		
        For the staging server you will also need to provide nginx basic auth.
		
        Sample usage:
		
         tykctl cloud login --ba-pass=<use this only is staging> --ba-pass=<use this in staging>
`
)

var (
	ErrTokenNotFound      = errors.New("no token found")
	ErrLoginFailed        = errors.New("login failed")
	ErrSignatureNotFound  = errors.New("signature not found")
	ErrPasswordIsRequired = errors.New("password is required")
	ErrNoOrganization     = errors.New("you do not have any organization")
)

// NewLoginCommand creates a new login command.
func NewLoginCommand(client internal.CloudClient) *cobra.Command {
	return internal.NewCmd(login).
		WithLongDescription(loginDesc).
		WithDescription("login to tyk cloud using password and email.").
		WithExample("tykctl cloud login --password=<your cloud password here> --email=<your email here>").
		WithFlagAdder(false, addLoginFlags).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			err := validateAndLogin(cmd.Flags())
			if err != nil {
				cmd.Println(err)
				return err
			}
			profile, err := initUserProfile(ctx, client)
			if err != nil {
				log.Println(err)
				return err
			}
			err = internal.SaveMapToConfig(profile)
			if err != nil {
				cmd.Println(err)
				return err
			}

			orgId := viper.GetString(org)
			if orgId == "" {
				cmd.Println("You need to create an organization here https://dashboard.cloud-ara.tyk.io/")
				return ErrNoOrganization
			}
			orgInfo, err := initOrgInfo(ctx, client, orgId)
			if err != nil {
				cmd.Println(err)
				return err
			}
			err = internal.SaveMapToConfig(orgInfo)
			if err != nil {
				cmd.Println(err)
				return err
			}
			cmd.Println("Authentication successful")
			return nil
		})

}

// initUserProfile will auto fetch user info such as:
// user roles,user team
func initUserProfile(ctx context.Context, client internal.CloudClient) (map[string]string, error) {
	userInfo, _, err := client.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	profile := getUserRole(userInfo.Roles)
	return profile, nil
}

// initOrgInfo will fetch the user organization and extract team and create a controllerUrl that
// the user can use to connect to tyk cloud depending on their region.
func initOrgInfo(ctx context.Context, client internal.CloudClient, prompt CloudPrompt, orgId string) (map[string]string, error) {
	info, _, err := client.GetOrgInfo(ctx, orgId)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	controllerUrl, err := util.GenerateUrlFromZone(info.Organisation.Zone)
	if err != nil {
		return m, err
	}
	m[controller] = controllerUrl
	m[org] = orgId
	team, err := prompt.teamPrompt(info.Organisation.Teams)

	if len(info.Organisation.Teams) == 1 {
		m[team] = info.Organisation.Teams[0].UID
	} else if len(info.Organisation.Teams) > 1 {
		team, err := prompt.teamPrompt(info.Organisation.Teams)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}
func selectDefaultTeam(prompt CloudPrompt, teams []cloud.Team) (*cloud.Team, error) {
	return prompt.teamPrompt(teams)
}

// getUserRole returns the user role.
func getUserRole(roles []internal.Role) map[string]string {
	roleList := []string{"org_admin", "team_admin", "team_member"}
	m := make(map[string]string)
	for _, role := range roles {
		contain := slices.Contains(roleList, role.Role)
		if contain {
			m[userRole] = role.Role
			if role.OrgID != "" {
				m[org] = role.OrgID
			}
			if role.TeamID != "" {
				m[team] = role.TeamID
			}
			return m

		}
	}
	return m
}

// addLoginFlags add the flags required by the login command.
func addLoginFlags(f *pflag.FlagSet) {
	f.StringP(email, "e", "", "email address you used to login into the dashboard")
	f.StringP(password, "p", "", "password you used to login into the dashboard")
	f.String(baUser, "", "Basic auth user.This should only be used for staging server")
	f.BoolP(interactive, "i", false, "login using the interactive mode.")
	err := viper.BindPFlag(baUser, f.Lookup(baUser))
	if err != nil {
		panic(err)
	}
	f.String(baPass, "", "Basic auth password")
	err = viper.BindPFlag(baPass, f.Lookup(baPass))
	if err != nil {
		panic(err)
	}
}

// dashboardLogin send a request to ara dashboard to get a token to use to authenticate all other requests.
func dashboardLogin(baseUrl, email, password, basicUser, basicPassword string) (*http.Response, error) {
	headers := map[string]string{
		contentType: applicationJson,
	}
	body := LoginBody{
		Email:    email,
		Password: password,
	}
	fullUrl, err := url.JoinPath(baseUrl, loginPath)
	if err != nil {
		return nil, err
	}
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
	return response, nil
}

// extractToken takes the response from ara and extract the token returned.
func extractToken(resp *http.Response) (string, error) {
	if resp.StatusCode != http.StatusOK && resp.Body != nil {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("login failed: %s\n", string(b))

	} else if resp.StatusCode != http.StatusOK {
		return "", ErrLoginFailed

	}
	var token string
	var cookieSignature string
	for _, cookie := range resp.Cookies() {
		switch cookie.Name {
		case cookieAuthorisation:
			token = cookie.Value

		case signature:

			cookieSignature = cookie.Value
		}

	}
	if len(token) == 0 {
		return "", ErrTokenNotFound
	}
	if cookieSignature == "" {
		return "", ErrSignatureNotFound
	}
	return fmt.Sprintf("%s.%s", token, cookieSignature), nil
}

// validateAndLogin validate cli flags and pass them to login.
func validateAndLogin(f *pflag.FlagSet) error {
	isInteractive, err := f.GetBool(interactive)
	if err != nil {
		return err
	}
	var loginBody *LoginBody
	if isInteractive {
		loginBody, err = loginInteractive()
		if err != nil {
			return err
		}
	} else {
		loginBody, err = loginWithFlag(f)
		if err != nil {
			return err
		}
	}
	err = util.ValidateEmail(loginBody.Email)
	if err != nil {
		return err
	}
	if util.StringIsEmpty(loginBody.Password) {
		return ErrPasswordIsRequired
	}
	baUser := viper.GetString(baUser)
	baPass := viper.GetString(baPass)
	err = getAndSaveToken(internal.DashboardUrl, loginBody.Email, loginBody.Password, baUser, baPass)
	if err != nil {
		return err
	}
	return nil
}

// loginWithFlag will extract email and password from the flags.
func loginWithFlag(f *pflag.FlagSet) (*LoginBody, error) {
	email, err := f.GetString(email)
	if err != nil {
		return nil, err
	}
	password, err := f.GetString(password)
	if err != nil {
		return nil, err
	}
	return &LoginBody{
		Email:    email,
		Password: password,
	}, nil
}

// loginInteractive will extract ask user to enter login details interactively.
func loginInteractive() (*LoginBody, error) {
	email, err := emailPrompt()
	if err != nil {
		return nil, err
	}
	password, err := passwordPrompt()
	if err != nil {
		return nil, err
	}
	return &LoginBody{
		Email:    email,
		Password: password,
	}, nil
}

// getAndSaveToken token to configuration file.
func getAndSaveToken(url, email, password, basicUser, basicPassword string) error {
	resp, err := dashboardLogin(url, email, password, basicUser, basicPassword)
	if err != nil {
		return err
	}
	token, err := extractToken(resp)

	if err != nil {
		return err
	}
	return internal.SaveToConfig("token", token)
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
