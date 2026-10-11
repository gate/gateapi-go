# OrderHistoryListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | [optional] 
**Symbol** | **string** | Symbol | [optional] 
**Exchange** | **string** | Exchange, supports us, hk, kr, and jp | [optional] 
**QuoteCurrency** | **string** | Quote currency | [optional] 
**FxRate** | **string** | Quote currency to USD exchange rate | [optional] 
**SymbolDesc** | **string** | Symbol description | [optional] 
**PriceType** | **string** | Price type (market &#x3D; market order, limit &#x3D; limit order) | [optional] 
**Status** | **int32** | Order status. - 0: Unknown. - 1: New order; the order has been validated and is waiting to be processed. - 2: Partially filled. - 3: Filled. - 4: Canceled. - 5: Order rejected. - 6: Cancellation rejected. - 7: Order update succeeded. - 8: Order update rejected. | [optional] 
**StatusDesc** | **string** | Order status description | [optional] 
**StatusDetail** | Pointer to [**OrderHistoryListItemStatusDetail**](OrderHistoryListItem_status_detail.md) |  | [optional] 
**FinishAs** | **int32** | Order completion reason. - 0: Unknown. - 1: No quote found. - 2: No match found. - 3: Trigger price exceeded. - 4: Order value exceeds the limit. - 5: Unknown rejection. - 6: Unknown cancellation. - 7: Canceled by user. - 8: Canceled after the order time in force expired. - 9: Trading is not supported in the current session. - 13: Canceled due to corporate action. | [optional] 
**Side** | **int32** | Side (1&#x3D;sell, 2&#x3D;buy) | [optional] 
**TimeInForce** | **string** | Time in force. - day: Day order. | [optional] 
**Volume** | **string** | Order quantity | [optional] 
**FillVolume** | **string** | Trading size | [optional] 
**Price** | **string** | Order price | [optional] 
**AvgFillPrice** | Pointer to **string** | Average fill price | [optional] 
**Commission** | **string** | fee | [optional] 
**TimeSetup** | **int64** | Order creation time (Unix timestamp, seconds) | [optional] 
**TimeDone** | **int64** | Order completion time (Unix timestamp in seconds) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


