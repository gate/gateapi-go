# ApiResponseExSkillGetBeginnerTaskListRespDataTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WelfareTaskId** | **int64** | Rewards Center task ID | [optional] 
**TaskCenterId** | **int64** | Task center task ID (fixed at 0 for registration tasks) | [optional] 
**TaskType** | **int32** | Task type: 1 &#x3D; KYC level-2 verification, 2 &#x3D; spot, 3 &#x3D; futures, 4 &#x3D; referral, 5 &#x3D; quantitative, 6 &#x3D; earn, 7 &#x3D; startup, 8 &#x3D; first deposit, 10 &#x3D; registration task, 11 &#x3D; onboarding task | [optional] 
**TaskName** | **string** | Task name | [optional] 
**TaskDesc** | **string** | Task description, may contain &lt;span&gt; highlight tags | [optional] 
**RewardNum** | **string** | Reward value | [optional] 
**RewardUnit** | **string** | Reward unit (e.g., USDT, BTC) | [optional] 
**PrizeType** | **int32** | Reward type: 1 &#x3D; points, 2 &#x3D; regular coupon, 3 &#x3D; VIP coupon | [optional] 
**Status** | **int32** | Task status: 0 &#x3D; unclaimed, 1 &#x3D; claimed, 2 &#x3D; reward pending, 3 &#x3D; rewarding, 4 &#x3D; completed, 5 &#x3D; expired | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


