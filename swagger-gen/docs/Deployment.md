# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessNonce** | **string** | Access nonce | [optional] [default to null]
**Blocked** | **
bool** | Blocked when true blocks the deployment from functioning depending on its kind | [optional] [default to null]
**BundleChannel** | **
string** | Name of the channel (e.g. \&quot;stable\&quot;, \&quot;beta\&quot;) for bundle versions of this deployment | [optional] [default to null]
**BundleVersion** | **string** | Bundle version applied to this deployment | [optional] [default to null]
**Certificates** | [**[]
SecureCert**](SecureCert.md) | Certificates to be deployed with the deployment | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) | Creation timestamp | [optional] [default to null]
**Deleted** | **
bool** | Deleted can be used to mark a deployment as deleted without removing the record. | [optional] [default to null]
**Driver** | **string** |  | [optional] [default to null]
**DriverMetaData** | [***Status**](Status.md) |  | [optional] [default to null]
**Entitlements** | [***Entitlements**](Entitlements.md) |  | [optional] [default to null]
**ExtraContext** | [***MetaDataStore**](MetaDataStore.md) |  | [optional] [default to null]
**FriendlyNames** | **map[string]
string** | DNS names generated that are used to actually target the ingresses | [optional] [default to null]
**Frozen** | **
bool** | Frozen when true prevents non-admin write operations on the deployment | [optional] [default to null]
**HasAWSSecrets** | **
bool** | HasAWSSecrets denotes whether or not the ExtraContext metadata contains a key ID and secret with which to deal with plugins on AWS. | [optional] [default to null]
**Ingresses** | **map[string]
string** | IP or hostname based ingresses for various services, organised by tag | [optional] [default to null]
**Keys** | [**[]SecureCert**](SecureCert.md) | Public and private keys for crypto ops | [optional] [default to null]
**Kind** | **string** |  | [default to null]
**LID** | **string** | LID is the parent Loadout(/environment) ID | [optional] [default to null]
**LastChecked** | [**
time.Time**](time.Time.md) | Last \&quot;sync\&quot; of data with the control plane | [optional] [default to null]
**LastUpdate** | [**time.Time**](time.Time.md) | Updated timestamp | [optional] [default to null]
**LinkedDeployments** | **map[string]
string** | LinkedDeployments uses &#x27;types.Linked*&#x27; consts as keys and values will be, for example, the UID of a parent home/CP deployment, for MDCB ingress. | [optional] [default to null]
**LoadoutName** | **string** | Cached name of the loadout this deployment is in | [optional] [default to null]
**Name** | **string** | Human readable name | [default to null]
**Namespace** | **string** | The namespace / slug of the deployment | [optional] [default to null]
**NoRollback** | **
bool** | Force deployment and don&#x27;t roll back on failure (WARNING: for debug only!) | [optional] [default to null]
**OID** | **string** | OID is the great-grandparent Organisation ID. | [optional] [default to null]
**Origin** | **
string** | If a federated or remote deployment, this should be set to the instance that has requested the deployment | [optional] [default to null]
**RevisionID** | **string** | RevisionID is essentially a deployment timestamp. | [optional] [default to null]
**RevisionRefs** | **[]
string** | Revisions to store changes to the deployment over time, last entry is the newest | [optional] [default to null]
**State** | **
string** | State represents the current state of the deployment runner finite state machine. | [optional] [default to null]
**TID** | **string** | TID is the grandparent Team ID. | [optional] [default to null]
**Tags** | **[]string** | Tags for sorting and listing | [optional] [default to null]
**TeamName** | **string** | Cached name of the team this deployment is in | [optional] [default to null]
**UID** | **
string** | Unique ID,shared by all objects to index outside of the storage engine | [optional] [default to null]
**ZoneCode** | **
string** | ZoneCode should identify the geographic location this deployment has been deployed to. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

