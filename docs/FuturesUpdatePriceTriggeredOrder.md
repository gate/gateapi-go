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
**AutoSize** | **string** | One-way Mode: auto_size is not required Hedge Mode partial closing (size≠0): auto_size is not required Hedge Mode full closing (size&#x3D;0): auto_size must be set, close_long for closing long positions, close_short for closing short positions | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


