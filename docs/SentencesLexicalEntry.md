# SentencesLexicalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrammaticalFeatures** | Pointer to **[]map[string]interface{}** | The different forms are correlated with meanings or functions which we text as &#39;features&#39; | [optional] 
**Language** | **string** | IANA language code | 
**LexicalCategory** | Pointer to [**LexicalCategory**](LexicalCategory.md) |  | [optional] 
**Sentences** | **[]map[string]interface{}** | A list of written or spoken rendering of examples of use of a word or text | 
**Text** | **string** | A given written or spoken realisation of an entry. | 

## Methods

### NewSentencesLexicalEntry

`func NewSentencesLexicalEntry(language string, sentences []map[string]interface{}, text string, ) *SentencesLexicalEntry`

NewSentencesLexicalEntry instantiates a new SentencesLexicalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSentencesLexicalEntryWithDefaults

`func NewSentencesLexicalEntryWithDefaults() *SentencesLexicalEntry`

NewSentencesLexicalEntryWithDefaults instantiates a new SentencesLexicalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrammaticalFeatures

`func (o *SentencesLexicalEntry) GetGrammaticalFeatures() []map[string]interface{}`

GetGrammaticalFeatures returns the GrammaticalFeatures field if non-nil, zero value otherwise.

### GetGrammaticalFeaturesOk

`func (o *SentencesLexicalEntry) GetGrammaticalFeaturesOk() (*[]map[string]interface{}, bool)`

GetGrammaticalFeaturesOk returns a tuple with the GrammaticalFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrammaticalFeatures

`func (o *SentencesLexicalEntry) SetGrammaticalFeatures(v []map[string]interface{})`

SetGrammaticalFeatures sets GrammaticalFeatures field to given value.

### HasGrammaticalFeatures

`func (o *SentencesLexicalEntry) HasGrammaticalFeatures() bool`

HasGrammaticalFeatures returns a boolean if a field has been set.

### GetLanguage

`func (o *SentencesLexicalEntry) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SentencesLexicalEntry) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SentencesLexicalEntry) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLexicalCategory

`func (o *SentencesLexicalEntry) GetLexicalCategory() LexicalCategory`

GetLexicalCategory returns the LexicalCategory field if non-nil, zero value otherwise.

### GetLexicalCategoryOk

`func (o *SentencesLexicalEntry) GetLexicalCategoryOk() (*LexicalCategory, bool)`

GetLexicalCategoryOk returns a tuple with the LexicalCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicalCategory

`func (o *SentencesLexicalEntry) SetLexicalCategory(v LexicalCategory)`

SetLexicalCategory sets LexicalCategory field to given value.

### HasLexicalCategory

`func (o *SentencesLexicalEntry) HasLexicalCategory() bool`

HasLexicalCategory returns a boolean if a field has been set.

### GetSentences

`func (o *SentencesLexicalEntry) GetSentences() []map[string]interface{}`

GetSentences returns the Sentences field if non-nil, zero value otherwise.

### GetSentencesOk

`func (o *SentencesLexicalEntry) GetSentencesOk() (*[]map[string]interface{}, bool)`

GetSentencesOk returns a tuple with the Sentences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentences

`func (o *SentencesLexicalEntry) SetSentences(v []map[string]interface{})`

SetSentences sets Sentences field to given value.


### GetText

`func (o *SentencesLexicalEntry) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SentencesLexicalEntry) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SentencesLexicalEntry) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


