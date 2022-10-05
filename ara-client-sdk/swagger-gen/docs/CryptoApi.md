# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessGrant**](CryptoApi.md#CreateAccessGrant) | **Post** /sec/grant | Adds an access grant to the server
[**CreateChallenge**](CryptoApi.md#CreateChallenge) | **
Post** /sec/challenge | Adds a new accessor key set to the Access Grant list based on an existing key set (AG)
[**UpdateChallenge**](CryptoApi.md#UpdateChallenge) | **
Put** /sec/challenge/{id}/change | Updates the password set for an access grant

# **CreateAccessGrant**

> InlineResponse20015 CreateAccessGrant(ctx, body)
> Adds an access grant to the server

Access grants represent encrypted, password-protected cryptographic private keys used by the server and clients to
access sensitive information in objects. This endpoint should only b used to add a new private key for the *server* to
use, all user-based access should use the `challenges` endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AccessGrant**](AccessGrant.md)| Access grant object |

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateChallenge**

> InlineResponse20015 CreateChallenge(ctx, body)
> Adds a new accessor key set to the Access Grant list based on an existing key set (AG)

Adds an access grant to the server based off an existing AG (e.g. using the server's own AG) to bootstrap another access
client. Will decrypt the server key, and re-encrypt it against the challenges sent with this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Challenge**](Challenge.md)| Access grant object |

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChallenge**

> InlineResponse20015 UpdateChallenge(ctx, body, id)
> Updates the password set for an access grant

Use this endpoint to update the password set on an access grant, if the access grant is expired, this endpoint will
reset the expiry time as well.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ChallengeReset**](ChallengeReset.md)| Challenge reset objct |
**id** | **string**| Access Grant ID |

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

