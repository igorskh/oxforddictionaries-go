# \SentencesApi

All URIs are relative to *https://od-api.oxforddictionaries.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SentencesSourceLangWordIdGet**](SentencesApi.md#SentencesSourceLangWordIdGet) | **Get** /sentences/{source_lang}/{word_id} | Retrieve real example sentences of a word in use



## SentencesSourceLangWordIdGet

> SentencesResults SentencesSourceLangWordIdGet(ctx, sourceLang, wordId).AppId(appId).AppKey(appKey).StrictMatch(strictMatch).Execute()

Retrieve real example sentences of a word in use



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
    sourceLang := "sourceLang_example" // string | Language code of the source language.
    wordId := "wordId_example" // string | The identifier for an Entry (case-sensitive). (default to "ace")
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter
    strictMatch := true // bool | Specifies whether diacritics must match exactly. If \"false\", near-homographs for the given word_id will also be selected (e.g., *rose* matches both *rose* and *rosé*; similarly *rosé* matches both). (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SentencesApi.SentencesSourceLangWordIdGet(context.Background(), sourceLang, wordId).AppId(appId).AppKey(appKey).StrictMatch(strictMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SentencesApi.SentencesSourceLangWordIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SentencesSourceLangWordIdGet`: SentencesResults
    fmt.Fprintf(os.Stdout, "Response from `SentencesApi.SentencesSourceLangWordIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language. | 
**wordId** | **string** | The identifier for an Entry (case-sensitive). | [default to &quot;ace&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiSentencesSourceLangWordIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 
 **strictMatch** | **bool** | Specifies whether diacritics must match exactly. If \&quot;false\&quot;, near-homographs for the given word_id will also be selected (e.g., *rose* matches both *rose* and *rosé*; similarly *rosé* matches both). | [default to false]

### Return type

[**SentencesResults**](SentencesResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

