# FuturesAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **string** | total is the balance after the user&#39;s accumulated deposit, withdraw, profit and loss (including realized profit and loss, fund, fee and referral rebate), excluding unrealized profit and loss.  total &#x3D; SUM(history_dnw, history_pnl, history_fee, history_refr, history_fund) | [optional] 
**UnrealisedPnl** | **string** | Unrealized PNL | [optional] 
**PositionMargin** | **string** | Position margin | [optional] 
**OrderMargin** | **string** | Order margin of unfinished orders | [optional] 
**Available** | **string** | Refers to the available withdrawal or trading amount in per-position, specifically the per-position available balance under the unified account that includes the credit line (which incorporates trial funds; since trial funds cannot be withdrawn, the actual withdrawal amount needs to deduct the trial fund portion when processing withdrawals) | [optional] 
**Point** | **string** | Point card amount | [optional] 
**Currency** | **string** | Settlement currency | [optional] 
**InDualMode** | **bool** | Whether Hedge Mode is enabled | [optional] 
**PositionMode** | **string** | Position mode: single - one-way, dual - dual-side, split - sub-positions (in_dual_mode is deprecated) | [optional] 
**EnableCredit** | **bool** | Whether portfolio margin account mode is enabled | [optional] 
**PositionInitialMargin** | **string** | Initial margin occupied by positions, applicable to unified account mode | [optional] 
**MaintenanceMargin** | **string** | Maintenance margin occupied by positions, applicable to new classic account margin mode and unified account mode | [optional] 
**Bonus** | **string** | Bonus | [optional] 
**EnableEvolvedClassic** | **bool** | Classic account margin mode, true-new mode, false-old mode | [optional] 
**CrossOrderMargin** | **string** | Cross margin order margin, applicable to new classic account margin mode | [optional] 
**CrossInitialMargin** | **string** | Cross margin initial margin, applicable to new classic account margin mode | [optional] 
**CrossMaintenanceMargin** | **string** | Cross margin maintenance margin, applicable to new classic account margin mode | [optional] 
**CrossUnrealisedPnl** | **string** | Cross margin unrealized P&amp;L, applicable to new classic account margin mode | [optional] 
**CrossAvailable** | **string** | Cross margin available balance, applicable to new classic account margin mode | [optional] 
**CrossMarginBalance** | **string** | Cross margin balance, applicable to new classic account margin mode | [optional] 
**CrossMmr** | **string** | Cross margin maintenance margin rate, applicable to new classic account margin mode | [optional] 
**CrossImr** | **string** | Cross margin initial margin rate, applicable to new classic account margin mode | [optional] 
**IsolatedPositionMargin** | **string** | Isolated position margin, applicable to new classic account margin mode | [optional] 
**EnableNewDualMode** | **bool** | Whether to open a new two-way position mode | [optional] 
**MarginMode** | **int32** | Margin mode, 0-classic margin mode, 1-cross-currency margin mode, 2-combined margin mode | [optional] 
**EnableTieredMm** | **bool** | Whether to enable tiered maintenance margin calculation | [optional] 
**History** | [**FuturesAccountHistory**](FuturesAccount_history.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


