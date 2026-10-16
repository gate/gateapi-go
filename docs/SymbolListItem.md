# SymbolListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Symbol | [optional] 
**Exchange** | **string** | Exchange. - us: US stocks. - hk: Hong Kong stocks. - kr: Korean stocks. - jp: Japanese stocks. | [optional] 
**ExchangeDesc** | **string** | Exchange description | [optional] 
**QuoteCurrency** | **string** | Quote currency | [optional] 
**QuoteCurrencyPrecision** | **int32** | Quote currency precision | [optional] 
**FxRate** | **string** | Quote currency to USD exchange rate | [optional] 
**SymbolDesc** | **string** | Symbol description | [optional] 
**Category** | **string** | Symbol category. - CS: Common stock. - ETF: Exchange-traded funds. - ADRC, ADR: Depositary receipts for foreign companies listed in the U.S. - ETV: Exchange-traded products. - PFD: Preferred stock. - ETS: Exchange-traded securities. - ETN: Exchange-traded notes. - FUND: Funds. | [optional] 
**AssetType** | **string** | Asset type. - STOCK: Stock. - ETF: Exchange-traded fund. | [optional] 
**TradeStatus** | **string** | Trading status. - pre_market: Pre-market. - open: Regular trading session. - post_market: Post-market. | [optional] 
**TradeMode** | **int32** | Current session trading mode. - 0: Trading disabled. - 1: Buy only. - 2: Sell only. - 4: Buy and sell supported. - 5: Trading suspended. - 6: Trading suspended. - 7: Trading suspended. - 8: Trading suspended. | [optional] 
**OrderFillTiming** | **int32** | Order fill timing (1&#x3D;immediate, 2&#x3D;after pre-market opens, 3&#x3D;after regular session opens) | [optional] 
**IconLink** | **string** | Icon URL | [optional] 
**QuoteCurrencySymbol** | **string** | Quote currency symbol | [optional] 
**PricePrecision** | **int32** | Price precision | [optional] 
**VolumePrecision** | **int32** | Quantity precision | [optional] 
**IsIpo** | **bool** | Whether it is an IPO symbol | [optional] 
**IpoPrice** | **string** | IPO price | [optional] 
**SellPriceProtection** | **string** | Sell price protection rate | [optional] 
**BuyPriceProtection** | **string** | Buy price protection rate | [optional] 
**SymbolDescs** | [**[]I18nTxt**](I18nTxt.md) | Multilingual symbol description | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


