package shared

import (
	"errors"

	"github.com/spf13/viper"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
)

const (
	xTykAuthorization = "x-tyk-authorization"
	gatewaySecret     = "gateway.secret"
	gatewayUrls       = "gateway.urls"
)

func AddGatewaySecret(config *apim.Configuration) {
	secret := viper.GetString(gatewaySecret)
	config.AddDefaultHeader(xTykAuthorization, secret)
}

func AddGatewayServers(config *apim.Configuration) error {
	urls := viper.GetStringSlice(gatewayUrls)
	if len(urls) == 0 {
		return errors.New("you have not set the url to connect to your gateway in your config")
	}

	for _, url := range urls {
		config.Servers = append(config.Servers, apim.ServerConfiguration{
			URL: url,
		})
	}

	return nil
}
