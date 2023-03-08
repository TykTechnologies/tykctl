package sharedCmd

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/cloudcmd"
	"github.com/TykTechnologies/tykctl/gatewaycmd"
	"github.com/TykTechnologies/tykctl/internal"
)

var cfgFile string

const rootDesc = `
Tykctl is a cli that can be used to interact with tyk cloud.

How to use:

All the commands should be prefixed with cloud:

tykctl cloud <subcommand here>
`

var (
	controller        = "controller"
	tykctl            = "tykctl"
	defaultConfigFile = ".tykctl.yaml"
	currentCloudUser  = "cloud.current_user"
	currentCloudToken = "cloud.current_token"
)

func NewRootCmd() *cobra.Command {
	return internal.NewCmd(tykctl).
		WithLongDescription(rootDesc).
		WithDescription("Tykctl is a cli that can be used to interact with tyk components (tyk cloud,tyk gateway and tyk dashboard.").
		WithFlagAdder(true, addGlobalPersistentFlags).
		WithCommands()
}

func Execute() {
	conf := cloud.Configuration{
		DefaultHeader: map[string]string{},
	}
	sdkClient := internal.NewCloudSdkClient(&conf)

	sdkClient.AddBeforeExecuteFunc(AddTokenAndBaseURL)
	sdkClient.AddBeforeRestyExecute(AddTokenAndBaseURLToResty)

	rootCmd := NewRootCmd()
	cloudFactory := internal.CloudFactory{
		Client: sdkClient,
		Prompt: internal.NewSurveyPrompt(),
		Config: internal.ViperConfig{},
	}

	rootCmd.AddCommand(cloudcmd.NewCloudCommand(cloudFactory))

	apiConfig := apim.Configuration{
		DefaultHeader: make(map[string]string),
		Debug:         false,
		Servers:       apim.ServerConfigurations{},
	}
	client := apim.NewAPIClient(&apiConfig)
	apimClient := internal.ApimClient{Client: client}

	rootCmd.AddCommand(gatewaycmd.NewGatewayCommand(apimClient))
	rootCmd.AddCommand(cloudcmd.NewCtxCmd())
	rootCmd.AddCommand(NewCheckoutCmd())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addGlobalPersistentFlags(f *pflag.FlagSet) {
	f.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tykctl.yaml)")
}

func init() {
	file := defaultConfigFile

	home, err := os.UserHomeDir()
	if err != nil {
		cobra.CheckErr(err)
	}

	err = CreateConfigFile(home, file)
	cobra.CheckErr(err)
	cobra.OnInitialize(initConfig)
}

// AddTokenAndBaseURL will add a user token from the configuration file to each request header.
func AddTokenAndBaseURL(client *cloud.APIClient, conf *cloud.Configuration) error {
	baseURL := viper.GetString(internal.CreateKeyFromPath("cloud", viper.GetString(currentCloudUser), controller))
	client.ChangeBasePath(baseURL)

	token := fmt.Sprintf("Bearer %s", viper.GetString(currentCloudToken))
	conf.AddDefaultHeader("Authorization", token)

	return nil
}

// AddTokenAndBaseURLToResty will add token and BaseUrl to the resty client.
func AddTokenAndBaseURLToResty(client *resty.Client) error {
	token := viper.GetString(currentCloudToken)

	client.SetBaseURL(internal.DashboardURL)
	client.SetAuthToken(token)

	return nil
}
