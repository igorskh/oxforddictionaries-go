# SentencesResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Additional Information provided by OUP | [optional] 
**Results** | Pointer to [**[]SentencesEntry**](SentencesEntry.md) | A list of entries and all the data related to them | [optional] 

## Methods

### NewSentencesResults

`func NewSentencesResults() *SentencesResults`

NewSentencesResults instantiates a new SentencesResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSentencesResultsWithDefaults

`func NewSentencesResultsWithDefaults() *SentencesResults`

NewSentencesResultsWithDefaults instantiates a new SentencesResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SentencesResults) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SentencesResults) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SentencesResults) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SentencesResults) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResults

`func (o *SentencesResults) GetResults() []SentencesEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SentencesResults) GetResultsOk() (*[]SentencesEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SentencesResults) SetResults(v []SentencesEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *SentencesResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


