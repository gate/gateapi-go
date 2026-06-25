# ContractStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **int64** | Stat timestamp | [optional] 
**LsrTaker** | **float64** | Long/short taker ratio | [optional] 
**LsrAccount** | **float64** | Long/short position user ratio | [optional] 
**LongLiqSize** | **string** | Long liquidation size (contracts) | [optional] 
**LongLiqAmount** | **float64** | Long liquidation amount (base currency) | [optional] 
**LongLiqUsd** | **float64** | Long liquidation volume (quote currency) | [optional] 
**LongLiqUsdNew** | **float64** | Long liquidations in quote currency; USDT settlement: long_liq_size × multiplier × mark price | [optional] 
**ShortLiqSize** | **string** | Short liquidation size (contracts) | [optional] 
**ShortLiqAmount** | **float64** | Short liquidation amount (base currency) | [optional] 
**ShortLiqUsd** | **float64** | Short liquidation volume (quote currency) | [optional] 
**ShortLiqUsdNew** | **float64** | Short liquidations in quote currency; USDT settlement: short_liq_size × multiplier × mark price | [optional] 
**OpenInterest** | **string** | Total open interest size (contracts) | [optional] 
**OpenInterestUsd** | **float64** | Total open interest volume (quote currency) | [optional] 
**TopLsrAccount** | **float64** | Top trader long/short account ratio | [optional] 
**TopLsrSize** | **string** | Top trader long/short position ratio | [optional] 
**MarkPrice** | **float64** | Mark price | [optional] 
**TopLongSize** | **string** | Top long open interest (contracts) | [optional] 
**TopShortSize** | **string** | Top short open interest (contracts) | [optional] 
**LongTakerSize** | **string** | Long taker trade volume (contracts) | [optional] 
**ShortTakerSize** | **string** | Short taker trade volume (contracts) | [optional] 
**TopLongAccount** | **int64** | Number of top long accounts (large holders) | [optional] 
**TopShortAccount** | **int64** | Number of top short accounts (large holders) | [optional] 
**LongUsers** | **string** | Number of users holding long positions | [optional] 
**ShortUsers** | **string** | Number of users holding short positions | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


