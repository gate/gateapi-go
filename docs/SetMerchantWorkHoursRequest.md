# SetMerchantWorkHoursRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkStatus** | **int32** | Working status. 0: resting, 1: working, 2: using custom working hours | 
**CycleType** | **string** | Custom working cycle; required when work_status is 2 | [optional] 
**DayOfWeek** | **string** | Weekly working days, comma-separated values 1-7 for Monday to Sunday; required when work_status is 2 and cycle_type is Weekly | [optional] 
**TimeZone** | **string** | UTC timezone offset, ranging from -12 to +14; required when work_status is 2 | [optional] 
**StartTime** | **string** | Custom working start time in HH:mm format; required when work_status is 2 and must not be later than end_time | [optional] 
**EndTime** | **string** | Custom working end time in HH:mm format; required when work_status is 2 and must not be earlier than start_time | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


