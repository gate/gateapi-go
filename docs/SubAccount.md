# SubAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remark** | **string** | Remark | [optional] 
**LoginName** | **string** | 子账户登陆名：仅支持字母、数字、下划线，不可包含其他非法字符。 | 
**Password** | **string** | The sub-account&#39;s password. (Default: the same as main account&#39;s password) | [optional] 
**Email** | **string** | The sub-account&#39;s email address. (Default: the same as main account&#39;s email address) | [optional] 
**State** | **int32** | Sub-account status: 1-normal, 2-locked | [optional] [readonly] 
**Type** | **int32** | Sub-account type: 1-Regular sub-account, 3-Cross margin sub-account | [optional] [readonly] 
**UserId** | **int64** | Sub-account user ID | [optional] [readonly] 
**CreateTime** | **int64** | Created time | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


