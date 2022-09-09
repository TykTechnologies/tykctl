/*
 * Ara Platform
 *
 * Ara is a server platform that can deploy Tyk and it's various components into multiple control planes such as K8s and CloudFormation. It is modular and designed to be extended to support multiple specific operational deployments, DNS providers and underlying storage providers
 *
 * API version: v0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Organisation struct {
	// Account ID; the account this organisation belongs to
	AccountID string `json:"AccountID,omitempty"`
	// Blocked when true blocks the organisation's deployments from functioning depending on their kind
	Blocked      bool           `json:"Blocked,omitempty"`
	Entitlements *Entitlements  `json:"Entitlements,omitempty"`
	ExtraContext *MetaDataStore `json:"ExtraContext,omitempty"`
	// Frozen when true prevents non-admin write operations on the org
	Frozen bool `json:"Frozen,omitempty"`
	// Meta data for the org, e.g. creation date - not required
	Meta map[string]interface{} `json:"Meta,omitempty"`
	// Human readable name
	Name string `json:"Name"`
	// Entitlement plan ID; if unset, will try and find a default or if only one present, will use that
	PlanID string `json:"PlanID,omitempty"`
	// Team list, create this with the Team creation API
	Teams []Team `json:"Teams,omitempty"`
	// Organisation GUID
	UID string `json:"UID,omitempty"`
	// Zone identifier, implied by the controller's zone
	Zone string `json:"Zone,omitempty"`
}
