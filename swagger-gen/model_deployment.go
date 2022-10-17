/*
 * Ara Platform
 *
 * Ara is a server platform that can deploy Tyk and it's various components into multiple control planes such as K8s and CloudFormation. It is modular and designed to be extended to support multiple specific operational deployments, DNS providers and underlying storage providers
 *
 * API version: v0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"time"
)

// Once deployed it also holds all the required data about the live deployment.
type Deployment struct {
	// Access nonce
	AccessNonce string `json:"AccessNonce,omitempty"`
	// Blocked when true blocks the deployment from functioning depending on its kind
	Blocked bool `json:"Blocked,omitempty"`
	// Name of the channel (e.g. \"stable\", \"beta\") for bundle versions of this deployment
	BundleChannel string `json:"BundleChannel,omitempty"`
	// Bundle version applied to this deployment
	BundleVersion string `json:"BundleVersion,omitempty"`
	// Certificates to be deployed with the deployment
	Certificates []SecureCert `json:"Certificates,omitempty"`
	// Creation timestamp
	Created time.Time `json:"Created,omitempty"`
	// Deleted can be used to mark a deployment as deleted without removing the record.
	Deleted        bool           `json:"Deleted,omitempty"`
	Driver         string         `json:"Driver,omitempty"`
	DriverMetaData *Status        `json:"DriverMetaData,omitempty"`
	Entitlements   *Entitlements  `json:"Entitlements,omitempty"`
	ExtraContext   *MetaDataStore `json:"ExtraContext,omitempty"`
	// DNS names generated that are used to actually target the ingresses
	FriendlyNames map[string]string `json:"FriendlyNames,omitempty"`
	// Frozen when true prevents non-admin write operations on the deployment
	Frozen bool `json:"Frozen,omitempty"`
	// HasAWSSecrets denotes whether or not the ExtraContext metadata contains a key ID and secret with which to deal with plugins on AWS.
	HasAWSSecrets bool `json:"HasAWSSecrets,omitempty"`
	// IP or hostname based ingresses for various services, organised by tag
	Ingresses map[string]string `json:"Ingresses,omitempty"`
	// Public and private keys for crypto ops
	Keys []SecureCert `json:"Keys,omitempty"`
	Kind string       `json:"Kind"`
	// LID is the parent Loadout(/environment) ID
	LID string `json:"LID,omitempty"`
	// Last \"sync\" of data with the control plane
	LastChecked time.Time `json:"LastChecked,omitempty"`
	// Updated timestamp
	LastUpdate time.Time `json:"LastUpdate,omitempty"`
	// LinkedDeployments uses 'types.Linked*' consts as keys and values will be, for example, the UID of a parent home/CP deployment, for MDCB ingress.
	LinkedDeployments map[string]string `json:"LinkedDeployments,omitempty"`
	// Cached name of the loadout this deployment is in
	LoadoutName string `json:"LoadoutName,omitempty"`
	// Human readable name
	Name string `json:"Name"`
	// The namespace / slug of the deployment
	Namespace string `json:"Namespace,omitempty"`
	// Force deployment and don't roll back on failure (WARNING: for debug only!)
	NoRollback bool `json:"NoRollback,omitempty"`
	// OID is the great-grandparent Organisation ID.
	OID string `json:"OID,omitempty"`
	// If a federated or remote deployment, this should be set to the instance that has requested the deployment
	Origin string `json:"Origin,omitempty"`
	// RevisionID is essentially a deployment timestamp.
	RevisionID string `json:"RevisionID,omitempty"`
	// Revisions to store changes to the deployment over time, last entry is the newest
	RevisionRefs []string `json:"RevisionRefs,omitempty"`
	// State represents the current state of the deployment runner finite state machine.
	State string `json:"State,omitempty"`
	// TID is the grandparent Team ID.
	TID string `json:"TID,omitempty"`
	// Tags for sorting and listing
	Tags []string `json:"Tags,omitempty"`
	// Cached name of the team this deployment is in
	TeamName string `json:"TeamName,omitempty"`
	// Unique ID,shared by all objects to index outside of the storage engine
	UID string `json:"UID,omitempty"`
	// ZoneCode should identify the geographic location this deployment has been deployed to.
	ZoneCode string `json:"ZoneCode"`
}