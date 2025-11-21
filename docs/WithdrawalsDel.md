# WithdrawalsDel

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
**BlockNumber** | **string** | Block Number | [optional] [readonly] 
**Status** | **string** | Transaction Status  - BCODE: Deposit Code Operation - CANCEL: Cancelled - CANCELPEND: Withdrawal Cancellation Pending - DMOVE: Pending Manual Review - DONE: Completed (Only considered truly on-chain when block_number &gt; 0) - EXTPEND: Sent and Waiting for Confirmation - FAIL: On-Chain Failure Pending Confirmation - FVERIFY: Facial Verification in Progress - LOCKED: Wallet-Side Order Locked - MANUAL: Pending Manual Review - REJECT: Rejected - REQUEST: Request in Progress - REVIEW: Under Review  | [optional] [readonly] 
**Chain** | **string** | Name of the chain used in withdrawals | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


