# OrderHistoryListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** |  | [optional] 
**Symbol** | **string** |  | [optional] 
**Exchange** | **string** | Exchange, supports us, hk, and kr | [optional] 
**QuoteCurrency** | **string** |  | [optional] 
**FxRate** | **string** | Quote currency to USD exchange rate | [optional] 
**SymbolDesc** | **string** |  | [optional] 
**PriceType** | **string** | Price type (market &#x3D; market order, limit &#x3D; limit order) | [optional] 
**Status** | **int32** | Order status | [optional] 
**StatusDesc** | **string** | Order status description | [optional] 
**StatusDetail** | Pointer to [**OrderHistoryListItemStatusDetail**](OrderHistoryListItem_status_detail.md) |  | [optional] 
**FinishAs** | **int32** | Order completion reason | [optional] 
**Side** | **int32** | Side (1&#x3D;sell, 2&#x3D;buy) | [optional] 
**TimeInForce** | **string** | Time in force. - day: Day order. | [optional] 
**Volume** | **string** |  | [optional] 
**FillVolume** | **string** |  | [optional] 
**Price** | **string** |  | [optional] 
**AvgFillPrice** | Pointer to **string** |  | [optional] 
**Commission** | **string** | fee | [optional] 
**TimeSetup** | **int64** |  | [optional] 
**TimeDone** | **int64** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


