# OrderDealItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Trade record primary key ID | [optional] 
**DealId** | **int64** | Trade ID | [optional] 
**OrderId** | **int64** | Order ID | [optional] 
**ClientOrderId** | **string** | Client-defined order ID | [optional] 
**UserId** | **int32** | User ID | [optional] 
**Symbol** | **string** | Symbol | [optional] 
**Exchange** | **string** | Exchange. - us: US stocks. - hk: Hong Kong stocks. - kr: Korean stocks. - jp: Japanese stocks. | [optional] 
**Side** | **int32** | Side (1&#x3D;sell, 2&#x3D;buy) | [optional] 
**Role** | **string** | Liquidity role. - maker: Maker. - taker: Taker. | [optional] 
**Price** | **string** | Trade price in the instrument trading currency | [optional] 
**Quantity** | **string** | Trading size | [optional] 
**PositionQty** | **string** | Position quantity after the trade | [optional] 
**AvgEntryPrice** | **string** | Average entry price | [optional] 
**Value** | **string** | Trade value in the instrument trading currency | [optional] 
**FxRate** | **string** | Quote currency to USD exchange rate | [optional] 
**Fee** | **string** | Fee in the instrument trading currency | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


