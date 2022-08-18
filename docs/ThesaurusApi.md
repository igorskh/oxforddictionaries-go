# \ThesaurusApi

All URIs are relative to *https://od-api.oxforddictionaries.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThesaurusSourceLangWordIdGet**](ThesaurusApi.md#ThesaurusSourceLangWordIdGet) | **Get** /thesaurus/{source_lang}/{word_id} | Retrieve words that are similar to a given word



## ThesaurusSourceLangWordIdGet

> Thesaurus ThesaurusSourceLangWordIdGet(ctx, sourceLang, wordId).AppId(appId).AppKey(appKey).Fields(fields).StrictMatch(strictMatch).Execute()

Retrieve words that are similar to a given word



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
    wordId := "wordId_example" // string | The identifier for an Entry (case-sensitive). (default to "ace")
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter
    fields := []string{"Fields_example"} // []string | A comma-separated list of data fields to return for the matched entries. What to return - if specified, either 'synonyms', 'antonyms' or 'synonyms,antonyms' It cannot be empty.  (optional)
    strictMatch := true // bool | Specifies whether diacritics must match exactly. If \"false\", near-homographs for the given word_id will also be selected (e.g., *rose* matches both *rose* and *rosé*; similarly *rosé* matches both). (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThesaurusApi.ThesaurusSourceLangWordIdGet(context.Background(), sourceLang, wordId).AppId(appId).AppKey(appKey).Fields(fields).StrictMatch(strictMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThesaurusApi.ThesaurusSourceLangWordIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThesaurusSourceLangWordIdGet`: Thesaurus
    fmt.Fprintf(os.Stdout, "Response from `ThesaurusApi.ThesaurusSourceLangWordIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language in a Thesaurus dataset. | 
**wordId** | **string** | The identifier for an Entry (case-sensitive). | [default to &quot;ace&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiThesaurusSourceLangWordIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 
 **fields** | **[]string** | A comma-separated list of data fields to return for the matched entries. What to return - if specified, either &#39;synonyms&#39;, &#39;antonyms&#39; or &#39;synonyms,antonyms&#39; It cannot be empty.  | 
 **strictMatch** | **bool** | Specifies whether diacritics must match exactly. If \&quot;false\&quot;, near-homographs for the given word_id will also be selected (e.g., *rose* matches both *rose* and *rosé*; similarly *rosé* matches both). | [default to false]

### Return type

[**Thesaurus**](Thesaurus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

