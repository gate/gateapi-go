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
**Status** | **int32** | Order status | [optional] 
**StatusDesc** | **string** | Order status description | [optional] 
**StatusDetail** | Pointer to [**OrderHistoryListItemStatusDetail**](OrderHistoryListItem_status_detail.md) |  | [optional] 
**FinishAs** | **int32** | Order completion reason | [optional] 
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


