# UnifiedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Unified account mode: - classic: Classic account mode - multi_currency: Multi-currency margin mode - portfolio: Portfolio margin mode - single_currency: Single-currency margin mode | [optional] 
**UserId** | **int64** | User ID | [optional] 
**RefreshTime** | **int64** | Last refresh time | [optional] 
**Locked** | **bool** | Whether the account is locked, valid in cross-currency margin/combined margin mode, false in other modes such as single-currency margin mode | [optional] 
**Balances** | [**map[string]UnifiedBalance**](UnifiedBalance.md) |  | [optional] 
**Total** | **string** | Total account assets converted to USD, i.e. the sum of &#x60;(available + freeze) * price&#x60; in all currencies (deprecated, to be removed, replaced by unified_account_total) | [optional] 
**Borrowed** | **string** | Total borrowed amount converted to USD, i.e. the sum of &#x60;borrowed * price&#x60; of all currencies (excluding point cards), valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**TotalInitialMargin** | **string** | Total initial margin (cross), effective in multi-currency margin/portfolio margin mode, 0 in single-currency margin mode | [optional] 
**TotalMarginBalance** | **string** | Total margin balance (cross), effective in multi-currency margin/portfolio margin mode, 0 in single-currency margin mode | [optional] 
**TotalMaintenanceMargin** | **string** | Total maintenance margin (cross), effective in multi-currency margin/portfolio margin mode, 0 in single-currency margin mode | [optional] 
**TotalInitialMarginRate** | **string** | Total initial margin rate (cross), effective in multi-currency margin/portfolio margin mode, 0 in single-currency margin mode | [optional] 
**TotalMaintenanceMarginRate** | **string** | Total maintenance margin rate (cross), effective in multi-currency margin/portfolio margin mode, 0 in single-currency margin mode | [optional] 
**TotalAvailableMargin** | **string** | Available margin amount, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**UnifiedAccountTotal** | **string** | Total unified account assets, includes both cross and isolated total assets in single-currency/multi-currency mode, only cross total assets in portfolio margin mode | [optional] 
**UnifiedAccountTotalLiab** | **string** | Total unified account borrowed, i.e. total cross borrowed, effective in multi-currency margin/portfolio margin mode, 0 in single-currency margin mode | [optional] 
**UnifiedAccountTotalEquity** | **string** | Total unified account equity, includes both cross and isolated total equity in single-currency/multi-currency mode, only cross total equity in portfolio margin mode | [optional] 
**Leverage** | **string** | Account leverage multiplier, effective in multi-currency/portfolio margin mode (deprecated). Currency leverage query API: GET /unified/leverage/user_currency_setting | [optional] [readonly] 
**SpotOrderLoss** | **string** | Spot Pending Order Loss, in USDT, effective only in Cross-Currency Margin Mode and Portfolio Margin Mode. | [optional] 
**OptionsOrderLoss** | **string** | Option Pending Order Loss, in USDT, effective only in Portfolio Margin Mode. | [optional] 
**SpotHedge** | **bool** | Spot hedging status: true - enabled, false - disabled | [optional] 
**UseFunding** | **bool** | Whether to use Earn funds as margin | [optional] 
**IsAllCollateral** | **bool** | Whether all currencies are used as margin: true - all currencies as margin, false - no | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


