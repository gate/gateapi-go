# SpotPricePutOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Order type，default to &#x60;limit&#x60;  - limit : Limit Order - market : Market Order | [optional] [default to TYPE_LIMIT]
**Side** | **string** | Order side  - buy: buy side - sell: sell side | 
**Price** | **string** | Order price | 
**Amount** | **string** | Trading quantity, refers to the trading quantity of the trading currency, i.e., the currency that needs to be traded, for example, the quantity of BTC in BTC_USDT. | 
**Account** | **string** | Trading account type. Unified account must be set to &#x60;unified&#x60;  - normal: spot trading - margin: margin trading - unified: unified account  | [default to ACCOUNT_NORMAL]
**TimeInForce** | **string** | time_in_force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only  | [optional] [default to TIME_IN_FORCE_GTC]
**AutoBorrow** | **bool** | Whether to borrow coins automatically | [optional] [default to false]
**AutoRepay** | **bool** | Whether to repay the loan automatically | [optional] [default to false]
**Text** | **string** | The source of the order, including: - web: Web - api: API call - app: Mobile app | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


