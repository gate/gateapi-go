# OptionsPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int32** | User ID | [optional] [readonly] 
**Underlying** | **string** | Underlying | [optional] [readonly] 
**UnderlyingPrice** | **string** | Underlying price (quote currency) | [optional] [readonly] 
**Contract** | **string** | Options contract name | [optional] [readonly] 
**Size** | **int64** | Position size (contract quantity) | [optional] [readonly] 
**EntryPrice** | **string** | Entry size (quote currency) | [optional] [readonly] 
**MarkPrice** | **string** | Current mark price (quote currency) | [optional] [readonly] 
**MarkIv** | **string** | Implied volatility | [optional] [readonly] 
**RealisedPnl** | **string** | Realized PnL | [optional] [readonly] 
**UnrealisedPnl** | **string** | Unrealized PNL | [optional] [readonly] 
**PendingOrders** | **int32** | Current pending order quantity | [optional] [readonly] 
**CloseOrder** | Pointer to [**OptionsPositionCloseOrder**](OptionsPosition_close_order.md) |  | [optional] 
**Delta** | **string** | Greek letter delta | [optional] [readonly] 
**Gamma** | **string** | Greek letter gamma | [optional] [readonly] 
**Vega** | **string** | Greek letter vega | [optional] [readonly] 
**Theta** | **string** | Greek letter theta | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


