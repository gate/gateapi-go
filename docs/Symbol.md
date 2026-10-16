# Symbol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Unique trading pair identifier in the format ExchangeType_BusinessType_Base_Counter. Stocks and ETFs use the market as ExchangeType (US/HK/KR/JP) and STOCK or ETF as the business type, for example US_STOCK_AAPL_USDT, HK_STOCK_00700_USDT, or US_ETF_SPY_USDT. | 
**ExchangeType** | **string** | Exchange type (BINANCE / OKX / GATE / BYBIT / KRAKEN / HYPERLIQUID / DERIBIT / LIGHTER). Stocks and ETFs return the listing market (US / HK / KR / JP) | 
**BusinessType** | **string** | Business type (SPOT: spot / FUTURE: futures / MARGIN: margin / STOCK: stocks / ETF: exchange-traded funds) | 
**State** | **string** | Status (&#x60;live&#x60; running / &#x60;suspend&#x60; paused). | 
**MinSize** | **string** | Minimum order quantity | 
**MinNotional** | **string** | Minimum Order Value | 
**LotSize** | **string** | Quantity Step | 
**TickSize** | **string** | Price Step | 
**MaxNumOrders** | **string** | maximumopen orderamount | 
**MaxMarketSize** | **string** | Maximum Market Order Quantity | 
**MaxLimitSize** | **string** | Maximum order quantity for limit orders. | 
**ContractSize** | **string** | Contract multiplier (deprecated; quantity is used uniformly) | 
**LiquidationFee** | **string** | Liquidation Fee Rate | 
**DelistTime** | **string** | Millisecond timestamp; &#x60;0&#x60; means not delisted. | 
**SupportRpi** | **string** | Whether RPI order placement is supported (true if supported; false otherwise) | [optional] 
**SupportCross** | **string** | Whether cross-margin order placement is supported (true if supported; false otherwise) | [optional] 
**Tag** | **string** | Tag information. Hyperliquid returns DEX information, stocks and ETFs return the security description, and other businesses return an empty string | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


