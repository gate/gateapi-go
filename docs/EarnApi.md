# EarnApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDualInvestmentPlans**](EarnApi.md#ListDualInvestmentPlans) | **Get** /earn/dual/investment_plan | Dual Investment product list
[**ListDualOrders**](EarnApi.md#ListDualOrders) | **Get** /earn/dual/orders | Dual Investment order list
[**PlaceDualOrder**](EarnApi.md#PlaceDualOrder) | **Post** /earn/dual/orders | Place Dual Investment order
[**ListDualBalance**](EarnApi.md#ListDualBalance) | **Get** /earn/dual/balance | Dual-Currency Earning Assets
[**FindCoin**](EarnApi.md#FindCoin) | **Get** /earn/staking/coins | Staking coins
[**SwapStakingCoin**](EarnApi.md#SwapStakingCoin) | **Post** /earn/staking/swap | On-chain token swap for earned coins
[**OrderList**](EarnApi.md#OrderList) | **Get** /earn/staking/order_list | List of on-chain coin-earning orders
[**AwardList**](EarnApi.md#AwardList) | **Get** /earn/staking/award_list | On-chain coin-earning dividend records
[**AssetList**](EarnApi.md#AssetList) | **Get** /earn/staking/assets | On-chain coin-earning assets
[**ListEarnFixedTermProducts**](EarnApi.md#ListEarnFixedTermProducts) | **Get** /earn/fixed-term/product | Get product list
[**ListEarnFixedTermProductsByAsset**](EarnApi.md#ListEarnFixedTermProductsByAsset) | **Get** /earn/fixed-term/product/{asset}/list | Get product list by single currency
[**ListEarnFixedTermLends**](EarnApi.md#ListEarnFixedTermLends) | **Get** /earn/fixed-term/user/lend | Subscription list
[**CreateEarnFixedTermLend**](EarnApi.md#CreateEarnFixedTermLend) | **Post** /earn/fixed-term/user/lend | Subscription
[**CreateEarnFixedTermPreRedeem**](EarnApi.md#CreateEarnFixedTermPreRedeem) | **Post** /earn/fixed-term/user/pre-redeem | Redeem
[**ListEarnFixedTermHistory**](EarnApi.md#ListEarnFixedTermHistory) | **Get** /earn/fixed-term/user/history | Subscription history


## ListDualInvestmentPlans

> []DualGetPlans ListDualInvestmentPlans(ctx, optional)

Dual Investment product list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListDualInvestmentPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDualInvestmentPlansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**planId** | **optional.Int64**| Financial project ID | 

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
    
    result, _, err := client.EarnApi.ListDualInvestmentPlans(ctx, nil)
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

[**[]DualGetPlans**](DualGetPlans.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDualOrders

> []DualGetOrders ListDualOrders(ctx, optional)

Dual Investment order list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListDualOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDualOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**from** | **optional.Int64**| Start settlement time | 
**to** | **optional.Int64**| End settlement time | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]

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
    
    result, _, err := client.EarnApi.ListDualOrders(ctx, nil)
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

[**[]DualGetOrders**](DualGetOrders.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PlaceDualOrder

> PlaceDualInvestmentOrder PlaceDualOrder(ctx, placeDualInvestmentOrderParams)

Place Dual Investment order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeDualInvestmentOrderParams** | [**PlaceDualInvestmentOrderParams**](PlaceDualInvestmentOrderParams.md)|  | 

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
    placeDualInvestmentOrderParams := gateapi.PlaceDualInvestmentOrderParams{} // PlaceDualInvestmentOrderParams - 
    
    result, _, err := client.EarnApi.PlaceDualOrder(ctx, placeDualInvestmentOrderParams)
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

[**PlaceDualInvestmentOrder**](PlaceDualInvestmentOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDualBalance

> DualGetBalance ListDualBalance(ctx, )

Dual-Currency Earning Assets

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
    
    result, _, err := client.EarnApi.ListDualBalance(ctx)
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

[**DualGetBalance**](DualGetBalance.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FindCoin

> []map[string]interface{} FindCoin(ctx, findCoin)

Staking coins

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**findCoin** | [**FindCoin**](FindCoin.md)|  | 

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
    findCoin := gateapi.FindCoin{} // FindCoin - 
    
    result, _, err := client.EarnApi.FindCoin(ctx, findCoin)
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

**[]map[string]interface{}**

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SwapStakingCoin

> SwapCoinStruct SwapStakingCoin(ctx, swapCoin)

On-chain token swap for earned coins

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**swapCoin** | [**SwapCoin**](SwapCoin.md)|  | 

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
    swapCoin := gateapi.SwapCoin{} // SwapCoin - 
    
    result, _, err := client.EarnApi.SwapStakingCoin(ctx, swapCoin)
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

[**SwapCoinStruct**](SwapCoinStruct.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## OrderList

> OrderListStruct OrderList(ctx, optional)

List of on-chain coin-earning orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **OrderListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrderListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pid** | **optional.Int32**| Product ID | 
**coin** | **optional.String**| Currency name | 
**type_** | **optional.Int32**| Type 0-staking 1-redemption | 
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
    
    result, _, err := client.EarnApi.OrderList(ctx, nil)
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

[**OrderListStruct**](OrderListStruct.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AwardList

> AwardListStruct AwardList(ctx, optional)

On-chain coin-earning dividend records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **AwardListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AwardListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pid** | **optional.Int32**| Product ID | 
**coin** | **optional.String**| Currency name | 
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
    
    result, _, err := client.EarnApi.AwardList(ctx, nil)
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

[**AwardListStruct**](AwardListStruct.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AssetList

> []map[string]interface{} AssetList(ctx, optional)

On-chain coin-earning assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **AssetListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssetListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**coin** | **optional.String**| Currency name | 

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
    
    result, _, err := client.EarnApi.AssetList(ctx, nil)
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

**[]map[string]interface{}**

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEarnFixedTermProducts

> ListEarnFixedTermProductsResponse ListEarnFixedTermProducts(ctx, page, limit, optional)

Get product list

Query fixed-term earn product list. Supports filtering by currency, product type, status, etc. Returns product interest rate, lock-up period, quota, and reward campaign information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **int32**| Page number | 
**limit** | **int32**| Page size | 
**optional** | **ListEarnFixedTermProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEarnFixedTermProductsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**asset** | **optional.String**| Currency | 
**type_** | **optional.Int32**| Product type: 1 for regular, 2 for VIP | 

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
    page := 1 // int32 - Page number
    limit := 100 // int32 - Page size
    
    result, _, err := client.EarnApi.ListEarnFixedTermProducts(ctx, page, limit, nil)
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

[**ListEarnFixedTermProductsResponse**](ListEarnFixedTermProductsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEarnFixedTermProductsByAsset

> ListEarnFixedTermProductsByAssetResponse ListEarnFixedTermProductsByAsset(ctx, asset, optional)

Get product list by single currency

Sort by product term in ascending order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asset** | **string**| Currency name, e.g., USDT, BTC | 
**optional** | **ListEarnFixedTermProductsByAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEarnFixedTermProductsByAssetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**type_** | **optional.String**| Product type: \&quot;\&quot; or 1 for regular product list, 2 for VIP product list, 0 for all products | 

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
    asset := "USDT" // string - Currency name, e.g., USDT, BTC
    
    result, _, err := client.EarnApi.ListEarnFixedTermProductsByAsset(ctx, asset, nil)
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

[**ListEarnFixedTermProductsByAssetResponse**](ListEarnFixedTermProductsByAssetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEarnFixedTermLends

> ListEarnFixedTermLendsResponse ListEarnFixedTermLends(ctx, orderType, page, limit, optional)

Subscription list

Query the user's fixed-term earn subscription order list. Supports filtering by product, currency, order type, etc. Returns order details, earnings, rewards, and interest rate boost coupon information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderType** | **string**| Order type: 1 for current orders, 2 for historical orders | 
**page** | **int32**| Page number | 
**limit** | **int32**| Page size | 
**optional** | **ListEarnFixedTermLendsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEarnFixedTermLendsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**productId** | **optional.Int32**| Product ID | 
**orderId** | **optional.Int64**| Order ID | 
**asset** | **optional.String**| Currency | 
**subBusiness** | **optional.Int32**| Sub-business | 
**businessFilter** | **optional.String**| Business filter conditions, JSON array format, e.g., [{\&quot;business\&quot;:1, \&quot;sub_business\&quot;: 0}]. business: 1 for regular, 2 for VIP | 

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
    orderType := "1" // string - Order type: 1 for current orders, 2 for historical orders
    page := 1 // int32 - Page number
    limit := 10 // int32 - Page size
    
    result, _, err := client.EarnApi.ListEarnFixedTermLends(ctx, orderType, page, limit, nil)
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

[**ListEarnFixedTermLendsResponse**](ListEarnFixedTermLendsResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateEarnFixedTermLend

> CreateEarnFixedTermLendResponse CreateEarnFixedTermLend(ctx, optional)

Subscription

Subscribe to a fixed-term earn product by specifying the product ID and subscription amount. Optionally enable auto-renewal and apply an interest rate boost coupon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CreateEarnFixedTermLendOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateEarnFixedTermLendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fixedTermLendRequest** | [**optional.Interface of FixedTermLendRequest**](FixedTermLendRequest.md)|  | 

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
    
    result, _, err := client.EarnApi.CreateEarnFixedTermLend(ctx, nil)
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

[**CreateEarnFixedTermLendResponse**](CreateEarnFixedTermLendResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateEarnFixedTermPreRedeem

> CreateEarnFixedTermPreRedeemResponse CreateEarnFixedTermPreRedeem(ctx, optional)

Redeem

Early redemption of a fixed-term earn order, order ID is required

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CreateEarnFixedTermPreRedeemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateEarnFixedTermPreRedeemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**earnFixedTermPreRedeemRequest** | [**optional.Interface of EarnFixedTermPreRedeemRequest**](EarnFixedTermPreRedeemRequest.md)|  | 

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
    
    result, _, err := client.EarnApi.CreateEarnFixedTermPreRedeem(ctx, nil)
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

[**CreateEarnFixedTermPreRedeemResponse**](CreateEarnFixedTermPreRedeemResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEarnFixedTermHistory

> ListEarnFixedTermHistoryResponse ListEarnFixedTermHistory(ctx, type_, page, limit, optional)

Subscription history

Query the user's fixed-term earn history records. Supports filtering by type (subscription, redemption, interest, bonus rewards) and time range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| 1 for subscription, 2 for redemption, 3 for interest, 4 for bonus reward | 
**page** | **int32**| Page number | 
**limit** | **int32**| Page size | 
**optional** | **ListEarnFixedTermHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEarnFixedTermHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**productId** | **optional.Int32**| Product ID | 
**orderId** | **optional.String**| Order ID | 
**asset** | **optional.String**| Currency | 
**startAt** | **optional.Int32**| Start timestamp | 
**endAt** | **optional.Int32**| End Timestamp | 
**subBusiness** | **optional.Int32**| Sub-business | 
**businessFilter** | **optional.String**| Business filter conditions, JSON array format, e.g., [{\&quot;business\&quot;:1, \&quot;sub_business\&quot;: 0}]. business: 1 for regular, 2 for VIP | 

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
    type_ := "1" // string - 1 for subscription, 2 for redemption, 3 for interest, 4 for bonus reward
    page := 1 // int32 - Page number
    limit := 10 // int32 - Page size
    
    result, _, err := client.EarnApi.ListEarnFixedTermHistory(ctx, type_, page, limit, nil)
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

[**ListEarnFixedTermHistoryResponse**](ListEarnFixedTermHistoryResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
