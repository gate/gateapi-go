# P2pChatListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]P2pChatMessage**](P2pChatMessage.md) | Message List | [optional] 
**Memo** | **string** | Payment tip (displayed on homepage only) | [optional] 
**HasHistory** | **bool** | Whether historical records exist | [optional] 
**Txid** | **int32** | Order ID | [optional] 
**SRVTM** | **int32** | Timestamp of the latest message. | [optional] 
**OrderStatus** | **string** | Raw order status in DB; typical values: &#x60;OPEN&#x60;, &#x60;PAID&#x60;, &#x60;LOCKED&#x60;, &#x60;ACCEPT&#x60;, &#x60;BCLOSED&#x60;, &#x60;CANCEL&#x60;, &#x60;BECANCEL&#x60;, &#x60;SCLOSED&#x60;, &#x60;SCANCEL&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


