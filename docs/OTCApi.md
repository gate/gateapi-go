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

Create stablecoin order

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

Retrieve the user's bank card list, used to select a bank card when placing an order. **Default card**: refer to the list item field `is_default` (1=default); there is no need to call the deprecated standalone \"default bank card\" endpoint. Corresponding Inner: `GET /bank_list` or `GET /bank/list`.

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

> OtcBankCreateResponse CreateOtcBank(ctx, bankAccountName, bankName, bankCountry, bankAddress, iban, swift, documentationFile, optional)

Create bank card

Bind a bank card. Under the Global entity, an account with a non-matching name may enter manual review (`status` pending) and require subsequent supplementary materials. Corresponding Inner: `POST /bank/create`. Fields and protocol are subject to the production form/gateway; in some environments `bank_account_name` is passed Base64-encoded, see the integration notes for details.

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
**documentationFile** | ***os.File*****os.File**| Account-opening proof file (jpg/jpeg/png/pdf, etc.; single file ≤4MB — subject to production environment). | 
**optional** | **CreateOtcBankOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOtcBankOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**remittanceLineNumber** | **optional.String**|  | 
**agentBankName** | **optional.String**|  | 
**agentBankSwift** | **optional.String**|  | 

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
    documentationFile := gateapi.*os.File{} // *os.File - Account-opening proof file (jpg/jpeg/png/pdf, etc.; single file ≤4MB — subject to production environment).
    
    result, _, err := client.OTCApi.CreateOtcBank(ctx, bankAccountName, bankName, bankCountry, bankAddress, iban, swift, documentationFile, nil)
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

**①** `bank_id` must be specified: after verifying that the card belongs to the current user and its status allows supplementation, returns the items to be supplemented and whether each sub-item is required, based on the user's **passed professional verification type** (personal/enterprise). Corresponding Inner: `GET /bank/bank_supplement_checklist`.

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

> OtcActionResponse SubmitOtcBankPersonalSupplement(ctx, bankId, idDocumentFront, idDocumentBack, addressProof)

Submit Bank Card Supplement Materials (Personal)

**Personal professional verification (type=1)** users submit non-same-person/supplementary materials. Must match `user_type=personal` returned by `GET /otc/bank/bank_supplement_checklist?bank_id=`, otherwise the request is rejected. **multipart/form-data** is recommended: each material item is a separate file field, with field names matching the checklist `code` (`id_document_front`, `id_document_back`, `address_proof`).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string**|  | 
**idDocumentFront** | **string**| ID document front-side file content (multipart file field, binary/Base64) | 
**idDocumentBack** | **string**| ID document back-side file content (multipart file field, binary/Base64) | 
**addressProof** | **string**| Proof-of-address file content (multipart file field, binary/Base64) | 

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
    idDocumentFront := "idDocumentFront_example" // string - ID document front-side file content (multipart file field, binary/Base64)
    idDocumentBack := "idDocumentBack_example" // string - ID document back-side file content (multipart file field, binary/Base64)
    addressProof := "addressProof_example" // string - Proof-of-address file content (multipart file field, binary/Base64)
    
    result, _, err := client.OTCApi.SubmitOtcBankPersonalSupplement(ctx, bankId, idDocumentFront, idDocumentBack, addressProof)
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

> OtcActionResponse SubmitOtcBankEnterpriseSupplement(ctx, bankId, certificate, shareHolders, passport, shareHoldingStructure, optional)

Submit Bank Card Supplement Materials (Enterprise)

**Enterprise professional verification (type=2)** users submit supplementary materials. Must match `user_type=enterprise` returned by the checklist. **multipart** file field names: `certificate`, `share_holders`, `passport`, `share_holding_structure`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string**|  | 
**certificate** | **string**| Business license / registration certificate file content (multipart file field, binary/Base64) | 
**shareHolders** | **string**| Register of shareholders file content (multipart file field, binary/Base64) | 
**passport** | **string**| Legal representative / shareholder passport file content (multipart file field, binary/Base64) | 
**shareHoldingStructure** | **string**| Ownership structure chart file content (multipart file field, binary/Base64) | 
**optional** | **SubmitOtcBankEnterpriseSupplementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubmitOtcBankEnterpriseSupplementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uid** | **optional.String**|  | 
**fundsStatement** | **optional.String**| Proof-of-funds file content (multipart file field, binary/Base64, optional) | 
**additional** | **optional.String**| Other supplementary material file content (multipart file field, binary/Base64, optional) | 

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
    certificate := "certificate_example" // string - Business license / registration certificate file content (multipart file field, binary/Base64)
    shareHolders := "shareHolders_example" // string - Register of shareholders file content (multipart file field, binary/Base64)
    passport := "passport_example" // string - Legal representative / shareholder passport file content (multipart file field, binary/Base64)
    shareHoldingStructure := "shareHoldingStructure_example" // string - Ownership structure chart file content (multipart file field, binary/Base64)
    
    result, _, err := client.OTCApi.SubmitOtcBankEnterpriseSupplement(ctx, bankId, certificate, shareHolders, passport, shareHoldingStructure, nil)
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

## MarkOtcOrderPaid

> OtcActionResponse MarkOtcOrderPaid(ctx, otcMarkOrderPaidRequest)

Mark fiat order as paid (deposit confirmation)

Mark a fiat buy order as paid (deposit confirmation). **The user's payment receipt must be uploaded**: `payment_receipt_file_key` is required; file format jpg / jpeg / png / pdf, single file no larger than 4MB (jointly validated by the server and gateway). The compatible field name `payment_receipt` is subject to the gateway/production environment. For the persisted field, see `otc_trade_record.payment_receipt_file_key`. The Pay Inner path is `POST .../pay/order_set_paid` (orders are usually associated via `client_order_id`); this OpenAPI path maps to Inner `POST /order/paid` and still uses `order_id` as the primary key—if the gateway unifies it to the merchant order number, the gateway documentation prevails.

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
**status** | **optional.String**| DONE: Completed CANCEL: Canceled PROCESSING: In Progress | 
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
