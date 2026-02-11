# AccountKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **int32** | API key status: 1 - Normal, 2 - Locked, 3 - Frozen (can only be modified, default is 1 when creating)API Key Status: 1 - Normal, 2 - Locked, 3 - Frozen (can only be modified; default is 1 upon creation) | [optional] 
**Mode** | **int32** | User mode: 1 - Classic mode, 2 - Legacy unified mode (can only be specified when creating, cannot be modified) | [optional] 
**Name** | **[]string** | API Key Remark | [optional] 
**CurrencyPairs** | **[]string** | Trading Pair Whitelist, Maximum 30 Pairs | [optional] 
**UserId** | **int64** | User ID | [optional] 
**IpWhitelist** | **[]string** | IP Whitelist | [optional] 
**Perms** | [**[]AccountKeyInfoPerms**](AccountKeyInfo_perms.md) |  | [optional] 
**Key** | [**AccountDetailKey**](AccountDetail_key.md) |  | [optional] 
**CreatedAt** | **string** | Created time | [optional] [readonly] 
**UpdatedAt** | **string** | Last Update Time | [optional] [readonly] 
**LastAccess** | **string** | Last Access Time | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


