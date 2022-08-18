# LexicalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compounds** | Pointer to **[]map[string]interface{}** | A list of written or spoken words | [optional] 
**DerivativeOf** | Pointer to **[]map[string]interface{}** | A list of written or spoken words | [optional] 
**Derivatives** | Pointer to **[]map[string]interface{}** | A list of written or spoken words | [optional] 
**Entries** | Pointer to [**[]Entry**](Entry.md) |  | [optional] 
**GrammaticalFeatures** | Pointer to **[]map[string]interface{}** | The different forms are correlated with meanings or functions which we text as &#39;features&#39; | [optional] 
**Language** | **string** | IANA language code | 
**LexicalCategory** | [**LexicalCategory**](LexicalCategory.md) |  | 
**Notes** | Pointer to **[]map[string]interface{}** | various types of notes that appear | [optional] 
**PhrasalVerbs** | Pointer to **[]map[string]interface{}** | A list of written or spoken words | [optional] 
**Phrases** | Pointer to **[]map[string]interface{}** | A list of written or spoken words | [optional] 
**Pronunciations** | Pointer to **[]map[string]interface{}** | A list of possible pronunciations of a word | [optional] 
**Root** | Pointer to **string** | Abstract root form from which this lexicalEntry is derived. | [optional] 
**Text** | **string** | A given written or spoken realisation of an entry. | 
**VariantForms** | Pointer to **[]map[string]interface{}** | Various words that are used interchangeably depending on the context, e.g &#39;aluminium&#39; and &#39;aluminum&#39; | [optional] 

## Methods

### NewLexicalEntry

`func NewLexicalEntry(language string, lexicalCategory LexicalCategory, text string, ) *LexicalEntry`

NewLexicalEntry instantiates a new LexicalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLexicalEntryWithDefaults

`func NewLexicalEntryWithDefaults() *LexicalEntry`

NewLexicalEntryWithDefaults instantiates a new LexicalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompounds

`func (o *LexicalEntry) GetCompounds() []map[string]interface{}`

GetCompounds returns the Compounds field if non-nil, zero value otherwise.

### GetCompoundsOk

`func (o *LexicalEntry) GetCompoundsOk() (*[]map[string]interface{}, bool)`

GetCompoundsOk returns a tuple with the Compounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompounds

`func (o *LexicalEntry) SetCompounds(v []map[string]interface{})`

SetCompounds sets Compounds field to given value.

### HasCompounds

`func (o *LexicalEntry) HasCompounds() bool`

HasCompounds returns a boolean if a field has been set.

### GetDerivativeOf

`func (o *LexicalEntry) GetDerivativeOf() []map[string]interface{}`

GetDerivativeOf returns the DerivativeOf field if non-nil, zero value otherwise.

### GetDerivativeOfOk

`func (o *LexicalEntry) GetDerivativeOfOk() (*[]map[string]interface{}, bool)`

GetDerivativeOfOk returns a tuple with the DerivativeOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeOf

`func (o *LexicalEntry) SetDerivativeOf(v []map[string]interface{})`

SetDerivativeOf sets DerivativeOf field to given value.

### HasDerivativeOf

`func (o *LexicalEntry) HasDerivativeOf() bool`

HasDerivativeOf returns a boolean if a field has been set.

### GetDerivatives

`func (o *LexicalEntry) GetDerivatives() []map[string]interface{}`

GetDerivatives returns the Derivatives field if non-nil, zero value otherwise.

### GetDerivativesOk

`func (o *LexicalEntry) GetDerivativesOk() (*[]map[string]interface{}, bool)`

GetDerivativesOk returns a tuple with the Derivatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivatives

`func (o *LexicalEntry) SetDerivatives(v []map[string]interface{})`

SetDerivatives sets Derivatives field to given value.

### HasDerivatives

`func (o *LexicalEntry) HasDerivatives() bool`

HasDerivatives returns a boolean if a field has been set.

### GetEntries

