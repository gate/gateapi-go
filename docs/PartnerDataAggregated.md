# PartnerDataAggregated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebateAmount** | **string** | 返佣金额，字符串格式保证精度  最多保留 6 位小数，去除尾零 | 
**TradeVolume** | **string** | 交易量，字符串格式保证精度  最多保留 6 位小数，去除尾零 | 
**NetFee** | **string** | 净手续费，字符串格式保证精度  最多保留 6 位小数，去除尾零 | 
**CustomerCount** | **int32** | Customer count (invited users) | 
**TradingUserCount** | Pointer to **string** | 交易人数，字符串形式（与线上 JSON 序列化一致）  仅在 business_type&#x3D;0（全部）时返回具体数值，其他业务类型返回 null | 
**TimeRangeDesc** | **string** | Time range description | 
**BusinessType** | **int32** | Business Type | 
**BusinessTypeDesc** | **string** | 业务类型描述，可取值：全部, 现货, 合约, Alpha, Web3, Perps(DEX), Exchange All, Web3 All, TradFi | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


