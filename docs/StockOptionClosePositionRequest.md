# StockOptionClosePositionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Option contract symbol to close, with an available position. | 
**Quantity** | **string** | Quantity to close as a positive decimal string. Must not exceed the available position and must satisfy the contract&#39;s minimum, maximum and quantity increment rules. | 
**OrderType** | [**StockOptionOrderType**](StockOptionOrderType.md) |  | [optional] 
**Price** | **string** | Positive limit closing price as a plain decimal string, up to 64 characters, without scientific notation. Ignored for market closing orders. | [optional] 
**ClientOrderId** | **string** | Custom closing order identifier. Generated automatically when omitted or whitespace-only. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


