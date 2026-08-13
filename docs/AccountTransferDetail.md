# AccountTransferDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** | Transfer transaction ID | [readonly] 
**Status** | **string** | Transfer status:  - &#x60;pending&#x60;: Processing - &#x60;success&#x60;: Successful - &#x60;fail&#x60;: Failed | [readonly] 
**Currency** | **string** | Transfer currency | [readonly] 
**Amount** | **string** | Transfer amount | [readonly] 
**FromAccount** | **string** | Source account type:  - &#x60;spot&#x60;: Spot account - &#x60;margin&#x60;: Margin account - &#x60;futures&#x60;: Perpetual futures account - &#x60;delivery&#x60;: Delivery futures account - &#x60;options&#x60;: Options account - &#x60;unknown&#x60;: Unrecognized account type | [readonly] 
**ToAccount** | **string** | Destination account type:  - &#x60;spot&#x60;: Spot account - &#x60;margin&#x60;: Margin account - &#x60;futures&#x60;: Perpetual futures account - &#x60;delivery&#x60;: Delivery futures account - &#x60;options&#x60;: Options account - &#x60;unknown&#x60;: Unrecognized account type | [readonly] 
**Settle** | Pointer to **string** | Settlement currency for futures, delivery, and options transfers; otherwise, null | [readonly] 
**CurrencyPair** | Pointer to **string** | Currency pair for margin transfers; otherwise, null | [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


