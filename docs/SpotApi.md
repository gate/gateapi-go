# SpotApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCurrencies**](SpotApi.md#ListCurrencies) | **Get** /spot/currencies | Query all currency information
[**GetCurrency**](SpotApi.md#GetCurrency) | **Get** /spot/currencies/{currency} | Query single currency information
[**ListCurrencyPairs**](SpotApi.md#ListCurrencyPairs) | **Get** /spot/currency_pairs | Query all supported currency pairs
[**GetCurrencyPair**](SpotApi.md#GetCurrencyPair) | **Get** /spot/currency_pairs/{currency_pair} | Query single currency pair details
[**ListTickers**](SpotApi.md#ListTickers) | **Get** /spot/tickers | Get currency pair ticker information
[**ListOrderBook**](SpotApi.md#ListOrderBook) | **Get** /spot/order_book | Get market depth information
[**ListTrades**](SpotApi.md#ListTrades) | **Get** /spot/trades | Query market transaction records
[**ListCandlesticks**](SpotApi.md#ListCandlesticks) | **Get** /spot/candlesticks | Market K-line chart
[**GetFee**](SpotApi.md#GetFee) | **Get** /spot/fee | Query account fee rates
[**GetBatchSpotFee**](SpotApi.md#GetBatchSpotFee) | **Get** /spot/batch_fee | Batch query account fee rates
[**ListSpotAccounts**](SpotApi.md#ListSpotAccounts) | **Get** /spot/accounts | List spot trading accounts
[**ListSpotAccountBook**](SpotApi.md#ListSpotAccountBook) | **Get** /spot/account_book | Query spot account transaction history
[**CreateBatchOrders**](SpotApi.md#CreateBatchOrders) | **Post** /spot/batch_orders | Batch place orders
[**ListAllOpenOrders**](SpotApi.md#ListAllOpenOrders) | **Get** /spot/open_orders | List all open orders
[**CreateCrossLiquidateOrder**](SpotApi.md#CreateCrossLiquidateOrder) | **Post** /spot/cross_liquidate_orders | Close position when cross-currency is disabled
[**ListOrders**](SpotApi.md#ListOrders) | **Get** /spot/orders | List orders
[**CreateOrder**](SpotApi.md#CreateOrder) | **Post** /spot/orders | Create an order
[**CancelOrders**](SpotApi.md#CancelOrders) | **Delete** /spot/orders | Cancel all &#x60;open&#x60; orders in specified currency pair
[**CancelBatchOrders**](SpotApi.md#CancelBatchOrders) | **Post** /spot/cancel_batch_orders | Cancel batch orders by specified ID list
[**GetOrder**](SpotApi.md#GetOrder) | **Get** /spot/orders/{order_id} | Query single order details
[**CancelOrder**](SpotApi.md#CancelOrder) | **Delete** /spot/orders/{order_id} | Cancel single order
[**AmendOrder**](SpotApi.md#AmendOrder) | **Patch** /spot/orders/{order_id} | Amend single order
[**ListMyTrades**](SpotApi.md#ListMyTrades) | **Get** /spot/my_trades | Query personal trading records
[**GetSystemTime**](SpotApi.md#GetSystemTime) | **Get** /spot/time | Get server current time
[**CountdownCancelAllSpot**](SpotApi.md#CountdownCancelAllSpot) | **Post** /spot/countdown_cancel_all | Countdown cancel orders
[**AmendBatchOrders**](SpotApi.md#AmendBatchOrders) | **Post** /spot/amend_batch_orders | Batch modification of orders
[**GetSpotInsuranceHistory**](SpotApi.md#GetSpotInsuranceHistory) | **Get** /spot/insurance_history | Query spot insurance fund historical data
[**ListSpotPriceTriggeredOrders**](SpotApi.md#ListSpotPriceTriggeredOrders) | **Get** /spot/price_orders | Query running auto order list
[**CreateSpotPriceTriggeredOrder**](SpotApi.md#CreateSpotPriceTriggeredOrder) | **Post** /spot/price_orders | Create price-triggered order
[**CancelSpotPriceTriggeredOrderList**](SpotApi.md#CancelSpotPriceTriggeredOrderList) | **Delete** /spot/price_orders | Cancel all auto orders
[**GetSpotPriceTriggeredOrder**](SpotApi.md#GetSpotPriceTriggeredOrder) | **Get** /spot/price_orders/{order_id} | Query single auto order details
[**CancelSpotPriceTriggeredOrder**](SpotApi.md#CancelSpotPriceTriggeredOrder) | **Delete** /spot/price_orders/{order_id} | Cancel single auto order


## ListCurrencies

> []Currency ListCurrencies(ctx, )

Query all currency information

When a currency corresponds to multiple chains, you can query the information of multiple chains through the `chains` field, such as the charging and recharge status, identification, etc. of the chain

### Required Parameters


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
    
    result, _, err := client.SpotApi.ListCurrencies(ctx)
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

[**[]Currency**](Currency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCurrency

> Currency GetCurrency(ctx, currency)

Query single currency information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Currency name | 

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
    currency := "GT" // string - Currency name
    
    result, _, err := client.SpotApi.GetCurrency(ctx, currency)
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

[**Currency**](Currency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCurrencyPairs

> []CurrencyPair ListCurrencyPairs(ctx, )

Query all supported currency pairs

### Required Parameters


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
    
    result, _, err := client.SpotApi.ListCurrencyPairs(ctx)
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

[**[]CurrencyPair**](CurrencyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCurrencyPair

> CurrencyPair GetCurrencyPair(ctx, currencyPair)

Query single currency pair details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Currency pair | 

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
    currencyPair := "ETH_BTC" // string - Currency pair
    
    result, _, err := client.SpotApi.GetCurrencyPair(ctx, currencyPair)
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

[**CurrencyPair**](CurrencyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListTickers

> []Ticker ListTickers(ctx, optional)

Get currency pair ticker information

If `currency_pair` is specified, only query that currency pair; otherwise return all information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListTickersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTickersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Currency pair | 
**timezone** | **optional.String**| Timezone | 

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
    
    result, _, err := client.SpotApi.ListTickers(ctx, nil)
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

[**[]Ticker**](Ticker.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOrderBook

> OrderBook ListOrderBook(ctx, currencyPair, optional)

Get market depth information

Market depth buy orders are sorted by price from high to low, sell orders are sorted from low to high

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Currency pair | 
**optional** | **ListOrderBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOrderBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**interval** | **optional.String**| Price precision for depth aggregation, 0 means no aggregation, defaults to 0 if not specified | [default to 0]
**limit** | **optional.Int32**| Number of depth levels | [default to 10]
**withId** | **optional.Bool**| Return order book update ID | [default to false]

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
    currencyPair := "BTC_USDT" // string - Currency pair
    
    result, _, err := client.SpotApi.ListOrderBook(ctx, currencyPair, nil)
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

[**OrderBook**](OrderBook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListTrades

> []Trade ListTrades(ctx, currencyPair, optional)

Query market transaction records

Supports querying by time range using `from` and `to` parameters or pagination based on `last_id`. By default, queries the last 30 days.  Pagination based on `last_id` is no longer recommended. If `last_id` is specified, the time range query parameters will be ignored.  When using limit&page pagination to retrieve data, the maximum number of pages is 100,000, that is, limit * (page - 1) <= 100,000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Currency pair | 
**optional** | **ListTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of items returned in list. Default: 100, minimum: 1, maximum: 1000 | [default to 100]
**lastId** | **optional.String**| Specify the currency name to query in batches, and support up to 100 pass parameters at a time | 
**reverse** | **optional.Bool**| Whether to retrieve data less than &#x60;last_id&#x60;. Default returns records greater than &#x60;last_id&#x60;.  Set to &#x60;true&#x60; to trace back market trade records, &#x60;false&#x60; to get latest trades.  No effect when &#x60;last_id&#x60; is not set. | [default to false]
**from** | **optional.Int64**| Start timestamp for the query | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**page** | **optional.Int32**| Page number | [default to 1]

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
    currencyPair := "BTC_USDT" // string - Currency pair
    
    result, _, err := client.SpotApi.ListTrades(ctx, currencyPair, nil)
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

[**[]Trade**](Trade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCandlesticks

> [][]string ListCandlesticks(ctx, currencyPair, optional)

Market K-line chart

Maximum of 1000 points can be returned in a query. Be sure not to exceed the limit when specifying from, to and interval

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Currency pair | 
**optional** | **ListCandlesticksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCandlesticksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of recent data points to return. &#x60;limit&#x60; conflicts with &#x60;from&#x60; and &#x60;to&#x60;. If either &#x60;from&#x60; or &#x60;to&#x60; is specified, request will be rejected. | [default to 100]
**from** | **optional.Int64**| Start time of candlesticks, formatted in Unix timestamp in seconds. Default to&#x60;to - 100 * interval&#x60; if not specified | 
**to** | **optional.Int64**| Specify the end time of the K-line chart, defaults to current time if not specified, note that the time format is Unix timestamp with second precision | 
**interval** | **optional.String**| Time interval between data points. Note that &#x60;30d&#x60; represents a calendar month, not aligned to 30 days | [default to 30m]

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
    currencyPair := "BTC_USDT" // string - Currency pair
    
    result, _, err := client.SpotApi.ListCandlesticks(ctx, currencyPair, nil)
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

[**[][]string**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetFee

> SpotFee GetFee(ctx, optional)

Query account fee rates

This API is deprecated. The new fee query API is `/wallet/fee`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetFeeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFeeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Specify currency pair to get more accurate fee settings.  This field is optional. Usually fee settings are the same for all currency pairs. | 

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
    
    result, _, err := client.SpotApi.GetFee(ctx, nil)
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

[**SpotFee**](SpotFee.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetBatchSpotFee

> map[string]SpotFee GetBatchSpotFee(ctx, currencyPairs)

Batch query account fee rates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPairs** | **string**| Maximum 50 currency pairs per request | 

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
    currencyPairs := "BTC_USDT,ETH_USDT" // string - Maximum 50 currency pairs per request
    
    result, _, err := client.SpotApi.GetBatchSpotFee(ctx, currencyPairs)
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

[**map[string]SpotFee**](SpotFee.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSpotAccounts

> []SpotAccount ListSpotAccounts(ctx, optional)

List spot trading accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSpotAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSpotAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Query by specified currency name | 

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
    
    result, _, err := client.SpotApi.ListSpotAccounts(ctx, nil)
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

[**[]SpotAccount**](SpotAccount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSpotAccountBook

> []SpotAccountBook ListSpotAccountBook(ctx, optional)

Query spot account transaction history

Record query time range cannot exceed 30 days.  When using limit&page pagination to retrieve data, the maximum number of pages is 100,000, that is, limit * (page - 1) <= 100,000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSpotAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSpotAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Query by specified currency name | 
**from** | **optional.Int64**| Start timestamp for the query | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**type_** | **optional.String**| Query by specified account change type. If not specified, all change types will be included. | 
**code** | **optional.String**| Specify account change code for query. If not specified, all change types are included. This parameter has higher priority than &#x60;type&#x60; | 

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
    
    result, _, err := client.SpotApi.ListSpotAccountBook(ctx, nil)
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

[**[]SpotAccountBook**](SpotAccountBook.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateBatchOrders

> []BatchOrder CreateBatchOrders(ctx, order, optional)

Batch place orders

Batch order requirements:  1. Custom order field `text` must be specified 2. Up to 4 currency pairs per request, with up to 10 orders per currency pair 3. Spot orders and margin orders cannot be mixed; all `account` fields in the same request must be identical

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**order** | [**[]Order**](Order.md)|  | 
**optional** | **CreateBatchOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBatchOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateExptime** | **optional.String**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    order := []gateapi.Order{gateapi.Order{}} // []Order - 
    
    result, _, err := client.SpotApi.CreateBatchOrders(ctx, order, nil)
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

[**[]BatchOrder**](BatchOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAllOpenOrders

> []OpenOrders ListAllOpenOrders(ctx, optional)

List all open orders

Query the current order list of all trading pairs. Please note that the paging parameter controls the number of pending orders in each trading pair. There is no paging control trading pairs. All trading pairs with pending orders will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListAllOpenOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllOpenOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in one page in each currency pair | [default to 100]
**account** | **optional.String**| Specify query account | 

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
    
    result, _, err := client.SpotApi.ListAllOpenOrders(ctx, nil)
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

[**[]OpenOrders**](OpenOrders.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateCrossLiquidateOrder

> Order CreateCrossLiquidateOrder(ctx, liquidateOrder)

Close position when cross-currency is disabled

Currently, only cross-margin accounts are supported to place buy orders for disabled currencies. Maximum buy quantity = (unpaid principal and interest - currency balance - the amount of the currency in pending orders) / 0.998

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liquidateOrder** | [**LiquidateOrder**](LiquidateOrder.md)|  | 

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
    liquidateOrder := gateapi.LiquidateOrder{} // LiquidateOrder - 
    
    result, _, err := client.SpotApi.CreateCrossLiquidateOrder(ctx, liquidateOrder)
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

[**Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOrders

> []Order ListOrders(ctx, currencyPair, status, optional)

List orders

Note that query results default to spot order lists for spot, unified account, and isolated margin accounts.  When `status` is set to `open` (i.e., when querying pending order lists), only `page` and `limit` pagination controls are supported. `limit` can only be set to a maximum of 100. The `side` parameter and time range query parameters `from` and `to` are not supported.  When `status` is set to `finished` (i.e., when querying historical orders), in addition to pagination queries, `from` and `to` time range queries are also supported. Additionally, the `side` parameter can be set to filter one-sided history.  Time range filter parameters are processed according to the order end time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Query by specified currency pair. Required for open orders, optional for filled orders | 
**status** | **string**| List orders based on status  &#x60;open&#x60; - order is waiting to be filled &#x60;finished&#x60; - order has been filled or cancelled  | 
**optional** | **ListOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records to be returned. If &#x60;status&#x60; is &#x60;open&#x60;, maximum of &#x60;limit&#x60; is 100 | [default to 100]
**account** | **optional.String**| Specify query account | 
**from** | **optional.Int64**| Start timestamp for the query | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
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
    currencyPair := "BTC_USDT" // string - Query by specified currency pair. Required for open orders, optional for filled orders
    status := "open" // string - List orders based on status  `open` - order is waiting to be filled `finished` - order has been filled or cancelled 
    
    result, _, err := client.SpotApi.ListOrders(ctx, currencyPair, status, nil)
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

[**[]Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateOrder

> Order CreateOrder(ctx, order, optional)

Create an order

Supports spot, margin, leverage, and cross-margin leverage orders. Use different accounts through the `account` field. Default is `spot`, which means using the spot account to place orders. If the user has a `unified` account, the default is to place orders with the unified account.  When using leveraged account trading (i.e., when `account` is set to `margin`), you can set `auto_borrow` to `true`. In case of insufficient account balance, the system will automatically execute `POST /margin/uni/loans` to borrow the insufficient amount. Whether assets obtained after leveraged order execution are automatically used to repay borrowing orders of the isolated margin account depends on the automatic repayment settings of the user's isolated margin account. Account automatic repayment settings can be queried and set through `/margin/auto_repay`.  When using unified account trading (i.e., when `account` is set to `unified`), `auto_borrow` can also be enabled to realize automatic borrowing of insufficient amounts. However, unlike the isolated margin account, whether unified account orders are automatically repaid depends on the `auto_repay` setting when placing the order. This setting only applies to the current order, meaning only assets obtained after order execution will be used to repay borrowing orders of the cross-margin account. Unified account ordering currently supports enabling both `auto_borrow` and `auto_repay` simultaneously.  Auto repayment will be triggered when the order ends, i.e., when `status` is `cancelled` or `closed`.  **Order Status**  The order status in pending orders is `open`, which remains `open` until all quantity is filled. If fully filled, the order ends and status becomes `closed`. If the order is cancelled before all transactions are completed, regardless of partial fills, the status will become `cancelled`.  **Iceberg Orders**  `iceberg` is used to set the displayed quantity of iceberg orders and does not support complete hiding. Note that hidden portions are charged according to the taker's fee rate.  **Self-Trade Prevention**  Set `stp_act` to determine the self-trade prevention strategy to use

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**order** | [**Order**](Order.md)|  | 
**optional** | **CreateOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateExptime** | **optional.String**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    order := gateapi.Order{} // Order - 
    
    result, _, err := client.SpotApi.CreateOrder(ctx, order, nil)
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

[**Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelOrders

> []OrderCancel CancelOrders(ctx, optional)

Cancel all `open` orders in specified currency pair

When the `account` parameter is not specified, all pending orders including spot, unified account, and isolated margin will be cancelled. When `currency_pair` is not specified, all trading pair pending orders will be cancelled. You can specify a particular account to cancel all pending orders under that account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CancelOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Currency pair | 
**side** | **optional.String**| Specify all bids or all asks, both included if not specified | 
**account** | **optional.String**| Specify account type  Classic account: All are included if not specified Unified account: Specify &#x60;unified&#x60; | 
**actionMode** | **optional.String**| Processing Mode  When placing an order, different fields are returned based on the action_mode  - &#x60;ACK&#x60;: Asynchronous mode, returns only key order fields - &#x60;RESULT&#x60;: No clearing information - &#x60;FULL&#x60;: Full mode (default) | 
**xGateExptime** | **optional.String**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    
    result, _, err := client.SpotApi.CancelOrders(ctx, nil)
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

[**[]OrderCancel**](OrderCancel.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelBatchOrders

> []CancelOrderResult CancelBatchOrders(ctx, cancelBatchOrder, optional)

Cancel batch orders by specified ID list

Multiple currency pairs can be specified, but maximum 20 orders are allowed per request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cancelBatchOrder** | [**[]CancelBatchOrder**](CancelBatchOrder.md)|  | 
**optional** | **CancelBatchOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelBatchOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateExptime** | **optional.String**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    cancelBatchOrder := []gateapi.CancelBatchOrder{gateapi.CancelBatchOrder{}} // []CancelBatchOrder - 
    
    result, _, err := client.SpotApi.CancelBatchOrders(ctx, cancelBatchOrder, nil)
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

[**[]CancelOrderResult**](CancelOrderResult.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOrder

> Order GetOrder(ctx, orderId, currencyPair, optional)

Query single order details

By default, queries orders for spot, unified account, and isolated margin accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| The order ID returned when the order was successfully created or the custom ID specified by the user&#39;s creation (i.e. the &#x60;text&#x60; field). Operations based on custom IDs can only be checked in pending orders. Only order ID can be used after the order is finished (transaction/cancel) | 
**currencyPair** | **string**| Specify the trading pair to query. This field is required when querying pending order records. This field can be omitted when querying filled order records. | 
**optional** | **GetOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**account** | **optional.String**| Specify query account | 

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
    orderId := "12345" // string - The order ID returned when the order was successfully created or the custom ID specified by the user's creation (i.e. the `text` field). Operations based on custom IDs can only be checked in pending orders. Only order ID can be used after the order is finished (transaction/cancel)
    currencyPair := "BTC_USDT" // string - Specify the trading pair to query. This field is required when querying pending order records. This field can be omitted when querying filled order records.
    
    result, _, err := client.SpotApi.GetOrder(ctx, orderId, currencyPair, nil)
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

[**Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelOrder

> Order CancelOrder(ctx, orderId, currencyPair, optional)

Cancel single order

By default, orders for spot, unified accounts and leveraged accounts are revoked.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| The order ID returned when the order was successfully created or the custom ID specified by the user&#39;s creation (i.e. the &#x60;text&#x60; field). Operations based on custom IDs can only be checked in pending orders. Only order ID can be used after the order is finished (transaction/cancel) | 
**currencyPair** | **string**| Currency pair | 
**optional** | **CancelOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**account** | **optional.String**| Specify query account | 
**actionMode** | **optional.String**| Processing Mode  When placing an order, different fields are returned based on the action_mode  - &#x60;ACK&#x60;: Asynchronous mode, returns only key order fields - &#x60;RESULT&#x60;: No clearing information - &#x60;FULL&#x60;: Full mode (default) | 
**xGateExptime** | **optional.String**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    orderId := "12345" // string - The order ID returned when the order was successfully created or the custom ID specified by the user's creation (i.e. the `text` field). Operations based on custom IDs can only be checked in pending orders. Only order ID can be used after the order is finished (transaction/cancel)
    currencyPair := "BTC_USDT" // string - Currency pair
    
    result, _, err := client.SpotApi.CancelOrder(ctx, orderId, currencyPair, nil)
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

[**Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AmendOrder

> Order AmendOrder(ctx, orderId, orderPatch, optional)

Amend single order

Modify orders in spot, unified account and isolated margin account by default.  Currently both request body and query support currency_pair and account parameters, but request body has higher priority.  currency_pair must be filled in one of the request body or query parameters.  About rate limit: Order modification and order creation share the same rate limit rules.  About matching priority: Only reducing the quantity does not affect the matching priority. Modifying the price or increasing the quantity will adjust the priority to the end of the new price level.  Note: Modifying the quantity to be less than the filled quantity will trigger a cancellation and isolated margin account by default.  Currently both request body and query support currency_pair and account parameters, but request body has higher priority.  currency_pair must be filled in one of the request body or query parameters.  About rate limit: Order modification and order creation share the same rate limit rules.  About matching priority: Only reducing the quantity does not affect the matching priority. Modifying the price or increasing the quantity will adjust the priority to the end of the new price level.  Note: Modifying the quantity to be less than the filled quantity will trigger a cancellation operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| The order ID returned when the order was successfully created or the custom ID specified by the user&#39;s creation (i.e. the &#x60;text&#x60; field). Operations based on custom IDs can only be checked in pending orders. Only order ID can be used after the order is finished (transaction/cancel) | 
**orderPatch** | [**OrderPatch**](OrderPatch.md)|  | 
**optional** | **AmendOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AmendOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Currency pair | 
**account** | **optional.String**| Specify query account | 
**xGateExptime** | **optional.String**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    orderId := "12345" // string - The order ID returned when the order was successfully created or the custom ID specified by the user's creation (i.e. the `text` field). Operations based on custom IDs can only be checked in pending orders. Only order ID can be used after the order is finished (transaction/cancel)
    orderPatch := gateapi.OrderPatch{} // OrderPatch - 
    
    result, _, err := client.SpotApi.AmendOrder(ctx, orderId, orderPatch, nil)
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

[**Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListMyTrades

> []Trade ListMyTrades(ctx, optional)

Query personal trading records

By default query of transaction records for spot, unified account and warehouse-by-site leverage accounts.  The history within a specified time range can be queried by specifying `from` or (and) `to`.  - If no time parameters are specified, only data for the last 7 days can be obtained. - If only any parameter of `from` or `to` is specified, only 7-day data from the start (or end) of the specified time is returned. - The range not allowed to exceed 30 days.  The parameters of the time range filter are processed according to the order end time.  The maximum number of pages when searching data using limit&page paging function is 100,0, that is, limit * (page - 1) <= 100,0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListMyTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMyTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Retrieve results with specified currency pair | 
**limit** | **optional.Int32**| Maximum number of items returned in list. Default: 100, minimum: 1, maximum: 1000 | [default to 100]
**page** | **optional.Int32**| Page number | [default to 1]
**orderId** | **optional.String**| Filter trades with specified order ID. &#x60;currency_pair&#x60; is also required if this field is present | 
**account** | **optional.String**| Specify query account | 
**from** | **optional.Int64**| Start timestamp for the query | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 

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
    
    result, _, err := client.SpotApi.ListMyTrades(ctx, nil)
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

[**[]Trade**](Trade.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSystemTime

> SystemTime GetSystemTime(ctx, )

Get server current time

### Required Parameters


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
    
    result, _, err := client.SpotApi.GetSystemTime(ctx)
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

[**SystemTime**](SystemTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CountdownCancelAllSpot

> TriggerTime CountdownCancelAllSpot(ctx, countdownCancelAllSpotTask)

Countdown cancel orders

Spot order heartbeat detection. If there is no \"cancel existing countdown\" or \"set new countdown\" when the user-set `timeout` time is reached, the related `spot pending orders` will be automatically cancelled. This interface can be called repeatedly to set a new countdown or cancel the countdown. Usage example: Repeat this interface at 30s intervals, setting the countdown `timeout` to `30 (seconds)` each time. If this interface is not called again within 30 seconds, all pending orders on the `market` you specified will be automatically cancelled. If no `market` is specified, all market cancelled. If the `timeout` is set to 0 within 30 seconds, the countdown timer will be terminated and the automatic order cancellation function will be cancelled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countdownCancelAllSpotTask** | [**CountdownCancelAllSpotTask**](CountdownCancelAllSpotTask.md)|  | 

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
    countdownCancelAllSpotTask := gateapi.CountdownCancelAllSpotTask{} // CountdownCancelAllSpotTask - 
    
    result, _, err := client.SpotApi.CountdownCancelAllSpot(ctx, countdownCancelAllSpotTask)
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

## AmendBatchOrders

> []BatchOrder AmendBatchOrders(ctx, batchAmendItem, optional)

Batch modification of orders

Modify orders in spot, unified account and isolated margin account by default. Modify uncompleted orders, up to 5 orders can be modified at a time. Request parameters should be passed in array format. If there are order modification failures during the batch modification process, the modification of the next order will continue to be executed, and the execution will return with the corresponding order failure information. The call order of batch modification orders is consistent with the order list order. The return content order of batch modification orders is consistent with the order list order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchAmendItem** | [**[]BatchAmendItem**](BatchAmendItem.md)|  | 
**optional** | **AmendBatchOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AmendBatchOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateExptime** | **optional.String**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    batchAmendItem := []gateapi.BatchAmendItem{gateapi.BatchAmendItem{}} // []BatchAmendItem - 
    
    result, _, err := client.SpotApi.AmendBatchOrders(ctx, batchAmendItem, nil)
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

[**[]BatchOrder**](BatchOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSpotInsuranceHistory

> []SpotInsuranceHistory GetSpotInsuranceHistory(ctx, business, currency, from, to, optional)

Query spot insurance fund historical data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**business** | **string**| Leverage business, margin - position by position; unified - unified account | 
**currency** | **string**| Currency | 
**from** | **int64**| Start timestamp in seconds | 
**to** | **int64**| End timestamp in seconds | 
**optional** | **GetSpotInsuranceHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSpotInsuranceHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| The maximum number of items returned in the list, the default value is 30 | [default to 30]

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
    business := "margin" // string - Leverage business, margin - position by position; unified - unified account
    currency := "BTC" // string - Currency
    from := 1547706332 // int64 - Start timestamp in seconds
    to := 1547706332 // int64 - End timestamp in seconds
    
    result, _, err := client.SpotApi.GetSpotInsuranceHistory(ctx, business, currency, from, to, nil)
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

[**[]SpotInsuranceHistory**](SpotInsuranceHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSpotPriceTriggeredOrders

> []SpotPriceTriggeredOrder ListSpotPriceTriggeredOrders(ctx, status, optional)

Query running auto order list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**status** | **string**| Query order list based on status | 
**optional** | **ListSpotPriceTriggeredOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSpotPriceTriggeredOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**market** | **optional.String**| Trading market | 
**account** | **optional.String**| Trading account type. Unified account must be set to &#x60;unified&#x60; | 
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
    status := "status_example" // string - Query order list based on status
    
    result, _, err := client.SpotApi.ListSpotPriceTriggeredOrders(ctx, status, nil)
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

[**[]SpotPriceTriggeredOrder**](SpotPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateSpotPriceTriggeredOrder

> TriggerOrderResponse CreateSpotPriceTriggeredOrder(ctx, spotPriceTriggeredOrder)

Create price-triggered order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spotPriceTriggeredOrder** | [**SpotPriceTriggeredOrder**](SpotPriceTriggeredOrder.md)|  | 

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
    spotPriceTriggeredOrder := gateapi.SpotPriceTriggeredOrder{} // SpotPriceTriggeredOrder - 
    
    result, _, err := client.SpotApi.CreateSpotPriceTriggeredOrder(ctx, spotPriceTriggeredOrder)
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

## CancelSpotPriceTriggeredOrderList

> []SpotPriceTriggeredOrder CancelSpotPriceTriggeredOrderList(ctx, optional)

Cancel all auto orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CancelSpotPriceTriggeredOrderListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelSpotPriceTriggeredOrderListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**market** | **optional.String**| Trading market | 
**account** | **optional.String**| Trading account type. Unified account must be set to &#x60;unified&#x60; | 

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
    
    result, _, err := client.SpotApi.CancelSpotPriceTriggeredOrderList(ctx, nil)
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

[**[]SpotPriceTriggeredOrder**](SpotPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSpotPriceTriggeredOrder

> SpotPriceTriggeredOrder GetSpotPriceTriggeredOrder(ctx, orderId)

Query single auto order details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
    orderId := "orderId_example" // string - ID returned when order is successfully created
    
    result, _, err := client.SpotApi.GetSpotPriceTriggeredOrder(ctx, orderId)
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

[**SpotPriceTriggeredOrder**](SpotPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelSpotPriceTriggeredOrder

> SpotPriceTriggeredOrder CancelSpotPriceTriggeredOrder(ctx, orderId)

Cancel single auto order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
    orderId := "orderId_example" // string - ID returned when order is successfully created
    
    result, _, err := client.SpotApi.CancelSpotPriceTriggeredOrder(ctx, orderId)
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

[**SpotPriceTriggeredOrder**](SpotPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
