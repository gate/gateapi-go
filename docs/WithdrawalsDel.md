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
**Status** | **string** | 交易状态  - BCODE: 充值码操作 - CANCEL: 已取消 - CANCELPEND: 取消提现中 - DMOVE: 待人工审核 - DONE: 完成 (block_number &gt; 0 才算真的上链完成) - EXTPEND: 已经发送等待确认 - FAIL: 链上失败等待确认 - FVERIFY: 人脸审核处理中 - LOCKED: 钱包侧锁单 - MANUAL: 待人工审核 - REJECT: 拒绝 - REQUEST: 请求中 - REVIEW: 审核中 | [optional] [readonly] 
**Chain** | **string** | Name of the chain used in withdrawals | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


