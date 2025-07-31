# LedgerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Record ID | [optional] [readonly] 
**Txid** | **string** | Hash record of the withdrawal | [optional] [readonly] 
**WithdrawOrderId** | **string** | User-defined order number for withdrawal. Default is empty. When not empty, the specified user-defined order number record will be queried | [optional] 
**Timestamp** | **string** | Operation time | [optional] [readonly] 
**Amount** | **string** | Token amount | 
**Currency** | **string** | Currency name | 
**Address** | **string** | Withdrawal address. Required for withdrawals | [optional] 
**Memo** | **string** | Additional remarks with regards to the withdrawal | [optional] 
**WithdrawId** | **string** | Withdrawal record ID starts with &#39;w&#39;, such as: w1879219868. When withdraw_id is not empty, only this specific withdrawal record will be queried, and time-based querying will be disabled | [optional] 
**AssetClass** | **string** | Withdrawal record currency type, empty by default. Supports users to query withdrawal records in main area and innovation area on demand. Valid values: SPOT, PILOT  SPOT: Main area PILOT: Innovation area | [optional] 
**Status** | **string** | Transaction status  - DONE: Completed - CANCEL: Cancelled - REQUEST: Requesting - MANUAL: Pending manual review - BCODE: GateCode operation - EXTPEND: Sent, waiting for confirmation - FAIL: Failed on chain, waiting for confirmation - INVALID: Invalid order - VERIFY: Verifying - PROCES: Processing - PEND: Processing - DMOVE: Pending manual review - REVIEW: Under review | [optional] [readonly] 
**Chain** | **string** | Name of the chain used in withdrawals | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


