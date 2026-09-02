# CrossexAccountBookRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Account Change Record ID | 
**UserId** | **string** | User ID | 
**BusinessId** | **string** | Business ID. Its meaning varies by &#x60;statement_type&#x60;. TRANSACTION: order ID TRADING_FEE: order ID LIQUIDATION_FEE: liquidation order ID FUNDING_FEE: position ID and funding fee settlement time For other types, it is a system-generated processing ID with no business meaning | 
**StatementType** | **string** | Account book entry type &#x60;TRANSACTION&#x60;: trade &#x60;TRADING_FEE&#x60;: trading fee &#x60;FUNDING_FEE&#x60;: futures funding fee &#x60;LIQUIDATION_FEE&#x60;: liquidation fee &#x60;TRANSFER_IN&#x60;: transfer in &#x60;TRANSFER_OUT&#x60;: transfer out &#x60;BANKRUPT_COMPENSATION&#x60;: bankruptcy compensation &#x60;AUTO_REPAY&#x60;: automatic repayment of margin position liabilities &#x60;INTEREST_ISOLATED&#x60;: interest entry &#x60;ACCOUNT_MODE_CHANGE&#x60;: balance change caused by an account mode switch &#x60;KRAKEN_CONVERSION&#x60;: conversion of other margin currencies to cover a negative KRAKEN_USD balance &#x60;OTHER&#x60;: other | 
**ExchangeType** | **string** | Exchange | 
**Coin** | **string** | Currency | 
**Symbol** | **string** | Trading Pair | [optional] 
**Change** | **string** | Change amount (positive values indicate an increase; negative values indicate a decrease) | 
**Balance** | **string** | Balance after change | 
**CreateTime** | **string** | Created time | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


