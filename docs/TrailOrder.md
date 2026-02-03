# TrailOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Order ID | [optional] [readonly] 
**UserId** | **int64** | User ID | [optional] [readonly] 
**User** | **int64** | User ID | [optional] [readonly] 
**Contract** | **string** | Contract name | [optional] 
**Settle** | **string** | Settle currency | [optional] 
**Amount** | **string** | Trading quantity in contracts, positive for buy, negative for sell | [optional] 
**IsGte** | **bool** | true: activate when market price &gt;&#x3D; activation price, false: &lt;&#x3D; activation price | [optional] 
**ActivationPrice** | **string** | Activation price, 0 means trigger immediately | [optional] 
**PriceType** | **int32** | Activation price type: 0-unknown, 1-latest price, 2-index price, 3-mark price | [optional] 
**PriceOffset** | **string** | Callback ratio or price distance, e.g., &#x60;0.1&#x60; or &#x60;0.1%&#x60; | [optional] 
**Text** | **string** | Custom field | [optional] 
**ReduceOnly** | **bool** | Reduce Position Only | [optional] 
**PositionRelated** | **bool** | Whether bound to position | [optional] 
**CreatedAt** | **int64** | Created time | [optional] [readonly] 
**ActivatedAt** | **int64** | Activation time | [optional] [readonly] 
**FinishedAt** | **int64** | End time | [optional] [readonly] 
**CreateTime** | **int64** | Created time | [optional] [readonly] 
**ActiveTime** | **int64** | Activation time | [optional] [readonly] 
**FinishTime** | **int64** | End time | [optional] [readonly] 
**Reason** | **string** | End reason | [optional] [readonly] 
**SuborderText** | **string** | Sub-order text field | [optional] [readonly] 
**IsDualMode** | **bool** | Whether dual position mode when creating order | [optional] [readonly] 
**TriggerPrice** | **string** | Trigger price | [optional] [readonly] 
**SuborderId** | **int64** | Sub-order ID | [optional] [readonly] 
**SideLabel** | **string** | Order direction label: long/short/open long/open short/close long/close short | [optional] [readonly] 
**OriginalStatus** | **int32** | Order status | [optional] [readonly] 
**Status** | **string** | Simplified order status: open/finished | [optional] [readonly] 
**PositionSideOutput** | **string** | Same as side_label, client requires consistency with other order types | [optional] [readonly] 
**UpdatedAt** | **int64** | Update time | [optional] [readonly] 
**ExtremumPrice** | **string** | Extremum price | [optional] [readonly] 
**StatusCode** | **string** | Status code value | [optional] [readonly] 
**CreatedAtPrecise** | **string** | Creation time (high precision, seconds.microseconds format) | [optional] [readonly] 
**FinishedAtPrecise** | **string** | End time (high precision, seconds.microseconds format) | [optional] [readonly] 
**ActivatedAtPrecise** | **string** | Activation time (high precision, seconds.microseconds format) | [optional] [readonly] 
**StatusLabel** | **string** | Status internationalization label (translated status text) | [optional] [readonly] 
**PosMarginMode** | **string** | Position margin mode: isolated/cross | [optional] [readonly] 
**PositionMode** | **string** | Position mode: single, dual, and dual_plus | [optional] [readonly] 
**ErrorLabel** | **string** | Error label | [optional] [readonly] 
**Leverage** | **string** | leverage | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


