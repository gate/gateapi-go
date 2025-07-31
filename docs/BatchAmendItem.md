# BatchAmendItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | The order ID returned upon successful creation or the custom ID specified by the user during creation (i.e., the &#39;text&#39; field) | 
**CurrencyPair** | **string** | Currency pair | 
**Account** | **string** | Default spot, unified account and warehouse-by-store leverage account | [optional] 
**Amount** | **string** | Trading Quantity. Only one of &#x60;amount&#x60; or &#x60;price&#x60; can be specified | [optional] 
**Price** | **string** | Trading Price. Only one of &#x60;amount&#x60; or &#x60;price&#x60; can be specified | [optional] 
**AmendText** | **string** | Custom info during order amendment | [optional] 
**ActionMode** | **string** | Processing Mode: When placing an order, different fields are returned based on action_mode. This field is only valid during the request and is not included in the response result ACK: Asynchronous mode, only returns key order fields RESULT: No clearing information FULL: Full mode (default) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


