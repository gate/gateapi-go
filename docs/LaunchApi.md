# LaunchApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLaunchPoolProjects**](LaunchApi.md#ListLaunchPoolProjects) | **Get** /launch/project-list | Query LaunchPool project list
[**CreateLaunchPoolOrder**](LaunchApi.md#CreateLaunchPoolOrder) | **Post** /launch/create-order | Create LaunchPool staking order
[**RedeemLaunchPool**](LaunchApi.md#RedeemLaunchPool) | **Post** /launch/redeem | Redeem LaunchPool staked assets
[**ListLaunchPoolPledgeRecords**](LaunchApi.md#ListLaunchPoolPledgeRecords) | **Get** /launch/user-pledge-records | Query user pledge records
[**ListLaunchPoolRewardRecords**](LaunchApi.md#ListLaunchPoolRewardRecords) | **Get** /launch/get-user-reward-records | Query user reward records
[**GetHodlerAirdropProjectList**](LaunchApi.md#GetHodlerAirdropProjectList) | **Get** /launch/hodler-airdrop/project-list | Check the list of HODLer Airdrop activities
[**HodlerAirdropOrder**](LaunchApi.md#HodlerAirdropOrder) | **Post** /launch/hodler-airdrop/order | Participate in the HODLer Airdrop event
[**GetHodlerAirdropUserOrderRecords**](LaunchApi.md#GetHodlerAirdropUserOrderRecords) | **Get** /launch/hodler-airdrop/user-order-records | Check HODLer Airdrop participation records
[**GetHodlerAirdropUserAirdropRecords**](LaunchApi.md#GetHodlerAirdropUserAirdropRecords) | **Get** /launch/hodler-airdrop/user-airdrop-records | Query HODLer Airdrop records
[**GetCandyDropActivityListV4**](LaunchApi.md#GetCandyDropActivityListV4) | **Get** /launch/candydrop/activity-list | Query activity list
[**RegisterCandyDropV4**](LaunchApi.md#RegisterCandyDropV4) | **Post** /launch/candydrop/register | Sign up for events
[**GetCandyDropActivityRulesV4**](LaunchApi.md#GetCandyDropActivityRulesV4) | **Get** /launch/candydrop/activity-rules | Query activity rules
[**GetCandyDropTaskProgressV4**](LaunchApi.md#GetCandyDropTaskProgressV4) | **Get** /launch/candydrop/task-progress | Query task completion progress
[**GetCandyDropParticipationRecordsV4**](LaunchApi.md#GetCandyDropParticipationRecordsV4) | **Get** /launch/candydrop/participation-records | Query participation records
[**GetCandyDropAirdropRecordsV4**](LaunchApi.md#GetCandyDropAirdropRecordsV4) | **Get** /launch/candydrop/airdrop-records | Query airdrop records


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

Check the list of HODLer Airdrop activities

Get the HODLer Airdrop activity list, which supports filtering by status, currency/project name, and participation status. This interface does not require user login, and logged in users can obtain personal participation information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetHodlerAirdropProjectListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHodlerAirdropProjectListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**status** | **optional.String**| Activity status filtering, optional values: ACTIVE (in progress + preheating), UNDERWAY (in progress), PREHEAT (preheating), FINISH (ended), return all if not passed | 
**keyword** | **optional.String**| Currency/project name keywords, fuzzy matching | 
**join** | **optional.Int32**| Participation filter: 0 all (default), 1 only participated | [default to 0]
**page** | **optional.Int32**| Page number, default 1 | [default to 1]
**size** | **optional.Int32**| Number of items per page, default 10 | [default to 10]

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

Participate in the HODLer Airdrop event

To participate in designated HODLer Airdrop activities, you need to hold GT. This interface requires user login authentication and must meet KYC requirements. It does not support sub-accounts and enterprise/institutional users.

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

Check HODLer Airdrop participation records

Query the user's HODLer Airdrop participation record and return the effective holdings and airdrop amount of each activity. This interface requires user login authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetHodlerAirdropUserOrderRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHodlerAirdropUserOrderRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**keyword** | **optional.String**| Currency name keyword filtering | 
**startTimest** | **optional.Int32**| Start timestamp (seconds) | 
**endTimest** | **optional.Int32**| end timestamp (seconds) | 
**page** | **optional.Int32**| Page number, default 1 | [default to 1]
**size** | **optional.Int32**| Number of items per page, default 10 | [default to 10]

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

Query HODLer Airdrop records

Query the HODLer Airdrop airdrop distribution record that the user has obtained, including basic airdrops, additional airdrops and automatic redemption status. This interface requires user login authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetHodlerAirdropUserAirdropRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHodlerAirdropUserAirdropRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**keyword** | **optional.String**| Currency name keyword filtering | 
**startTimest** | **optional.Int32**| Start timestamp (seconds) | 
**endTimest** | **optional.Int32**| end timestamp (seconds) | 
**page** | **optional.Int32**| Page number, default 1 | [default to 1]
**size** | **optional.Int32**| Number of items per page, default 10 | [default to 10]

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

Query activity list

Supports multi-dimensional filtering of CandyDrop activities, and each query returns the top ten data sorted by the list. No login required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCandyDropActivityListV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCandyDropActivityListV4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**status** | **optional.String**| Activity status filtering: ongoing (in progress), upcoming (about to start), ended (ended), if not passed, all will be returned | 
**ruleName** | **optional.String**| Task type filtering: spot (spot), futures (contract), deposit (recharge), invite (invitation), trading_bot (trading robot), simple_earn (Yu Bibao), first_deposit (first deposit), alpha (Alpha), flash_swap (flash swap), tradfi (TradFi), etf (ETF) | 
**registerStatus** | **optional.String**| Participation status screening: registered (already participated), unregistered (not participated), if not passed, all will be returned | 
**currency** | **optional.String**| Currency name filter | 
**limit** | **optional.Int32**| Number of items returned, default 10, maximum 30 | [default to 10]
**offset** | **optional.Int32**| Offset, default 0 | [default to 0]

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

Sign up for events

Sign up for select CandyDrop events. Login is required and API Key signature authentication is required.

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

Query activity rules

Query the rules of a specific activity, including prize pool and corresponding task data. No login required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCandyDropActivityRulesV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCandyDropActivityRulesV4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**activityId** | **optional.Int64**| Activity ID, choose one from currency, at least one of them must be passed | 
**currency** | **optional.String**| Project/currency name, choose one from activity_id, at least one of them must be passed | 

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

Query task completion progress

Check the completion progress of tasks that are in progress and have been registered/participated. Login required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCandyDropTaskProgressV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCandyDropTaskProgressV4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**activityId** | **optional.Int64**| Activity ID, choose one from currency, at least one of them must be passed | 
**currency** | **optional.String**| Project/currency name, choose one from activity_id, at least one of them must be passed | 

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

Query participation records

Query the user's CandyDrop participation details. Login required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCandyDropParticipationRecordsV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCandyDropParticipationRecordsV4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Currency name filter | 
**status** | **optional.String**| Status filtering: ongoing (in progress), awaiting_draw (to be drawn), won (already won), not_win (not won) | 
**startTime** | **optional.Int64**| Start time (Unix timestamp seconds) | 
**endTime** | **optional.Int64**| End time (Unix timestamp seconds) | 
**page** | **optional.Int32**| Page number, default 1 | [default to 1]
**limit** | **optional.Int32**| Number of items per page, default 10, maximum 30 | [default to 10]

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

Query airdrop records

Query the user's CandyDrop airdrop details. Login required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCandyDropAirdropRecordsV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCandyDropAirdropRecordsV4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Currency name filter | 
**startTime** | **optional.Int64**| Start time (Unix timestamp seconds) | 
**endTime** | **optional.Int64**| End time (Unix timestamp seconds) | 
**page** | **optional.Int32**| Page number, default 1 | [default to 1]
**limit** | **optional.Int32**| Number of items per page, default 10, maximum 30 | [default to 10]

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
