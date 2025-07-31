# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Futures contract | [optional] 
**Type** | **string** | Contract type: inverse - inverse contract, direct - direct contract | [optional] 
**QuantoMultiplier** | **string** | Multiplier used in converting from invoicing to settlement currency | [optional] 
**LeverageMin** | **string** | Minimum leverage | [optional] 
**LeverageMax** | **string** | Maximum leverage | [optional] 
**MaintenanceRate** | **string** | Maintenance rate of margin | [optional] 
**MarkType** | **string** | Mark price type: internal - internal trading price, index - external index price | [optional] 
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
**OrderSizeMin** | **int64** | Minimum order size allowed by the contract | [optional] 
**OrderSizeMax** | **int64** | Maximum order size allowed by the contract | [optional] 
**OrderPriceDeviate** | **string** | Maximum allowed deviation between order price and current mark price. The order price &#x60;order_price&#x60; must satisfy the following condition:      abs(order_price - mark_price) &lt;&#x3D; mark_price * order_price_deviate | [optional] 
**RefDiscountRate** | **string** | Trading fee discount for referred users | [optional] 
**RefRebateRate** | **string** | Commission rate for referrers | [optional] 
**OrderbookId** | **int64** | Orderbook update ID | [optional] 
**TradeId** | **int64** | Current trade ID | [optional] 
**TradeSize** | **int64** | Historical cumulative trading volume | [optional] 
**PositionSize** | **int64** | Current total long position size | [optional] 
**ConfigChangeTime** | **float64** | Last configuration update time | [optional] 
**InDelisting** | **bool** | &#x60;in_delisting&#x3D;true&#x60; and position_size&gt;0 indicates the contract is in delisting transition period &#x60;in_delisting&#x3D;true&#x60; and position_size&#x3D;0 indicates the contract is delisted | [optional] 
**OrdersLimit** | **int32** | Maximum number of pending orders | [optional] 
**EnableBonus** | **bool** | Whether bonus is enabled | [optional] 
**EnableCredit** | **bool** | Whether portfolio margin account is enabled | [optional] 
**CreateTime** | **float64** | Created time of the contract | [optional] 
**FundingCapRatio** | **string** | The factor for the maximum of the funding rate. Maximum of funding rate &#x3D; (1/market maximum leverage - maintenance margin rate) * funding_cap_ratio | [optional] 
**Status** | **string** | Contract status types include: prelaunch (pre-launch), trading (active), delisting (delisting), delisted (delisted) | [optional] 
**LaunchTime** | **int64** | Contract expiry timestamp | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


