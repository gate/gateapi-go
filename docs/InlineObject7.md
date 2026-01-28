# InlineObject7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CryptoCurrency** | **string** | Cryptocurrency | 
**FiatCurrency** | **string** | Fiat currency | 
**OrderTab** | **string** | Order tab, default is pending (pending: Processing (pending: AND status in (&#39;OPEN&#39;,  &#39;PAID&#39;, &#39;LOCKED&#39;, &#39;TEMP&#39;)); dispute: In dispute (status in (&#39;ACCEPT&#39;,  &#39;BCLOSED&#39;, &#39;CANCEL&#39;, &#39;BECANCEL&#39;, &#39;SCLOSED&#39;, &#39;SCANCEL&#39;))) | [optional] 
**SelectType** | **string** | Buy/Sell (sell&#x3D;Sell, buy&#x3D;Buy, others&#x3D;All) | [optional] 
**Status** | **string** | 订单状态（dispute: 申诉订单； closed: ACCEPT、BCLOSED； cancel： CANCEL、BECANCEL、SCLOSED、SCANCEL； locked: LOCKED； open: OPEN； paid： PAID； completed： CANCEL、BECANCEL、SCLOSED、SCANCEL、ACCEPT、BCLOSED） | [optional] 
**Txid** | **int32** | Order ID | [optional] 
**StartTime** | **int32** | Start timestamp, default is 00:00 89 days ago | [optional] 
**EndTime** | **int32** | End timestamp, default is 23:59:59 today | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


