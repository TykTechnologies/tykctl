# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminTeams**](TeamsApi.md#AdminTeams) | **Get** /admin/teams | Fetches a list of teams available in the system
[**CreateTeam**](TeamsApi.md#CreateTeam) | **Post** /org/{oid}/team | Create a team
[**DeleteTeam**](TeamsApi.md#DeleteTeam) | **Delete** /org/{oid}/team/{tid} | Delete a team
[**GetDeploymentsForTeams**](TeamsApi.md#GetDeploymentsForTeams) | **
Get** /org/{oid}/team/{tid}/deployments | Fetches a list of deployments for a team
[**GetTeam**](TeamsApi.md#GetTeam) | **Get** /org/{oid}/team/{tid} | Fetch a team
[**GetTeams**](TeamsApi.md#GetTeams) | **Get** /org/{oid}/team | Fetches a list of deployments for an organisation
[**UpdateTeam**](TeamsApi.md#UpdateTeam) | **Put** /org/{oid}/team/{tid} | Update a team

# **AdminTeams**

> InlineResponse2008 AdminTeams(ctx, optional)
> Fetches a list of teams available in the system

Fetches a list of teams in the system irrespective of organisations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***TeamsApiAdminTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsApiAdminTeamsOpts struct
Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.String**| Page number for pagination |
**perPage** | **optional.String**| Number of items to be shown per page |

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTeam**

> InlineResponse2011 CreateTeam(ctx, body, oid)
> Create a team

Creates a new team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Team**](Team.md)| team object |
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTeam**

> InlineResponse2011 DeleteTeam(ctx, oid, tid, optional)
> Delete a team

Deletes a team (add `cascade` option to de-provision)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |
**optional** | ***TeamsApiDeleteTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsApiDeleteTeamOpts struct
Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------

**cascade** | **optional.String**| set to any value to have the deployment set for this object de-provisioned |

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentsForTeams**

> InlineResponse200 GetDeploymentsForTeams(ctx, oid, tid)
> Fetches a list of deployments for a team

Fetches a list of deployments for a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeam**

> InlineResponse2011 GetTeam(ctx, oid, tid)
> Fetch a team

Fetches a single team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeams**

> InlineResponse20017 GetTeams(ctx, oid)
> Fetches a list of deployments for an organisation

Fetches a list of deployments for an organisation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeam**

> InlineResponse2011 UpdateTeam(ctx, body, oid, tid)
> Update a team

Updates a single team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Team**](Team.md)| team object |
**oid** | **string**| Organisation ID |
**tid** | **string**| Team ID |

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

