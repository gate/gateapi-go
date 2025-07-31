# UnifiedCollateralReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollateralType** | **int32** | User-set collateral mode: 0(all)-All currencies as collateral, 1(custom)-Custom currencies as collateral. When collateral_type is 0(all), enable_list and disable_list parameters are invalid | [optional] 
**EnableList** | **[]string** | Currency list, where collateral_type&#x3D;1(custom) indicates the addition logic | [optional] 
**DisableList** | **[]string** | Disable list, indicating the disable logic | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


