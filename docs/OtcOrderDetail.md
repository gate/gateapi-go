# OtcOrderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | 
**Uid** | **string** | User ID | 
**Type** | **string** | Order Type | 
**FiatCurrency** | **string** | Fiat type | 
**FiatAmount** | **string** | Fiat amount | 
**CryptoCurrency** | **string** | Stablecoin | 
**CryptoAmount** | **string** | Stablecoin amount | 
**Rate** | **string** | Exchange rate | 
**TransferRemark** | **string** | Transfer remark (mutually exclusive with reference_code; empty string when the deposit buy order has a reference code) | 
**ReferenceCode** | **string** | Unique bank transfer reference code for deposit buy orders (SGB deposit scenario; mutually exclusive with transfer_remark) | [optional] 
**Status** | **string** | Status | 
**DbStatus** | **string** |  | 
**CreateTime** | **string** | Created time | 
**Memo** | **string** | Cancellation or rejection reason | 
**Side** | **string** | Quote direction | 
**PromotionCode** | **string** | Promotion code | 
**TradeNo** | **string** | Trade number | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


