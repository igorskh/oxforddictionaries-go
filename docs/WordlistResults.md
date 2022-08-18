# WordlistResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of a word | 
**MatchString** | Pointer to **string** |  | [optional] 
**MatchType** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** | Name of region. | [optional] 
**Word** | **string** | (DEPRECATED) A given written or spoken realisation of an entry, lowercased. | 

## Methods

### NewWordlistResults

`func NewWordlistResults(id string, word string, ) *WordlistResults`

NewWordlistResults instantiates a new WordlistResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordlistResultsWithDefaults

`func NewWordlistResultsWithDefaults() *WordlistResults`

NewWordlistResultsWithDefaults instantiates a new WordlistResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WordlistResults) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WordlistResults) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WordlistResults) SetId(v string)`

SetId sets Id field to given value.


### GetMatchString

`func (o *WordlistResults) GetMatchString() string`

GetMatchString returns the MatchString field if non-nil, zero value otherwise.

### GetMatchStringOk

`func (o *WordlistResults) GetMatchStringOk() (*string, bool)`

GetMatchStringOk returns a tuple with the MatchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchString

`func (o *WordlistResults) SetMatchString(v string)`

SetMatchString sets MatchString field to given value.

### HasMatchString

`func (o *WordlistResults) HasMatchString() bool`

HasMatchString returns a boolean if a field has been set.

### GetMatchType

`func (o *WordlistResults) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *WordlistResults) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *WordlistResults) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *WordlistResults) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetRegion

`func (o *WordlistResults) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *WordlistResults) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *WordlistResults) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *WordlistResults) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetWord

`func (o *WordlistResults) GetWord() string`

GetWord returns the Word field if non-nil, zero value otherwise.

### GetWordOk

`func (o *WordlistResults) GetWordOk() (*string, bool)`

GetWordOk returns a tuple with the Word field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWord

`func (o *WordlistResults) SetWord(v string)`

SetWord sets Word field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


