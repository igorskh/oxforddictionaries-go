# RetrieveEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Additional Information provided by OUP | [optional] 
**Results** | Pointer to [**[]HeadwordEntry**](HeadwordEntry.md) | A list of entries and all the data related to them | [optional] 

## Methods

### NewRetrieveEntry

`func NewRetrieveEntry() *RetrieveEntry`

NewRetrieveEntry instantiates a new RetrieveEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveEntryWithDefaults

`func NewRetrieveEntryWithDefaults() *RetrieveEntry`

NewRetrieveEntryWithDefaults instantiates a new RetrieveEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RetrieveEntry) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RetrieveEntry) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RetrieveEntry) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RetrieveEntry) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResults

`func (o *RetrieveEntry) GetResults() []HeadwordEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RetrieveEntry) GetResultsOk() (*[]HeadwordEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RetrieveEntry) SetResults(v []HeadwordEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *RetrieveEntry) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


