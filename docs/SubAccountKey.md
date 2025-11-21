# SubAccountKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | User ID | [optional] [readonly] 
**Mode** | **int32** | Mode: 1 - classic 2 - portfolio account | [optional] 
**Name** | **string** | API Key Name | [optional] 
**Perms** | [**[]SubAccountKeyPerms**](SubAccountKey_perms.md) |  | [optional] 
**IpWhitelist** | **[]string** | IP whitelist (list will be cleared if no value is passed) | [optional] 
**Key** | **string** | API Key | [optional] [readonly] 
**State** | **int32** | Status: 1-Normal 2-Frozen 3-Locked | [optional] [readonly] 
**CreatedAt** | **int64** | Created time | [optional] [readonly] 
**UpdatedAt** | **int64** | Last Update Time | [optional] [readonly] 
**LastAccess** | **int64** | Last Access Time | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


