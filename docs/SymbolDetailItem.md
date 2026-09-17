# SymbolDetailItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Symbol | [optional] 
**Exchange** | **string** | Exchange, supports us, hk, kr, and jp | [optional] 
**ExchangeDesc** | **string** | Exchange description | [optional] 
**QuoteCurrency** | **string** | Quote currency | [optional] 
**QuoteCurrencyPrecision** | **int32** | Quote currency precision | [optional] 
**FxRate** | **string** | Quote currency to USD exchange rate | [optional] 
**SymbolDesc** | **string** | Symbol description | [optional] 
**Category** | **string** | Symbol category. - CS: Common stock. - ETF: Exchange-traded funds. - ADRC, ADR: Depositary receipts for foreign companies listed in the U.S. - ETV: Exchange-traded products. - PFD: Preferred stock. - ETS: Exchange-traded securities. - ETN: Exchange-traded notes. - FUND: Funds. | [optional] 
**AssetType** | **string** | Asset type. - STOCK: Stock. - ETF: Exchange-traded fund. | [optional] 
**SettlementCurrency** | **string** | Settlement currency | [optional] 
**MaxOrderVolume** | **string** | Maximum order quantity | [optional] 
**StepOrderVolume** | **string** | Order step size | [optional] 
**MinOrderVolume** | **string** | Minimum order quantity | [optional] 
**PricePrecision** | **int32** | Price precision | [optional] 
**VolumePrecision** | **int32** | Quantity precision | [optional] 
**IsIpo** | **bool** | Whether it is an IPO symbol | [optional] 
**IpoPrice** | **string** | IPO price | [optional] 
**PriceProtection** | **string** | Price protection range | [optional] 
**SellPriceProtection** | **string** | Sell price protection rate | [optional] 
**BuyPriceProtection** | **string** | Buy price protection rate | [optional] 
**SlippageRate** | **string** | Slippage | [optional] 
**CommissionRate** | **string** | Fee Rate | [optional] 
**TradeStatus** | **string** | Trading status. - pre_market: Pre-market. - open: Regular trading session. - post_market: Post-market. - closed: Market closed. - gt_lp: GT LP session. | [optional] 
**TradeMode** | **int32** | Current session trading mode. - 0: Trading disabled. - 1: Buy only. - 2: Sell only. - 4: Buy and sell supported. | [optional] 
**OrderFillTiming** | **int32** | Order fill timing (1&#x3D;immediate, 2&#x3D;after pre-market opens, 3&#x3D;after regular session opens) | [optional] 
**SymbolDescs** | [**[]SymbolDetailItemSymbolDescs**](SymbolDetailItem_symbol_descs.md) | Multilingual symbol description | [optional] 
**IconLink** | **string** | Icon URL | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


