# P2pMerchantUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSelf** | **bool** | Whether self | [optional] 
**UserTimest** | **string** | User registration time (formatted string) | [optional] 
**CounterpartiesNum** | **int32** | Number of counterparties | [optional] 
**EmailVerified** | **string** | Whether email is verified. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**Verified** | **string** | Whether KYC is completed. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**HasPhone** | **string** | Whether a phone number is bound. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
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
**IsBlack** | **int32** | Whether the user is blocked. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**IsFollow** | **int32** | Whether you follow this user. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**HaveTraded** | **int32** | Whether you have traded with this user before. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**BizUid** | **string** | Encrypted UID | [optional] 
**BlueVip** | **int32** | Blue V Crown Shield | [optional] 
**WorkStatus** | **int32** | Merchant work status | [optional] 
**RegistrationDays** | **int32** | Registration days | [optional] 
**FirstTradeDays** | **int32** | Days since first trade | [optional] 
**NeedReplenish** | **int32** | Whether additional margin is required. &#x60;1&#x60;: yes; &#x60;0&#x60;: no. | [optional] 
**MerchantInfo** | [**P2pMerchantMarketInfo**](P2pMerchantMarketInfo.md) |  | [optional] 
**OnlineStatus** | **int32** | Merchant online status: &#x60;1&#x60; online; &#x60;0&#x60; offline. | [optional] 
**WorkHours** | Pointer to [**map[string]interface{}**](.md) | Merchant online status details | [optional] 
**TransactionsMonth** | **float32** | 30-day transaction volume | [optional] 
**TransactionsAll** | **float32** | Total transaction volume | [optional] 
**TradeVersatile** | **bool** | Single user or composite user | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


