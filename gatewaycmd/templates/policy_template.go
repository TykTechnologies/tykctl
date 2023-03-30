package templates

import (
	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/util"
)

func createLeanPolicy() *apim.Policy {
	policy := apim.NewPolicy()
	policy.OrgId = util.GetStrPtr("53ac07777cbb8c2d53000002")
	policy.Rate = util.GetFloat64Ptr(3)
	policy.Per = util.GetFloat64Ptr(1)
	policy.QuotaMax = util.GetInt64(1000)
	policy.QuotaRenewalRate = util.GetInt64(90000)
	accessRight := map[string]apim.AccessDefinition{
		"{API-ID}": {
			AllowedUrls: []apim.AccessSpec{},
			ApiId:       util.GetStrPtr("{API-ID}"),
			ApiName:     util.GetStrPtr("{API-NAME}"),
			Limit:       nil,
			Versions:    []string{"Default"},
		},
	}
	policy.AccessRights = &accessRight
	policy.Active = util.GetBoolPtr(true)
	policy.IsInactive = util.GetBoolPtr(false)
	policy.Tags = []string{}

	return policy
}
