# PlaceBizPushOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyType** | **string** | Cryptocurrency symbol. | 
**ExchangeType** | **string** | Fiat currency | 
**Type** | **string** | Ad operation type. &#x60;0&#x60;: publish sell ad; &#x60;1&#x60;: publish buy ad; &#x60;2&#x60;: edit sell ad; &#x60;3&#x60;: edit buy ad. | 
**UnitPrice** | **string** | Per-unit price in fixed-price mode. | 
**Number** | **string** | Ad amount priced in &#x60;currencyType&#x60;. | 
**PayType** | **string** | Payment types, comma-separated; from pay type list &#x60;pay_type&#x60;, e.g. &#x60;bank&#x60;, &#x60;alipay&#x60;, &#x60;wechat&#x60;, &#x60;paypal&#x60;, &#x60;swift&#x60;, &#x60;wu&#x60;. | 
**PayTypeJson** | **string** | JSON map of payment type -&gt; user&#39;s payment method ID. | [optional] 
**RateFixed** | **string** | Price type: &#x60;0&#x60; floating; &#x60;1&#x60; fixed. | [optional] 
**Oid** | **string** | Pass ad ID when editing; omit or empty when publishing a new ad. | [optional] 
**MinAmount** | **string** | Minimum trade amount in &#x60;exchangeType&#x60;. | 
**MaxAmount** | **string** | Maximum amount per trade in &#x60;exchangeType&#x60; fiat units. | 
**TierLimit** | **string** | Minimum counterparty VIP level; &#x60;0&#x60; means no requirement. | [optional] 
**VerifiedLimit** | **string** | Minimum counterparty verification level; &#x60;0&#x60; means no limit. | [optional] 
**RegTimeLimit** | **string** | Minimum counterparty account age in days; &#x60;0&#x60; means no limit. | [optional] 
**AdvertisersLimit** | **string** | Whether trading with the advertiser is restricted. &#x60;0&#x60;: no; &#x60;1&#x60;: yes. | [optional] 
**ExpireMin** | **string** | Payment timeout in minutes. | [optional] 
**TradeTips** | **string** | Ad trading terms shown to the taker. | [optional] 
**AutoReply** | **string** | Auto-reply message after order creation. | [optional] 
**MinCompletedLimit** | **string** | Minimum completed orders for counterparty; &#x60;-1&#x60; unlimited. | [optional] 
**MaxCompletedLimit** | **string** | Maximum completed orders for counterparty; &#x60;-1&#x60; unlimited. | [optional] 
**CompletedRateLimit** | **string** | Counterparty minimum 30-day completion rate; &#x60;-1&#x60; means no limit. | [optional] 
**UserCountryLimit** | **string** | KYC nationality restriction; &#x60;-1&#x60; means no restriction. | [optional] 
**UserOrderLimit** | **string** | Maximum concurrent orders allowed for the counterparty. &#x60;-1&#x60;: unlimited. | [optional] 
**RateReferenceId** | **string** | Floating price reference. &#x60;1&#x60;: platform reference; &#x60;2&#x60;: Gate reference; &#x60;3&#x60;: spot reference. | [optional] 
**RateOffset** | **string** | Absolute floating offset ratio, e.g. &#x60;0.5&#x60; means 0.5%. | [optional] 
**FloatTrend** | **string** | Floating direction: &#x60;0&#x60; markup; &#x60;1&#x60; markdown. | [optional] 
**TeamPaymentUid** | **string** | Team payee UID; optional for non-team merchants. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


