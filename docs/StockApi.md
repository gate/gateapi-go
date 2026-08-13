# StockApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryStockUserAssets**](StockApi.md#QueryStockUserAssets) | **Get** /stock/users/assets | Query user assets
[**QueryStockSymbols**](StockApi.md#QueryStockSymbols) | **Get** /stock/symbols | Query symbol list
[**QueryStockSymbolDetail**](StockApi.md#QueryStockSymbolDetail) | **Get** /stock/symbols/detail | Query symbol details
[**QueryStockOrderBook**](StockApi.md#QueryStockOrderBook) | **Get** /stock/market/{symbol}/orderbook | Query market order book
[**QueryStockOrderList**](StockApi.md#QueryStockOrderList) | **Get** /stock/orders | Query open order list
[**CreateStockOrder**](StockApi.md#CreateStockOrder) | **Post** /stock/orders | Create order
[**DeleteAllStockOrders**](StockApi.md#DeleteAllStockOrders) | **Delete** /stock/orders | Cancel all open orders
[**QueryStockOrderHistory**](StockApi.md#QueryStockOrderHistory) | **Get** /stock/orders/history | Query historical order list
[**UpdateStockOrder**](StockApi.md#UpdateStockOrder) | **Put** /stock/orders/{order_id} | Modify order
[**DeleteStockOrder**](StockApi.md#DeleteStockOrder) | **Delete** /stock/orders/{order_id} | Cancel order
[**QueryStockPositions**](StockApi.md#QueryStockPositions) | **Get** /stock/positions | Query current position list
[**CloseStockPosition**](StockApi.md#CloseStockPosition) | **Post** /stock/positions/close | Close position
[**QueryStockTransactions**](StockApi.md#QueryStockTransactions) | **Get** /stock/transactions | Query transaction records
[**CreateStockTransaction**](StockApi.md#CreateStockTransaction) | **Post** /stock/transactions | Fund transfer
[**QueryStockExchanges**](StockApi.md#QueryStockExchanges) | **Get** /stock/exchanges | Query supported exchanges
[**QueryStockFeeRate**](StockApi.md#QueryStockFeeRate) | **Get** /stock/fee-rate | Query fee rates for Japanese and Korean stocks


## QueryStockUserAssets

> UserAssetResp2 QueryStockUserAssets(ctx, optional)

Query user assets

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryStockUserAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryStockUserAssetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pnlCalcType** | **optional.Int32**| PnL calculation cost type. Defaults to average cost price when omitted (1 &#x3D; average cost price, 2 &#x3D; diluted cost price) | 
**pnlCalcPrice** | **optional.Int32**| PnL calculation price type. Defaults to intraday price when omitted (1 &#x3D; intraday price, 2 &#x3D; latest extended-hours price) | 

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
    
    result, _, err := client.StockApi.QueryStockUserAssets(ctx, nil)
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

[**UserAssetResp2**](UserAssetResp_2.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryStockSymbols

> Symbols2 QueryStockSymbols(ctx, optional)

Query symbol list

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryStockSymbolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryStockSymbolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbols** | **optional.String**| Symbol list, multiple separated by commas | 
**exchange** | **optional.String**| Exchange, supports us, hk, and kr | 
**withDescI18n** | **optional.Bool**| Whether to return multilingual symbol description | 
**page** | **optional.Int32**| Page number, defaults to 1 | 
**pageSize** | **optional.Int32**| Page size, defaults to 10, max 500; server caps at 500 | 

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
    
    result, _, err := client.StockApi.QueryStockSymbols(ctx, nil)
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

[**Symbols2**](Symbols_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryStockSymbolDetail

> SymbolDetail QueryStockSymbolDetail(ctx, optional)

Query symbol details

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryStockSymbolDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryStockSymbolDetailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbols** | **optional.String**| Symbol list, multiple separated by commas | 
**exchange** | **optional.String**| Exchange, supports us, hk, and kr | 
**page** | **optional.Int32**| Page number, defaults to 1 | 
**pageSize** | **optional.Int32**| Page size, defaults to 10, max 500; server caps at 500 | 

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
    
    result, _, err := client.StockApi.QueryStockSymbolDetail(ctx, nil)
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

[**SymbolDetail**](SymbolDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryStockOrderBook

> OrderBook2 QueryStockOrderBook(ctx, symbol)

Query market order book

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Symbol | 

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
    symbol := "AAPL" // string - Symbol
    
    result, _, err := client.StockApi.QueryStockOrderBook(ctx, symbol)
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

[**OrderBook2**](OrderBook_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryStockOrderList

> OrderList2 QueryStockOrderList(ctx, optional)

Query open order list

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryStockOrderListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryStockOrderListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbol** | **optional.String**| Symbol | 

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
    
    result, _, err := client.StockApi.QueryStockOrderList(ctx, nil)
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

[**OrderList2**](OrderList_2.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateStockOrder

> CreateOrder2 CreateStockOrder(ctx, tradFiSpotOrderRequest)

Create order

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradFiSpotOrderRequest** | [**TradFiSpotOrderRequest**](TradFiSpotOrderRequest.md)|  | 

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
    tradFiSpotOrderRequest := gateapi.TradFiSpotOrderRequest{} // TradFiSpotOrderRequest - 
    
    result, _, err := client.StockApi.CreateStockOrder(ctx, tradFiSpotOrderRequest)
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

[**CreateOrder2**](CreateOrder_2.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteAllStockOrders

> DeleteOrder DeleteAllStockOrders(ctx, )

Cancel all open orders

Rate limit: 5 qps.

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
    
    result, _, err := client.StockApi.DeleteAllStockOrders(ctx)
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

[**DeleteOrder**](DeleteOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryStockOrderHistory

> OrderHistoryList2 QueryStockOrderHistory(ctx, optional)

Query historical order list

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryStockOrderHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryStockOrderHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbol** | **optional.String**| Symbol | 
**orderIds** | **optional.String**| Order ID list, multiple separated by commas; max 20, each must be a positive integer | 
**beginTime** | **optional.Int32**| Start time (Unix timestamp, seconds). When both begin_time and end_time are provided, end_time must be &gt;&#x3D; begin_time, query range must not exceed 3 months. | 
**endTime** | **optional.Int32**| End time (Unix timestamp, seconds). When both begin_time and end_time are provided, end_time must be &gt;&#x3D; begin_time, query range must not exceed 3 months. | 
**side** | **optional.Int32**| Side (1&#x3D;sell, 2&#x3D;buy) | 
**page** | **optional.Int32**| Page number, defaults to 1 | 
**pageSize** | **optional.Int32**| Page size, defaults to 10, max 500; server caps at 500 | 

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
    
    result, _, err := client.StockApi.QueryStockOrderHistory(ctx, nil)
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

[**OrderHistoryList2**](OrderHistoryList_2.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateStockOrder

> UpdateOrder2 UpdateStockOrder(ctx, orderId, tradFiSpotOrderUpdateRequest)

Modify order

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64**| Order ID | 
**tradFiSpotOrderUpdateRequest** | [**TradFiSpotOrderUpdateRequest**](TradFiSpotOrderUpdateRequest.md)|  | 

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
    orderId := 123456 // int64 - Order ID
    tradFiSpotOrderUpdateRequest := gateapi.TradFiSpotOrderUpdateRequest{} // TradFiSpotOrderUpdateRequest - 
    
    result, _, err := client.StockApi.UpdateStockOrder(ctx, orderId, tradFiSpotOrderUpdateRequest)
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

[**UpdateOrder2**](UpdateOrder_2.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteStockOrder

> DeleteOrder DeleteStockOrder(ctx, orderId)

Cancel order

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64**| Order ID | 

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
    orderId := 123456 // int64 - Order ID
    
    result, _, err := client.StockApi.DeleteStockOrder(ctx, orderId)
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

[**DeleteOrder**](DeleteOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryStockPositions

> PositionList2 QueryStockPositions(ctx, optional)

Query current position list

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryStockPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryStockPositionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pnlCalcType** | **optional.Int32**| PnL calculation cost type. Defaults to average cost price when omitted (1 &#x3D; average cost price, 2 &#x3D; diluted cost price) | 
**pnlCalcPrice** | **optional.Int32**| PnL calculation price type. Defaults to intraday price when omitted (1 &#x3D; intraday price, 2 &#x3D; latest extended-hours price) | 
**symbol** | **optional.String**| Symbol | 
**exchange** | **optional.String**| Exchange, supports us, hk, and kr | 

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
    
    result, _, err := client.StockApi.QueryStockPositions(ctx, nil)
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

[**PositionList2**](PositionList_2.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CloseStockPosition

> ClosePosition CloseStockPosition(ctx, tradFiSpotClosePositionRequest)

Close position

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradFiSpotClosePositionRequest** | [**TradFiSpotClosePositionRequest**](TradFiSpotClosePositionRequest.md)|  | 

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
    tradFiSpotClosePositionRequest := gateapi.TradFiSpotClosePositionRequest{} // TradFiSpotClosePositionRequest - 
    
    result, _, err := client.StockApi.CloseStockPosition(ctx, tradFiSpotClosePositionRequest)
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

[**ClosePosition**](ClosePosition.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryStockTransactions

> TransactionList2 QueryStockTransactions(ctx, optional)

Query transaction records

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryStockTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryStockTransactionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**beginTime** | **optional.Int64**| Start time (Unix timestamp, seconds). When both begin_time and end_time are provided, end_time must be &gt;&#x3D; begin_time, query range must not exceed 3 months. | 
**endTime** | **optional.Int64**| End time (Unix timestamp, seconds). When both begin_time and end_time are provided, end_time must be &gt;&#x3D; begin_time, query range must not exceed 3 months. | 
**refId** | **optional.String**| Business idempotent ID. When ref_id is provided, the server queries by ref_id, ignoring other parameters such as begin_time, end_time, type, page, page_size | 
**type_** | **optional.String**| Transaction type | 
**page** | **optional.Int32**| Page number, defaults to 1 | 
**pageSize** | **optional.Int32**| Page size, defaults to 10, max 500; server caps at 500 | 

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
    
    result, _, err := client.StockApi.QueryStockTransactions(ctx, nil)
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

[**TransactionList2**](TransactionList_2.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateStockTransaction

> CreateTransaction2 CreateStockTransaction(ctx, tradFiSpotTransactionRequest)

Fund transfer

Rate limit: 5 qps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradFiSpotTransactionRequest** | [**TradFiSpotTransactionRequest**](TradFiSpotTransactionRequest.md)|  | 

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
    tradFiSpotTransactionRequest := gateapi.TradFiSpotTransactionRequest{} // TradFiSpotTransactionRequest - 
    
    result, _, err := client.StockApi.CreateStockTransaction(ctx, tradFiSpotTransactionRequest)
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

[**CreateTransaction2**](CreateTransaction_2.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryStockExchanges

> Exchanges QueryStockExchanges(ctx, )

Query supported exchanges

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
    ctx := context.Background()
    
    result, _, err := client.StockApi.QueryStockExchanges(ctx)
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

[**Exchanges**](Exchanges.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryStockFeeRate

> FeeRate QueryStockFeeRate(ctx, )

Query fee rates for Japanese and Korean stocks

Query fee rates for Japanese and Korean stocks. Rate limit: 5 qps.

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
    ctx := context.Background()
    
    result, _, err := client.StockApi.QueryStockFeeRate(ctx)
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

[**FeeRate**](FeeRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
