# \WordsApi

All URIs are relative to *https://od-api.oxforddictionaries.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WordsSourceLangGet**](WordsApi.md#WordsSourceLangGet) | **Get** /words/{source_lang} | Check if an inflected form exists in the dictionary and retrieve the entries data of its root form.



## WordsSourceLangGet

> RetrieveEntry WordsSourceLangGet(ctx, sourceLang).Q(q).AppId(appId).AppKey(appKey).Fields(fields).GrammaticalFeatures(grammaticalFeatures).LexicalCategory(lexicalCategory).Domains(domains).Registers(registers).Execute()

Check if an inflected form exists in the dictionary and retrieve the entries data of its root form.



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
    q := "q_example" // string | The word or inflection to retrieve. This parameter is NOT case-sensitive. (default to "swimming")
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter
    fields := []string{"Fields_example"} // []string | A comma-separated list of data fields to return for the matched entries. Valid field names are 'definitions', 'domains', 'etymologies', 'examples', 'pronunciations', 'regions', 'registers' and 'variantForms'. (a) If not specified, all available fields for each word_id are returned. (b) If specified and empty, the minimum payload for each word_id is returned. (c) If more field names are specified, then the minimum payload plus the specified fields for each word_id are returned.  (optional)
    grammaticalFeatures := "grammaticalFeatures_example" // string | Selection filter: a comma-separated list of grammatical features ids to match on (default: all features). The available grammatical features for a given language (or language pair) can be obtained from the /grammaticalfeatures/ endpoint.  The filter keeps all the entries in the response whose grammaticalFeatures \"id\" matches the values in the grammaticalFeatures parameter. ```@javascript {   \"grammaticalFeatures\": [                           {                               \"id\": \"mass\",                               \"text\": \"Mass\",                               \"type\": \"Countability\"                           }                       ] } ```  (optional)
    lexicalCategory := "lexicalCategory_example" // string | Selection filter: a comma-separated list of lexical categories ids to match on (default: all categories). The available lexical categories for a given language (or language pair) can be obtained from the /lexicalcategories/ endpoint.  The filter keeps all the entries in the response whose lexicalCategory \"id\" matches the values in the lexicalCategory parameter. ```@javascript {   \"lexicalCategory\": {                   \"id\": \"adjective\",                   \"text\": \"Adjective\"               } } ```  (optional)
    domains := "domains_example" // string | Selection filter: a comma-separated list of domains ids to match on (default: all domains). The available domains for a given language (or language pair) can be obtained from the /domains/ endpoint.  The filter keeps all the senses and subsenses in the response whose domains \"id\" matches the values in the domains parameter.  ```@javascript {   \"domains\": [     {       \"id\": \"jazz\",       \"text\": \"Jazz\"      }] } ```  (optional)
    registers := "registers_example" // string | Selection filter: a comma-separated list of registers ids to match on (default: all registers). The available registers for a given language (or language pair) can be obtained from the /registers/ endpoint.  The filter keeps all the senses and subsenses in the response whose registers \"id\" matches the values in the registers parameter.  ```@javascript {   \"registers\": [     {       \"id\": \"informal\",       \"text\": \"Informal\"      }] } ```  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WordsApi.WordsSourceLangGet(context.Background(), sourceLang).Q(q).AppId(appId).AppKey(appKey).Fields(fields).GrammaticalFeatures(grammaticalFeatures).LexicalCategory(lexicalCategory).Domains(domains).Registers(registers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WordsApi.WordsSourceLangGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WordsSourceLangGet`: RetrieveEntry
    fmt.Fprintf(os.Stdout, "Response from `WordsApi.WordsSourceLangGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language in a monolingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWordsSourceLangGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | The word or inflection to retrieve. This parameter is NOT case-sensitive. | [default to &quot;swimming&quot;]
 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 
 **fields** | **[]string** | A comma-separated list of data fields to return for the matched entries. Valid field names are &#39;definitions&#39;, &#39;domains&#39;, &#39;etymologies&#39;, &#39;examples&#39;, &#39;pronunciations&#39;, &#39;regions&#39;, &#39;registers&#39; and &#39;variantForms&#39;. (a) If not specified, all available fields for each word_id are returned. (b) If specified and empty, the minimum payload for each word_id is returned. (c) If more field names are specified, then the minimum payload plus the specified fields for each word_id are returned.  | 
 **grammaticalFeatures** | **string** | Selection filter: a comma-separated list of grammatical features ids to match on (default: all features). The available grammatical features for a given language (or language pair) can be obtained from the /grammaticalfeatures/ endpoint.  The filter keeps all the entries in the response whose grammaticalFeatures \&quot;id\&quot; matches the values in the grammaticalFeatures parameter. &#x60;&#x60;&#x60;@javascript {   \&quot;grammaticalFeatures\&quot;: [                           {                               \&quot;id\&quot;: \&quot;mass\&quot;,                               \&quot;text\&quot;: \&quot;Mass\&quot;,                               \&quot;type\&quot;: \&quot;Countability\&quot;                           }                       ] } &#x60;&#x60;&#x60;  | 
 **lexicalCategory** | **string** | Selection filter: a comma-separated list of lexical categories ids to match on (default: all categories). The available lexical categories for a given language (or language pair) can be obtained from the /lexicalcategories/ endpoint.  The filter keeps all the entries in the response whose lexicalCategory \&quot;id\&quot; matches the values in the lexicalCategory parameter. &#x60;&#x60;&#x60;@javascript {   \&quot;lexicalCategory\&quot;: {                   \&quot;id\&quot;: \&quot;adjective\&quot;,                   \&quot;text\&quot;: \&quot;Adjective\&quot;               } } &#x60;&#x60;&#x60;  | 
 **domains** | **string** | Selection filter: a comma-separated list of domains ids to match on (default: all domains). The available domains for a given language (or language pair) can be obtained from the /domains/ endpoint.  The filter keeps all the senses and subsenses in the response whose domains \&quot;id\&quot; matches the values in the domains parameter.  &#x60;&#x60;&#x60;@javascript {   \&quot;domains\&quot;: [     {       \&quot;id\&quot;: \&quot;jazz\&quot;,       \&quot;text\&quot;: \&quot;Jazz\&quot;      }] } &#x60;&#x60;&#x60;  | 
 **registers** | **string** | Selection filter: a comma-separated list of registers ids to match on (default: all registers). The available registers for a given language (or language pair) can be obtained from the /registers/ endpoint.  The filter keeps all the senses and subsenses in the response whose registers \&quot;id\&quot; matches the values in the registers parameter.  &#x60;&#x60;&#x60;@javascript {   \&quot;registers\&quot;: [     {       \&quot;id\&quot;: \&quot;informal\&quot;,       \&quot;text\&quot;: \&quot;Informal\&quot;      }] } &#x60;&#x60;&#x60;  | 

### Return type

[**RetrieveEntry**](RetrieveEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

