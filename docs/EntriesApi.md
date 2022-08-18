# \EntriesApi

All URIs are relative to *https://od-api.oxforddictionaries.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntriesSourceLangWordIdGet**](EntriesApi.md#EntriesSourceLangWordIdGet) | **Get** /entries/{source_lang}/{word_id} | Retrieve dictionary information for a given word



## EntriesSourceLangWordIdGet

> RetrieveEntry EntriesSourceLangWordIdGet(ctx, sourceLang, wordId).AppId(appId).AppKey(appKey).Fields(fields).GrammaticalFeatures(grammaticalFeatures).LexicalCategory(lexicalCategory).Domains(domains).Registers(registers).StrictMatch(strictMatch).Execute()

Retrieve dictionary information for a given word



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
    fields := []string{"Fields_example"} // []string | A comma-separated list of data fields to return for the matched entries. Valid field names are 'definitions', 'domains', 'etymologies', 'examples', 'pronunciations', 'regions', 'registers' and 'variantForms'. (a) If not specified, all available fields for each word_id are returned. (b) If specified and empty, the minimum payload for each word_id is returned. (c) If more field names are specified, then the minimum payload plus the specified fields for each word_id are returned.  (optional)
    grammaticalFeatures := "grammaticalFeatures_example" // string | Selection filter: a comma-separated list of grammatical features ids to match on (default: all features). The available grammatical features for a given language (or language pair) can be obtained from the /grammaticalfeatures/ endpoint.  The filter keeps all the entries in the response whose grammaticalFeatures \"id\" matches the values in the grammaticalFeatures parameter. ```@javascript {   \"grammaticalFeatures\": [                           {                               \"id\": \"mass\",                               \"text\": \"Mass\",                               \"type\": \"Countability\"                           }                       ] } ```  (optional)
    lexicalCategory := "lexicalCategory_example" // string | Selection filter: a comma-separated list of lexical categories ids to match on (default: all categories). The available lexical categories for a given language (or language pair) can be obtained from the /lexicalcategories/ endpoint.  The filter keeps all the entries in the response whose lexicalCategory \"id\" matches the values in the lexicalCategory parameter. ```@javascript {   \"lexicalCategory\": {                   \"id\": \"adjective\",                   \"text\": \"Adjective\"               } } ```  (optional)
    domains := "domains_example" // string | Selection filter: a comma-separated list of domains ids to match on (default: all domains). The available domains for a given language (or language pair) can be obtained from the /domains/ endpoint.  The filter keeps all the senses and subsenses in the response whose domains \"id\" matches the values in the domains parameter.  ```@javascript {   \"domains\": [     {       \"id\": \"jazz\",       \"text\": \"Jazz\"      }] } ```  (optional)
    registers := "registers_example" // string | Selection filter: a comma-separated list of registers ids to match on (default: all registers). The available registers for a given language (or language pair) can be obtained from the /registers/ endpoint.  The filter keeps all the senses and subsenses in the response whose registers \"id\" matches the values in the registers parameter.  ```@javascript {   \"registers\": [     {       \"id\": \"informal\",       \"text\": \"Informal\"      }] } ```  (optional)
    strictMatch := true // bool | Specifies whether diacritics must match exactly. If \"false\", near-homographs for the given word_id will also be selected (e.g., *rose* matches both *rose* and *rosé*; similarly *rosé* matches both). (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntriesApi.EntriesSourceLangWordIdGet(context.Background(), sourceLang, wordId).AppId(appId).AppKey(appKey).Fields(fields).GrammaticalFeatures(grammaticalFeatures).LexicalCategory(lexicalCategory).Domains(domains).Registers(registers).StrictMatch(strictMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.EntriesSourceLangWordIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntriesSourceLangWordIdGet`: RetrieveEntry
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.EntriesSourceLangWordIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language in a monolingual dataset. | 
**wordId** | **string** | The identifier for an Entry (case-sensitive). | [default to &quot;ace&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiEntriesSourceLangWordIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 
 **fields** | **[]string** | A comma-separated list of data fields to return for the matched entries. Valid field names are &#39;definitions&#39;, &#39;domains&#39;, &#39;etymologies&#39;, &#39;examples&#39;, &#39;pronunciations&#39;, &#39;regions&#39;, &#39;registers&#39; and &#39;variantForms&#39;. (a) If not specified, all available fields for each word_id are returned. (b) If specified and empty, the minimum payload for each word_id is returned. (c) If more field names are specified, then the minimum payload plus the specified fields for each word_id are returned.  | 
 **grammaticalFeatures** | **string** | Selection filter: a comma-separated list of grammatical features ids to match on (default: all features). The available grammatical features for a given language (or language pair) can be obtained from the /grammaticalfeatures/ endpoint.  The filter keeps all the entries in the response whose grammaticalFeatures \&quot;id\&quot; matches the values in the grammaticalFeatures parameter. &#x60;&#x60;&#x60;@javascript {   \&quot;grammaticalFeatures\&quot;: [                           {                               \&quot;id\&quot;: \&quot;mass\&quot;,                               \&quot;text\&quot;: \&quot;Mass\&quot;,                               \&quot;type\&quot;: \&quot;Countability\&quot;                           }                       ] } &#x60;&#x60;&#x60;  | 
 **lexicalCategory** | **string** | Selection filter: a comma-separated list of lexical categories ids to match on (default: all categories). The available lexical categories for a given language (or language pair) can be obtained from the /lexicalcategories/ endpoint.  The filter keeps all the entries in the response whose lexicalCategory \&quot;id\&quot; matches the values in the lexicalCategory parameter. &#x60;&#x60;&#x60;@javascript {   \&quot;lexicalCategory\&quot;: {                   \&quot;id\&quot;: \&quot;adjective\&quot;,                   \&quot;text\&quot;: \&quot;Adjective\&quot;               } } &#x60;&#x60;&#x60;  | 
 **domains** | **string** | Selection filter: a comma-separated list of domains ids to match on (default: all domains). The available domains for a given language (or language pair) can be obtained from the /domains/ endpoint.  The filter keeps all the senses and subsenses in the response whose domains \&quot;id\&quot; matches the values in the domains parameter.  &#x60;&#x60;&#x60;@javascript {   \&quot;domains\&quot;: [     {       \&quot;id\&quot;: \&quot;jazz\&quot;,       \&quot;text\&quot;: \&quot;Jazz\&quot;      }] } &#x60;&#x60;&#x60;  | 
 **registers** | **string** | Selection filter: a comma-separated list of registers ids to match on (default: all registers). The available registers for a given language (or language pair) can be obtained from the /registers/ endpoint.  The filter keeps all the senses and subsenses in the response whose registers \&quot;id\&quot; matches the values in the registers parameter.  &#x60;&#x60;&#x60;@javascript {   \&quot;registers\&quot;: [     {       \&quot;id\&quot;: \&quot;informal\&quot;,       \&quot;text\&quot;: \&quot;Informal\&quot;      }] } &#x60;&#x60;&#x60;  | 
 **strictMatch** | **bool** | Specifies whether diacritics must match exactly. If \&quot;false\&quot;, near-homographs for the given word_id will also be selected (e.g., *rose* matches both *rose* and *rosé*; similarly *rosé* matches both). | [default to false]

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

