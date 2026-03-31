# AutoInvestPlanCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanName** | **string** | Plan name。Length: 0-50 characters | [optional] 
**PlanDes** | **string** | Plan description | [optional] 
**PlanMoney** | **string** | Pricing currency，SupportUSDT，BTC | 
**PlanAmount** | **string** | Per PeriodAuto InvestAmount，Must &gt; 0，and not exceedThePricing currencyConfigurationSingleMaximumAmount | 
**PlanPeriodType** | **string** | Enum: daily, weekly, biweekly, monthly, hourly, 4-hourly | 
**PlanPeriodDay** | **int64** | Cycle day. For monthly: day of month (1-30); For weekly/biweekly: day of week (1-7, 1&#x3D;Monday); For daily/hourly/4-hourly: ignored | 
**PlanPeriodHour** | **int64** | Execution hourAuto Invest 0-23 | 
**Items** | [**[]AutoInvestPlanCreateItems**](AutoInvestPlanCreate_items.md) | Investment portfolio, asset cannot be repeated; Sum of all items&#39; ratios must be 100 | 
**FundSource** | **string** | Fund source: spot or earn, default: spot | [optional] 
**FundFlow** | **string** | Fund flow direction: auto_invest or earn, default: auto_invest | [optional] 
**Type** | **int64** | 0 Normal creation, 1 QuickInvestment | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


