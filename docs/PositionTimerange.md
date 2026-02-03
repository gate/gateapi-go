# PositionTimerange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | **string** | Futures contract | [optional] [readonly] 
**Size** | **string** | Position size | [optional] [readonly] 
**Leverage** | **string** | Position leverage. 0 means cross margin; positive number means isolated margin | [optional] 
**RiskLimit** | **string** | Position risk limit | [optional] 
**LeverageMax** | **string** | the maximum permissible leverage given to the current positon value: the higher positon value, the lower maximum permissible leverage | [optional] [readonly] 
**MaintenanceRate** | **string** | The maintenance margin rate of the first tier of risk limit sheet | [optional] [readonly] 
**Margin** | **string** | Position margin | [optional] 
**LiqPrice** | **string** | Liquidation price | [optional] [readonly] 
**RealisedPnl** | **string** | Realized PnL | [optional] [readonly] 
**HistoryPnl** | **string** | Total realized PnL from closed positions | [optional] [readonly] 
**LastClosePnl** | **string** | PNL of last position close | [optional] [readonly] 
**RealisedPoint** | **string** | Realized POINT PNL | [optional] [readonly] 
**HistoryPoint** | **string** | History realized POINT PNL | [optional] [readonly] 
**Mode** | **string** | Position mode, including: - &#x60;single&#x60;: One-way Mode - &#x60;dual_long&#x60;: Long position in Hedge Mode - &#x60;dual_short&#x60;: Short position in Hedge Mode | [optional] 
**CrossLeverageLimit** | **string** | Cross margin leverage (valid only when &#x60;leverage&#x60; is 0) | [optional] 
**EntryPrice** | **string** | Entry price | [optional] [readonly] 
**Time** | **int64** | Timestamp | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


