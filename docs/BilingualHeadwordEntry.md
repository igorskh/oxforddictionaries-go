# BilingualHeadwordEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of a word | 
**Language** | **string** | IANA language code | 
**LexicalEntries** | [**[]BilingualLexicalEntry**](BilingualLexicalEntry.md) | A grouping of various senses containing translations in a specific language, and a lexical category that relates to a word | 
**Pronunciations** | Pointer to **[]map[string]interface{}** | A list of possible pronunciations of a word | [optional] 
**Type** | Pointer to **string** | The json object type. Could be &#39;headword&#39;, &#39;inflection&#39; or &#39;phrase&#39; | [optional] 
**Word** | **string** | (DEPRECATED) A given written or spoken realisation of an entry, lowercased. | 

## Methods

### NewBilingualHeadwordEntry

`func NewBilingualHeadwordEntry(id string, language string, lexicalEntries []BilingualLexicalEntry, word string, ) *BilingualHeadwordEntry`

NewBilingualHeadwordEntry instantiates a new BilingualHeadwordEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBilingualHeadwordEntryWithDefaults

`func NewBilingualHeadwordEntryWithDefaults() *BilingualHeadwordEntry`

NewBilingualHeadwordEntryWithDefaults instantiates a new BilingualHeadwordEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BilingualHeadwordEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BilingualHeadwordEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BilingualHeadwordEntry) SetId(v string)`

SetId sets Id field to given value.


### GetLanguage

`func (o *BilingualHeadwordEntry) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BilingualHeadwordEntry) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BilingualHeadwordEntry) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLexicalEntries

`func (o *BilingualHeadwordEntry) GetLexicalEntries() []BilingualLexicalEntry`

GetLexicalEntries returns the LexicalEntries field if non-nil, zero value otherwise.

### GetLexicalEntriesOk

`func (o *BilingualHeadwordEntry) GetLexicalEntriesOk() (*[]BilingualLexicalEntry, bool)`

GetLexicalEntriesOk returns a tuple with the LexicalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicalEntries

`func (o *BilingualHeadwordEntry) SetLexicalEntries(v []BilingualLexicalEntry)`

SetLexicalEntries sets LexicalEntries field to given value.


### GetPronunciations

`func (o *BilingualHeadwordEntry) GetPronunciations() []map[string]interface{}`

GetPronunciations returns the Pronunciations field if non-nil, zero value otherwise.

### GetPronunciationsOk

`func (o *BilingualHeadwordEntry) GetPronunciationsOk() (*[]map[string]interface{}, bool)`

GetPronunciationsOk returns a tuple with the Pronunciations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronunciations

`func (o *BilingualHeadwordEntry) SetPronunciations(v []map[string]interface{})`

SetPronunciations sets Pronunciations field to given value.

### HasPronunciations

`func (o *BilingualHeadwordEntry) HasPronunciations() bool`

HasPronunciations returns a boolean if a field has been set.

### GetType

`func (o *BilingualHeadwordEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BilingualHeadwordEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BilingualHeadwordEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BilingualHeadwordEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWord

`func (o *BilingualHeadwordEntry) GetWord() string`

GetWord returns the Word field if non-nil, zero value otherwise.

### GetWordOk

`func (o *BilingualHeadwordEntry) GetWordOk() (*string, bool)`

GetWordOk returns a tuple with the Word field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWord

`func (o *BilingualHeadwordEntry) SetWord(v string)`

SetWord sets Word field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


