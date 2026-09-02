# OtcUploadPreUploadData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileKey** | **string** | Base64 temporary object path; pass back **unchanged** on business submit—do not decode | 
**Url** | **string** | S3 direct upload URL | 
**Fields** | [**OtcUploadPreUploadPolicyFields**](OtcUploadPreUploadPolicyFields.md) |  | 
**ExpiresIn** | **int32** | Policy validity period in seconds; currently 5400 (90 minutes); aligns with &#x60;expiration&#x60; in &#x60;fields.Policy&#x60;; call this endpoint again after expiry | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


