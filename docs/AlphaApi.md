# AlphaApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAlphaAccounts**](AlphaApi.md#ListAlphaAccounts) | **Get** /alpha/accounts | Alpha Account API
[**ListAlphaAccountBook**](AlphaApi.md#ListAlphaAccountBook) | **Get** /alpha/account_book | Alpha Account Transaction History API
[**QuoteAlphaOrder**](AlphaApi.md#QuoteAlphaOrder) | **Post** /alpha/quote | Alpha Quote API
[**ListAlphaOrder**](AlphaApi.md#ListAlphaOrder) | **Get** /alpha/orders | Alpha Order List API
[**PlaceAlphaOrder**](AlphaApi.md#PlaceAlphaOrder) | **Post** /alpha/orders | Alpha Order API
[**GetAlphaOrder**](AlphaApi.md#GetAlphaOrder) | **Get** /alpha/order | Alpha Single Order Query API
[**ListAlphaCurrencies**](AlphaApi.md#ListAlphaCurrencies) | **Get** /alpha/currencies | Query currency information
[**ListAlphaTickers**](AlphaApi.md#ListAlphaTickers) | **Get** /alpha/tickers | Query currency ticker
[**ListAlphaTokens**](AlphaApi.md#ListAlphaTokens) | **Get** /alpha/tokens | Query Token Information


## ListAlphaAccounts

> []AccountsResponse ListAlphaAccounts(ctx, )

Alpha Account API

Query position assets

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
    
    result, _, err := client.AlphaApi.ListAlphaAccounts(ctx)
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

[**[]AccountsResponse**](AccountsResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAlphaAccountBook

> []AccountBookResponse ListAlphaAccountBook(ctx, from, optional)

Alpha Account Transaction History API

Query asset transactions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**from** | **int64**| Start timestamp for the query | 
**optional** | **ListAlphaAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAlphaAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**page** | **optional.Int32**| Page number | 
**limit** | **optional.Int32**| Maximum 100 items per page | 

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
    from := 56 // int64 - Start timestamp for the query
    
    result, _, err := client.AlphaApi.ListAlphaAccountBook(ctx, from, nil)
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

[**[]AccountBookResponse**](AccountBookResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QuoteAlphaOrder

> QuoteResponse QuoteAlphaOrder(ctx, quoteRequest)

Alpha Quote API

The quote_id returned by the price inquiry interface is valid for one minute; an order must be placed within this minute. If it times out, a new price inquiry is required. Rate-limited at 10 requests per second per user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteRequest** | [**QuoteRequest**](QuoteRequest.md)|  | 

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
    quoteRequest := gateapi.QuoteRequest{} // QuoteRequest - 
    
    result, _, err := client.AlphaApi.QuoteAlphaOrder(ctx, quoteRequest)
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

[**QuoteResponse**](QuoteResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAlphaOrder

> []OrderResponse ListAlphaOrder(ctx, optional)

Alpha Order List API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListAlphaOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAlphaOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Trading symbol | 
**side** | **optional.String**| Buy or sell orders - buy - sell | 
**status** | **optional.Int32**| Order Status - &#x60;0&#x60; : All - &#x60;1&#x60; : Processing - &#x60;2&#x60; : Successful - &#x60;3&#x60; : Failed - &#x60;4&#x60; : Cancelled - &#x60;5&#x60; : Buy order placed but transfer not completed - &#x60;6&#x60; : Order cancelled but transfer not completed | 
**from** | **optional.Int64**| Start time for order query | 
**to** | **optional.Int64**| End time for order query, defaults to current time if not specified | 
**limit** | **optional.Int32**| Maximum number of items returned. Default: 100, minimum: 1, maximum: 100 | [default to 100]
**page** | **optional.Int32**| Page number | [default to 1]

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
    
    result, _, err := client.AlphaApi.ListAlphaOrder(ctx, nil)
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

[**[]OrderResponse**](OrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PlaceAlphaOrder

> PlaceOrderResponse PlaceAlphaOrder(ctx, placeOrderRequest)

Alpha Order API

Order placement interface, rate-limited at 5 requests per second per user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeOrderRequest** | [**PlaceOrderRequest**](PlaceOrderRequest.md)|  | 

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
    placeOrderRequest := gateapi.PlaceOrderRequest{} // PlaceOrderRequest - 
    
    result, _, err := client.AlphaApi.PlaceAlphaOrder(ctx, placeOrderRequest)
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

[**PlaceOrderResponse**](PlaceOrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetAlphaOrder

> OrderResponse GetAlphaOrder(ctx, orderId)

Alpha Single Order Query API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Order ID | 

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
    orderId := "fdaf12321" // string - Order ID
    
    result, _, err := client.AlphaApi.GetAlphaOrder(ctx, orderId)
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

[**OrderResponse**](OrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAlphaCurrencies

> []AlphaCurrency ListAlphaCurrencies(ctx, optional)

Query currency information

When currency is provided, query and return specified currency information; when currency is not provided, return paginated currency list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListAlphaCurrenciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAlphaCurrenciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Query currency information by currency symbol | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**page** | **optional.Int32**| Page number | [default to 1]

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
    
    result, _, err := client.AlphaApi.ListAlphaCurrencies(ctx, nil)
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

[**[]AlphaCurrency**](AlphaCurrency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAlphaTickers

> []AlphaTicker ListAlphaTickers(ctx, optional)

Query currency ticker

When currency is provided, returns ticker information for the specified currency. When currency is not provided, returns paginated ticker list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListAlphaTickersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAlphaTickersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Query by specified currency name | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**page** | **optional.Int32**| Page number | [default to 1]

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
    
    result, _, err := client.AlphaApi.ListAlphaTickers(ctx, nil)
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

[**[]AlphaTicker**](AlphaTicker.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAlphaTokens

> []Tokens ListAlphaTokens(ctx, optional)

Query Token Information

Supports passing chain, platform, and address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListAlphaTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAlphaTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**chain** | **optional.String**| Query Token List by Chain  - solana - eth - bsc - base - world - sui - arbitrum - avalanche - polygon - linea - optimism - zksync - gatelayer | 
**launchPlatform** | **optional.String**| Query Token List by Launch Platform  - meteora_dbc - fourmeme - moonshot - pump - raydium_launchlab - letsbonk - gatefun - virtuals | 
**address** | **optional.String**| Query Token List by Contract Address | 
**page** | **optional.Int32**| Page number | [default to 1]

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
    
    result, _, err := client.AlphaApi.ListAlphaTokens(ctx, nil)
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

[**[]Tokens**](Tokens.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
