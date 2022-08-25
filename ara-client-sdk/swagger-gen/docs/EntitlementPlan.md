# EntitlementPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | [**map[string]CounterEntitlement**](CounterEntitlement.md) | Counters represent a maximum of a class that can be consumed. | [optional] [default to null]
**Default_** | **bool** | Not used | [optional] [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | Last time the plan was changed | [optional] [default to null]
**Name** | **string** | Human-readable name of the plan | [optional] [default to null]
**Region** | [***RegionEntitlement**](RegionEntitlement.md) |  | [optional] [default to null]
**Runtimes** | [**map[string]RuntimeEntitlement**](RuntimeEntitlement.md) | Runtimes represent abstract resources consumed by deployed stacks  In Phase 1: Control Plane deployments consume runtimes at the organisation level Edge deployments consume runtimes at the loadout/environment level | [optional] [default to null]
**Toggles** | [**map[string]EnabledEntitlement**](EnabledEntitlement.md) | Toggles represent values that can be either on or off. | [optional] [default to null]
**UID** | **string** | ID of the plan | [optional] [default to null]
**Values** | [***map[string]interface{}**](map.md) |  | [optional] [default to null]
**Zone** | **string** | Zone tag, not stored, but populated implicitly | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

