package shared

import (
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
		urls = append(urls, "http://localhost:8080")
	}

	for _, url := range urls {
		config.Servers = append(config.Servers, apim.ServerConfiguration{
			URL: url,
		})
	}

	return nil
}
