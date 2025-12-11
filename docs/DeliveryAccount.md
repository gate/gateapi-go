# DeliveryAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **string** | Balance, only applicable to classic contract account.The balance is the sum of all historical fund flows, including historical transfers in and out, closing settlements, and transaction fee expenses, but does not include upl of positions.total &#x3D; SUM(history_dnw, history_pnl, history_fee, history_refr, history_fund) | [optional] 
**UnrealisedPnl** | **string** | Unrealized PNL | [optional] 
**PositionMargin** | **string** | Deprecated | [optional] 
**OrderMargin** | **string** | initial margin of all open orders | [optional] 
**Available** | **string** | Available amount for transfer or trading, which includes credit limits under the unified account (includes experience funds; experience funds cannot be transferred, so when transferring, the transfer amount must deduct experience funds) | [optional] 
**Point** | **string** | Point card amount | [optional] 
**Currency** | **string** | Settlement currency | [optional] 
**InDualMode** | **bool** | Whether Hedge Mode is enabled | [optional] 
**EnableCredit** | **bool** | Whether portfolio margin account mode is enabled | [optional] 
**PositionInitialMargin** | **string** | Initial margin occupied by positions, applicable to unified account mode | [optional] 
**MaintenanceMargin** | **string** | Maintenance margin occupied by positions, applicable to new classic account margin mode and unified account mode | [optional] 
**Bonus** | **string** | Bonus | [optional] 
**EnableEvolvedClassic** | **bool** | Deprecated | [optional] 
**CrossOrderMargin** | **string** | Cross margin order margin, applicable to new classic account margin mode | [optional] 
**CrossInitialMargin** | **string** | Cross margin initial margin, applicable to new classic account margin mode | [optional] 
**CrossMaintenanceMargin** | **string** | Cross margin maintenance margin, applicable to new classic account margin mode | [optional] 
**CrossUnrealisedPnl** | **string** | Cross margin unrealized P&amp;L, applicable to new classic account margin mode | [optional] 
**CrossAvailable** | **string** | Cross margin available balance, applicable to new classic account margin mode | [optional] 
**CrossMarginBalance** | **string** | Cross margin balance, applicable to new classic account margin mode | [optional] 
**CrossMmr** | **string** | Cross margin maintenance margin rate, applicable to new classic account margin mode | [optional] 
**CrossImr** | **string** | Cross margin initial margin rate, applicable to new classic account margin mode | [optional] 
**IsolatedPositionMargin** | **string** | Isolated position margin, applicable to new classic account margin mode | [optional] 
**EnableNewDualMode** | **bool** | Deprecated | [optional] 
**MarginMode** | **int32** | Margin mode of the account 0: classic future account or Classic Spot Margin Mode of unified account; 1:  Multi-Currency Margin Mode; 2:  Portoforlio Margin Mode; 3:  Single-Currency Margin Mode | [optional] 
**EnableTieredMm** | **bool** | Whether to enable tiered maintenance margin calculation | [optional] 
**History** | [**FuturesAccountHistory**](FuturesAccount_history.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


