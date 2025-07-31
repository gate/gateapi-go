# MarginApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMarginAccounts**](MarginApi.md#ListMarginAccounts) | **Get** /margin/accounts | Margin account list
[**ListMarginAccountBook**](MarginApi.md#ListMarginAccountBook) | **Get** /margin/account_book | Query margin account balance change history
[**ListFundingAccounts**](MarginApi.md#ListFundingAccounts) | **Get** /margin/funding_accounts | Funding account list
[**GetAutoRepayStatus**](MarginApi.md#GetAutoRepayStatus) | **Get** /margin/auto_repay | Query user auto repayment settings
[**SetAutoRepay**](MarginApi.md#SetAutoRepay) | **Post** /margin/auto_repay | Update user auto repayment settings
[**GetMarginTransferable**](MarginApi.md#GetMarginTransferable) | **Get** /margin/transferable | Get maximum transferable amount for isolated margin
[**GetUserMarginTier**](MarginApi.md#GetUserMarginTier) | **Get** /margin/user/loan_margin_tiers | Query user&#39;s own leverage lending tiers in current market
[**GetMarketMarginTier**](MarginApi.md#GetMarketMarginTier) | **Get** /margin/loan_margin_tiers | Query current market leverage lending tiers
[**SetUserMarketLeverage**](MarginApi.md#SetUserMarketLeverage) | **Post** /margin/leverage/user_market_setting | Set user market leverage multiplier
[**ListMarginUserAccount**](MarginApi.md#ListMarginUserAccount) | **Get** /margin/user/account | Query user&#39;s isolated margin account list
[**ListCrossMarginLoans**](MarginApi.md#ListCrossMarginLoans) | **Get** /margin/cross/loans | Query cross margin borrow history (deprecated)
[**ListCrossMarginRepayments**](MarginApi.md#ListCrossMarginRepayments) | **Get** /margin/cross/repayments | Retrieve cross margin repayments. (deprecated)


## ListMarginAccounts

> []MarginAccount ListMarginAccounts(ctx, optional)

Margin account list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListMarginAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMarginAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Currency pair | 

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
    
    result, _, err := client.MarginApi.ListMarginAccounts(ctx, nil)
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

[**[]MarginAccount**](MarginAccount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListMarginAccountBook

> []MarginAccountBook ListMarginAccountBook(ctx, optional)

Query margin account balance change history

Currently only provides transfer history to and from margin accounts. Query time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListMarginAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMarginAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Query history for specified currency. If &#x60;currency&#x60; is specified, &#x60;currency_pair&#x60; must also be specified. | 
**currencyPair** | **optional.String**| Specify margin account currency pair. Used in combination with &#x60;currency&#x60;. Ignored if &#x60;currency&#x60; is not specified | 
**type_** | **optional.String**| Query by specified account change type. If not specified, all change types will be included. | 
**from** | **optional.Int64**| Start timestamp for the query | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]

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
    
    result, _, err := client.MarginApi.ListMarginAccountBook(ctx, nil)
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

[**[]MarginAccountBook**](MarginAccountBook.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFundingAccounts

> []FundingAccount ListFundingAccounts(ctx, optional)

Funding account list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListFundingAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFundingAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Query by specified currency name | 

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
    
    result, _, err := client.MarginApi.ListFundingAccounts(ctx, nil)
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

[**[]FundingAccount**](FundingAccount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetAutoRepayStatus

> AutoRepaySetting GetAutoRepayStatus(ctx, )

Query user auto repayment settings

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
    
    result, _, err := client.MarginApi.GetAutoRepayStatus(ctx)
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

[**AutoRepaySetting**](AutoRepaySetting.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SetAutoRepay

> AutoRepaySetting SetAutoRepay(ctx, status)

Update user auto repayment settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**status** | **string**| Whether to enable auto repayment: &#x60;on&#x60; - enabled, &#x60;off&#x60; - disabled | 

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
    status := "on" // string - Whether to enable auto repayment: `on` - enabled, `off` - disabled
    
    result, _, err := client.MarginApi.SetAutoRepay(ctx, status)
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

[**AutoRepaySetting**](AutoRepaySetting.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetMarginTransferable

> MarginTransferable GetMarginTransferable(ctx, currency, optional)

Get maximum transferable amount for isolated margin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Query by specified currency name | 
**optional** | **GetMarginTransferableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMarginTransferableOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Currency pair | 

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
    currency := "BTC" // string - Query by specified currency name
    
    result, _, err := client.MarginApi.GetMarginTransferable(ctx, currency, nil)
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

[**MarginTransferable**](MarginTransferable.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserMarginTier

> []MarginLeverageTier GetUserMarginTier(ctx, currencyPair)

Query user's own leverage lending tiers in current market

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Currency pair | 

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
    currencyPair := "BTC_USDT" // string - Currency pair
    
    result, _, err := client.MarginApi.GetUserMarginTier(ctx, currencyPair)
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

[**[]MarginLeverageTier**](MarginLeverageTier.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetMarketMarginTier

> []MarginLeverageTier GetMarketMarginTier(ctx, currencyPair)

Query current market leverage lending tiers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Currency pair | 

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
    currencyPair := "BTC_USDT" // string - Currency pair
    
    result, _, err := client.MarginApi.GetMarketMarginTier(ctx, currencyPair)
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

[**[]MarginLeverageTier**](MarginLeverageTier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SetUserMarketLeverage

> SetUserMarketLeverage(ctx, marginMarketLeverage)

Set user market leverage multiplier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marginMarketLeverage** | [**MarginMarketLeverage**](MarginMarketLeverage.md)|  | 

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
    marginMarketLeverage := gateapi.MarginMarketLeverage{} // MarginMarketLeverage - 
    
    result, _, err := client.MarginApi.SetUserMarketLeverage(ctx, marginMarketLeverage)
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

## ListMarginUserAccount

> []MarginAccount ListMarginUserAccount(ctx, optional)

Query user's isolated margin account list

Supports querying risk ratio isolated accounts and margin ratio isolated accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListMarginUserAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMarginUserAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Currency pair | 

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
    
    result, _, err := client.MarginApi.ListMarginUserAccount(ctx, nil)
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

[**[]MarginAccount**](MarginAccount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossMarginLoans

> []CrossMarginLoan ListCrossMarginLoans(ctx, status, optional)

Query cross margin borrow history (deprecated)

Sorted by creation time in descending order by default. Set `reverse=false` for ascending order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**status** | **int32**| Filter by status. Supported values are 2 and 3. (deprecated.) | 
**optional** | **ListCrossMarginLoansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossMarginLoansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Query by specified currency, includes all currencies if not specified | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**reverse** | **optional.Bool**| Whether to sort in descending order, which is the default. Set &#x60;reverse&#x3D;false&#x60; to return ascending results | [default to true]

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
    status := 56 // int32 - Filter by status. Supported values are 2 and 3. (deprecated.)
    
    result, _, err := client.MarginApi.ListCrossMarginLoans(ctx, status, nil)
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

[**[]CrossMarginLoan**](CrossMarginLoan.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCrossMarginRepayments

> []CrossMarginRepayment ListCrossMarginRepayments(ctx, optional)

Retrieve cross margin repayments. (deprecated)

Sorted by creation time in descending order by default. Set `reverse=false` for ascending order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListCrossMarginRepaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCrossMarginRepaymentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**|  | 
**loanId** | **optional.String**|  | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**reverse** | **optional.Bool**| Whether to sort in descending order, which is the default. Set &#x60;reverse&#x3D;false&#x60; to return ascending results | [default to true]

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
    
    result, _, err := client.MarginApi.ListCrossMarginRepayments(ctx, nil)
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

[**[]CrossMarginRepayment**](CrossMarginRepayment.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
