# FixedTermProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Product ID | [optional] 
**Name** | **string** | Product name | [optional] 
**Asset** | **string** | Currency | [optional] 
**LockUpPeriod** | **int32** | Lock-up period (in days) | [optional] 
**MinLendAmount** | **string** | Minimum earn amount | [optional] 
**UserMaxLendAmount** | **string** | User maximum earn limit | [optional] 
**TotalLendAmount** | **string** | Platform earn limit | [optional] 
**YearRate** | **string** | Annual interest rate | [optional] 
**Type** | **int32** | Product type: 1 for regular, 2 for VIP | [optional] 
**PreRedeem** | **int32** | Whether early redemption is supported: 0 for not supported, 1 for supported | [optional] 
**Reinvest** | **int32** | Whether auto-renewal is supported: 0 for not supported, 1 for supported | [optional] 
**RedeemAccount** | **int32** | Whether fixed-to-flexible conversion is supported: 0 for not supported, 1 for supported | [optional] 
**MinVip** | **int32** | Minimum VIP level requirement, 0-16, 0 means no restriction | [optional] 
**MaxVip** | **int32** | Maximum VIP level requirement (0-16), 0 means no restriction | [optional] 
**Status** | **int32** | Product status: 1 for unlisted, 2 for listed, 3 for delisted | [optional] 
**CreateTime** | **string** | Created time | [optional] 
**UserMaxLendVolume** | **string** | User maximum earn amount | [optional] 
**UserTotalAmount** | **string** | Total amount the user has invested in earn products | [optional] 
**SaleStatus** | **int32** | Sale status: 1 for on sale, 2 for sold out | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


