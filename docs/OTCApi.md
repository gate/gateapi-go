# OTCApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOtcQuote**](OTCApi.md#CreateOtcQuote) | **Post** /otc/quote | Fiat and stablecoin quote
[**CreateOtcOrder**](OTCApi.md#CreateOtcOrder) | **Post** /otc/order/create | Create fiat order
[**CreateStableCoinOrder**](OTCApi.md#CreateStableCoinOrder) | **Post** /otc/stable_coin/order/create | Create stablecoin order
[**GetBankListInnerPath**](OTCApi.md#GetBankListInnerPath) | **Get** /otc/bank/list | Get user bank card list
[**CreateOtcBank**](OTCApi.md#CreateOtcBank) | **Post** /otc/bank/create | Create bank card
[**DeleteOtcBank**](OTCApi.md#DeleteOtcBank) | **Post** /otc/bank/delete | Delete bank card
[**SetDefaultOtcBank**](OTCApi.md#SetDefaultOtcBank) | **Post** /otc/bank/set_default | Set default bank card
[**GetOtcBankSupplementChecklist**](OTCApi.md#GetOtcBankSupplementChecklist) | **Get** /otc/bank/bank_supplement_checklist | Query the checklist of materials to supplement for a bank card
[**SubmitOtcBankPersonalSupplement**](OTCApi.md#SubmitOtcBankPersonalSupplement) | **Post** /otc/bank/personal/bank_supplement | Submit Bank Card Supplement Materials (Personal)
[**SubmitOtcBankEnterpriseSupplement**](OTCApi.md#SubmitOtcBankEnterpriseSupplement) | **Post** /otc/bank/enterprise/bank_supplement | Submit Bank Card Supplement Materials (Enterprise)
[**CreateOtcUploadPreUpload**](OTCApi.md#CreateOtcUploadPreUpload) | **Post** /otc/upload/pre_upload | Pre-upload file (temporary bucket)
[**MarkOtcOrderPaid**](OTCApi.md#MarkOtcOrderPaid) | **Post** /otc/order/paid | Mark fiat order as paid (deposit confirmation)
[**CancelOtcOrder**](OTCApi.md#CancelOtcOrder) | **Post** /otc/order/cancel | Fiat order cancellation
[**ListOtcOrders**](OTCApi.md#ListOtcOrders) | **Get** /otc/order/list | Fiat order list
[**ListStableCoinOrders**](OTCApi.md#ListStableCoinOrders) | **Get** /otc/stable_coin/order/list | Stablecoin order list
[**GetOtcOrderDetail**](OTCApi.md#GetOtcOrderDetail) | **Get** /otc/order/detail | Fiat order details


## CreateOtcQuote

> OtcQuoteResponse CreateOtcQuote(ctx, otcQuoteRequest)

Fiat and stablecoin quote

Create fiat and stablecoin quotes, supporting both PAY and GET directions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otcQuoteRequest** | [**OtcQuoteRequest**](OtcQuoteRequest.md)|  | 

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
    otcQuoteRequest := gateapi.OtcQuoteRequest{} // OtcQuoteRequest - 
    
    result, _, err := client.OTCApi.CreateOtcQuote(ctx, otcQuoteRequest)
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

[**OtcQuoteResponse**](OtcQuoteResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateOtcOrder

> OtcActionResponse CreateOtcOrder(ctx, otcOrderRequest)

Create fiat order

Create a fiat order, supporting BUY for on-ramp and SELL for off-ramp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otcOrderRequest** | [**OtcOrderRequest**](OtcOrderRequest.md)|  | 

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
    otcOrderRequest := gateapi.OtcOrderRequest{} // OtcOrderRequest - 
    
    result, _, err := client.OTCApi.CreateOtcOrder(ctx, otcOrderRequest)
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

[**OtcActionResponse**](OtcActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateStableCoinOrder

> OtcStableCoinOrderCreateResponse CreateStableCoinOrder(ctx, otcStableCoinOrderRequest)

Create stablecoin order

Create a stablecoin order. All request body fields except `promotion_code` are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otcStableCoinOrderRequest** | [**OtcStableCoinOrderRequest**](OtcStableCoinOrderRequest.md)|  | 

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
    otcStableCoinOrderRequest := gateapi.OtcStableCoinOrderRequest{} // OtcStableCoinOrderRequest - 
    
    result, _, err := client.OTCApi.CreateStableCoinOrder(ctx, otcStableCoinOrderRequest)
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

[**OtcStableCoinOrderCreateResponse**](OtcStableCoinOrderCreateResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetBankListInnerPath

> OtcBankListResponse GetBankListInnerPath(ctx, )

Get user bank card list

List the user's bank cards for selecting a card when placing an order. **Default card**: use the `is_default` field in each list item (`1` indicates the default). The deprecated standalone default-bank-card endpoint is no longer required.

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
    
    result, _, err := client.OTCApi.GetBankListInnerPath(ctx)
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

[**OtcBankListResponse**](OtcBankListResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateOtcBank

> OtcBankCreateResponse CreateOtcBank(ctx, bankAccountName, bankName, bankCountry, bankAddress, iban, swift, optional)

Create bank card

Bind a bank card. Under the Global entity, non-same-name accounts may enter manual review (`status` pending) and require supplementary materials later. Corresponds to Inner: `POST /bank/create`. Fields and protocol follow the live form/gateway; `bank_account_name` may be Base64-encoded in some environments—see integration notes.  Account-opening proof supports two methods (choose one):  1. **Pre-upload (recommended)**: call `POST /otc/upload/pre_upload` (`scene=bank`) to obtain a temporary-bucket Policy and upload directly to S3, then pass `documentation_file_key` + `file_type` in this endpoint; 2. **Multipart direct upload**: pass the `documentation_file` file field; the server writes directly to the production bucket.  When using pre-upload, the server validates object existence and that the uid in the `file_key` path matches the caller; after validation, the object is moved to the production bucket and persisted. Cross-user references return `Invalid parameters file_key`; incomplete direct upload returns `Invalid parameters file not uploaded`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankAccountName** | **string**|  | 
**bankName** | **string**|  | 
**bankCountry** | **string**|  | 
**bankAddress** | **string**|  | 
**iban** | **string**|  | 
**swift** | **string**|  | 
**optional** | **CreateOtcBankOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOtcBankOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**remittanceLineNumber** | **optional.String**|  | 
**agentBankName** | **optional.String**|  | 
**agentBankSwift** | **optional.String**|  | 
**documentationFile** | **optional.String**| Multipart direct upload; mutually exclusive with documentation_file_key | 
**documentationFileKey** | **optional.String**| Pre-upload mode; file_key returned by pre_upload (plaintext or base64 accepted) | 
**fileType** | **optional.String**| Required when using documentation_file_key; plaintext MIME or its base64 | 

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
    bankAccountName := "bankAccountName_example" // string - 
    bankName := "bankName_example" // string - 
    bankCountry := "bankCountry_example" // string - 
    bankAddress := "bankAddress_example" // string - 
    iban := "iban_example" // string - 
    swift := "swift_example" // string - 
    
    result, _, err := client.OTCApi.CreateOtcBank(ctx, bankAccountName, bankName, bankCountry, bankAddress, iban, swift, nil)
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

[**OtcBankCreateResponse**](OtcBankCreateResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteOtcBank

> OtcActionResponse DeleteOtcBank(ctx, otcBankIdRequest)

Delete bank card

Delete the specified bank card. Corresponds to Inner: `POST /bank/delete`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otcBankIdRequest** | [**OtcBankIdRequest**](OtcBankIdRequest.md)|  | 

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
    otcBankIdRequest := gateapi.OtcBankIdRequest{} // OtcBankIdRequest - 
    
    result, _, err := client.OTCApi.DeleteOtcBank(ctx, otcBankIdRequest)
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

[**OtcActionResponse**](OtcActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SetDefaultOtcBank

> OtcActionResponse SetDefaultOtcBank(ctx, otcBankIdRequest)

Set default bank card

Set the specified bank card as default. Corresponds to Inner: `POST /bank/set_default`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otcBankIdRequest** | [**OtcBankIdRequest**](OtcBankIdRequest.md)|  | 

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
    otcBankIdRequest := gateapi.OtcBankIdRequest{} // OtcBankIdRequest - 
    
    result, _, err := client.OTCApi.SetDefaultOtcBank(ctx, otcBankIdRequest)
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

[**OtcActionResponse**](OtcActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOtcBankSupplementChecklist

> OtcBankSupplementChecklistResponse GetOtcBankSupplementChecklist(ctx, bankId)

Query the checklist of materials to supplement for a bank card

**①** `bank_id` must be specified. After verifying that the card belongs to the current user and its status allows supplementary documents, the endpoint returns the required items based on the user's **approved advanced verification type** (personal/enterprise); each item's `description` states the submission requirements. Corresponding Inner endpoint: `GET /bank/bank_supplement_checklist`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string**| Bank card ID (otc_rds / the id returned by the list endpoint). | 

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
    bankId := "bankId_example" // string - Bank card ID (otc_rds / the id returned by the list endpoint).
    
    result, _, err := client.OTCApi.GetOtcBankSupplementChecklist(ctx, bankId)
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

[**OtcBankSupplementChecklistResponse**](OtcBankSupplementChecklistResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SubmitOtcBankPersonalSupplement

> OtcActionResponse SubmitOtcBankPersonalSupplement(ctx, bankId, optional)

Submit Bank Card Supplement Materials (Personal)

**Personal professional verification (type=1)** users submit non-same-person/supplementary materials. Must match `user_type=personal` from `GET /otc/bank/bank_supplement_checklist?bank_id=`; otherwise rejected.  Two submission methods (can be mixed):  1. **Pre-upload (recommended)**: call `POST /otc/upload/pre_upload` (`scene=bank`) to upload to the temporary bucket, then fill file items by category in the `relationship_proof` JSON; pass **`key` as plaintext** object path (`base64_decode(pre_upload.file_key)`, e.g. `otc_temp/{uid}/bank/xxx.png`), and `file_type` as plaintext MIME; the server base64-encodes before persistence—do not pass base64 `file_key` directly; 2. **Multipart direct upload**: one file field per material item; field names match checklist `code` (`id_document_front`, `id_document_back`, `address_proof`).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string**|  | 
**optional** | **SubmitOtcBankPersonalSupplementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubmitOtcBankPersonalSupplementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**idDocumentFront** | **optional.String**| ID document front-side file content (multipart file field, binary/Base64) | 
**idDocumentBack** | **optional.String**| ID document back-side file content (multipart file field, binary/Base64) | 
**addressProof** | **optional.String**| Proof-of-address file content (multipart file field, binary/Base64) | 
**relationshipProof** | **optional.String**| Optional. JSON string of relationship_proof. | 

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
    bankId := "bankId_example" // string - 
    
    result, _, err := client.OTCApi.SubmitOtcBankPersonalSupplement(ctx, bankId, nil)
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

[**OtcActionResponse**](OtcActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SubmitOtcBankEnterpriseSupplement

> OtcActionResponse SubmitOtcBankEnterpriseSupplement(ctx, bankId, optional)

Submit Bank Card Supplement Materials (Enterprise)

**Enterprise professional verification (type=2)** users submit supplementary materials. Must match `user_type=enterprise` from the checklist.  Two submission methods (can be mixed):  1. **Pre-upload (recommended)**: call `POST /otc/upload/pre_upload` (`scene=bank`), fill file items by category in `relationship_proof`; pass **`key` as plaintext** object path (`base64_decode(pre_upload.file_key)`), and `file_type` as plaintext MIME; 2. **Multipart direct upload**: file field names `certificate`, `share_holders`, `passport`, `share_holding_structure`; optional `funds_statement`, `additional`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string**|  | 
**optional** | **SubmitOtcBankEnterpriseSupplementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubmitOtcBankEnterpriseSupplementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uid** | **optional.String**|  | 
**certificate** | **optional.String**| Business license / registration certificate file content (multipart file field, binary/Base64) | 
**shareHolders** | **optional.String**| Register of shareholders file content (multipart file field, binary/Base64) | 
**passport** | **optional.String**| Legal representative / shareholder passport file content (multipart file field, binary/Base64) | 
**shareHoldingStructure** | **optional.String**| Ownership structure chart file content (multipart file field, binary/Base64) | 
**fundsStatement** | **optional.String**| Proof-of-funds file content (multipart file field, binary/Base64, optional) | 
**additional** | **optional.String**| Other supplementary material file content (multipart file field, binary/Base64, optional) | 
**relationshipProof** | **optional.String**| Optional. JSON string of relationship_proof. | 

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
    bankId := "bankId_example" // string - 
    
    result, _, err := client.OTCApi.SubmitOtcBankEnterpriseSupplement(ctx, bankId, nil)
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

[**OtcActionResponse**](OtcActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateOtcUploadPreUpload

> OtcUploadPreUploadResponse CreateOtcUploadPreUpload(ctx, otcUploadPreUploadRequest)

Pre-upload file (temporary bucket)

After selecting a file, the client calls this endpoint first to obtain a temporary-bucket POST Policy and `file_key`; then upload directly to S3 using the returned `url` and `fields` (success HTTP 204); finally, in business submit endpoints (e.g. `POST /otc/order/paid`, `POST /otc/bank/create`), pass the **same base64 `file_key` unchanged** (do not decode). The server validates ownership and object existence, then moves to the production bucket and persists. Unsubmitted files remain in the temporary bucket and are reclaimed by lifecycle rules.  Corresponds to Inner: `POST /upload/pre_upload`.  **`content_type` must be sent as base64** (plaintext containing `/` may be blocked by the gateway). Only the following MIME types are supported:  | MIME | base64 | Extension | | --- | --- | --- | | image/png | aW1hZ2UvcG5n | .png | | image/jpeg | aW1hZ2UvanBlZw== | .jpeg | | image/jpg | aW1hZ2UvanBn | .jpg | | application/pdf | YXBwbGljYXRpb24vcGRm | .pdf |  **`scene` mapping to downstream endpoints**:  | scene | Typical use | | --- | --- | | general | Fiat buy payment receipt (`payment_receipt_file_key` in `POST /otc/order/paid`) | | bank | Add card, bank card supplementary materials | | assessment | Professional verification materials | | credit | Credit limit increase materials |  **Credential validity**: response `expires_in` is **5400 seconds (90 minutes)**; `fields.Policy` `expiration` matches it. Complete the S3 direct upload within this window; after expiry, call this endpoint again for a new credential.  **File size**: the S3 POST Policy enforces `content-length-range` **1 byte ~ 10MB** (10485760 bytes). Uploads exceeding the limit are rejected by S3; all `scene` values share this limit.  **Direct S3 upload**: `url` is the upload address; send each key-value pair in `fields` unchanged as form-data; the `file` field must be last. Object path is generated as `otc_temp/{uid}/{scene}/{unique filename}`; uid is taken from the login session.  This endpoint returns `content type is required.` when `content_type` is missing. Ownership and object-existence checks for `file_key` are performed by the subsequent business submission endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otcUploadPreUploadRequest** | [**OtcUploadPreUploadRequest**](OtcUploadPreUploadRequest.md)|  | 

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
    otcUploadPreUploadRequest := gateapi.OtcUploadPreUploadRequest{} // OtcUploadPreUploadRequest - 
    
    result, _, err := client.OTCApi.CreateOtcUploadPreUpload(ctx, otcUploadPreUploadRequest)
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

[**OtcUploadPreUploadResponse**](OtcUploadPreUploadResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MarkOtcOrderPaid

> OtcActionResponse MarkOtcOrderPaid(ctx, otcMarkOrderPaidRequest)

Mark fiat order as paid (deposit confirmation)

Mark a fiat buy order as paid (deposit confirmation). **A user payment receipt must be uploaded**: `payment_receipt_file_key` is required; supported formats are jpg / jpeg / png / pdf, with a maximum size of 10 MB per file (validated jointly by the service and gateway). The compatible field name `payment_receipt` depends on the gateway and production contract. The persisted field is `otc_trade_record.payment_receipt_file_key`. The Pay Inner path is `POST .../pay/order_set_paid` (which commonly identifies orders by `client_order_id`); the Inner path corresponding to this OpenAPI operation, `POST /order/paid`, still primarily uses `order_id`. If the gateway standardizes on the merchant order ID, follow the gateway documentation.  **Recommended pre-upload flow**: first call `POST /otc/upload/pre_upload` (`scene=general`) and upload directly to the temporary bucket, then pass the returned **base64 `file_key` unchanged** (do not decode) to this endpoint. The service validates the uid and object existence before moving the object to the production bucket. A cross-user key returns `Invalid parameters file_key`; an object that has not been uploaded returns `Invalid parameters file not uploaded`. The legacy flow using a base64 key for an object uploaded directly to the production bucket remains supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otcMarkOrderPaidRequest** | [**OtcMarkOrderPaidRequest**](OtcMarkOrderPaidRequest.md)|  | 

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
    otcMarkOrderPaidRequest := gateapi.OtcMarkOrderPaidRequest{} // OtcMarkOrderPaidRequest - 
    
    result, _, err := client.OTCApi.MarkOtcOrderPaid(ctx, otcMarkOrderPaidRequest)
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

[**OtcActionResponse**](OtcActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelOtcOrder

> OtcActionResponse CancelOtcOrder(ctx, orderId)

Fiat order cancellation

Cancel fiat order

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
    orderId := "orderId_example" // string - Order ID
    
    result, _, err := client.OTCApi.CancelOtcOrder(ctx, orderId)
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

[**OtcActionResponse**](OtcActionResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOtcOrders

> OtcOrderListResponse ListOtcOrders(ctx, optional)

Fiat order list

Query the fiat order list with filters such as type, currency, time range, and status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListOtcOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOtcOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**type_** | **optional.String**| BUY for on-ramp, SELL for off-ramp | 
**fiatCurrency** | **optional.String**| Fiat currency | 
**cryptoCurrency** | **optional.String**| Digital currency | 
**startTime** | **optional.String**| starttime   for example : 2025-09-09 | 
**endTime** | **optional.String**| endtime  for example :2025-09-09 | 
**status** | **optional.String**| DONE: completed CANCEL: canceled PROCESSING: in progress DISBURSED: disbursed | 
**pn** | **optional.String**| Page number | 
**ps** | **optional.String**| Number of items per page | 

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
    
    result, _, err := client.OTCApi.ListOtcOrders(ctx, nil)
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

[**OtcOrderListResponse**](OtcOrderListResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListStableCoinOrders

> OtcStableCoinOrderListResponse ListStableCoinOrders(ctx, optional)

Stablecoin order list

Query stablecoin order list with filtering by currency, time range, status, etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListStableCoinOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListStableCoinOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pageSize** | **optional.String**| Number of records per page | 
**pageNumber** | **optional.String**| Page number | 
**coinName** | **optional.String**| ordercurrency | 
**startTime** | **optional.String**| Start Time | 
**endTime** | **optional.String**| End time | 
**status** | **optional.String**| Status: PROCESSING: in progress / DONE：completed / FAILED: failed | 

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
    
    result, _, err := client.OTCApi.ListStableCoinOrders(ctx, nil)
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

[**OtcStableCoinOrderListResponse**](OtcStableCoinOrderListResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOtcOrderDetail

> OtcOrderDetailResponse GetOtcOrderDetail(ctx, orderId)

Fiat order details

Query fiat order details

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
    orderId := "orderId_example" // string - Order ID
    
    result, _, err := client.OTCApi.GetOtcOrderDetail(ctx, orderId)
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

[**OtcOrderDetailResponse**](OtcOrderDetailResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
