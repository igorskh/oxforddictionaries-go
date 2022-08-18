# BilingualSense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Antonyms** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Constructions** | Pointer to [**[]BilingualSenseConstructions**](BilingualSenseConstructions.md) | A construction provides information about typical syntax used of this sense. Each construction may optionally have one or more examples. Constructions may be translated if there is an equivalent in the target language. | [optional] 
**CrossReferenceMarkers** | Pointer to **[]string** |  | [optional] 
**CrossReferences** | Pointer to **[]map[string]interface{}** | A reference to another word that is closely related, might provide additional information about the subject, has a variant spelling or is an abbreviated form of it. | [optional] 
**DatasetCrossLinks** | Pointer to [**[]DatasetCrossLink**](DatasetCrossLink.md) | List of links to entries in other datasets | [optional] 
**Definitions** | Pointer to **[]string** |  | [optional] 
**DomainClasses** | Pointer to [**DomainClassesList**](DomainClassesList.md) |  | [optional] 
**Domains** | Pointer to [**DomainsList**](DomainsList.md) |  | [optional] 
**Etymologies** | Pointer to **[]string** |  | [optional] 
**Examples** | Pointer to **[]map[string]interface{}** | A list of written or spoken rendering of examples of use of a word or text | [optional] 
**Id** | Pointer to **string** | The id of the sense that is required for the delete procedure | [optional] 
**Inflections** | Pointer to [**[]InflectedForm**](InflectedForm.md) | A list of inflected forms for a sense. | [optional] 
**Notes** | Pointer to **[]map[string]interface{}** | various types of notes that appear | [optional] 
**Pronunciations** | Pointer to **[]map[string]interface{}** | A list of possible pronunciations of a word | [optional] 
**Regions** | Pointer to [**RegionsList**](RegionsList.md) |  | [optional] 
**Registers** | Pointer to [**RegistersList**](RegistersList.md) |  | [optional] 
**SemanticClasses** | Pointer to [**SemanticClassesList**](SemanticClassesList.md) |  | [optional] 
**ShortDefinitions** | Pointer to **[]string** |  | [optional] 
**Subsenses** | Pointer to [**[]BilingualSense**](BilingualSense.md) | Ordered list of subsenses of a sense | [optional] 
**Synonyms** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ThesaurusLinks** | Pointer to [**[]ThesaurusLink**](ThesaurusLink.md) | Ordered list of links to the Thesaurus Dictionary | [optional] 
**Translations** | Pointer to **[]map[string]interface{}** | A list of written or spoken rendering of the meaning of a word or text in another language(s) | [optional] 
**VariantForms** | Pointer to **[]map[string]interface{}** | Various words that are used interchangeably depending on the context, e.g &#39;aluminium&#39; and &#39;aluminum&#39; | [optional] 

## Methods

### NewBilingualSense

`func NewBilingualSense() *BilingualSense`

NewBilingualSense instantiates a new BilingualSense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBilingualSenseWithDefaults

`func NewBilingualSenseWithDefaults() *BilingualSense`

NewBilingualSenseWithDefaults instantiates a new BilingualSense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntonyms

`func (o *BilingualSense) GetAntonyms() []map[string]interface{}`

GetAntonyms returns the Antonyms field if non-nil, zero value otherwise.

### GetAntonymsOk

`func (o *BilingualSense) GetAntonymsOk() (*[]map[string]interface{}, bool)`

GetAntonymsOk returns a tuple with the Antonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntonyms

`func (o *BilingualSense) SetAntonyms(v []map[string]interface{})`

SetAntonyms sets Antonyms field to given value.

### HasAntonyms

`func (o *BilingualSense) HasAntonyms() bool`

HasAntonyms returns a boolean if a field has been set.

### GetConstructions

`func (o *BilingualSense) GetConstructions() []BilingualSenseConstructions`

GetConstructions returns the Constructions field if non-nil, zero value otherwise.

### GetConstructionsOk

`func (o *BilingualSense) GetConstructionsOk() (*[]BilingualSenseConstructions, bool)`

GetConstructionsOk returns a tuple with the Constructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstructions

`func (o *BilingualSense) SetConstructions(v []BilingualSenseConstructions)`

SetConstructions sets Constructions field to given value.

### HasConstructions

`func (o *BilingualSense) HasConstructions() bool`

HasConstructions returns a boolean if a field has been set.

### GetCrossReferenceMarkers

`func (o *BilingualSense) GetCrossReferenceMarkers() []string`

GetCrossReferenceMarkers returns the CrossReferenceMarkers field if non-nil, zero value otherwise.

### GetCrossReferenceMarkersOk

`func (o *BilingualSense) GetCrossReferenceMarkersOk() (*[]string, bool)`

GetCrossReferenceMarkersOk returns a tuple with the CrossReferenceMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossReferenceMarkers

`func (o *BilingualSense) SetCrossReferenceMarkers(v []string)`

SetCrossReferenceMarkers sets CrossReferenceMarkers field to given value.

