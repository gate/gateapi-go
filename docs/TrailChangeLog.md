# TrailChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | **int64** | Update time | [optional] [readonly] 
**Amount** | **string** | Trading quantity in contracts, positive for buy, negative for sell | [optional] [readonly] 
**IsGte** | **bool** | true: activate when market price &gt;&#x3D; activation price, false: &lt;&#x3D; activation price | [optional] [readonly] 
**ActivationPrice** | **string** | Activation price, 0 means trigger immediately | [optional] [readonly] 
**PriceType** | **int32** | Activation price type: 0-unknown, 1-latest price, 2-index price, 3-mark price | [optional] [readonly] 
**PriceOffset** | **string** | Callback ratio or price distance, e.g., &#x60;0.1&#x60; or &#x60;0.1%&#x60; | [optional] [readonly] 
**IsCreate** | **bool** | true - order creation, false - order modification | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


