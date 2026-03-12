# OrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | [optional] 
**TxHash** | **string** | Transaction Hash | [optional] 
**Side** | **string** | Buy or sell orders - buy - sell | [optional] 
**UsdtAmount** | **string** | Amount (USDT) | [optional] 
**Currency** | **string** | Token | [optional] 
**CurrencyAmount** | **string** | Token amount | [optional] 
**Status** | **int32** | Order Status - &#x60;0&#x60; : All - &#x60;1&#x60; : Processing - &#x60;2&#x60; : Successful - &#x60;3&#x60; : Failed - &#x60;4&#x60; : Cancelled - &#x60;5&#x60; : Buy order placed but transfer not completed - &#x60;6&#x60; : Order cancelled but transfer not completed | [optional] 
**GasMode** | **string** | Trading mode affects slippage selection - &#x60;speed&#x60; : Smart mode - &#x60;custom&#x60; : Custom mode, uses &#x60;slippage&#x60; parameter | [optional] 
**Chain** | **string** | Blockchain | [optional] 
**GasFee** | **string** | Gas Fee (USDT-based) | [optional] 
**TransactionFee** | **string** | Trading Fee (USDT-based) | [optional] 
**FailedReason** | **string** | Failure reason (if applicable) | [optional] 
**CreateTime** | **int64** | Creation timestamp | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


