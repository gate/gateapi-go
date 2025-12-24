# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Futures contract | [optional] 
**Type** | **string** | Contract type: inverse - inverse contract, direct - direct contract | [optional] 
**QuantoMultiplier** | **string** | The contract multiplier indicates how many units of the underlying asset the face value of one contract represents. | [optional] 
**LeverageMin** | **string** | Minimum leverage | [optional] 
**LeverageMax** | **string** | Maximum leverage | [optional] 
**MaintenanceRate** | **string** | The maintenance margin rate of the first tier of risk limit sheet | [optional] 
**MarkType** | **string** | Deprecated | [optional] 
**MarkPrice** | **string** | Current mark price | [optional] 
**IndexPrice** | **string** | Current index price | [optional] 
**LastPrice** | **string** | Last trading price | [optional] 
**MakerFeeRate** | **string** | Maker fee rate, negative values indicate rebates | [optional] 
**TakerFeeRate** | **string** | Taker fee rate | [optional] 
**OrderPriceRound** | **string** | Minimum order price increment | [optional] 
**MarkPriceRound** | **string** | Minimum mark price increment | [optional] 
**FundingRate** | **string** | Current funding rate | [optional] 
**FundingInterval** | **int32** | Funding application interval, unit in seconds | [optional] 
**FundingNextApply** | **float64** | Next funding time | [optional] 
**RiskLimitBase** | **string** | Base risk limit (deprecated) | [optional] 
**RiskLimitStep** | **string** | Risk limit adjustment step (deprecated) | [optional] 
**RiskLimitMax** | **string** | Maximum risk limit allowed by the contract (deprecated). It is recommended to use /futures/{settle}/risk_limit_tiers to query risk limits | [optional] 
**OrderSizeMin** | **string** | Minimum order size allowed by the contract | [optional] 
**OrderSizeMax** | **string** | Maximum order size allowed by the contract | [optional] 
**OrderPriceDeviate** | **string** | Maximum allowed deviation between order price and current mark price. The order price &#x60;order_price&#x60; must satisfy the following condition:      abs(order_price - mark_price) &lt;&#x3D; mark_price * order_price_deviate | [optional] 
**RefDiscountRate** | **string** | Trading fee discount for referred users | [optional] 
**RefRebateRate** | **string** | Commission rate for referrers | [optional] 
**OrderbookId** | **int64** | Orderbook update ID | [optional] 
**TradeId** | **int64** | Current trade ID | [optional] 
**TradeSize** | **string** | Historical cumulative trading volume | [optional] 
**PositionSize** | **string** | Current total long position size | [optional] 
**ConfigChangeTime** | **float64** | Last configuration update time | [optional] 
**InDelisting** | **bool** | &#x60;in_delisting&#x3D;true&#x60; and position_size&gt;0 indicates the contract is in delisting transition period &#x60;in_delisting&#x3D;true&#x60; and position_size&#x3D;0 indicates the contract is delisted | [optional] 
**OrdersLimit** | **int32** | Maximum number of pending orders | [optional] 
**EnableBonus** | **bool** | Whether bonus is enabled | [optional] 
**EnableCredit** | **bool** | Whether portfolio margin account is enabled | [optional] 
**CreateTime** | **float64** | Created time of the contract | [optional] 
**FundingCapRatio** | **string** | Deprecated | [optional] 
**Status** | **string** | Contract status types include: prelaunch (pre-launch), trading (active), delisting (delisting), delisted (delisted), circuit_breaker (circuit breaker) | [optional] 
**LaunchTime** | **int64** | Contract expiry timestamp | [optional] 
**DelistingTime** | **int64** | Timestamp when contract enters reduce-only state | [optional] 
**DelistedTime** | **int64** | Contract delisting time | [optional] 
**MarketOrderSlipRatio** | **string** | The maximum slippage allowed for market orders, with the slippage rate calculated based on the latest market price | [optional] 
**MarketOrderSizeMax** | **string** | The maximum number of contracts supported for market orders, with a default value of 0. When the default value is used, the maximum number of contracts is limited by the &#x60;order_size_max&#x60; field | [optional] 
**FundingRateLimit** | **string** | Upper and lower limits of funding rate | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


