# FixedTermLendOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Subscription record ID | [optional] 
**Business** | **int32** | Business type: 1 for regular, 2 for VIP | [optional] 
**OrderId** | **int64** | Order ID | [optional] 
**UserId** | **int64** | User ID | [optional] 
**Asset** | **string** | Currency | [optional] 
**ProductId** | **int32** | Product ID | [optional] 
**LockUpPeriod** | **int32** | Lock-up period (in days) | [optional] 
**Principal** | **string** | Subscription principal | [optional] 
**YearRate** | **string** | Annual interest rate | [optional] 
**ProductType** | **int32** | Product type: 1 for regular, 2 for VIP | [optional] 
**Interest** | **string** | Accrued interest | [optional] 
**Status** | **int32** | Order status: 1 for holding, 2 for redeemed, 3 for matured, 4 for settled | [optional] 
**ReinvestStatus** | **int32** | Auto-renewal status: 0 for disabled, 1 for enabled | [optional] 
**RedeemAccountType** | **int32** | Redemption payout account type: 1 for spot account | [optional] 
**OriginOrder** | **string** | Original order ID, linked to previous order IDs in auto-renewal scenarios | [optional] 
**RedeemType** | **int32** | Redemption type: 1 for early redemption, 2 for maturity redemption | [optional] 
**RedeemTime** | **string** | Redemption time | [optional] 
**FinishTime** | **string** | Expiration time | [optional] 
**CreateTime** | **string** | Created time | [optional] 
**YearRatePerent** | **string** | Annual interest rate percentage display value | [optional] 
**TotalYearRatePercent** | **string** | Comprehensive annualized yield percentage (including interest rate boost, rewards, etc.) | [optional] 
**TotalInterest** | **string** | Total earnings (including interest and bonus rewards) | [optional] 
**ProductInfo** | [**FixedTermProductInfo**](FixedTermProductInfo.md) |  | [optional] 
**BonusInfo** | [**FixedTermBonusInfo**](FixedTermBonusInfo.md) |  | [optional] 
**CouponInfo** | [**FixedTermCouponInfo**](FixedTermCouponInfo.md) |  | [optional] 
**RedeemAt** | **int32** | Redemption timestamp (in seconds) | [optional] 
**FinishAt** | **int32** | Expiration timestamp (in seconds) | [optional] 
**CreateAt** | **int32** | Creation timestamp (in seconds) | [optional] 
**Icon** | **string** | Currency icon URL | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


