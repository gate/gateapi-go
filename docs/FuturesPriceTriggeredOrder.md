# FuturesPriceTriggeredOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initial** | [**FuturesInitialOrder**](FuturesInitialOrder.md) |  | 
**Trigger** | [**FuturesPriceTrigger**](FuturesPriceTrigger.md) |  | 
**Id** | **int64** | Auto order ID | [optional] [readonly] 
**User** | **int32** | User ID | [optional] [readonly] 
**CreateTime** | **float64** | Created time | [optional] [readonly] 
**FinishTime** | **float64** | End time | [optional] [readonly] 
**TradeId** | **int64** | ID of the order created after trigger | [optional] [readonly] 
**Status** | **string** | Order status  - &#x60;open&#x60;: Active - &#x60;finished&#x60;: Finished - &#x60;inactive&#x60;: Inactive, only applies to order take-profit/stop-loss - &#x60;invalid&#x60;: Invalid, only applies to order take-profit/stop-loss | [optional] [readonly] 
**FinishAs** | **string** | Finish status: cancelled - Cancelled; succeeded - Succeeded; failed - Failed; expired - Expired | [optional] [readonly] 
**Reason** | **string** | Additional description of how the order was completed | [optional] [readonly] 
**OrderType** | **string** | Types of take-profit and stop-loss orders, including:  - &#x60;close-long-order&#x60;: Order take-profit/stop-loss, close long position - &#x60;close-short-order&#x60;: Order take-profit/stop-loss, close short position - &#x60;close-long-position&#x60;: Position take-profit/stop-loss, used to close all long positions - &#x60;close-short-position&#x60;: Position take-profit/stop-loss, used to close all short positions - &#x60;plan-close-long-position&#x60;: Position plan take-profit/stop-loss, used to close all or partial long positions - &#x60;plan-close-short-position&#x60;: Position plan take-profit/stop-loss, used to close all or partial short positions  The two types of order take-profit/stop-loss are read-only and cannot be passed in requests | [optional] 
**MeOrderId** | **int64** | Corresponding order ID for order take-profit/stop-loss orders | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


