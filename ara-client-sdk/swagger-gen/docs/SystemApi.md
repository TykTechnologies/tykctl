# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminZones**](SystemApi.md#AdminZones) | **Get** /admin/zones | Fetch all visible zones
[**AllManifests**](SystemApi.md#AllManifests) | **Get** /bundle/manifests | Fetch all manifest objects
[**GetBundle**](SystemApi.md#GetBundle) | **Get** /bundle/{kind}/channel/{channel}/driver/{driver}/version/{version} | Fetch a versioned bundle object
[**GetManifest**](SystemApi.md#GetManifest) | **Get** /bundle/{kind}/channel/{channel}/driver/{driver}/manifest | Fetch a manifest object

# **AdminZones**
> InlineResponse2009 AdminZones(ctx, optional)
Fetch all visible zones

Will return a list of available zones and their capabilities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemApiAdminZonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemApiAdminZonesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| control_discovery.filter, determines which zones AsZoneList is going to return | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AllManifests**
> InlineResponse20010 AllManifests(ctx, )
Fetch all manifest objects

Fetches a list of entire manifest objects registered, including all the bundles in it.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundle**
> InlineResponse20012 GetBundle(ctx, kind, channel, driver, version)
Fetch a versioned bundle object

Fetches a bundle object of a given version with its components.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kind** | **string**| Deployment Kind | 
  **channel** | **string**| Manifest Channel | 
  **driver** | **string**| Manifest Driver | 
  **version** | **string**| Bundle Version (pass \&quot;latest\&quot; for the latest one) | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetManifest**
> InlineResponse20011 GetManifest(ctx, kind, channel, driver)
Fetch a manifest object

Fetches an entire manifest object including all the bundles in it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kind** | **string**| Deployment Kind | 
  **channel** | **string**| Manifest Channel | 
  **driver** | **string**| Manifest Driver | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

