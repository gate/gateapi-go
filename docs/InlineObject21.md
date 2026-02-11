# InlineObject21

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Client-defined Order ID, supports letters (a-z), numbers (0-9), symbols (-, _) only | [optional] 
**Symbol** | **string** | Unique Identifier for Exchange_Business_Base_Counter Examples: - To place a spot order for ADA/USDT on Binance: Use unique identifier &#x60;BINANCE_SPOT_ADA_USDT&#x60;; - To place a USDT-margin perpetual contract order for ADA/USDT on OKX: Use unique identifier &#x60;OKX_FUTURE_ADA_USDT&#x60;; - To place a spot margin order for ADA/USDT on Gate.io: Use unique identifier &#x60;GATE_MARGIN_ADA_USDT&#x60;; Currently supports three order types: Spot Orders, USDT-margin Perpetual Contract Orders, and Spot Margin Orders | 
**Side** | **string** | BUY, SELL | 
**Type** | **string** | Order type (default: &#x60;LIMIT&#x60;; supported types: &#x60;LIMIT&#x60;, &#x60;MARKET&#x60;) | [optional] [default to TYPE_LIMIT]
**TimeInForce** | **string** | Default GTC, supports enumerated types: GTC, IOC, FOK, POC GTC: GoodTillCancelled IOC: ImmediateOrCancelled FOK: FillOrKill POC: PendingOrCancelled or PostOnly | [optional] [default to TIME_IN_FORCE_GTC]
**Qty** | **string** | Order quantity (required unless spot market buy) | [optional] 
**Price** | **string** | Limit Order Price (Required for Limit Orders) | [optional] 
**QuoteQty** | **string** | Order quote quantity; required for spot and margin market buy orders | [optional] 
**ReduceOnly** | **string** | Reduce-only: &#x60;true&#x60; or &#x60;false&#x60; | [optional] 
**PositionSide** | **string** | Position side: &#x60;NONE&#x60;, &#x60;LONG&#x60;, &#x60;SHORT&#x60; Defaults to &#x60;NONE&#x60; (single position mode) if not specified | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