### HasCrossReferenceMarkers

`func (o *BilingualSense) HasCrossReferenceMarkers() bool`

HasCrossReferenceMarkers returns a boolean if a field has been set.

### GetCrossReferences

`func (o *BilingualSense) GetCrossReferences() []map[string]interface{}`

GetCrossReferences returns the CrossReferences field if non-nil, zero value otherwise.

### GetCrossReferencesOk

`func (o *BilingualSense) GetCrossReferencesOk() (*[]map[string]interface{}, bool)`

GetCrossReferencesOk returns a tuple with the CrossReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossReferences

`func (o *BilingualSense) SetCrossReferences(v []map[string]interface{})`

SetCrossReferences sets CrossReferences field to given value.

### HasCrossReferences

`func (o *BilingualSense) HasCrossReferences() bool`

HasCrossReferences returns a boolean if a field has been set.

### GetDatasetCrossLinks

`func (o *BilingualSense) GetDatasetCrossLinks() []DatasetCrossLink`

GetDatasetCrossLinks returns the DatasetCrossLinks field if non-nil, zero value otherwise.

### GetDatasetCrossLinksOk

`func (o *BilingualSense) GetDatasetCrossLinksOk() (*[]DatasetCrossLink, bool)`

GetDatasetCrossLinksOk returns a tuple with the DatasetCrossLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetCrossLinks

`func (o *BilingualSense) SetDatasetCrossLinks(v []DatasetCrossLink)`

SetDatasetCrossLinks sets DatasetCrossLinks field to given value.

### HasDatasetCrossLinks

`func (o *BilingualSense) HasDatasetCrossLinks() bool`

HasDatasetCrossLinks returns a boolean if a field has been set.

### GetDefinitions

`func (o *BilingualSense) GetDefinitions() []string`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *BilingualSense) GetDefinitionsOk() (*[]string, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *BilingualSense) SetDefinitions(v []string)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *BilingualSense) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetDomainClasses

`func (o *BilingualSense) GetDomainClasses() DomainClassesList`

GetDomainClasses returns the DomainClasses field if non-nil, zero value otherwise.

### GetDomainClassesOk

`func (o *BilingualSense) GetDomainClassesOk() (*DomainClassesList, bool)`

GetDomainClassesOk returns a tuple with the DomainClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClasses

`func (o *BilingualSense) SetDomainClasses(v DomainClassesList)`

SetDomainClasses sets DomainClasses field to given value.

### HasDomainClasses

`func (o *BilingualSense) HasDomainClasses() bool`

HasDomainClasses returns a boolean if a field has been set.

### GetDomains

`func (o *BilingualSense) GetDomains() DomainsList`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *BilingualSense) GetDomainsOk() (*DomainsList, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *BilingualSense) SetDomains(v DomainsList)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *BilingualSense) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEtymologies

`func (o *BilingualSense) GetEtymologies() []string`

GetEtymologies returns the Etymologies field if non-nil, zero value otherwise.

### GetEtymologiesOk

`func (o *BilingualSense) GetEtymologiesOk() (*[]string, bool)`

GetEtymologiesOk returns a tuple with the Etymologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtymologies

`func (o *BilingualSense) SetEtymologies(v []string)`

SetEtymologies sets Etymologies field to given value.

### HasEtymologies

`func (o *BilingualSense) HasEtymologies() bool`

HasEtymologies returns a boolean if a field has been set.

### GetExamples

`func (o *BilingualSense) GetExamples() []map[string]interface{}`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *BilingualSense) GetExamplesOk() (*[]map[string]interface{}, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *BilingualSense) SetExamples(v []map[string]interface{})`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *BilingualSense) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetId

`func (o *BilingualSense) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BilingualSense) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BilingualSense) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BilingualSense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInflections

`func (o *BilingualSense) GetInflections() []InflectedForm`

GetInflections returns the Inflections field if non-nil, zero value otherwise.

### GetInflectionsOk

`func (o *BilingualSense) GetInflectionsOk() (*[]InflectedForm, bool)`

GetInflectionsOk returns a tuple with the Inflections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflections

`func (o *BilingualSense) SetInflections(v []InflectedForm)`

SetInflections sets Inflections field to given value.

### HasInflections

`func (o *BilingualSense) HasInflections() bool`

HasInflections returns a boolean if a field has been set.

### GetNotes

