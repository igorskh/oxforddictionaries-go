# ThesaurusEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrammaticalFeatures** | Pointer to **[]map[string]interface{}** | The different forms are correlated with meanings or functions which we text as &#39;features&#39; | [optional] 
**HomographNumber** | Pointer to **string** | Identifies the homograph grouping. The last two digits identify different entries of the same homograph. The first one/two digits identify the homograph number. | [optional] 
**Senses** | Pointer to [**[]ThesaurusSense**](ThesaurusSense.md) | Complete list of senses | [optional] 
**VariantForms** | Pointer to **[]map[string]interface{}** | Various words that are used interchangeably depending on the context, e.g &#39;aluminium&#39; and &#39;aluminum&#39; | [optional] 

## Methods

### NewThesaurusEntry

`func NewThesaurusEntry() *ThesaurusEntry`

NewThesaurusEntry instantiates a new ThesaurusEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThesaurusEntryWithDefaults

`func NewThesaurusEntryWithDefaults() *ThesaurusEntry`

NewThesaurusEntryWithDefaults instantiates a new ThesaurusEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrammaticalFeatures

`func (o *ThesaurusEntry) GetGrammaticalFeatures() []map[string]interface{}`

GetGrammaticalFeatures returns the GrammaticalFeatures field if non-nil, zero value otherwise.

### GetGrammaticalFeaturesOk

`func (o *ThesaurusEntry) GetGrammaticalFeaturesOk() (*[]map[string]interface{}, bool)`

GetGrammaticalFeaturesOk returns a tuple with the GrammaticalFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrammaticalFeatures

`func (o *ThesaurusEntry) SetGrammaticalFeatures(v []map[string]interface{})`

SetGrammaticalFeatures sets GrammaticalFeatures field to given value.

### HasGrammaticalFeatures

`func (o *ThesaurusEntry) HasGrammaticalFeatures() bool`

HasGrammaticalFeatures returns a boolean if a field has been set.

### GetHomographNumber

`func (o *ThesaurusEntry) GetHomographNumber() string`

GetHomographNumber returns the HomographNumber field if non-nil, zero value otherwise.

### GetHomographNumberOk

`func (o *ThesaurusEntry) GetHomographNumberOk() (*string, bool)`

GetHomographNumberOk returns a tuple with the HomographNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomographNumber

`func (o *ThesaurusEntry) SetHomographNumber(v string)`

SetHomographNumber sets HomographNumber field to given value.

### HasHomographNumber

`func (o *ThesaurusEntry) HasHomographNumber() bool`

HasHomographNumber returns a boolean if a field has been set.

### GetSenses

`func (o *ThesaurusEntry) GetSenses() []ThesaurusSense`

GetSenses returns the Senses field if non-nil, zero value otherwise.

### GetSensesOk

`func (o *ThesaurusEntry) GetSensesOk() (*[]ThesaurusSense, bool)`

GetSensesOk returns a tuple with the Senses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenses

`func (o *ThesaurusEntry) SetSenses(v []ThesaurusSense)`

SetSenses sets Senses field to given value.

### HasSenses

`func (o *ThesaurusEntry) HasSenses() bool`

HasSenses returns a boolean if a field has been set.

### GetVariantForms

`func (o *ThesaurusEntry) GetVariantForms() []map[string]interface{}`

GetVariantForms returns the VariantForms field if non-nil, zero value otherwise.

### GetVariantFormsOk

`func (o *ThesaurusEntry) GetVariantFormsOk() (*[]map[string]interface{}, bool)`

GetVariantFormsOk returns a tuple with the VariantForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantForms

`func (o *ThesaurusEntry) SetVariantForms(v []map[string]interface{})`

SetVariantForms sets VariantForms field to given value.

### HasVariantForms

`func (o *ThesaurusEntry) HasVariantForms() bool`

HasVariantForms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


