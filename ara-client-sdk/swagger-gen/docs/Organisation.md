# Organisation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | **string** | Account ID; the account this organisation belongs to | [optional] [default to null]
**Blocked** | **bool** | Blocked when true blocks the organisation&#x27;s deployments from functioning depending on their kind | [optional] [default to null]
**Entitlements** | [***Entitlements**](Entitlements.md) |  | [optional] [default to null]
**ExtraContext** | [***MetaDataStore**](MetaDataStore.md) |  | [optional] [default to null]
**Frozen** | **bool** | Frozen when true prevents non-admin write operations on the org | [optional] [default to null]
**Meta** | [**map[string]interface{}**](interface{}.md) | Meta data for the org, e.g. creation date - not required | [optional] [default to null]
**Name** | **string** | Human readable name | [default to null]
**PlanID** | **string** | Entitlement plan ID; if unset, will try and find a default or if only one present, will use that | [optional] [default to null]
**Teams** | [**[]Team**](Team.md) | Team list, create this with the Team creation API | [optional] [default to null]
**UID** | **string** | Organisation GUID | [optional] [default to null]
**Zone** | **string** | Zone identifier, implied by the controller&#x27;s zone | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

