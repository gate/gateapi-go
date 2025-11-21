# UidPushOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Order ID | [optional] 
**PushUid** | **int64** | Initiator User ID | [optional] 
**ReceiveUid** | **int64** | Recipient User ID | [optional] 
**Currency** | **string** | Currency name | [optional] 
**Amount** | **string** | Transfer amount | [optional] 
**CreateTime** | **int64** | Created time | [optional] 
**Status** | **string** | Withdrawal status:  - CREATING: Creating - PENDING: Waiting for recipient (Please contact the recipient to accept the transfer on Gate official website) - CANCELLING: Cancelling - CANCELLED: Cancelled - REFUSING: Refusing - REFUSED: Refused - RECEIVING: Receiving - RECEIVED: Success | [optional] 
**Message** | **string** | PENDING reason tips | [optional] 
**TransactionType** | **string** | Order Type | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


