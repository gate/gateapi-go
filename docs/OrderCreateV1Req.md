# OrderCreateV1Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**[]CreateParam**](CreateParam.md) | Sell ​​side list, at least one item; each item is the currency and amount &#x60;amount&#x60; to be swapped out. | 
**To** | [**[]CreateParam**](CreateParam.md) | Target side list, at least one item; each item is the target currency and **amount** &#x60;amount&#x60; (non-proportional). The structural semantics are different from &#x60;OrderPreviewV1Req.to&#x60; (&#x60;PreviewToParam&#x60;, including &#x60;ratio&#x60;), so do not mix them. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


