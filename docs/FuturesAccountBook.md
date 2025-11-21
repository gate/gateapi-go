# FuturesAccountBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **float64** | Change time | [optional] 
**Change** | **string** | Change amount | [optional] 
**Balance** | **string** | Balance after change | [optional] 
**Type** | **string** | Change types:  - dnw: Deposit and withdrawal - pnl: Profit and loss from position reduction - fee: Trading fees - refr: Referrer rebates - fund: Funding fees - point_dnw: Point card deposit and withdrawal - point_fee: Point card trading fees - point_refr: Point card referrer rebates - bonus_offset: Trial fund deduction | [optional] 
**Text** | **string** | Comment | [optional] 
**Contract** | **string** | Futures contract, the field is only available for data after 2023-10-30 | [optional] 
**TradeId** | **string** | trade id | [optional] 
**Id** | **string** | Account change record ID | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