`func (o *BilingualSense) GetNotes() []map[string]interface{}`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BilingualSense) GetNotesOk() (*[]map[string]interface{}, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BilingualSense) SetNotes(v []map[string]interface{})`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BilingualSense) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPronunciations

`func (o *BilingualSense) GetPronunciations() []map[string]interface{}`

GetPronunciations returns the Pronunciations field if non-nil, zero value otherwise.

### GetPronunciationsOk

`func (o *BilingualSense) GetPronunciationsOk() (*[]map[string]interface{}, bool)`

GetPronunciationsOk returns a tuple with the Pronunciations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronunciations

`func (o *BilingualSense) SetPronunciations(v []map[string]interface{})`

SetPronunciations sets Pronunciations field to given value.

### HasPronunciations

`func (o *BilingualSense) HasPronunciations() bool`

HasPronunciations returns a boolean if a field has been set.

### GetRegions

`func (o *BilingualSense) GetRegions() RegionsList`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *BilingualSense) GetRegionsOk() (*RegionsList, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *BilingualSense) SetRegions(v RegionsList)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *BilingualSense) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRegisters

`func (o *BilingualSense) GetRegisters() RegistersList`

GetRegisters returns the Registers field if non-nil, zero value otherwise.

### GetRegistersOk

`func (o *BilingualSense) GetRegistersOk() (*RegistersList, bool)`

GetRegistersOk returns a tuple with the Registers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisters

`func (o *BilingualSense) SetRegisters(v RegistersList)`

SetRegisters sets Registers field to given value.

### HasRegisters

`func (o *BilingualSense) HasRegisters() bool`

HasRegisters returns a boolean if a field has been set.

### GetSemanticClasses

`func (o *BilingualSense) GetSemanticClasses() SemanticClassesList`

GetSemanticClasses returns the SemanticClasses field if non-nil, zero value otherwise.

### GetSemanticClassesOk

`func (o *BilingualSense) GetSemanticClassesOk() (*SemanticClassesList, bool)`

GetSemanticClassesOk returns a tuple with the SemanticClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticClasses

`func (o *BilingualSense) SetSemanticClasses(v SemanticClassesList)`

SetSemanticClasses sets SemanticClasses field to given value.

### HasSemanticClasses

`func (o *BilingualSense) HasSemanticClasses() bool`

HasSemanticClasses returns a boolean if a field has been set.

### GetShortDefinitions

`func (o *BilingualSense) GetShortDefinitions() []string`

GetShortDefinitions returns the ShortDefinitions field if non-nil, zero value otherwise.

### GetShortDefinitionsOk

`func (o *BilingualSense) GetShortDefinitionsOk() (*[]string, bool)`

GetShortDefinitionsOk returns a tuple with the ShortDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDefinitions

`func (o *BilingualSense) SetShortDefinitions(v []string)`

SetShortDefinitions sets ShortDefinitions field to given value.

### HasShortDefinitions

`func (o *BilingualSense) HasShortDefinitions() bool`

HasShortDefinitions returns a boolean if a field has been set.

### GetSubsenses

`func (o *BilingualSense) GetSubsenses() []BilingualSense`

GetSubsenses returns the Subsenses field if non-nil, zero value otherwise.

### GetSubsensesOk

`func (o *BilingualSense) GetSubsensesOk() (*[]BilingualSense, bool)`

GetSubsensesOk returns a tuple with the Subsenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsenses

`func (o *BilingualSense) SetSubsenses(v []BilingualSense)`

SetSubsenses sets Subsenses field to given value.

### HasSubsenses

`func (o *BilingualSense) HasSubsenses() bool`

HasSubsenses returns a boolean if a field has been set.

### GetSynonyms

`func (o *BilingualSense) GetSynonyms() []map[string]interface{}`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *BilingualSense) GetSynonymsOk() (*[]map[string]interface{}, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *BilingualSense) SetSynonyms(v []map[string]interface{})`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *BilingualSense) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetThesaurusLinks

`func (o *BilingualSense) GetThesaurusLinks() []ThesaurusLink`

GetThesaurusLinks returns the ThesaurusLinks field if non-nil, zero value otherwise.

### GetThesaurusLinksOk

`func (o *BilingualSense) GetThesaurusLinksOk() (*[]ThesaurusLink, bool)`

GetThesaurusLinksOk returns a tuple with the ThesaurusLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThesaurusLinks

`func (o *BilingualSense) SetThesaurusLinks(v []ThesaurusLink)`

SetThesaurusLinks sets ThesaurusLinks field to given value.

### HasThesaurusLinks

`func (o *BilingualSense) HasThesaurusLinks() bool`

HasThesaurusLinks returns a boolean if a field has been set.

### GetTranslations

`func (o *BilingualSense) GetTranslations() []map[string]interface{}`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *BilingualSense) GetTranslationsOk() (*[]map[string]interface{}, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *BilingualSense) SetTranslations(v []map[string]interface{})`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *BilingualSense) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### GetVariantForms

`func (o *BilingualSense) GetVariantForms() []map[string]interface{}`

GetVariantForms returns the VariantForms field if non-nil, zero value otherwise.

### GetVariantFormsOk

`func (o *BilingualSense) GetVariantFormsOk() (*[]map[string]interface{}, bool)`

GetVariantFormsOk returns a tuple with the VariantForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantForms

`func (o *BilingualSense) SetVariantForms(v []map[string]interface{})`

SetVariantForms sets VariantForms field to given value.

### HasVariantForms

`func (o *BilingualSense) HasVariantForms() bool`

HasVariantForms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


