# BilingualSenseConstructions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to [**DomainsList**](DomainsList.md) |  | [optional] 
**Examples** | Pointer to **[][]string** |  | [optional] 
**Notes** | Pointer to **[]map[string]interface{}** | various types of notes that appear | [optional] 
**Regions** | Pointer to [**RegionsList**](RegionsList.md) |  | [optional] 
**Registers** | Pointer to [**RegistersList**](RegistersList.md) |  | [optional] 
**Text** | **string** | The construction text | 
**Translations** | Pointer to **[]map[string]interface{}** | A list of written or spoken rendering of the meaning of a word or text in another language(s) | [optional] 

## Methods

### NewBilingualSenseConstructions

`func NewBilingualSenseConstructions(text string, ) *BilingualSenseConstructions`

NewBilingualSenseConstructions instantiates a new BilingualSenseConstructions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBilingualSenseConstructionsWithDefaults

`func NewBilingualSenseConstructionsWithDefaults() *BilingualSenseConstructions`

NewBilingualSenseConstructionsWithDefaults instantiates a new BilingualSenseConstructions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *BilingualSenseConstructions) GetDomains() DomainsList`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *BilingualSenseConstructions) GetDomainsOk() (*DomainsList, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *BilingualSenseConstructions) SetDomains(v DomainsList)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *BilingualSenseConstructions) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetExamples

`func (o *BilingualSenseConstructions) GetExamples() [][]string`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *BilingualSenseConstructions) GetExamplesOk() (*[][]string, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *BilingualSenseConstructions) SetExamples(v [][]string)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *BilingualSenseConstructions) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetNotes

`func (o *BilingualSenseConstructions) GetNotes() []map[string]interface{}`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BilingualSenseConstructions) GetNotesOk() (*[]map[string]interface{}, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BilingualSenseConstructions) SetNotes(v []map[string]interface{})`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BilingualSenseConstructions) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRegions

`func (o *BilingualSenseConstructions) GetRegions() RegionsList`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *BilingualSenseConstructions) GetRegionsOk() (*RegionsList, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *BilingualSenseConstructions) SetRegions(v RegionsList)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *BilingualSenseConstructions) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRegisters

`func (o *BilingualSenseConstructions) GetRegisters() RegistersList`

GetRegisters returns the Registers field if non-nil, zero value otherwise.

### GetRegistersOk

`func (o *BilingualSenseConstructions) GetRegistersOk() (*RegistersList, bool)`

GetRegistersOk returns a tuple with the Registers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisters

`func (o *BilingualSenseConstructions) SetRegisters(v RegistersList)`

SetRegisters sets Registers field to given value.

### HasRegisters

`func (o *BilingualSenseConstructions) HasRegisters() bool`

HasRegisters returns a boolean if a field has been set.

### GetText

`func (o *BilingualSenseConstructions) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *BilingualSenseConstructions) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *BilingualSenseConstructions) SetText(v string)`

SetText sets Text field to given value.


### GetTranslations

`func (o *BilingualSenseConstructions) GetTranslations() []map[string]interface{}`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *BilingualSenseConstructions) GetTranslationsOk() (*[]map[string]interface{}, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *BilingualSenseConstructions) SetTranslations(v []map[string]interface{})`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *BilingualSenseConstructions) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


