# UnifiedLoanRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID | [optional] [readonly] 
**Type** | **string** | Type: &#x60;borrow&#x60; - borrow, &#x60;repay&#x60; - repay | [optional] [readonly] 
**RepaymentType** | **string** | Repayment type: none - No repayment type, manual_repay - Manual repayment, auto_repay - Automatic repayment, cancel_auto_repay - Automatic repayment after order cancellation, different_currencies_repayment - Cross-currency repayment | [optional] [readonly] 
**BorrowType** | **string** | Borrowing type, returned when querying loan records: manual_borrow - Manual borrowing, auto_borrow - Automatic borrowing | [optional] 
**CurrencyPair** | **string** | Currency pair | [optional] [readonly] 
**Currency** | **string** | Currency | [optional] [readonly] 
**Amount** | **string** | Borrow or repayment amount | [optional] [readonly] 
**CreateTime** | **int64** | Created time | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