`func (o *LexicalEntry) GetEntries() []Entry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *LexicalEntry) GetEntriesOk() (*[]Entry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *LexicalEntry) SetEntries(v []Entry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *LexicalEntry) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetGrammaticalFeatures

`func (o *LexicalEntry) GetGrammaticalFeatures() []map[string]interface{}`

GetGrammaticalFeatures returns the GrammaticalFeatures field if non-nil, zero value otherwise.

### GetGrammaticalFeaturesOk

`func (o *LexicalEntry) GetGrammaticalFeaturesOk() (*[]map[string]interface{}, bool)`

GetGrammaticalFeaturesOk returns a tuple with the GrammaticalFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrammaticalFeatures

`func (o *LexicalEntry) SetGrammaticalFeatures(v []map[string]interface{})`

SetGrammaticalFeatures sets GrammaticalFeatures field to given value.

### HasGrammaticalFeatures

`func (o *LexicalEntry) HasGrammaticalFeatures() bool`

HasGrammaticalFeatures returns a boolean if a field has been set.

### GetLanguage

`func (o *LexicalEntry) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *LexicalEntry) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *LexicalEntry) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLexicalCategory

`func (o *LexicalEntry) GetLexicalCategory() LexicalCategory`

GetLexicalCategory returns the LexicalCategory field if non-nil, zero value otherwise.

### GetLexicalCategoryOk

`func (o *LexicalEntry) GetLexicalCategoryOk() (*LexicalCategory, bool)`

GetLexicalCategoryOk returns a tuple with the LexicalCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicalCategory

`func (o *LexicalEntry) SetLexicalCategory(v LexicalCategory)`

SetLexicalCategory sets LexicalCategory field to given value.


### GetNotes

`func (o *LexicalEntry) GetNotes() []map[string]interface{}`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *LexicalEntry) GetNotesOk() (*[]map[string]interface{}, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *LexicalEntry) SetNotes(v []map[string]interface{})`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *LexicalEntry) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPhrasalVerbs

`func (o *LexicalEntry) GetPhrasalVerbs() []map[string]interface{}`

GetPhrasalVerbs returns the PhrasalVerbs field if non-nil, zero value otherwise.

### GetPhrasalVerbsOk

`func (o *LexicalEntry) GetPhrasalVerbsOk() (*[]map[string]interface{}, bool)`

GetPhrasalVerbsOk returns a tuple with the PhrasalVerbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhrasalVerbs

`func (o *LexicalEntry) SetPhrasalVerbs(v []map[string]interface{})`

SetPhrasalVerbs sets PhrasalVerbs field to given value.

### HasPhrasalVerbs

`func (o *LexicalEntry) HasPhrasalVerbs() bool`

HasPhrasalVerbs returns a boolean if a field has been set.

### GetPhrases

`func (o *LexicalEntry) GetPhrases() []map[string]interface{}`

GetPhrases returns the Phrases field if non-nil, zero value otherwise.

### GetPhrasesOk

`func (o *LexicalEntry) GetPhrasesOk() (*[]map[string]interface{}, bool)`

GetPhrasesOk returns a tuple with the Phrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhrases

`func (o *LexicalEntry) SetPhrases(v []map[string]interface{})`

SetPhrases sets Phrases field to given value.

### HasPhrases

`func (o *LexicalEntry) HasPhrases() bool`

HasPhrases returns a boolean if a field has been set.

### GetPronunciations

`func (o *LexicalEntry) GetPronunciations() []map[string]interface{}`

GetPronunciations returns the Pronunciations field if non-nil, zero value otherwise.

### GetPronunciationsOk

`func (o *LexicalEntry) GetPronunciationsOk() (*[]map[string]interface{}, bool)`

GetPronunciationsOk returns a tuple with the Pronunciations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronunciations

`func (o *LexicalEntry) SetPronunciations(v []map[string]interface{})`

SetPronunciations sets Pronunciations field to given value.

### HasPronunciations

`func (o *LexicalEntry) HasPronunciations() bool`

HasPronunciations returns a boolean if a field has been set.

### GetRoot

`func (o *LexicalEntry) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *LexicalEntry) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *LexicalEntry) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *LexicalEntry) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetText

`func (o *LexicalEntry) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *LexicalEntry) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *LexicalEntry) SetText(v string)`

SetText sets Text field to given value.


### GetVariantForms

`func (o *LexicalEntry) GetVariantForms() []map[string]interface{}`

GetVariantForms returns the VariantForms field if non-nil, zero value otherwise.

### GetVariantFormsOk

`func (o *LexicalEntry) GetVariantFormsOk() (*[]map[string]interface{}, bool)`

GetVariantFormsOk returns a tuple with the VariantForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantForms

`func (o *LexicalEntry) SetVariantForms(v []map[string]interface{})`

SetVariantForms sets VariantForms field to given value.

### HasVariantForms

`func (o *LexicalEntry) HasVariantForms() bool`

HasVariantForms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


