# FuturesTrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Fill ID | [optional] 
**CreateTime** | **float64** | Fill Time | [optional] 
**CreateTimeMs** | **float64** | Trade time, with millisecond precision to 3 decimal places | [optional] 
**Contract** | **string** | Futures contract | [optional] 
**Size** | **int64** | Trading size | [optional] 
**Price** | **string** | Trade price (quote currency) | [optional] 
**IsInternal** | **bool** | Whether it is an internal trade. Internal trade refers to the takeover of liquidation orders by the insurance fund and ADL users. Since it is not a normal matching on the market depth, the trade price may deviate from the market, and it will not be recorded in the K-line. If it is not an internal trade, this field will not be returned | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


