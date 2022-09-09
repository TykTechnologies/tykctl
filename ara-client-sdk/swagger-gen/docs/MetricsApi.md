# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GuestGetMetric**](MetricsApi.md#GuestGetMetric) | **Get** /guest/{id}/metric/{kind} | Fetch a specific metric from the guest deployment
[**GuestGetMetricSum**](MetricsApi.md#GuestGetMetricSum) | **Get** /guest/{id}/metricsum | Fetch a sum of metrics from the guest deployment

# **GuestGetMetric**
> Payload GuestGetMetric(ctx, id, kind)
Fetch a specific metric from the guest deployment

Queries the guest deployment's driver for a specific metrics value, returning it

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Deployment ID | 
  **kind** | **string**| Metric kind | 

### Return type

[**Payload**](Payload.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GuestGetMetricSum**
> Payload GuestGetMetricSum(ctx, id, kinds)
Fetch a sum of metrics from the guest deployment

Queries the guest deployment's driver for a set of metrics in \"kinds\" param, returns the sum

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Deployment ID | 
  **kinds** | [**[]string**](string.md)| List of metric kinds | 

### Return type

[**Payload**](Payload.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

