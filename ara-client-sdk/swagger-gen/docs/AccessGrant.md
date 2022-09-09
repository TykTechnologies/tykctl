# AccessGrant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | [**time.Time**](time.Time.md) | When the access grant expires | [optional] [default to null]
**Keys** | [**map[string]Secret**](Secret.md) | The map of tag-name to encrypted pem-encoded private key | [optional] [default to null]
**NeverExpire** | **bool** | Set to &#x60;true&#x60; to skip expiry checks | [optional] [default to null]
**UID** | **string** | THe ID of the access grant | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

