# UnifiedBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **string** | Cross available balance, deducted futures isolated margin occupation and frozen amount (futures isolated occupation, i.e. futures isolated balance), effective in single-currency/multi-currency/portfolio margin mode. | [optional] 
**Freeze** | **string** | Frozen amount, effective in single-currency/multi-currency/portfolio margin mode | [optional] 
**Borrowed** | **string** | Borrowed amount, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**NegativeLiab** | **string** | Negative balance borrowing, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**FuturesPosLiab** | **string** | Contract opening position borrowing currency (abandoned, to be offline field) | [optional] 
**Equity** | **string** | Currency equity amount (cross), effective in single-currency/multi-currency/portfolio margin mode | [optional] 
**TotalFreeze** | **string** | Total frozen (deprecated, to be removed) | [optional] 
**TotalLiab** | **string** | Total borrowed amount, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**SpotInUse** | **string** | The amount of spot hedging is valid in the combined margin mode, and is 0 in other margin modes such as single currency and cross-currency margin modes | [optional] 
**Funding** | **string** | Uniloan financial management amount, effective when turned on as a unified account margin switch | [optional] 
**FundingVersion** | **string** | Funding version | [optional] 
**CrossBalance** | **string** | Full margin balance is valid in single currency margin mode, and is 0 in other modes such as cross currency margin/combined margin mode | [optional] 
**IsoBalance** | **string** | Futures isolated balance, effective in single-currency and multi-currency margin mode, 0 in portfolio margin mode | [optional] 
**Im** | **string** | Cross initial margin, only effective for USDT in single-currency margin mode, 0 in multi-currency/portfolio margin mode | [optional] 
**Mm** | **string** | Cross maintenance margin, only effective for USDT in single-currency margin mode, 0 in multi-currency/portfolio margin mode | [optional] 
**Imr** | **string** | Cross initial margin rate, only effective for USDT in single-currency margin mode, 0 in multi-currency/portfolio margin mode | [optional] 
**Mmr** | **string** | Cross maintenance margin rate, only effective for USDT in single-currency margin mode, 0 in multi-currency/portfolio margin mode | [optional] 
**MarginBalance** | **string** | Cross margin balance, only effective for USDT in single-currency margin mode, 0 in multi-currency/portfolio margin mode | [optional] 
**AvailableMargin** | **string** | Cross available margin, only effective for USDT in single-currency margin mode, 0 in multi-currency/portfolio margin mode | [optional] 
**EnabledCollateral** | **bool** | Currency enabled as margin: true - Enabled, false - Disabled | [optional] 
**BalanceVersion** | **float32** | Balance version number | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


