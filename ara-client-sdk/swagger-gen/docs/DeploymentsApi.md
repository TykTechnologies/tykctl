# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminDeploy**](DeploymentsApi.md#AdminDeploy) | **Post** /admin/deployments/{id}/deploy | Deploys any deployment by ID
[**AdminDeployments**](DeploymentsApi.md#AdminDeployments) | **Get** /admin/deployments | Fetches a list of deployments available in the system
[**AdminDeploymentsUpdate**](DeploymentsApi.md#AdminDeploymentsUpdate) | **Put** /admin/deployments | Update a deployment via the admin endpoint
[**AdminFetchDeployment**](DeploymentsApi.md#AdminFetchDeployment) | **Get** /admin/deployments/{id} | Fetches any deployment by ID
[**AdminFetchLinkedEdgeDeployments**](DeploymentsApi.md#AdminFetchLinkedEdgeDeployments) | **Get** /admin/deployments/{id}/linkededges | Fetches an Edge deployment linked to a CP
[**CreateDeployment**](DeploymentsApi.md#CreateDeployment) | **Post** /org/{oid}/team/{tid}/loadout/{lid}/deployment | Create a new deployment
[**DestroyDeployment**](DeploymentsApi.md#DestroyDeployment) | **Delete** /org/{oid}/team/{tid}/loadout/{lid}/deployment/{id} | De-provisions a deployment
[**DestroyGuestDeployment**](DeploymentsApi.md#DestroyGuestDeployment) | **Delete** /guest/deploy | Destroy a guest deployment
[**GetDeployment**](DeploymentsApi.md#GetDeployment) | **Get** /org/{oid}/team/{tid}/loadout/{lid}/deployment/{id} | Fetch a deployment
[**RestartDeployment**](DeploymentsApi.md#RestartDeployment) | **Put** /org/{oid}/team/{tid}/loadout/{lid}/deployment/{id}/restart | Restart a deployment
[**StartDeployment**](DeploymentsApi.md#StartDeployment) | **Post** /org/{oid}/team/{tid}/loadout/{lid}/deployment/{id}/deploy | Run a deployment (attempt to make it live)
[**StartGuestDeployment**](DeploymentsApi.md#StartGuestDeployment) | **Post** /guest/deploy | Run a guest deployment (attempt to make it live)
[**UpdateDeployment**](DeploymentsApi.md#UpdateDeployment) | **Put** /org/{oid}/team/{tid}/loadout/{lid}/deployment/{id} | Update a deployment

# **AdminDeploy**
> InlineResponse2001 AdminDeploy(ctx, id)
Deploys any deployment by ID

Deploys any deployment by ID with admin auth

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Deployment ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDeployments**
> InlineResponse200 AdminDeployments(ctx, optional)
Fetches a list of deployments available in the system

Fetches a list of deployments in the system irrespective of organisations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeploymentsApiAdminDeploymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiAdminDeploymentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**| Page number for pagination | 
 **perPage** | **optional.String**| Number of items to be shown per page | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDeploymentsUpdate**
> InlineResponse200 AdminDeploymentsUpdate(ctx, )
Update a deployment via the admin endpoint

Update a deployment via the admin endpoint

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminFetchDeployment**
> InlineResponse2001 AdminFetchDeployment(ctx, id)
Fetches any deployment by ID

Fetches any deployment by ID with admin auth

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Deployment ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminFetchLinkedEdgeDeployments**
> InlineResponse2001 AdminFetchLinkedEdgeDeployments(ctx, id, optional)
Fetches an Edge deployment linked to a CP

Fetches the first Edge deployment that is linked to the CP with the given UID, filtered by the given state, via admin auth

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Deployment ID of the parent CP | 
 **optional** | ***DeploymentsApiAdminFetchLinkedEdgeDeploymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiAdminFetchLinkedEdgeDeploymentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**| Filter for linked Edge(s) currently in this state | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDeployment**
> InlineResponse2001 CreateDeployment(ctx, body, oid, tid, lid)
Create a new deployment

Create a new deployment, will not deploy it to a control plane, but can be deployed once created. Requires at least a `Name`, `Zones` and `Kind` in the deployment object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Deployment**](Deployment.md)| deployment object | 
  **oid** | **string**| Organisation ID | 
  **tid** | **string**| Team ID | 
  **lid** | **string**| Loadout ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DestroyDeployment**
> InlineResponse2001 DestroyDeployment(ctx, oid, tid, lid, id, optional)
De-provisions a deployment

If the deployment is live, then the deployment will be de-provisioned, also optionally deletes it

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oid** | **string**| Organisation ID | 
  **tid** | **string**| Team ID | 
  **lid** | **string**| Loadout ID | 
  **id** | **string**| Deployment ID | 
 **optional** | ***DeploymentsApiDestroyDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiDestroyDeploymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **delete** | **optional.Bool**| mark deployment as deleted | 
 **purge** | **optional.Bool**| purge deployment from storage | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DestroyGuestDeployment**
> InlineResponse2001 DestroyGuestDeployment(ctx, body)
Destroy a guest deployment

Will stop a deployment operation for a deployment object as a guest.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Deployment**](Deployment.md)| deployment object | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployment**
> InlineResponse2001 GetDeployment(ctx, oid, tid, lid, id, optional)
Fetch a deployment

Update a deployment, will not deploy it to a control plane, call `/deploy` to apply the changes to the installation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oid** | **string**| Organisation ID | 
  **tid** | **string**| Team ID | 
  **lid** | **string**| Loadout ID | 
  **id** | **string**| Deployment ID | 
 **optional** | ***DeploymentsApiGetDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiGetDeploymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **access** | **optional.Bool**| Access nonce | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartDeployment**
> InlineResponse2001 RestartDeployment(ctx, oid, tid, lid, id)
Restart a deployment

Restarts a provisioned deployment's application resources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oid** | **string**| Organisation ID | 
  **tid** | **string**| Team ID | 
  **lid** | **string**| Loadout ID | 
  **id** | **string**| Deployment ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartDeployment**
> InlineResponse2001 StartDeployment(ctx, oid, tid, lid, id)
Run a deployment (attempt to make it live)

Will kick off a deployment operation for a specific deployment ID, this request returns almost immediately, for status information poll the deployment object for a change in the 'Status' field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oid** | **string**| Organisation ID | 
  **tid** | **string**| Team ID | 
  **lid** | **string**| Loadout ID | 
  **id** | **string**| Deployment ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartGuestDeployment**
> InlineResponse2001 StartGuestDeployment(ctx, body)
Run a guest deployment (attempt to make it live)

Will run a deployment operation for a deployment object as a guest, this request returns almost immediately, for status information poll the deployment object for a change in the 'Status' field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Deployment**](Deployment.md)| deployment object | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeployment**
> InlineResponse2001 UpdateDeployment(ctx, body, oid, tid, lid, id, optional)
Update a deployment

Fetch a deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Deployment**](Deployment.md)| A basic deployment object, use the &#x60;ExtraContext&#x60; field annotations to add options | 
  **oid** | **string**| Organisation ID | 
  **tid** | **string**| Team ID | 
  **lid** | **string**| Loadout ID | 
  **id** | **string**| Deployment ID | 
 **optional** | ***DeploymentsApiUpdateDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiUpdateDeploymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **retainSecrets** | **optional.**| preserve any previously specified AWS secrets | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

