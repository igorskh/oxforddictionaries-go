# InflectedForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to [**DomainsList**](DomainsList.md) |  | [optional] 
**GrammaticalFeatures** | Pointer to **[]map[string]interface{}** | The different forms are correlated with meanings or functions which we text as &#39;features&#39; | [optional] 
**InflectedForm** | **string** | Canonical form of an inflection | 
**LexicalCategory** | Pointer to [**LexicalCategory**](LexicalCategory.md) |  | [optional] 
**Pronunciations** | Pointer to **[]map[string]interface{}** | A list of possible pronunciations of a word | [optional] 
**Regions** | Pointer to [**RegionsList**](RegionsList.md) |  | [optional] 
**Registers** | Pointer to [**RegistersList**](RegistersList.md) |  | [optional] 

## Methods

### NewInflectedForm

`func NewInflectedForm(inflectedForm string, ) *InflectedForm`

NewInflectedForm instantiates a new InflectedForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInflectedFormWithDefaults

`func NewInflectedFormWithDefaults() *InflectedForm`

NewInflectedFormWithDefaults instantiates a new InflectedForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *InflectedForm) GetDomains() DomainsList`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *InflectedForm) GetDomainsOk() (*DomainsList, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *InflectedForm) SetDomains(v DomainsList)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *InflectedForm) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetGrammaticalFeatures

`func (o *InflectedForm) GetGrammaticalFeatures() []map[string]interface{}`

GetGrammaticalFeatures returns the GrammaticalFeatures field if non-nil, zero value otherwise.

### GetGrammaticalFeaturesOk

`func (o *InflectedForm) GetGrammaticalFeaturesOk() (*[]map[string]interface{}, bool)`

GetGrammaticalFeaturesOk returns a tuple with the GrammaticalFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrammaticalFeatures

`func (o *InflectedForm) SetGrammaticalFeatures(v []map[string]interface{})`

SetGrammaticalFeatures sets GrammaticalFeatures field to given value.

### HasGrammaticalFeatures

`func (o *InflectedForm) HasGrammaticalFeatures() bool`

HasGrammaticalFeatures returns a boolean if a field has been set.

### GetInflectedForm

`func (o *InflectedForm) GetInflectedForm() string`

GetInflectedForm returns the InflectedForm field if non-nil, zero value otherwise.

### GetInflectedFormOk

`func (o *InflectedForm) GetInflectedFormOk() (*string, bool)`

GetInflectedFormOk returns a tuple with the InflectedForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflectedForm

`func (o *InflectedForm) SetInflectedForm(v string)`

SetInflectedForm sets InflectedForm field to given value.


### GetLexicalCategory

`func (o *InflectedForm) GetLexicalCategory() LexicalCategory`

GetLexicalCategory returns the LexicalCategory field if non-nil, zero value otherwise.

### GetLexicalCategoryOk

`func (o *InflectedForm) GetLexicalCategoryOk() (*LexicalCategory, bool)`

GetLexicalCategoryOk returns a tuple with the LexicalCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicalCategory

`func (o *InflectedForm) SetLexicalCategory(v LexicalCategory)`

SetLexicalCategory sets LexicalCategory field to given value.

### HasLexicalCategory

`func (o *InflectedForm) HasLexicalCategory() bool`

HasLexicalCategory returns a boolean if a field has been set.

### GetPronunciations

`func (o *InflectedForm) GetPronunciations() []map[string]interface{}`

GetPronunciations returns the Pronunciations field if non-nil, zero value otherwise.

### GetPronunciationsOk

`func (o *InflectedForm) GetPronunciationsOk() (*[]map[string]interface{}, bool)`

GetPronunciationsOk returns a tuple with the Pronunciations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronunciations

`func (o *InflectedForm) SetPronunciations(v []map[string]interface{})`

SetPronunciations sets Pronunciations field to given value.

### HasPronunciations

`func (o *InflectedForm) HasPronunciations() bool`

HasPronunciations returns a boolean if a field has been set.

### GetRegions

`func (o *InflectedForm) GetRegions() RegionsList`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *InflectedForm) GetRegionsOk() (*RegionsList, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *InflectedForm) SetRegions(v RegionsList)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *InflectedForm) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRegisters

`func (o *InflectedForm) GetRegisters() RegistersList`

GetRegisters returns the Registers field if non-nil, zero value otherwise.

### GetRegistersOk

`func (o *InflectedForm) GetRegistersOk() (*RegistersList, bool)`

GetRegistersOk returns a tuple with the Registers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisters

`func (o *InflectedForm) SetRegisters(v RegistersList)`

SetRegisters sets Registers field to given value.

### HasRegisters

`func (o *InflectedForm) HasRegisters() bool`

HasRegisters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


