# SymbolListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Symbol | [optional] 
**Exchange** | **string** | Exchange, supports us, hk, and kr | [optional] 
**ExchangeDesc** | **string** | Exchange description | [optional] 
**QuoteCurrency** | **string** | Quote currency | [optional] 
**QuoteCurrencyPrecision** | **int32** | Quote currency precision | [optional] 
**FxRate** | **string** | Quote currency to USD exchange rate | [optional] 
**SymbolDesc** | **string** | Symbol description | [optional] 
**Category** | **string** | Category | [optional] 
**TradeStatus** | **string** | Trading status. - pre_market: Pre-market. - open: Regular trading session. - post_market: Post-market. - closed: Market closed. - gt_lp: GT LP session. | [optional] 
**TradeMode** | **int32** | Current session trading mode. - 0: Trading disabled. - 1: Buy only. - 2: Sell only. - 4: Buy and sell supported. | [optional] 
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


