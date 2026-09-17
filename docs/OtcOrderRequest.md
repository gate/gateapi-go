# OtcOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | BUY for on-ramp, SELL for off-ramp | 
**Side** | **string** | The side returned by the quote endpoint (used for order validation). For backward compatibility, &#x60;FIAT&#x60;/&#x60;CRYPTO&#x60; or &#x60;PAY&#x60;/&#x60;GET&#x60; are accepted; new integrations should use the value returned by the quote response. | 
**CryptoCurrency** | **string** | Cryptocurrency (supported currencies can be queried from the OTC web fiat quote page) | 
**FiatCurrency** | **string** | Fiat currency (supported currencies can be queried from the OTC web fiat quote page) | 
**CryptoAmount** | **string** | Amount of cryptocurrency | 
**FiatAmount** | **string** | Fiat amount | 
**PromotionCode** | **string** | Promotion code | [optional] 
**QuoteToken** | **string** | Parameter returned by the quote API | 
**BankId** | **string** | Bank card ID used to place the order. Select one from the list returned by &#x60;GET /otc/bank/list&#x60;; the default card has &#x60;is_default&#x3D;1&#x60;. | 
**ReceiveType** | **string** | Name used for the remittance. Allowed values depend on the user type: Corporate users: YOU (remit in your company&#39;s name), GATE (remit in Gate&#39;s name), RECIPIENT (remit in the recipient&#39;s name); Individual users: GATE (remit in Gate&#39;s name), PERSON (remit in the user&#39;s own name). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


