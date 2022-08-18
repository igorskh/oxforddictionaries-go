# BilingualEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrossReferenceMarkers** | Pointer to **[]string** |  | [optional] 
**CrossReferences** | Pointer to **[]map[string]interface{}** | A reference to another word that is closely related, might provide additional information about the subject, has a variant spelling or is an abbreviated form of it. | [optional] 
**Etymologies** | Pointer to **[]string** |  | [optional] 
**GrammaticalFeatures** | Pointer to **[]map[string]interface{}** | The different forms are correlated with meanings or functions which we text as &#39;features&#39; | [optional] 
**HomographNumber** | Pointer to **string** | Identifies the homograph grouping. The last two digits identify different entries of the same homograph. The first one/two digits identify the homograph number. | [optional] 
**Inflections** | Pointer to [**[]InflectedForm**](InflectedForm.md) | A list of inflected forms for an Entry. | [optional] 
**Notes** | Pointer to **[]map[string]interface{}** | various types of notes that appear | [optional] 
**Pronunciations** | Pointer to **[]map[string]interface{}** | A list of possible pronunciations of a word | [optional] 
**Senses** | Pointer to [**[]BilingualSense**](BilingualSense.md) | Complete list of senses for bilingual entries | [optional] 
**VariantForms** | Pointer to **[]map[string]interface{}** | Various words that are used interchangeably depending on the context, e.g &#39;aluminium&#39; and &#39;aluminum&#39; | [optional] 

## Methods

### NewBilingualEntry

`func NewBilingualEntry() *BilingualEntry`

NewBilingualEntry instantiates a new BilingualEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBilingualEntryWithDefaults

`func NewBilingualEntryWithDefaults() *BilingualEntry`

NewBilingualEntryWithDefaults instantiates a new BilingualEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrossReferenceMarkers

`func (o *BilingualEntry) GetCrossReferenceMarkers() []string`

GetCrossReferenceMarkers returns the CrossReferenceMarkers field if non-nil, zero value otherwise.

### GetCrossReferenceMarkersOk

`func (o *BilingualEntry) GetCrossReferenceMarkersOk() (*[]string, bool)`

GetCrossReferenceMarkersOk returns a tuple with the CrossReferenceMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossReferenceMarkers

`func (o *BilingualEntry) SetCrossReferenceMarkers(v []string)`

SetCrossReferenceMarkers sets CrossReferenceMarkers field to given value.

### HasCrossReferenceMarkers

`func (o *BilingualEntry) HasCrossReferenceMarkers() bool`

HasCrossReferenceMarkers returns a boolean if a field has been set.

### GetCrossReferences

`func (o *BilingualEntry) GetCrossReferences() []map[string]interface{}`

GetCrossReferences returns the CrossReferences field if non-nil, zero value otherwise.

### GetCrossReferencesOk

`func (o *BilingualEntry) GetCrossReferencesOk() (*[]map[string]interface{}, bool)`

GetCrossReferencesOk returns a tuple with the CrossReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossReferences

`func (o *BilingualEntry) SetCrossReferences(v []map[string]interface{})`

SetCrossReferences sets CrossReferences field to given value.

### HasCrossReferences

`func (o *BilingualEntry) HasCrossReferences() bool`

HasCrossReferences returns a boolean if a field has been set.

### GetEtymologies

`func (o *BilingualEntry) GetEtymologies() []string`

GetEtymologies returns the Etymologies field if non-nil, zero value otherwise.

### GetEtymologiesOk

`func (o *BilingualEntry) GetEtymologiesOk() (*[]string, bool)`

GetEtymologiesOk returns a tuple with the Etymologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtymologies

`func (o *BilingualEntry) SetEtymologies(v []string)`

SetEtymologies sets Etymologies field to given value.

### HasEtymologies

`func (o *BilingualEntry) HasEtymologies() bool`

HasEtymologies returns a boolean if a field has been set.

### GetGrammaticalFeatures

`func (o *BilingualEntry) GetGrammaticalFeatures() []map[string]interface{}`

