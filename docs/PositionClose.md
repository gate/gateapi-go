# PositionClose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **float64** | Position close time | [optional] [readonly] 
**Contract** | **string** | Futures contract | [optional] [readonly] 
**Side** | **string** | Position side  - &#x60;long&#x60;: Long position - &#x60;short&#x60;: Short position | [optional] [readonly] 
**Pnl** | **string** | PnL | [optional] [readonly] 
**PnlPnl** | **string** | PNL - Position P/L | [optional] [readonly] 
**PnlFund** | **string** | PNL - Funding Fees | [optional] [readonly] 
**PnlFee** | **string** | PNL - Transaction Fees | [optional] [readonly] 
**Text** | **string** | Source of close order. See &#x60;order.text&#x60; field for specific values | [optional] [readonly] 
**MaxSize** | **string** | Max Trade Size | [optional] [readonly] 
**AccumSize** | **string** | Cumulative closed position volume | [optional] [readonly] 
**FirstOpenTime** | **int64** | First Open Time | [optional] [readonly] 
**LongPrice** | **string** | When side is &#39;long&#39;, it indicates the opening average price; when side is &#39;short&#39;, it indicates the closing average price | [optional] [readonly] 
**ShortPrice** | **string** | When side is &#39;long&#39;, it indicates the closing average price; when side is &#39;short&#39;, it indicates the opening average price | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


