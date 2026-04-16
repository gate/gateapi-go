# OrderCreateV1Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**[]CreateParam**](CreateParam.md) | 卖出侧列表，至少一项；每项为要换出的币种及数量 &#x60;amount&#x60;。 | 
**To** | [**[]CreateParam**](CreateParam.md) | 目标侧列表，至少一项；每项为目标币种及**数量** &#x60;amount&#x60;（非比例）。 与 &#x60;OrderPreviewV1Req.to&#x60;（&#x60;PreviewToParam&#x60;，含 &#x60;ratio&#x60;）结构语义不同，勿混用。 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


