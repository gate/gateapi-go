# StockOptionOrderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID. | [optional] 
**Symbol** | **string** | Full option contract symbol. | [optional] 
**SymbolDisplay** | **string** | Option contract name for display. | [optional] 
**Underlying** | **string** | Option underlying symbol. | [optional] 
**Side** | [**StockOptionOrderSide**](StockOptionOrderSide.md) |  | [optional] 
**OrderType** | [**StockOptionOrderType**](StockOptionOrderType.md) |  | [optional] 
**Status** | **int32** | Order status in order details. The details endpoint does not filter by status. - 1: New, validated and awaiting processing. - 2: Partially filled, with a remaining unfilled quantity. - 3: Filled, terminal. - 4: Cancelled, terminal; the order may have been partially filled. - 5: Rejected, terminal. Current orders include only 1 and 2; order history includes only 3, 4 and 5. | [optional] 
**StatusDesc** | **string** | Human-readable order status description in the requested language. | [optional] 
**TimeInForce** | [**StockOptionTimeInForce**](StockOptionTimeInForce.md) |  | [optional] 
**Quantity** | **string** | Order quantity. | [optional] 
**FilledQuantity** | **string** | Filled quantity. | [optional] 
**Price** | **string** | Order price. | [optional] 
**TimeSetup** | **int64** | Order creation time. | [optional] 
**AvgFillPrice** | **string** | Average fill price. | [optional] 
**Fee** | **string** | Trading fee. | [optional] 
**TimeDone** | **int64** | Order completion time. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


