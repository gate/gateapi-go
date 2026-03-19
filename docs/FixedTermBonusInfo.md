# FixedTermBonusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Activity ID | [optional] 
**ProductId** | **int32** | Associated product ID | [optional] 
**Asset** | **string** | Product currency | [optional] 
**BonusAsset** | **string** | Reward currency | [optional] 
**KycLimit** | **string** | KYC level restrictions, comma-separated | [optional] 
**LadderApr** | [**[]LadderApr**](LadderApr.md) | Tiered annual interest rate | [optional] 
**TotalBonusAmount** | **string** | Total reward amount | [optional] 
**UserTotalBonusAmount** | **string** | Maximum reward per user | [optional] 
**Status** | **int32** | Activity status: 1 for unlisted, 2 for listed, 3 for delisted | [optional] 
**StartTime** | **string** | Activity start time | [optional] 
**EndTime** | **string** | Activity end time | [optional] 
**CreateTime** | **string** | Created time | [optional] 
**StartAt** | **int32** | Activity start timestamp (in seconds) | [optional] 
**EndAt** | **int32** | Activity end timestamp (in seconds) | [optional] 
**TotalIssuedAmount** | **string** | Total rewards distributed | [optional] 
**UserTotalIssuedAmount** | **string** | Total rewards distributed to the user | [optional] 
**BonusAssetPrice** | **string** | Reward currency price (denominated in USDT) | [optional] 
**ProductAssetPrice** | **string** | Product currency price (denominated in USDT) | [optional] 
**ProductYearRate** | **string** | Product base annual interest rate | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


