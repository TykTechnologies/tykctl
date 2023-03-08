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

func AddGatewayServers(config *apim.Configuration) {
	urls := viper.GetStringSlice(gatewayUrls)
	for _, url := range urls {
		config.Servers = append(config.Servers, apim.ServerConfiguration{
			URL: url,
		})
	}
}
