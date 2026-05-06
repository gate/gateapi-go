# GetCompletedTransactionListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CryptoCurrency** | **string** | Cryptocurrency symbol. | 
**FiatCurrency** | **string** | Fiat currency | 
**SelectType** | **string** | Order side filter: &#x60;buy&#x60; buy orders; &#x60;sell&#x60; sell orders; empty: all. | [optional] 
**Status** | **string** | Order status filter. &#x60;closed&#x60;: filled (&#x60;ACCEPT&#x60;, &#x60;BCLOSED&#x60;); &#x60;cancel&#x60;: canceled (&#x60;CANCEL&#x60;, &#x60;BECANCEL&#x60;, &#x60;SCLOSED&#x60;, &#x60;SCANCEL&#x60;); &#x60;locked&#x60;: locked (&#x60;LOCKED&#x60;); &#x60;open&#x60;: unpaid (&#x60;OPEN&#x60;); &#x60;paid&#x60;: paid (&#x60;PAID&#x60;); &#x60;completed&#x60;: finished or canceled (&#x60;CANCEL&#x60;, &#x60;BECANCEL&#x60;, &#x60;SCLOSED&#x60;, &#x60;SCANCEL&#x60;, &#x60;ACCEPT&#x60;, &#x60;BCLOSED&#x60;); Empty or omitted uses the endpoint default range. | [optional] 
**Txid** | **int32** | Order ID | [optional] 
**StartTime** | **int32** | Start timestamp, default is 00:00 89 days ago | [optional] 
**EndTime** | **int32** | End timestamp, default is 23:59:59 today | [optional] 
**QueryDispute** | **int32** | Whether to flag dispute status in the response. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**Page** | **int32** | Page number starting at 1; values below 1 are treated as 1. | [optional] 
**PerPage** | **int32** | Orders per page; default 10, max 200. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


