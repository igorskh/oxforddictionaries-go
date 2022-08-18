# Thesaurus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Additional Information provided by OUP | [optional] 
**Results** | Pointer to [**[]HeadwordThesaurus**](HeadwordThesaurus.md) | A list of found synonyms or antonyms | [optional] 

## Methods

### NewThesaurus

`func NewThesaurus() *Thesaurus`

NewThesaurus instantiates a new Thesaurus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThesaurusWithDefaults

`func NewThesaurusWithDefaults() *Thesaurus`

NewThesaurusWithDefaults instantiates a new Thesaurus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Thesaurus) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Thesaurus) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Thesaurus) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Thesaurus) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResults

`func (o *Thesaurus) GetResults() []HeadwordThesaurus`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Thesaurus) GetResultsOk() (*[]HeadwordThesaurus, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Thesaurus) SetResults(v []HeadwordThesaurus)`

SetResults sets Results field to given value.

### HasResults

`func (o *Thesaurus) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


