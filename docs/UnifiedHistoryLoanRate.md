# UnifiedHistoryLoanRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Currency name | [optional] 
**Tier** | **string** | VIP level for the floating rate to be retrieved | [optional] 
**TierUpRate** | **string** | Floating rate corresponding to VIP level | [optional] 
**Rates** | [**[]UnifiedHistoryLoanRateRates**](UnifiedHistoryLoanRate_rates.md) | Historical interest rate information, one data point per hour, array size determined by page and limit parameters from the API request, sorted by time from recent to distant | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


