# OtcStableCoinOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayCoin** | **string** | Currency paid by the user. Supported currencies can be queried from the OTC web stablecoin quote page. | 
**GetCoin** | **string** | Currency to be received by the user. Supported currencies can be queried from the OTC web stablecoin quote page. | 
**PayAmount** | **string** | User payment currency amount | 
**GetAmount** | **string** | Amount of currency received by the user | 
**Side** | **string** | The side returned by the quote endpoint (used for order validation). For backward compatibility, &#x60;PAY&#x60;/&#x60;GET&#x60; are accepted; new integrations should use the value returned by the quote response. | 
**PromotionCode** | **string** | Promotion code (optional) | [optional] 
**QuoteToken** | **string** | Parameter returned by the quote API | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


