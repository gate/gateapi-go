# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Currency symbol | [optional] 
**Name** | **string** | Currency name | [optional] 
**Delisted** | **bool** | Whether currency is de-listed | [optional] 
**WithdrawDisabled** | **bool** | Whether currency&#39;s withdrawal is disabled (deprecated) | [optional] 
**WithdrawDelayed** | **bool** | Whether currency&#39;s withdrawal is delayed (deprecated) | [optional] 
**DepositDisabled** | **bool** | Whether currency&#39;s deposit is disabled (deprecated) | [optional] 
**TradeDisabled** | **bool** | Whether currency&#39;s trading is disabled | [optional] 
**FixedRate** | **string** | Fixed fee rate. Only for fixed rate currencies, not valid for normal currencies | [optional] 
**Chain** | **string** | The main chain corresponding to the coin | [optional] 
**Chains** | [**[]SpotCurrencyChain**](SpotCurrencyChain.md) | All links corresponding to coins | [optional] 
**TotalSupply** | **string** | Total supply | [optional] 
**MarketCap** | **string** | Market cap | [optional] 
**Category** | **[]string** | 币种分类  - stocks: 股票 - metals: 金属 - indices: 指数 - forex: 外汇 - commodities: 大宗商品 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


