# PartnerDataAggregated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebateAmount** | **string** | Rebate amount as a string for precision. Up to 6 decimal places; trailing zeros removed. | 
**TradeVolume** | **string** | Trading volume as a string for precision. Up to 6 decimal places; trailing zeros removed. | 
**NetFee** | **string** | Net fee as a string for precision. Up to 6 decimal places; trailing zeros removed. | 
**CustomerCount** | **int32** | Customer count (invited users) | 
**TradingUserCount** | Pointer to **string** | Transaction participant count​ (string format, consistent with online JSON serialization) only returns a specific value when business_type&#x3D;0(all), and returns nullfor other business types. | 
**TimeRangeDesc** | **string** | Time range description | 
**BusinessType** | **int32** | Business Type | 
**BusinessTypeDesc** | **string** | Business type description; allowed values: All, Spot, Futures, Alpha, Web3, Perps (DEX), Exchange All, Web3 All, TradFi | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


