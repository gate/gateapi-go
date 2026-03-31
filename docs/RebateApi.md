# RebateApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgencyTransactionHistory**](RebateApi.md#AgencyTransactionHistory) | **Get** /rebate/agency/transaction_history | Broker obtains transaction history of recommended users
[**AgencyCommissionsHistory**](RebateApi.md#AgencyCommissionsHistory) | **Get** /rebate/agency/commission_history | Broker obtains rebate history of recommended users
[**PartnerTransactionHistory**](RebateApi.md#PartnerTransactionHistory) | **Get** /rebate/partner/transaction_history | Partner obtains transaction history of recommended users
[**PartnerCommissionsHistory**](RebateApi.md#PartnerCommissionsHistory) | **Get** /rebate/partner/commission_history | Partner obtains rebate records of recommended users
[**PartnerSubList**](RebateApi.md#PartnerSubList) | **Get** /rebate/partner/sub_list | Partner subordinate list
[**RebateBrokerCommissionHistory**](RebateApi.md#RebateBrokerCommissionHistory) | **Get** /rebate/broker/commission_history | Broker obtains user&#39;s rebate records
[**RebateBrokerTransactionHistory**](RebateApi.md#RebateBrokerTransactionHistory) | **Get** /rebate/broker/transaction_history | Broker obtains user&#39;s trading history
[**RebateUserInfo**](RebateApi.md#RebateUserInfo) | **Get** /rebate/user/info | User obtains rebate information
[**UserSubRelation**](RebateApi.md#UserSubRelation) | **Get** /rebate/user/sub_relation | User subordinate relationship
[**GetPartnerApplicationRecent**](RebateApi.md#GetPartnerApplicationRecent) | **Get** /rebate/partner/applications/recent | Get recent partner application records
[**GetPartnerEligibility**](RebateApi.md#GetPartnerEligibility) | **Get** /rebate/partner/eligibility | Check partner application eligibility
[**GetPartnerAgentDataAggregated**](RebateApi.md#GetPartnerAgentDataAggregated) | **Get** /rebate/partner/data/aggregated | Aggregated partner agent statistics


## AgencyTransactionHistory

> []AgencyTransactionHistory AgencyTransactionHistory(ctx, optional)

Broker obtains transaction history of recommended users

Record query time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **AgencyTransactionHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgencyTransactionHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Specify the trading pair. If not specified, returns all trading pairs | 
**userId** | **optional.Int64**| User ID. If not specified, all user records will be returned | 
**from** | **optional.Int64**| Start time for querying records. If not specified, defaults to 7 days before current time | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

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
    
    result, _, err := client.RebateApi.AgencyTransactionHistory(ctx, nil)
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

[**[]AgencyTransactionHistory**](AgencyTransactionHistory.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AgencyCommissionsHistory

> []AgencyCommissionHistory AgencyCommissionsHistory(ctx, optional)

Broker obtains rebate history of recommended users

Record query time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **AgencyCommissionsHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgencyCommissionsHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Specify the currency. If not specified, returns all currencies | 
**commissionType** | **optional.Int32**| Rebate type: 1 - Direct rebate, 2 - Indirect rebate, 3 - Self rebate | 
**userId** | **optional.Int64**| User ID. If not specified, all user records will be returned | 
**from** | **optional.Int64**| Start time for querying records. If not specified, defaults to 7 days before current time | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

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
    
    result, _, err := client.RebateApi.AgencyCommissionsHistory(ctx, nil)
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

[**[]AgencyCommissionHistory**](AgencyCommissionHistory.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PartnerTransactionHistory

> PartnerTransactionHistory PartnerTransactionHistory(ctx, optional)

Partner obtains transaction history of recommended users

Record query time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **PartnerTransactionHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PartnerTransactionHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Specify the trading pair. If not specified, returns all trading pairs | 
**userId** | **optional.Int64**| User ID. If not specified, all user records will be returned | 
**from** | **optional.Int64**| Start time for querying records. If not specified, defaults to 7 days before current time | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

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
    
    result, _, err := client.RebateApi.PartnerTransactionHistory(ctx, nil)
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

[**PartnerTransactionHistory**](PartnerTransactionHistory.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PartnerCommissionsHistory

> PartnerCommissionHistory PartnerCommissionsHistory(ctx, optional)

Partner obtains rebate records of recommended users

Record query time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **PartnerCommissionsHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PartnerCommissionsHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Specify the currency. If not specified, returns all currencies | 
**userId** | **optional.Int64**| User ID. If not specified, all user records will be returned | 
**from** | **optional.Int64**| Start time for querying records. If not specified, defaults to 7 days before current time | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

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
    
    result, _, err := client.RebateApi.PartnerCommissionsHistory(ctx, nil)
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

[**PartnerCommissionHistory**](PartnerCommissionHistory.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PartnerSubList

> PartnerSubList PartnerSubList(ctx, optional)

Partner subordinate list

Including sub-agents, direct customers, and indirect customers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **PartnerSubListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PartnerSubListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**userId** | **optional.Int64**| User ID. If not specified, all user records will be returned | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

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
    
    result, _, err := client.RebateApi.PartnerSubList(ctx, nil)
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

[**PartnerSubList**](PartnerSubList.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RebateBrokerCommissionHistory

> []BrokerCommission RebateBrokerCommissionHistory(ctx, optional)

Broker obtains user's rebate records

Record query time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **RebateBrokerCommissionHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RebateBrokerCommissionHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**userId** | **optional.Int64**| User ID. If not specified, all user records will be returned | 
**from** | **optional.Int64**| Start time of the query record. If not specified, defaults to 30 days before the current time | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 

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
    
    result, _, err := client.RebateApi.RebateBrokerCommissionHistory(ctx, nil)
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

[**[]BrokerCommission**](BrokerCommission.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RebateBrokerTransactionHistory

> []BrokerTransactionHistory RebateBrokerTransactionHistory(ctx, optional)

Broker obtains user's trading history

Record query time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **RebateBrokerTransactionHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RebateBrokerTransactionHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**userId** | **optional.Int64**| User ID. If not specified, all user records will be returned | 
**from** | **optional.Int64**| Start time of the query record. If not specified, defaults to 30 days before the current time | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 

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
    
    result, _, err := client.RebateApi.RebateBrokerTransactionHistory(ctx, nil)
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

[**[]BrokerTransactionHistory**](BrokerTransactionHistory.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RebateUserInfo

> []RebateUserInfo RebateUserInfo(ctx, )

User obtains rebate information

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
    
    result, _, err := client.RebateApi.RebateUserInfo(ctx)
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

[**[]RebateUserInfo**](RebateUserInfo.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UserSubRelation

> UserSubRelation UserSubRelation(ctx, userIdList)

User subordinate relationship

Query whether the specified user is within the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdList** | **string**| Query user ID list, separated by commas. If more than 100, only 100 will be returned | 

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
    userIdList := "1, 2, 3" // string - Query user ID list, separated by commas. If more than 100, only 100 will be returned
    
    result, _, err := client.RebateApi.UserSubRelation(ctx, userIdList)
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

[**UserSubRelation**](UserSubRelation.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPartnerApplicationRecent

> PartnerApplicationResponse GetPartnerApplicationRecent(ctx, )

Get recent partner application records

获取当前用户最近的合伙人申请记录。  此接口返回用户最近 30 天内的申请记录，包括申请状态、审核信息、申请材料等详细信息。

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
    
    result, _, err := client.RebateApi.GetPartnerApplicationRecent(ctx)
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

[**PartnerApplicationResponse**](PartnerApplicationResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPartnerEligibility

> EligibilityResponse GetPartnerEligibility(ctx, )

Check partner application eligibility

检查当前用户是否有资格申请成为合伙人。  此接口会检查多个条件： - 账户状态（是否被封禁） - 是否为子账号 - 是否已经是合伙人 - KYC 认证状态 - 是否在其他代理商的邀请链下 - 是否在黑名单中 - 其他业务规则限制

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
    
    result, _, err := client.RebateApi.GetPartnerEligibility(ctx)
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

[**EligibilityResponse**](EligibilityResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPartnerAgentDataAggregated

> PartnerDataAggregatedResponse GetPartnerAgentDataAggregated(ctx, optional)

Aggregated partner agent statistics

查询指定时间范围内合伙人代理的数据聚合统计，包括返佣金额、交易量、净手续费、客户数和交易人数。  **注意事项：** - 交易人数 `trading_user_count` 仅在 `business_type=0`（全部）时返回 - 时间参数使用 UTC+8 时区 - 如不传时间参数，默认查询近 7 天数据 - 仅限合伙人代理访问，子账号无权限

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetPartnerAgentDataAggregatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPartnerAgentDataAggregatedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**startDate** | **optional.String**| 查询开始时间，格式：yyyy-mm-dd hh:ii:ss（UTC+8）  不传时默认为近 7 日开始时间 | 
**endDate** | **optional.String**| 查询结束时间，格式：yyyy-mm-dd hh:ii:ss（UTC+8）  不传时默认为近 7 日结束时间 | 
**businessType** | **optional.Int32**| Business type filter: - 0: All (default) - 1: Spot - 2: Futures - 3: Alpha - 4: Web3 - 5: Perps (DEX) - 6: Exchange All - 7: Web3 All - 8: TradFi | [default to 0]

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
    
    result, _, err := client.RebateApi.GetPartnerAgentDataAggregated(ctx, nil)
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

[**PartnerDataAggregatedResponse**](PartnerDataAggregatedResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
