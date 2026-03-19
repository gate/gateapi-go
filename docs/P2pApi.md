# P2pApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**P2pMerchantAccountGetUserInfo**](P2pApi.md#P2pMerchantAccountGetUserInfo) | **Post** /p2p/merchant/account/get_user_info | Get account information
[**P2pMerchantAccountGetCounterpartyUserInfo**](P2pApi.md#P2pMerchantAccountGetCounterpartyUserInfo) | **Post** /p2p/merchant/account/get_counterparty_user_info | Get counterparty information
[**P2pMerchantAccountGetMyselfPayment**](P2pApi.md#P2pMerchantAccountGetMyselfPayment) | **Post** /p2p/merchant/account/get_myself_payment | Get payment method list
[**P2pMerchantTransactionGetPendingTransactionList**](P2pApi.md#P2pMerchantTransactionGetPendingTransactionList) | **Post** /p2p/merchant/transaction/get_pending_transaction_list | Get pending orders
[**P2pMerchantTransactionGetCompletedTransactionList**](P2pApi.md#P2pMerchantTransactionGetCompletedTransactionList) | **Post** /p2p/merchant/transaction/get_completed_transaction_list | Get all/historical orders
[**P2pMerchantTransactionGetTransactionDetails**](P2pApi.md#P2pMerchantTransactionGetTransactionDetails) | **Post** /p2p/merchant/transaction/get_transaction_details | Query order details
[**P2pMerchantTransactionConfirmPayment**](P2pApi.md#P2pMerchantTransactionConfirmPayment) | **Post** /p2p/merchant/transaction/confirm-payment | Confirm payment
[**P2pMerchantTransactionConfirmReceipt**](P2pApi.md#P2pMerchantTransactionConfirmReceipt) | **Post** /p2p/merchant/transaction/confirm-receipt | Confirm receipt
[**P2pMerchantTransactionCancel**](P2pApi.md#P2pMerchantTransactionCancel) | **Post** /p2p/merchant/transaction/cancel | Cancel order
[**P2pMerchantBooksPlaceBizPushOrder**](P2pApi.md#P2pMerchantBooksPlaceBizPushOrder) | **Post** /p2p/merchant/books/place_biz_push_order | Publish ad order
[**P2pMerchantBooksAdsUpdateStatus**](P2pApi.md#P2pMerchantBooksAdsUpdateStatus) | **Post** /p2p/merchant/books/ads_update_status | Update ad status
[**P2pMerchantBooksAdsDetail**](P2pApi.md#P2pMerchantBooksAdsDetail) | **Post** /p2p/merchant/books/ads_detail | Query ad details
[**P2pMerchantBooksMyAdsList**](P2pApi.md#P2pMerchantBooksMyAdsList) | **Post** /p2p/merchant/books/my_ads_list | Get my ad list
[**P2pMerchantBooksAdsList**](P2pApi.md#P2pMerchantBooksAdsList) | **Post** /p2p/merchant/books/ads_list | Get Advertisement List
[**P2pMerchantChatGetChatsList**](P2pApi.md#P2pMerchantChatGetChatsList) | **Post** /p2p/merchant/chat/get_chats_list | Get chat history
[**P2pMerchantChatSendChatMessage**](P2pApi.md#P2pMerchantChatSendChatMessage) | **Post** /p2p/merchant/chat/send_chat_message | Send text message
[**P2pMerchantChatUploadChatFile**](P2pApi.md#P2pMerchantChatUploadChatFile) | **Post** /p2p/merchant/chat/upload_chat_file | Upload chat file


## P2pMerchantAccountGetUserInfo

> P2pMerchantUserInfoResponse P2pMerchantAccountGetUserInfo(ctx, )

Get account information

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
    
    result, _, err := client.P2pApi.P2pMerchantAccountGetUserInfo(ctx)
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

[**P2pMerchantUserInfoResponse**](P2pMerchantUserInfoResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantAccountGetCounterpartyUserInfo

> P2pCounterpartyUserInfoResponse P2pMerchantAccountGetCounterpartyUserInfo(ctx, getCounterpartyUserInfoRequest)

Get counterparty information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getCounterpartyUserInfoRequest** | [**GetCounterpartyUserInfoRequest**](GetCounterpartyUserInfoRequest.md)|  | 

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
    getCounterpartyUserInfoRequest := gateapi.GetCounterpartyUserInfoRequest{} // GetCounterpartyUserInfoRequest - 
    
    result, _, err := client.P2pApi.P2pMerchantAccountGetCounterpartyUserInfo(ctx, getCounterpartyUserInfoRequest)
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

[**P2pCounterpartyUserInfoResponse**](P2pCounterpartyUserInfoResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantAccountGetMyselfPayment

> P2pPaymentMethodsResponse P2pMerchantAccountGetMyselfPayment(ctx, optional)

Get payment method list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **P2pMerchantAccountGetMyselfPaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantAccountGetMyselfPaymentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getMyselfPaymentRequest** | [**optional.Interface of GetMyselfPaymentRequest**](GetMyselfPaymentRequest.md)|  | 

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
    
    result, _, err := client.P2pApi.P2pMerchantAccountGetMyselfPayment(ctx, nil)
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

[**P2pPaymentMethodsResponse**](P2pPaymentMethodsResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionGetPendingTransactionList

> P2pTransactionListResponse P2pMerchantTransactionGetPendingTransactionList(ctx, getPendingTransactionListRequest)

Get pending orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getPendingTransactionListRequest** | [**GetPendingTransactionListRequest**](GetPendingTransactionListRequest.md)|  | 

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
    getPendingTransactionListRequest := gateapi.GetPendingTransactionListRequest{} // GetPendingTransactionListRequest - 
    
    result, _, err := client.P2pApi.P2pMerchantTransactionGetPendingTransactionList(ctx, getPendingTransactionListRequest)
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

[**P2pTransactionListResponse**](P2pTransactionListResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionGetCompletedTransactionList

> P2pTransactionListResponse P2pMerchantTransactionGetCompletedTransactionList(ctx, getCompletedTransactionListRequest)

Get all/historical orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getCompletedTransactionListRequest** | [**GetCompletedTransactionListRequest**](GetCompletedTransactionListRequest.md)|  | 

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
    getCompletedTransactionListRequest := gateapi.GetCompletedTransactionListRequest{} // GetCompletedTransactionListRequest - 
    
    result, _, err := client.P2pApi.P2pMerchantTransactionGetCompletedTransactionList(ctx, getCompletedTransactionListRequest)
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

[**P2pTransactionListResponse**](P2pTransactionListResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionGetTransactionDetails

> P2pTransactionDetailResponse P2pMerchantTransactionGetTransactionDetails(ctx, getTransactionDetailsRequest)

Query order details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getTransactionDetailsRequest** | [**GetTransactionDetailsRequest**](GetTransactionDetailsRequest.md)|  | 

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
    getTransactionDetailsRequest := gateapi.GetTransactionDetailsRequest{} // GetTransactionDetailsRequest - 
    
    result, _, err := client.P2pApi.P2pMerchantTransactionGetTransactionDetails(ctx, getTransactionDetailsRequest)
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

[**P2pTransactionDetailResponse**](P2pTransactionDetailResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionConfirmPayment

> P2pTransactionActionResponse P2pMerchantTransactionConfirmPayment(ctx, confirmPayment)

Confirm payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**confirmPayment** | [**ConfirmPayment**](ConfirmPayment.md)|  | 

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
    confirmPayment := gateapi.ConfirmPayment{} // ConfirmPayment - 
    
    result, _, err := client.P2pApi.P2pMerchantTransactionConfirmPayment(ctx, confirmPayment)
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

[**P2pTransactionActionResponse**](P2pTransactionActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionConfirmReceipt

> P2pTransactionActionResponse P2pMerchantTransactionConfirmReceipt(ctx, confirmReceipt)

Confirm receipt

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**confirmReceipt** | [**ConfirmReceipt**](ConfirmReceipt.md)|  | 

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
    confirmReceipt := gateapi.ConfirmReceipt{} // ConfirmReceipt - 
    
    result, _, err := client.P2pApi.P2pMerchantTransactionConfirmReceipt(ctx, confirmReceipt)
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

[**P2pTransactionActionResponse**](P2pTransactionActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionCancel

> P2pTransactionActionResponse P2pMerchantTransactionCancel(ctx, cancelOrder)

Cancel order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cancelOrder** | [**CancelOrder**](CancelOrder.md)|  | 

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
    cancelOrder := gateapi.CancelOrder{} // CancelOrder - 
    
    result, _, err := client.P2pApi.P2pMerchantTransactionCancel(ctx, cancelOrder)
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

[**P2pTransactionActionResponse**](P2pTransactionActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantBooksPlaceBizPushOrder

> map[string]interface{} P2pMerchantBooksPlaceBizPushOrder(ctx, placeBizPushOrder)

Publish ad order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeBizPushOrder** | [**PlaceBizPushOrder**](PlaceBizPushOrder.md)|  | 

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
    placeBizPushOrder := gateapi.PlaceBizPushOrder{} // PlaceBizPushOrder - 
    
    result, _, err := client.P2pApi.P2pMerchantBooksPlaceBizPushOrder(ctx, placeBizPushOrder)
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

**map[string]interface{}**

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantBooksAdsUpdateStatus

> P2pAdsUpdateStatusResponse P2pMerchantBooksAdsUpdateStatus(ctx, adsUpdateStatus, optional)

Update ad status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adsUpdateStatus** | [**AdsUpdateStatus**](AdsUpdateStatus.md)|  | 
**optional** | **P2pMerchantBooksAdsUpdateStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantBooksAdsUpdateStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**tradeType** | **optional.String**| Project-Id-Version: GateApiTools 1.0.0 Report-Msgid-Bugs-To: EMAIL@ADDRESS POT-Creation-Date: 2025-11-12 18:14+0800 PO-Revision-Date: 2019-01-02 17:30+0800 Last-Translator: FULL NAME &lt;EMAIL@ADDRESS&gt; Language: en Language-Team: en &lt;L@li.org&gt; Plural-Forms: nplurals&#x3D;2; plural&#x3D;(n !&#x3D;1) MIME-Version: 1.0 Content-Type: text/plain; charset&#x3D;utf-8 Content-Transfer-Encoding: 8bit Generated-By: Babel 2.8.0  | 

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
    adsUpdateStatus := gateapi.AdsUpdateStatus{} // AdsUpdateStatus - 
    
    result, _, err := client.P2pApi.P2pMerchantBooksAdsUpdateStatus(ctx, adsUpdateStatus, nil)
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

[**P2pAdsUpdateStatusResponse**](P2pAdsUpdateStatusResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantBooksAdsDetail

> P2pAdDetailResponse P2pMerchantBooksAdsDetail(ctx, adsDetailRequest)

Query ad details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adsDetailRequest** | [**AdsDetailRequest**](AdsDetailRequest.md)|  | 

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
    adsDetailRequest := gateapi.AdsDetailRequest{} // AdsDetailRequest - 
    
    result, _, err := client.P2pApi.P2pMerchantBooksAdsDetail(ctx, adsDetailRequest)
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

[**P2pAdDetailResponse**](P2pAdDetailResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantBooksMyAdsList

> P2pMyAdsListResponse P2pMerchantBooksMyAdsList(ctx, optional)

Get my ad list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **P2pMerchantBooksMyAdsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantBooksMyAdsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**myAdsListRequest** | [**optional.Interface of MyAdsListRequest**](MyAdsListRequest.md)|  | 

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
    
    result, _, err := client.P2pApi.P2pMerchantBooksMyAdsList(ctx, nil)
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

[**P2pMyAdsListResponse**](P2pMyAdsListResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantBooksAdsList

> P2pAdsListResponse P2pMerchantBooksAdsList(ctx, adsListRequest)

Get Advertisement List

Project-Id-Version: GateApiTools 1.0.0 Report-Msgid-Bugs-To: EMAIL@ADDRESS POT-Creation-Date: 2025-11-12 18:14+0800 PO-Revision-Date: 2019-01-02 17:30+0800 Last-Translator: FULL NAME <EMAIL@ADDRESS> Language: en Language-Team: en <L@li.org> Plural-Forms: nplurals=2; plural=(n !=1) MIME-Version: 1.0 Content-Type: text/plain; charset=utf-8 Content-Transfer-Encoding: 8bit Generated-By: Babel 2.8.0 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adsListRequest** | [**AdsListRequest**](AdsListRequest.md)|  | 

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
    adsListRequest := gateapi.AdsListRequest{} // AdsListRequest - 
    
    result, _, err := client.P2pApi.P2pMerchantBooksAdsList(ctx, adsListRequest)
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

[**P2pAdsListResponse**](P2pAdsListResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantChatGetChatsList

> P2pChatListResponse P2pMerchantChatGetChatsList(ctx, getChatsListRequest)

Get chat history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getChatsListRequest** | [**GetChatsListRequest**](GetChatsListRequest.md)|  | 

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
    getChatsListRequest := gateapi.GetChatsListRequest{} // GetChatsListRequest - 
    
    result, _, err := client.P2pApi.P2pMerchantChatGetChatsList(ctx, getChatsListRequest)
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

[**P2pChatListResponse**](P2pChatListResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantChatSendChatMessage

> P2pSendChatMessageResponse P2pMerchantChatSendChatMessage(ctx, sendChatMessageRequest)

Send text message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sendChatMessageRequest** | [**SendChatMessageRequest**](SendChatMessageRequest.md)|  | 

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
    sendChatMessageRequest := gateapi.SendChatMessageRequest{} // SendChatMessageRequest - 
    
    result, _, err := client.P2pApi.P2pMerchantChatSendChatMessage(ctx, sendChatMessageRequest)
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

[**P2pSendChatMessageResponse**](P2pSendChatMessageResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantChatUploadChatFile

> P2pUploadChatFileResponse P2pMerchantChatUploadChatFile(ctx, uploadChatFile)

Upload chat file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadChatFile** | [**UploadChatFile**](UploadChatFile.md)|  | 

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
    uploadChatFile := gateapi.UploadChatFile{} // UploadChatFile - 
    
    result, _, err := client.P2pApi.P2pMerchantChatUploadChatFile(ctx, uploadChatFile)
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

[**P2pUploadChatFileResponse**](P2pUploadChatFileResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
