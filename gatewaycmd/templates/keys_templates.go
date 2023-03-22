package templates

import (
	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/util"
)

func createLeanKeyTemplate() *apim.SessionState {
	sessionState := apim.NewSessionState()
	sessionState.Allowance = util.GetFloat64Ptr(1000)
	sessionState.Rate = util.GetFloat64Ptr(1000)
	sessionState.Per = util.GetFloat64Ptr(1)
	sessionState.Expires = util.GetInt64(-1)
	sessionState.QuotaMax = util.GetInt64(-1)
	sessionState.OrgId = util.GetStrPtr("1")
	sessionState.QuotaRenews = util.GetInt64(1449051461)
	sessionState.QuotaRemaining = util.GetInt64(-1)
	sessionState.QuotaRenewalRate = util.GetInt64(60)
	accessRight := map[string]apim.AccessDefinition{
		"{API-ID}": {
			AllowedUrls: nil,
			ApiId:       util.GetStrPtr("{API-ID}"),
			ApiName:     util.GetStrPtr("{API-NAME}"),
			Limit:       nil,
			Versions:    []string{"Default"},
		},
	}
	sessionState.AccessRights = &accessRight
	return sessionState
}
