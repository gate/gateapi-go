# MarginAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyPair** | **string** | Currency pair | [optional] 
**AccountType** | **string** | Account type: risk - risk rate account, mmr - maintenance margin rate account, inactive - market not activated | [optional] 
**Leverage** | **string** | User&#39;s current market leverage multiplier | [optional] 
**Locked** | **bool** | Whether the account is locked | [optional] 
**Risk** | **string** | Current risk rate of the margin account (returned when the account is a risk rate account) | [optional] 
**Mmr** | **string** | Leveraged Account Current Maintenance Margin Rate (returned when the Account is Account) | [optional] 
**Base** | [**MarginAccountCurrency**](MarginAccountCurrency.md) |  | [optional] 
**Quote** | [**MarginAccountCurrency**](MarginAccountCurrency.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


