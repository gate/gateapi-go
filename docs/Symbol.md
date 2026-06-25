# Symbol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Unique trading pair identifier in the form ExchangeType_BusinessType_Base_Counter. | 
**ExchangeType** | **string** | Venue bucket (&#x60;BINANCE&#x60; / &#x60;OKX&#x60; / &#x60;GATE&#x60; / &#x60;BYBIT&#x60; / &#x60;KRAKEN&#x60; / &#x60;HYPERLIQUID&#x60;). | 
**BusinessType** | **string** | Business type (&#x60;SPOT&#x60; Spot / &#x60;FUTURE&#x60; Futures / &#x60;MARGIN&#x60; Margin). | 
**State** | **string** | Status (&#x60;live&#x60; running / &#x60;suspend&#x60; paused). | 
**MinSize** | **string** | Minimum order size allowed by the contract | 
**MinNotional** | **string** | Minimum Order Value | 
**LotSize** | **string** | Quantity Step | 
**TickSize** | **string** | Price Step | 
**MaxNumOrders** | **string** | maximumopen orderamount | 
**MaxMarketSize** | **string** | Maximum Market Order Quantity | 
**MaxLimitSize** | **string** | Maximum order quantity for limit orders. | 
**ContractSize** | **string** | Contract Multiplier | 
**LiquidationFee** | **string** | Liquidation Fee Rate | 
**DelistTime** | **string** | Millisecond timestamp; &#x60;0&#x60; means not delisted. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


