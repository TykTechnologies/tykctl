# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminGetOrgTotalUsage**](OrganisationsApi.md#AdminGetOrgTotalUsage) | **
Get** /admin/org/{oid}/total/usage/{kind}/{t0}/{t1} | Fetches throughput usage of org
[**AdminGetUsageByOrgOverRange**](OrganisationsApi.md#AdminGetUsageByOrgOverRange) | **
Get** /admin/org/usage/{kind}/{t0}/{t1} | Fetches usage totals for orgs, by org
[**AdminOrgLicenseUpdate**](OrganisationsApi.md#AdminOrgLicenseUpdate) | **
Patch** /admin/org/{oid}/license-update | Update licenses in this organisation&#x27;s deployments
[**AdminOrganisationBlock**](OrganisationsApi.md#AdminOrganisationBlock) | **
Put** /admin/org/{oid}/block | Block an organisation admin action
[**AdminOrganisationDelete**](OrganisationsApi.md#AdminOrganisationDelete) | **
Delete** /admin/org/{oid} | Delete an organisation via the admin endpoint
[**AdminOrganisationFreeze**](OrganisationsApi.md#AdminOrganisationFreeze) | **
Put** /admin/org/{oid}/freeze | Freeze an organisation admin action
[**AdminOrganisationRead**](OrganisationsApi.md#AdminOrganisationRead) | **
Get** /admin/org/{oid} | Fetches a single organisation available in the system
[**AdminOrganisationUnblock**](OrganisationsApi.md#AdminOrganisationUnblock) | **
Delete** /admin/org/{oid}/block | Unblock an organisation admin action
[**AdminOrganisationUnfreeze**](OrganisationsApi.md#AdminOrganisationUnfreeze) | **
Delete** /admin/org/{oid}/freeze | Unfreeze an organisation admin action
[**AdminOrganisationUpdate**](OrganisationsApi.md#AdminOrganisationUpdate) | **
Put** /admin/org/{oid} | Change an organisation object via the admin endpoint
[**AdminOrganisationsCreate**](OrganisationsApi.md#AdminOrganisationsCreate) | **
Post** /admin/orgs | Create an organisation via the admin endpoint
[**AdminOrganisationsRead**](OrganisationsApi.md#AdminOrganisationsRead) | **
Get** /admin/orgs | Fetches a list of organisations available in the system
[**CreateOrg**](OrganisationsApi.md#CreateOrg) | **Post** /org | Create a new organisation
[**DeleteOrg**](OrganisationsApi.md#DeleteOrg) | **Delete** /org/{oid} | Delete an org
[**GetDeploymentsForOrg**](OrganisationsApi.md#GetDeploymentsForOrg) | **
Get** /org/{oid}/deployments | Fetches a list of deployments for an organisation
[**GetOrg**](OrganisationsApi.md#GetOrg) | **Get** /org/{oid} | Fetch an org object
[**GetOrgTotalUsage**](OrganisationsApi.md#GetOrgTotalUsage) | **
Get** /org/{oid}/total/usage/{kind}/{t0}/{t1} | Fetches throughput usage of org
[**GetOrgs**](OrganisationsApi.md#GetOrgs) | **Get** /org | Fetch all organisations
[**GetUsageByOrgOverRange**](OrganisationsApi.md#GetUsageByOrgOverRange) | **
Get** /org/usage/{kind}/{t0}/{t1} | Fetches usage totals for orgs, by org
[**GetUsagePerDepForOrgOverRange**](OrganisationsApi.md#GetUsagePerDepForOrgOverRange) | **
Get** /org/{oid}/usage/{kind}/{t0}/{t1} | Fetches usage totals for an org, by deployment
[**UpdateOrg**](OrganisationsApi.md#UpdateOrg) | **Put** /org/{oid} | Update an organisation

# **AdminGetOrgTotalUsage**

> InlineResponse2006 AdminGetOrgTotalUsage(ctx, oid, kind, t0, t1)
> Fetches throughput usage of org

Fetches throughput usage of org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**kind** | **string**| Report kind |
**t0** | **int64**| start time |
**t1** | **int64**| end time |

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGetUsageByOrgOverRange**

> InlineResponse2004 AdminGetUsageByOrgOverRange(ctx, kind, t0, t1)
> Fetches usage totals for orgs, by org

Fetches usage totals for orgs, by org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string**| Report kind |
**t0** | **int64**| start time |
**t1** | **int64**| end time |

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrgLicenseUpdate**

> Payload AdminOrgLicenseUpdate(ctx, oid, time)
> Update licenses in this organisation's deployments

Update the licenses on all of this organisation's deployments to the new subscription period end timestamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**time** | **int64**| Unix timestamp for license end |

### Return type

[**Payload**](Payload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrganisationBlock**

> InlineResponse2005 AdminOrganisationBlock(ctx, oid)
> Block an organisation admin action

Sets a blocked flag on the organisation, which blocks its resources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrganisationDelete**

> InlineResponse2005 AdminOrganisationDelete(ctx, oid, optional)
> Delete an organisation via the admin endpoint

Delete an organisation via the admin endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**optional** | ***OrganisationsApiAdminOrganisationDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationsApiAdminOrganisationDeleteOpts struct
Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------

**cascade** | **optional.String**| set to any value to have the entire deployment set for this object de-provisioned |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrganisationFreeze**

> InlineResponse2005 AdminOrganisationFreeze(ctx, oid)
> Freeze an organisation admin action

Sets a freeze flag on the organisation, which prevents organisation's non-admin write actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrganisationRead**

> InlineResponse2005 AdminOrganisationRead(ctx, oid)
> Fetches a single organisation available in the system

Fetches a single organisation in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrganisationUnblock**

> InlineResponse2005 AdminOrganisationUnblock(ctx, oid)
> Unblock an organisation admin action

Unsets a block flag on the organisation, which unblocks its resources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrganisationUnfreeze**

> InlineResponse2005 AdminOrganisationUnfreeze(ctx, oid)
> Unfreeze an organisation admin action

Unsets a freeze flag on the organisation, which allows organisation's non-admin write actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrganisationUpdate**

> InlineResponse2005 AdminOrganisationUpdate(ctx, body, oid)
> Change an organisation object via the admin endpoint

Change an organisation object via the admin endpoint, requires the OrgUpdateProperties object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OrgUpdateProperties**](OrgUpdateProperties.md)| organisation update properties object |
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrganisationsCreate**

> InlineResponse2005 AdminOrganisationsCreate(ctx, body)
> Create an organisation via the admin endpoint

Create an organisation via the admin endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Organisation**](Organisation.md)| organisation object |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrganisationsRead**

> InlineResponse2007 AdminOrganisationsRead(ctx, optional)
> Fetches a list of organisations available in the system

Fetches a list of organisations in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***OrganisationsApiAdminOrganisationsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationsApiAdminOrganisationsReadOpts struct
Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Number of items to be shown per page |
**after** | **optional.String**| To set a starting point for pagination, only return objects after this given ID |

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrg**

> InlineResponse2005 CreateOrg(ctx, body)
> Create a new organisation

Organisations are a high-level wrapper object that represent a tenant in a specific geographic location (where
their `Home` deployments sit, orgs which require `Home` deployments in different locations require separate organisation
objects with each geographic controller)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Organisation**](Organisation.md)| organisation object |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrg**

> InlineResponse2005 DeleteOrg(ctx, oid, optional)
> Delete an org

Deletes an organisation object, will *not* de-provision unless `cascade` is set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**optional** | ***OrganisationsApiDeleteOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationsApiDeleteOrgOpts struct
Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------

**cascade** | **optional.String**| set to any value to have the entire deployment set for this object de-provisioned |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentsForOrg**

> InlineResponse200 GetDeploymentsForOrg(ctx, oid)
> Fetches a list of deployments for an organisation

Fetches a list of deployments for an organisation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrg**

> InlineResponse2005 GetOrg(ctx, oid)
> Fetch an org object

Fetches a entire org object including teams, loadouts and deployments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgTotalUsage**

> InlineResponse2006 GetOrgTotalUsage(ctx, oid, kind, t0, t1)
> Fetches throughput usage of org

Fetches throughput usage of org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**kind** | **string**| Report kind |
**t0** | **int64**| start time |
**t1** | **int64**| end time |

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgs**

> InlineResponse20014 GetOrgs(ctx, )
> Fetch all organisations

Fetches a list of Organisations objects (warning, this is a large object!)

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageByOrgOverRange**

> InlineResponse2004 GetUsageByOrgOverRange(ctx, kind, t0, t1)
> Fetches usage totals for orgs, by org

Fetches usage totals for orgs, by org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string**| Report kind |
**t0** | **int64**| start time |
**t1** | **int64**| end time |

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsagePerDepForOrgOverRange**

> InlineResponse20018 GetUsagePerDepForOrgOverRange(ctx, oid, kind, t0, t1)
> Fetches usage totals for an org, by deployment

Fetches usage totals for an org, by deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string**| Organisation ID |
**kind** | **string**| Report kind |
**t0** | **int64**| start time |
**t1** | **int64**| end time |

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrg**

> InlineResponse2005 UpdateOrg(ctx, body, oid)
> Update an organisation

Updates an organisation object, requires the OrgUpdateProperties object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OrgUpdateProperties**](OrgUpdateProperties.md)| organisation update properties object |
**oid** | **string**| Organisation ID |

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

