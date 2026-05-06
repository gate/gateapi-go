# GetPendingTransactionListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CryptoCurrency** | **string** | Cryptocurrency symbol. | 
**FiatCurrency** | **string** | Fiat currency | 
**OrderTab** | **string** | Order tab: &#x60;pending&#x60; in progress (&#x60;OPEN&#x60;, &#x60;PAID&#x60;, &#x60;LOCKED&#x60;, &#x60;TEMP&#x60;); &#x60;dispute&#x60; in dispute; default &#x60;pending&#x60;. | [optional] 
**SelectType** | **string** | Order side filter: &#x60;buy&#x60; buy orders; &#x60;sell&#x60; sell orders; empty: all. | [optional] 
**Status** | **string** | Order status filter. &#x60;open&#x60; unpaid (&#x60;OPEN&#x60;); &#x60;paid&#x60; paid (&#x60;PAID&#x60;); &#x60;locked&#x60; locked (&#x60;LOCKED&#x60;); &#x60;dispute&#x60; in dispute; empty or omitted uses the default range for &#x60;order_tab&#x60;. | [optional] 
**Txid** | **int32** | Order ID | [optional] 
**StartTime** | **int32** | Start timestamp, default is 00:00 89 days ago | [optional] 
**EndTime** | **int32** | End timestamp, default is 23:59:59 today | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


