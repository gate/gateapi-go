# ApiResponseExSkillGetUserIdentityResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | Business error code: 0 &#x3D; success, 1001 &#x3D; existing user, 1002 &#x3D; risk-controlled user, 1003 &#x3D; sub-account, 1004 &#x3D; agent, 1005 &#x3D; market maker, 1006 &#x3D; enterprise account, 1008 &#x3D; not logged in | [optional] 
**Label** | **string** | Error identifier code. Empty string on success, machine-readable error label on error | [optional] 
**Message** | **string** | Error description | [optional] 
**Data** | [**map[string]interface{}**](.md) | Empty object {} on success, null on failure | [optional] 
**Timestamp** | **int64** | Server timestamp (milliseconds) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


