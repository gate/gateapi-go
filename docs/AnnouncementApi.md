# AnnouncementApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAnnouncementArticles**](AnnouncementApi.md#ListAnnouncementArticles) | **Post** /ann/list_article | List announcement articles


## ListAnnouncementArticles

> AnnouncementArticleListResponse ListAnnouncementArticles(ctx, announcementArticleListRequest)

List announcement articles

Query announcement articles with pagination and filters for title, category, language, time, and other criteria. Send query parameters in the JSON request body. Both page and size are optional and must be strings when provided. No API key is required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**announcementArticleListRequest** | [**AnnouncementArticleListRequest**](AnnouncementArticleListRequest.md)|  | 

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
    announcementArticleListRequest := gateapi.AnnouncementArticleListRequest{} // AnnouncementArticleListRequest - 
    
    result, _, err := client.AnnouncementApi.ListAnnouncementArticles(ctx, announcementArticleListRequest)
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

[**AnnouncementArticleListResponse**](AnnouncementArticleListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
