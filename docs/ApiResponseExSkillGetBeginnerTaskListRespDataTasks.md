# ApiResponseExSkillGetBeginnerTaskListRespDataTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WelfareTaskId** | **int64** | Rewards Center task ID | [optional] 
**TaskCenterId** | **int64** | Task center task ID (fixed at 0 for registration tasks) | 
**TaskType** | **int32** | 任务类型：1&#x3D;KYC二级认证 2&#x3D;现货 3&#x3D;合约 4&#x3D;邀请 5&#x3D;量化 6&#x3D;余币宝 7&#x3D;startup 8&#x3D;首次入金 10&#x3D;注册任务 11&#x3D;引导任务 23&#x3D;下载任务 | [optional] 
**TaskName** | **string** | Task name | [optional] 
**TaskDesc** | **string** | Task description, may contain &lt;span&gt; highlight tags | [optional] 
**RewardNum** | **string** | Reward value | [optional] 
**RewardUnit** | **string** | Reward unit (e.g., USDT, BTC) | [optional] 
**PrizeType** | **int32** | Reward type: 1 &#x3D; points, 2 &#x3D; regular coupon, 3 &#x3D; VIP coupon | [optional] 
**Status** | **int32** | 任务状态：0&#x3D;未领取（典型为待领取下载任务） 1&#x3D;已领取/进行中 2&#x3D;已完成待领奖 3&#x3D;发奖中 4&#x3D;已完成/已结算 5&#x3D;已过期 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


