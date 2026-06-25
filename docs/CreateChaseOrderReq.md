# CreateChaseOrderReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | **string** | Contract name; server-side converted to uppercase | 
**Settle** | **string** | Settle currency, overridden by the path parameter and converted to lowercase | [optional] 
**Amount** | **string** | Total order size in contracts, decimal string. Positive for buy, negative for sell. Cannot be 0 | 
**PriceLimit** | **string** | 最高追逐价，合法十进制字符串；未设置限价时请传 \&quot;0\&quot; | 
**OffsetLimit** | **string** | Maximum chasing distance from the best price, mutually exclusive with price_limit | [optional] 
**ReduceOnly** | **bool** | Whether reduce only | [optional] 
**Text** | **string** | Optional custom tag | [optional] 
**IsDualMode** | **bool** | Whether dual-position mode is enabled | [optional] 
**PriceType** | **int64** | Price type: 1 best bid/ask, 2 distance from best bid/ask | [optional] 
**PriceGapType** | **int64** | Used when price_type &#x3D;&#x3D; 2: 1 absolute price gap, 2 percentage | [optional] 
**PriceGapValue** | **string** | Price gap value paired with price_gap_type | [optional] 
**PosMarginMode** | **string** | Position margin mode, e.g. isolated or cross | [optional] 
**PositionMode** | **string** | Position mode (e.g. single, dual, dual_plus) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


