# InlineObject7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Side** | **string** | PAY/GET quote direction. PAY means user inputs pay amount, GET means user inputs get amount. If PAY, pay_amount is required. If GET, get_amount is required | 
**PayCoin** | **string** | Currency the user pays. Supported currencies can be found on the OTC web quote page. | 
**GetCoin** | **string** | Currency the user receives. Supported currencies can be found on the OTC web quote page. | 
**PayAmount** | **string** | User payment currency amount | [optional] 
**GetAmount** | **string** | Amount of currency received by the user | [optional] 
**CreateQuoteToken** | **string** | Create quote token: 0: quote preview only; 1: generate quote token for order placement. | [optional] 
**PromotionCode** | **string** | Promotion code (optional) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


