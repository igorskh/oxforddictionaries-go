# LanguagesResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | Name of region. | [optional] 
**Source** | Pointer to **string** | Name of source dictionary. | [optional] 
**SourceLanguage** | Pointer to [**LanguagesSourceLanguage**](LanguagesSourceLanguage.md) |  | [optional] 
**TargetLanguage** | Pointer to [**LanguagesTargetLanguage**](LanguagesTargetLanguage.md) |  | [optional] 
**Type** | Pointer to **string** | whether monolingual or bilingual. | [optional] 

## Methods

### NewLanguagesResults

`func NewLanguagesResults() *LanguagesResults`

NewLanguagesResults instantiates a new LanguagesResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguagesResultsWithDefaults

`func NewLanguagesResultsWithDefaults() *LanguagesResults`

NewLanguagesResultsWithDefaults instantiates a new LanguagesResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *LanguagesResults) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LanguagesResults) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LanguagesResults) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LanguagesResults) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSource

`func (o *LanguagesResults) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LanguagesResults) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LanguagesResults) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LanguagesResults) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceLanguage

`func (o *LanguagesResults) GetSourceLanguage() LanguagesSourceLanguage`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *LanguagesResults) GetSourceLanguageOk() (*LanguagesSourceLanguage, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *LanguagesResults) SetSourceLanguage(v LanguagesSourceLanguage)`

SetSourceLanguage sets SourceLanguage field to given value.

### HasSourceLanguage

`func (o *LanguagesResults) HasSourceLanguage() bool`

HasSourceLanguage returns a boolean if a field has been set.

### GetTargetLanguage

`func (o *LanguagesResults) GetTargetLanguage() LanguagesTargetLanguage`

GetTargetLanguage returns the TargetLanguage field if non-nil, zero value otherwise.

### GetTargetLanguageOk

`func (o *LanguagesResults) GetTargetLanguageOk() (*LanguagesTargetLanguage, bool)`

GetTargetLanguageOk returns a tuple with the TargetLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguage

`func (o *LanguagesResults) SetTargetLanguage(v LanguagesTargetLanguage)`

SetTargetLanguage sets TargetLanguage field to given value.

### HasTargetLanguage

`func (o *LanguagesResults) HasTargetLanguage() bool`

HasTargetLanguage returns a boolean if a field has been set.

### GetType

`func (o *LanguagesResults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LanguagesResults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LanguagesResults) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LanguagesResults) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


