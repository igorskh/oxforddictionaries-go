# Wordlist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**WordlistMetadata**](WordlistMetadata.md) |  | [optional] 
**Results** | Pointer to [**[]WordlistResults**](WordlistResults.md) | A list of found words | [optional] 

## Methods

### NewWordlist

`func NewWordlist() *Wordlist`

NewWordlist instantiates a new Wordlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordlistWithDefaults

`func NewWordlistWithDefaults() *Wordlist`

NewWordlistWithDefaults instantiates a new Wordlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Wordlist) GetMetadata() WordlistMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Wordlist) GetMetadataOk() (*WordlistMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Wordlist) SetMetadata(v WordlistMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Wordlist) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResults

`func (o *Wordlist) GetResults() []WordlistResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Wordlist) GetResultsOk() (*[]WordlistResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Wordlist) SetResults(v []WordlistResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *Wordlist) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


