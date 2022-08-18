# \LemmasApi

All URIs are relative to *https://od-api.oxforddictionaries.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LemmasSourceLangWordIdGet**](LemmasApi.md#LemmasSourceLangWordIdGet) | **Get** /lemmas/{source_lang}/{word_id} | Check a word exists in the dictionary and retrieve its root form



## LemmasSourceLangWordIdGet

> Lemmatron LemmasSourceLangWordIdGet(ctx, sourceLang, wordId).AppId(appId).AppKey(appKey).GrammaticalFeatures(grammaticalFeatures).LexicalCategory(lexicalCategory).Execute()

Check a word exists in the dictionary and retrieve its root form



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
    wordId := "wordId_example" // string | The identifier for an Entry (case-sensitive). (default to "ace")
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter
    grammaticalFeatures := "grammaticalFeatures_example" // string | Selection filter: a comma-separated list of grammatical features ids to match on (default: all features). The available grammatical features for a given language (or language pair) can be obtained from the /grammaticalfeatures/ endpoint.  The filter keeps all the entries in the response whose grammaticalFeatures \"id\" matches the values in the grammaticalFeatures parameter. ```@javascript {   \"grammaticalFeatures\": [                           {                               \"id\": \"mass\",                               \"text\": \"Mass\",                               \"type\": \"Countability\"                           }                       ] } ```  (optional)
    lexicalCategory := "lexicalCategory_example" // string | Selection filter: a comma-separated list of lexical categories ids to match on (default: all categories). The available lexical categories for a given language (or language pair) can be obtained from the /lexicalcategories/ endpoint.  The filter keeps all the entries in the response whose lexicalCategory \"id\" matches the values in the lexicalCategory parameter. ```@javascript {   \"lexicalCategory\": {                   \"id\": \"adjective\",                   \"text\": \"Adjective\"               } } ```  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LemmasApi.LemmasSourceLangWordIdGet(context.Background(), sourceLang, wordId).AppId(appId).AppKey(appKey).GrammaticalFeatures(grammaticalFeatures).LexicalCategory(lexicalCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LemmasApi.LemmasSourceLangWordIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LemmasSourceLangWordIdGet`: Lemmatron
    fmt.Fprintf(os.Stdout, "Response from `LemmasApi.LemmasSourceLangWordIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language in a monolingual dataset. | 
**wordId** | **string** | The identifier for an Entry (case-sensitive). | [default to &quot;ace&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiLemmasSourceLangWordIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 
 **grammaticalFeatures** | **string** | Selection filter: a comma-separated list of grammatical features ids to match on (default: all features). The available grammatical features for a given language (or language pair) can be obtained from the /grammaticalfeatures/ endpoint.  The filter keeps all the entries in the response whose grammaticalFeatures \&quot;id\&quot; matches the values in the grammaticalFeatures parameter. &#x60;&#x60;&#x60;@javascript {   \&quot;grammaticalFeatures\&quot;: [                           {                               \&quot;id\&quot;: \&quot;mass\&quot;,                               \&quot;text\&quot;: \&quot;Mass\&quot;,                               \&quot;type\&quot;: \&quot;Countability\&quot;                           }                       ] } &#x60;&#x60;&#x60;  | 
 **lexicalCategory** | **string** | Selection filter: a comma-separated list of lexical categories ids to match on (default: all categories). The available lexical categories for a given language (or language pair) can be obtained from the /lexicalcategories/ endpoint.  The filter keeps all the entries in the response whose lexicalCategory \&quot;id\&quot; matches the values in the lexicalCategory parameter. &#x60;&#x60;&#x60;@javascript {   \&quot;lexicalCategory\&quot;: {                   \&quot;id\&quot;: \&quot;adjective\&quot;,                   \&quot;text\&quot;: \&quot;Adjective\&quot;               } } &#x60;&#x60;&#x60;  | 

### Return type

[**Lemmatron**](Lemmatron.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

