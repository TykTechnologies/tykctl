# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminCreateEnt**](EntitlementsApi.md#AdminCreateEnt) | **
Post** /admin/entitlement-plans | Create an entitlement plan as an admin
[**AdminEntitlementList**](EntitlementsApi.md#AdminEntitlementList) | **
Get** /admin/entitlement-plans | Get a list of all entitlement plans as an admin
[**AdminGetEnt**](EntitlementsApi.md#AdminGetEnt) | **
Get** /admin/entitlement-plans/{eid} | Fetch an entitlement plan object as an admin
[**ConsumeEntitlement**](EntitlementsApi.md#ConsumeEntitlement) | **
Post** /org/{oid}/entitlement | Consume/return an entitlement in an organisation
[**CreateEnt**](EntitlementsApi.md#CreateEnt) | **Post** /entitlement-plans | Fetch an entitlement plan object
[**DeleteEnt**](EntitlementsApi.md#DeleteEnt) | **Delete** /entitlement-plans/{eid} | Delete an entitlement plan object
[**EntitlementList**](EntitlementsApi.md#EntitlementList) | **
Get** /entitlement-plans | Get a list of all entitlement plans
[**GetEnt**](EntitlementsApi.md#GetEnt) | **Get** /entitlement-plans/{eid} | Fetch an entitlement plan object
[**UpdateEnt**](EntitlementsApi.md#UpdateEnt) | **Put** /entitlement-plans/{eid} | Fetch an entitlement plan object

# **AdminCreateEnt**

> InlineResponse201 AdminCreateEnt(ctx, body)
> Create an entitlement plan as an admin

Creates an entitlement plan as an admin (with admin token)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EntitlementPlan**](EntitlementPlan.md)| entitlement object |

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminEntitlementList**

> InlineResponse2002 AdminEntitlementList(ctx, )
> Get a list of all entitlement plans as an admin

Get a list of all entitlement plans as an admin (with admin token)

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGetEnt**

> InlineResponse201 AdminGetEnt(ctx, eid)
> Fetch an entitlement plan object as an admin

Fetches an entitlement plan as an admin (with admin token)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eid** | **string**| Plan ID |

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConsumeEntitlement**

> InlineResponse20015 ConsumeEntitlement(ctx, oid, code, optional)
> Consume/return an entitlement in an organisation

Consumes/returns 1 or more of a given entitlement within an organisation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**code** | **string**| Entitlement code to consume, e.g. MaxTeamMemberCount. |
**optional** | ***EntitlementsApiConsumeEntitlementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EntitlementsApiConsumeEntitlementOpts struct
Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------

**count** | **optional.Int32**| Number of this entitlement code to consume if positive, or return if negative. Optional,
defaults to 1. |

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEnt**

> InlineResponse201 CreateEnt(ctx, body)
> Fetch an entitlement plan object

Fetches an entitlement plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EntitlementPlan**](EntitlementPlan.md)| entitlement object |

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnt**

> InlineResponse201 DeleteEnt(ctx, eid)
> Delete an entitlement plan object

Deletes an entitlement plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eid** | **string**| Plan ID |

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EntitlementList**

> InlineResponse2002 EntitlementList(ctx, )
> Get a list of all entitlement plans

Get a list of all entitlement plans

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnt**

> InlineResponse201 GetEnt(ctx, eid)
> Fetch an entitlement plan object

Fetches an entitlement plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eid** | **string**| Plan ID |

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnt**

> InlineResponse201 UpdateEnt(ctx, body, eid)
> Fetch an entitlement plan object

Fetches an entitlement plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EntitlementPlan**](EntitlementPlan.md)| entitlement object |
**eid** | **string**| Plan ID |

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

