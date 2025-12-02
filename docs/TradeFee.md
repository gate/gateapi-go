# TradeFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | User ID | [optional] 
**TakerFee** | **string** | spot taker fee rate | [optional] 
**MakerFee** | **string** | spot maker fee rate | [optional] 
**RpiMakerFee** | **string** | spot RPI MM maker fee rate | [optional] 
**GtDiscount** | **bool** | Whether GT deduction discount is enabled | [optional] 
**GtTakerFee** | **string** | Taker fee rate if using GT deduction. It will be 0 if GT deduction is disabled | [optional] 
**GtMakerFee** | **string** | Maker fee rate with GT deduction. Returns 0 if GT deduction is disabled | [optional] 
**LoanFee** | **string** | Loan fee rate of margin lending | [optional] 
**PointType** | **string** | Point card type: 0 - Original version, 1 - New version since 202009 | [optional] 
**FuturesTakerFee** | **string** | Perpetual contract taker fee rate | [optional] 
**FuturesMakerFee** | **string** | Perpetual contract maker fee rate | [optional] 
**FuturesRpiMakerFee** | **string** | contract RPI MM maker fee rate | [optional] 
**DeliveryTakerFee** | **string** | Delivery contract taker fee rate | [optional] 
**DeliveryMakerFee** | **string** | Delivery contract maker fee rate | [optional] 
**DebitFee** | **int32** | Deduction types for rates, 1 - GT deduction, 2 - Point card deduction, 3 - VIP rates | [optional] 
**RpiMm** | **int32** | RPI MM Level | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


