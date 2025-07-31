# CreateMultiCollateralOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | [optional] 
**OrderType** | **string** | current - current rate, fixed - fixed rate, defaults to current if not specified | [optional] 
**FixedType** | **string** | Fixed interest rate lending period: 7d - 7 days, 30d - 30 days. Required for fixed rate | [optional] 
**FixedRate** | **string** | Fixed interest rate, required for fixed rate | [optional] 
**AutoRenew** | **bool** | Fixed interest rate, auto-renewal | [optional] 
**AutoRepay** | **bool** | Fixed interest rate, auto-repayment | [optional] 
**BorrowCurrency** | **string** | Borrowed currency | 
**BorrowAmount** | **string** | Borrowed amount | 
**CollateralCurrencies** | [**[]CollateralCurrency**](CollateralCurrency.md) | Collateral currency and amount | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


