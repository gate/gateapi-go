# SquareApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSquareAiSearch**](SquareApi.md#ListSquareAiSearch) | **Get** /social/message/search | AI MCP Dynamic Search
[**ListLiveReplay**](SquareApi.md#ListLiveReplay) | **Get** /social/live/tag_coin_live_replay | Gate AI Assistant live stream data retrieval


## ListSquareAiSearch

> InlineResponse2009 ListSquareAiSearch(ctx, optional)

AI MCP Dynamic Search

Dynamic search endpoint for AI MCP platform. All parameters are passed via query string. Returns simplified fields. Designed for LLM function calling / MCP tool scenarios.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSquareAiSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSquareAiSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**keyword** | **optional.String**| Search keyword (currency name or content keyword, e.g., BTC, ETH) | 
**currency** | **optional.String**| Filter by currency (exact currency code, e.g., BTC, ETH, SOL) | 
**timeRange** | **optional.Int32**| Time range: 0 &#x3D; unlimited (default), 1 &#x3D; last day, 2 &#x3D; last week, 3 &#x3D; last month | [default to 0]
**sort** | **optional.Int32**| Sort order: 0 &#x3D; most popular (default), 1 &#x3D; latest | [default to 0]
**limit** | **optional.Int32**| Return count, 1-50, default 10 | [default to 10]
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
    
    result, _, err := client.SquareApi.ListSquareAiSearch(ctx, nil)
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

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLiveReplay

> InlineResponse20010 ListLiveReplay(ctx, optional)

Gate AI Assistant live stream data retrieval

AI assistant live stream/replay search endpoint. Filters live rooms or replay videos by business tags and currency. Each record in the returned list is distinguished by content_type: streaming = live broadcast (live field has value), video = replay video (video field has value). The number of results is controlled by the limit parameter (max 10), no additional pagination needed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListLiveReplayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLiveReplayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**tag** | **optional.String**| Business type filter. Available values: Market Analysis, Hot Topics, Blockchain, Others | 
**coin** | **optional.String**| Filter by currency name (e.g., BTC, ETH) | 
**sort** | **optional.String**| Sort order: hot &#x3D; most popular (default), new &#x3D; latest | [default to hot]
**limit** | **optional.Int32**| Return count, 1-10, default 3 | [default to 3]

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
    
    result, _, err := client.SquareApi.ListLiveReplay(ctx, nil)
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

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
