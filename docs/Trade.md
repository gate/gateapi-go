# Trade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Fill ID | [optional] 
**CreateTime** | **string** | Fill Time | [optional] 
**CreateTimeMs** | **string** | Trading time, with millisecond precision | [optional] 
**CurrencyPair** | **string** | Currency pair | [optional] 
**Side** | **string** | Buy or sell order | [optional] 
**Role** | **string** | Trade role, not returned in public endpoints | [optional] 
**Amount** | **string** | Trade amount | [optional] 
**Price** | **string** | Order price | [optional] 
**OrderId** | **string** | Related order ID, not returned in public endpoints | [optional] 
**Fee** | **string** | Fee deducted, not returned in public endpoints | [optional] 
**FeeCurrency** | **string** | Fee currency unit, not returned in public endpoints | [optional] 
**PointFee** | **string** | Points used to deduct fee, not returned in public endpoints | [optional] 
**GtFee** | **string** | GT used to deduct fee, not returned in public endpoints | [optional] 
**AmendText** | **string** | The custom data that the user remarked when amending the order | [optional] 
**SequenceId** | **string** | Consecutive trade ID within a single market. Used to track and identify trades in the specific market | [optional] 
**Text** | **string** | Order&#39;s Custom Information. This field is not returned by public interfaces. The scenarios pm_liquidate, comb_margin_liquidate, and scm_liquidate represent full-account forced liquidation orders. liquidate represents isolated-account forced liquidation orders. | [optional] 
**Deal** | **string** | Total Executed Value | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


