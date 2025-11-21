# OptionsApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOptionsUnderlyings**](OptionsApi.md#ListOptionsUnderlyings) | **Get** /options/underlyings | List all underlying assets
[**ListOptionsExpirations**](OptionsApi.md#ListOptionsExpirations) | **Get** /options/expirations | List all expiration dates
[**ListOptionsContracts**](OptionsApi.md#ListOptionsContracts) | **Get** /options/contracts | List all contracts for specified underlying and expiration date
[**GetOptionsContract**](OptionsApi.md#GetOptionsContract) | **Get** /options/contracts/{contract} | Query specified contract details
[**ListOptionsSettlements**](OptionsApi.md#ListOptionsSettlements) | **Get** /options/settlements | List settlement history
[**GetOptionsSettlement**](OptionsApi.md#GetOptionsSettlement) | **Get** /options/settlements/{contract} | Get specified contract settlement information
[**ListMyOptionsSettlements**](OptionsApi.md#ListMyOptionsSettlements) | **Get** /options/my_settlements | Query personal settlement records
[**ListOptionsOrderBook**](OptionsApi.md#ListOptionsOrderBook) | **Get** /options/order_book | Query options contract order book
[**ListOptionsTickers**](OptionsApi.md#ListOptionsTickers) | **Get** /options/tickers | Query options market ticker information
[**ListOptionsUnderlyingTickers**](OptionsApi.md#ListOptionsUnderlyingTickers) | **Get** /options/underlying/tickers/{underlying} | Query underlying ticker information
[**ListOptionsCandlesticks**](OptionsApi.md#ListOptionsCandlesticks) | **Get** /options/candlesticks | Options contract market candlestick chart
[**ListOptionsUnderlyingCandlesticks**](OptionsApi.md#ListOptionsUnderlyingCandlesticks) | **Get** /options/underlying/candlesticks | Underlying index price candlestick chart
[**ListOptionsTrades**](OptionsApi.md#ListOptionsTrades) | **Get** /options/trades | Market trade records
[**ListOptionsAccount**](OptionsApi.md#ListOptionsAccount) | **Get** /options/accounts | Query account information
[**ListOptionsAccountBook**](OptionsApi.md#ListOptionsAccountBook) | **Get** /options/account_book | Query account change history
[**ListOptionsPositions**](OptionsApi.md#ListOptionsPositions) | **Get** /options/positions | List user&#39;s positions of specified underlying
[**GetOptionsPosition**](OptionsApi.md#GetOptionsPosition) | **Get** /options/positions/{contract} | Get specified contract position
[**ListOptionsPositionClose**](OptionsApi.md#ListOptionsPositionClose) | **Get** /options/position_close | List user&#39;s liquidation history of specified underlying
[**ListOptionsOrders**](OptionsApi.md#ListOptionsOrders) | **Get** /options/orders | List options orders
[**CreateOptionsOrder**](OptionsApi.md#CreateOptionsOrder) | **Post** /options/orders | Create an options order
[**CancelOptionsOrders**](OptionsApi.md#CancelOptionsOrders) | **Delete** /options/orders | Cancel all orders with &#39;open&#39; status
[**GetOptionsOrder**](OptionsApi.md#GetOptionsOrder) | **Get** /options/orders/{order_id} | Query single order details
[**CancelOptionsOrder**](OptionsApi.md#CancelOptionsOrder) | **Delete** /options/orders/{order_id} | Cancel single order
[**CountdownCancelAllOptions**](OptionsApi.md#CountdownCancelAllOptions) | **Post** /options/countdown_cancel_all | Countdown cancel orders
[**ListMyOptionsTrades**](OptionsApi.md#ListMyOptionsTrades) | **Get** /options/my_trades | Query personal trading records
[**GetOptionsMMP**](OptionsApi.md#GetOptionsMMP) | **Get** /options/mmp | MMP Query.
[**SetOptionsMMP**](OptionsApi.md#SetOptionsMMP) | **Post** /options/mmp | MMP Settings
[**ResetOptionsMMP**](OptionsApi.md#ResetOptionsMMP) | **Post** /options/mmp/reset | MMP Reset


## ListOptionsUnderlyings

> []OptionsUnderlying ListOptionsUnderlyings(ctx, )

List all underlying assets

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
    
    result, _, err := client.OptionsApi.ListOptionsUnderlyings(ctx)
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

[**[]OptionsUnderlying**](OptionsUnderlying.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsExpirations

> []int64 ListOptionsExpirations(ctx, underlying)

List all expiration dates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Underlying (Obtained by listing underlying endpoint) | 

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
    underlying := "BTC_USDT" // string - Underlying (Obtained by listing underlying endpoint)
    
    result, _, err := client.OptionsApi.ListOptionsExpirations(ctx, underlying)
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

**[]int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsContracts

> []OptionsContract ListOptionsContracts(ctx, underlying, optional)

List all contracts for specified underlying and expiration date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Underlying (Obtained by listing underlying endpoint) | 
**optional** | **ListOptionsContractsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOptionsContractsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**expiration** | **optional.Int64**| Unix timestamp of expiration date | 

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
    underlying := "BTC_USDT" // string - Underlying (Obtained by listing underlying endpoint)
    
    result, _, err := client.OptionsApi.ListOptionsContracts(ctx, underlying, nil)
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

[**[]OptionsContract**](OptionsContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOptionsContract

> OptionsContract GetOptionsContract(ctx, contract)

Query specified contract details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string**|  | 

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
    contract := "BTC_USDT-20211130-65000-C" // string - 
    
    result, _, err := client.OptionsApi.GetOptionsContract(ctx, contract)
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

[**OptionsContract**](OptionsContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsSettlements

> []OptionsSettlement ListOptionsSettlements(ctx, underlying, optional)

List settlement history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Underlying (Obtained by listing underlying endpoint) | 
**optional** | **ListOptionsSettlementsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOptionsSettlementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 

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
    underlying := "BTC_USDT" // string - Underlying (Obtained by listing underlying endpoint)
    
    result, _, err := client.OptionsApi.ListOptionsSettlements(ctx, underlying, nil)
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

[**[]OptionsSettlement**](OptionsSettlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOptionsSettlement

> OptionsSettlement GetOptionsSettlement(ctx, contract, underlying, at)

Get specified contract settlement information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string**|  | 
**underlying** | **string**| Underlying (Obtained by listing underlying endpoint) | 
**at** | **int64**|  | 

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
    contract := "BTC_USDT-20211130-65000-C" // string - 
    underlying := "BTC_USDT" // string - Underlying (Obtained by listing underlying endpoint)
    at := 56 // int64 - 
    
    result, _, err := client.OptionsApi.GetOptionsSettlement(ctx, contract, underlying, at)
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

[**OptionsSettlement**](OptionsSettlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListMyOptionsSettlements

> []OptionsMySettlements ListMyOptionsSettlements(ctx, underlying, optional)

Query personal settlement records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Underlying (Obtained by listing underlying endpoint) | 
**optional** | **ListMyOptionsSettlementsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMyOptionsSettlementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Options contract name | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 

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
    underlying := "BTC_USDT" // string - Underlying (Obtained by listing underlying endpoint)
    
    result, _, err := client.OptionsApi.ListMyOptionsSettlements(ctx, underlying, nil)
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

[**[]OptionsMySettlements**](OptionsMySettlements.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsOrderBook

> FuturesOrderBook ListOptionsOrderBook(ctx, contract, optional)

Query options contract order book

Bids will be sorted by price from high to low, while asks sorted reversely

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string**| Options contract name | 
**optional** | **ListOptionsOrderBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOptionsOrderBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**interval** | **optional.String**| Price precision for depth aggregation, 0 means no aggregation, defaults to 0 if not specified | [default to 0]
**limit** | **optional.Int32**| Number of depth levels | [default to 10]
**withId** | **optional.Bool**| Whether to return depth update ID. This ID increments by 1 each time the depth changes | [default to false]

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
    contract := "BTC_USDT-20210916-5000-C" // string - Options contract name
    
    result, _, err := client.OptionsApi.ListOptionsOrderBook(ctx, contract, nil)
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

## ListOptionsTickers

> []OptionsTicker ListOptionsTickers(ctx, underlying)

Query options market ticker information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Underlying (Obtained by listing underlying endpoint) | 

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
    underlying := "BTC_USDT" // string - Underlying (Obtained by listing underlying endpoint)
    
    result, _, err := client.OptionsApi.ListOptionsTickers(ctx, underlying)
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

[**[]OptionsTicker**](OptionsTicker.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsUnderlyingTickers

> OptionsUnderlyingTicker ListOptionsUnderlyingTickers(ctx, underlying)

Query underlying ticker information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Underlying | 

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
    underlying := "BTC_USDT" // string - Underlying
    
    result, _, err := client.OptionsApi.ListOptionsUnderlyingTickers(ctx, underlying)
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

[**OptionsUnderlyingTicker**](OptionsUnderlyingTicker.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsCandlesticks

> []OptionsCandlestick ListOptionsCandlesticks(ctx, contract, optional)

Options contract market candlestick chart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string**| Options contract name | 
**optional** | **ListOptionsCandlesticksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOptionsCandlesticksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 
**interval** | **optional.String**| Time interval between data points | [default to 5m]

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
    contract := "BTC_USDT-20210916-5000-C" // string - Options contract name
    
    result, _, err := client.OptionsApi.ListOptionsCandlesticks(ctx, contract, nil)
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

[**[]OptionsCandlestick**](OptionsCandlestick.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsUnderlyingCandlesticks

> []FuturesCandlestick ListOptionsUnderlyingCandlesticks(ctx, underlying, optional)

Underlying index price candlestick chart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Underlying (Obtained by listing underlying endpoint) | 
**optional** | **ListOptionsUnderlyingCandlesticksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOptionsUnderlyingCandlesticksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 
**interval** | **optional.String**| Time interval between data points | [default to 5m]

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
    underlying := "BTC_USDT" // string - Underlying (Obtained by listing underlying endpoint)
    
    result, _, err := client.OptionsApi.ListOptionsUnderlyingCandlesticks(ctx, underlying, nil)
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

[**[]FuturesCandlestick**](FuturesCandlestick.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsTrades

> []FuturesTrade ListOptionsTrades(ctx, optional)

Market trade records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListOptionsTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOptionsTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Options contract name | 
**type_** | **optional.String**| &#x60;C&#x60; for call, &#x60;P&#x60; for put | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 

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
    
    result, _, err := client.OptionsApi.ListOptionsTrades(ctx, nil)
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

## ListOptionsAccount

> OptionsAccount ListOptionsAccount(ctx, )

Query account information

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
    
    result, _, err := client.OptionsApi.ListOptionsAccount(ctx)
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

[**OptionsAccount**](OptionsAccount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsAccountBook

> []OptionsAccountBook ListOptionsAccountBook(ctx, optional)

Query account change history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListOptionsAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOptionsAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 
**type_** | **optional.String**| Change types: - dnw: Deposit &amp; Withdrawal - prem: Trading premium - fee: Trading fee - refr: Referrer rebate - set: Settlement P&amp;L | 

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
    
    result, _, err := client.OptionsApi.ListOptionsAccountBook(ctx, nil)
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

[**[]OptionsAccountBook**](OptionsAccountBook.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsPositions

> []OptionsPosition ListOptionsPositions(ctx, optional)

List user's positions of specified underlying

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListOptionsPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOptionsPositionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**underlying** | **optional.String**| Underlying | 

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
    
    result, _, err := client.OptionsApi.ListOptionsPositions(ctx, nil)
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

[**[]OptionsPosition**](OptionsPosition.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOptionsPosition

> OptionsPosition GetOptionsPosition(ctx, contract)

Get specified contract position

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string**|  | 

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
    contract := "BTC_USDT-20211130-65000-C" // string - 
    
    result, _, err := client.OptionsApi.GetOptionsPosition(ctx, contract)
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

[**OptionsPosition**](OptionsPosition.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsPositionClose

> []OptionsPositionClose ListOptionsPositionClose(ctx, underlying, optional)

List user's liquidation history of specified underlying

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Underlying (Obtained by listing underlying endpoint) | 
**optional** | **ListOptionsPositionCloseOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOptionsPositionCloseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Options contract name | 

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
    underlying := "BTC_USDT" // string - Underlying (Obtained by listing underlying endpoint)
    
    result, _, err := client.OptionsApi.ListOptionsPositionClose(ctx, underlying, nil)
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

[**[]OptionsPositionClose**](OptionsPositionClose.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOptionsOrders

> []OptionsOrder ListOptionsOrders(ctx, status, optional)

List options orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**status** | **string**| Query order list based on status | 
**optional** | **ListOptionsOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOptionsOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Options contract name | 
**underlying** | **optional.String**| Underlying | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 

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
    status := "open" // string - Query order list based on status
    
    result, _, err := client.OptionsApi.ListOptionsOrders(ctx, status, nil)
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

[**[]OptionsOrder**](OptionsOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateOptionsOrder

> OptionsOrder CreateOptionsOrder(ctx, optionsOrder)

Create an options order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optionsOrder** | [**OptionsOrder**](OptionsOrder.md)|  | 

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
    optionsOrder := gateapi.OptionsOrder{} // OptionsOrder - 
    
    result, _, err := client.OptionsApi.CreateOptionsOrder(ctx, optionsOrder)
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

[**OptionsOrder**](OptionsOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelOptionsOrders

> []OptionsOrder CancelOptionsOrders(ctx, optional)

Cancel all orders with 'open' status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CancelOptionsOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelOptionsOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Options contract name | 
**underlying** | **optional.String**| Underlying | 
**side** | **optional.String**| Specify all bids or all asks, both included if not specified | 

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
    
    result, _, err := client.OptionsApi.CancelOptionsOrders(ctx, nil)
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

[**[]OptionsOrder**](OptionsOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOptionsOrder

> OptionsOrder GetOptionsOrder(ctx, orderId)

Query single order details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64**| Order ID returned when order is successfully created | 

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
    orderId := 12345 // int64 - Order ID returned when order is successfully created
    
    result, _, err := client.OptionsApi.GetOptionsOrder(ctx, orderId)
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

[**OptionsOrder**](OptionsOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelOptionsOrder

> OptionsOrder CancelOptionsOrder(ctx, orderId)

Cancel single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64**| Order ID returned when order is successfully created | 

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
    orderId := 12345 // int64 - Order ID returned when order is successfully created
    
    result, _, err := client.OptionsApi.CancelOptionsOrder(ctx, orderId)
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

[**OptionsOrder**](OptionsOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CountdownCancelAllOptions

> TriggerTime CountdownCancelAllOptions(ctx, countdownCancelAllOptionsTask)

Countdown cancel orders

Option order heartbeat detection, when the `timeout` time set by the user is reached, if the existing countdown is not canceled or a new countdown is set, the related `option pending order` will be automatically canceled.  This interface can be called repeatedly to set a new countdown or cancel the countdown.  Usage example: Repeat this interface at intervals of 30 seconds, with each countdown `timeout` set to 30 (seconds).  If this interface is not called again within 30 seconds, all pending orders on the `underlying` `contract` you specified will be automatically cancelled. If `underlying` `contract` is not specified, user will be automatically cancelled  If `timeout` is set to 0 within 30 seconds, the countdown timer will expire and the automatic order cancellation function will be cancelled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countdownCancelAllOptionsTask** | [**CountdownCancelAllOptionsTask**](CountdownCancelAllOptionsTask.md)|  | 

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
    countdownCancelAllOptionsTask := gateapi.CountdownCancelAllOptionsTask{} // CountdownCancelAllOptionsTask - 
    
    result, _, err := client.OptionsApi.CountdownCancelAllOptions(ctx, countdownCancelAllOptionsTask)
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

[**TriggerTime**](triggerTime.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListMyOptionsTrades

> []OptionsMyTrade ListMyOptionsTrades(ctx, underlying, optional)

Query personal trading records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**underlying** | **string**| Underlying (Obtained by listing underlying endpoint) | 
**optional** | **ListMyOptionsTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMyOptionsTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contract** | **optional.String**| Options contract name | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 

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
    underlying := "BTC_USDT" // string - Underlying (Obtained by listing underlying endpoint)
    
    result, _, err := client.OptionsApi.ListMyOptionsTrades(ctx, underlying, nil)
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

[**[]OptionsMyTrade**](OptionsMyTrade.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOptionsMMP

> []OptionsMmp GetOptionsMMP(ctx, optional)

MMP Query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetOptionsMMPOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOptionsMMPOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**underlying** | **optional.String**| Underlying | 

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
    
    result, _, err := client.OptionsApi.GetOptionsMMP(ctx, nil)
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

[**[]OptionsMmp**](OptionsMMP.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SetOptionsMMP

> OptionsMmp SetOptionsMMP(ctx, optionsMmp)

MMP Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optionsMmp** | [**OptionsMmp**](OptionsMmp.md)|  | 

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
    optionsMmp := gateapi.OptionsMmp{} // OptionsMmp - 
    
    result, _, err := client.OptionsApi.SetOptionsMMP(ctx, optionsMmp)
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

[**OptionsMmp**](OptionsMMP.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ResetOptionsMMP

> OptionsMmp ResetOptionsMMP(ctx, optionsMmpReset)

MMP Reset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optionsMmpReset** | [**OptionsMmpReset**](OptionsMmpReset.md)|  | 

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
    optionsMmpReset := gateapi.OptionsMmpReset{} // OptionsMmpReset - 
    
    result, _, err := client.OptionsApi.ResetOptionsMMP(ctx, optionsMmpReset)
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

[**OptionsMmp**](OptionsMMP.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
