# DeliveryApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDeliveryContracts**](DeliveryApi.md#ListDeliveryContracts) | **Get** /delivery/{settle}/contracts | Query all futures contracts
[**GetDeliveryContract**](DeliveryApi.md#GetDeliveryContract) | **Get** /delivery/{settle}/contracts/{contract} | Query single contract information
[**ListDeliveryOrderBook**](DeliveryApi.md#ListDeliveryOrderBook) | **Get** /delivery/{settle}/order_book | Query futures market depth information
[**ListDeliveryTrades**](DeliveryApi.md#ListDeliveryTrades) | **Get** /delivery/{settle}/trades | Futures market transaction records
[**ListDeliveryCandlesticks**](DeliveryApi.md#ListDeliveryCandlesticks) | **Get** /delivery/{settle}/candlesticks | Futures market K-line chart
[**ListDeliveryTickers**](DeliveryApi.md#ListDeliveryTickers) | **Get** /delivery/{settle}/tickers | Get all futures trading statistics
[**ListDeliveryInsuranceLedger**](DeliveryApi.md#ListDeliveryInsuranceLedger) | **Get** /delivery/{settle}/insurance | Futures market insurance fund history
[**ListDeliveryAccounts**](DeliveryApi.md#ListDeliveryAccounts) | **Get** /delivery/{settle}/accounts | Get futures account
[**ListDeliveryAccountBook**](DeliveryApi.md#ListDeliveryAccountBook) | **Get** /delivery/{settle}/account_book | Query futures account change history
[**ListDeliveryPositions**](DeliveryApi.md#ListDeliveryPositions) | **Get** /delivery/{settle}/positions | Get user position list
[**GetDeliveryPosition**](DeliveryApi.md#GetDeliveryPosition) | **Get** /delivery/{settle}/positions/{contract} | Get single position information
[**UpdateDeliveryPositionMargin**](DeliveryApi.md#UpdateDeliveryPositionMargin) | **Post** /delivery/{settle}/positions/{contract}/margin | Update position margin
[**UpdateDeliveryPositionLeverage**](DeliveryApi.md#UpdateDeliveryPositionLeverage) | **Post** /delivery/{settle}/positions/{contract}/leverage | Update position leverage
[**UpdateDeliveryPositionRiskLimit**](DeliveryApi.md#UpdateDeliveryPositionRiskLimit) | **Post** /delivery/{settle}/positions/{contract}/risk_limit | Update position risk limit
[**ListDeliveryOrders**](DeliveryApi.md#ListDeliveryOrders) | **Get** /delivery/{settle}/orders | Query futures order list
[**CreateDeliveryOrder**](DeliveryApi.md#CreateDeliveryOrder) | **Post** /delivery/{settle}/orders | Place futures order
[**CancelDeliveryOrders**](DeliveryApi.md#CancelDeliveryOrders) | **Delete** /delivery/{settle}/orders | Cancel all orders with &#39;open&#39; status
[**GetDeliveryOrder**](DeliveryApi.md#GetDeliveryOrder) | **Get** /delivery/{settle}/orders/{order_id} | Query single order details
[**CancelDeliveryOrder**](DeliveryApi.md#CancelDeliveryOrder) | **Delete** /delivery/{settle}/orders/{order_id} | Cancel single order
[**GetMyDeliveryTrades**](DeliveryApi.md#GetMyDeliveryTrades) | **Get** /delivery/{settle}/my_trades | Query personal trading records
[**ListDeliveryPositionClose**](DeliveryApi.md#ListDeliveryPositionClose) | **Get** /delivery/{settle}/position_close | Query position close history
[**ListDeliveryLiquidates**](DeliveryApi.md#ListDeliveryLiquidates) | **Get** /delivery/{settle}/liquidates | Query liquidation history
[**ListDeliverySettlements**](DeliveryApi.md#ListDeliverySettlements) | **Get** /delivery/{settle}/settlements | Query settlement records
[**ListDeliveryRiskLimitTiers**](DeliveryApi.md#ListDeliveryRiskLimitTiers) | **Get** /delivery/{settle}/risk_limit_tiers | Query risk limit tiers
[**ListPriceTriggeredDeliveryOrders**](DeliveryApi.md#ListPriceTriggeredDeliveryOrders) | **Get** /delivery/{settle}/price_orders | Query auto order list
[**CreatePriceTriggeredDeliveryOrder**](DeliveryApi.md#CreatePriceTriggeredDeliveryOrder) | **Post** /delivery/{settle}/price_orders | Create price-triggered order
[**CancelPriceTriggeredDeliveryOrderList**](DeliveryApi.md#CancelPriceTriggeredDeliveryOrderList) | **Delete** /delivery/{settle}/price_orders | Cancel all auto orders
[**GetPriceTriggeredDeliveryOrder**](DeliveryApi.md#GetPriceTriggeredDeliveryOrder) | **Get** /delivery/{settle}/price_orders/{order_id} | Query single auto order details
[**CancelPriceTriggeredDeliveryOrder**](DeliveryApi.md#CancelPriceTriggeredDeliveryOrder) | **Delete** /delivery/{settle}/price_orders/{order_id} | Cancel single auto order


## ListDeliveryContracts

> []DeliveryContract ListDeliveryContracts(ctx, settle)

Query all futures contracts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliveryContracts(ctx, settle)
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

[**[]DeliveryContract**](DeliveryContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetDeliveryContract

> DeliveryContract GetDeliveryContract(ctx, settle, contract)

Query single contract information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**contract** | **string**| Futures contract | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    settle := "usdt" // string - Settle currency
    contract := "BTC_USDT_20200814" // string - Futures contract
    
    result, _, err := client.DeliveryApi.GetDeliveryContract(ctx, settle, contract)
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

[**DeliveryContract**](DeliveryContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryOrderBook

> FuturesOrderBook ListDeliveryOrderBook(ctx, settle, contract, optional)

Query futures market depth information

Bids will be sorted by price from high to low, while asks sorted reversely

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**contract** | **string**| Futures contract | 
**optional** | **ListDeliveryOrderBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliveryOrderBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**interval** | **optional.String**| Price precision for depth aggregation, 0 means no aggregation, defaults to 0 if not specified | [default to 0]
**limit** | **optional.Int32**| Number of depth levels | [default to 10]
**withId** | **optional.Bool**| Whether to return depth update ID. This ID increments by 1 each time depth changes | [default to false]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    settle := "usdt" // string - Settle currency
    contract := "BTC_USDT_20200814" // string - Futures contract
    
    result, _, err := client.DeliveryApi.ListDeliveryOrderBook(ctx, settle, contract, nil)
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

[**FuturesOrderBook**](FuturesOrderBook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryTrades

> []FuturesTrade ListDeliveryTrades(ctx, settle, contract, optional)

Futures market transaction records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**contract** | **string**| Futures contract | 
**optional** | **ListDeliveryTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliveryTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**lastId** | **optional.String**| 以上个列表的最后一条记录的 ID 作为下个列表的起点。 该字段不再继续支持，新的请求请使用 &#x60;from&#x60; 和 &#x60;to&#x60; 字段来限定时间范围 | 
**from** | **optional.Int64**| Specify starting time in Unix seconds. If not specified, &#x60;to&#x60; and &#x60;limit&#x60; will be used to limit response items. If items between &#x60;from&#x60; and &#x60;to&#x60; are more than &#x60;limit&#x60;, only &#x60;limit&#x60; number will be returned.  | 
**to** | **optional.Int64**| Specify end time in Unix seconds, default to current time. | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    settle := "usdt" // string - Settle currency
    contract := "BTC_USDT_20200814" // string - Futures contract
    
    result, _, err := client.DeliveryApi.ListDeliveryTrades(ctx, settle, contract, nil)
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

[**[]FuturesTrade**](FuturesTrade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryCandlesticks

> []DeliveryCandlestick ListDeliveryCandlesticks(ctx, settle, contract, optional)

Futures market K-line chart

Return specified contract candlesticks. If prefix `contract` with `mark_`, the contract's mark price candlesticks are returned; if prefix with `index_`, index price candlesticks will be returned.  Maximum of 2000 points are returned in one query. Be sure not to exceed the limit when specifying `from`, `to` and `interval`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**contract** | **string**| Futures contract | 
**optional** | **ListDeliveryCandlesticksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliveryCandlesticksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**from** | **optional.Int64**| Start time of candlesticks, formatted in Unix timestamp in seconds. Default to&#x60;to - 100 * interval&#x60; if not specified | 
**to** | **optional.Int64**| Specify the end time of the K-line chart, defaults to current time if not specified, note that the time format is Unix timestamp with second precision | 
**limit** | **optional.Int32**| Maximum number of recent data points to return. &#x60;limit&#x60; conflicts with &#x60;from&#x60; and &#x60;to&#x60;. If either &#x60;from&#x60; or &#x60;to&#x60; is specified, request will be rejected. | [default to 100]
**interval** | **optional.String**| Time interval between data points, note that 1w represents a natural week, 7d time is aligned with Unix initial timeTime interval between data points, note that 1w represents a natural week, 7d time is aligned with Unix initial timeweek, 7d time is aligned with Unix initial time | [default to 5m]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    settle := "usdt" // string - Settle currency
    contract := "BTC_USDT_20200814" // string - Futures contract
    
    result, _, err := client.DeliveryApi.ListDeliveryCandlesticks(ctx, settle, contract, nil)
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

[**[]DeliveryCandlestick**](DeliveryCandlestick.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryTickers

> []DeliveryTicker ListDeliveryTickers(ctx, settle, optional)

Get all futures trading statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**optional** | **ListDeliveryTickersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliveryTickersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Futures contract | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliveryTickers(ctx, settle, nil)
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

[**[]DeliveryTicker**](DeliveryTicker.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryInsuranceLedger

> []InsuranceRecord ListDeliveryInsuranceLedger(ctx, settle, optional)

Futures market insurance fund history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**optional** | **ListDeliveryInsuranceLedgerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliveryInsuranceLedgerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliveryInsuranceLedger(ctx, settle, nil)
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

[**[]InsuranceRecord**](InsuranceRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryAccounts

> FuturesAccount ListDeliveryAccounts(ctx, settle)

Get futures account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliveryAccounts(ctx, settle)
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

[**FuturesAccount**](FuturesAccount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryAccountBook

> []FuturesAccountBook ListDeliveryAccountBook(ctx, settle, optional)

Query futures account change history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**optional** | **ListDeliveryAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliveryAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 
**type_** | **optional.String**| Changing Type: - dnw: Deposit &amp; Withdraw - pnl: Profit &amp; Loss by reducing position - fee: Trading fee - refr: Referrer rebate - fund: Funding - point_dnw: point_fee: POINT Trading fee - point_refr: POINT Referrer rebate | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliveryAccountBook(ctx, settle, nil)
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

[**[]FuturesAccountBook**](FuturesAccountBook.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryPositions

> []Position ListDeliveryPositions(ctx, settle)

Get user position list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliveryPositions(ctx, settle)
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

[**[]Position**](Position.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetDeliveryPosition

> Position GetDeliveryPosition(ctx, settle, contract)

Get single position information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**contract** | **string**| Futures contract | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    contract := "BTC_USDT_20200814" // string - Futures contract
    
    result, _, err := client.DeliveryApi.GetDeliveryPosition(ctx, settle, contract)
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

[**Position**](Position.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateDeliveryPositionMargin

> Position UpdateDeliveryPositionMargin(ctx, settle, contract, change)

Update position margin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**contract** | **string**| Futures contract | 
**change** | **string**| Margin change amount, positive number increases, negative number decreases | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    contract := "BTC_USDT_20200814" // string - Futures contract
    change := "0.01" // string - Margin change amount, positive number increases, negative number decreases
    
    result, _, err := client.DeliveryApi.UpdateDeliveryPositionMargin(ctx, settle, contract, change)
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

[**Position**](Position.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateDeliveryPositionLeverage

> Position UpdateDeliveryPositionLeverage(ctx, settle, contract, leverage)

Update position leverage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**contract** | **string**| Futures contract | 
**leverage** | **string**| New position leverage | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    contract := "BTC_USDT_20200814" // string - Futures contract
    leverage := "10" // string - New position leverage
    
    result, _, err := client.DeliveryApi.UpdateDeliveryPositionLeverage(ctx, settle, contract, leverage)
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

[**Position**](Position.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateDeliveryPositionRiskLimit

> Position UpdateDeliveryPositionRiskLimit(ctx, settle, contract, riskLimit)

Update position risk limit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**contract** | **string**| Futures contract | 
**riskLimit** | **string**| New position risk limit | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    contract := "BTC_USDT_20200814" // string - Futures contract
    riskLimit := "10" // string - New position risk limit
    
    result, _, err := client.DeliveryApi.UpdateDeliveryPositionRiskLimit(ctx, settle, contract, riskLimit)
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

[**Position**](Position.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryOrders

> []FuturesOrder ListDeliveryOrders(ctx, settle, status, optional)

Query futures order list

Zero-fill orders cannot be retrieved 10 minutes after order cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**status** | **string**| Query order list based on status | 
**optional** | **ListDeliveryOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliveryOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Futures contract | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**lastId** | **optional.String**| Specify the currency name to query in batches, and support up to 100 pass parameters at a time | 
**countTotal** | **optional.Int32**| Whether to return total number matched, defaults to 0 (no return) | [default to 0]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    status := "open" // string - Query order list based on status
    
    result, _, err := client.DeliveryApi.ListDeliveryOrders(ctx, settle, status, nil)
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

[**[]FuturesOrder**](FuturesOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateDeliveryOrder

> FuturesOrder CreateDeliveryOrder(ctx, settle, futuresOrder)

Place futures order

Zero-fill orders cannot be retrieved 10 minutes after order cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**futuresOrder** | [**FuturesOrder**](FuturesOrder.md)|  | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    futuresOrder := gateapi.FuturesOrder{} // FuturesOrder - 
    
    result, _, err := client.DeliveryApi.CreateDeliveryOrder(ctx, settle, futuresOrder)
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

[**FuturesOrder**](FuturesOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelDeliveryOrders

> []FuturesOrder CancelDeliveryOrders(ctx, settle, contract, optional)

Cancel all orders with 'open' status

Zero-fill orders cannot be retrieved 10 minutes after order cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**contract** | **string**| Futures contract | 
**optional** | **CancelDeliveryOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelDeliveryOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**side** | **optional.String**| Specify all bids or all asks, both included if not specified | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    contract := "BTC_USDT_20200814" // string - Futures contract
    
    result, _, err := client.DeliveryApi.CancelDeliveryOrders(ctx, settle, contract, nil)
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

[**[]FuturesOrder**](FuturesOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetDeliveryOrder

> FuturesOrder GetDeliveryOrder(ctx, settle, orderId)

Query single order details

Zero-fill orders cannot be retrieved 10 minutes after order cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**orderId** | **string**| ID returned when order is successfully created | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    orderId := "12345" // string - ID returned when order is successfully created
    
    result, _, err := client.DeliveryApi.GetDeliveryOrder(ctx, settle, orderId)
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

[**FuturesOrder**](FuturesOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelDeliveryOrder

> FuturesOrder CancelDeliveryOrder(ctx, settle, orderId)

Cancel single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**orderId** | **string**| ID returned when order is successfully created | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    orderId := "12345" // string - ID returned when order is successfully created
    
    result, _, err := client.DeliveryApi.CancelDeliveryOrder(ctx, settle, orderId)
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

[**FuturesOrder**](FuturesOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetMyDeliveryTrades

> []MyFuturesTrade GetMyDeliveryTrades(ctx, settle, optional)

Query personal trading records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**optional** | **GetMyDeliveryTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMyDeliveryTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Futures contract | 
**order** | **optional.Int64**| Futures order ID, return related data only if specified | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**lastId** | **optional.String**| Specify the currency name to query in batches, and support up to 100 pass parameters at a time | 
**countTotal** | **optional.Int32**| Whether to return total number matched, defaults to 0 (no return) | [default to 0]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.GetMyDeliveryTrades(ctx, settle, nil)
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

[**[]MyFuturesTrade**](MyFuturesTrade.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryPositionClose

> []PositionClose ListDeliveryPositionClose(ctx, settle, optional)

Query position close history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**optional** | **ListDeliveryPositionCloseOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliveryPositionCloseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Futures contract | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliveryPositionClose(ctx, settle, nil)
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

[**[]PositionClose**](PositionClose.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryLiquidates

> []FuturesLiquidate ListDeliveryLiquidates(ctx, settle, optional)

Query liquidation history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**optional** | **ListDeliveryLiquidatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliveryLiquidatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Futures contract | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**at** | **optional.Int32**| Specify liquidation timestamp | [default to 0]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliveryLiquidates(ctx, settle, nil)
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

[**[]FuturesLiquidate**](FuturesLiquidate.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliverySettlements

> []DeliverySettlement ListDeliverySettlements(ctx, settle, optional)

Query settlement records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**optional** | **ListDeliverySettlementsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliverySettlementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Futures contract | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**at** | **optional.Int32**| Specify settlement timestamp | [default to 0]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliverySettlements(ctx, settle, nil)
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

[**[]DeliverySettlement**](DeliverySettlement.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeliveryRiskLimitTiers

> []FuturesLimitRiskTiers ListDeliveryRiskLimitTiers(ctx, settle, optional)

Query risk limit tiers

When the 'contract' parameter is not passed, the default is to query the risk limits for the top 100 markets.'Limit' and 'offset' correspond to pagination queries at the market level, not to the length of the returned array. This only takes effect empty.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**optional** | **ListDeliveryRiskLimitTiersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeliveryRiskLimitTiersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Futures contract | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliveryRiskLimitTiers(ctx, settle, nil)
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

[**[]FuturesLimitRiskTiers**](FuturesLimitRiskTiers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListPriceTriggeredDeliveryOrders

> []FuturesPriceTriggeredOrder ListPriceTriggeredDeliveryOrders(ctx, settle, status, optional)

Query auto order list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**status** | **string**| Query order list based on status | 
**optional** | **ListPriceTriggeredDeliveryOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPriceTriggeredDeliveryOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Futures contract, return related data only if specified | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    status := "status_example" // string - Query order list based on status
    
    result, _, err := client.DeliveryApi.ListPriceTriggeredDeliveryOrders(ctx, settle, status, nil)
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

[**[]FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreatePriceTriggeredDeliveryOrder

> TriggerOrderResponse CreatePriceTriggeredDeliveryOrder(ctx, settle, futuresPriceTriggeredOrder)

Create price-triggered order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**futuresPriceTriggeredOrder** | [**FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)|  | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    futuresPriceTriggeredOrder := gateapi.FuturesPriceTriggeredOrder{} // FuturesPriceTriggeredOrder - 
    
    result, _, err := client.DeliveryApi.CreatePriceTriggeredDeliveryOrder(ctx, settle, futuresPriceTriggeredOrder)
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

[**TriggerOrderResponse**](TriggerOrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelPriceTriggeredDeliveryOrderList

> []FuturesPriceTriggeredOrder CancelPriceTriggeredDeliveryOrderList(ctx, settle, contract)

Cancel all auto orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**contract** | **string**| Futures contract | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    contract := "BTC_USDT" // string - Futures contract
    
    result, _, err := client.DeliveryApi.CancelPriceTriggeredDeliveryOrderList(ctx, settle, contract)
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

[**[]FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPriceTriggeredDeliveryOrder

> FuturesPriceTriggeredOrder GetPriceTriggeredDeliveryOrder(ctx, settle, orderId)

Query single auto order details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**orderId** | **string**| ID returned when order is successfully created | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    orderId := "orderId_example" // string - ID returned when order is successfully created
    
    result, _, err := client.DeliveryApi.GetPriceTriggeredDeliveryOrder(ctx, settle, orderId)
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

[**FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelPriceTriggeredDeliveryOrder

> FuturesPriceTriggeredOrder CancelPriceTriggeredDeliveryOrder(ctx, settle, orderId)

Cancel single auto order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settle** | **string**| Settle currency | 
**orderId** | **string**| ID returned when order is successfully created | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    settle := "usdt" // string - Settle currency
    orderId := "orderId_example" // string - ID returned when order is successfully created
    
    result, _, err := client.DeliveryApi.CancelPriceTriggeredDeliveryOrder(ctx, settle, orderId)
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

[**FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
