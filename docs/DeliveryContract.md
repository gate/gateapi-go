# DeliveryContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Futures contract | [optional] 
**Underlying** | **string** | Underlying | [optional] 
**Cycle** | **string** | Cycle type, e.g. WEEKLY, QUARTERLY | [optional] 
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
**BasisRate** | **string** | Fair basis rate | [optional] 
**BasisValue** | **string** | Fair basis value | [optional] 
**BasisImpactValue** | **string** | Funding used for calculating impact bid, ask price | [optional] 
**SettlePrice** | **string** | Settle price | [optional] 
**SettlePriceInterval** | **int32** | Settle price update interval | [optional] 
**SettlePriceDuration** | **int32** | Settle price update duration in seconds | [optional] 
**ExpireTime** | **int64** | Contract expiry timestamp | [optional] 
**RiskLimitBase** | **string** | Risk limit base | [optional] 
**RiskLimitStep** | **string** | Step of adjusting risk limit | [optional] 
**RiskLimitMax** | **string** | Maximum risk limit the contract allowed | [optional] 
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
**InDelisting** | **bool** | Contract is delisting | [optional] 
**OrdersLimit** | **int32** | Maximum number of pending orders | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


