# OrderListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | [optional] 
**Symbol** | **string** | Symbol | [optional] 
**Exchange** | **string** | Exchange, supports us, hk, and kr | [optional] 
**QuoteCurrency** | **string** | Quote currency | [optional] 
**FxRate** | **string** | Quote currency to USD exchange rate | [optional] 
**SymbolDesc** | **string** | Symbol description | [optional] 
**TradeStatus** | **string** | Trading status. - pre_market: Pre-market. - open: Regular trading session. - post_market: Post-market. - closed: Market closed. - gt_lp: GT LP session. | [optional] 
**TradeMode** | **int32** | Current session trading mode. - 0: Trading disabled. - 1: Buy only. - 2: Sell only. - 4: Buy and sell supported. | [optional] 
**PriceType** | **string** | Price type (market &#x3D; market order, limit &#x3D; limit order) | [optional] 
**Side** | **int32** | Side (1&#x3D;sell, 2&#x3D;buy) | [optional] 
**Status** | **int32** | Order status | [optional] 
**Volume** | **string** | Order quantity | [optional] 
**FillVolume** | **string** | Trading size | [optional] 
**Price** | **string** | Order price | [optional] 
**TimeSetup** | **int64** | Order creation time (Unix timestamp, seconds) | [optional] 
**TimeUpdate** | **int64** | Order update time (Unix timestamp, seconds) | [optional] 
**MaxOrderVolume** | **string** | Maximum order quantity | [optional] 
**StepOrderVolume** | **string** | Order step size | [optional] 
**MinOrderVolume** | **string** | Minimum order quantity | [optional] 
**PricePrecision** | **int32** | Price precision | [optional] 
**PriceProtection** | **string** | Price protection range | [optional] 
**SellPriceProtection** | **string** | Sell price protection rate | [optional] 
**BuyPriceProtection** | **string** | Buy price protection rate | [optional] 
**CommissionRate** | **string** | Fee Rate | [optional] 
**SlippageRate** | **string** | Slippage | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


