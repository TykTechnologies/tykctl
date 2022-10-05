# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminLoadouts**](LoadoutsApi.md#AdminLoadouts) | **
Get** /admin/loadouts | Fetches a list of loadouts available in the system
[**AllocateRuntimeToLoadout**](LoadoutsApi.md#AllocateRuntimeToLoadout) | **
Post** /org/{oid}/team/{tid}/loadout/{lid}/runtime | Allocate a runtime to a loadout
[**CreateLoadout**](LoadoutsApi.md#CreateLoadout) | **Post** /org/{oid}/team/{tid}/loadout | Creates a loadout
[**DeleteLoadout**](LoadoutsApi.md#DeleteLoadout) | **Delete** /org/{oid}/team/{tid}/loadout/{lid} | Deletes a loadout
[**GetDeploymentsForLoadout**](LoadoutsApi.md#GetDeploymentsForLoadout) | **
Get** /org/{oid}/team/{tid}/loadout/{lid}/deployments | Fetches a list of deployments for a loadout
[**GetLoadout**](LoadoutsApi.md#GetLoadout) | **Get** /org/{oid}/team/{tid}/loadout/{lid} | Fetches a loadout
[**GetLoadouts**](LoadoutsApi.md#GetLoadouts) | **
Get** /org/{oid}/team/{tid}/loadout | Fetches a list of loadouts for a team
[**GetOrgLoadouts**](LoadoutsApi.md#GetOrgLoadouts) | **
Get** /org/{oid}/loadouts | Fetches a list of loadouts for an organisation
[**UpdateLoadout**](LoadoutsApi.md#UpdateLoadout) | **Put** /org/{oid}/team/{tid}/loadout/{lid} | Updates a loadout

# **AdminLoadouts**

> InlineResponse2003 AdminLoadouts(ctx, optional)
> Fetches a list of loadouts available in the system

Fetches a list of loadouts in the system irrespective of organisations and/or teams

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***LoadoutsApiAdminLoadoutsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoadoutsApiAdminLoadoutsOpts struct
Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.String**| Page number for pagination |
**perPage** | **optional.String**| Number of items to be shown per page |

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AllocateRuntimeToLoadout**

> InlineResponse2012 AllocateRuntimeToLoadout(ctx, oid, tid, lid, optional)
> Allocate a runtime to a loadout

Allocates a runtime entitlement to a single loadout from its parent organisation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |
**lid** | **string**| Loadout ID |
**optional** | ***LoadoutsApiAllocateRuntimeToLoadoutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoadoutsApiAllocateRuntimeToLoadoutOpts struct
Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------

**kind** | **optional.String**| Runtime entitlement kind to allocate, e.g. gp.small. If blank, deallocates the current
one, if any. |

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadout**

> InlineResponse2012 CreateLoadout(ctx, body, oid, tid)
> Creates a loadout

Creates a loadout

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Loadout**](Loadout.md)| loadout object |
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadout**

> InlineResponse2012 DeleteLoadout(ctx, oid, tid, lid, optional)
> Deletes a loadout

Deletes a loadout, add the `cascade` option to have the loadout de-provisioned

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |
**lid** | **string**| Loadout ID |
**optional** | ***LoadoutsApiDeleteLoadoutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoadoutsApiDeleteLoadoutOpts struct
Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------

**cascade** | **optional.String**| set to any value to have the deployment set for this object de-provisioned |

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentsForLoadout**

> InlineResponse200 GetDeploymentsForLoadout(ctx, oid, tid, lid)
> Fetches a list of deployments for a loadout

Fetches a list of deployments for a loadout

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |
**lid** | **string**| Loadout ID |

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadout**

> InlineResponse2012 GetLoadout(ctx, oid, tid, lid)
> Fetches a loadout

Fetches a loadout

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |
**lid** | **string**| Loadout ID |

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadouts**

> InlineResponse20016 GetLoadouts(ctx, oid, tid)
> Fetches a list of loadouts for a team

Fetches a list of loadouts for a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgLoadouts**

> InlineResponse20016 GetOrgLoadouts(ctx, oid)
> Fetches a list of loadouts for an organisation

Fetches a list of loadouts for an organisation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadout**

> InlineResponse2012 UpdateLoadout(ctx, body, oid, tid, lid)
> Updates a loadout

Updates a loadout

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Loadout**](Loadout.md)| loadout object |
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |
**lid** | **string**| Loadout ID |

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

