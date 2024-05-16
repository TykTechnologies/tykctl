package sharedcmd

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
	"github.com/TykTechnologies/tykctl/configcmd"
	"github.com/TykTechnologies/tykctl/gatewaycmd"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"

	cc "github.com/ivanpirog/coloredcobra"
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
	currentCloudToken = "cloud.current_token"

	configDefault = "default"
)

func NewRootCmd() *cobra.Command {
	builder := internal.NewCmd(tykctl).
		WithLongDescription(rootDesc).
		WithDescription("Tykctl is a cli that can be used to interact with tyk components (tyk cloud,tyk gateway and tyk dashboard.").
		WithFlagAdder(true, addGlobalPersistentFlags)

	return builder.WithCommands()
}

func createConfigFiles() error {
	dir, err := internal.GetCoreDir()
	if err != nil {
		return err
	}

	configDir, err := internal.GetDefaultConfigDir()
	if err != nil {
		return err
	}

	err = util.CheckDirectory(configDir)
	if err != nil {
		return err
	}

	err = internal.CreateFile(dir, internal.CoreConfigFileName)
	if err != nil {
		return err
	}

	v, err := internal.CreateViper(dir, internal.CoreConfig)
	if err != nil {
		return err
	}

	activeConf := v.GetString(internal.CurrentConfig)
	if util.StringIsEmpty(activeConf) {
		activeConf = configDefault
		v.Set(internal.CurrentConfig, activeConf)
	}

	err = internal.CreateConfigFile(configDir, activeConf)
	if err != nil {
		return err
	}

	return v.WriteConfig()
}

func configCloud() internal.CloudFactory {
	conf := cloud.Configuration{
		DefaultHeader: map[string]string{},
	}

	sdkClient := internal.NewCloudSdkClient(&conf)
	cloudFactory := internal.CloudFactory{
		Client: sdkClient,
		Prompt: internal.NewSurveyPrompt(),
		Config: internal.ViperConfig{},
	}

	sdkClient.AddBeforeExecuteFunc(AddTokenAndBaseURL)
	sdkClient.AddBeforeRestyExecute(AddTokenAndBaseURLToResty)

	return cloudFactory
}

func configGateway() internal.ApimClient {
	apiConfig := apim.Configuration{
		DefaultHeader: make(map[string]string),
		Debug:         false,
		Servers:       apim.ServerConfigurations{},
	}
	client := apim.NewAPIClient(&apiConfig)
	apimClient := internal.ApimClient{Client: client}

	return apimClient
}

func Execute() {
	err := createConfigFiles()
	if err != nil {
		os.Exit(1)
	}

	cobra.OnInitialize(initConfig)

	rootCmd := NewRootCmd()

	cc.Init(&cc.Config{
		RootCmd:       rootCmd,
		Headings:      cc.HiCyan + cc.Bold + cc.Underline,
		Commands:      cc.HiYellow + cc.Bold,
		CmdShortDescr: cc.HiBlue,
		Example:       cc.HiGreen + cc.Italic,
		ExecName:      cc.Red + cc.Bold,
		Flags:         cc.HiMagenta + cc.Bold,
		FlagsDescr:    cc.HiBlue,
	})

	v, err := internal.CreateCoreViper()
	if err != nil {
		return
	}

	cloudFactory := configCloud()

	service := v.GetString(internal.CurrentService)
	switch service {
	case internal.Cloud:
		rootCmd.AddCommand(cloudcmd.CloudCommands(cloudFactory)...)
		rootCmd.AddCommand(cloudcmd.NewCloudCommand(cloudFactory))
	case internal.Gateway:
		apimClient := configGateway()
		rootCmd.AddCommand(gatewaycmd.GatewayCommands(apimClient)...)
		rootCmd.AddCommand(gatewaycmd.NewGatewayCommand(apimClient))
	default:
		rootCmd.AddCommand(cloudcmd.NewCloudCommand(cloudFactory))

		apimClient := configGateway()
		rootCmd.AddCommand(gatewaycmd.NewGatewayCommand(apimClient))
	}

	rootCmd.AddCommand(cloudcmd.NewCtxCmd())
	rootCmd.AddCommand(NewCheckoutCmd())

	configPrompt := configcmd.PickConfigPrompt{}
	fileConfig := internal.FileConfigEntry{}
	rootCmd.AddCommand(configcmd.NewConfigCmd(configPrompt, fileConfig, cloudFactory))

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addGlobalPersistentFlags(f *pflag.FlagSet) {
	f.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tykctl/config/config_default.yaml)")
}

// AddTokenAndBaseURL will add a user token from the configuration file to each request header.
func AddTokenAndBaseURL(client *cloud.APIClient, conf *cloud.Configuration) error {
	baseURL := viper.GetString(internal.CreateKeyFromPath("cloud", controller))
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
