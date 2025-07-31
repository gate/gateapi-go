# SpotPriceTriggeredOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trigger** | [**SpotPriceTrigger**](SpotPriceTrigger.md) |  | 
**Put** | [**SpotPricePutOrder**](SpotPricePutOrder.md) |  | 
**Id** | **int64** | Auto order ID | [optional] [readonly] 
**User** | **int32** | User ID | [optional] [readonly] 
**Market** | **string** | Market | 
**Ctime** | **int64** | Created time | [optional] [readonly] 
**Ftime** | **int64** | End time | [optional] [readonly] 
**FiredOrderId** | **int64** | ID of the order created after trigger | [optional] [readonly] 
**Status** | **string** | Status  - open: Running - cancelled: Manually cancelled - finish: Successfully completed - failed: Failed to execute - expired: Expired | [optional] [readonly] 
**Reason** | **string** | Additional description of how the order was completed | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


