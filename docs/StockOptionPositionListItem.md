# StockOptionPositionListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Full option contract symbol. | [optional] 
**SymbolDisplay** | **string** | Option contract name for display. | [optional] 
**Underlying** | **string** | Option underlying symbol. | [optional] 
**Side** | **string** | Position side. - long: Long position. - short: Short position. | [optional] 
**OpenTime** | **int64** | Position opening time as a Unix timestamp in seconds; 0 when unavailable. | [optional] 
**Quantity** | **string** | Position quantity. | [optional] 
**Available** | **string** | Available quantity; currently equal to quantity. | [optional] 
**Frozen** | **string** | Frozen quantity; may be empty when not calculated separately. | [optional] 
**AvgEntryPrice** | **string** | Average entry price. | [optional] 
**LastPrice** | **string** | Last traded price; 0 when no valid value is available. | [optional] 
**MarkPrice** | **string** | Mark price; 0 when no valid value is available. | [optional] 
**PreviousClose** | **string** | Previous closing price. If no valid positive value exists, this field may be empty and daily PnL calculation is skipped. | [optional] 
**MarketValue** | **string** | Market value &#x3D; valuation price × signed position quantity × contract multiplier. The valuation price is selected by &#x60;price_type&#x60; (&#x60;latest&#x60; uses last_price; &#x60;mark&#x60; uses mark_price). | [optional] 
**UnrealizedPnl** | **string** | Unrealized PnL &#x3D; (valuation price - average entry price) × signed position quantity × contract multiplier. Returns 0 if the average entry price or quantity is invalid. The valuation price follows the same rules as market_value. | [optional] 
**TodayPnl** | **string** | Daily PnL for this contract. Requires a valid previous_close; 0 if unavailable. The valuation price follows the same rules as market_value. | [optional] 
**TodayBuyQty** | **string** | Quantity bought today (user&#39;s daily trading totals). | [optional] 
**TodaySellQty** | **string** | Quantity sold today (user&#39;s daily trading totals). | [optional] 
**TodayBuyValue** | **string** | Amount bought today (user&#39;s daily trading totals). | [optional] 
**TodaySellValue** | **string** | Amount sold today (user&#39;s daily trading totals). | [optional] 
**TradeStatus** | [**StockOptionTradeStatus**](StockOptionTradeStatus.md) |  | [optional] 
**TradeMode** | [**StockOptionTradeMode**](StockOptionTradeMode.md) |  | [optional] 
**Multiplier** | **string** | Contract multiplier. | [optional] 
**PricePrecision** | **int32** | Number of decimal places for prices. | [optional] 
**MinOrderVolume** | **string** | Minimum order quantity. | [optional] 
**MaxOrderVolume** | **string** | Maximum order quantity. | [optional] 
**StepOrderVolume** | **string** | Order quantity increment. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


