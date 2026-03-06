# GetCompletedTransactionListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CryptoCurrency** | **string** | Cryptocurrency | 
**FiatCurrency** | **string** | Fiat currency | 
**SelectType** | **string** | Buy/Sell (sell&#x3D;Sell, buy&#x3D;Buy, others&#x3D;All) | [optional] 
**Status** | **string** | Order Status (dispute: Disputed Order; closed: ACCEPT, BCLOSED; cancel: CANCEL, BECANCEL, SCLOSED, SCANCEL; locked: LOCKED; open: OPEN; paid: PAID; completed: CANCEL, BECANCEL, SCLOSED, SCANCEL, ACCEPT, BCLOSED) | [optional] 
**Txid** | **int32** | Order ID | [optional] 
**StartTime** | **int32** | Start timestamp, default is 00:00 89 days ago | [optional] 
**EndTime** | **int32** | End timestamp, default is 23:59:59 today | [optional] 
**QueryDispute** | **int32** | 1: Include appeal status, 0: None | [optional] 
**Page** | **int32** | page number | [optional] 
**PerPage** | **int32** | Number of orders per page | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


