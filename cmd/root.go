package cmd

import (
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

const rootDesc = `
Tykctl is a cli that can be used to interact with all tyk components (tyk cloud,tyk gateway and tyk dashboard).

The cli is grouped into services.
For example to use the tyk cloud services you should prefix all your subcommands with:
tykctl cloud <subcommand here>

Currently we only support tyk cloud.
`

func NewRootCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(tykctl).WithLongDescription(rootDesc).
		WithDescription("access all tyk service via the tykctl.").
		WithFlagAdder(true, addGlobalPersistentFlags).
		WithFlagAdder(false, addRootLocalFlags).
		WithCommands(NewCloudCommand(client), NewCtxCmd())
}

func Execute() {
	conf := cloud.Configuration{
		DefaultHeader: map[string]string{},
	}
	sdkClient := internal.NewCloudSdkClient(&conf)
	sdkClient.AddBeforeExecuteFunc(AddTokenAndBaseUrl)
	sdkClient.AddBeforeRestyExecute(AddTokenAndBaseUrlToResty)
	rootCmd := NewRootCmd(sdkClient)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addRootLocalFlags(f *pflag.FlagSet) {
	f.BoolP("toggle", "t", false, "Help message for toggle")
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

// AddTokenAndBaseUrl will add a user token from the configuration file to each request header.
func AddTokenAndBaseUrl(client *cloud.APIClient, conf *cloud.Configuration) error {
	baseUrl := viper.GetString(controller)
	client.ChangeBasePath(baseUrl)
	token := fmt.Sprintf("Bearer %s", viper.GetString(token))
	conf.AddDefaultHeader("Authorization", token)
	return nil

}

// AddTokenAndBaseUrlToResty will add token and BaseUrl to the resty client.
func AddTokenAndBaseUrlToResty(client *resty.Client) error {
	token := viper.GetString(token)
	client.SetBaseURL(internal.DashboardUrl)
	client.SetAuthToken(token)
	return nil
}
