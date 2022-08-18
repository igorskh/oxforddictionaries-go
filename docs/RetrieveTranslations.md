# RetrieveTranslations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Additional Information provided by OUP | [optional] 
**Results** | Pointer to [**[]BilingualHeadwordEntry**](BilingualHeadwordEntry.md) | A list of bilingual entries and all the data related to them | [optional] 

## Methods

### NewRetrieveTranslations

`func NewRetrieveTranslations() *RetrieveTranslations`

NewRetrieveTranslations instantiates a new RetrieveTranslations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveTranslationsWithDefaults

`func NewRetrieveTranslationsWithDefaults() *RetrieveTranslations`

NewRetrieveTranslationsWithDefaults instantiates a new RetrieveTranslations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RetrieveTranslations) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RetrieveTranslations) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RetrieveTranslations) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RetrieveTranslations) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResults

`func (o *RetrieveTranslations) GetResults() []BilingualHeadwordEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RetrieveTranslations) GetResultsOk() (*[]BilingualHeadwordEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RetrieveTranslations) SetResults(v []BilingualHeadwordEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *RetrieveTranslations) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


