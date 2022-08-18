# SenseConstructions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to [**DomainsList**](DomainsList.md) |  | [optional] 
**Examples** | Pointer to **[][]string** |  | [optional] 
**Notes** | Pointer to **[]map[string]interface{}** | various types of notes that appear | [optional] 
**Regions** | Pointer to [**RegionsList**](RegionsList.md) |  | [optional] 
**Registers** | Pointer to [**RegistersList**](RegistersList.md) |  | [optional] 
**Text** | **string** | The construction text | 

## Methods

### NewSenseConstructions

`func NewSenseConstructions(text string, ) *SenseConstructions`

NewSenseConstructions instantiates a new SenseConstructions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSenseConstructionsWithDefaults

`func NewSenseConstructionsWithDefaults() *SenseConstructions`

NewSenseConstructionsWithDefaults instantiates a new SenseConstructions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *SenseConstructions) GetDomains() DomainsList`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *SenseConstructions) GetDomainsOk() (*DomainsList, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *SenseConstructions) SetDomains(v DomainsList)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *SenseConstructions) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetExamples

`func (o *SenseConstructions) GetExamples() [][]string`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *SenseConstructions) GetExamplesOk() (*[][]string, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *SenseConstructions) SetExamples(v [][]string)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *SenseConstructions) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetNotes

`func (o *SenseConstructions) GetNotes() []map[string]interface{}`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SenseConstructions) GetNotesOk() (*[]map[string]interface{}, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SenseConstructions) SetNotes(v []map[string]interface{})`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SenseConstructions) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRegions

`func (o *SenseConstructions) GetRegions() RegionsList`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *SenseConstructions) GetRegionsOk() (*RegionsList, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *SenseConstructions) SetRegions(v RegionsList)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *SenseConstructions) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRegisters

`func (o *SenseConstructions) GetRegisters() RegistersList`

GetRegisters returns the Registers field if non-nil, zero value otherwise.

### GetRegistersOk

`func (o *SenseConstructions) GetRegistersOk() (*RegistersList, bool)`

GetRegistersOk returns a tuple with the Registers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisters

`func (o *SenseConstructions) SetRegisters(v RegistersList)`

SetRegisters sets Registers field to given value.

### HasRegisters

`func (o *SenseConstructions) HasRegisters() bool`

HasRegisters returns a boolean if a field has been set.

### GetText

`func (o *SenseConstructions) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SenseConstructions) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SenseConstructions) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


