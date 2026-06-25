# P2pSendChatMessageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SRVTM** | **int32** | Timestamp when message was successfully sent (current timestamp) | [optional] 
**Txid** | **int32** | Order ID | [optional] 
**ConversationId** | **string** | Chat ID, formatted as both parties&#39; UIDs concatenated in ascending order | [optional] 
**MsgType** | **int32** | Message content type when risk control is hit. 0: text | [optional] 
**RiskType** | **int32** | Risk control display type. 1: off-platform traffic diversion risk; returned only when risk control is hit | [optional] 
**ToastMsg** | **string** | Risk control prompt message; returned only when risk_type&#x3D;1 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


