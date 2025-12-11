# DeliveryPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int64** | User ID | [optional] [readonly] 
**Contract** | **string** | Futures contract | [optional] [readonly] 
**Size** | **int64** | Position size | [optional] [readonly] 
**Leverage** | **string** | Position leverage. 0 means cross margin; positive number means isolated margin | [optional] 
**RiskLimit** | **string** | Position risk limit | [optional] 
**LeverageMax** | **string** | Maximum leverage under current risk limit | [optional] [readonly] 
**MaintenanceRate** | **string** | The maintenance margin rate of the first tier of risk limit sheet | [optional] [readonly] 
**Value** | **string** | Position value calculated in settlement currency | [optional] [readonly] 
**Margin** | **string** | Position margin | [optional] 
**EntryPrice** | **string** | Entry price | [optional] [readonly] 
**LiqPrice** | **string** | Liquidation price | [optional] [readonly] 
**MarkPrice** | **string** | Current mark price | [optional] [readonly] 
**InitialMargin** | **string** | The initial margin occupied by the position, applicable to the portfolio margin account | [optional] [readonly] 
**MaintenanceMargin** | **string** | Maintenance margin required for the position, applicable to portfolio margin account | [optional] [readonly] 
**UnrealisedPnl** | **string** | Unrealized PNL | [optional] [readonly] 
**RealisedPnl** | **string** | Realized PnL | [optional] [readonly] 
**PnlPnl** | **string** | Realized PNL - Position P/L | [optional] [readonly] 
**PnlFund** | **string** | Realized PNL - Funding Fees | [optional] [readonly] 
**PnlFee** | **string** | Realized PNL - Transaction Fees | [optional] [readonly] 
**HistoryPnl** | **string** | Total realized PnL from closed positions | [optional] [readonly] 
**LastClosePnl** | **string** | PNL of last position close | [optional] [readonly] 
**RealisedPoint** | **string** | Realized POINT PNL | [optional] [readonly] 
**HistoryPoint** | **string** | History realized POINT PNL | [optional] [readonly] 
**AdlRanking** | **int32** | Ranking of auto deleveraging, a total of 1-5 grades, &#x60;1&#x60; is the highest, &#x60;5&#x60; is the lowest, and &#x60;6&#x60; is the special case when there is no position held or in liquidation | [optional] [readonly] 
**PendingOrders** | **int32** | Current pending order quantity | [optional] [readonly] 
**CloseOrder** | Pointer to [**PositionCloseOrder**](Position_close_order.md) |  | [optional] 
**Mode** | **string** | Position mode, including:  - &#x60;single&#x60;: One-way Mode - &#x60;dual_long&#x60;: Long position in Hedge Mode - &#x60;dual_short&#x60;: Short position in Hedge Mode | [optional] 
**CrossLeverageLimit** | **string** | Cross margin leverage (valid only when &#x60;leverage&#x60; is 0) | [optional] 
**UpdateTime** | **int64** | Last update time | [optional] [readonly] 
**UpdateId** | **int64** | Update ID. The value increments by 1 each time the position is updated | [optional] [readonly] 
**OpenTime** | **int64** | First Open Time | [optional] 
**RiskLimitTable** | **string** | Risk limit table ID | [optional] [readonly] 
**AverageMaintenanceRate** | **string** | Average maintenance margin rate | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