GetGrammaticalFeatures returns the GrammaticalFeatures field if non-nil, zero value otherwise.

### GetGrammaticalFeaturesOk

`func (o *BilingualEntry) GetGrammaticalFeaturesOk() (*[]map[string]interface{}, bool)`

GetGrammaticalFeaturesOk returns a tuple with the GrammaticalFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrammaticalFeatures

`func (o *BilingualEntry) SetGrammaticalFeatures(v []map[string]interface{})`

SetGrammaticalFeatures sets GrammaticalFeatures field to given value.

### HasGrammaticalFeatures

`func (o *BilingualEntry) HasGrammaticalFeatures() bool`

HasGrammaticalFeatures returns a boolean if a field has been set.

### GetHomographNumber

`func (o *BilingualEntry) GetHomographNumber() string`

GetHomographNumber returns the HomographNumber field if non-nil, zero value otherwise.

### GetHomographNumberOk

`func (o *BilingualEntry) GetHomographNumberOk() (*string, bool)`

GetHomographNumberOk returns a tuple with the HomographNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomographNumber

`func (o *BilingualEntry) SetHomographNumber(v string)`

SetHomographNumber sets HomographNumber field to given value.

### HasHomographNumber

`func (o *BilingualEntry) HasHomographNumber() bool`

HasHomographNumber returns a boolean if a field has been set.

### GetInflections

`func (o *BilingualEntry) GetInflections() []InflectedForm`

GetInflections returns the Inflections field if non-nil, zero value otherwise.

### GetInflectionsOk

`func (o *BilingualEntry) GetInflectionsOk() (*[]InflectedForm, bool)`

GetInflectionsOk returns a tuple with the Inflections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflections

`func (o *BilingualEntry) SetInflections(v []InflectedForm)`

SetInflections sets Inflections field to given value.

### HasInflections

`func (o *BilingualEntry) HasInflections() bool`

HasInflections returns a boolean if a field has been set.

### GetNotes

`func (o *BilingualEntry) GetNotes() []map[string]interface{}`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BilingualEntry) GetNotesOk() (*[]map[string]interface{}, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BilingualEntry) SetNotes(v []map[string]interface{})`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BilingualEntry) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPronunciations

`func (o *BilingualEntry) GetPronunciations() []map[string]interface{}`

GetPronunciations returns the Pronunciations field if non-nil, zero value otherwise.

### GetPronunciationsOk

`func (o *BilingualEntry) GetPronunciationsOk() (*[]map[string]interface{}, bool)`

GetPronunciationsOk returns a tuple with the Pronunciations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronunciations

`func (o *BilingualEntry) SetPronunciations(v []map[string]interface{})`

SetPronunciations sets Pronunciations field to given value.

### HasPronunciations

`func (o *BilingualEntry) HasPronunciations() bool`

HasPronunciations returns a boolean if a field has been set.

### GetSenses

`func (o *BilingualEntry) GetSenses() []BilingualSense`

GetSenses returns the Senses field if non-nil, zero value otherwise.

### GetSensesOk

`func (o *BilingualEntry) GetSensesOk() (*[]BilingualSense, bool)`

GetSensesOk returns a tuple with the Senses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenses

`func (o *BilingualEntry) SetSenses(v []BilingualSense)`

SetSenses sets Senses field to given value.

### HasSenses

`func (o *BilingualEntry) HasSenses() bool`

HasSenses returns a boolean if a field has been set.

### GetVariantForms

`func (o *BilingualEntry) GetVariantForms() []map[string]interface{}`

GetVariantForms returns the VariantForms field if non-nil, zero value otherwise.

### GetVariantFormsOk

`func (o *BilingualEntry) GetVariantFormsOk() (*[]map[string]interface{}, bool)`

GetVariantFormsOk returns a tuple with the VariantForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantForms

`func (o *BilingualEntry) SetVariantForms(v []map[string]interface{})`

SetVariantForms sets VariantForms field to given value.

### HasVariantForms

`func (o *BilingualEntry) HasVariantForms() bool`

HasVariantForms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


