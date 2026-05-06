# CrossexOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User ID | 
**OrderId** | **string** | Order ID | 
**Text** | **string** | Client-defined order ID. | 
**State** | **string** | 订单状态：  NEW：订单已通过校验，等待发送到交易所  OPEN：订单已挂在交易所订单簿上  PARTIALLY_FILLED：订单已部分成交  FILLED：订单已完全成交  FAIL：CrossEx 内部校验未通过，请查看 reason 字段了解失败原因  REJECT：订单被交易所拒绝，请查看 reason 字段了解失败原因 | 
**Symbol** | **string** | Unique trading pair identifiers, e.g. &#x60;BINANCE_SPOT_BTC_USDT&#x60;, &#x60;BINANCE_FUTURE_BTC_USDT&#x60;. | 
**Side** | **string** | Side (&#x60;BUY&#x60; buy / &#x60;SELL&#x60; sell). | 
**Type** | **string** | Order type (&#x60;LIMIT&#x60; limit / &#x60;MARKET&#x60; market). | 
**Attribute** | **string** | Order attributes (&#x60;COMMON&#x60; normal / &#x60;LIQ&#x60; liquidation takeover / &#x60;REDUCE&#x60; liquidation reduction / &#x60;ADL&#x60; auto-deleverage). | 
**ExchangeType** | **string** | Exchange type (&#x60;BINANCE&#x60; / &#x60;OKX&#x60; / &#x60;GATE&#x60; / &#x60;BYBIT&#x60;). | 
**BusinessType** | **string** | Business type (&#x60;SPOT&#x60; Spot / &#x60;FUTURE&#x60; Futures / &#x60;MARGIN&#x60; Margin). | 
**Qty** | **string** | Order quantity in the base currency. | 
**QuoteQty** | **string** | Order quantity in the quote currency. | 
**Price** | **string** | Order price. | 
**TimeInForce** | **string** | Time in force (default &#x60;GTC&#x60;; enum: &#x60;GTC&#x60; / &#x60;IOC&#x60; / &#x60;FOK&#x60; / &#x60;POC&#x60;). | 
**ExecutedQty** | **string** | Filled base amount. | 
**ExecutedAmount** | **string** | Filled quote amount. | 
**ExecutedAvgPrice** | **string** | Average Filled Price | 
**FeeCoin** | **string** | Fee currency | 
**Fee** | **string** | Fee amount. | 
**ReduceOnly** | **string** | Reduce-only order (&#x60;\&quot;true\&quot;&#x60; or &#x60;\&quot;false\&quot;&#x60;). | 
**Leverage** | **string** | Order leverage multiplier. | 
**Reason** | **string** | Failure reason description. | 
**LastExecutedQty** | **string** | Base quantity of the latest fill. | 
**LastExecutedPrice** | **string** | Price of the latest fill. | 
**LastExecutedAmount** | **string** | Quote amount of the latest fill. | 
**PositionSide** | **string** | Position side (&#x60;NONE&#x60; flat / &#x60;LONG&#x60; long / &#x60;SHORT&#x60; short). | 
**CreateTime** | **string** | Created time | 
**UpdateTime** | **string** | Update time | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


