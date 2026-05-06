# P2pChatMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Order status when sending a message. Typical values: &#x60;OPEN&#x60;, &#x60;PAID&#x60;, &#x60;LOCKED&#x60;, &#x60;ACCEPT&#x60;, &#x60;BCLOSED&#x60;, &#x60;CANCEL&#x60;, &#x60;BECANCEL&#x60;, &#x60;SCLOSED&#x60;, &#x60;SCANCEL&#x60;. | [optional] 
**Text** | **string** | Message content | [optional] 
**PaymentVoucher** | **[]string** | Payment voucher | [optional] 
**ReasonId** | **int32** | Cancel reason ID. &#x60;1&#x60; no longer want to buy; &#x60;2&#x60; cannot reach seller; &#x60;3&#x60; will not pay; &#x60;4&#x60; seller account not real; &#x60;5&#x60; payout account issue; &#x60;6&#x60; price mismatch; &#x60;7&#x60; mutually agreed cancel; &#x60;8&#x60; poor communication; &#x60;9&#x60; other; &#x60;10&#x60; seller cannot release with refund; &#x60;11&#x60; terms not met; &#x60;12&#x60; seller payout risk-controlled. | [optional] 
**ToastId** | **int32** | Cancellation reason popup | [optional] 
**ReasonMemo** | **string** | Cancel reason description. | [optional] 
**CancelTime** | **int32** | Cancellation time | [optional] 
**SellerConfirm** | **int32** | Seller confirmation of cancel reason: &#x60;0&#x60; pending; &#x60;1&#x60; confirmed; &#x60;2&#x60; rejected. | [optional] 
**Id** | **string** | Payment method information ID | [optional] 
**AccountDes** | **string** | Payment method description | [optional] 
**PayType** | **string** | Payment method type | [optional] 
**File** | **string** | Payment method file link | [optional] 
**FileKey** | **string** | Payment method file key | [optional] 
**Account** | **string** | Payment account or masked payment account. | [optional] 
**Memo** | **string** | Payment method note | [optional] 
**Code** | **string** | Payment method code | [optional] 
**MemoExt** | **string** | Payment method additional note | [optional] 
**TradeTips** | **string** | Payment method tip | [optional] 
**RealName** | **string** | Payment method username | [optional] 
**IsDelete** | **int32** | Whether the payment method was deleted. &#x60;1&#x60;: deleted; &#x60;0&#x60;: not deleted. | [optional] 
**PayName** | **string** | Payment method full name | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


