# FuturesInitialOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | **string** | Futures contract | 
**Size** | **int64** | Represents the number of contracts that need to be closed, full closing: size&#x3D;0 Partial closing: plan-close-short-position size&gt;0  Partial closing: plan-close-long-position size&lt;0 | [optional] 
**Price** | **string** | Order price. Set to 0 to use market price | 
**Close** | **bool** | In One-way Mode, when closing all positions, this must be set to true to perform the closing operation When partially closing positions in One-way Mode or Hedge Mode, you can omit close or set close&#x3D;false | [optional] [default to false]
**Tif** | **string** | Time in force strategy, default is gtc, market orders currently only support ioc mode  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled | [optional] [default to TIF_GTC]
**Text** | **string** | The source of the order, including: - web: Web - api: API call - app: Mobile app | [optional] 
**ReduceOnly** | **bool** | When set to true, perform automatic position reduction operation. Set to true to ensure that the order will not open a new position, and is only used to close or reduce positions | [optional] [default to false]
**AutoSize** | **string** | One-way Mode: auto_size is not required Hedge Mode full closing (size&#x3D;0): auto_size must be set, close_long for closing long positions, close_short for closing short positions Hedge Mode partial closing (size≠0): auto_size is not required | [optional] 
**IsReduceOnly** | **bool** | Is the order reduce-only | [optional] [readonly] 
**IsClose** | **bool** | Is the order to close position | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


