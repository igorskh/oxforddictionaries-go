# UtilityLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Additional Information provided by OUP | [optional] 
**Results** | Pointer to [**map[string]UtilityLabelsResults**](UtilityLabelsResults.md) | Mapping of labels available. | [optional] 

## Methods

### NewUtilityLabels

`func NewUtilityLabels() *UtilityLabels`

NewUtilityLabels instantiates a new UtilityLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilityLabelsWithDefaults

`func NewUtilityLabelsWithDefaults() *UtilityLabels`

NewUtilityLabelsWithDefaults instantiates a new UtilityLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *UtilityLabels) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UtilityLabels) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UtilityLabels) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UtilityLabels) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResults

`func (o *UtilityLabels) GetResults() map[string]UtilityLabelsResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UtilityLabels) GetResultsOk() (*map[string]UtilityLabelsResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UtilityLabels) SetResults(v map[string]UtilityLabelsResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *UtilityLabels) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


