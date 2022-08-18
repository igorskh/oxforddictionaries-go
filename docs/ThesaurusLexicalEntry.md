# ThesaurusLexicalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]ThesaurusEntry**](ThesaurusEntry.md) |  | [optional] 
**GrammaticalFeatures** | Pointer to **[]map[string]interface{}** | The different forms are correlated with meanings or functions which we text as &#39;features&#39; | [optional] 
**Language** | **string** | IANA language code | 
**LexicalCategory** | [**LexicalCategory**](LexicalCategory.md) |  | 
**Text** | **string** | A given written or spoken realisation of an entry. | 
**VariantForms** | Pointer to **[]map[string]interface{}** | Various words that are used interchangeably depending on the context, e.g &#39;aluminium&#39; and &#39;aluminum&#39; | [optional] 

## Methods

### NewThesaurusLexicalEntry

`func NewThesaurusLexicalEntry(language string, lexicalCategory LexicalCategory, text string, ) *ThesaurusLexicalEntry`

NewThesaurusLexicalEntry instantiates a new ThesaurusLexicalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThesaurusLexicalEntryWithDefaults

`func NewThesaurusLexicalEntryWithDefaults() *ThesaurusLexicalEntry`

NewThesaurusLexicalEntryWithDefaults instantiates a new ThesaurusLexicalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *ThesaurusLexicalEntry) GetEntries() []ThesaurusEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ThesaurusLexicalEntry) GetEntriesOk() (*[]ThesaurusEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ThesaurusLexicalEntry) SetEntries(v []ThesaurusEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ThesaurusLexicalEntry) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetGrammaticalFeatures

`func (o *ThesaurusLexicalEntry) GetGrammaticalFeatures() []map[string]interface{}`

GetGrammaticalFeatures returns the GrammaticalFeatures field if non-nil, zero value otherwise.

### GetGrammaticalFeaturesOk

`func (o *ThesaurusLexicalEntry) GetGrammaticalFeaturesOk() (*[]map[string]interface{}, bool)`

GetGrammaticalFeaturesOk returns a tuple with the GrammaticalFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrammaticalFeatures

`func (o *ThesaurusLexicalEntry) SetGrammaticalFeatures(v []map[string]interface{})`

SetGrammaticalFeatures sets GrammaticalFeatures field to given value.

### HasGrammaticalFeatures

`func (o *ThesaurusLexicalEntry) HasGrammaticalFeatures() bool`

HasGrammaticalFeatures returns a boolean if a field has been set.

### GetLanguage

`func (o *ThesaurusLexicalEntry) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ThesaurusLexicalEntry) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ThesaurusLexicalEntry) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLexicalCategory

`func (o *ThesaurusLexicalEntry) GetLexicalCategory() LexicalCategory`

GetLexicalCategory returns the LexicalCategory field if non-nil, zero value otherwise.

### GetLexicalCategoryOk

`func (o *ThesaurusLexicalEntry) GetLexicalCategoryOk() (*LexicalCategory, bool)`

GetLexicalCategoryOk returns a tuple with the LexicalCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicalCategory

`func (o *ThesaurusLexicalEntry) SetLexicalCategory(v LexicalCategory)`

SetLexicalCategory sets LexicalCategory field to given value.


### GetText

`func (o *ThesaurusLexicalEntry) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ThesaurusLexicalEntry) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ThesaurusLexicalEntry) SetText(v string)`

SetText sets Text field to given value.


### GetVariantForms

`func (o *ThesaurusLexicalEntry) GetVariantForms() []map[string]interface{}`

GetVariantForms returns the VariantForms field if non-nil, zero value otherwise.

### GetVariantFormsOk

`func (o *ThesaurusLexicalEntry) GetVariantFormsOk() (*[]map[string]interface{}, bool)`

GetVariantFormsOk returns a tuple with the VariantForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantForms

`func (o *ThesaurusLexicalEntry) SetVariantForms(v []map[string]interface{})`

SetVariantForms sets VariantForms field to given value.

### HasVariantForms

`func (o *ThesaurusLexicalEntry) HasVariantForms() bool`

HasVariantForms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


