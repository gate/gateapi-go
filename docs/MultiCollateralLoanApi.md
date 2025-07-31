# MultiCollateralLoanApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMultiCollateralOrders**](MultiCollateralLoanApi.md#ListMultiCollateralOrders) | **Get** /loan/multi_collateral/orders | Query multi-currency collateral order list
[**CreateMultiCollateral**](MultiCollateralLoanApi.md#CreateMultiCollateral) | **Post** /loan/multi_collateral/orders | Place multi-currency collateral order
[**GetMultiCollateralOrderDetail**](MultiCollateralLoanApi.md#GetMultiCollateralOrderDetail) | **Get** /loan/multi_collateral/orders/{order_id} | Query order details
[**ListMultiRepayRecords**](MultiCollateralLoanApi.md#ListMultiRepayRecords) | **Get** /loan/multi_collateral/repay | Query multi-currency collateral repayment records
[**RepayMultiCollateralLoan**](MultiCollateralLoanApi.md#RepayMultiCollateralLoan) | **Post** /loan/multi_collateral/repay | Multi-currency collateral repayment
[**ListMultiCollateralRecords**](MultiCollateralLoanApi.md#ListMultiCollateralRecords) | **Get** /loan/multi_collateral/mortgage | Query collateral adjustment records
[**OperateMultiCollateral**](MultiCollateralLoanApi.md#OperateMultiCollateral) | **Post** /loan/multi_collateral/mortgage | Add or withdraw collateral
[**ListUserCurrencyQuota**](MultiCollateralLoanApi.md#ListUserCurrencyQuota) | **Get** /loan/multi_collateral/currency_quota | Query user&#39;s collateral and borrowing currency quota information
[**ListMultiCollateralCurrencies**](MultiCollateralLoanApi.md#ListMultiCollateralCurrencies) | **Get** /loan/multi_collateral/currencies | Query supported borrowing and collateral currencies for multi-currency collateral
[**GetMultiCollateralLtv**](MultiCollateralLoanApi.md#GetMultiCollateralLtv) | **Get** /loan/multi_collateral/ltv | Query collateralization ratio information
[**GetMultiCollateralFixRate**](MultiCollateralLoanApi.md#GetMultiCollateralFixRate) | **Get** /loan/multi_collateral/fixed_rate | Query currency&#39;s 7-day and 30-day fixed interest rates
[**GetMultiCollateralCurrentRate**](MultiCollateralLoanApi.md#GetMultiCollateralCurrentRate) | **Get** /loan/multi_collateral/current_rate | Query currency&#39;s current interest rate


## ListMultiCollateralOrders

> []MultiCollateralOrder ListMultiCollateralOrders(ctx, optional)

Query multi-currency collateral order list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListMultiCollateralOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMultiCollateralOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 10]
**sort** | **optional.String**| Sort type: &#x60;time_desc&#x60; - Created time descending (default), &#x60;ltv_asc&#x60; - Collateral ratio ascending, &#x60;ltv_desc&#x60; - Collateral ratio descending. | 
**orderType** | **optional.String**| Order type: current - Query current orders, fixed - Query fixed orders, defaults to current orders if not specifiedOrder type: current - Query current orders, fixed - Query fixed orders, defaults to current orders if not specified | 

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
    
    result, _, err := client.MultiCollateralLoanApi.ListMultiCollateralOrders(ctx, nil)
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

[**[]MultiCollateralOrder**](MultiCollateralOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateMultiCollateral

> OrderResp CreateMultiCollateral(ctx, createMultiCollateralOrder)

Place multi-currency collateral order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createMultiCollateralOrder** | [**CreateMultiCollateralOrder**](CreateMultiCollateralOrder.md)|  | 

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
    createMultiCollateralOrder := gateapi.CreateMultiCollateralOrder{} // CreateMultiCollateralOrder - 
    
    result, _, err := client.MultiCollateralLoanApi.CreateMultiCollateral(ctx, createMultiCollateralOrder)
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

## GetMultiCollateralOrderDetail

> MultiCollateralOrder GetMultiCollateralOrderDetail(ctx, orderId)

Query order details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Order ID returned when order is successfully created | 

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
    orderId := "12345" // string - Order ID returned when order is successfully created
    
    result, _, err := client.MultiCollateralLoanApi.GetMultiCollateralOrderDetail(ctx, orderId)
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

[**MultiCollateralOrder**](MultiCollateralOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListMultiRepayRecords

> []MultiRepayRecord ListMultiRepayRecords(ctx, type_, optional)

Query multi-currency collateral repayment records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| Operation type: repay - Regular repayment, liquidate - Liquidation | 
**optional** | **ListMultiRepayRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMultiRepayRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**borrowCurrency** | **optional.String**| Borrowed currency | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 10]
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
    type_ := "repay" // string - Operation type: repay - Regular repayment, liquidate - Liquidation
    
    result, _, err := client.MultiCollateralLoanApi.ListMultiRepayRecords(ctx, type_, nil)
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

[**[]MultiRepayRecord**](MultiRepayRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RepayMultiCollateralLoan

> MultiRepayResp RepayMultiCollateralLoan(ctx, repayMultiLoan)

Multi-currency collateral repayment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repayMultiLoan** | [**RepayMultiLoan**](RepayMultiLoan.md)|  | 

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
    repayMultiLoan := gateapi.RepayMultiLoan{} // RepayMultiLoan - 
    
    result, _, err := client.MultiCollateralLoanApi.RepayMultiCollateralLoan(ctx, repayMultiLoan)
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

[**MultiRepayResp**](MultiRepayResp.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListMultiCollateralRecords

> []MultiCollateralRecord ListMultiCollateralRecords(ctx, optional)

Query collateral adjustment records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListMultiCollateralRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMultiCollateralRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 10]
**from** | **optional.Int64**| Start timestamp for the query | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
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
    
    result, _, err := client.MultiCollateralLoanApi.ListMultiCollateralRecords(ctx, nil)
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

[**[]MultiCollateralRecord**](MultiCollateralRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## OperateMultiCollateral

> CollateralAdjustRes OperateMultiCollateral(ctx, collateralAdjust)

Add or withdraw collateral

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collateralAdjust** | [**CollateralAdjust**](CollateralAdjust.md)|  | 

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
    collateralAdjust := gateapi.CollateralAdjust{} // CollateralAdjust - 
    
    result, _, err := client.MultiCollateralLoanApi.OperateMultiCollateral(ctx, collateralAdjust)
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

[**CollateralAdjustRes**](CollateralAdjustRes.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserCurrencyQuota

> []CurrencyQuota ListUserCurrencyQuota(ctx, type_, currency)

Query user's collateral and borrowing currency quota information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| Currency type: collateral - Collateral currency, borrow - Borrowing currency | 
**currency** | **string**| When it is a collateral currency, multiple currencies can be provided separated by commas; when it is a borrowing currency, only one currency can be provided. | 

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
    type_ := "collateral" // string - Currency type: collateral - Collateral currency, borrow - Borrowing currency
    currency := "BTC" // string - When it is a collateral currency, multiple currencies can be provided separated by commas; when it is a borrowing currency, only one currency can be provided.
    
    result, _, err := client.MultiCollateralLoanApi.ListUserCurrencyQuota(ctx, type_, currency)
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

[**[]CurrencyQuota**](CurrencyQuota.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListMultiCollateralCurrencies

> MultiCollateralCurrency ListMultiCollateralCurrencies(ctx, )

Query supported borrowing and collateral currencies for multi-currency collateral

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
    
    result, _, err := client.MultiCollateralLoanApi.ListMultiCollateralCurrencies(ctx)
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

[**MultiCollateralCurrency**](MultiCollateralCurrency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetMultiCollateralLtv

> CollateralLtv GetMultiCollateralLtv(ctx, )

Query collateralization ratio information

Multi-currency collateral ratio is fixed, independent of currency

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
    
    result, _, err := client.MultiCollateralLoanApi.GetMultiCollateralLtv(ctx)
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

[**CollateralLtv**](CollateralLtv.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetMultiCollateralFixRate

> []CollateralFixRate GetMultiCollateralFixRate(ctx, )

Query currency's 7-day and 30-day fixed interest rates

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
    
    result, _, err := client.MultiCollateralLoanApi.GetMultiCollateralFixRate(ctx)
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

[**[]CollateralFixRate**](CollateralFixRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetMultiCollateralCurrentRate

> []CollateralCurrentRate GetMultiCollateralCurrentRate(ctx, currencies, optional)

Query currency's current interest rate

Query currency's current interest rate for the previous hour, current interest rate updates hourly

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencies** | [**[]string**](string.md)| Specify currency name query array, separated by commas, maximum 100 items | 
**optional** | **GetMultiCollateralCurrentRateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMultiCollateralCurrentRateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**vipLevel** | **optional.String**| VIP level, defaults to 0 if not specified | [default to 0]

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
    currencies := []string{"[\"BTC\",\"GT\"]"} // []string - Specify currency name query array, separated by commas, maximum 100 items
    
    result, _, err := client.MultiCollateralLoanApi.GetMultiCollateralCurrentRate(ctx, currencies, nil)
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

[**[]CollateralCurrentRate**](CollateralCurrentRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
