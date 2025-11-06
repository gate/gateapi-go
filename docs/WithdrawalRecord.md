# WithdrawalRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Record ID | [optional] [readonly] 
**Txid** | **string** | Hash record of the withdrawal | [optional] [readonly] 
**BlockNumber** | **string** | Block Number | [optional] [readonly] 
**WithdrawOrderId** | **string** | Client order id, up to 32 length and can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.)  | [optional] 
**Timestamp** | **string** | Operation time | [optional] [readonly] 
**Amount** | **string** | Token amount | 
**Fee** | **string** | fee | [optional] [readonly] 
**Currency** | **string** | Currency name | 
**Address** | **string** | Withdrawal address | [optional] 
**Type** | **string** | Business Type | [optional] 
**FailReason** | **string** | Reason for withdrawal failure. Has a value when status &#x3D; CANCEL, empty for all other statuses | [optional] 
**Timestamp2** | **string** | Withdrawal final time, i.e.: withdrawal cancellation time or withdrawal success time When status &#x3D; CANCEL, corresponds to cancellation time When status &#x3D; DONE and block_number &gt; 0, it is the withdrawal success time | [optional] 
**Memo** | **string** | Additional remarks with regards to the withdrawal | [optional] 
**Status** | **string** | Transaction Status  - BCODE: Deposit Code Operation - CANCEL: Cancelled - CANCELPEND: Withdrawal Cancellation Pending - DONE: Completed (Only considered truly on-chain when block_number &gt; 0) - EXTPEND: Sent and Waiting for Confirmation - FAIL: On-Chain Failure Pending Confirmation - FVERIFY: Facial Verification in Progress - LOCKED: Wallet-Side Order Locked - MANUAL: Pending Manual Review - REJECT: Rejected - REQUEST: Request in Progress - REVIEW: Under Review  | [optional] [readonly] 
**Chain** | **string** | Name of the chain used in withdrawals | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


