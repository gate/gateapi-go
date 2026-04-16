# LaunchApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLaunchPoolProjects**](LaunchApi.md#ListLaunchPoolProjects) | **Get** /launch/project-list | Query LaunchPool project list
[**CreateLaunchPoolOrder**](LaunchApi.md#CreateLaunchPoolOrder) | **Post** /launch/create-order | Create LaunchPool staking order
[**RedeemLaunchPool**](LaunchApi.md#RedeemLaunchPool) | **Post** /launch/redeem | Redeem LaunchPool staked assets
[**ListLaunchPoolPledgeRecords**](LaunchApi.md#ListLaunchPoolPledgeRecords) | **Get** /launch/user-pledge-records | Query user pledge records
[**ListLaunchPoolRewardRecords**](LaunchApi.md#ListLaunchPoolRewardRecords) | **Get** /launch/get-user-reward-records | Query user reward records
[**GetHodlerAirdropProjectList**](LaunchApi.md#GetHodlerAirdropProjectList) | **Get** /launch/hodler-airdrop/project-list | 查询HODLer Airdrop活动列表
[**HodlerAirdropOrder**](LaunchApi.md#HodlerAirdropOrder) | **Post** /launch/hodler-airdrop/order | 参与HODLer Airdrop活动
[**GetHodlerAirdropUserOrderRecords**](LaunchApi.md#GetHodlerAirdropUserOrderRecords) | **Get** /launch/hodler-airdrop/user-order-records | 查询HODLer Airdrop参与记录
[**GetHodlerAirdropUserAirdropRecords**](LaunchApi.md#GetHodlerAirdropUserAirdropRecords) | **Get** /launch/hodler-airdrop/user-airdrop-records | 查询HODLer Airdrop空投记录
[**GetCandyDropActivityListV4**](LaunchApi.md#GetCandyDropActivityListV4) | **Get** /launch/candydrop/activity-list | 查询活动列表
[**RegisterCandyDropV4**](LaunchApi.md#RegisterCandyDropV4) | **Post** /launch/candydrop/register | 报名参与活动
[**GetCandyDropActivityRulesV4**](LaunchApi.md#GetCandyDropActivityRulesV4) | **Get** /launch/candydrop/activity-rules | 查询活动规则
[**GetCandyDropTaskProgressV4**](LaunchApi.md#GetCandyDropTaskProgressV4) | **Get** /launch/candydrop/task-progress | 查询任务完成进度
[**GetCandyDropParticipationRecordsV4**](LaunchApi.md#GetCandyDropParticipationRecordsV4) | **Get** /launch/candydrop/participation-records | 查询参与记录
[**GetCandyDropAirdropRecordsV4**](LaunchApi.md#GetCandyDropAirdropRecordsV4) | **Get** /launch/candydrop/airdrop-records | 查询空投记录


## ListLaunchPoolProjects

> []LaunchPoolV4Project ListLaunchPoolProjects(ctx, optional)

Query LaunchPool project list

Retrieve the list of available LaunchPool projects, including basic project information and reward pool configuration. This endpoint does not require user authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListLaunchPoolProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLaunchPoolProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**status** | **optional.Int32**| Filter by project status: 0 for all, 1 for ongoing, 2 for warming up, 3 for ended, 4 for ongoing and warming up | 
**mortgageCoin** | **optional.String**| Exact match by staking currency | 
**searchCoin** | **optional.String**| Fuzzy match by reward currency and name | 
**limitRule** | **optional.Int32**| Limit rule: 0 for regular pool, 1 for beginner pool | 
**sortType** | **optional.Int32**| Sort type: 1 for max APR descending, 2 for max APR ascending | 
**page** | **optional.Int32**| Page number, default 1 | [default to 1]
**pageSize** | **optional.Int32**| Number of items per page, default 10, maximum 30 | [default to 10]

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
    ctx := context.Background()
    
    result, _, err := client.LaunchApi.ListLaunchPoolProjects(ctx, nil)
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

[**[]LaunchPoolV4Project**](LaunchPoolV4Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateLaunchPoolOrder

> LaunchPoolV4CreateOrderResponse CreateLaunchPoolOrder(ctx, createOrderV4)

Create LaunchPool staking order

Create a new staking order for asset staking mining. This endpoint requires API Key signature authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createOrderV4** | [**CreateOrderV4**](CreateOrderV4.md)|  | 

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
    createOrderV4 := gateapi.CreateOrderV4{} // CreateOrderV4 - 
    
    result, _, err := client.LaunchApi.CreateLaunchPoolOrder(ctx, createOrderV4)
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

[**LaunchPoolV4CreateOrderResponse**](LaunchPoolV4CreateOrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RedeemLaunchPool

> RedeemLaunchPoolResponse RedeemLaunchPool(ctx, redeemV4)

Redeem LaunchPool staked assets

Redeem staked assets and end staking mining. This endpoint requires API Key signature authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**redeemV4** | [**RedeemV4**](RedeemV4.md)|  | 

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
    redeemV4 := gateapi.RedeemV4{} // RedeemV4 - 
    
    result, _, err := client.LaunchApi.RedeemLaunchPool(ctx, redeemV4)
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

[**RedeemLaunchPoolResponse**](RedeemLaunchPoolResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLaunchPoolPledgeRecords

> []LaunchPoolV4PledgeRecord ListLaunchPoolPledgeRecords(ctx, optional)

Query user pledge records

Query user's staking and redemption operation records. This endpoint requires user authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListLaunchPoolPledgeRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLaunchPoolPledgeRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int64**| Page number, default 1 | [default to 1]
**pageSize** | **optional.Int64**| Number of items per page, maximum 30 | [default to 10]
**type_** | **optional.Int32**| Type: 1 for pledge, 2 for redemption | 
**startTime** | **optional.String**| Start time, format: YYYY-MM-DD HH:MM:SS | 
**endTime** | **optional.String**| End time, format: YYYY-MM-DD HH:MM:SS | 
**coin** | **optional.String**| Collateral currency | 

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
    
    result, _, err := client.LaunchApi.ListLaunchPoolPledgeRecords(ctx, nil)
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

[**[]LaunchPoolV4PledgeRecord**](LaunchPoolV4PledgeRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLaunchPoolRewardRecords

> []LaunchPoolV4RewardRecord ListLaunchPoolRewardRecords(ctx, optional)

Query user reward records

Query the user's staking reward records. This endpoint requires user authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListLaunchPoolRewardRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLaunchPoolRewardRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int64**| Page number, default 1 | [default to 1]
**pageSize** | **optional.Int64**| Number of items per page, maximum 30 | [default to 10]
**startTime** | **optional.Int64**| Start timestamp | 
**endTime** | **optional.Int64**| End Timestamp | 
**coin** | **optional.String**| Reward currency | 

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
    
    result, _, err := client.LaunchApi.ListLaunchPoolRewardRecords(ctx, nil)
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

[**[]LaunchPoolV4RewardRecord**](LaunchPoolV4RewardRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetHodlerAirdropProjectList

> []HodlerAirdropV4ProjectItem GetHodlerAirdropProjectList(ctx, optional)

查询HODLer Airdrop活动列表

获取HODLer Airdrop活动列表，支持按状态、币种/项目名称、参与情况筛选。此接口无需用户登录，登录用户可获取个人参与信息。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetHodlerAirdropProjectListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHodlerAirdropProjectListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**status** | **optional.String**| 活动状态筛选，可选值：ACTIVE（进行中+预热中）、UNDERWAY（进行中）、PREHEAT（预热中）、FINISH（已结束），不传返回全部 | 
**keyword** | **optional.String**| 币种/项目名称关键词，模糊匹配 | 
**join** | **optional.Int32**| 参与情况筛选：0全部（默认），1仅已参与 | [default to 0]
**page** | **optional.Int32**| 页码，默认1 | [default to 1]
**size** | **optional.Int32**| 每页条数，默认10 | [default to 10]

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
    ctx := context.Background()
    
    result, _, err := client.LaunchApi.GetHodlerAirdropProjectList(ctx, nil)
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

[**[]HodlerAirdropV4ProjectItem**](HodlerAirdropV4ProjectItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## HodlerAirdropOrder

> HodlerAirdropV4OrderResponse HodlerAirdropOrder(ctx, hodlerAirdropV4OrderRequest)

参与HODLer Airdrop活动

参与指定的HODLer Airdrop活动，需持有GT。此接口需要用户登录认证，且须满足KYC要求，不支持子账户、企业/机构用户。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hodlerAirdropV4OrderRequest** | [**HodlerAirdropV4OrderRequest**](HodlerAirdropV4OrderRequest.md)|  | 

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
    hodlerAirdropV4OrderRequest := gateapi.HodlerAirdropV4OrderRequest{} // HodlerAirdropV4OrderRequest - 
    
    result, _, err := client.LaunchApi.HodlerAirdropOrder(ctx, hodlerAirdropV4OrderRequest)
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

[**HodlerAirdropV4OrderResponse**](HodlerAirdropV4OrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetHodlerAirdropUserOrderRecords

> []HodlerAirdropV4UserOrderRecord GetHodlerAirdropUserOrderRecords(ctx, optional)

查询HODLer Airdrop参与记录

查询用户的HODLer Airdrop参与记录，返回每个活动的有效持仓和空投金额。此接口需要用户登录认证。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetHodlerAirdropUserOrderRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHodlerAirdropUserOrderRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**keyword** | **optional.String**| 币种名称关键词筛选 | 
**startTimest** | **optional.Int32**| 开始时间戳（秒） | 
**endTimest** | **optional.Int32**| 结束时间戳（秒） | 
**page** | **optional.Int32**| 页码，默认1 | [default to 1]
**size** | **optional.Int32**| 每页条数，默认10 | [default to 10]

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
    
    result, _, err := client.LaunchApi.GetHodlerAirdropUserOrderRecords(ctx, nil)
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

[**[]HodlerAirdropV4UserOrderRecord**](HodlerAirdropV4UserOrderRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetHodlerAirdropUserAirdropRecords

> []HodlerAirdropV4UserAirdropRecord GetHodlerAirdropUserAirdropRecords(ctx, optional)

查询HODLer Airdrop空投记录

查询用户已获得的HODLer Airdrop空投发放记录，包含基础空投、额外空投和自动兑换状态。此接口需要用户登录认证。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetHodlerAirdropUserAirdropRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHodlerAirdropUserAirdropRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**keyword** | **optional.String**| 币种名称关键词筛选 | 
**startTimest** | **optional.Int32**| 开始时间戳（秒） | 
**endTimest** | **optional.Int32**| 结束时间戳（秒） | 
**page** | **optional.Int32**| 页码，默认1 | [default to 1]
**size** | **optional.Int32**| 每页条数，默认10 | [default to 10]

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
    
    result, _, err := client.LaunchApi.GetHodlerAirdropUserAirdropRecords(ctx, nil)
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

[**[]HodlerAirdropV4UserAirdropRecord**](HodlerAirdropV4UserAirdropRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCandyDropActivityListV4

> []CandyDropV4ActivityCd01 GetCandyDropActivityListV4(ctx, optional)

查询活动列表

支持多维度筛选 CandyDrop 活动，每次查询返回列表排序的前十条数据。不需要登录。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCandyDropActivityListV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCandyDropActivityListV4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**status** | **optional.String**| 活动状态筛选：ongoing(进行中)、upcoming(即将开始)、ended(已结束)，不传则返回全部 | 
**ruleName** | **optional.String**| 任务类型筛选：spot(现货)、futures(合约)、deposit(充值)、invite(邀请)、trading_bot(交易机器人)、simple_earn(余币宝)、first_deposit(首笔入金)、alpha(Alpha)、flash_swap(闪兑)、tradfi(TradFi)、etf(ETF) | 
**registerStatus** | **optional.String**| 参与情况筛选：registered(已参与)、unregistered(未参与)，不传则返回全部 | 
**currency** | **optional.String**| 币种名称筛选 | 
**limit** | **optional.Int32**| 返回条数，默认10，最大30 | [default to 10]
**offset** | **optional.Int32**| 偏移量，默认0 | [default to 0]

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
    ctx := context.Background()
    
    result, _, err := client.LaunchApi.GetCandyDropActivityListV4(ctx, nil)
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

[**[]CandyDropV4ActivityCd01**](CandyDropV4Activity_cd01.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RegisterCandyDropV4

> CandyDropV4RegisterRespCd02 RegisterCandyDropV4(ctx, candyDropV4RegisterReqCd02)

报名参与活动

报名参与特定 CandyDrop 活动。需要登录，需要 API Key 签名认证。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**candyDropV4RegisterReqCd02** | [**CandyDropV4RegisterReqCd02**](CandyDropV4RegisterReqCd02.md)|  | 

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
    candyDropV4RegisterReqCd02 := gateapi.CandyDropV4RegisterReqCd02{} // CandyDropV4RegisterReqCd02 - 
    
    result, _, err := client.LaunchApi.RegisterCandyDropV4(ctx, candyDropV4RegisterReqCd02)
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

[**CandyDropV4RegisterRespCd02**](CandyDropV4RegisterResp_cd02.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCandyDropActivityRulesV4

> CandyDropV4ActivityRulesCd03 GetCandyDropActivityRulesV4(ctx, optional)

查询活动规则

查询特定活动的规则，包括奖池及对应任务数据。不需要登录。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCandyDropActivityRulesV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCandyDropActivityRulesV4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**activityId** | **optional.Int64**| 活动ID，与 currency 二选一，至少须传其一 | 
**currency** | **optional.String**| 项目/币种名称，与 activity_id 二选一，至少须传其一 | 

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
    ctx := context.Background()
    
    result, _, err := client.LaunchApi.GetCandyDropActivityRulesV4(ctx, nil)
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

[**CandyDropV4ActivityRulesCd03**](CandyDropV4ActivityRules_cd03.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCandyDropTaskProgressV4

> CandyDropV4TaskProgressCd04 GetCandyDropTaskProgressV4(ctx, optional)

查询任务完成进度

查询进行中且已报名/参与的任务完成进度。需要登录。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCandyDropTaskProgressV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCandyDropTaskProgressV4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**activityId** | **optional.Int64**| 活动ID，与 currency 二选一，至少须传其一 | 
**currency** | **optional.String**| 项目/币种名称，与 activity_id 二选一，至少须传其一 | 

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
    
    result, _, err := client.LaunchApi.GetCandyDropTaskProgressV4(ctx, nil)
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

[**CandyDropV4TaskProgressCd04**](CandyDropV4TaskProgress_cd04.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCandyDropParticipationRecordsV4

> []CandyDropV4ParticipationRecordCd05 GetCandyDropParticipationRecordsV4(ctx, optional)

查询参与记录

查询用户的 CandyDrop 参与详情。需要登录。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCandyDropParticipationRecordsV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCandyDropParticipationRecordsV4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| 币种名称筛选 | 
**status** | **optional.String**| 状态筛选：ongoing(进行中)、awaiting_draw(待开奖)、won(已中奖)、not_win(未中奖) | 
**startTime** | **optional.Int64**| 开始时间（Unix 时间戳秒） | 
**endTime** | **optional.Int64**| 结束时间（Unix 时间戳秒） | 
**page** | **optional.Int32**| 页码，默认1 | [default to 1]
**limit** | **optional.Int32**| 每页条数，默认10，最大30 | [default to 10]

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
    
    result, _, err := client.LaunchApi.GetCandyDropParticipationRecordsV4(ctx, nil)
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

[**[]CandyDropV4ParticipationRecordCd05**](CandyDropV4ParticipationRecord_cd05.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCandyDropAirdropRecordsV4

> []CandyDropV4AirdropRecordCd06 GetCandyDropAirdropRecordsV4(ctx, optional)

查询空投记录

查询用户的 CandyDrop 空投详情。需要登录。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCandyDropAirdropRecordsV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCandyDropAirdropRecordsV4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| 币种名称筛选 | 
**startTime** | **optional.Int64**| 开始时间（Unix 时间戳秒） | 
**endTime** | **optional.Int64**| 结束时间（Unix 时间戳秒） | 
**page** | **optional.Int32**| 页码，默认1 | [default to 1]
**limit** | **optional.Int32**| 每页条数，默认10，最大30 | [default to 10]

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
    
    result, _, err := client.LaunchApi.GetCandyDropAirdropRecordsV4(ctx, nil)
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

[**[]CandyDropV4AirdropRecordCd06**](CandyDropV4AirdropRecord_cd06.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
