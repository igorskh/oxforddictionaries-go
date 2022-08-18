# InflectionLexicalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrammaticalFeatures** | Pointer to **[]map[string]interface{}** | The different forms are correlated with meanings or functions which we text as &#39;features&#39; | [optional] 
**Inflections** | Pointer to [**[]InflectedForm**](InflectedForm.md) | A list of inflected forms for a lexicalEntry. | [optional] 
**Language** | **string** | IANA language code | 
**LexicalCategory** | [**LexicalCategory**](LexicalCategory.md) |  | 

## Methods

### NewInflectionLexicalEntry

`func NewInflectionLexicalEntry(language string, lexicalCategory LexicalCategory, ) *InflectionLexicalEntry`

NewInflectionLexicalEntry instantiates a new InflectionLexicalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInflectionLexicalEntryWithDefaults

`func NewInflectionLexicalEntryWithDefaults() *InflectionLexicalEntry`

NewInflectionLexicalEntryWithDefaults instantiates a new InflectionLexicalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrammaticalFeatures

`func (o *InflectionLexicalEntry) GetGrammaticalFeatures() []map[string]interface{}`

GetGrammaticalFeatures returns the GrammaticalFeatures field if non-nil, zero value otherwise.

### GetGrammaticalFeaturesOk

`func (o *InflectionLexicalEntry) GetGrammaticalFeaturesOk() (*[]map[string]interface{}, bool)`

GetGrammaticalFeaturesOk returns a tuple with the GrammaticalFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrammaticalFeatures

`func (o *InflectionLexicalEntry) SetGrammaticalFeatures(v []map[string]interface{})`

SetGrammaticalFeatures sets GrammaticalFeatures field to given value.

### HasGrammaticalFeatures

`func (o *InflectionLexicalEntry) HasGrammaticalFeatures() bool`

HasGrammaticalFeatures returns a boolean if a field has been set.

### GetInflections

`func (o *InflectionLexicalEntry) GetInflections() []InflectedForm`

GetInflections returns the Inflections field if non-nil, zero value otherwise.

### GetInflectionsOk

`func (o *InflectionLexicalEntry) GetInflectionsOk() (*[]InflectedForm, bool)`

GetInflectionsOk returns a tuple with the Inflections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflections

`func (o *InflectionLexicalEntry) SetInflections(v []InflectedForm)`

SetInflections sets Inflections field to given value.

### HasInflections

`func (o *InflectionLexicalEntry) HasInflections() bool`

HasInflections returns a boolean if a field has been set.

### GetLanguage

`func (o *InflectionLexicalEntry) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *InflectionLexicalEntry) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *InflectionLexicalEntry) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLexicalCategory

`func (o *InflectionLexicalEntry) GetLexicalCategory() LexicalCategory`

GetLexicalCategory returns the LexicalCategory field if non-nil, zero value otherwise.

### GetLexicalCategoryOk

`func (o *InflectionLexicalEntry) GetLexicalCategoryOk() (*LexicalCategory, bool)`

GetLexicalCategoryOk returns a tuple with the LexicalCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicalCategory

`func (o *InflectionLexicalEntry) SetLexicalCategory(v LexicalCategory)`

SetLexicalCategory sets LexicalCategory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


