# FuturesAutoDeleverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **int64** | Automatic deleveraging time | [optional] [readonly] 
**User** | **int64** | User ID | [optional] [readonly] 
**OrderId** | **int64** | Order ID. Order IDs before 2023-02-20 are null | [optional] [readonly] 
**Contract** | **string** | Futures contract | [optional] [readonly] 
**Leverage** | **string** |  leverage for isolated margin. 0 means cross margin. For leverage of cross margin, please refer to &#x60;cross_leverage_limit&#x60;. | [optional] [readonly] 
**CrossLeverageLimit** | **string** | leverage for cross margin | [optional] [readonly] 
**EntryPrice** | **string** | Average entry price | [optional] [readonly] 
**FillPrice** | **string** | Average fill price | [optional] [readonly] 
**TradeSize** | **string** | Trading size | [optional] [readonly] 
**PositionSize** | **string** | Positions after auto-deleveraging | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


