# PlaceBizPushOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyType** | **string** | Cryptocurrency symbol. | 
**ExchangeType** | **string** | Fiat currency | 
**Type** | **string** | Ad operation type. &#x60;0&#x60;: publish sell ad; &#x60;1&#x60;: publish buy ad; &#x60;2&#x60;: edit sell ad; &#x60;3&#x60;: edit buy ad. | 
**UnitPrice** | **string** | Per-unit price in fixed-price mode. | 
**Number** | **string** | Ad amount priced in &#x60;currencyType&#x60;. | 
**PayType** | **string** | Payment types enabled for the ad, comma-separated; values can be obtained from &#x60;pay_type&#x60; in the payment method list, e.g. &#x60;bank&#x60;, &#x60;alipay&#x60;, &#x60;wechat&#x60;, &#x60;paypal&#x60;, &#x60;swift&#x60;, &#x60;wu&#x60;. &#x60;pay_type_json&#x60; uses the types in this field as keys to specify the corresponding payment accounts. | 
**PayTypeJson** | **string** | JSON string of specific payment accounts corresponding to &#x60;payType&#x60;. Each key is a payment type listed in &#x60;payType&#x60;, and each value is the current user&#39;s payment method ID for that type. For example, when &#x60;payType&#x60; is &#x60;bank,swift&#x60;, this field can be {\&quot;bank\&quot;:\&quot;10001\&quot;,\&quot;swift\&quot;:\&quot;10002\&quot;}. | [optional] 
**RateFixed** | **string** | Price type: &#x60;0&#x60; floating; &#x60;1&#x60; fixed. | [optional] 
**Oid** | **string** | Pass ad ID when editing; omit or empty when publishing a new ad. | [optional] 
**MinAmount** | **string** | Minimum quantity per order, denominated by currencyType; required when limitBasis is not passed or is 0 | [optional] 
**MaxAmount** | **string** | Maximum quantity per order, denominated by currencyType; required when limitBasis is not passed or is 0 | [optional] 
**LimitBasis** | **int32** | Trading limit unit. 0: by crypto quantity, 1: by fiat amount; defaults to 0 when not passed for a new ad. The limit unit of an existing ad cannot be changed when editing; a fiat-limit ad must keep passing 1 when edited | [optional] 
**FiatMinAmount** | **string** | Minimum amount per order, denominated by exchangeType; required when limitBasis is 1 | [optional] 
**FiatMaxAmount** | **string** | Maximum amount per order, denominated by exchangeType; required when limitBasis is 1, and must not exceed the total fiat value of the ad quantity converted at the price | [optional] 
**TierLimit** | **string** | Minimum counterparty VIP level; &#x60;0&#x60; means no requirement. | [optional] 
**VerifiedLimit** | **string** | Minimum counterparty verification level; &#x60;0&#x60; means no limit. | [optional] 
**RegTimeLimit** | **string** | Minimum counterparty account age in days; &#x60;0&#x60; means no limit. | [optional] 
**AdvertisersLimit** | **string** | Whether trading with the advertiser is restricted. &#x60;0&#x60;: no; &#x60;1&#x60;: yes. | [optional] 
**PolymarketLimit** | **int32** | Whether to restrict trading with Polymarket users. 0: no restriction, 1: restricted | [optional] 
**ExpireMin** | **string** | Payment timeout in minutes. | [optional] 
**TradeTips** | **string** | Advertisement trade terms displayed to ordering users; goes through off-platform traffic diversion risk control on submission, and when hit, the advertisement is not saved and code 70305102 is returned | [optional] 
**AutoReply** | **string** | Auto reply content after order creation; goes through off-platform traffic diversion risk control on submission, and when hit, the advertisement is not saved and code 70305102 is returned | [optional] 
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


