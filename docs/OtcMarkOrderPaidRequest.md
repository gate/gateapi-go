# OtcMarkOrderPaidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | 
**ClientOrderId** | **string** | Client order ID (used by some gateway/Inner Pay paths, optional) | [optional] 
**PaymentReceiptFileKey** | **string** | User payment receipt: **required**. Stored as a file_key. Single file; jpg/jpeg/png/pdf; ≤4MB. | 
**PaymentReceipt** | **string** | Alias compatible with &#x60;payment_receipt_file_key&#x60; (depends on the gateway&#39;s external field name) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


