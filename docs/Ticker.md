# Ticker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyPair** | **string** | Currency pair | [optional] 
**Last** | **string** | Last trading price | [optional] 
**LowestAsk** | **string** | Recent lowest ask | [optional] 
**LowestSize** | **string** | Latest seller&#39;s lowest price quantity; not available for batch queries; available for single queries, empty if no data | [optional] 
**HighestBid** | **string** | Recent highest bid | [optional] 
**HighestSize** | **string** | Latest buyer&#39;s highest price quantity; not available for batch queries; available for single queries, empty if no data | [optional] 
**ChangePercentage** | **string** | 24h price change percentage (negative for decrease, e.g., -7.45) | [optional] 
**ChangeUtc0** | **string** | UTC+0 timezone, 24h price change percentage, negative for decline (e.g., -7.45) | [optional] 
**ChangeUtc8** | **string** | UTC+8 timezone, 24h price change percentage, negative for decline (e.g., -7.45) | [optional] 
**BaseVolume** | **string** | Base currency trading volume in the last 24h | [optional] 
**QuoteVolume** | **string** | Quote currency trading volume in the last 24h | [optional] 
**High24h** | **string** | 24h High | [optional] 
**Low24h** | **string** | 24h Low | [optional] 
**EtfNetValue** | **string** | ETF net value | [optional] 
**EtfPreNetValue** | Pointer to **string** | ETF net value at previous rebalancing point | [optional] 
**EtfPreTimestamp** | Pointer to **int64** | ETF previous rebalancing time | [optional] 
**EtfLeverage** | Pointer to **string** | ETF current leverage | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


