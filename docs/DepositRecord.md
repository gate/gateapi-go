# DepositRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Record ID | [optional] [readonly] 
**Txid** | **string** | Hash record of the withdrawal | [optional] [readonly] 
**Timestamp** | **string** | Operation time | [optional] [readonly] 
**Amount** | **string** | Token amount | 
**Currency** | **string** | Currency name | 
**Address** | **string** | Withdrawal address. Required for withdrawals | [optional] 
**Memo** | **string** | Additional remarks with regards to the withdrawal | [optional] 
**Status** | **string** | Transaction Status  - BLOCKED: Deposit Blocked - DEP_CREDITED: Deposit Credited, Withdrawal Pending Unlock - DONE: Funds Credited to Spot Account - INVALID: Invalid Transaction - MANUAL: Manual Review Required - PEND: Processing - REVIEW: Under Compliance Review - TRACK: Tracking Block Confirmations, Pending Spot Account Credit | [optional] [readonly] 
**RefundStatus** | **string** | Blocked deposit refund status. This field is returned only when the deposit record has a blocked deposit refund record with a non-empty refund status. Not returned when there is no refund record or the refund status is empty - REFUNDING: Refund in progress - REFUNDED: Refund completed - REFUND_FAILED: Refund failed - REJECTED: Refund rejected | [optional] [readonly] 
**Chain** | **string** | Name of the chain used in withdrawals | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


