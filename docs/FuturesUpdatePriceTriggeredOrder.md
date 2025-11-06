# FuturesUpdatePriceTriggeredOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settle** | **string** | Settlement Currency (e.g., USDT, BTC) | [optional] [readonly] 
**OrderId** | **int32** | ID of the Pending Take-Profit/Stop-Loss Trigger Order | [optional] [readonly] 
**Contact** | **string** | The order ID of the modified price-triggered order. This ID is returned upon successful creation of the price-triggered order. Note: This ID must be passed in both the request path and request body. | [optional] 
**Size** | **int64** | Modified Contract Quantity. Full Close: 0; Partial Close: Positive/Negative values indicate direction (consistent with the creation interface logic). | [optional] 
**Price** | **string** | Represents the modified trading price. A value of 0 indicates a market order. | [optional] 
**TriggerPrice** | **string** | Modified Trigger Price | [optional] 
**PriceType** | **int32** | Reference price type. 0 - Latest trade price, 1 - Mark price, 2 - Index price | [optional] 
**AutoSize** | **string** | 单仓模式不需设置auto_size 双仓模式部分平仓(size≠0)时，不需设置auto_size 双仓模式全部平仓(size&#x3D;0)时，必须设置auto_size，close_long 平多头， close_short 平空头 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


