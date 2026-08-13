# SpotPovOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Order ID | [readonly] 
**CurrencyPair** | **string** | Currency pair | [readonly] 
**Side** | **string** | Buy or sell order | [readonly] 
**Amount** | **string** | Trade amount | [readonly] 
**ParticipationRate** | **int32** | Target participation rate as a percentage. Allowed values: 5, 10, 20, and 40 | [readonly] 
**Ttl** | **string** | Time to live. Valid values: 1h, 6h, 12h, 1d, 2d, 3d, 4d, 5d, 6d, and 7d | [readonly] 
**LimitPrice** | **string** | Limit price. If omitted, the market price is used | [optional] [readonly] 
**TriggerPrice** | **string** | Trigger price. If omitted, the order is triggered immediately | [optional] [readonly] 
**Status** | **string** | Order status  - CREATED: Created - CANCELING: Canceling - RUNNING: Running - COMPLETED: Completed - EXPIRED: Expired - TERMINATED: Terminated | [readonly] 
**TerminatedAs** | **string** | Order termination reason code | [optional] [readonly] 
**StartTimeMs** | **int64** | Order execution start time in milliseconds | [optional] [readonly] 
**EndTimeMs** | **int64** | Order execution end time in milliseconds | [optional] [readonly] 
**ExpireTimeMs** | **int64** | Order expiration time in milliseconds | [optional] [readonly] 
**CreateTimeMs** | **int64** | Creation time of order (in milliseconds) | [readonly] 
**UpdateTimeMs** | **int64** | Last modification time of order (in milliseconds) | [optional] [readonly] 
**Text** | **string** | Order custom information. Users can set custom ID with this field. Custom fields must meet the following conditions:  1. Must start with &#x60;t-&#x60; 2. Excluding &#x60;t-&#x60;, length cannot exceed 28 bytes 3. Can only contain numbers, letters, underscore(_), hyphen(-) or dot(.)  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


