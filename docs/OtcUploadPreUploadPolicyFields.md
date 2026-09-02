# OtcUploadPreUploadPolicyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Plaintext temporary object path, identical to base64_decode(file_key) | 
**ContentType** | **string** | Must match the decoded content_type from the pre-upload request | 
**XAmzCredential** | **string** | AWS temporary credential and scope; submit them unchanged during direct upload | 
**XAmzAlgorithm** | **string** | AWS signing algorithm; submit it unchanged during direct upload | 
**XAmzDate** | **string** | AWS signing timestamp; submit it unchanged during direct upload | 
**Policy** | **string** | Base64-encoded S3 POST Policy; submit it unchanged during direct upload | 
**XAmzSignature** | **string** | S3 POST Policy signature; submit it unchanged during direct upload | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


