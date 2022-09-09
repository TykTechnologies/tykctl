/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ara-client-sdk/internal"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"net/mail"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: loginViaDashboard,
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("email", "e", "", "User email for auth")
	loginCmd.Flags().StringP("password", "p", "", "User password for auth")
	loginCmd.Flags().String("ba-user", "", "Basic auth user")
	loginCmd.Flags().String("ba-pass", "", "Basic auth password")
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
	cmd.Println(url)
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
	username, err := cmd.Flags().GetString("ba-user")
	if err != nil {
		cmd.Println(err)
		return err
	}
	userpass, err := cmd.Flags().GetString("ba-pass")
	if err != nil {
		cmd.Println(err)
		return err
	}
	if len(username) > 0 && len(userpass) > 0 {
		req.SetBasicAuth(username, userpass)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		cmd.Println(err)
		return err
	}
	if resp.StatusCode != 200 {
		cmd.PrintErrf("Login failed: %d\n", resp.StatusCode)
		if resp.Body != nil {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
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
		//Mask:     '*',
	}

	return passPrompt.Run()
}
