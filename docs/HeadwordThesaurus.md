# HeadwordThesaurus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of a word | 
**Language** | **string** | IANA language code | 
**LexicalEntries** | [**[]ThesaurusLexicalEntry**](ThesaurusLexicalEntry.md) | A grouping of various senses in a specific language, and a lexical category that relates to a word | 
**Type** | Pointer to **string** | The json object type. Could be &#39;headword&#39;, &#39;inflection&#39; or &#39;phrase&#39; | [optional] 
**Word** | **string** | (DEPRECATED) A given written or spoken realisation of an entry, lowercased. | 

## Methods

### NewHeadwordThesaurus

`func NewHeadwordThesaurus(id string, language string, lexicalEntries []ThesaurusLexicalEntry, word string, ) *HeadwordThesaurus`

NewHeadwordThesaurus instantiates a new HeadwordThesaurus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadwordThesaurusWithDefaults

`func NewHeadwordThesaurusWithDefaults() *HeadwordThesaurus`

NewHeadwordThesaurusWithDefaults instantiates a new HeadwordThesaurus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HeadwordThesaurus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeadwordThesaurus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeadwordThesaurus) SetId(v string)`

SetId sets Id field to given value.


### GetLanguage

`func (o *HeadwordThesaurus) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *HeadwordThesaurus) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *HeadwordThesaurus) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLexicalEntries

`func (o *HeadwordThesaurus) GetLexicalEntries() []ThesaurusLexicalEntry`

GetLexicalEntries returns the LexicalEntries field if non-nil, zero value otherwise.

### GetLexicalEntriesOk

`func (o *HeadwordThesaurus) GetLexicalEntriesOk() (*[]ThesaurusLexicalEntry, bool)`

GetLexicalEntriesOk returns a tuple with the LexicalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicalEntries

`func (o *HeadwordThesaurus) SetLexicalEntries(v []ThesaurusLexicalEntry)`

SetLexicalEntries sets LexicalEntries field to given value.


### GetType

`func (o *HeadwordThesaurus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeadwordThesaurus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeadwordThesaurus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HeadwordThesaurus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWord

`func (o *HeadwordThesaurus) GetWord() string`

GetWord returns the Word field if non-nil, zero value otherwise.

### GetWordOk

`func (o *HeadwordThesaurus) GetWordOk() (*string, bool)`

GetWordOk returns a tuple with the Word field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWord

`func (o *HeadwordThesaurus) SetWord(v string)`

SetWord sets Word field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


