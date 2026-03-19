# P2pAdDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | **string** | Price | [optional] 
**Type** | **string** | Buy/Sell order | [optional] 
**Amount** | **string** | Cryptocurrency amount | [optional] 
**MinAmount** | **string** | Minimum limit | [optional] 
**MaxAmount** | **string** | Maximum limit | [optional] 
**Total** | **string** | Fiat amount | [optional] 
**PayAli** | **int32** | Whether Alipay payment is supported | [optional] 
**PayBank** | **int32** | Whether bank payment is supported | [optional] 
**PayPaypal** | **int32** | Whether PayPal payment is supported | [optional] 
**PayWechat** | **int32** | Whether WeChat payment is supported | [optional] 
**PayTypeNum** | **string** | Payment method ID list | [optional] 
**PayTypeJson** | **string** | Payment method list | [optional] 
**LockedAmount** | **string** | Locked amount | [optional] 
**Orderid** | **int32** | Order ID | [optional] 
**Timestamp** | **int32** | Created time | [optional] 
**CurrencyType** | **string** | Cryptocurrency type | [optional] 
**WantType** | **string** | Fiat type | [optional] 
**HideRate** | **string** | Hidden price | [optional] 
**TradeTips** | **string** | Trading terms | [optional] 
**AutoReply** | **string** | Auto reply | [optional] 
**NewHand** | **string** | Merchant-friendly order | [optional] 
**RateRefId** | **int32** | Floating price reference ID: 1&#x3D;Platform reference price, 3&#x3D;Spot reference price (≤0 means fixed price, &gt;0 means floating price) | [optional] 
**RateOffset** | **float32** | Floating ratio (absolute value) | [optional] 
**Status** | **string** | Status | [optional] 
**RateFixed** | **int32** | 0&#x3D;Floating, 1&#x3D;Fixed | [optional] 
**FloatTrend** | **int32** | 0&#x3D;Upward float, 1&#x3D;Downward float | [optional] 
**ExpireMin** | **int32** | Timeout (minutes) | [optional] 
**TierLimit** | **int32** | Tier limit | [optional] 
**RegTimeLimit** | **int32** | Registration time limit | [optional] 
**AdvertisersLimit** | **int32** | Do not trade with advertisers, advertiser limit: 0&#x3D;No limit, 1&#x3D;Limit | [optional] 
**VerifiedLimit** | **int32** | kyclimit | [optional] 
**MinCompletedLimit** | **int32** | Minimum limit of completed orders | [optional] 
**MaxCompletedLimit** | **int32** | Maximum limit of completed orders | [optional] 
**UserOrdersLimit** | **int32** | Order count limit | [optional] 
**CompletedRateLimit** | **float32** | 30-day completion rate limit | [optional] 
**UserCountryLimit** | **int32** | KYC nationality restriction | [optional] 
**LimitCountryCn** | **string** | Restricted nationality (Chinese) | [optional] 
**LimitCountryEn** | **string** | Restricted nationality (English) | [optional] 
**IsHedge** | **int32** | Whether auto delegation | [optional] 
**HidePayment** | **int32** | Whether to hide payment method | [optional] 
**Fee** | **int32** | fee | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


