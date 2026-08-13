# OtcOrderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | 
**Uid** | **string** | User ID | 
**Type** | **string** | Order Type | 
**FiatCurrency** | **string** | Fiat currency | 
**FiatAmount** | **string** | Fiat amount | 
**CryptoCurrency** | **string** | Digital currency | 
**CryptoAmount** | **string** | Cryptocurrency amount | 
**Rate** | **string** | Exchange rate | 
**BankAccountName** | **string** | User payment/receiving name | [optional] 
**BankName** | **string** | User payment/receiving bank name | [optional] 
**BankCountry** | **string** | User payment/receiving bank country | [optional] 
**BankAddress** | **string** | User payment/receiving bank address | [optional] 
**BankAccountNumberIban** | **string** | User payment/receiving bank account number/IBAN | [optional] 
**SwiftCode** | **string** | User payment/receiving bank SWIFT code | [optional] 
**IntermediateBankName** | **string** | User payment/receiving intermediary bank name | [optional] 
**IntermediaryBankSwiftCode** | **string** | User payment/receiving intermediary bank SWIFT code | [optional] 
**GateBankAccountName** | **string** | Gate beneficiary name, shown for BUY only | [optional] 
**GateBankName** | **string** | Gate beneficiary bank name, shown for BUY only | [optional] 
**GateBankCountry** | **string** | Gate beneficiary bank country, shown for BUY only | [optional] 
**GateBankAddress** | **string** | Gate beneficiary bank address, shown for BUY only | [optional] 
**GateBankAccountNumberIban** | **string** | Gate beneficiary bank account number/IBAN, shown for BUY only | [optional] 
**GateSwiftCode** | **string** | Gate beneficiary bank SWIFT code, shown for BUY only | [optional] 
**GateIntermediaryBankName** | **string** | Gate beneficiary intermediary bank name, shown for BUY only | [optional] 
**GateIntermediaryBankSwiftCode** | **string** | Gate beneficiary intermediary bank SWIFT code, shown for BUY only | [optional] 
**GateTransferRemark** | **string** | Transfer remark (mutually exclusive with &#x60;gate_reference_code&#x60;; empty when a BUY deposit order has a reference code), shown for BUY only | [optional] 
**GateReferenceCode** | **string** | Be sure to include the reference code when making the transfer so that your order can be processed promptly. (Mutually exclusive with &#x60;gate_transfer_remark&#x60;.) | [optional] 
**Status** | **string** | Status | 
**CreateTime** | **string** | Created time | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


