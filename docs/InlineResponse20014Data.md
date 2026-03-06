# InlineResponse20014Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSelf** | **bool** | Whether self | [optional] 
**UserTimest** | **string** | User registration time (formatted string) | [optional] 
**CounterpartiesNum** | **int32** | Number of counterparties | [optional] 
**EmailVerified** | **string** | Whether email is verified | [optional] 
**Verified** | **string** | Whether KYC verification is completed | [optional] 
**HasPhone** | **string** | Whether phone is bound | [optional] 
**UserName** | **string** | Username | [optional] 
**UserNote** | **string** | User note information | [optional] 
**CompleteTransactions** | **string** | Total completed orders | [optional] 
**PaidTransactions** | **string** | Number of completed buy orders | [optional] 
**AcceptedTransactions** | **string** | Number of completed sell orders | [optional] 
**TransactionsUsedTime** | **string** | Average time to confirm receipt | [optional] 
**CancelledUsedTimeMonth** | **string** | Cancellation time in last 30 days | [optional] 
**CompleteTransactionsMonth** | **string** | Number of completed orders in last 30 days | [optional] 
**CompleteRateMonth** | **float32** | Completion rate in last 30 days | [optional] 
**OrdersBuyRateMonth** | **float32** | Buy order ratio in last 30 days | [optional] 
**IsBlack** | **int32** | Whether blocked | [optional] 
**IsFollow** | **int32** | Whether following | [optional] 
**HaveTraded** | **int32** | Whether traded with self | [optional] 
**BizUid** | **string** | Encrypted UID | [optional] 
**BlueVip** | **int32** | Blue V Crown Shield | [optional] 
**WorkStatus** | **int32** | Merchant work status | [optional] 
**RegistrationDays** | **int32** | Registration days | [optional] 
**FirstTradeDays** | **int32** | Days since first trade | [optional] 
**NeedReplenish** | **int32** | Whether margin replenishment is needed | [optional] 
**MerchantInfo** | [**InlineResponse20014DataMerchantInfo**](inline_response_200_14_data_merchant_info.md) |  | [optional] 
**OnlineStatus** | **int32** | Merchant online status | [optional] 
**WorkHours** | Pointer to [**map[string]interface{}**](.md) | Merchant online status details | [optional] 
**TransactionsMonth** | **float32** | 30-day transaction volume | [optional] 
**TransactionsAll** | **float32** | Total transaction volume | [optional] 
**TradeVersatile** | **bool** | Single user or composite user | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


