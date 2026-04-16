# BotApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAIHubStrategyRecommend**](BotApi.md#GetAIHubStrategyRecommend) | **Get** /bot/strategy/recommend | 获取 AIHub 策略推荐
[**PostAIHubSpotGridCreate**](BotApi.md#PostAIHubSpotGridCreate) | **Post** /bot/spot-grid/create | 创建现货网格
[**PostAIHubMarginGridCreate**](BotApi.md#PostAIHubMarginGridCreate) | **Post** /bot/margin-grid/create | 创建杠杆网格
[**PostAIHubInfiniteGridCreate**](BotApi.md#PostAIHubInfiniteGridCreate) | **Post** /bot/infinite-grid/create | 创建无限网格
[**PostAIHubFuturesGridCreate**](BotApi.md#PostAIHubFuturesGridCreate) | **Post** /bot/futures-grid/create | 创建合约网格
[**PostAIHubSpotMartingaleCreate**](BotApi.md#PostAIHubSpotMartingaleCreate) | **Post** /bot/spot-martingale/create | 创建现货马丁
[**PostAIHubContractMartingaleCreate**](BotApi.md#PostAIHubContractMartingaleCreate) | **Post** /bot/contract-martingale/create | 创建合约马丁
[**GetAIHubPortfolioRunning**](BotApi.md#GetAIHubPortfolioRunning) | **Get** /bot/portfolio/running | 查询运行中策略列表
[**GetAIHubPortfolioDetail**](BotApi.md#GetAIHubPortfolioDetail) | **Get** /bot/portfolio/detail | 查询单策略详情
[**PostAIHubPortfolioStop**](BotApi.md#PostAIHubPortfolioStop) | **Post** /bot/portfolio/stop | 终止单个运行中策略


## GetAIHubStrategyRecommend

> AiHubDiscoverSuccessResponse GetAIHubStrategyRecommend(ctx, optional)

获取 AIHub 策略推荐

discover 域唯一正式接口。  支持场景： - `top1` - `bundle` - `filter` - `refresh`  约束： - 主动推荐池仅包含 `spot_grid`、`futures_grid`、`spot_martingale` - 可返回但不主动推荐 `infinite_grid`、`margin_grid` - 不得返回 `contract_martingale`、`smart-position`、`spot-future-arbitrage` - `scene=filter` 时只允许按 `market`、`backtest_apr_gte`、`max_drawdown_lte` 过滤 - `scene=refresh` 通过 `refresh_recommendation_id` 承接刷新上下文；正式最小格式只要求 `strategy_type|market` - 若上游直接透传上一条推荐的 `recommendation_id`，其中第三段 `backtest_id` 当前会被忽略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetAIHubStrategyRecommendOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAIHubStrategyRecommendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**market** | **optional.String**| 交易对，例如 &#x60;BTC_USDT&#x60; | 
**strategyType** | **optional.String**| 推荐目标策略类型；&#x60;contract_martingale&#x60; 不允许 | 
**direction** | **optional.String**| 行情方向 | 
**investAmount** | **optional.String**| 投入金额，字符串透传 | 
**scene** | **optional.String**| 推荐场景；为空时 bot-service 可按实现逻辑自动推断 | 
**refreshRecommendationId** | **optional.String**| 推荐刷新上下文。&#x60;scene&#x3D;refresh&#x60; 时使用；当 &#x60;scene&#x60; 为空但该字段存在时，bot-service 也会自动判定为 &#x60;refresh&#x60;。 正式最小格式为 &#x60;strategy_type|market&#x60;；若直接透传上一条推荐的 &#x60;recommendation_id&#x60;，第三段 &#x60;backtest_id&#x60; 会被忽略。 | 
**limit** | **optional.Int32**| 返回数量；&#x60;scene&#x3D;filter&#x60; 时实际结果最多 10 条 | 
**maxDrawdownLte** | **optional.String**| 最大回撤上限 | 
**backtestAprGte** | **optional.String**| 回测年化下限 | 
**xGateServiceId** | **optional.String**| 调用来源标识；如有需要由 APIv4 注入 | 
**xGateAppLang** | **optional.String**| 语言上下文，例如 &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| 请求链路 ID；调用方可透传 | 
**xTraceId** | **optional.String**| trace header；可由 APIv4 统一生成 | 

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
    
    result, _, err := client.BotApi.GetAIHubStrategyRecommend(ctx, nil)
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

[**AiHubDiscoverSuccessResponse**](AIHubDiscoverSuccessResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostAIHubSpotGridCreate

> AiHubCreateSuccessResponse PostAIHubSpotGridCreate(ctx, spotGridCreateRequest, optional)

创建现货网格

根据传入参数创建现货网格策略。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spotGridCreateRequest** | [**SpotGridCreateRequest**](SpotGridCreateRequest.md)|  | 
**optional** | **PostAIHubSpotGridCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAIHubSpotGridCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateServiceId** | **optional.String**| 调用来源标识；如有需要由 APIv4 注入 | 
**xGateAppLang** | **optional.String**| 语言上下文，例如 &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| 请求链路 ID；调用方可透传 | 
**xTraceId** | **optional.String**| trace header；可由 APIv4 统一生成 | 

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
    spotGridCreateRequest := gateapi.SpotGridCreateRequest{} // SpotGridCreateRequest - 
    
    result, _, err := client.BotApi.PostAIHubSpotGridCreate(ctx, spotGridCreateRequest, nil)
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

[**AiHubCreateSuccessResponse**](AIHubCreateSuccessResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostAIHubMarginGridCreate

> AiHubCreateSuccessResponse PostAIHubMarginGridCreate(ctx, marginGridCreateRequest, optional)

创建杠杆网格

根据传入参数创建杠杆网格策略。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marginGridCreateRequest** | [**MarginGridCreateRequest**](MarginGridCreateRequest.md)|  | 
**optional** | **PostAIHubMarginGridCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAIHubMarginGridCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateServiceId** | **optional.String**| 调用来源标识；如有需要由 APIv4 注入 | 
**xGateAppLang** | **optional.String**| 语言上下文，例如 &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| 请求链路 ID；调用方可透传 | 
**xTraceId** | **optional.String**| trace header；可由 APIv4 统一生成 | 

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
    marginGridCreateRequest := gateapi.MarginGridCreateRequest{} // MarginGridCreateRequest - 
    
    result, _, err := client.BotApi.PostAIHubMarginGridCreate(ctx, marginGridCreateRequest, nil)
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

[**AiHubCreateSuccessResponse**](AIHubCreateSuccessResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostAIHubInfiniteGridCreate

> AiHubCreateSuccessResponse PostAIHubInfiniteGridCreate(ctx, infiniteGridCreateRequest, optional)

创建无限网格

根据传入参数创建无限网格策略。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infiniteGridCreateRequest** | [**InfiniteGridCreateRequest**](InfiniteGridCreateRequest.md)|  | 
**optional** | **PostAIHubInfiniteGridCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAIHubInfiniteGridCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateServiceId** | **optional.String**| 调用来源标识；如有需要由 APIv4 注入 | 
**xGateAppLang** | **optional.String**| 语言上下文，例如 &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| 请求链路 ID；调用方可透传 | 
**xTraceId** | **optional.String**| trace header；可由 APIv4 统一生成 | 

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
    infiniteGridCreateRequest := gateapi.InfiniteGridCreateRequest{} // InfiniteGridCreateRequest - 
    
    result, _, err := client.BotApi.PostAIHubInfiniteGridCreate(ctx, infiniteGridCreateRequest, nil)
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

[**AiHubCreateSuccessResponse**](AIHubCreateSuccessResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostAIHubFuturesGridCreate

> AiHubCreateSuccessResponse PostAIHubFuturesGridCreate(ctx, futuresGridCreateRequest, optional)

创建合约网格

根据传入参数创建合约网格策略。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**futuresGridCreateRequest** | [**FuturesGridCreateRequest**](FuturesGridCreateRequest.md)|  | 
**optional** | **PostAIHubFuturesGridCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAIHubFuturesGridCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateServiceId** | **optional.String**| 调用来源标识；如有需要由 APIv4 注入 | 
**xGateAppLang** | **optional.String**| 语言上下文，例如 &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| 请求链路 ID；调用方可透传 | 
**xTraceId** | **optional.String**| trace header；可由 APIv4 统一生成 | 

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
    futuresGridCreateRequest := gateapi.FuturesGridCreateRequest{} // FuturesGridCreateRequest - 
    
    result, _, err := client.BotApi.PostAIHubFuturesGridCreate(ctx, futuresGridCreateRequest, nil)
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

[**AiHubCreateSuccessResponse**](AIHubCreateSuccessResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostAIHubSpotMartingaleCreate

> AiHubCreateSuccessResponse PostAIHubSpotMartingaleCreate(ctx, spotMartingaleCreateRequest, optional)

创建现货马丁

根据传入参数创建现货马丁策略。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spotMartingaleCreateRequest** | [**SpotMartingaleCreateRequest**](SpotMartingaleCreateRequest.md)|  | 
**optional** | **PostAIHubSpotMartingaleCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAIHubSpotMartingaleCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateServiceId** | **optional.String**| 调用来源标识；如有需要由 APIv4 注入 | 
**xGateAppLang** | **optional.String**| 语言上下文，例如 &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| 请求链路 ID；调用方可透传 | 
**xTraceId** | **optional.String**| trace header；可由 APIv4 统一生成 | 

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
    spotMartingaleCreateRequest := gateapi.SpotMartingaleCreateRequest{} // SpotMartingaleCreateRequest - 
    
    result, _, err := client.BotApi.PostAIHubSpotMartingaleCreate(ctx, spotMartingaleCreateRequest, nil)
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

[**AiHubCreateSuccessResponse**](AIHubCreateSuccessResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostAIHubContractMartingaleCreate

> AiHubCreateSuccessResponse PostAIHubContractMartingaleCreate(ctx, contractMartingaleCreateRequest, optional)

创建合约马丁

根据传入参数创建合约马丁策略。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractMartingaleCreateRequest** | [**ContractMartingaleCreateRequest**](ContractMartingaleCreateRequest.md)|  | 
**optional** | **PostAIHubContractMartingaleCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAIHubContractMartingaleCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateServiceId** | **optional.String**| 调用来源标识；如有需要由 APIv4 注入 | 
**xGateAppLang** | **optional.String**| 语言上下文，例如 &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| 请求链路 ID；调用方可透传 | 
**xTraceId** | **optional.String**| trace header；可由 APIv4 统一生成 | 

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
    contractMartingaleCreateRequest := gateapi.ContractMartingaleCreateRequest{} // ContractMartingaleCreateRequest - 
    
    result, _, err := client.BotApi.PostAIHubContractMartingaleCreate(ctx, contractMartingaleCreateRequest, nil)
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

[**AiHubCreateSuccessResponse**](AIHubCreateSuccessResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetAIHubPortfolioRunning

> AiHubPortfolioRunningSuccessResponse GetAIHubPortfolioRunning(ctx, optional)

查询运行中策略列表

查询当前用户运行中的 AIHub 策略列表，支持按策略类型、交易对和分页条件过滤。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetAIHubPortfolioRunningOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAIHubPortfolioRunningOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**strategyType** | **optional.String**| 按策略类型过滤 | 
**market** | **optional.String**| 按交易对过滤 | 
**page** | **optional.Int32**| 页码，默认 1 | [default to 1]
**pageSize** | **optional.Int32**| 分页大小，默认 20，最大 50 | [default to 20]
**xGateServiceId** | **optional.String**| 调用来源标识；如有需要由 APIv4 注入 | 
**xGateAppLang** | **optional.String**| 语言上下文，例如 &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| 请求链路 ID；调用方可透传 | 
**xTraceId** | **optional.String**| trace header；可由 APIv4 统一生成 | 

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
    
    result, _, err := client.BotApi.GetAIHubPortfolioRunning(ctx, nil)
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

[**AiHubPortfolioRunningSuccessResponse**](AIHubPortfolioRunningSuccessResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetAIHubPortfolioDetail

> AiHubPortfolioDetailSuccessResponse GetAIHubPortfolioDetail(ctx, strategyId, strategyType, optional)

查询单策略详情

请求中必须同时传 `strategy_id` 与 `strategy_type`，其中 `strategy_type` 用于按策略类型分发到底层详情实现。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyId** | **string**| 策略 ID | 
**strategyType** | **string**| 策略类型；用于底层详情分发 | 
**optional** | **GetAIHubPortfolioDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAIHubPortfolioDetailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateServiceId** | **optional.String**| 调用来源标识；如有需要由 APIv4 注入 | 
**xGateAppLang** | **optional.String**| 语言上下文，例如 &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| 请求链路 ID；调用方可透传 | 
**xTraceId** | **optional.String**| trace header；可由 APIv4 统一生成 | 

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
    strategyId := "strategyId_example" // string - 策略 ID
    strategyType := "strategyType_example" // string - 策略类型；用于底层详情分发
    
    result, _, err := client.BotApi.GetAIHubPortfolioDetail(ctx, strategyId, strategyType, nil)
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

[**AiHubPortfolioDetailSuccessResponse**](AIHubPortfolioDetailSuccessResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostAIHubPortfolioStop

> AiHubPortfolioStopSuccessResponse PostAIHubPortfolioStop(ctx, aiHubPortfolioStopRequest, optional)

终止单个运行中策略

单次请求只允许终止一个策略。 风险提示与二次确认由 OpenClaw 上层承担；本接口只负责执行 stop。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aiHubPortfolioStopRequest** | [**AiHubPortfolioStopRequest**](AiHubPortfolioStopRequest.md)|  | 
**optional** | **PostAIHubPortfolioStopOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAIHubPortfolioStopOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateServiceId** | **optional.String**| 调用来源标识；如有需要由 APIv4 注入 | 
**xGateAppLang** | **optional.String**| 语言上下文，例如 &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| 请求链路 ID；调用方可透传 | 
**xTraceId** | **optional.String**| trace header；可由 APIv4 统一生成 | 

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
    aiHubPortfolioStopRequest := gateapi.AiHubPortfolioStopRequest{} // AiHubPortfolioStopRequest - 
    
    result, _, err := client.BotApi.PostAIHubPortfolioStop(ctx, aiHubPortfolioStopRequest, nil)
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

[**AiHubPortfolioStopSuccessResponse**](AIHubPortfolioStopSuccessResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
