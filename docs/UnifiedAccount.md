# UnifiedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | User ID | [optional] 
**RefreshTime** | **int64** | Last refresh time | [optional] 
**Locked** | **bool** | Whether the account is locked, valid in cross-currency margin/combined margin mode, false in other modes such as single-currency margin mode | [optional] 
**Balances** | [**map[string]UnifiedBalance**](UnifiedBalance.md) |  | [optional] 
**Total** | **string** | Total account assets converted to USD, i.e. the sum of &#x60;(available + freeze) * price&#x60; in all currencies (deprecated, to be removed, replaced by unified_account_total) | [optional] 
**Borrowed** | **string** | Total borrowed amount converted to USD, i.e. the sum of &#x60;borrowed * price&#x60; of all currencies (excluding point cards), valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**TotalInitialMargin** | **string** | Total initial margin, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**TotalMarginBalance** | **string** | Total margin balance, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**TotalMaintenanceMargin** | **string** | Total maintenance margin is valid in cross-currency margin/combined margin mode, and is 0 in other modes such as single-currency margin mode | [optional] 
**TotalInitialMarginRate** | **string** | Total initial margin rate, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**TotalMaintenanceMarginRate** | **string** | Total maintenance margin rate, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**TotalAvailableMargin** | **string** | Available margin amount, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**UnifiedAccountTotal** | **string** | Total unified account assets, valid in single currency margin/cross-currency margin/combined margin mode | [optional] 
**UnifiedAccountTotalLiab** | **string** | Total unified account borrowed amount, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**UnifiedAccountTotalEquity** | **string** | Total unified account equity, valid in single currency margin/cross-currency margin/combined margin mode | [optional] 
**Leverage** | **string** | Actual leverage ratio, valid in cross-currency margin/combined margin mode | [optional] [readonly] 
**SpotOrderLoss** | **string** | Spot Pending Order Loss, in USDT, effective only in Cross-Currency Margin Mode and Portfolio Margin Mode. | [optional] 
**OptionsOrderLoss** | **string** | Option Pending Order Loss, in USDT, effective only in Portfolio Margin Mode. | [optional] 
**SpotHedge** | **bool** | Spot hedging status: true - enabled, false - disabled | [optional] 
**UseFunding** | **bool** | Whether to use Earn funds as margin | [optional] 
**IsAllCollateral** | **bool** | Whether all currencies are used as margin: true - all currencies as margin, false - no | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


