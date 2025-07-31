# CollateralLoanApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCollateralLoanOrders**](CollateralLoanApi.md#ListCollateralLoanOrders) | **Get** /loan/collateral/orders | Query collateral loan order list
[**CreateCollateralLoan**](CollateralLoanApi.md#CreateCollateralLoan) | **Post** /loan/collateral/orders | Place collateral loan order
[**GetCollateralLoanOrderDetail**](CollateralLoanApi.md#GetCollateralLoanOrderDetail) | **Get** /loan/collateral/orders/{order_id} | Query single order details
[**RepayCollateralLoan**](CollateralLoanApi.md#RepayCollateralLoan) | **Post** /loan/collateral/repay | Collateral loan repayment
[**ListRepayRecords**](CollateralLoanApi.md#ListRepayRecords) | **Get** /loan/collateral/repay_records | Query collateral loan repayment records
[**ListCollateralRecords**](CollateralLoanApi.md#ListCollateralRecords) | **Get** /loan/collateral/collaterals | Query collateral adjustment records
[**OperateCollateral**](CollateralLoanApi.md#OperateCollateral) | **Post** /loan/collateral/collaterals | Increase or redeem collateral
[**GetUserTotalAmount**](CollateralLoanApi.md#GetUserTotalAmount) | **Get** /loan/collateral/total_amount | Query user&#39;s total borrowing and collateral amount
[**GetUserLtvInfo**](CollateralLoanApi.md#GetUserLtvInfo) | **Get** /loan/collateral/ltv | Query user&#39;s collateralization ratio and remaining borrowable currencies
[**ListCollateralCurrencies**](CollateralLoanApi.md#ListCollateralCurrencies) | **Get** /loan/collateral/currencies | Query supported borrowing and collateral currencies


## ListCollateralLoanOrders

> []CollateralOrder ListCollateralLoanOrders(ctx, optional)

Query collateral loan order list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCollateralLoanOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCollateralLoanOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**collateralCurrency** | **optional.String**| Collateral currency | 
**borrowCurrency** | **optional.String**| Borrowed currency | 

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
    
    result, _, err := client.CollateralLoanApi.ListCollateralLoanOrders(ctx, nil)
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

[**[]CollateralOrder**](CollateralOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateCollateralLoan

> OrderResp CreateCollateralLoan(ctx, createCollateralOrder)

Place collateral loan order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createCollateralOrder** | [**CreateCollateralOrder**](CreateCollateralOrder.md)|  | 

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
    createCollateralOrder := gateapi.CreateCollateralOrder{} // CreateCollateralOrder - 
    
    result, _, err := client.CollateralLoanApi.CreateCollateralLoan(ctx, createCollateralOrder)
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

[**OrderResp**](OrderResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCollateralLoanOrderDetail

> CollateralOrder GetCollateralLoanOrderDetail(ctx, orderId)

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
    orderId := 100001 // int64 - Order ID returned when order is successfully created
    
    result, _, err := client.CollateralLoanApi.GetCollateralLoanOrderDetail(ctx, orderId)
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

[**CollateralOrder**](CollateralOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RepayCollateralLoan

> RepayResp RepayCollateralLoan(ctx, repayLoan)

Collateral loan repayment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repayLoan** | [**RepayLoan**](RepayLoan.md)|  | 

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
    repayLoan := gateapi.RepayLoan{} // RepayLoan - 
    
    result, _, err := client.CollateralLoanApi.RepayCollateralLoan(ctx, repayLoan)
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

[**RepayResp**](RepayResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListRepayRecords

> []RepayRecord ListRepayRecords(ctx, source, optional)

Query collateral loan repayment records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string**| Operation type: repay - Regular repayment, liquidate - Liquidation | 
**optional** | **ListRepayRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRepayRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**borrowCurrency** | **optional.String**| Borrowed currency | 
**collateralCurrency** | **optional.String**| Collateral currency | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
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
    source := "repay" // string - Operation type: repay - Regular repayment, liquidate - Liquidation
    
    result, _, err := client.CollateralLoanApi.ListRepayRecords(ctx, source, nil)
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

[**[]RepayRecord**](RepayRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCollateralRecords

> []CollateralRecord ListCollateralRecords(ctx, optional)

Query collateral adjustment records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCollateralRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCollateralRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**from** | **optional.Int64**| Start timestamp for the query | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**borrowCurrency** | **optional.String**| Borrowed currency | 
**collateralCurrency** | **optional.String**| Collateral currency | 

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
    
    result, _, err := client.CollateralLoanApi.ListCollateralRecords(ctx, nil)
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

[**[]CollateralRecord**](CollateralRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## OperateCollateral

> OperateCollateral(ctx, collateralAlign)

Increase or redeem collateral

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collateralAlign** | [**CollateralAlign**](CollateralAlign.md)|  | 

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
    collateralAlign := gateapi.CollateralAlign{} // CollateralAlign - 
    
    result, _, err := client.CollateralLoanApi.OperateCollateral(ctx, collateralAlign)
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

 (empty response body)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserTotalAmount

> UserTotalAmount GetUserTotalAmount(ctx, )

Query user's total borrowing and collateral amount

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
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    
    result, _, err := client.CollateralLoanApi.GetUserTotalAmount(ctx)
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

[**UserTotalAmount**](UserTotalAmount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserLtvInfo

> UserLtvInfo GetUserLtvInfo(ctx, collateralCurrency, borrowCurrency)

Query user's collateralization ratio and remaining borrowable currencies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collateralCurrency** | **string**| Collateral currency | 
**borrowCurrency** | **string**| Borrowed currency | 

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
    collateralCurrency := "BTC" // string - Collateral currency
    borrowCurrency := "USDT" // string - Borrowed currency
    
    result, _, err := client.CollateralLoanApi.GetUserLtvInfo(ctx, collateralCurrency, borrowCurrency)
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

[**UserLtvInfo**](UserLtvInfo.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCollateralCurrencies

> []CollateralLoanCurrency ListCollateralCurrencies(ctx, optional)

Query supported borrowing and collateral currencies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCollateralCurrenciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCollateralCurrenciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**loanCurrency** | **optional.String**| Parameter loan_currency. If omitted, returns all supported borrowing currencies; if provided, returns the array of collateral currencies supported for that borrowing currency | 

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
    
    result, _, err := client.CollateralLoanApi.ListCollateralCurrencies(ctx, nil)
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

[**[]CollateralLoanCurrency**](CollateralLoanCurrency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
