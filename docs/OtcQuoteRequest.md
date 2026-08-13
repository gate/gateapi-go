# OtcQuoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Side** | **string** | PAY: specify the payment amount (&#x60;pay_amount&#x60; is required); GET: specify the receive amount (&#x60;get_amount&#x60; is required). | 
**PayCoin** | **string** | Payment currency. Supported currencies are available on the OTC web quote page. | 
**GetCoin** | **string** | Receive currency. Supported currencies are available on the OTC web quote page. | 
**PayAmount** | **string** | User payment currency amount | [optional] 
**GetAmount** | **string** | Amount of currency received by the user | [optional] 
**CreateQuoteToken** | **string** | Create quote token: 0: quote preview only; 1: generate quote token for order placement. | [optional] 
**PromotionCode** | **string** | Promotion code | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


