# OtcOrderListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **string** | Current time | [optional] 
**Timestamp** | **int32** | Current timestamp | [optional] 
**OrderId** | **string** | orderId | [optional] 
**TradeNo** | **string** | Trade number | [optional] 
**Type** | **string** | Quote direction buy/sell/all | [optional] 
**Status** | **string** | Order Status | [optional] 
**DbStatus** | **string** |  | [optional] 
**FiatCurrency** | **string** | Fiat type | [optional] 
**FiatCurrencyInfo** | [**OtcOrderListFiatCurrencyInfo**](OtcOrderListFiatCurrencyInfo.md) |  | [optional] 
**FiatAmount** | **string** | Fiat amount | [optional] 
**CryptoCurrency** | **string** | Stablecoin | [optional] 
**CryptoCurrencyInfo** | [**OtcOrderListCryptoCurrencyInfo**](OtcOrderListCryptoCurrencyInfo.md) |  | [optional] 
**CryptoAmount** | **string** | Stablecoin amount | [optional] 
**Rate** | **string** | Exchange rate | [optional] 
**TransferRemark** | **string** | Transfer remark (mutually exclusive with reference_code; empty string when the deposit buy order has a reference code) | [optional] 
**ReferenceCode** | **string** | Unique bank transfer reference code for deposit buy orders (SGB deposit scenario) | [optional] 
**GateBankAccountIban** | **string** | Bank account | [optional] 
**PromotionCode** | **string** | Promotion code | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


