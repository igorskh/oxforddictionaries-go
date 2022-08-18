# HeadwordInflection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of a word | 
**Language** | **string** | IANA language code | 
**LexicalEntries** | [**[]InflectionLexicalEntry**](InflectionLexicalEntry.md) | A grouping of various senses in a specific language, and a lexical category that relates to a word | 
**Text** | Pointer to **string** | A given written or spoken realisation of an entry. | [optional] 

## Methods

### NewHeadwordInflection

`func NewHeadwordInflection(id string, language string, lexicalEntries []InflectionLexicalEntry, ) *HeadwordInflection`

NewHeadwordInflection instantiates a new HeadwordInflection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadwordInflectionWithDefaults

`func NewHeadwordInflectionWithDefaults() *HeadwordInflection`

NewHeadwordInflectionWithDefaults instantiates a new HeadwordInflection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HeadwordInflection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeadwordInflection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeadwordInflection) SetId(v string)`

SetId sets Id field to given value.


### GetLanguage

`func (o *HeadwordInflection) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *HeadwordInflection) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *HeadwordInflection) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLexicalEntries

`func (o *HeadwordInflection) GetLexicalEntries() []InflectionLexicalEntry`

GetLexicalEntries returns the LexicalEntries field if non-nil, zero value otherwise.

### GetLexicalEntriesOk

`func (o *HeadwordInflection) GetLexicalEntriesOk() (*[]InflectionLexicalEntry, bool)`

GetLexicalEntriesOk returns a tuple with the LexicalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicalEntries

`func (o *HeadwordInflection) SetLexicalEntries(v []InflectionLexicalEntry)`

SetLexicalEntries sets LexicalEntries field to given value.


### GetText

`func (o *HeadwordInflection) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *HeadwordInflection) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *HeadwordInflection) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *HeadwordInflection) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


