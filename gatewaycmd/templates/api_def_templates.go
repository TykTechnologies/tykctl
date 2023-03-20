package templates

import (
	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/util"
)

func createHTTpBinApiDefinition() *apim.APIDefinition {
	api := apim.APIDefinition{}
	api.Active = util.GetBoolPtr(true)
	api.Name = util.GetStrPtr("test1")
	api.ApiId = util.GetStrPtr("test1")
	api.UseKeyless = util.GetBoolPtr(true)
	api.Proxy = new(apim.APIDefinitionProxy)
	api.Proxy.ListenPath = util.GetStrPtr("/test1/")
	api.Proxy.TargetUrl = util.GetStrPtr("https://httpbin.org")
	api.Proxy.StripListenPath = util.GetBoolPtr(true)
	api.Proxy.PreserveHostHeader = util.GetBoolPtr(true)
	api.OrgId = util.GetStrPtr("default")
	api.VersionData = new(apim.APIDefinitionVersionData)
	api.VersionData.NotVersioned = util.GetBoolPtr(true)
	dt := map[string]apim.VersionInfo{
		"Default": {
			Name: util.GetStrPtr("Default"),
		},
	}
	api.VersionData.Versions = &dt
	return &api
}

func createFullApiDefinition() *apim.APIDefinition {
	api := apim.APIDefinition{}

	return &api
}
