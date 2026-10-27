# AiHubPortfolioDetailData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StrategyId** | **string** |  | 
**StrategyType** | [**StrategyType**](StrategyType.md) |  | 
**Market** | **string** | Single market, or empty string for a portfolio; use assets/markets for the latter. | 
**Status** | **string** |  | 
**BaseInfo** | [**AiHubPortfolioBaseInfo**](AIHubPortfolioBaseInfo.md) |  | 
**Metrics** | [**AiHubPortfolioMetrics**](AIHubPortfolioMetrics.md) |  | 
**Position** | Pointer to [**AiHubPortfolioPosition**](AIHubPortfolioPosition.md) |  | [optional] 
**StopSupported** | **bool** |  | 
**Assets** | **[]string** |  | [optional] 
**Markets** | **[]string** |  | [optional] 
**Positions** | [**[]AiHubPortfolioAssetPosition**](AIHubPortfolioAssetPosition.md) |  | [optional] 
**PositionsStatus** | **string** |  | [optional] 
**Configuration** | [**map[string]interface{}**](.md) | Product-specific saved configuration. spot_martingale exposes AIHubSpotMartingaleConfiguration; absence is unverified. Runtime recovery internals are excluded. | [optional] 
**StopProgress** | [**AiHubStopProgress**](AIHubStopProgress.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


