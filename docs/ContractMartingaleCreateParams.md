# ContractMartingaleCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestAmount** | **string** | Margin allocated; the server converts it to initial contract size using live contract price, contract multiplier, and minimum lot size. | 
**PriceDeviation** | **string** |  | 
**MaxOrders** | **int32** |  | 
**TakeProfitRatio** | **string** |  | 
**Direction** | [**ContractMartingaleDirection**](ContractMartingaleDirection.md) |  | 
**Leverage** | **string** |  | 
**StopLossPrice** | **string** | Legacy field name. The AIHub &#x60;contract_martingale&#x60; creation path does not map this field today; follow contract martingale rules from the underlying API. MCP tooling must match bot-service behavior. | [optional] 
**ProfitSharingRatio** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


