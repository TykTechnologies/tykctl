/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/swagger-gen"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"net/mail"
)

const loginDesc = `
This command will login into your cloud account and set the token in your config file.

Note: The token will only last for 30 minute you will need to login again after 30 minutes.

You will be prompted to provide your email and  password to login.

When using the cloud service you should always run this command first as each command will require a token.

For the staging server you will also need to provide nginx basic auth.

Sample usage:
tykctl cloud login --ba-pass=<use this only is staging> --ba-pass=<use this in staging>

`

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:     "login",
	Short:   "Login into tyk cloud",
	Long:    loginDesc,
	Example: `tykctl cloud login --ba-pass=<use this only in staging> --ba-pass=<use this only in staging>`,
	RunE:    loginViaDashboard,
}

func init() {
	///cloudCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("email", "e", "", "User email for auth")
	loginCmd.Flags().StringP("password", "p", "", "User password for auth")
	loginCmd.Flags().String("ba-user", "", "Basic auth user")
	viper.BindPFlag("ba-user", loginCmd.Flags().Lookup("ba-user"))
	loginCmd.Flags().String("ba-pass", "", "Basic auth password")
	viper.BindPFlag("ba-pass", loginCmd.Flags().Lookup("ba-pass"))
	loginCmd.Flags().String("dashboard", "https://dash.ara-staging.tyk.technology", "Url to connect to the dashboard")
	viper.BindPFlag("dashboard", loginCmd.Flags().Lookup("dashboard"))

	//loginCmd.MarkFlagRequired("dashboard")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func loginViaDashboard(cmd *cobra.Command, args []string) error {
	dash := viper.GetString("dashboard")
	if len(dash) == 0 {
		cmd.Println("dashboard url is required")
		return errors.New("dashboard url is required")
	}
	url := fmt.Sprintf("%s/api/login", dash)
	//check if email was passed as an argument
	email, err := cmd.Flags().GetString("email")
	if err != nil {
		cmd.Println("error parsing the email")
		return nil
	}
	if len(email) == 0 {
		email, err = emailPrompt()
		if err != nil {
			cmd.Println(err)
			return err
		}
	}
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		cmd.Println("error parsing the password")
		return err
	}
	if len(password) == 0 {
		password, err = passwordPrompt()
		if err != nil {
			cmd.Println(err)
			return err
		}
	}
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    email,
		Password: password,
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	req, err := internal.CreatePostRequest(url, body, headers)
	if err != nil {
		cmd.Println(err)
		return err
	}
	username := viper.GetString("ba-user")

	userpass := viper.GetString("ba-pass")

	if len(username) > 0 && len(userpass) > 0 {
		req.SetBasicAuth(username, userpass)
	}
	log.Println(url)
	client := &http.Client{}
	s.Prefix = "login in "
	s.Start()
	resp, err := client.Do(req)
	if err != nil {
		cmd.Println(err)
		return err
	}
	s.Stop()
	if resp.StatusCode != 200 {
		cmd.PrintErrf("Login failed: %d\n", resp.StatusCode)
		if resp.Body != nil {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				message := err.Error()
				if myerr, ok := err.(swagger.GenericSwaggerError); ok {
					message = string(myerr.Body())
					// handle myerr
				}
				cmd.Println(message)
				return err
			}
			cmd.PrintErrf("Login failed: %s\n", string(b))
		}

		return errors.New("login failed")
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
		cmd.Println("Token was not found")
		return errors.New("no token found")
	}
	if signature == "" {
		cmd.PrintErrln("Signature not found")
		return errors.New("signature not found")
	}
	jwt := fmt.Sprintf("%s.%s", token, signature)
	viper.Set("token", jwt)
	if err := viper.WriteConfig(); err != nil {
		cmd.PrintErrf("Couldn't write config: %s\n", err.Error())
		return err
	}

	cmd.Println("Authentication successful")
	return nil
	///if email is empty then prompt for an email

}

func emailPrompt() (string, error) {
	prompt := promptui.Prompt{
		Label:    "Enter dashboard user email",
		Validate: validateEmail,
	}
	return prompt.Run()
}

func validateEmail(email string) error {
	if len(email) == 0 {
		return errors.New("cannot be empty string")
	}
	_, err := mail.ParseAddress(email)
	return err
}
func validateNotEmpty(email string) error {
	if len(email) == 0 {
		return errors.New("cannot be empty string")
	}
	return nil
}

func passwordPrompt() (string, error) {
	passPrompt := promptui.Prompt{
		Label:    "Enter dashboard user password",
		Validate: validateNotEmpty,
		Mask:     '*',
		//Mask:     '*',
	}

	return passPrompt.Run()
}
