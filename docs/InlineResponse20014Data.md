# InlineResponse20014Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSelf** | **bool** | Whether self | 
**UserTimest** | **string** | User registration time (formatted string) | 
**CounterpartiesNum** | **int32** | Number of counterparties | 
**EmailVerified** | **string** | Whether email is verified | 
**Verified** | **string** | Whether KYC verification is completed | 
**HasPhone** | **string** | Whether phone is bound | 
**UserName** | **string** | Username | 
**UserNote** | **string** | User note information | 
**CompleteTransactions** | **string** | Total completed orders | 
**PaidTransactions** | **string** | Number of completed buy orders | 
**AcceptedTransactions** | **string** | Number of completed sell orders | 
**TransactionsUsedTime** | **string** | Average time to confirm receipt | 
**CancelledUsedTimeMonth** | **string** | Cancellation time in last 30 days | 
**CompleteTransactionsMonth** | **string** | Number of completed orders in last 30 days | 
**CompleteRateMonth** | **int32** | Completion rate in last 30 days | 
**OrdersBuyRateMonth** | **int32** | Buy order ratio in last 30 days | 
**IsBlack** | **int32** | Whether blocked | 
**IsFollow** | **int32** | Whether following | 
**HaveTraded** | **int32** | Whether traded with self | 
**BizUid** | **string** | Encrypted UID | 
**BlueVip** | **int32** | Blue V Crown Shield | 
**WorkStatus** | **int32** | Merchant work status | 
**RegistrationDays** | **int32** | Registration days | 
**FirstTradeDays** | **int32** | Days since first trade | 
**NeedReplenish** | **int32** | Whether margin replenishment is needed | 
**MerchantInfo** | [**InlineResponse20014DataMerchantInfo**](inline_response_200_14_data_merchant_info.md) |  | 
**OnlineStatus** | **int32** | Merchant online status | 
**WorkHours** | Pointer to [**map[string]interface{}**](.md) | Merchant online status details | 
**TransactionsMonth** | **int32** | 30-day transaction volume | 
**TransactionsAll** | **int32** | Total transaction volume | 
**TradeVersatile** | **bool** | Single user or composite user | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


