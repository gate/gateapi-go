# TransactionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | **string** | Asset | [optional] 
**Symbol** | **string** | Symbol | [optional] 
**SymbolDisplay** | **string** | Symbol display name | [optional] 
**Type** | **string** | Transaction type. - deposit: Funds transfer in. - withdraw: Funds transfer out. - fee: Trading fee. - dividend: Dividend payout. - sell: Stock sale credit. - buy: Stock purchase debit. - award: Airdrop reward. - stock_transfer_in: Stock transfer in. - stock_transfer_out: Stock transfer out. | [optional] 
**TypeDesc** | **string** | Transaction type description | [optional] 
**Change** | **string** | Change amount | [optional] 
**Balance** | **string** | Balance after change | [optional] 
**RefId** | **string** | Business idempotent ID | [optional] 
**Time** | **int64** | Unix timestamp (seconds) | [optional] 
**UnitText** | **string** | Unit display text | [optional] 
**Detail** | **map[string]map[string]interface{}** | Business details | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


