# ThesaurusSense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Antonyms** | Pointer to **[]map[string]interface{}** |  | [optional] 
**DomainClasses** | Pointer to [**DomainClassesList**](DomainClassesList.md) |  | [optional] 
**Domains** | Pointer to [**DomainsList**](DomainsList.md) |  | [optional] 
**Examples** | Pointer to **[]map[string]interface{}** | A list of written or spoken rendering of examples of use of a word or text | [optional] 
**Id** | Pointer to **string** | The id of the sense that is required for the delete procedure | [optional] 
**Regions** | Pointer to [**RegionsList**](RegionsList.md) |  | [optional] 
**Registers** | Pointer to [**RegistersList**](RegistersList.md) |  | [optional] 
**SemanticClasses** | Pointer to [**SemanticClassesList**](SemanticClassesList.md) |  | [optional] 
**Subsenses** | Pointer to [**[]ThesaurusSense**](ThesaurusSense.md) | subsenses of word | [optional] 
**Synonyms** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewThesaurusSense

`func NewThesaurusSense() *ThesaurusSense`

NewThesaurusSense instantiates a new ThesaurusSense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThesaurusSenseWithDefaults

`func NewThesaurusSenseWithDefaults() *ThesaurusSense`

NewThesaurusSenseWithDefaults instantiates a new ThesaurusSense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntonyms

`func (o *ThesaurusSense) GetAntonyms() []map[string]interface{}`

GetAntonyms returns the Antonyms field if non-nil, zero value otherwise.

### GetAntonymsOk

`func (o *ThesaurusSense) GetAntonymsOk() (*[]map[string]interface{}, bool)`

GetAntonymsOk returns a tuple with the Antonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntonyms

`func (o *ThesaurusSense) SetAntonyms(v []map[string]interface{})`

SetAntonyms sets Antonyms field to given value.

### HasAntonyms

`func (o *ThesaurusSense) HasAntonyms() bool`

HasAntonyms returns a boolean if a field has been set.

### GetDomainClasses

`func (o *ThesaurusSense) GetDomainClasses() DomainClassesList`

GetDomainClasses returns the DomainClasses field if non-nil, zero value otherwise.

### GetDomainClassesOk

`func (o *ThesaurusSense) GetDomainClassesOk() (*DomainClassesList, bool)`

GetDomainClassesOk returns a tuple with the DomainClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClasses

`func (o *ThesaurusSense) SetDomainClasses(v DomainClassesList)`

SetDomainClasses sets DomainClasses field to given value.

### HasDomainClasses

`func (o *ThesaurusSense) HasDomainClasses() bool`

HasDomainClasses returns a boolean if a field has been set.

### GetDomains

`func (o *ThesaurusSense) GetDomains() DomainsList`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ThesaurusSense) GetDomainsOk() (*DomainsList, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ThesaurusSense) SetDomains(v DomainsList)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ThesaurusSense) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetExamples

`func (o *ThesaurusSense) GetExamples() []map[string]interface{}`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *ThesaurusSense) GetExamplesOk() (*[]map[string]interface{}, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *ThesaurusSense) SetExamples(v []map[string]interface{})`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *ThesaurusSense) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetId

`func (o *ThesaurusSense) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThesaurusSense) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThesaurusSense) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThesaurusSense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegions

`func (o *ThesaurusSense) GetRegions() RegionsList`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ThesaurusSense) GetRegionsOk() (*RegionsList, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ThesaurusSense) SetRegions(v RegionsList)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ThesaurusSense) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRegisters

`func (o *ThesaurusSense) GetRegisters() RegistersList`

GetRegisters returns the Registers field if non-nil, zero value otherwise.

### GetRegistersOk

`func (o *ThesaurusSense) GetRegistersOk() (*RegistersList, bool)`

GetRegistersOk returns a tuple with the Registers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisters

`func (o *ThesaurusSense) SetRegisters(v RegistersList)`

SetRegisters sets Registers field to given value.

### HasRegisters

`func (o *ThesaurusSense) HasRegisters() bool`

HasRegisters returns a boolean if a field has been set.

### GetSemanticClasses

`func (o *ThesaurusSense) GetSemanticClasses() SemanticClassesList`

GetSemanticClasses returns the SemanticClasses field if non-nil, zero value otherwise.

### GetSemanticClassesOk

`func (o *ThesaurusSense) GetSemanticClassesOk() (*SemanticClassesList, bool)`

GetSemanticClassesOk returns a tuple with the SemanticClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticClasses

`func (o *ThesaurusSense) SetSemanticClasses(v SemanticClassesList)`

SetSemanticClasses sets SemanticClasses field to given value.

### HasSemanticClasses

`func (o *ThesaurusSense) HasSemanticClasses() bool`

HasSemanticClasses returns a boolean if a field has been set.

### GetSubsenses

`func (o *ThesaurusSense) GetSubsenses() []ThesaurusSense`

GetSubsenses returns the Subsenses field if non-nil, zero value otherwise.

### GetSubsensesOk

`func (o *ThesaurusSense) GetSubsensesOk() (*[]ThesaurusSense, bool)`

GetSubsensesOk returns a tuple with the Subsenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsenses

`func (o *ThesaurusSense) SetSubsenses(v []ThesaurusSense)`

SetSubsenses sets Subsenses field to given value.

### HasSubsenses

`func (o *ThesaurusSense) HasSubsenses() bool`

HasSubsenses returns a boolean if a field has been set.

### GetSynonyms

`func (o *ThesaurusSense) GetSynonyms() []map[string]interface{}`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *ThesaurusSense) GetSynonymsOk() (*[]map[string]interface{}, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *ThesaurusSense) SetSynonyms(v []map[string]interface{})`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *ThesaurusSense) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


