/*
 * Ara Platform
 *
 * Ara is a server platform that can deploy Tyk and it's various components into multiple control planes such as K8s and CloudFormation. It is modular and designed to be extended to support multiple specific operational deployments, DNS providers and underlying storage providers
 *
 * API version: v0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Loadout struct {
	// Blocked when true blocks the loadout's deployments from functioning depending on their kind
	Blocked bool `json:"Blocked,omitempty"`
	// The Deployment field should not be accessed directly, instead use the helper methods; GetDeployments, SetDeployment & AppendDeployment
	Deployments  []Deployment   `json:"Deployments,omitempty"`
	Entitlements *Entitlements  `json:"Entitlements,omitempty"`
	ExtraContext *MetaDataStore `json:"ExtraContext,omitempty"`
	// Frozen when true prevents non-admin write operations on the loadout
	Frozen      bool   `json:"Frozen,omitempty"`
	Name        string `json:"Name,omitempty"`
	OID         string `json:"OID,omitempty"`
	RuntimeKind string `json:"RuntimeKind,omitempty"`
	TID         string `json:"TID,omitempty"`
	// Cached name of the team this loadout is in
	TeamName string `json:"TeamName,omitempty"`
	UID      string `json:"UID,omitempty"`
}