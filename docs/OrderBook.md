# OrderBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Order book ID, which is updated whenever the order book is changed. Valid only when &#x60;with_id&#x60; is set to &#x60;true&#x60; | [optional] 
**Current** | **int64** | The timestamp of the response data being generated (in milliseconds) | [optional] 
**Update** | **int64** | The timestamp of when the orderbook last changed (in milliseconds) | [optional] 
**Asks** | [**[][]string**](array.md) | Ask Depth | 
**Bids** | [**[][]string**](array.md) | Bid Depth | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


