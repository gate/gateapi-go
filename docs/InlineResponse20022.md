# InlineResponse20022

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Order ID | 
**Text** | **string** | Client Custom ID | 
**FromAccountType** | **string** | Source &#x60;from&#x60; account (CROSSEX_BINANCE, CROSSEX_OKX, CROSSEX_GATE, CROSSEX, SPOT) | 
**ToAccountType** | **string** |  | 
**Coin** | **string** | Currency | 
**Amount** | **string** | Transfer amount, the amount requested for the transfer | 
**ActualReceive** | **string** | Actual credited amount (has a value when status &#x3D; SUCCESS; empty for other statuses) | [optional] 
**Status** | **string** | Transfer Status - &#x60;FAIL&#x60;: Failed - &#x60;SUCCESS&#x60;: Successful - &#x60;PENDING&#x60;: Transfer in Progress | 
**FailReason** | **string** | Failure reason (has a value when status &#x3D; FAIL; empty for other statuses) | [optional] 
**CreateTime** | **int32** | Creation time of order | 
**UpdateTime** | **int32** | OrderUpdateTime | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


