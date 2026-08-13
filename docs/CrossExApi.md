# CrossExApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCrossexRuleSymbols**](CrossExApi.md#ListCrossexRuleSymbols) | **Get** /crossex/rule/symbols | Query symbol information
[**ListCrossexRuleRiskLimits**](CrossExApi.md#ListCrossexRuleRiskLimits) | **Get** /crossex/rule/risk_limits | Query risk limit information
[**ListCrossexTransferCoins**](CrossExApi.md#ListCrossexTransferCoins) | **Get** /crossex/transfers/coin | Query supported transfer currencies
[**ListCrossexTransfers**](CrossExApi.md#ListCrossexTransfers) | **Get** /crossex/transfers | Query Fund Transfer History
[**CreateCrossexTransfer**](CrossExApi.md#CreateCrossexTransfer) | **Post** /crossex/transfers | Fund Transfer
[**CreateCrossexOrder**](CrossExApi.md#CreateCrossexOrder) | **Post** /crossex/orders | Create order
[**CancelBatchCrossexOrders**](CrossExApi.md#CancelBatchCrossexOrders) | **Post** /crossex/batch_cancel_orders | Batch cancel orders
[**GetCrossexOrder**](CrossExApi.md#GetCrossexOrder) | **Get** /crossex/orders/{order_id} | Query order details
[**UpdateCrossexOrder**](CrossExApi.md#UpdateCrossexOrder) | **Put** /crossex/orders/{order_id} | Modify Order
[**CancelCrossexOrder**](CrossExApi.md#CancelCrossexOrder) | **Delete** /crossex/orders/{order_id} | Cancel Order
[**CreateCrossexConvertQuote**](CrossExApi.md#CreateCrossexConvertQuote) | **Post** /crossex/convert/quote | Flash Swap Inquiry
[**CreateCrossexConvertOrder**](CrossExApi.md#CreateCrossexConvertOrder) | **Post** /crossex/convert/orders | Flash Swap Transaction
[**GetCrossexAccount**](CrossExApi.md#GetCrossexAccount) | **Get** /crossex/accounts | Query Account Assets
[**UpdateCrossexAccount**](CrossExApi.md#UpdateCrossexAccount) | **Put** /crossex/accounts | Modify Account Contract Position Mode and Account Mode
[**GetCrossexPositionsLeverage**](CrossExApi.md#GetCrossexPositionsLeverage) | **Get** /crossex/positions/leverage | Query Contract Trading Pair Leverage Multiplier
[**UpdateCrossexPositionsLeverage**](CrossExApi.md#UpdateCrossexPositionsLeverage) | **Post** /crossex/positions/leverage | Modify Contract Trading Pair Leverage Multiplier
[**GetCrossexMarginPositionsLeverage**](CrossExApi.md#GetCrossexMarginPositionsLeverage) | **Get** /crossex/margin_positions/leverage | Query Leveraged Trading Pair Leverage Multiplier
[**UpdateCrossexMarginPositionsLeverage**](CrossExApi.md#UpdateCrossexMarginPositionsLeverage) | **Post** /crossex/margin_positions/leverage | Modify Leveraged Trading Pair Leverage Multiplier
[**CloseCrossexPosition**](CrossExApi.md#CloseCrossexPosition) | **Post** /crossex/position | Full Close Position
[**GetCrossexPositionsMarginMode**](CrossExApi.md#GetCrossexPositionsMarginMode) | **Get** /crossex/positions/margin_mode | Get futures position margin mode
[**UpdateCrossexPositionsMarginMode**](CrossExApi.md#UpdateCrossexPositionsMarginMode) | **Post** /crossex/positions/margin_mode | Update futures position margin mode
[**UpdateCrossexPositionsMargin**](CrossExApi.md#UpdateCrossexPositionsMargin) | **Post** /crossex/positions/margin | Increase or decrease isolated margin
[**GetCrossexInterestRate**](CrossExApi.md#GetCrossexInterestRate) | **Get** /crossex/interest_rate | Query margin asset interest rates
[**GetCrossexFee**](CrossExApi.md#GetCrossexFee) | **Get** /crossex/fee | Query User Fee Rates
[**ListCrossexPositions**](CrossExApi.md#ListCrossexPositions) | **Get** /crossex/positions | Query Contract Positions
[**ListCrossexMarginPositions**](CrossExApi.md#ListCrossexMarginPositions) | **Get** /crossex/margin_positions | Query Leveraged Positions
[**ListCrossexAdlRank**](CrossExApi.md#ListCrossexAdlRank) | **Get** /crossex/adl_rank | Query ADL Position Reduction Ranking
[**ListCrossexOpenOrders**](CrossExApi.md#ListCrossexOpenOrders) | **Get** /crossex/open_orders | Query All Current Open Orders
[**ListCrossexHistoryOrders**](CrossExApi.md#ListCrossexHistoryOrders) | **Get** /crossex/history_orders | Query order history
[**ListCrossexHistoryPositions**](CrossExApi.md#ListCrossexHistoryPositions) | **Get** /crossex/history_positions | Query Contract Position History
[**ListCrossexHistoryMarginPositions**](CrossExApi.md#ListCrossexHistoryMarginPositions) | **Get** /crossex/history_margin_positions | Query Leveraged Position History
[**ListCrossexHistoryMarginInterests**](CrossExApi.md#ListCrossexHistoryMarginInterests) | **Get** /crossex/history_margin_interests | Query Leveraged Interest Deduction History
[**ListCrossexHistoryTrades**](CrossExApi.md#ListCrossexHistoryTrades) | **Get** /crossex/history_trades | Query filled history
[**ListCrossexAccountBook**](CrossExApi.md#ListCrossexAccountBook) | **Get** /crossex/account_book | Query Account Asset Change History
[**ListCrossexCoinDiscountRate**](CrossExApi.md#ListCrossexCoinDiscountRate) | **Get** /crossex/coin_discount_rate | Query Currency Discount Rate
[**ListCrossexMarketTickers**](CrossExApi.md#ListCrossexMarketTickers) | **Get** /crossex/market/tickers | Get exchange tickers
[**ListCrossexMarketFundingInfo**](CrossExApi.md#ListCrossexMarketFundingInfo) | **Get** /crossex/market/funding_info | Get exchange futures funding rate information


## ListCrossexRuleSymbols

> []Symbol ListCrossexRuleSymbols(ctx, optional)

Query symbol information

Query Trading Pair Information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexRuleSymbolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexRuleSymbolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbols** | **optional.String**| List of trading pairs, comma-separated. Example: BINANCE_FUTURE_ADA_USDT,OKX_FUTURE_ADA_USDT | 

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
    
    result, _, err := client.CrossExApi.ListCrossexRuleSymbols(ctx, nil)
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

[**[]Symbol**](Symbol.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexRuleRiskLimits

> []CrossexRiskLimit ListCrossexRuleRiskLimits(ctx, symbols)

Query risk limit information

Query risk limit information for futures/margin trading pairs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbols** | **string**| Trading Pair List, multiple separated by commas Example values: BINANCE_FUTURE_ADA_USDT,GATE_MARGIN_ADA_USDT | 

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
    symbols := "BINANCE_FUTURE_AAVE_USDT" // string - Trading Pair List, multiple separated by commas Example values: BINANCE_FUTURE_ADA_USDT,GATE_MARGIN_ADA_USDT
    
    result, _, err := client.CrossExApi.ListCrossexRuleRiskLimits(ctx, symbols)
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

[**[]CrossexRiskLimit**](CrossexRiskLimit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexTransferCoins

> []CrossexTransferCoin ListCrossexTransferCoins(ctx, optional)

Query supported transfer currencies

`est_fee`: On-chain withdrawal fee. When a fund transfer involves an on-chain withdrawal, the exchange charges this fee. This value is for reference only; the actual fee charged by the exchange applies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexTransferCoinsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexTransferCoinsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**coin** | **optional.String**| Query by specified currency name | 

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
    
    result, _, err := client.CrossExApi.ListCrossexTransferCoins(ctx, nil)
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

[**[]CrossexTransferCoin**](CrossexTransferCoin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexTransfers

> []CrossexTransferRecord ListCrossexTransfers(ctx, optional)

Query Fund Transfer History

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexTransfersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**coin** | **optional.String**| Query by specified currency name | 
**orderId** | **optional.String**| Supports querying by the order ID returned when creating an order (tx_id), as well as a user-defined custom ID specified at creation (text) | 
**from** | **optional.Int32**| Start timestamp for the query | 
**to** | **optional.Int32**| End timestamp for the query, defaults to current time if not specified | 
**page** | **optional.Int32**| Page number | 
**limit** | **optional.Int32**| Maximum number returned by list, max 1000 | 

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
    
    result, _, err := client.CrossExApi.ListCrossexTransfers(ctx, nil)
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

[**[]CrossexTransferRecord**](CrossexTransferRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateCrossexTransfer

> CrossexTransferResponse CreateCrossexTransfer(ctx, optional)

Fund Transfer

Rate limit: 10 requests per 10 seconds - In cross-exchange mode, when transferring USDT, either `from` or `to` must be `SPOT`, and the other side must be `CROSSEX`.   If `CROSSEX_${exchange_type}` (e.g. `CROSSEX_GATE`) is provided, it will be automatically treated as `CROSSEX`. - In isolated exchange mode, when transferring USDT, either `from` or `to` must be `CROSSEX_${exchange_type}`, and the other side must be `SPOT` or `CROSSEX_${exchange_type}`.   If `CROSSEX` is provided, it will be automatically treated as `CROSSEX_GATE`. - When transferring non-USDT assets to or from CrossEx, neither `from` nor `to` can be `CROSSEX`; `CROSSEX_${exchange_type}` must be explicitly specified. - When transferring non-USDT assets, transfers between `CROSSEX_{exchange_type}` accounts are supported, for example: from = `CROSSEX_BINANCE`, to = `CROSSEX_GATE` - When either side of the transfer is `CROSSEX_KRAKEN`, only USDT is supported for now. - When either side of the transfer is `CROSSEX_HYPERLIQUID`, the other side must be `SPOT`, and only USDC is supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CreateCrossexTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCrossexTransferOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexTransferRequest** | [**optional.Interface of CrossexTransferRequest**](CrossexTransferRequest.md)|  | 

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
    
    result, _, err := client.CrossExApi.CreateCrossexTransfer(ctx, nil)
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

[**CrossexTransferResponse**](CrossexTransferResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateCrossexOrder

> CrossexOrderActionResponse CreateCrossexOrder(ctx, optional)

Create order

Rate Limit: 100 requests per 10 seconds, maximum 1,000 open orders per user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CreateCrossexOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCrossexOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexOrderRequest** | [**optional.Interface of CrossexOrderRequest**](CrossexOrderRequest.md)|  | 

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
    
    result, _, err := client.CrossExApi.CreateCrossexOrder(ctx, nil)
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

[**CrossexOrderActionResponse**](CrossexOrderActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelBatchCrossexOrders

> []CrossexBatchCancelOrderResponse CancelBatchCrossexOrders(ctx, crossexBatchCancelOrderRequest)

Batch cancel orders

Cancel multiple specified orders. Either order_id or text is required; if both are provided, order_id takes precedence. Rate limit: 100 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crossexBatchCancelOrderRequest** | [**[]CrossexBatchCancelOrderRequest**](CrossexBatchCancelOrderRequest.md)|  | 

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
    crossexBatchCancelOrderRequest := []gateapi.CrossexBatchCancelOrderRequest{gateapi.CrossexBatchCancelOrderRequest{}} // []CrossexBatchCancelOrderRequest - 
    
    result, _, err := client.CrossExApi.CancelBatchCrossexOrders(ctx, crossexBatchCancelOrderRequest)
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

[**[]CrossexBatchCancelOrderResponse**](CrossexBatchCancelOrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCrossexOrder

> CrossexOrder GetCrossexOrder(ctx, orderId)

Query order details

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| 1. Supports querying order IDs returned when creating orders 2. Supports custom IDs specified by users when creating orders (i.e., the text field) | 

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
    orderId := "2048522992198912" // string - 1. Supports querying order IDs returned when creating orders 2. Supports custom IDs specified by users when creating orders (i.e., the text field)
    
    result, _, err := client.CrossExApi.GetCrossexOrder(ctx, orderId)
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

[**CrossexOrder**](CrossexOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCrossexOrder

> CrossexOrderActionResponse UpdateCrossexOrder(ctx, orderId, optional)

Modify Order

Rate Limit: 100 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Support Order ID or Text for Modify Order | 
**optional** | **UpdateCrossexOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCrossexOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexOrderUpdateRequest** | [**optional.Interface of CrossexOrderUpdateRequest**](CrossexOrderUpdateRequest.md)|  | 

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
    orderId := "orderId_example" // string - Support Order ID or Text for Modify Order
    
    result, _, err := client.CrossExApi.UpdateCrossexOrder(ctx, orderId, nil)
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

[**CrossexOrderActionResponse**](CrossexOrderActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelCrossexOrder

> CrossexOrderActionResponse CancelCrossexOrder(ctx, orderId)

Cancel Order

Rate Limit: 100 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Support Order ID or Text for Cancel Order | 

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
    orderId := "orderId_example" // string - Support Order ID or Text for Cancel Order
    
    result, _, err := client.CrossExApi.CancelCrossexOrder(ctx, orderId)
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

[**CrossexOrderActionResponse**](CrossexOrderActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateCrossexConvertQuote

> CrossexConvertQuoteResponse CreateCrossexConvertQuote(ctx, optional)

Flash Swap Inquiry

Rate limit: 100 requests per day For HYPERLIQUID, swaps between `HYPERLIQUID_USDC` and `CROSSEX_USDT` are supported. Flash Swap in isolated exchange mode is not currently supported for HYPERLIQUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CreateCrossexConvertQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCrossexConvertQuoteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexConvertQuoteRequest** | [**optional.Interface of CrossexConvertQuoteRequest**](CrossexConvertQuoteRequest.md)|  | 

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
    
    result, _, err := client.CrossExApi.CreateCrossexConvertQuote(ctx, nil)
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

[**CrossexConvertQuoteResponse**](CrossexConvertQuoteResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateCrossexConvertOrder

> CrossexConvertOrderResponse CreateCrossexConvertOrder(ctx, optional)

Flash Swap Transaction

Rate limit: 10 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CreateCrossexConvertOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCrossexConvertOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexConvertOrderRequest** | [**optional.Interface of CrossexConvertOrderRequest**](CrossexConvertOrderRequest.md)|  | 

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
    
    result, _, err := client.CrossExApi.CreateCrossexConvertOrder(ctx, nil)
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

[**CrossexConvertOrderResponse**](CrossexConvertOrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCrossexAccount

> CrossexAccount GetCrossexAccount(ctx, optional)

Query Account Assets

Rate Limit: 200 requests per 10 seconds If 100% ≤ initial_margin_rate < 110%, transferring out the margin currency is prohibited. If initial_margin_rate < 100%, the system will automatically cancel orders; only closing positions is allowed, not opening new ones. If maintenance_margin_rate ≤ 100%, the system will force liquidation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCrossexAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCrossexAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**exchangeType** | **optional.String**| Trading venue identifier. Omit in cross-exchange mode; required in isolated-per-venue mode (&#x60;BINANCE&#x60; / &#x60;OKX&#x60; / &#x60;GATE&#x60; / &#x60;BYBIT&#x60; / &#x60;KRAKEN&#x60; / &#x60;HYPERLIQUID&#x60; / &#x60;DERIBIT&#x60;). | 

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
    
    result, _, err := client.CrossExApi.GetCrossexAccount(ctx, nil)
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

[**CrossexAccount**](CrossexAccount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCrossexAccount

> CrossexAccountUpdateResponse UpdateCrossexAccount(ctx, optional)

Modify Account Contract Position Mode and Account Mode

Rate Limit: 100 requests per 60 seconds. position_mode+exchange_type modifies contract position mode (exchange_type is required when the user's account mode is split exchange); account_mode modifies the user's account mode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **UpdateCrossexAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCrossexAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexAccountUpdateRequest** | [**optional.Interface of CrossexAccountUpdateRequest**](CrossexAccountUpdateRequest.md)|  | 

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
    
    result, _, err := client.CrossExApi.UpdateCrossexAccount(ctx, nil)
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

[**CrossexAccountUpdateResponse**](CrossexAccountUpdateResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCrossexPositionsLeverage

> map[string]string GetCrossexPositionsLeverage(ctx, optional)

Query Contract Trading Pair Leverage Multiplier

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCrossexPositionsLeverageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCrossexPositionsLeverageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbols** | **optional.String**| Trading Pair List, multiple separated by commas | 

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
    
    result, _, err := client.CrossExApi.GetCrossexPositionsLeverage(ctx, nil)
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

**map[string]string**

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCrossexPositionsLeverage

> CrossexLeverageResponse UpdateCrossexPositionsLeverage(ctx, optional)

Modify Contract Trading Pair Leverage Multiplier

Rate Limit: 100 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **UpdateCrossexPositionsLeverageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCrossexPositionsLeverageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexLeverageRequest** | [**optional.Interface of CrossexLeverageRequest**](CrossexLeverageRequest.md)|  | 

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
    
    result, _, err := client.CrossExApi.UpdateCrossexPositionsLeverage(ctx, nil)
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

[**CrossexLeverageResponse**](CrossexLeverageResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCrossexMarginPositionsLeverage

> map[string]string GetCrossexMarginPositionsLeverage(ctx, optional)

Query Leveraged Trading Pair Leverage Multiplier

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCrossexMarginPositionsLeverageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCrossexMarginPositionsLeverageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbols** | **optional.String**| Trading Pair List, multiple separated by commas | 

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
    
    result, _, err := client.CrossExApi.GetCrossexMarginPositionsLeverage(ctx, nil)
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

**map[string]string**

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCrossexMarginPositionsLeverage

> CrossexLeverageResponse UpdateCrossexMarginPositionsLeverage(ctx, optional)

Modify Leveraged Trading Pair Leverage Multiplier

Rate Limit: 100 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **UpdateCrossexMarginPositionsLeverageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCrossexMarginPositionsLeverageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexLeverageRequest** | [**optional.Interface of CrossexLeverageRequest**](CrossexLeverageRequest.md)|  | 

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
    
    result, _, err := client.CrossExApi.UpdateCrossexMarginPositionsLeverage(ctx, nil)
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

[**CrossexLeverageResponse**](CrossexLeverageResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CloseCrossexPosition

> CrossexOrderActionResponse CloseCrossexPosition(ctx, optional)

Full Close Position

Rate limit: 100 requests per day. Automatic position-closing rules. FUTURE and MARGIN positions are supported.  Before using this endpoint, ensure that the following prerequisite is met: - There are no open orders for the symbol in the current account. - Once the prerequisite is met, the system checks whether the position meets either of the following conditions: - Less than the minimum notional amount (minNotional) - Less than the minimum order size (minSize)  When either condition is met, the system automatically creates a closing order and immediately closes the entire position. This endpoint prevents positions that are too small to be submitted to an exchange from becoming stranded and ensures that small positions can be closed when they fall below the threshold.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CloseCrossexPositionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloseCrossexPositionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexClosePositionRequest** | [**optional.Interface of CrossexClosePositionRequest**](CrossexClosePositionRequest.md)|  | 

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
    
    result, _, err := client.CrossExApi.CloseCrossexPosition(ctx, nil)
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

[**CrossexOrderActionResponse**](CrossexOrderActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCrossexPositionsMarginMode

> CrossexMarginModeResponse GetCrossexPositionsMarginMode(ctx, symbol)

Get futures position margin mode

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Futures trading pair | 

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
    symbol := "HYPERLIQUID_FUTURE_CXMT_USDC" // string - Futures trading pair
    
    result, _, err := client.CrossExApi.GetCrossexPositionsMarginMode(ctx, symbol)
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

[**CrossexMarginModeResponse**](CrossexMarginModeResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCrossexPositionsMarginMode

> CrossexMarginModeResponse UpdateCrossexPositionsMarginMode(ctx, optional)

Update futures position margin mode

Rate limit: 100 requests per 10 seconds. Only Hyperliquid futures trading pairs are supported. The margin mode cannot be changed while open orders or positions exist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **UpdateCrossexPositionsMarginModeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCrossexPositionsMarginModeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexMarginModeRequest** | [**optional.Interface of CrossexMarginModeRequest**](CrossexMarginModeRequest.md)|  | 

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
    
    result, _, err := client.CrossExApi.UpdateCrossexPositionsMarginMode(ctx, nil)
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

[**CrossexMarginModeResponse**](CrossexMarginModeResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCrossexPositionsMargin

> CrossexIsolatedMarginResponse UpdateCrossexPositionsMargin(ctx, optional)

Increase or decrease isolated margin

Rate limit: 100 requests per 10 seconds. Only Hyperliquid isolated futures positions are supported. Positive values increase margin, while negative values decrease margin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **UpdateCrossexPositionsMarginOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCrossexPositionsMarginOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**crossexIsolatedMarginRequest** | [**optional.Interface of CrossexIsolatedMarginRequest**](CrossexIsolatedMarginRequest.md)|  | 

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
    
    result, _, err := client.CrossExApi.UpdateCrossexPositionsMargin(ctx, nil)
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

[**CrossexIsolatedMarginResponse**](CrossexIsolatedMarginResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCrossexInterestRate

> []CrossexInterestRate GetCrossexInterestRate(ctx, optional)

Query margin asset interest rates

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetCrossexInterestRateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCrossexInterestRateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**coin** | **optional.String**| Query by specified currency name | 
**exchangeType** | **optional.String**| Exchange | 

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
    
    result, _, err := client.CrossExApi.GetCrossexInterestRate(ctx, nil)
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

[**[]CrossexInterestRate**](CrossexInterestRate.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCrossexFee

> []InlineResponse200 GetCrossexFee(ctx, )

Query User Fee Rates

Rate Limit: 200 requests per 10 seconds

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
    
    result, _, err := client.CrossExApi.GetCrossexFee(ctx)
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

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexPositions

> []CrossexPosition ListCrossexPositions(ctx, optional)

Query Contract Positions

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexPositionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbol** | **optional.String**| Trading Pair | 
**exchangeType** | **optional.String**| Exchange | 

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
    
    result, _, err := client.CrossExApi.ListCrossexPositions(ctx, nil)
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

[**[]CrossexPosition**](CrossexPosition.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexMarginPositions

> []CrossexMarginPosition ListCrossexMarginPositions(ctx, optional)

Query Leveraged Positions

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexMarginPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexMarginPositionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbol** | **optional.String**| Currency pair | 
**exchangeType** | **optional.String**| Exchange | 

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
    
    result, _, err := client.CrossExApi.ListCrossexMarginPositions(ctx, nil)
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

[**[]CrossexMarginPosition**](CrossexMarginPosition.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexAdlRank

> CrossexAdlRank ListCrossexAdlRank(ctx, symbol)

Query ADL Position Reduction Ranking

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Trading Pair | 

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
    symbol := "BINANCE_FUTURE_ADA_USDT" // string - Trading Pair
    
    result, _, err := client.CrossExApi.ListCrossexAdlRank(ctx, symbol)
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

[**CrossexAdlRank**](CrossexAdlRank.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexOpenOrders

> []CrossexOrder ListCrossexOpenOrders(ctx, optional)

Query All Current Open Orders

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexOpenOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexOpenOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbol** | **optional.String**| Trading Pair | 
**exchangeType** | **optional.String**| Exchange | 
**businessType** | **optional.String**| Business Type | 

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
    
    result, _, err := client.CrossExApi.ListCrossexOpenOrders(ctx, nil)
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

[**[]CrossexOrder**](CrossexOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexHistoryOrders

> []CrossexOrder ListCrossexHistoryOrders(ctx, optional)

Query order history

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexHistoryOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexHistoryOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | 
**symbol** | **optional.String**| Currency pair | 
**from** | **optional.Int32**| Start Millisecond Timestamp | 
**to** | **optional.Int32**| End Millisecond Timestamp | 
**attributes** | **optional.String**| Order attributes (&#x60;COMMON&#x60; normal / &#x60;LIQ&#x60; liquidation takeover / &#x60;REDUCE&#x60; liquidation reduction / &#x60;ADL&#x60; auto-deleverage / &#x60;SETTLEMENT&#x60; delisting settlement). Multiple values, comma-separated. | 

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
    
    result, _, err := client.CrossExApi.ListCrossexHistoryOrders(ctx, nil)
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

[**[]CrossexOrder**](CrossexOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexHistoryPositions

> []CrossexHistoricalPosition ListCrossexHistoryPositions(ctx, optional)

Query Contract Position History

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexHistoryPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexHistoryPositionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | 
**limit** | **optional.Int32**| Maximum number returned by list, max 1000 | 
**symbol** | **optional.String**| Currency pair | 
**from** | **optional.Int32**| Start Millisecond Timestamp | 
**to** | **optional.Int32**| End Millisecond Timestamp | 

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
    
    result, _, err := client.CrossExApi.ListCrossexHistoryPositions(ctx, nil)
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

[**[]CrossexHistoricalPosition**](CrossexHistoricalPosition.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexHistoryMarginPositions

> []CrossexHistoricalMarginPosition ListCrossexHistoryMarginPositions(ctx, optional)

Query Leveraged Position History

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexHistoryMarginPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexHistoryMarginPositionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | 
**limit** | **optional.Int32**| Maximum number returned by list, max 1000 | 
**symbol** | **optional.String**| Currency pair | 
**from** | **optional.Int32**| Start Millisecond Timestamp | 
**to** | **optional.Int32**| End Millisecond Timestamp | 

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
    
    result, _, err := client.CrossExApi.ListCrossexHistoryMarginPositions(ctx, nil)
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

[**[]CrossexHistoricalMarginPosition**](CrossexHistoricalMarginPosition.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexHistoryMarginInterests

> []CrossexMarginInterestRecord ListCrossexHistoryMarginInterests(ctx, optional)

Query Leveraged Interest Deduction History

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexHistoryMarginInterestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexHistoryMarginInterestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbol** | **optional.String**| Currency pair | 
**from** | **optional.Int32**| Project-Id-Version: GateApiTools 1.0.0 Report-Msgid-Bugs-To: EMAIL@ADDRESS POT-Creation-Date: 2025-11-12 18:14+0800 PO-Revision-Date: 2019-01-02 17:30+0800 Last-Translator: FULL NAME &lt;EMAIL@ADDRESS&gt; Language: en Language-Team: en &lt;L@li.org&gt; Plural-Forms: nplurals&#x3D;2; plural&#x3D;(n !&#x3D;1) MIME-Version: 1.0 Content-Type: text/plain; charset&#x3D;utf-8 Content-Transfer-Encoding: 8bit Generated-By: Babel 2.8.0  | 
**to** | **optional.Int32**| Project-Id-Version: GateApiTools 1.0.0 Report-Msgid-Bugs-To: EMAIL@ADDRESS POT-Creation-Date: 2025-11-12 18:14+0800 PO-Revision-Date: 2019-01-02 17:30+0800 Last-Translator: FULL NAME &lt;EMAIL@ADDRESS&gt; Language: en Language-Team: en &lt;L@li.org&gt; Plural-Forms: nplurals&#x3D;2; plural&#x3D;(n !&#x3D;1) MIME-Version: 1.0 Content-Type: text/plain; charset&#x3D;utf-8 Content-Transfer-Encoding: 8bit Generated-By: Babel 2.8.0  | 
**page** | **optional.Int32**| Project-Id-Version: GateApiTools 1.0.0 Report-Msgid-Bugs-To: EMAIL@ADDRESS POT-Creation-Date: 2025-11-12 18:14+0800 PO-Revision-Date: 2019-01-02 17:30+0800 Last-Translator: FULL NAME &lt;EMAIL@ADDRESS&gt; Language: en Language-Team: en &lt;L@li.org&gt; Plural-Forms: nplurals&#x3D;2; plural&#x3D;(n !&#x3D;1) MIME-Version: 1.0 Content-Type: text/plain; charset&#x3D;utf-8 Content-Transfer-Encoding: 8bit Generated-By: Babel 2.8.0  | 
**limit** | **optional.Int32**| Project-Id-Version: GateApiTools 1.0.0 Report-Msgid-Bugs-To: EMAIL@ADDRESS POT-Creation-Date: 2025-11-12 18:14+0800 PO-Revision-Date: 2019-01-02 17:30+0800 Last-Translator: FULL NAME &lt;EMAIL@ADDRESS&gt; Language: en Language-Team: en &lt;L@li.org&gt; Plural-Forms: nplurals&#x3D;2; plural&#x3D;(n !&#x3D;1) MIME-Version: 1.0 Content-Type: text/plain; charset&#x3D;utf-8 Content-Transfer-Encoding: 8bit Generated-By: Babel 2.8.0  | 
**exchangeType** | **optional.String**| Project-Id-Version: GateApiTools 1.0.0 Report-Msgid-Bugs-To: EMAIL@ADDRESS POT-Creation-Date: 2025-11-12 18:14+0800 PO-Revision-Date: 2019-01-02 17:30+0800 Last-Translator: FULL NAME &lt;EMAIL@ADDRESS&gt; Language: en Language-Team: en &lt;L@li.org&gt; Plural-Forms: nplurals&#x3D;2; plural&#x3D;(n !&#x3D;1) MIME-Version: 1.0 Content-Type: text/plain; charset&#x3D;utf-8 Content-Transfer-Encoding: 8bit Generated-By: Babel 2.8.0  | 

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
    
    result, _, err := client.CrossExApi.ListCrossexHistoryMarginInterests(ctx, nil)
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

[**[]CrossexMarginInterestRecord**](CrossexMarginInterestRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexHistoryTrades

> []CrossexTrade ListCrossexHistoryTrades(ctx, optional)

Query filled history

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexHistoryTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexHistoryTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | 
**limit** | **optional.Int32**| Maximum number returned by list, max 1000 | 
**symbol** | **optional.String**| Currency pair | 
**from** | **optional.Int32**| Start Millisecond Timestamp | 
**to** | **optional.Int32**| End Millisecond Timestamp | 

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
    
    result, _, err := client.CrossExApi.ListCrossexHistoryTrades(ctx, nil)
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

[**[]CrossexTrade**](CrossexTrade.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexAccountBook

> []CrossexAccountBookRecord ListCrossexAccountBook(ctx, optional)

Query Account Asset Change History

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | 
**limit** | **optional.Int32**| Maximum number returned by list, max 1000 | 
**coin** | **optional.String**| Query by specified currency name | 
**statementType** | **optional.String**| Bill entry type. The filter accepts the same values returned in the response. | 
**from** | **optional.Int32**| Start Millisecond Timestamp | 
**to** | **optional.Int32**| End Millisecond Timestamp | 

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
    
    result, _, err := client.CrossExApi.ListCrossexAccountBook(ctx, nil)
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

[**[]CrossexAccountBookRecord**](CrossexAccountBookRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexCoinDiscountRate

> []CrossexCoinDiscountRate ListCrossexCoinDiscountRate(ctx, optional)

Query Currency Discount Rate

Rate Limit: 200 requests per 10 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexCoinDiscountRateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexCoinDiscountRateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**coin** | **optional.String**| Query by specified currency name | 
**exchangeType** | **optional.String**| OKX/GATE/BINANCE/BYBIT/KRAKEN/HYPERLIQUID/DERIBIT | 

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
    
    result, _, err := client.CrossExApi.ListCrossexCoinDiscountRate(ctx, nil)
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

[**[]CrossexCoinDiscountRate**](CrossexCoinDiscountRate.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexMarketTickers

> []InlineResponse2001 ListCrossexMarketTickers(ctx, optional)

Get exchange tickers

Rate limit: 1 request per second - Margin trading pairs cannot be passed directly as parameters. For example, `GATE_MARGIN_BTC_USDT` is invalid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexMarketTickersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexMarketTickersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbols** | **optional.String**| Trading Pair List, multiple separated by commas | 

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
    
    result, _, err := client.CrossExApi.ListCrossexMarketTickers(ctx, nil)
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

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossexMarketFundingInfo

> []InlineResponse2002 ListCrossexMarketFundingInfo(ctx, optional)

Get exchange futures funding rate information

Rate limit: 1 request per second - For `Deribit`, `funding_rate` is the current real-time rate calculated over an 8-hour period.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossexMarketFundingInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossexMarketFundingInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**symbols** | **optional.String**| Trading Pair List, multiple separated by commas | 

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
    
    result, _, err := client.CrossExApi.ListCrossexMarketFundingInfo(ctx, nil)
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

[**[]InlineResponse2002**](inline_response_200_2.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
