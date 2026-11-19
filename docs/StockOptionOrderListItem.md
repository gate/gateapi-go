# StockOptionOrderListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID. | [optional] 
**Symbol** | **string** | Full option contract symbol. | [optional] 
**SymbolDisplay** | **string** | Option contract name for display. | [optional] 
**Underlying** | **string** | Option underlying symbol. | [optional] 
**Side** | [**StockOptionOrderSide**](StockOptionOrderSide.md) |  | [optional] 
**OrderType** | [**StockOptionOrderType**](StockOptionOrderType.md) |  | [optional] 
**Status** | **int32** | Current order status: - 1: New, validated and awaiting processing. - 2: Partially filled, with a remaining unfilled quantity. | [optional] 
**StatusDesc** | **string** | Human-readable order status description in the requested language. | [optional] 
**TimeInForce** | [**StockOptionTimeInForce**](StockOptionTimeInForce.md) |  | [optional] 
**Quantity** | **string** | Order quantity. | [optional] 
**FilledQuantity** | **string** | Filled quantity. | [optional] 
**Price** | **string** | Order price. | [optional] 
**TimeSetup** | **int64** | Order creation time. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


