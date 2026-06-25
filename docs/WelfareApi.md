# WelfareApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserIdentity**](WelfareApi.md#GetUserIdentity) | **Get** /rewards/getUserIdentity | Get user identity
[**GetBeginnerTaskList**](WelfareApi.md#GetBeginnerTaskList) | **Get** /rewards/getBeginnerTaskList | Get beginner task list
[**ClaimTask**](WelfareApi.md#ClaimTask) | **Post** /rewards/claimTask | Get the task
[**ClaimReward**](WelfareApi.md#ClaimReward) | **Post** /rewards/claimReward | Receive mission rewards


## GetUserIdentity

> ApiResponseExSkillGetUserIdentityResp GetUserIdentity(ctx, )

Get user identity

Validate whether the user is eligible for new user rewards. Returns the corresponding error code on validation failure, data is an empty object on success.

### Required Parameters


### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gate/gateapi-go/v7"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    
    result, _, err := client.WelfareApi.GetUserIdentity(ctx)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**ApiResponseExSkillGetUserIdentityResp**](ApiResponse_ExSkillGetUserIdentityResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetBeginnerTaskList

> ApiResponseExSkillGetBeginnerTaskListResp GetBeginnerTaskList(ctx, )

Get beginner task list

Get the current user's onboarding task list. By default, the registration tasks (type=10) and guidance tasks (type=11) that have been assigned to the current user are returned. The registration tasks are ranked first and the guidance tasks are ranked behind. When the user does not have a download task and the system determines that the user has not downloaded the app, a \"Download task to be received\" card will be dynamically added. Among them: - `task_type = 23` for download tasks - `status = 0` for download tasks to be collected Results are cached for 60 seconds. Task lists are limited in number (usually no more than 10) and do not require pagination.

### Required Parameters


### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gate/gateapi-go/v7"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    
    result, _, err := client.WelfareApi.GetBeginnerTaskList(ctx)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**ApiResponseExSkillGetBeginnerTaskListResp**](ApiResponse_ExSkillGetBeginnerTaskListResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ClaimTask

> ApiResponseExSkillClaimTaskResp ClaimTask(ctx, exSkillClaimTaskReq)

Get the task

Receive a single welfare task. The current main scene is for new customer download tasks to be collected, but the interface itself supports new customer registration, guidance, and advanced task types. Processing flow: 1. Read the logged in user 2. Verify user qualifications 3. Risk control verification (event code `task_center`) 4. Verify task configuration and task center tasks 5. Verify whether there is an ongoing task 6. If it is a download task, verify whether the App has been downloaded 7. Write `welfare_user_tasks_xx` 8. Report to task center Risk control transparent transmission fields: - Old fields: `user_id`, `ip`, `const_id`, `is_async`, `action`, `task_id` - New fields: `req_method`, `traceid` - Among them:   - `req_method` from `X-Gate-Request-Source`   - `ip` from `X-Gate-Ip`   - `traceid` from `X-Gate-Trace-Id`   - `const_id` is fixed to an empty string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exSkillClaimTaskReq** | [**ExSkillClaimTaskReq**](ExSkillClaimTaskReq.md)|  | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gate/gateapi-go/v7"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    exSkillClaimTaskReq := gateapi.ExSkillClaimTaskReq{} // ExSkillClaimTaskReq - 
    
    result, _, err := client.WelfareApi.ClaimTask(ctx, exSkillClaimTaskReq)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**ApiResponseExSkillClaimTaskResp**](ApiResponse_ExSkillClaimTaskResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ClaimReward

> ApiResponseExSkillClaimRewardResp ClaimReward(ctx, exSkillClaimRewardReq)

Receive mission rewards

Receive rewards from a single welfare task. Processing flow: 1. Read the logged in user 2. Verify user qualifications 3. Query `welfare_user_tasks_xx` and require the task status to be `StatusDone(2)` 4. Risk control verification (event code `index_page_check`) 5. Query task details and reward information in the task center 6. If the reward is m and n is selected from the prize pool, then `has_m_n_task = true` will be returned and the reward will not be actually issued. 7. Ordinary rewards will enter the original reward collection logic of the Welfare Center Risk control transparent transmission fields: - Old fields: `user_id`, `ip`, `const_id`, `is_async` - New fields: `req_method`, `traceid` - Among them:   - `req_method` from `X-Gate-Request-Source`   - `ip` from `X-Gate-Ip`   - `traceid` from `X-Gate-Trace-Id`   - `const_id` is fixed to an empty string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exSkillClaimRewardReq** | [**ExSkillClaimRewardReq**](ExSkillClaimRewardReq.md)|  | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gate/gateapi-go/v7"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    exSkillClaimRewardReq := gateapi.ExSkillClaimRewardReq{} // ExSkillClaimRewardReq - 
    
    result, _, err := client.WelfareApi.ClaimReward(ctx, exSkillClaimRewardReq)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**ApiResponseExSkillClaimRewardResp**](ApiResponse_ExSkillClaimRewardResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
