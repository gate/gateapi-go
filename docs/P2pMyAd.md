# P2pMyAd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Ad side: &#x60;buy&#x60; buy-crypto ad; &#x60;sell&#x60; sell-crypto ad. | [optional] 
**Rate** | **string** | Price | [optional] 
**OriginalRate** | **string** | Original price | [optional] 
**Amount** | **string** | Remaining crypto amount on the ad. | [optional] 
**Total** | **string** | Remaining fiat amount of ad | [optional] 
**LimitTotal** | **string** | Single order limit range (cryptocurrency) | [optional] 
**LimitFiat** | **string** | Single order limit range (fiat) | [optional] 
**MinAmount** | **string** | Minimum quantity per order | [optional] 
**MaxAmount** | **string** | Maximum quantity per order | [optional] 
**PayTypeNum** | **string** | Payment method ID list | [optional] 
**PayTypeJson** | **string** | JSON map of payment type -&gt; payment method ID. | [optional] 
**ExpireMin** | **string** | Ad expiration time (minutes) | [optional] 
**TierLimit** | **string** | VIP limit | [optional] 
**AdvertisersLimit** | **int32** | Whether trading with the advertiser is restricted. &#x60;0&#x60;: no; &#x60;1&#x60;: yes. | [optional] 
**RegTimeLimit** | **int32** | Registration time limit | [optional] 
**VerifiedLimit** | **int32** | KYC level limit | [optional] 
**MinCompletedLimit** | **int32** | Minimum limit of completed orders by counterparty | [optional] 
**MaxCompletedLimit** | **int32** | Maximum limit of completed orders by counterparty | [optional] 
**UserCountryLimit** | **int32** | KYC nationality restriction | [optional] 
**CompletedRateLimit** | **float32** | 30-day completion rate limit | [optional] 
**UserOrdersLimit** | **int32** | Maximum order limit for counterparty | [optional] 
**HidePayment** | **string** | Whether payment methods are hidden. &#x60;1&#x60;: hidden; &#x60;0&#x60;: visible. | [optional] 
**CurrencyType** | **string** | Cryptocurrency symbol. | [optional] 
**WantType** | **string** | Fiat currency | [optional] 
**TradeTips** | **string** | Trading terms | [optional] 
**NewHand** | **int32** | Special ad type. &#x60;0&#x60; normal; &#x60;1&#x60; newcomer guide; &#x60;2&#x60; newcomer discount; &#x60;3&#x60; featured promo; &#x60;4&#x60; KOL ad; &#x60;5&#x60; coupon ad. | [optional] 
**Id** | **string** | Advertisement ID. | [optional] 
**Status** | **string** | Ad status: &#x60;OPEN&#x60; listed; &#x60;OFFLIN&#x60; delisted; &#x60;CLOSED&#x60; closed; &#x60;CANCEL&#x60; canceled. | [optional] 
**LockedAmount** | **string** | Ad frozen amount | [optional] 
**HideRate** | **string** | Hidden price | [optional] 
**IsOutTime** | **int32** | Whether the ad timed out. &#x60;1&#x60;: timed out; &#x60;0&#x60;: not yet. | [optional] 
**RateRefId** | **int32** | Floating reference: &#x60;1&#x60; platform; &#x60;2&#x60; Gate; &#x60;3&#x60; spot; &#x60;&lt;&#x3D; 0&#x60; means fixed price. | [optional] 
**RateOffset** | **string** | Floating ratio | [optional] 
**RateFixed** | **int32** | Price type: &#x60;0&#x60; floating; &#x60;1&#x60; fixed. | [optional] 
**FloatTrend** | **int32** | Floating direction: &#x60;0&#x60; markup; &#x60;1&#x60; markdown. | [optional] 
**InDispute** | **int32** | Whether the ad had a disputed trade. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**AutoReply** | **string** | Auto reply data | [optional] 
**Timestamp** | **int32** | Ad creation time | [optional] 
**IsHedge** | **int32** | Whether auto-delegation is enabled. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


