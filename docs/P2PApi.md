# P2PApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**P2pMerchantAccountGetUserInfo**](P2PApi.md#P2pMerchantAccountGetUserInfo) | **Post** /p2p/merchant/account/get_user_info | Get account information
[**P2pMerchantAccountGetCounterpartyUserInfo**](P2PApi.md#P2pMerchantAccountGetCounterpartyUserInfo) | **Post** /p2p/merchant/account/get_counterparty_user_info | Get counterparty information
[**P2pMerchantAccountGetMyselfPayment**](P2PApi.md#P2pMerchantAccountGetMyselfPayment) | **Post** /p2p/merchant/account/get_myself_payment | Get payment method list
[**P2pMerchantTransactionGetPendingTransactionList**](P2PApi.md#P2pMerchantTransactionGetPendingTransactionList) | **Post** /p2p/merchant/transaction/get_pending_transaction_list | Get pending orders
[**P2pMerchantTransactionGetCompletedTransactionList**](P2PApi.md#P2pMerchantTransactionGetCompletedTransactionList) | **Post** /p2p/merchant/transaction/get_completed_transaction_list | Get all/historical orders
[**P2pMerchantTransactionGetTransactionDetails**](P2PApi.md#P2pMerchantTransactionGetTransactionDetails) | **Post** /p2p/merchant/transaction/get_transaction_details | Query order details
[**P2pMerchantTransactionConfirmPayment**](P2PApi.md#P2pMerchantTransactionConfirmPayment) | **Post** /p2p/merchant/transaction/confirm-payment | Confirm payment
[**P2pMerchantTransactionConfirmReceipt**](P2PApi.md#P2pMerchantTransactionConfirmReceipt) | **Post** /p2p/merchant/transaction/confirm-receipt | Confirm receipt
[**P2pMerchantTransactionCancel**](P2PApi.md#P2pMerchantTransactionCancel) | **Post** /p2p/merchant/transaction/cancel | Cancel order
[**P2pMerchantBooksPlaceBizPushOrder**](P2PApi.md#P2pMerchantBooksPlaceBizPushOrder) | **Post** /p2p/merchant/books/place_biz_push_order | Publish ad order
[**P2pMerchantBooksAdsUpdateStatus**](P2PApi.md#P2pMerchantBooksAdsUpdateStatus) | **Post** /p2p/merchant/books/ads_update_status | Update ad status
[**P2pMerchantBooksAdsDetail**](P2PApi.md#P2pMerchantBooksAdsDetail) | **Post** /p2p/merchant/books/ads_detail | Query ad details
[**P2pMerchantBooksMyAdsList**](P2PApi.md#P2pMerchantBooksMyAdsList) | **Post** /p2p/merchant/books/my_ads_list | Get my ad list
[**P2pMerchantChatGetChatsList**](P2PApi.md#P2pMerchantChatGetChatsList) | **Post** /p2p/merchant/chat/get_chats_list | Get chat history
[**P2pMerchantChatSendChatMessage**](P2PApi.md#P2pMerchantChatSendChatMessage) | **Post** /p2p/merchant/chat/send_chat_message | Send text message
[**P2pMerchantChatUploadChatFile**](P2PApi.md#P2pMerchantChatUploadChatFile) | **Post** /p2p/merchant/chat/upload_chat_file | Upload chat file


## P2pMerchantAccountGetUserInfo

> InlineResponse20013 P2pMerchantAccountGetUserInfo(ctx, )

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
    ctx := context.Background()
    
    result, _, err := client.P2PApi.P2pMerchantAccountGetUserInfo(ctx)
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

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantAccountGetCounterpartyUserInfo

> InlineResponse20014 P2pMerchantAccountGetCounterpartyUserInfo(ctx, bizUid)

Get counterparty information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bizUid** | **string**| Counterparty UID (encrypted) | 

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
    bizUid := "bizUid_example" // string - Counterparty UID (encrypted)
    
    result, _, err := client.P2PApi.P2pMerchantAccountGetCounterpartyUserInfo(ctx, bizUid)
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

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantAccountGetMyselfPayment

> InlineResponse20015 P2pMerchantAccountGetMyselfPayment(ctx, optional)

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
**fiat** | **optional.String**| Fiat currency | 

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
    
    result, _, err := client.P2PApi.P2pMerchantAccountGetMyselfPayment(ctx, nil)
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

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionGetPendingTransactionList

> InlineResponse20016 P2pMerchantTransactionGetPendingTransactionList(ctx, cryptoCurrency, fiatCurrency, optional)

Get pending orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cryptoCurrency** | **string**| Cryptocurrency | 
**fiatCurrency** | **string**| Fiat currency | 
**optional** | **P2pMerchantTransactionGetPendingTransactionListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantTransactionGetPendingTransactionListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**orderTab** | **optional.String**| Order tab, default is pending (pending: Processing (pending: AND status in (&#39;OPEN&#39;,  &#39;PAID&#39;, &#39;LOCKED&#39;, &#39;TEMP&#39;)); dispute: In dispute (status in (&#39;ACCEPT&#39;,  &#39;BCLOSED&#39;, &#39;CANCEL&#39;, &#39;BECANCEL&#39;, &#39;SCLOSED&#39;, &#39;SCANCEL&#39;))) | 
**selectType** | **optional.String**| Buy/Sell (sell&#x3D;Sell, buy&#x3D;Buy, others&#x3D;All) | 
**status** | **optional.String**| 订单状态（dispute: 申诉订单； closed: ACCEPT、BCLOSED； cancel： CANCEL、BECANCEL、SCLOSED、SCANCEL； locked: LOCKED； open: OPEN； paid： PAID； completed： CANCEL、BECANCEL、SCLOSED、SCANCEL、ACCEPT、BCLOSED） | 
**txid** | **optional.Int32**| Order ID | 
**startTime** | **optional.Int32**| Start timestamp, default is 00:00 89 days ago | 
**endTime** | **optional.Int32**| End timestamp, default is 23:59:59 today | 

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
    cryptoCurrency := "cryptoCurrency_example" // string - Cryptocurrency
    fiatCurrency := "fiatCurrency_example" // string - Fiat currency
    
    result, _, err := client.P2PApi.P2pMerchantTransactionGetPendingTransactionList(ctx, cryptoCurrency, fiatCurrency, nil)
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

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionGetCompletedTransactionList

> InlineResponse20016 P2pMerchantTransactionGetCompletedTransactionList(ctx, cryptoCurrency, fiatCurrency, optional)

Get all/historical orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cryptoCurrency** | **string**| Cryptocurrency | 
**fiatCurrency** | **string**| Fiat currency | 
**optional** | **P2pMerchantTransactionGetCompletedTransactionListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantTransactionGetCompletedTransactionListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**selectType** | **optional.String**| Buy/Sell (sell&#x3D;Sell, buy&#x3D;Buy, others&#x3D;All) | 
**status** | **optional.String**| 订单状态（dispute: 申诉订单； closed: ACCEPT、BCLOSED； cancel： CANCEL、BECANCEL、SCLOSED、SCANCEL； locked: LOCKED； open: OPEN； paid： PAID； completed： CANCEL、BECANCEL、SCLOSED、SCANCEL、ACCEPT、BCLOSED） | 
**txid** | **optional.Int32**| Order ID | 
**startTime** | **optional.Int32**| Start timestamp, default is 00:00 89 days ago | 
**endTime** | **optional.Int32**| End timestamp, default is 23:59:59 today | 
**queryDispute** | **optional.Int32**| 1: Include appeal status, 0: None | 
**page** | **optional.Int32**| page number | 
**perPage** | **optional.Int32**| Number of orders per page | 

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
    cryptoCurrency := "cryptoCurrency_example" // string - Cryptocurrency
    fiatCurrency := "fiatCurrency_example" // string - Fiat currency
    
    result, _, err := client.P2PApi.P2pMerchantTransactionGetCompletedTransactionList(ctx, cryptoCurrency, fiatCurrency, nil)
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

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionGetTransactionDetails

> InlineResponse20017 P2pMerchantTransactionGetTransactionDetails(ctx, txid, optional)

Query order details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txid** | **int32**| Order ID | 
**optional** | **P2pMerchantTransactionGetTransactionDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantTransactionGetTransactionDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**channel** | **optional.String**| Empty or web3 | 

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
    txid := 56 // int32 - Order ID
    
    result, _, err := client.P2PApi.P2pMerchantTransactionGetTransactionDetails(ctx, txid, nil)
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

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionConfirmPayment

> InlineResponse2007 P2pMerchantTransactionConfirmPayment(ctx, optional)

Confirm payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **P2pMerchantTransactionConfirmPaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantTransactionConfirmPaymentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**inlineObject10** | [**optional.Interface of InlineObject10**](InlineObject10.md)|  | 

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
    
    result, _, err := client.P2PApi.P2pMerchantTransactionConfirmPayment(ctx, nil)
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

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionConfirmReceipt

> InlineResponse2007 P2pMerchantTransactionConfirmReceipt(ctx, optional)

Confirm receipt

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **P2pMerchantTransactionConfirmReceiptOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantTransactionConfirmReceiptOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**inlineObject11** | [**optional.Interface of InlineObject11**](InlineObject11.md)|  | 

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
    
    result, _, err := client.P2PApi.P2pMerchantTransactionConfirmReceipt(ctx, nil)
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

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantTransactionCancel

> InlineResponse2007 P2pMerchantTransactionCancel(ctx, optional)

Cancel order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **P2pMerchantTransactionCancelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantTransactionCancelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**inlineObject12** | [**optional.Interface of InlineObject12**](InlineObject12.md)|  | 

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
    
    result, _, err := client.P2PApi.P2pMerchantTransactionCancel(ctx, nil)
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

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantBooksPlaceBizPushOrder

> map[string]interface{} P2pMerchantBooksPlaceBizPushOrder(ctx, currencyType, exchangeType, type_, unitPrice, number, minAmount, maxAmount, optional)

Publish ad order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyType** | **string**| Cryptocurrency | 
**exchangeType** | **string**| Fiat currency | 
**type_** | **string**| Ad type: 0&#x3D;Sell, 1&#x3D;Buy, 2&#x3D;Edit sell, 3&#x3D;Edit buy | 
**unitPrice** | **string**| Unit price | 
**number** | **string**| Size | 
**minAmount** | **string**| Minimum transaction amount per order | 
**maxAmount** | **string**| Maximum transaction amount per order | 
**optional** | **P2pMerchantBooksPlaceBizPushOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantBooksPlaceBizPushOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**payType** | **optional.String**| Payment method | 
**payTypeJson** | **optional.String**| Payment method JSON string | 
**rateFixed** | **optional.String**| Price type: 0-Floating price, 1-Fixed price | 
**oid** | **optional.String**| Ad ID when editing | 
**tierLimit** | **optional.String**| Order tier limit | 
**verifiedLimit** | **optional.String**| Verification level limit | 
**regTimeLimit** | **optional.String**| Registration time limit | 
**advertisersLimit** | **optional.String**| Advertiser restriction | 
**hidePayment** | **optional.String**| Whether to hide payment method: 1&#x3D;Yes, 0&#x3D;No | 
**expireMin** | **optional.String**| Ad expiration time (minutes) | 
**tradeTips** | **optional.String**| Trading terms | 
**autoReply** | **optional.String**| Auto reply | 
**minCompletedLimit** | **optional.String**| Minimum limit of completed orders | 
**maxCompletedLimit** | **optional.String**| Maximum limit of completed orders | 
**completedRateLimit** | **optional.String**| 30-day completion rate limit | 
**userCountryLimit** | **optional.String**| KYC nationality restriction | 
**userOrderLimit** | **optional.String**| Order count limit | 
**rateReferenceId** | **optional.String**| Reference exchange rate ID | 
**rateOffset** | **optional.String**| Reference exchange rate offset | 
**floatTrend** | **optional.String**| 444 | 

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
    currencyType := "currencyType_example" // string - Cryptocurrency
    exchangeType := "exchangeType_example" // string - Fiat currency
    type_ := "type__example" // string - Ad type: 0=Sell, 1=Buy, 2=Edit sell, 3=Edit buy
    unitPrice := "unitPrice_example" // string - Unit price
    number := "number_example" // string - Size
    minAmount := "minAmount_example" // string - Minimum transaction amount per order
    maxAmount := "maxAmount_example" // string - Maximum transaction amount per order
    
    result, _, err := client.P2PApi.P2pMerchantBooksPlaceBizPushOrder(ctx, currencyType, exchangeType, type_, unitPrice, number, minAmount, maxAmount, nil)
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

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantBooksAdsUpdateStatus

> InlineResponse20018 P2pMerchantBooksAdsUpdateStatus(ctx, advNo, advStatus, optional)

Update ad status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**advNo** | **int32**| Ad ID | 
**advStatus** | **int32**| Ad status: 1&#x3D;Active, 3&#x3D;Inactive, 4&#x3D;Closed | 
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
    ctx := context.Background()
    advNo := 56 // int32 - Ad ID
    advStatus := 56 // int32 - Ad status: 1=Active, 3=Inactive, 4=Closed
    
    result, _, err := client.P2PApi.P2pMerchantBooksAdsUpdateStatus(ctx, advNo, advStatus, nil)
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

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantBooksAdsDetail

> InlineResponse20019 P2pMerchantBooksAdsDetail(ctx, advNo)

Query ad details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**advNo** | **string**|  | 

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
    advNo := "advNo_example" // string - 
    
    result, _, err := client.P2PApi.P2pMerchantBooksAdsDetail(ctx, advNo)
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

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantBooksMyAdsList

> InlineResponse20020 P2pMerchantBooksMyAdsList(ctx, optional)

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
**asset** | **optional.String**| Cryptocurrency | 
**fiatUnit** | **optional.String**| Fiat currency | 
**tradeType** | **optional.String**| Buy/Sell | 

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
    
    result, _, err := client.P2PApi.P2pMerchantBooksMyAdsList(ctx, nil)
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

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantChatGetChatsList

> InlineResponse20021 P2pMerchantChatGetChatsList(ctx, txid, optional)

Get chat history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txid** | **int32**| Order ID | 
**optional** | **P2pMerchantChatGetChatsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantChatGetChatsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**lastreceived** | **optional.Int32**| Pagination timestamp (forward) | 
**firstreceived** | **optional.Int32**| Pagination timestamp (backward) | 

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
    txid := 56 // int32 - Order ID
    
    result, _, err := client.P2PApi.P2pMerchantChatGetChatsList(ctx, txid, nil)
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

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantChatSendChatMessage

> InlineResponse20022 P2pMerchantChatSendChatMessage(ctx, txid, message, optional)

Send text message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txid** | **int32**| Order ID | 
**message** | **string**| Message content | 
**optional** | **P2pMerchantChatSendChatMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a P2pMerchantChatSendChatMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**type_** | **optional.Int32**| 0&#x3D;Text, 1&#x3D;File (video or image), default is 0 if not provided | 

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
    txid := 56 // int32 - Order ID
    message := "message_example" // string - Message content
    
    result, _, err := client.P2PApi.P2pMerchantChatSendChatMessage(ctx, txid, message, nil)
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

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## P2pMerchantChatUploadChatFile

> InlineResponse20023 P2pMerchantChatUploadChatFile(ctx, imageContentType, base64Img)

Upload chat file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageContentType** | **string**| File type, currently only images and videos are supported | 
**base64Img** | **string**| File content (base64 encoded) | 

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
    imageContentType := "imageContentType_example" // string - File type, currently only images and videos are supported
    base64Img := "base64Img_example" // string - File content (base64 encoded)
    
    result, _, err := client.P2PApi.P2pMerchantChatUploadChatFile(ctx, imageContentType, base64Img)
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

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
