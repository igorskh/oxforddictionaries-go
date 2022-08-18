# LemmatronLexicalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrammaticalFeatures** | Pointer to **[]map[string]interface{}** | The different forms are correlated with meanings or functions which we text as &#39;features&#39; | [optional] 
**InflectionOf** | **[]map[string]interface{}** | A grouping of the modifications of a word to express different grammatical categories | 
**Language** | **string** | IANA language code | 
**LexicalCategory** | [**LexicalCategory**](LexicalCategory.md) |  | 
**Text** | **string** | A given written or spoken realisation of an entry. | 

## Methods

### NewLemmatronLexicalEntry

`func NewLemmatronLexicalEntry(inflectionOf []map[string]interface{}, language string, lexicalCategory LexicalCategory, text string, ) *LemmatronLexicalEntry`

NewLemmatronLexicalEntry instantiates a new LemmatronLexicalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLemmatronLexicalEntryWithDefaults

`func NewLemmatronLexicalEntryWithDefaults() *LemmatronLexicalEntry`

NewLemmatronLexicalEntryWithDefaults instantiates a new LemmatronLexicalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrammaticalFeatures

`func (o *LemmatronLexicalEntry) GetGrammaticalFeatures() []map[string]interface{}`

GetGrammaticalFeatures returns the GrammaticalFeatures field if non-nil, zero value otherwise.

### GetGrammaticalFeaturesOk

`func (o *LemmatronLexicalEntry) GetGrammaticalFeaturesOk() (*[]map[string]interface{}, bool)`

GetGrammaticalFeaturesOk returns a tuple with the GrammaticalFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrammaticalFeatures

`func (o *LemmatronLexicalEntry) SetGrammaticalFeatures(v []map[string]interface{})`

SetGrammaticalFeatures sets GrammaticalFeatures field to given value.

### HasGrammaticalFeatures

`func (o *LemmatronLexicalEntry) HasGrammaticalFeatures() bool`

HasGrammaticalFeatures returns a boolean if a field has been set.

### GetInflectionOf

`func (o *LemmatronLexicalEntry) GetInflectionOf() []map[string]interface{}`

GetInflectionOf returns the InflectionOf field if non-nil, zero value otherwise.

### GetInflectionOfOk

`func (o *LemmatronLexicalEntry) GetInflectionOfOk() (*[]map[string]interface{}, bool)`

GetInflectionOfOk returns a tuple with the InflectionOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflectionOf

`func (o *LemmatronLexicalEntry) SetInflectionOf(v []map[string]interface{})`

SetInflectionOf sets InflectionOf field to given value.


### GetLanguage

`func (o *LemmatronLexicalEntry) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *LemmatronLexicalEntry) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *LemmatronLexicalEntry) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLexicalCategory

`func (o *LemmatronLexicalEntry) GetLexicalCategory() LexicalCategory`

GetLexicalCategory returns the LexicalCategory field if non-nil, zero value otherwise.

### GetLexicalCategoryOk

`func (o *LemmatronLexicalEntry) GetLexicalCategoryOk() (*LexicalCategory, bool)`

GetLexicalCategoryOk returns a tuple with the LexicalCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicalCategory

`func (o *LemmatronLexicalEntry) SetLexicalCategory(v LexicalCategory)`

SetLexicalCategory sets LexicalCategory field to given value.


### GetText

`func (o *LemmatronLexicalEntry) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *LemmatronLexicalEntry) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *LemmatronLexicalEntry) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


