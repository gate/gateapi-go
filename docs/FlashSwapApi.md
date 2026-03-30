# FlashSwapApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFlashSwapCurrencyPair**](FlashSwapApi.md#ListFlashSwapCurrencyPair) | **Get** /flash_swap/currency_pairs | List All Supported Currency Pairs In Flash Swap
[**ListFlashSwapOrders**](FlashSwapApi.md#ListFlashSwapOrders) | **Get** /flash_swap/orders | Query flash swap order list
[**CreateFlashSwapOrder**](FlashSwapApi.md#CreateFlashSwapOrder) | **Post** /flash_swap/orders | Create a flash swap order
[**GetFlashSwapOrder**](FlashSwapApi.md#GetFlashSwapOrder) | **Get** /flash_swap/orders/{order_id} | Query single flash swap order
[**PreviewFlashSwapOrder**](FlashSwapApi.md#PreviewFlashSwapOrder) | **Post** /flash_swap/orders/preview | Flash swap order preview
[**CreateFlashSwapMultiCurrencyManyToOneOrder**](FlashSwapApi.md#CreateFlashSwapMultiCurrencyManyToOneOrder) | **Post** /flash_swap/multi-currency/many-to-one/order/create | Flash Swap - Multi-currency exchange - Place order (many-to-one)
[**PreviewFlashSwapMultiCurrencyManyToOneOrder**](FlashSwapApi.md#PreviewFlashSwapMultiCurrencyManyToOneOrder) | **Post** /flash_swap/multi-currency/many-to-one/order/preview | Flash Swap - Multi-currency exchange - Preview (many-to-one)
[**CreateFlashSwapOrderV1**](FlashSwapApi.md#CreateFlashSwapOrderV1) | **Post** /flash_swap/order/create | Flash Swap - Place order (one-to-one)
[**CreateFlashSwapMultiCurrencyOneToManyOrder**](FlashSwapApi.md#CreateFlashSwapMultiCurrencyOneToManyOrder) | **Post** /flash_swap/multi-currency/one-to-many/order/create | Flash Swap - Multi-currency exchange - Place order (one-to-many)
[**PreviewFlashSwapMultiCurrencyOneToManyOrder**](FlashSwapApi.md#PreviewFlashSwapMultiCurrencyOneToManyOrder) | **Post** /flash_swap/multi-currency/one-to-many/order/preview | Flash Swap - Multi-currency exchange - Preview (one-to-many)
[**PreviewFlashSwapOrderV1**](FlashSwapApi.md#PreviewFlashSwapOrderV1) | **Get** /flash_swap/order/preview | Flash Swap - Preview (one-to-one)


## ListFlashSwapCurrencyPair

> []FlashSwapCurrencyPair ListFlashSwapCurrencyPair(ctx, optional)

List All Supported Currency Pairs In Flash Swap

`BTC_GT` represents selling BTC and buying GT. The limits for each currency may vary across different currency pairs.  It is not necessary that two currencies that can be swapped instantaneously can be exchanged with each other. For example, it is possible to sell BTC and buy GT, but it does not necessarily mean that GT can be sold to buy BTC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListFlashSwapCurrencyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFlashSwapCurrencyPairOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Query by specified currency name | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of items returned. Default: 1000, minimum: 1, maximum: 1000 | [default to 1000]

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
    
    result, _, err := client.FlashSwapApi.ListFlashSwapCurrencyPair(ctx, nil)
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

[**[]FlashSwapCurrencyPair**](FlashSwapCurrencyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFlashSwapOrders

> []FlashSwapOrder ListFlashSwapOrders(ctx, optional)

Query flash swap order list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListFlashSwapOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFlashSwapOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**status** | **optional.Int32**| Flash swap order status  &#x60;1&#x60; - success &#x60;2&#x60; - failed | 
**sellCurrency** | **optional.String**| Asset name to sell - Retrieved from API &#x60;GET /flash_swap/currencies&#x60; for supported flash swap currencies | 
**buyCurrency** | **optional.String**| Asset name to buy - Retrieved from API &#x60;GET /flash_swap/currencies&#x60; for supported flash swap currencies | 
**reverse** | **optional.Bool**| Sort by ID in ascending or descending order, default &#x60;true&#x60; - &#x60;true&#x60;: ID descending order (most recent data first) - &#x60;false&#x60;: ID ascending order (oldest data first) | 
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
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    
    result, _, err := client.FlashSwapApi.ListFlashSwapOrders(ctx, nil)
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

[**[]FlashSwapOrder**](FlashSwapOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateFlashSwapOrder

> FlashSwapOrder CreateFlashSwapOrder(ctx, flashSwapOrderRequest)

Create a flash swap order

Initiate a flash swap preview in advance because order creation requires a preview result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flashSwapOrderRequest** | [**FlashSwapOrderRequest**](FlashSwapOrderRequest.md)|  | 

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
    flashSwapOrderRequest := gateapi.FlashSwapOrderRequest{} // FlashSwapOrderRequest - 
    
    result, _, err := client.FlashSwapApi.CreateFlashSwapOrder(ctx, flashSwapOrderRequest)
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

[**FlashSwapOrder**](FlashSwapOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetFlashSwapOrder

> FlashSwapOrder GetFlashSwapOrder(ctx, orderId)

Query single flash swap order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32**| Flash swap order ID | 

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
    orderId := 1 // int32 - Flash swap order ID
    
    result, _, err := client.FlashSwapApi.GetFlashSwapOrder(ctx, orderId)
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

[**FlashSwapOrder**](FlashSwapOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PreviewFlashSwapOrder

> FlashSwapOrderPreview PreviewFlashSwapOrder(ctx, flashSwapPreviewRequest)

Flash swap order preview

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flashSwapPreviewRequest** | [**FlashSwapPreviewRequest**](FlashSwapPreviewRequest.md)|  | 

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
    flashSwapPreviewRequest := gateapi.FlashSwapPreviewRequest{} // FlashSwapPreviewRequest - 
    
    result, _, err := client.FlashSwapApi.PreviewFlashSwapOrder(ctx, flashSwapPreviewRequest)
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

[**FlashSwapOrderPreview**](FlashSwapOrderPreview.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateFlashSwapMultiCurrencyManyToOneOrder

> FlashSwapMultiCurrencyManyToOneOrderCreateResp CreateFlashSwapMultiCurrencyManyToOneOrder(ctx, flashSwapMultiCurrencyManyToOneOrderCreateReq)

Flash Swap - Multi-currency exchange - Place order (many-to-one)

Create a multi-currency to single target currency exchange order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flashSwapMultiCurrencyManyToOneOrderCreateReq** | [**FlashSwapMultiCurrencyManyToOneOrderCreateReq**](FlashSwapMultiCurrencyManyToOneOrderCreateReq.md)|  | 

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
    flashSwapMultiCurrencyManyToOneOrderCreateReq := gateapi.FlashSwapMultiCurrencyManyToOneOrderCreateReq{} // FlashSwapMultiCurrencyManyToOneOrderCreateReq - 
    
    result, _, err := client.FlashSwapApi.CreateFlashSwapMultiCurrencyManyToOneOrder(ctx, flashSwapMultiCurrencyManyToOneOrderCreateReq)
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

[**FlashSwapMultiCurrencyManyToOneOrderCreateResp**](FlashSwapMultiCurrencyManyToOneOrderCreateResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PreviewFlashSwapMultiCurrencyManyToOneOrder

> FlashSwapMultiCurrencyManyToOneOrderPreviewResp PreviewFlashSwapMultiCurrencyManyToOneOrder(ctx, flashSwapMultiCurrencyManyToOneOrderPreviewReq)

Flash Swap - Multi-currency exchange - Preview (many-to-one)

Preview quote for multi-currency to single target currency exchange

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flashSwapMultiCurrencyManyToOneOrderPreviewReq** | [**FlashSwapMultiCurrencyManyToOneOrderPreviewReq**](FlashSwapMultiCurrencyManyToOneOrderPreviewReq.md)|  | 

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
    flashSwapMultiCurrencyManyToOneOrderPreviewReq := gateapi.FlashSwapMultiCurrencyManyToOneOrderPreviewReq{} // FlashSwapMultiCurrencyManyToOneOrderPreviewReq - 
    
    result, _, err := client.FlashSwapApi.PreviewFlashSwapMultiCurrencyManyToOneOrder(ctx, flashSwapMultiCurrencyManyToOneOrderPreviewReq)
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

[**FlashSwapMultiCurrencyManyToOneOrderPreviewResp**](FlashSwapMultiCurrencyManyToOneOrderPreviewResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateFlashSwapOrderV1

> FlashSwapOrderCreateResp CreateFlashSwapOrderV1(ctx, flashSwapOrderCreateReq)

Flash Swap - Place order (one-to-one)

Submit a one-to-one flash swap order. A quote_id must be obtained from the preview endpoint first

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flashSwapOrderCreateReq** | [**FlashSwapOrderCreateReq**](FlashSwapOrderCreateReq.md)|  | 

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
    flashSwapOrderCreateReq := gateapi.FlashSwapOrderCreateReq{} // FlashSwapOrderCreateReq - 
    
    result, _, err := client.FlashSwapApi.CreateFlashSwapOrderV1(ctx, flashSwapOrderCreateReq)
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

[**FlashSwapOrderCreateResp**](FlashSwapOrderCreateResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateFlashSwapMultiCurrencyOneToManyOrder

> FlashSwapMultiCurrencyOneToManyOrderCreateResp CreateFlashSwapMultiCurrencyOneToManyOrder(ctx, flashSwapMultiCurrencyOneToManyOrderCreateReq)

Flash Swap - Multi-currency exchange - Place order (one-to-many)

Create a single currency to multiple target currencies exchange order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flashSwapMultiCurrencyOneToManyOrderCreateReq** | [**FlashSwapMultiCurrencyOneToManyOrderCreateReq**](FlashSwapMultiCurrencyOneToManyOrderCreateReq.md)|  | 

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
    flashSwapMultiCurrencyOneToManyOrderCreateReq := gateapi.FlashSwapMultiCurrencyOneToManyOrderCreateReq{} // FlashSwapMultiCurrencyOneToManyOrderCreateReq - 
    
    result, _, err := client.FlashSwapApi.CreateFlashSwapMultiCurrencyOneToManyOrder(ctx, flashSwapMultiCurrencyOneToManyOrderCreateReq)
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

[**FlashSwapMultiCurrencyOneToManyOrderCreateResp**](FlashSwapMultiCurrencyOneToManyOrderCreateResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PreviewFlashSwapMultiCurrencyOneToManyOrder

> FlashSwapMultiCurrencyOneToManyOrderPreviewResp PreviewFlashSwapMultiCurrencyOneToManyOrder(ctx, flashSwapMultiCurrencyOneToManyOrderPreviewReq)

Flash Swap - Multi-currency exchange - Preview (one-to-many)

Preview quote for single currency to multiple target currencies exchange

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flashSwapMultiCurrencyOneToManyOrderPreviewReq** | [**FlashSwapMultiCurrencyOneToManyOrderPreviewReq**](FlashSwapMultiCurrencyOneToManyOrderPreviewReq.md)|  | 

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
    flashSwapMultiCurrencyOneToManyOrderPreviewReq := gateapi.FlashSwapMultiCurrencyOneToManyOrderPreviewReq{} // FlashSwapMultiCurrencyOneToManyOrderPreviewReq - 
    
    result, _, err := client.FlashSwapApi.PreviewFlashSwapMultiCurrencyOneToManyOrder(ctx, flashSwapMultiCurrencyOneToManyOrderPreviewReq)
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

[**FlashSwapMultiCurrencyOneToManyOrderPreviewResp**](FlashSwapMultiCurrencyOneToManyOrderPreviewResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PreviewFlashSwapOrderV1

> FlashSwapOrderPreviewResp PreviewFlashSwapOrderV1(ctx, sellAsset, buyAsset, optional)

Flash Swap - Preview (one-to-one)

Get one-to-one flash swap quote. Either sell_amount or buy_amount must be specified

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellAsset** | **string**| Currency to sell | 
**buyAsset** | **string**| Currency to buy | 
**optional** | **PreviewFlashSwapOrderV1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PreviewFlashSwapOrderV1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sellAmount** | **optional.String**| Sell amount, either this or buy_amount must be specified | 
**buyAmount** | **optional.String**| Buy amount, either this or sell_amount must be specified | 

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
    sellAsset := "sellAsset_example" // string - Currency to sell
    buyAsset := "buyAsset_example" // string - Currency to buy
    
    result, _, err := client.FlashSwapApi.PreviewFlashSwapOrderV1(ctx, sellAsset, buyAsset, nil)
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

[**FlashSwapOrderPreviewResp**](FlashSwapOrderPreviewResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
