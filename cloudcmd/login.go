package cloudcmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"io"
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
	ErrorNoRoleFound      = errors.New("role not found")
)

// NewLoginCommand creates a new login command.
func NewLoginCommand(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(login).
		WithLongDescription(loginDesc).
		WithDescription("login to tyk cloud using password and email.").
		WithExample("tykctl cloud login --password=<your cloud password here> --email=<your email here>").
		WithFlagAdder(false, addLoginFlags).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			err := validateAndLogin(cmd.Flags())
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			info, err := GetUserInfo(ctx, factory.Client)
			if err != nil {
				return err
			}
			role, err := getUserRole(info.Roles)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			err = internal.SaveMapToCloudUserContext(info.ID, role.RoleToMap())
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}

			orgId := viper.GetString(internal.CreateKeyFromPath(cloudPath, info.ID, org))
			if orgId == "" {
				cmd.Println("You need to create an organization here https://dashboard.cloud-ara.tyk.io/")
				return ErrNoOrganization
			}
			orgInfo, err := initOrgInfo(ctx, factory.Client, factory.Prompt, orgId)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			err = internal.SaveMapToCloudUserContext(info.ID, orgInfo.OrgInitToMap())
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			err = internal.SaveToConfig(currentCloudUser, info.ID)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			cmd.Println("Authentication successful")
			return nil
		})

}

// initUserProfile will auto fetch user info such as:
// user roles,user team.
func initUserProfile(ctx context.Context, client internal.CloudClient) (*internal.Role, error) {
	userInfo, _, err := client.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	return getUserRole(userInfo.Roles)
}

func GetUserInfo(ctx context.Context, client internal.CloudClient) (*internal.UserInfo, error) {
	userInfo, _, err := client.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

// / getUserRole returns the user role.
func getUserRole(roles []internal.Role) (*internal.Role, error) {
	roleList := []string{"org_admin", "team_admin", "team_member"}
	for _, role := range roles {
		contain := slices.Contains(roleList, role.Role)
		if contain {
			return &role, nil
		}
	}
	return nil, ErrorNoRoleFound
}

// initOrgInfo will fetch the user organization and extract team and create a controllerUrl that
// the user can use to connect to tyk cloud depending on their region.
func initOrgInfo(ctx context.Context, client internal.CloudClient, prompt internal.CloudPrompt, orgId string) (*internal.OrgInit, error) {
	info, _, err := client.GetOrgInfo(ctx, orgId)
	if err != nil {
		return nil, err
	}
	controllerUrl, err := util.GenerateUrlFromZone(info.Organisation.Zone)
	if err != nil {
		return nil, err
	}
	var orgInit internal.OrgInit
	orgInit.Controller = controllerUrl
	orgInit.Org = orgId
	selectedTeam, err := prompt.TeamPrompt(info.Organisation.Teams)
	if err != nil {
		return nil, err
	}
	if selectedTeam != nil {
		orgInit.Team = selectedTeam.UID
	}

	return &orgInit, nil
}

// addLoginFlags add the flags required by the login command.
func addLoginFlags(f *pflag.FlagSet) {
	f.StringP(email, "e", "", "email address you used to login into the dashboard")
	f.StringP(password, "p", "", "password you used to login into the dashboard")
	f.String(baUser, "", "Basic auth user.This should only be used for staging server")
	f.BoolP(interactive, "i", false, "login using the interactive mode.")
	err := viper.BindPFlag(internal.CreateKeyFromPath(cloudPath, baUser), f.Lookup(baUser))
	if err != nil {
		panic(err)
	}
	f.String(baPass, "", "Basic auth password")
	err = viper.BindPFlag(internal.CreateKeyFromPath(cloudPath, baPass), f.Lookup(baPass))
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
	baUser := viper.GetString(internal.CreateKeyFromPath(cloudPath, baUser))
	baPass := viper.GetString(internal.CreateKeyFromPath(cloudPath, baPass))
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
	email, err := internal.EmailPrompt()
	if err != nil {
		return nil, err
	}
	password, err := internal.PasswordPrompt()
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
	return internal.SaveToConfig(currentCloudToken, token)
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
