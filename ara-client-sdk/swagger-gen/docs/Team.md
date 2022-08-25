# Team

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | **bool** | Blocked when true blocks the team&#x27;s deployments from functioning depending on their kind | [optional] [default to null]
**Entitlements** | [***Entitlements**](Entitlements.md) |  | [optional] [default to null]
**ExtraContext** | [***MetaDataStore**](MetaDataStore.md) |  | [optional] [default to null]
**Frozen** | **bool** | Frozen when true prevents non-admin write operations on the team | [optional] [default to null]
**Loadouts** | [**[]Loadout**](Loadout.md) | The team Loadouts (collections of deployments) The Loadouts field should not be accessed directly, instead use the helper methods; SetLoadout &amp; AppendLoadout | [optional] [default to null]
**Name** | **string** | The human-readable name of the team | [default to null]
**OID** | **string** | Organisation ID | [optional] [default to null]
**Organisation** | [***TeamOrganisation**](Team_Organisation.md) |  | [optional] [default to null]
**UID** | **string** | The unique ID of the team | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

