# BotApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAIHubStrategyRecommend**](BotApi.md#GetAIHubStrategyRecommend) | **Get** /bot/strategy/recommend | Get AIHub strategy recommendations
[**PostAIHubSpotGridCreate**](BotApi.md#PostAIHubSpotGridCreate) | **Post** /bot/spot-grid/create | Create spot grid
[**PostAIHubMarginGridCreate**](BotApi.md#PostAIHubMarginGridCreate) | **Post** /bot/margin-grid/create | Create a lever grid
[**PostAIHubInfiniteGridCreate**](BotApi.md#PostAIHubInfiniteGridCreate) | **Post** /bot/infinite-grid/create | Create infinite grid
[**PostAIHubFuturesGridCreate**](BotApi.md#PostAIHubFuturesGridCreate) | **Post** /bot/futures-grid/create | Create a contract grid
[**PostAIHubSpotMartingaleCreate**](BotApi.md#PostAIHubSpotMartingaleCreate) | **Post** /bot/spot-martingale/create | Create Spot Martin
[**PostAIHubContractMartingaleCreate**](BotApi.md#PostAIHubContractMartingaleCreate) | **Post** /bot/contract-martingale/create | Create contract martin
[**GetAIHubPortfolioRunning**](BotApi.md#GetAIHubPortfolioRunning) | **Get** /bot/portfolio/running | Query the list of running policies
[**GetAIHubPortfolioDetail**](BotApi.md#GetAIHubPortfolioDetail) | **Get** /bot/portfolio/detail | Query order policy details
[**PostAIHubPortfolioStop**](BotApi.md#PostAIHubPortfolioStop) | **Post** /bot/portfolio/stop | Terminate a single running policy


## GetAIHubStrategyRecommend

> AiHubDiscoverSuccessResponse GetAIHubStrategyRecommend(ctx, optional)

Get AIHub strategy recommendations

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
**market** | **optional.String**| Trading pair, such as &#x60;BTC_USDT&#x60; | 
**strategyType** | **optional.String**| Recommended target policy type; &#x60;contract_martingale&#x60; not allowed | 
**direction** | **optional.String**| Market direction | 
**investAmount** | **optional.String**| Investment amount, string transparent transmission | 
**scene** | **optional.String**| Recommended scenario; when empty, bot-service can automatically infer according to the implementation logic. | 
**refreshRecommendationId** | **optional.String**| It is recommended to refresh the context. Used when &#x60;scene&#x3D;refresh&#x60; is used; when &#x60;scene&#x60; is empty but the field exists, bot-service will also automatically determine as &#x60;refresh&#x60;. The official minimum format is &#x60;strategy_type|market&#x60;; if the &#x60;recommendation_id&#x60; of the previous recommendation is directly passed through, the third paragraph &#x60;backtest_id&#x60; will be ignored. | 
**limit** | **optional.Int32**| Return quantity; when &#x60;scene&#x3D;filter&#x60; is used, the actual results are up to 10 | 
**maxDrawdownLte** | **optional.String**| Maximum drawdown limit | 
**backtestAprGte** | **optional.String**| Backtest annualized lower limit | 
**xGateServiceId** | **optional.String**| Call source identifier; injected by APIv4 if necessary | 
**xGateAppLang** | **optional.String**| Language context, such as &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| Request link ID; caller can transmit transparently | 
**xTraceId** | **optional.String**| trace header; can be generated uniformly by APIv4 | 

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

Create spot grid

Create a spot grid strategy based on the incoming parameters.

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
**xGateServiceId** | **optional.String**| Call source identifier; injected by APIv4 if necessary | 
**xGateAppLang** | **optional.String**| Language context, such as &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| Request link ID; caller can transmit transparently | 
**xTraceId** | **optional.String**| trace header; can be generated uniformly by APIv4 | 

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

Create a lever grid

Create a leverage grid strategy based on the passed parameters.

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
**xGateServiceId** | **optional.String**| Call source identifier; injected by APIv4 if necessary | 
**xGateAppLang** | **optional.String**| Language context, such as &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| Request link ID; caller can transmit transparently | 
**xTraceId** | **optional.String**| trace header; can be generated uniformly by APIv4 | 

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

Create infinite grid

Create an infinite grid strategy based on passed parameters.

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
**xGateServiceId** | **optional.String**| Call source identifier; injected by APIv4 if necessary | 
**xGateAppLang** | **optional.String**| Language context, such as &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| Request link ID; caller can transmit transparently | 
**xTraceId** | **optional.String**| trace header; can be generated uniformly by APIv4 | 

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

Create a contract grid

Create a contract grid strategy based on the incoming parameters.

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
**xGateServiceId** | **optional.String**| Call source identifier; injected by APIv4 if necessary | 
**xGateAppLang** | **optional.String**| Language context, such as &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| Request link ID; caller can transmit transparently | 
**xTraceId** | **optional.String**| trace header; can be generated uniformly by APIv4 | 

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

Create Spot Martin

根据传入参数创建现货马丁策略。  止损口径与 App / `MartingaleBot` 一致： - 使用 **`create_params.stop_loss_per_cycle`**（每轮止损比例，小数字符串），**不要**使用 `stop_loss_price` 表达创建侧止损。 - 详情页展示的「止损价」由引擎按轮次计算；创建侧可选 **`create_params.trigger_price`**（触发价）。

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
**xGateServiceId** | **optional.String**| Call source identifier; injected by APIv4 if necessary | 
**xGateAppLang** | **optional.String**| Language context, such as &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| Request link ID; caller can transmit transparently | 
**xTraceId** | **optional.String**| trace header; can be generated uniformly by APIv4 | 

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

Create contract martin

Create a contract Martin strategy based on the input parameters.

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
**xGateServiceId** | **optional.String**| Call source identifier; injected by APIv4 if necessary | 
**xGateAppLang** | **optional.String**| Language context, such as &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| Request link ID; caller can transmit transparently | 
**xTraceId** | **optional.String**| trace header; can be generated uniformly by APIv4 | 

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

Query the list of running policies

Query the list of AIHub strategies currently running by the user, and support filtering by strategy type, trading pair and paging conditions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetAIHubPortfolioRunningOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAIHubPortfolioRunningOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**strategyType** | **optional.String**| Filter by policy type | 
**market** | **optional.String**| Filter by trading pair | 
**page** | **optional.Int32**| Page number, default 1 | [default to 1]
**pageSize** | **optional.Int32**| Paging size, default 20, maximum 50 | [default to 20]
**xGateServiceId** | **optional.String**| Call source identifier; injected by APIv4 if necessary | 
**xGateAppLang** | **optional.String**| Language context, such as &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| Request link ID; caller can transmit transparently | 
**xTraceId** | **optional.String**| trace header; can be generated uniformly by APIv4 | 

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

Query order policy details

Both `strategy_id` and `strategy_type` must be passed in the request, where `strategy_type` is used to distribute to the underlying detailed implementation by strategy type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyId** | **string**| Policy ID | 
**strategyType** | **string**| Policy type; used for underlying detail distribution | 
**optional** | **GetAIHubPortfolioDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAIHubPortfolioDetailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateServiceId** | **optional.String**| Call source identifier; injected by APIv4 if necessary | 
**xGateAppLang** | **optional.String**| Language context, such as &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| Request link ID; caller can transmit transparently | 
**xTraceId** | **optional.String**| trace header; can be generated uniformly by APIv4 | 

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
    strategyId := "strategyId_example" // string - Policy ID
    strategyType := "strategyType_example" // string - Policy type; used for underlying detail distribution
    
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

Terminate a single running policy

Only one policy is allowed to be terminated per request. Risk warning and secondary confirmation are borne by the upper layer of OpenClaw; this interface is only responsible for executing stop.

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
**xGateServiceId** | **optional.String**| Call source identifier; injected by APIv4 if necessary | 
**xGateAppLang** | **optional.String**| Language context, such as &#x60;zh-CN&#x60; / &#x60;en-US&#x60; | 
**xRequestId** | **optional.String**| Request link ID; caller can transmit transparently | 
**xTraceId** | **optional.String**| trace header; can be generated uniformly by APIv4 | 

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
