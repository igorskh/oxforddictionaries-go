# Filters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Additional Information provided by OUP | [optional] 
**Results** | Pointer to [**FiltersResults**](FiltersResults.md) |  | [optional] 

## Methods

### NewFilters

`func NewFilters() *Filters`

NewFilters instantiates a new Filters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersWithDefaults

`func NewFiltersWithDefaults() *Filters`

NewFiltersWithDefaults instantiates a new Filters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Filters) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Filters) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Filters) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Filters) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResults

`func (o *Filters) GetResults() FiltersResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Filters) GetResultsOk() (*FiltersResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Filters) SetResults(v FiltersResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *Filters) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


