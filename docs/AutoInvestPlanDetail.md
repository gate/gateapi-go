# AutoInvestPlanDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Plan ID | 
**Version** | **int64** | PlanVersion | [optional] 
**Name** | **string** | Plan name | 
**CreateTime** | **int64** | Creation time (Unix timestamp) | 
**UpdateTime** | **int64** | Update time (Unix timestamp) | 
**UserId** | **int64** | User ID | 
**Money** | **string** | Quote Currency | 
**Amount** | **string** | Per PeriodInvestment amount | 
**PeriodType** | **string** | Cycle type（e.g., monthly） | 
**PeriodDay** | **int64** | Cycle day | 
**PeriodHour** | **int64** | CycleHours | 
**Portfolio** | [**[]AutoInvestPortfolioItem**](AutoInvestPortfolioItem.md) | InvestmentPortfolio | 
**NextTime** | **int64** | Next execution time (Unix timestamp) | 
**Period** | **int64** | Executed periods | 
**FundSource** | **string** | Fund source（spot/earn） | 
**FundFlow** | **string** | Fund flow direction (auto_invest/earn) | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


