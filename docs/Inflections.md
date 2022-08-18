# Inflections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Additional Information provided by OUP | [optional] 
**Results** | Pointer to [**[]HeadwordInflection**](HeadwordInflection.md) | A list of headwords matching a given word | [optional] 

## Methods

### NewInflections

`func NewInflections() *Inflections`

NewInflections instantiates a new Inflections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInflectionsWithDefaults

`func NewInflectionsWithDefaults() *Inflections`

NewInflectionsWithDefaults instantiates a new Inflections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Inflections) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Inflections) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Inflections) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Inflections) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResults

`func (o *Inflections) GetResults() []HeadwordInflection`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Inflections) GetResultsOk() (*[]HeadwordInflection, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Inflections) SetResults(v []HeadwordInflection)`

SetResults sets Results field to given value.

### HasResults

`func (o *Inflections) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


