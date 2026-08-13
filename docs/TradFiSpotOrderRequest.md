# TradFiSpotOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | **string** | Order quantity | 
**Symbol** | **string** | Symbol | 
**Side** | **int32** | Side (1&#x3D;sell, 2&#x3D;buy) | 
**PriceType** | **string** | Price type (market &#x3D; market order, limit &#x3D; limit order) | 
**TradingSession** | **string** | Trading session. Limit orders support only All, while market orders support only Regular. | 
**TimeInForce** | **string** | Time in force. - day: Day order. | 
**Price** | **string** | Order price, used for limit orders | [optional] 
**ClientOrderId** | **string** | Client-defined order ID | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


