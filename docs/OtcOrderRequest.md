# OtcOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | BUY for on-ramp, SELL for off-ramp | 
**Side** | **string** | Quote direction returned by the quote API (used for order validation) | 
**CryptoCurrency** | **string** | Cryptocurrency (supported currencies can be queried from the OTC web fiat quote page) | 
**FiatCurrency** | **string** | Fiat currency (supported currencies can be queried from the OTC web fiat quote page) | 
**CryptoAmount** | **string** | Amount of cryptocurrency | 
**FiatAmount** | **string** | Fiat amount | 
**PromotionCode** | **string** | Promotion code | [optional] 
**QuoteToken** | **string** | Parameter returned by the quote API | 
**BankId** | **string** | Bank card ID used for the order (retrieved via the default bank card API) | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


