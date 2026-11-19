# StockOptionsApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryOptionFeeRate**](StockOptionsApi.md#QueryOptionFeeRate) | **Get** /stock_options/fee-rate | Get option trading fee rates
[**QueryOptionSymbols**](StockOptionsApi.md#QueryOptionSymbols) | **Get** /stock_options/underlyings | List option underlyings
[**QueryOptionContracts**](StockOptionsApi.md#QueryOptionContracts) | **Get** /stock_options/contracts | List option contracts
[**QueryOptionExpirations**](StockOptionsApi.md#QueryOptionExpirations) | **Get** /stock_options/market/expirations | List option expiration dates
[**QueryOptionChain**](StockOptionsApi.md#QueryOptionChain) | **Get** /stock_options/market/chain | Get the option chain
[**QueryOptionOrderBook**](StockOptionsApi.md#QueryOptionOrderBook) | **Get** /stock_options/market/{symbol}/orderbook | Get the option order book
[**QueryOptionPositions**](StockOptionsApi.md#QueryOptionPositions) | **Get** /stock_options/positions | List current option positions
[**CloseOptionPosition**](StockOptionsApi.md#CloseOptionPosition) | **Post** /stock_options/positions/close | Close an option position
[**QueryOptionOrderList**](StockOptionsApi.md#QueryOptionOrderList) | **Get** /stock_options/orders | List current option orders
[**CreateOptionOrder**](StockOptionsApi.md#CreateOptionOrder) | **Post** /stock_options/orders | Create an option order
[**CancelAllOptionOrders**](StockOptionsApi.md#CancelAllOptionOrders) | **Delete** /stock_options/orders | Cancel all option limit orders
[**QueryOptionOrderDetail**](StockOptionsApi.md#QueryOptionOrderDetail) | **Get** /stock_options/orders/{order_id} | Get option order details
[**UpdateOptionOrder**](StockOptionsApi.md#UpdateOptionOrder) | **Put** /stock_options/orders/{order_id} | Amend an option limit order
[**CancelOptionOrder**](StockOptionsApi.md#CancelOptionOrder) | **Delete** /stock_options/orders/{order_id} | Cancel an option limit order
[**QueryOptionOrderHistory**](StockOptionsApi.md#QueryOptionOrderHistory) | **Get** /stock_options/orders/history | List historical option orders


## QueryOptionFeeRate

> StockOptionConfigResponse QueryOptionFeeRate(ctx, )

Get option trading fee rates

Query the fee configuration for options trading before placing an order. Returns the rates and applicable minimum and maximum fees for SEC, TAF, OCC, ORF and CAT charges.

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
    
    result, _, err := client.StockOptionsApi.QueryOptionFeeRate(ctx)
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

[**StockOptionConfigResponse**](StockOptionConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryOptionSymbols

> StockOptionSymbolListResponse QueryOptionSymbols(ctx, optional)

List option underlyings

Query available option underlyings before looking up contracts, expiration dates and option chains. Returns underlying symbols, names, multilingual descriptions, icons, price precision and order quantity limits. `underlyings` accepts multiple comma-separated symbols; omit it to query all available underlyings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryOptionSymbolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryOptionSymbolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**underlyings** | **optional.String**| Comma-separated option underlying symbols, for example AAPL,TSLA. Omit to query all available underlyings. | 

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
    
    result, _, err := client.StockOptionsApi.QueryOptionSymbols(ctx, nil)
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

[**StockOptionSymbolListResponse**](StockOptionSymbolListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryOptionContracts

> StockOptionContractListResponse QueryOptionContracts(ctx, underlying, optional)

List option contracts

Query option contracts for an underlying, including contract symbols, strike prices, expiration dates and trading rules. `underlying` is required. Filter by expiration date and call or put type; only US market (`us`) contracts are currently available. Results are paginated in ascending contract symbol order, with total records, total pages and the current page. The default and maximum `page_size` are both 1000. Use price precision and minimum and maximum order quantities to check an order before submission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Option underlying symbol, for example AAPL. This is one underlying symbol, not an option contract symbol. | 
**optional** | **QueryOptionContractsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryOptionContractsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**exchange** | **optional.String**| Trading market filter. Only US market (us) contracts are currently available. Omitted or empty values use us. - us: US market. | 
**expirationDate** | **optional.String**| Option expiration date in YYYY-MM-DD format. Omit to disable expiration date filtering. | 
**optionType** | [**optional.Interface of StockOptionOptionType**](.md)| Option type filter for the contract list. Omit for both types. - call: Call option. - put: Put option. | 
**page** | **optional.Int32**| Page number, starting from 1. Defaults to 1 when omitted. | [default to 1]
**pageSize** | **optional.Int32**| Contracts per page. Defaults to 1000 when omitted; values above 1000 are capped at 1000. | [default to 1000]

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
    underlying := "AAPL" // string - Option underlying symbol, for example AAPL. This is one underlying symbol, not an option contract symbol.
    
    result, _, err := client.StockOptionsApi.QueryOptionContracts(ctx, underlying, nil)
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

[**StockOptionContractListResponse**](StockOptionContractListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryOptionExpirations

> StockOptionExpirationListResponse QueryOptionExpirations(ctx, underlying)

List option expiration dates

Query expiration dates for an option underlying before selecting a contract expiration. `underlying` is required. Returns dates and their Unix timestamps; use the selected date to query the option chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Option underlying symbol, for example AAPL. This is one underlying symbol, not an option contract symbol. | 

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
    underlying := "AAPL" // string - Option underlying symbol, for example AAPL. This is one underlying symbol, not an option contract symbol.
    
    result, _, err := client.StockOptionsApi.QueryOptionExpirations(ctx, underlying)
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

[**StockOptionExpirationListResponse**](StockOptionExpirationListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryOptionChain

> StockOptionOptionChainResponse QueryOptionChain(ctx, underlying, expirationDate, optional)

Get the option chain

Query option quotes at different strike prices for the same underlying and expiration date to compare call and put contracts. `underlying` and `expiration_date` are required. Set `option_type` to `call` or `put` to filter; omit it or use `all` for both types. Returns contract symbols, strike prices, bid and ask quotes, last prices, mark prices, implied volatility and Greeks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Option underlying symbol, for example AAPL. This is one underlying symbol, not an option contract symbol. | 
**expirationDate** | **string**| Required option expiration date in YYYY-MM-DD format, obtained from the expiration date list. | 
**optional** | **QueryOptionChainOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryOptionChainOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**optionType** | **optional.String**| Option type filter for the option chain. - call: Call options only. - put: Put options only. - all: Both types, equivalent to omitting this parameter. | 

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
    underlying := "AAPL" // string - Option underlying symbol, for example AAPL. This is one underlying symbol, not an option contract symbol.
    expirationDate := "2026-08-14" // string - Required option expiration date in YYYY-MM-DD format, obtained from the expiration date list.
    
    result, _, err := client.StockOptionsApi.QueryOptionChain(ctx, underlying, expirationDate, nil)
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

[**StockOptionOptionChainResponse**](StockOptionOptionChainResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryOptionOrderBook

> StockOptionOrderBookResponse QueryOptionOrderBook(ctx, symbol)

Get the option order book

Query the best bid and ask for an option contract. `symbol` is the option contract symbol. Each of `bids` and `asks` contains at most one price level, without order quantities. A side without an available quote is an empty array. Both sides are empty arrays when the contract is not found.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Option contract symbol | 

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
    symbol := "AAPL260814P00130000" // string - Option contract symbol
    
    result, _, err := client.StockOptionsApi.QueryOptionOrderBook(ctx, symbol)
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

[**StockOptionOrderBookResponse**](StockOptionOrderBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryOptionPositions

> StockOptionPositionListResponse QueryOptionPositions(ctx, optional)

List current option positions

Query the current user's option positions, market values and profit and loss. Filter by `underlying` or `symbol`. A specified `symbol` must identify an existing option contract; no position returns an empty list. `price_type` selects the price for market value, unrealized PnL and today's PnL: `latest` (default) uses the last traded price; `mark` uses the mark price. Aggregate `today_pnl` includes current positions and contracts fully closed today. Results are sorted by `open_time` descending, then `symbol` ascending.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryOptionPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryOptionPositionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**underlying** | **optional.String**| Option underlying symbol, for example AAPL. Omit to disable underlying filtering. Current orders support only one underlying. | 
**symbol** | **optional.String**| Full option contract symbol, for example AAPL260814P00130000. Omit to disable contract filtering. | 
**priceType** | [**optional.Interface of StockOptionPositionPnlPriceType**](.md)| Valuation price type for position market value, unrealized PnL, and daily PnL. Options: &#x60;latest&#x60; (last traded price) or &#x60;mark&#x60; (mark price). Defaults to &#x60;latest&#x60; when omitted or empty. | 

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
    
    result, _, err := client.StockOptionsApi.QueryOptionPositions(ctx, nil)
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

[**StockOptionPositionListResponse**](StockOptionPositionListResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CloseOptionPosition

> StockOptionClosePositionResponse CloseOptionPosition(ctx, stockOptionClosePositionRequest)

Close an option position

Submit a closing order for part or all of an option position. `symbol` and `quantity` are required. The positive quantity must not exceed the available position and must satisfy the contract's minimum, maximum and quantity increment rules. Omit `order_type` or pass an empty value for a `market` order. A `limit` order requires a positive `price`; a market order ignores `price`. The closing side is derived from the position and time in force is fixed to `day`. The response returns the closing `order_id`; query the order for its subsequent status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockOptionClosePositionRequest** | [**StockOptionClosePositionRequest**](StockOptionClosePositionRequest.md)|  | 

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
    stockOptionClosePositionRequest := gateapi.StockOptionClosePositionRequest{} // StockOptionClosePositionRequest - 
    
    result, _, err := client.StockOptionsApi.CloseOptionPosition(ctx, stockOptionClosePositionRequest)
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

[**StockOptionClosePositionResponse**](StockOptionClosePositionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryOptionOrderList

> StockOptionOrderListResponse QueryOptionOrderList(ctx, optional)

List current option orders

Query the current user's unfinished option orders, including open and partially filled orders. Filter by `underlying`, `symbol` or `side`. `underlying` accepts only one symbol, not a comma-separated list. `status` includes only 1 (new) and 2 (partially filled). Query order history for terminal orders. Matching orders are sorted by creation time `time_setup` descending. This endpoint has no pagination parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryOptionOrderListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryOptionOrderListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**underlying** | **optional.String**| Option underlying symbol, for example AAPL. Omit to disable underlying filtering. Current orders support only one underlying. | 
**symbol** | **optional.String**| Full option contract symbol, for example AAPL260814P00130000. Omit to disable contract filtering. | 
**side** | [**optional.Interface of StockOptionOrderListSide**](.md)| Order side filter. Omit to query both sides. - 1: Sell orders only. - 2: Buy orders only. | 

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
    
    result, _, err := client.StockOptionsApi.QueryOptionOrderList(ctx, nil)
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

[**StockOptionOrderListResponse**](StockOptionOrderListResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateOptionOrder

> StockOptionPutOrderResponse CreateOptionOrder(ctx, stockOptionPutOrderRequest)

Create an option order

Submit a buy or sell order for an option contract and return the order ID on success. `symbol`, `side`, `order_type`, `time_in_force` and `quantity` are required. `side` is 1 (sell) or 2 (buy). Limit orders require a positive `price` and support `day` or `gtc`. Market orders ignore `price` and support only `day`. Quantity is a positive decimal string within the contract's minimum and maximum order quantities. Sell orders require sufficient available long positions. `client_order_id` is optional and is generated automatically when omitted or whitespace-only. Query current orders or order details to track the returned order ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockOptionPutOrderRequest** | [**StockOptionPutOrderRequest**](StockOptionPutOrderRequest.md)|  | 

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
    stockOptionPutOrderRequest := gateapi.StockOptionPutOrderRequest{} // StockOptionPutOrderRequest - 
    
    result, _, err := client.StockOptionsApi.CreateOptionOrder(ctx, stockOptionPutOrderRequest)
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

[**StockOptionPutOrderResponse**](StockOptionPutOrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelAllOptionOrders

> StockOptionCancelAllOrdersResponse CancelAllOptionOrders(ctx, optional)

Cancel all option limit orders

Cancel all or a filtered set of the current user's option limit orders. The request body may be omitted or an empty object. Without filters, all limit orders are cancelled. Filter by `underlying`, `symbol` or `side`. `side` supports 1 (sell) and 2 (buy); omit it to include both sides. The response reports the cancellation request result. Query the orders to confirm their subsequent status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CancelAllOptionOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelAllOptionOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**stockOptionCancelAllOrdersRequest** | [**optional.Interface of StockOptionCancelAllOrdersRequest**](StockOptionCancelAllOrdersRequest.md)|  | 

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
    
    result, _, err := client.StockOptionsApi.CancelAllOptionOrders(ctx, nil)
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

[**StockOptionCancelAllOrdersResponse**](StockOptionCancelAllOrdersResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryOptionOrderDetail

> StockOptionOrderDetailResponse QueryOptionOrderDetail(ctx, orderId)

Get option order details

Query one of the current user's orders by ID to track its status, filled quantity and average fill price. `order_id` is required and must be a positive integer represented as a string. Both current and finished orders can be queried, without filtering by terminal status. Returns price, quantity, filled quantity, status description, average fill price, fees and timestamps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Option order ID as a positive integer string, obtained from order submission or query results. | 

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
    orderId := "8000000001" // string - Option order ID as a positive integer string, obtained from order submission or query results.
    
    result, _, err := client.StockOptionsApi.QueryOptionOrderDetail(ctx, orderId)
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

[**StockOptionOrderDetailResponse**](StockOptionOrderDetailResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateOptionOrder

> StockOptionUpdateOrderResponse UpdateOptionOrder(ctx, orderId, stockOptionUpdateOrderRequest)

Amend an option limit order

Update the price and total quantity of one of the current user's option limit orders. The path `order_id` and body `price` and `quantity` are required. Market orders cannot be amended. `quantity` is the new total order quantity, not the remaining unfilled quantity, and must be at least `filled_quantity`. Amending a sell order also requires sufficient available positions for the new unfilled quantity. Returns the order ID; query the order for its subsequent status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Option order ID as a positive integer string, obtained from order submission or query results. | 
**stockOptionUpdateOrderRequest** | [**StockOptionUpdateOrderRequest**](StockOptionUpdateOrderRequest.md)|  | 

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
    orderId := "72057594037983901" // string - Option order ID as a positive integer string, obtained from order submission or query results.
    stockOptionUpdateOrderRequest := gateapi.StockOptionUpdateOrderRequest{} // StockOptionUpdateOrderRequest - 
    
    result, _, err := client.StockOptionsApi.UpdateOptionOrder(ctx, orderId, stockOptionUpdateOrderRequest)
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

[**StockOptionUpdateOrderResponse**](StockOptionUpdateOrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelOptionOrder

> StockOptionCancelOrderResponse CancelOptionOrder(ctx, orderId)

Cancel an option limit order

Cancel one of the current user's option limit orders by ID. `order_id` is required. Only limit orders can be cancelled; already filled quantities are unaffected. Returns the order ID. Query order details or order history for the final status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Option order ID as a positive integer string, obtained from order submission or query results. | 

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
    orderId := "72057594037983901" // string - Option order ID as a positive integer string, obtained from order submission or query results.
    
    result, _, err := client.StockOptionsApi.CancelOptionOrder(ctx, orderId)
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

[**StockOptionCancelOrderResponse**](StockOptionCancelOrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QueryOptionOrderHistory

> StockOptionOrderHistoryResponse QueryOptionOrderHistory(ctx, optional)

List historical option orders

Query the current user's finished option orders with pagination to check orders, fills and fees. `status` includes only 3 (filled), 4 (cancelled) and 5 (rejected). Filter by underlying, contract, side or order creation time. `begin_time` and `end_time` are Unix timestamps in seconds; when both are provided, the start must not be later than the end. Returns total records, total pages and the current page. The default and maximum `page_size` are both 100.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **QueryOptionOrderHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryOptionOrderHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**underlying** | **optional.String**| Option underlying symbol, for example AAPL. Omit to disable underlying filtering. Current orders support only one underlying. | 
**symbol** | **optional.String**| Full option contract symbol, for example AAPL260814P00130000. Omit to disable contract filtering. | 
**beginTime** | **optional.Int64**| Order creation time filter start, as a Unix timestamp in seconds. Omit for no lower bound; must not exceed end_time when both are provided. | 
**endTime** | **optional.Int64**| Order creation time filter end, as a Unix timestamp in seconds. Omit for no upper bound; must not precede begin_time when both are provided. | 
**side** | [**optional.Interface of StockOptionOrderListSide**](.md)| Order side filter. Omit to query both sides. - 1: Sell orders only. - 2: Buy orders only. | 
**page** | **optional.Int32**| Page number, starting from 1. Defaults to 1 when omitted. | [default to 1]
**pageSize** | **optional.Int32**| Number of items per page, up to 100. Defaults to 100 when omitted; values above 100 are capped at 100. | [default to 100]

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
    
    result, _, err := client.StockOptionsApi.QueryOptionOrderHistory(ctx, nil)
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

[**StockOptionOrderHistoryResponse**](StockOptionOrderHistoryResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
