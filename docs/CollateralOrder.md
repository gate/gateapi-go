# CollateralOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int64** | Order ID | [optional] 
**CollateralCurrency** | **string** | Collateral currency | [optional] 
**CollateralAmount** | **string** | Collateral amount | [optional] 
**BorrowCurrency** | **string** | Borrowed currency | [optional] 
**BorrowAmount** | **string** | Borrowed amount | [optional] 
**RepaidAmount** | **string** | Repaid amount | [optional] 
**RepaidPrincipal** | **string** | Repaid principal | [optional] 
**RepaidInterest** | **string** | Repaid interest | [optional] 
**InitLtv** | **string** | Initial collateralization rate | [optional] 
**CurrentLtv** | **string** | Current collateralization rate | [optional] 
**LiquidateLtv** | **string** | Liquidation collateralization rate | [optional] 
**Status** | **string** | Order status: - initial: Initial state after placing the order - collateral_deducted: Collateral deduction successful - collateral_returning: Loan failed - Collateral return pending - lent: Loan successful - repaying: Repayment in progress - liquidating: Liquidation in progress - finished: Order completed - closed_liquidated: Liquidation and repayment completed | [optional] 
**BorrowTime** | **int64** | Borrowing time, timestamp in seconds | [optional] 
**LeftRepayTotal** | **string** | Outstanding principal and interest (outstanding principal + outstanding interest) | [optional] 
**LeftRepayPrincipal** | **string** | Outstanding principal | [optional] 
**LeftRepayInterest** | **string** | Outstanding interest | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


