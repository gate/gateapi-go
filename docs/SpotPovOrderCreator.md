# SpotPovOrderCreator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyPair** | **string** | Currency pair | 
**Side** | **string** | Buy or sell order | 
**Amount** | **string** | Trade amount | 
**ParticipationRate** | **int32** | Target participation rate as a percentage. Valid values: 5, 10, 20, and 40 | 
**Ttl** | **string** | Time to live. Valid values: 1h, 6h, 12h, 1d, 2d, 3d, 4d, 5d, 6d, and 7d | 
**LimitPrice** | **string** | Limit price. If omitted, the market price is used | [optional] 
**TriggerPrice** | **string** | Trigger price. If omitted, the order is triggered immediately | [optional] 
**Text** | **string** | Order custom information. Users can set custom ID with this field. Custom fields must meet the following conditions:  1. Must start with &#x60;t-&#x60; 2. Excluding &#x60;t-&#x60;, length cannot exceed 28 bytes 3. Can only contain numbers, letters, underscore(_), hyphen(-) or dot(.)  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


