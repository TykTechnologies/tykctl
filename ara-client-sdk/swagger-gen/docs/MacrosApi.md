# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteMacro**](MacrosApi.md#ExecuteMacro) | **Post** /macro/{name} | runs a macro operation on a set of deployments

# **ExecuteMacro**
> Payload ExecuteMacro(ctx, body, name)
runs a macro operation on a set of deployments

runs a macro operation on a set of deployments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**map[string]string**](map.md)| options | 
  **name** | **string**| macro name | 

### Return type

[**Payload**](Payload.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

