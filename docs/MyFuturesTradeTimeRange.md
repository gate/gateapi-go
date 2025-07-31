# MyFuturesTradeTimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeId** | **string** | Fill ID | [optional] 
**CreateTime** | **float64** | Fill Time | [optional] 
**Contract** | **string** | Futures contract | [optional] 
**OrderId** | **string** | Related order ID | [optional] 
**Size** | **int64** | Trading size | [optional] 
**CloseSize** | **int64** | Number of closed positions:  close_size&#x3D;0 &amp;&amp; size＞0 Open long position close_size&#x3D;0 &amp;&amp; size＜0 Open short position close_size&gt;0 &amp;&amp; size&gt;0 &amp;&amp; size &lt;&#x3D; close_size Close short position close_size&gt;0 &amp;&amp; size&gt;0 &amp;&amp; size &gt; close_size Close short position and open long position close_size&lt;0 &amp;&amp; size&lt;0 &amp;&amp; size &gt;&#x3D; close_size Close long position close_size&lt;0 &amp;&amp; size&lt;0 &amp;&amp; size &lt; close_size Close long position and open short position | [optional] 
**Price** | **string** | Fill Price | [optional] 
**Role** | **string** | Trade role. taker - taker, maker - maker | [optional] 
**Text** | **string** | Order custom information | [optional] 
**Fee** | **string** | Trade fee | [optional] 
**PointFee** | **string** | Points used to deduct trade fee | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


