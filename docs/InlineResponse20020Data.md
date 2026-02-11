# InlineResponse20020Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | **string** | Price | 
**Type** | **string** | Buy/Sell order | 
**Amount** | **string** | Cryptocurrency amount | 
**MinAmount** | **string** | Minimum limit | 
**MaxAmount** | **string** | Maximum limit | 
**Total** | **string** | Fiat amount | 
**PayAli** | **int32** | Whether Alipay payment is supported | 
**PayBank** | **int32** | Whether bank payment is supported | 
**PayPaypal** | **int32** | Whether PayPal payment is supported | 
**PayWechat** | **int32** | Whether WeChat payment is supported | 
**PayTypeNum** | **string** | Payment method ID list | 
**PayTypeJson** | **string** | Payment method list | 
**LockedAmount** | **string** | Locked amount | 
**Orderid** | **int32** | Order ID | 
**Timestamp** | **int32** | Created time | 
**CurrencyType** | **string** | Cryptocurrency type | 
**WantType** | **string** | Fiat type | 
**HideRate** | **string** | Hidden price | 
**TradeTips** | **string** | Trading terms | 
**AutoReply** | **string** | Auto reply | 
**NewHand** | **string** | Merchant-friendly order | 
**RateRefId** | **int32** | Floating price reference ID: 1&#x3D;Platform reference price, 3&#x3D;Spot reference price (≤0 means fixed price, &gt;0 means floating price) | 
**RateOffset** | **float32** | Floating ratio (absolute value) | 
**Status** | **string** | Status | 
**RateFixed** | **int32** | 0&#x3D;Floating, 1&#x3D;Fixed | 
**FloatTrend** | **int32** | 0&#x3D;Upward float, 1&#x3D;Downward float | 
**ExpireMin** | **int32** | Timeout (minutes) | 
**TierLimit** | **int32** | Tier limit | 
**RegTimeLimit** | **int32** | Registration time limit | 
**AdvertisersLimit** | **int32** | Do not trade with advertisers, advertiser limit: 0&#x3D;No limit, 1&#x3D;Limit | 
**VerifiedLimit** | **int32** | kyclimit | 
**MinCompletedLimit** | **int32** | Minimum limit of completed orders | 
**MaxCompletedLimit** | **int32** | Maximum limit of completed orders | 
**UserOrdersLimit** | **int32** | Order count limit | 
**CompletedRateLimit** | **int32** | 30-day completion rate limit | 
**UserCountryLimit** | **int32** | KYC nationality restriction | 
**LimitCountryCn** | **string** | Restricted nationality (Chinese) | 
**LimitCountryEn** | **string** | Restricted nationality (English) | 
**IsHedge** | **int32** | Whether auto delegation | 
**HidePayment** | **int32** | Whether to hide payment method | 
**Fee** | **int32** | fee | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


