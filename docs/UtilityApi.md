# \UtilityApi

All URIs are relative to *https://od-api.oxforddictionaries.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsSourceLangDomainsTargetLangDomainsGet**](UtilityApi.md#DomainsSourceLangDomainsTargetLangDomainsGet) | **Get** /domains/{source_lang_domains}/{target_lang_domains} | Lists available domains in a bilingual dataset
[**DomainsSourceLangGet**](UtilityApi.md#DomainsSourceLangGet) | **Get** /domains/{source_lang} | Lists available domains in a monolingual dataset
[**FieldsEndpointGet**](UtilityApi.md#FieldsEndpointGet) | **Get** /fields/{endpoint} | Lists available fields for specific endpoint
[**FieldsGet**](UtilityApi.md#FieldsGet) | **Get** /fields | Lists available fields
[**FiltersEndpointGet**](UtilityApi.md#FiltersEndpointGet) | **Get** /filters/{endpoint} | Lists available filters for specific endpoint
[**FiltersGet**](UtilityApi.md#FiltersGet) | **Get** /filters | Lists available filters
[**GrammaticalFeaturesSourceLangGet**](UtilityApi.md#GrammaticalFeaturesSourceLangGet) | **Get** /grammaticalFeatures/{source_lang} | Lists available grammatical features in a monolingual dataset
[**GrammaticalFeaturesSourceLangGrammaticalTargetLangGrammaticalGet**](UtilityApi.md#GrammaticalFeaturesSourceLangGrammaticalTargetLangGrammaticalGet) | **Get** /grammaticalFeatures/{source_lang_grammatical}/{target_lang_grammatical} | Lists available grammatical features in a bilingual dataset
[**LanguagesGet**](UtilityApi.md#LanguagesGet) | **Get** /languages | Returns the names of Dictionaries in the API
[**LexicalCategoriesSourceLangGet**](UtilityApi.md#LexicalCategoriesSourceLangGet) | **Get** /lexicalCategories/{source_lang} | Lists available lexical categories in a monolingual dataset
[**LexicalCategoriesSourceLangLexicalTargetLangLexicalGet**](UtilityApi.md#LexicalCategoriesSourceLangLexicalTargetLangLexicalGet) | **Get** /lexicalCategories/{source_lang_lexical}/{target_lang_lexical} | Lists available lexical categories in a bilingual dataset
[**RegistersSourceLangGet**](UtilityApi.md#RegistersSourceLangGet) | **Get** /registers/{source_lang} | Lists available registers in a  monolingual dataset
[**RegistersSourceLangRegistersTargetLangRegistersGet**](UtilityApi.md#RegistersSourceLangRegistersTargetLangRegistersGet) | **Get** /registers/{source_lang_registers}/{target_lang_registers} | Lists available registers in a bilingual dataset



## DomainsSourceLangDomainsTargetLangDomainsGet

> UtilityLabels DomainsSourceLangDomainsTargetLangDomainsGet(ctx, sourceLangDomains, targetLangDomains).AppId(appId).AppKey(appKey).Execute()

Lists available domains in a bilingual dataset



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
    sourceLangDomains := "sourceLangDomains_example" // string | Language code of the source language in a bilingual dataset.
    targetLangDomains := "targetLangDomains_example" // string | Language code of the target language in a bilingual dataset.
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.DomainsSourceLangDomainsTargetLangDomainsGet(context.Background(), sourceLangDomains, targetLangDomains).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.DomainsSourceLangDomainsTargetLangDomainsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsSourceLangDomainsTargetLangDomainsGet`: UtilityLabels
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.DomainsSourceLangDomainsTargetLangDomainsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLangDomains** | **string** | Language code of the source language in a bilingual dataset. | 
**targetLangDomains** | **string** | Language code of the target language in a bilingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsSourceLangDomainsTargetLangDomainsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsSourceLangGet

> UtilityLabels DomainsSourceLangGet(ctx, sourceLang).AppId(appId).AppKey(appKey).Execute()

Lists available domains in a monolingual dataset



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
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.DomainsSourceLangGet(context.Background(), sourceLang).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.DomainsSourceLangGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsSourceLangGet`: UtilityLabels
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.DomainsSourceLangGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language in a monolingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsSourceLangGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldsEndpointGet

> Filters FieldsEndpointGet(ctx, endpoint).AppId(appId).AppKey(appKey).Execute()

Lists available fields for specific endpoint



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
    endpoint := "endpoint_example" // string | Name of the endpoint
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.FieldsEndpointGet(context.Background(), endpoint).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.FieldsEndpointGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FieldsEndpointGet`: Filters
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.FieldsEndpointGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpoint** | **string** | Name of the endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldsEndpointGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**Filters**](Filters.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldsGet

> Filters FieldsGet(ctx).AppId(appId).AppKey(appKey).Execute()

Lists available fields



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
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.FieldsGet(context.Background()).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.FieldsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FieldsGet`: Filters
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.FieldsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFieldsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**Filters**](Filters.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FiltersEndpointGet

> Filters FiltersEndpointGet(ctx, endpoint).AppId(appId).AppKey(appKey).Execute()

Lists available filters for specific endpoint



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
    endpoint := "endpoint_example" // string | Name of the endpoint.
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.FiltersEndpointGet(context.Background(), endpoint).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.FiltersEndpointGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FiltersEndpointGet`: Filters
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.FiltersEndpointGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpoint** | **string** | Name of the endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFiltersEndpointGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**Filters**](Filters.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FiltersGet

> Filters FiltersGet(ctx).AppId(appId).AppKey(appKey).Execute()

Lists available filters



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
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.FiltersGet(context.Background()).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.FiltersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FiltersGet`: Filters
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.FiltersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFiltersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**Filters**](Filters.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrammaticalFeaturesSourceLangGet

> UtilityLabels GrammaticalFeaturesSourceLangGet(ctx, sourceLang).AppId(appId).AppKey(appKey).Execute()

Lists available grammatical features in a monolingual dataset



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
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.GrammaticalFeaturesSourceLangGet(context.Background(), sourceLang).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.GrammaticalFeaturesSourceLangGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GrammaticalFeaturesSourceLangGet`: UtilityLabels
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.GrammaticalFeaturesSourceLangGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language in a monolingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrammaticalFeaturesSourceLangGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrammaticalFeaturesSourceLangGrammaticalTargetLangGrammaticalGet

> UtilityLabels GrammaticalFeaturesSourceLangGrammaticalTargetLangGrammaticalGet(ctx, sourceLangGrammatical, targetLangGrammatical).AppId(appId).AppKey(appKey).Execute()

Lists available grammatical features in a bilingual dataset



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
    sourceLangGrammatical := "sourceLangGrammatical_example" // string | Language code of the source language in a bilingual dataset.
    targetLangGrammatical := "targetLangGrammatical_example" // string | Language code of the target language in a bilingual dataset.
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.GrammaticalFeaturesSourceLangGrammaticalTargetLangGrammaticalGet(context.Background(), sourceLangGrammatical, targetLangGrammatical).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.GrammaticalFeaturesSourceLangGrammaticalTargetLangGrammaticalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GrammaticalFeaturesSourceLangGrammaticalTargetLangGrammaticalGet`: UtilityLabels
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.GrammaticalFeaturesSourceLangGrammaticalTargetLangGrammaticalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLangGrammatical** | **string** | Language code of the source language in a bilingual dataset. | 
**targetLangGrammatical** | **string** | Language code of the target language in a bilingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrammaticalFeaturesSourceLangGrammaticalTargetLangGrammaticalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LanguagesGet

> Languages LanguagesGet(ctx).AppId(appId).AppKey(appKey).SourceLanguage(sourceLanguage).TargetLanguage(targetLanguage).Execute()

Returns the names of Dictionaries in the API



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
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter
    sourceLanguage := "sourceLanguage_example" // string | Source Language. If provided, output will be filtered by sourceLanguage. (optional)
    targetLanguage := "targetLanguage_example" // string | Target Language. If provided, output will be filtered by targetLanguage. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.LanguagesGet(context.Background()).AppId(appId).AppKey(appKey).SourceLanguage(sourceLanguage).TargetLanguage(targetLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.LanguagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LanguagesGet`: Languages
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.LanguagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLanguagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 
 **sourceLanguage** | **string** | Source Language. If provided, output will be filtered by sourceLanguage. | 
 **targetLanguage** | **string** | Target Language. If provided, output will be filtered by targetLanguage. | 

### Return type

[**Languages**](Languages.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LexicalCategoriesSourceLangGet

> UtilityLabels LexicalCategoriesSourceLangGet(ctx, sourceLang).AppId(appId).AppKey(appKey).Execute()

Lists available lexical categories in a monolingual dataset



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
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.LexicalCategoriesSourceLangGet(context.Background(), sourceLang).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.LexicalCategoriesSourceLangGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LexicalCategoriesSourceLangGet`: UtilityLabels
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.LexicalCategoriesSourceLangGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language in a monolingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLexicalCategoriesSourceLangGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LexicalCategoriesSourceLangLexicalTargetLangLexicalGet

> UtilityLabels LexicalCategoriesSourceLangLexicalTargetLangLexicalGet(ctx, sourceLangLexical, targetLangLexical).AppId(appId).AppKey(appKey).Execute()

Lists available lexical categories in a bilingual dataset



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
    sourceLangLexical := "sourceLangLexical_example" // string | Language code of the source language in a bilingual dataset.
    targetLangLexical := "targetLangLexical_example" // string | Language code of the target language in a bilingual dataset.
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.LexicalCategoriesSourceLangLexicalTargetLangLexicalGet(context.Background(), sourceLangLexical, targetLangLexical).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.LexicalCategoriesSourceLangLexicalTargetLangLexicalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LexicalCategoriesSourceLangLexicalTargetLangLexicalGet`: UtilityLabels
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.LexicalCategoriesSourceLangLexicalTargetLangLexicalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLangLexical** | **string** | Language code of the source language in a bilingual dataset. | 
**targetLangLexical** | **string** | Language code of the target language in a bilingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLexicalCategoriesSourceLangLexicalTargetLangLexicalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistersSourceLangGet

> UtilityLabels RegistersSourceLangGet(ctx, sourceLang).AppId(appId).AppKey(appKey).Execute()

Lists available registers in a  monolingual dataset



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
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.RegistersSourceLangGet(context.Background(), sourceLang).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.RegistersSourceLangGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistersSourceLangGet`: UtilityLabels
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.RegistersSourceLangGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLang** | **string** | Language code of the source language in a monolingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistersSourceLangGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistersSourceLangRegistersTargetLangRegistersGet

> UtilityLabels RegistersSourceLangRegistersTargetLangRegistersGet(ctx, sourceLangRegisters, targetLangRegisters).AppId(appId).AppKey(appKey).Execute()

Lists available registers in a bilingual dataset



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
    sourceLangRegisters := "sourceLangRegisters_example" // string | Language code of the source language in a bilingual dataset.
    targetLangRegisters := "targetLangRegisters_example" // string | Language code of the target language in a bilingual dataset.
    appId := "appId_example" // string | App ID Authentication Parameter
    appKey := "appKey_example" // string | App Key Authentication Parameter

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilityApi.RegistersSourceLangRegistersTargetLangRegistersGet(context.Background(), sourceLangRegisters, targetLangRegisters).AppId(appId).AppKey(appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.RegistersSourceLangRegistersTargetLangRegistersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistersSourceLangRegistersTargetLangRegistersGet`: UtilityLabels
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.RegistersSourceLangRegistersTargetLangRegistersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceLangRegisters** | **string** | Language code of the source language in a bilingual dataset. | 
**targetLangRegisters** | **string** | Language code of the target language in a bilingual dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistersSourceLangRegistersTargetLangRegistersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appId** | **string** | App ID Authentication Parameter | 
 **appKey** | **string** | App Key Authentication Parameter | 

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

