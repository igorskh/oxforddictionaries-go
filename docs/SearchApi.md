# \SearchApi

All URIs are relative to *https://od-api.oxforddictionaries.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSourceLangGet**](SearchApi.md#SearchSourceLangGet) | **Get** /search/{source_lang} | Retrieves possible dictionary matches to an input string
[**SearchThesaurusSourceLangGet**](SearchApi.md#SearchThesaurusSourceLangGet) | **Get** /search/thesaurus/{source_lang} | Retrieves possible dictionary matches to an input string
[**SearchTranslationsSourceLangSearchTargetLangSearchGet**](SearchApi.md#SearchTranslationsSourceLangSearchTargetLangSearchGet) | **Get** /search/translations/{source_lang_search}/{target_lang_search} | Retrieves possible headwords with translations



## SearchSourceLangGet

> Wordlist SearchSourceLangGet(ctx, sourceLang).Q(q).AppId(appId).AppKey(appKey).Prefix(prefix).Limit(limit).Offset(offset).Execute()

Retrieves possible dictionary matches to an input string



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceLang := "sourceLang_example" // string | Language code of the source language in a monolingual dataset.
    q := "q_example" // string | Search string
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter
    prefix := true // bool | Use prefix=true to return only results that start with the value of the \"q\" parameter. (optional)
    limit := int32(56) // int32 | Restricts number of returned results. Default and max. is 5000. (optional)
    offset := int32(56) // int32 | Pagination - results offset.  The sum of offset and limit must not exceed 10000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.SearchSourceLangGet(context.Background(), sourceLang).Q(q).AppId(appId).AppKey(appKey).Prefix(prefix).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchSourceLangGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSourceLangGet`: Wordlist
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchSourceLangGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language in a monolingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSourceLangGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Search string | 
 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 
 **prefix** | **bool** | Use prefix&#x3D;true to return only results that start with the value of the \&quot;q\&quot; parameter. | 
 **limit** | **int32** | Restricts number of returned results. Default and max. is 5000. | 
 **offset** | **int32** | Pagination - results offset.  The sum of offset and limit must not exceed 10000. | 

### Return type

[**Wordlist**](Wordlist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchThesaurusSourceLangGet

> Wordlist SearchThesaurusSourceLangGet(ctx, sourceLang).Q(q).AppId(appId).AppKey(appKey).Prefix(prefix).Limit(limit).Offset(offset).Execute()

Retrieves possible dictionary matches to an input string



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceLang := "sourceLang_example" // string | Language code of the source language in a Thesaurus dataset.
    q := "q_example" // string | Search string
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter
    prefix := true // bool | Use prefix=true to return only results that start with the value of the \"q\" parameter. (optional)
    limit := int32(56) // int32 | Restricts number of returned results. Default and max. is 5000. (optional)
    offset := int32(56) // int32 | Pagination - results offset.  The sum of offset and limit must not exceed 10000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.SearchThesaurusSourceLangGet(context.Background(), sourceLang).Q(q).AppId(appId).AppKey(appKey).Prefix(prefix).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchThesaurusSourceLangGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchThesaurusSourceLangGet`: Wordlist
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchThesaurusSourceLangGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language in a Thesaurus dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchThesaurusSourceLangGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Search string | 
 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 
 **prefix** | **bool** | Use prefix&#x3D;true to return only results that start with the value of the \&quot;q\&quot; parameter. | 
 **limit** | **int32** | Restricts number of returned results. Default and max. is 5000. | 
 **offset** | **int32** | Pagination - results offset.  The sum of offset and limit must not exceed 10000. | 

### Return type

[**Wordlist**](Wordlist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTranslationsSourceLangSearchTargetLangSearchGet

> Wordlist SearchTranslationsSourceLangSearchTargetLangSearchGet(ctx, sourceLangSearch, targetLangSearch).Q(q).AppId(appId).AppKey(appKey).Prefix(prefix).Limit(limit).Offset(offset).Execute()

Retrieves possible headwords with translations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceLangSearch := "sourceLangSearch_example" // string | Language code of the source language in a bilingual dataset.
    targetLangSearch := "targetLangSearch_example" // string | Language code of the target language in a bilingual dataset.
    q := "q_example" // string | Search string
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter
    prefix := true // bool | Use prefix=true to return only results that start with the value of the \"q\" parameter. (optional)
    limit := int32(56) // int32 | Restricts number of returned results. Default and max. is 5000. (optional)
    offset := int32(56) // int32 | pagination - results offset.  The sum of offset and limit must not exceed 10000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.SearchTranslationsSourceLangSearchTargetLangSearchGet(context.Background(), sourceLangSearch, targetLangSearch).Q(q).AppId(appId).AppKey(appKey).Prefix(prefix).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchTranslationsSourceLangSearchTargetLangSearchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchTranslationsSourceLangSearchTargetLangSearchGet`: Wordlist
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchTranslationsSourceLangSearchTargetLangSearchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLangSearch** | **string** | Language code of the source language in a bilingual dataset. | 
**targetLangSearch** | **string** | Language code of the target language in a bilingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchTranslationsSourceLangSearchTargetLangSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** | Search string | 
 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 
 **prefix** | **bool** | Use prefix&#x3D;true to return only results that start with the value of the \&quot;q\&quot; parameter. | 
 **limit** | **int32** | Restricts number of returned results. Default and max. is 5000. | 
 **offset** | **int32** | pagination - results offset.  The sum of offset and limit must not exceed 10000. | 

### Return type

[**Wordlist**](Wordlist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

