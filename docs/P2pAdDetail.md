# P2pAdDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | **string** | Advertisement price. | [optional] 
**Type** | **string** | Ad side: &#x60;buy&#x60; buy-crypto ad; &#x60;sell&#x60; sell-crypto ad. | [optional] 
**Amount** | **string** | Remaining crypto amount on the ad. | [optional] 
**MinAmount** | **string** | Minimum trade amount in &#x60;want_type&#x60;. | [optional] 
**MaxAmount** | **string** | Maximum trade amount priced in &#x60;want_type&#x60;. | [optional] 
**Total** | **string** | Fiat amount | [optional] 
**PayAli** | **int32** | Whether Alipay is supported. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**PayBank** | **int32** | Whether bank transfer is supported. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**PayPaypal** | **int32** | Whether PayPal is supported. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**PayWechat** | **int32** | Whether WeChat Pay is supported. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**PayTypeNum** | **string** | Payment method ID list | [optional] 
**PayTypeJson** | **string** | JSON map of payment type -&gt; payment method ID. | [optional] 
**LockedAmount** | **string** | Locked amount | [optional] 
**Orderid** | **int32** | Order ID | [optional] 
**Timestamp** | **int32** | Created time | [optional] 
**CurrencyType** | **string** | Cryptocurrency symbol. | [optional] 
**WantType** | **string** | Fiat type | [optional] 
**HideRate** | **string** | Hidden price | [optional] 
**TradeTips** | **string** | Trading terms | [optional] 
**AutoReply** | **string** | Auto reply | [optional] 
**RateRefId** | **int32** | Floating reference: &#x60;1&#x60; platform; &#x60;2&#x60; Gate; &#x60;3&#x60; spot; &#x60;&lt;&#x3D; 0&#x60; means fixed price. | [optional] 
**RateOffset** | **float32** | Floating ratio (absolute value) | [optional] 
**Status** | **string** | Ad status: &#x60;OPEN&#x60; listed; &#x60;OFFLIN&#x60; delisted; &#x60;CLOSED&#x60; closed; &#x60;CANCEL&#x60; canceled. | [optional] 
**RateFixed** | **int32** | Price type: &#x60;0&#x60; floating; &#x60;1&#x60; fixed. | [optional] 
**FloatTrend** | **int32** | Floating direction: &#x60;0&#x60; markup; &#x60;1&#x60; markdown. | [optional] 
**ExpireMin** | **int32** | Timeout (minutes) | [optional] 
**TierLimit** | **int32** | Tier limit | [optional] 
**RegTimeLimit** | **int32** | Registration time limit | [optional] 
**AdvertisersLimit** | **int32** | Whether trading with the advertiser is restricted. &#x60;0&#x60;: no; &#x60;1&#x60;: yes. | [optional] 
**MinCompletedLimit** | **int32** | Minimum limit of completed orders | [optional] 
**MaxCompletedLimit** | **int32** | Maximum limit of completed orders | [optional] 
**UserOrdersLimit** | **int32** | Order count limit | [optional] 
**CompletedRateLimit** | **float32** | 30-day completion rate limit | [optional] 
**LimitCountryCn** | **string** | Restricted nationality (Chinese) | [optional] 
**LimitCountryEn** | **string** | Restricted nationality (English) | [optional] 
**IsHedge** | **int32** | Whether auto-delegation is enabled. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**HidePayment** | **int32** | Whether payment methods are hidden. &#x60;1&#x60;: hidden; &#x60;0&#x60;: visible. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


