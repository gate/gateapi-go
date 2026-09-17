# AnnouncementArticleListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TitleQuery** | **string** | Article title to use when querying announcements. | [optional] 
**Page** | **string** | Page number. Optional. Pass as a string, for example \&quot;1\&quot;. | [optional] 
**Size** | **string** | Number of articles to return per page. Optional. Pass as a string, for example \&quot;5\&quot;. | [optional] 
**Tags** | **string** | Article tags | [optional] 
**Timer** | **string** | Query announcements from the last N days, with N passed as a string. For example, on the 10th, pass \&quot;10\&quot; to query announcements from the 1st through the 10th. | [optional] 
**CateName** | **string** | Announcement category name. | [optional] 
**CateLevel** | **string** | Category level, passed as a string: \&quot;1\&quot; or \&quot;2\&quot;. | [optional] 
**SubWebsiteId** | **string** | Subsite ID, passed as a string: \&quot;0\&quot; for the main site or \&quot;177\&quot; for the Turkey site. Defaults to \&quot;0\&quot;. | [optional] [default to 0]
**Pinned** | **int32** | Whether to include pinned articles: 1 to include them or 0 to exclude them. Defaults to 1. | [optional] [default to 1]
**UpdateAfter** | **int32** | Query announcement articles updated after this timestamp. | [optional] 
**Lang** | **string** | Language code, for example \&quot;cn\&quot;. | [optional] 
**FilterEmptyContent** | **int32** | Whether to exclude articles with empty content in the current language: 0 to keep them or 1 to exclude them. Defaults to 1. | [optional] [default to 1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


