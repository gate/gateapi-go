# CrossexAccountBookRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Account Change Record ID | 
**UserId** | **string** | User ID | 
**BusinessId** | **string** | Business ID | 
**Type** | **string** | Change type | &#x60;TRANSACTION&#x60; trade &#x60;TRADING_FEE&#x60; fee &#x60;FUNDING_FEE&#x60; futures funding fee &#x60;LIQUIDATION_FEE&#x60; liquidation fee &#x60;TRANSFER_IN&#x60; transfer in &#x60;TRANSFER_OUT&#x60; transfer out &#x60;BANKRUPT_COMPENSATION&#x60; bankruptcy compensation &#x60;AUTO_REPAY&#x60; margin position auto-repay | 
**ExchangeType** | **string** | Exchange | 
**Coin** | **string** | Currency | 
**Change** | **string** | Change amount (positive indicates transfer in; negative indicates transfer out) | 
**Balance** | **string** | Balance after change | 
**CreateTime** | **string** | Created time | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


