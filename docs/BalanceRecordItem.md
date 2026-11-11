# BalanceRecordItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Balance record primary key ID | [optional] 
**UserId** | **int32** | User ID | [optional] 
**Change** | **string** | Cash amount change. Positive values are credits; negative values are debits. | [optional] 
**Balance** | **string** | Cash balance snapshot after the change | [optional] 
**StockChange** | **string** | Stock quantity change, parsed from detail.amount for stock quantity records; may be an empty string for other records | [optional] 
**Type** | **string** | Raw record type. Common values include: - deposit: Funds transferred in. - withdraw: Funds transferred out. - fee: Trading fee. - dividend: Dividend payment. - sell: Proceeds from selling stocks. - buy: Payment for buying stocks. - award: Airdrop reward. - stock_transfer_in: Stocks transferred in. - stock_transfer_out: Stocks transferred out. | [optional] 
**Business** | **string** | Business identifier | [optional] 
**BusinessId** | **int64** | Related business ID | [optional] 
**Symbol** | **string** | Symbol | [optional] 
**Detail** | **string** | Raw business details as a JSON string; returned only when include_detail&#x3D;true is requested | [optional] 
**Version** | **int32** | Record version | [optional] 
**TimeCreate** | **int64** | Business event time (Unix timestamp in seconds) | [optional] 
**CreateTime** | **int64** | Record creation time (Unix timestamp in seconds) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


