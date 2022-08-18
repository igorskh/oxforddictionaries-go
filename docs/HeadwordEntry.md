# HeadwordEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of a word | 
**Language** | **string** | IANA language code | 
**LexicalEntries** | [**[]LexicalEntry**](LexicalEntry.md) | A grouping of various senses in a specific language, and a lexical category that relates to a word | 
**Pronunciations** | Pointer to **[]map[string]interface{}** | A list of possible pronunciations of a word | [optional] 
**Type** | Pointer to **string** | The json object type. Could be &#39;headword&#39;, &#39;inflection&#39; or &#39;phrase&#39; | [optional] 
**Word** | **string** | (DEPRECATED) A given written or spoken realisation of an entry, lowercased. | 

## Methods

### NewHeadwordEntry

`func NewHeadwordEntry(id string, language string, lexicalEntries []LexicalEntry, word string, ) *HeadwordEntry`

NewHeadwordEntry instantiates a new HeadwordEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadwordEntryWithDefaults

`func NewHeadwordEntryWithDefaults() *HeadwordEntry`

NewHeadwordEntryWithDefaults instantiates a new HeadwordEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HeadwordEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeadwordEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeadwordEntry) SetId(v string)`

SetId sets Id field to given value.


### GetLanguage

`func (o *HeadwordEntry) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *HeadwordEntry) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *HeadwordEntry) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLexicalEntries

`func (o *HeadwordEntry) GetLexicalEntries() []LexicalEntry`

GetLexicalEntries returns the LexicalEntries field if non-nil, zero value otherwise.

### GetLexicalEntriesOk

`func (o *HeadwordEntry) GetLexicalEntriesOk() (*[]LexicalEntry, bool)`

GetLexicalEntriesOk returns a tuple with the LexicalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicalEntries

`func (o *HeadwordEntry) SetLexicalEntries(v []LexicalEntry)`

SetLexicalEntries sets LexicalEntries field to given value.


### GetPronunciations

`func (o *HeadwordEntry) GetPronunciations() []map[string]interface{}`

GetPronunciations returns the Pronunciations field if non-nil, zero value otherwise.

### GetPronunciationsOk

`func (o *HeadwordEntry) GetPronunciationsOk() (*[]map[string]interface{}, bool)`

GetPronunciationsOk returns a tuple with the Pronunciations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronunciations

`func (o *HeadwordEntry) SetPronunciations(v []map[string]interface{})`

SetPronunciations sets Pronunciations field to given value.

### HasPronunciations

`func (o *HeadwordEntry) HasPronunciations() bool`

HasPronunciations returns a boolean if a field has been set.

### GetType

`func (o *HeadwordEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeadwordEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeadwordEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HeadwordEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWord

`func (o *HeadwordEntry) GetWord() string`

GetWord returns the Word field if non-nil, zero value otherwise.

### GetWordOk

`func (o *HeadwordEntry) GetWordOk() (*string, bool)`

GetWordOk returns a tuple with the Word field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWord

`func (o *HeadwordEntry) SetWord(v string)`

SetWord sets Word field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


