# CrossexOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**OrderId** | **string** | Order ID | 
**Text** | **string** | Customer-defined order ID | 
**State** | **string** | Order State:  NEW: The order is legal and waiting to be sent to the exchange  OPEN: The order has been placed on the orderbook of the exchange  PARTIALLY_FILLED: The order has been partially completed  FILLED: The order has been fully executed  FAIL: The order verification in CrossEx did not pass. Please check the order reason  REJECT：The order was rejected by the exchange. Please check the order reason | 
**Symbol** | **string** | Trading pair unique identifier ,example: BINANCE_SPOT_BTC_USDT, BINANCE_FUTURE_BTC_USDT | 
**Side** | **string** | Side(BUY,SELL) | 
**Type** | **string** | Type(LIMIT, MARKET) | 
**Attribute** | **string** | COMMON, LIQ, REDUCE, ADL | 
**ExchangeType** | **string** | Exchange type(BINANCE,OKX,GATE,BYBIT) | 
**BusinessType** | **string** | Business type(SPOT,FUTURE,MARGIN) | 
**Qty** | **string** | Order base quantity | 
**QuoteQty** | **string** | Order quote quantity | 
**Price** | **string** | Order price | 
**TimeInForce** | **string** | Timeinforce (default GTC, enums:GTC,IOC,FOK,POC) | 
**ExecutedQty** | **string** | Executed quantity | 
**ExecutedAmount** | **string** | Executed quote quantity | 
**ExecutedAvgPrice** | **string** | Average transaction price | 
**FeeCoin** | **string** | Transaction fee coin | 
**Fee** | **string** | Transaction fee amount | 
**ReduceOnly** | **string** | Reduce position orders only, \&quot;true\&quot; or \&quot;false\&quot; | 
**Leverage** | **string** | Order leverage | 
**Reason** | **string** | Fail message | 
**LastExecutedQty** | **string** | Last transaction quantity | 
**LastExecutedPrice** | **string** | Last transaction price | 
**LastExecutedAmount** | **string** | Last transaction amount | 
**PositionSide** | **string** | Position side(NONE/LONG/SHORT) | 
**CreateTime** | **string** | Create time | 
**UpdateTime** | **string** | Update time | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


