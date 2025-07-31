# WalletApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCurrencyChains**](WalletApi.md#ListCurrencyChains) | **Get** /wallet/currency_chains | Query chains supported for specified currency
[**GetDepositAddress**](WalletApi.md#GetDepositAddress) | **Get** /wallet/deposit_address | Generate currency deposit address
[**ListWithdrawals**](WalletApi.md#ListWithdrawals) | **Get** /wallet/withdrawals | Get withdrawal records
[**ListDeposits**](WalletApi.md#ListDeposits) | **Get** /wallet/deposits | Get deposit records
[**Transfer**](WalletApi.md#Transfer) | **Post** /wallet/transfers | Transfer between trading accounts
[**ListSubAccountTransfers**](WalletApi.md#ListSubAccountTransfers) | **Get** /wallet/sub_account_transfers | Get transfer records between main and sub accounts
[**TransferWithSubAccount**](WalletApi.md#TransferWithSubAccount) | **Post** /wallet/sub_account_transfers | Transfer between main and sub accounts
[**SubAccountToSubAccount**](WalletApi.md#SubAccountToSubAccount) | **Post** /wallet/sub_account_to_sub_account | Transfer between sub-accounts
[**GetTransferOrderStatus**](WalletApi.md#GetTransferOrderStatus) | **Get** /wallet/order_status | Transfer status query
[**ListWithdrawStatus**](WalletApi.md#ListWithdrawStatus) | **Get** /wallet/withdraw_status | Query withdrawal status
[**ListSubAccountBalances**](WalletApi.md#ListSubAccountBalances) | **Get** /wallet/sub_account_balances | Query sub-account balance information
[**ListSubAccountMarginBalances**](WalletApi.md#ListSubAccountMarginBalances) | **Get** /wallet/sub_account_margin_balances | Query sub-account isolated margin account balance information
[**ListSubAccountFuturesBalances**](WalletApi.md#ListSubAccountFuturesBalances) | **Get** /wallet/sub_account_futures_balances | Query sub-account perpetual futures account balance information
[**ListSubAccountCrossMarginBalances**](WalletApi.md#ListSubAccountCrossMarginBalances) | **Get** /wallet/sub_account_cross_margin_balances | Query sub-account cross margin account balance information
[**ListSavedAddress**](WalletApi.md#ListSavedAddress) | **Get** /wallet/saved_address | Query withdrawal address whitelist
[**GetTradeFee**](WalletApi.md#GetTradeFee) | **Get** /wallet/fee | Query personal trading fees
[**GetTotalBalance**](WalletApi.md#GetTotalBalance) | **Get** /wallet/total_balance | Query personal account totals
[**ListSmallBalance**](WalletApi.md#ListSmallBalance) | **Get** /wallet/small_balance | Get list of convertible small balance currencies
[**ConvertSmallBalance**](WalletApi.md#ConvertSmallBalance) | **Post** /wallet/small_balance | Convert small balance currency
[**ListSmallBalanceHistory**](WalletApi.md#ListSmallBalanceHistory) | **Get** /wallet/small_balance_history | Get convertible small balance currency history
[**ListPushOrders**](WalletApi.md#ListPushOrders) | **Get** /wallet/push | Get UID transfer history


## ListCurrencyChains

> []CurrencyChain ListCurrencyChains(ctx, currency)

Query chains supported for specified currency

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Currency name | 

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
    currency := "GT" // string - Currency name
    
    result, _, err := client.WalletApi.ListCurrencyChains(ctx, currency)
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

[**[]CurrencyChain**](CurrencyChain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetDepositAddress

> DepositAddress GetDepositAddress(ctx, currency)

Generate currency deposit address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Currency name | 

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
    currency := "USDT" // string - Currency name
    
    result, _, err := client.WalletApi.GetDepositAddress(ctx, currency)
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

[**DepositAddress**](DepositAddress.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListWithdrawals

> []WithdrawalRecord ListWithdrawals(ctx, optional)

Get withdrawal records

Record query time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListWithdrawalsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWithdrawalsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Specify the currency. If not specified, returns all currencies | 
**withdrawId** | **optional.String**| Withdrawal record ID starts with &#39;w&#39;, such as: w1879219868. When withdraw_id is not empty, only this specific withdrawal record will be queried, and time-based querying will be disabled | 
**assetClass** | **optional.String**| Currency type of withdrawal record, empty by default. Supports querying withdrawal records in main zone and innovation zone on demand. Value range: SPOT, PILOT  SPOT: Main Zone PILOT: Innovation Zone | 
**withdrawOrderId** | **optional.String**| User-defined order number for withdrawal. Default is empty. When not empty, the specified user-defined order number record will be queried | 
**from** | **optional.Int64**| Start time for querying records, defaults to 7 days before current time if not specified | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

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
    
    result, _, err := client.WalletApi.ListWithdrawals(ctx, nil)
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

[**[]WithdrawalRecord**](WithdrawalRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeposits

> []DepositRecord ListDeposits(ctx, optional)

Get deposit records

Record query time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListDepositsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDepositsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Specify the currency. If not specified, returns all currencies | 
**from** | **optional.Int64**| Start time for querying records, defaults to 7 days before current time if not specified | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**limit** | **optional.Int32**| Maximum number of entries returned in the list, limited to 500 transactions | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

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
    
    result, _, err := client.WalletApi.ListDeposits(ctx, nil)
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

[**[]DepositRecord**](DepositRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## Transfer

> TransactionId Transfer(ctx, transfer)

Transfer between trading accounts

Balance transfers between personal trading accounts. Currently supports the following transfer operations:  1. Spot account - Margin account 2. Spot account - Perpetual futures account 3. Spot account - Delivery futures account 4. Spot account - Options account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transfer** | [**Transfer**](Transfer.md)|  | 

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
    transfer := gateapi.Transfer{} // Transfer - 
    
    result, _, err := client.WalletApi.Transfer(ctx, transfer)
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

[**TransactionId**](TransactionID.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSubAccountTransfers

> []SubAccountTransferRecordItem ListSubAccountTransfers(ctx, optional)

Get transfer records between main and sub accounts

Record query time range cannot exceed 30 days  > Note: Only records after 2020-04-10 can be retrieved

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSubAccountTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubAccountTransfersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**subUid** | **optional.String**| Sub-account user ID, you can query multiple records separated by &#x60;,&#x60;. If not specified, it will return records of all sub-accounts | 
**from** | **optional.Int64**| Start time for querying records, defaults to 7 days before current time if not specified | 
**to** | **optional.Int64**| End timestamp for the query, defaults to current time if not specified | 
**limit** | **optional.Int32**| Maximum number of records returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

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
    
    result, _, err := client.WalletApi.ListSubAccountTransfers(ctx, nil)
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

[**[]SubAccountTransferRecordItem**](SubAccountTransferRecordItem.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## TransferWithSubAccount

> TransactionId TransferWithSubAccount(ctx, subAccountTransfer)

Transfer between main and sub accounts

Supports transfers to/from sub-account's spot or futures accounts. Note that regardless of which sub-account is operated, only the main account's spot account is used

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subAccountTransfer** | [**SubAccountTransfer**](SubAccountTransfer.md)|  | 

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
    subAccountTransfer := gateapi.SubAccountTransfer{} // SubAccountTransfer - 
    
    result, _, err := client.WalletApi.TransferWithSubAccount(ctx, subAccountTransfer)
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

[**TransactionId**](TransactionID.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SubAccountToSubAccount

> TransactionId SubAccountToSubAccount(ctx, subAccountToSubAccount)

Transfer between sub-accounts

Supports balance transfers between two sub-accounts under the same main account. You can use either the main account's API Key or the source sub-account's API Key to perform the operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subAccountToSubAccount** | [**SubAccountToSubAccount**](SubAccountToSubAccount.md)|  | 

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
    subAccountToSubAccount := gateapi.SubAccountToSubAccount{} // SubAccountToSubAccount - 
    
    result, _, err := client.WalletApi.SubAccountToSubAccount(ctx, subAccountToSubAccount)
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

[**TransactionId**](TransactionID.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetTransferOrderStatus

> TransferOrderStatus GetTransferOrderStatus(ctx, optional)

Transfer status query

Supports querying transfer status based on user-defined client_order_id or tx_id returned by the transfer interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetTransferOrderStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTransferOrderStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**clientOrderId** | **optional.String**| Customer-defined ID to prevent duplicate transfers. Can be a combination of letters (case-sensitive), numbers, hyphens &#39;-&#39;, and underscores &#39;_&#39;. Can be pure letters or pure numbers with length between 1-64 characters | 
**txId** | **optional.String**| Transfer operation number, cannot be empty at the same time as client_order_id | 

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
    
    result, _, err := client.WalletApi.GetTransferOrderStatus(ctx, nil)
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

[**TransferOrderStatus**](TransferOrderStatus.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListWithdrawStatus

> []WithdrawStatus ListWithdrawStatus(ctx, optional)

Query withdrawal status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListWithdrawStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWithdrawStatusOpts struct

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
    
    result, _, err := client.WalletApi.ListWithdrawStatus(ctx, nil)
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

[**[]WithdrawStatus**](WithdrawStatus.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSubAccountBalances

> []SubAccountBalance ListSubAccountBalances(ctx, optional)

Query sub-account balance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSubAccountBalancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubAccountBalancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**subUid** | **optional.String**| Sub-account user ID, you can query multiple records separated by &#x60;,&#x60;. If not specified, it will return records of all sub-accounts | 

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
    
    result, _, err := client.WalletApi.ListSubAccountBalances(ctx, nil)
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

[**[]SubAccountBalance**](SubAccountBalance.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSubAccountMarginBalances

> []SubAccountMarginBalance ListSubAccountMarginBalances(ctx, optional)

Query sub-account isolated margin account balance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSubAccountMarginBalancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubAccountMarginBalancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**subUid** | **optional.String**| Sub-account user ID, you can query multiple records separated by &#x60;,&#x60;. If not specified, it will return records of all sub-accounts | 

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
    
    result, _, err := client.WalletApi.ListSubAccountMarginBalances(ctx, nil)
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

[**[]SubAccountMarginBalance**](SubAccountMarginBalance.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSubAccountFuturesBalances

> []SubAccountFuturesBalance ListSubAccountFuturesBalances(ctx, optional)

Query sub-account perpetual futures account balance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSubAccountFuturesBalancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubAccountFuturesBalancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**subUid** | **optional.String**| Sub-account user ID, you can query multiple records separated by &#x60;,&#x60;. If not specified, it will return records of all sub-accounts | 
**settle** | **optional.String**| Query balance of specified settlement currency | 

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
    
    result, _, err := client.WalletApi.ListSubAccountFuturesBalances(ctx, nil)
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

[**[]SubAccountFuturesBalance**](SubAccountFuturesBalance.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSubAccountCrossMarginBalances

> []SubAccountCrossMarginBalance ListSubAccountCrossMarginBalances(ctx, optional)

Query sub-account cross margin account balance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSubAccountCrossMarginBalancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubAccountCrossMarginBalancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**subUid** | **optional.String**| Sub-account user ID, you can query multiple records separated by &#x60;,&#x60;. If not specified, it will return records of all sub-accounts | 

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
    
    result, _, err := client.WalletApi.ListSubAccountCrossMarginBalances(ctx, nil)
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

[**[]SubAccountCrossMarginBalance**](SubAccountCrossMarginBalance.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSavedAddress

> []SavedAddress ListSavedAddress(ctx, currency, optional)

Query withdrawal address whitelist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Currency | 
**optional** | **ListSavedAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSavedAddressOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**chain** | **optional.String**| Chain name | [default to ]
**limit** | **optional.String**| Maximum number returned, up to 100 | [default to 50]
**page** | **optional.Int32**| Page number | [default to 1]

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
    currency := "USDT" // string - Currency
    
    result, _, err := client.WalletApi.ListSavedAddress(ctx, currency, nil)
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

[**[]SavedAddress**](SavedAddress.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetTradeFee

> TradeFee GetTradeFee(ctx, optional)

Query personal trading fees

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetTradeFeeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTradeFeeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Specify currency pair to get more accurate fee settings.  This field is optional. Usually fee settings are the same for all currency pairs. | 
**settle** | **optional.String**| Specify the settlement currency of the contract to get more accurate fee settings.  This field is optional. Generally, the fee settings for all settlement currencies are the same. | 

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
    
    result, _, err := client.WalletApi.GetTradeFee(ctx, nil)
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

[**TradeFee**](TradeFee.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetTotalBalance

> TotalBalance GetTotalBalance(ctx, optional)

Query personal account totals

This query endpoint returns the total *estimated value* of all currencies in each account converted to the input currency. Exchange rates and related account balance information may be cached for up to 1 minute. It is not recommended to use this interface data for real-time calculations.  For real-time calculations, query the corresponding balance interface based on account type, such as:  - `GET /spot/accounts` to query spot account - `GET /margin/accounts` to query margin account - `GET /futures/{settle}/accounts` to query futures account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetTotalBalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTotalBalanceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Target currency type for statistical conversion. Accepts BTC, CNY, USD, and USDT. USDT is the default value | [default to USDT]

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
    
    result, _, err := client.WalletApi.GetTotalBalance(ctx, nil)
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

[**TotalBalance**](TotalBalance.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSmallBalance

> []SmallBalance ListSmallBalance(ctx, )

Get list of convertible small balance currencies

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
    
    result, _, err := client.WalletApi.ListSmallBalance(ctx)
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

[**[]SmallBalance**](SmallBalance.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ConvertSmallBalance

> ConvertSmallBalance(ctx, convertSmallBalance)

Convert small balance currency

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**convertSmallBalance** | [**ConvertSmallBalance**](ConvertSmallBalance.md)|  | 

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
    convertSmallBalance := gateapi.ConvertSmallBalance{} // ConvertSmallBalance - 
    
    result, _, err := client.WalletApi.ConvertSmallBalance(ctx, convertSmallBalance)
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

## ListSmallBalanceHistory

> []SmallBalanceHistory ListSmallBalanceHistory(ctx, optional)

Get convertible small balance currency history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSmallBalanceHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSmallBalanceHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Currency to convert | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of items returned. Default: 100, minimum: 1, maximum: 100 | [default to 100]

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
    
    result, _, err := client.WalletApi.ListSmallBalanceHistory(ctx, nil)
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

[**[]SmallBalanceHistory**](SmallBalanceHistory.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListPushOrders

> []UidPushOrder ListPushOrders(ctx, optional)

Get UID transfer history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListPushOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPushOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **optional.Int32**| Order ID | 
**from** | **optional.Int32**| Start time for querying records. If not specified, defaults to 7 days before the current time. Unix timestamp in seconds | 
**to** | **optional.Int32**| End time for querying records. If not specified, defaults to the current time. Unix timestamp in seconds | 
**limit** | **optional.Int32**| Maximum number of items returned in the list, default value is 100 | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
**transactionType** | **optional.String**| Order type returned in the list: &#x60;withdraw&#x60;, &#x60;deposit&#x60;. Default is &#x60;withdraw&#x60;. | [default to withdraw]

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
    
    result, _, err := client.WalletApi.ListPushOrders(ctx, nil)
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

[**[]UidPushOrder**](UidPushOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
