# Sense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Antonyms** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Constructions** | Pointer to [**[]SenseConstructions**](SenseConstructions.md) | A construction provides information about typical syntax used of this sense. Each construction may optionally have one or more examples. | [optional] 
**CrossReferenceMarkers** | Pointer to **[]string** |  | [optional] 
**CrossReferences** | Pointer to **[]map[string]interface{}** | A reference to another word that is closely related, might provide additional information about the subject, has a variant spelling or is an abbreviated form of it. | [optional] 
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
**Subsenses** | Pointer to [**[]Sense**](Sense.md) | Ordered list of subsenses of a sense | [optional] 
**Synonyms** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ThesaurusLinks** | Pointer to [**[]ThesaurusLink**](ThesaurusLink.md) | Ordered list of links to the Thesaurus Dictionary | [optional] 
**VariantForms** | Pointer to **[]map[string]interface{}** | Various words that are used interchangeably depending on the context, e.g &#39;aluminium&#39; and &#39;aluminum&#39; | [optional] 

## Methods

### NewSense

`func NewSense() *Sense`

NewSense instantiates a new Sense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSenseWithDefaults

`func NewSenseWithDefaults() *Sense`

NewSenseWithDefaults instantiates a new Sense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntonyms

`func (o *Sense) GetAntonyms() []map[string]interface{}`

GetAntonyms returns the Antonyms field if non-nil, zero value otherwise.

### GetAntonymsOk

`func (o *Sense) GetAntonymsOk() (*[]map[string]interface{}, bool)`

GetAntonymsOk returns a tuple with the Antonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntonyms

`func (o *Sense) SetAntonyms(v []map[string]interface{})`

SetAntonyms sets Antonyms field to given value.

### HasAntonyms

`func (o *Sense) HasAntonyms() bool`

HasAntonyms returns a boolean if a field has been set.

### GetConstructions

`func (o *Sense) GetConstructions() []SenseConstructions`

GetConstructions returns the Constructions field if non-nil, zero value otherwise.

### GetConstructionsOk

`func (o *Sense) GetConstructionsOk() (*[]SenseConstructions, bool)`

GetConstructionsOk returns a tuple with the Constructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstructions

`func (o *Sense) SetConstructions(v []SenseConstructions)`

SetConstructions sets Constructions field to given value.

### HasConstructions

`func (o *Sense) HasConstructions() bool`

HasConstructions returns a boolean if a field has been set.

### GetCrossReferenceMarkers

`func (o *Sense) GetCrossReferenceMarkers() []string`

GetCrossReferenceMarkers returns the CrossReferenceMarkers field if non-nil, zero value otherwise.

### GetCrossReferenceMarkersOk

`func (o *Sense) GetCrossReferenceMarkersOk() (*[]string, bool)`

GetCrossReferenceMarkersOk returns a tuple with the CrossReferenceMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossReferenceMarkers

`func (o *Sense) SetCrossReferenceMarkers(v []string)`

SetCrossReferenceMarkers sets CrossReferenceMarkers field to given value.

### HasCrossReferenceMarkers

`func (o *Sense) HasCrossReferenceMarkers() bool`

HasCrossReferenceMarkers returns a boolean if a field has been set.

### GetCrossReferences

`func (o *Sense) GetCrossReferences() []map[string]interface{}`

GetCrossReferences returns the CrossReferences field if non-nil, zero value otherwise.

### GetCrossReferencesOk

`func (o *Sense) GetCrossReferencesOk() (*[]map[string]interface{}, bool)`

GetCrossReferencesOk returns a tuple with the CrossReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossReferences

`func (o *Sense) SetCrossReferences(v []map[string]interface{})`

SetCrossReferences sets CrossReferences field to given value.

### HasCrossReferences

`func (o *Sense) HasCrossReferences() bool`

HasCrossReferences returns a boolean if a field has been set.

### GetDefinitions

`func (o *Sense) GetDefinitions() []string`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *Sense) GetDefinitionsOk() (*[]string, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *Sense) SetDefinitions(v []string)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *Sense) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetDomainClasses

`func (o *Sense) GetDomainClasses() DomainClassesList`

GetDomainClasses returns the DomainClasses field if non-nil, zero value otherwise.

### GetDomainClassesOk

`func (o *Sense) GetDomainClassesOk() (*DomainClassesList, bool)`

GetDomainClassesOk returns a tuple with the DomainClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClasses

`func (o *Sense) SetDomainClasses(v DomainClassesList)`

SetDomainClasses sets DomainClasses field to given value.

### HasDomainClasses

`func (o *Sense) HasDomainClasses() bool`

HasDomainClasses returns a boolean if a field has been set.

### GetDomains

`func (o *Sense) GetDomains() DomainsList`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Sense) GetDomainsOk() (*DomainsList, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *Sense) SetDomains(v DomainsList)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *Sense) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEtymologies

`func (o *Sense) GetEtymologies() []string`

GetEtymologies returns the Etymologies field if non-nil, zero value otherwise.

### GetEtymologiesOk

`func (o *Sense) GetEtymologiesOk() (*[]string, bool)`

GetEtymologiesOk returns a tuple with the Etymologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtymologies

`func (o *Sense) SetEtymologies(v []string)`

SetEtymologies sets Etymologies field to given value.

### HasEtymologies

`func (o *Sense) HasEtymologies() bool`

HasEtymologies returns a boolean if a field has been set.

### GetExamples

`func (o *Sense) GetExamples() []map[string]interface{}`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *Sense) GetExamplesOk() (*[]map[string]interface{}, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *Sense) SetExamples(v []map[string]interface{})`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *Sense) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetId

`func (o *Sense) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sense) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sense) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Sense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInflections

`func (o *Sense) GetInflections() []InflectedForm`

GetInflections returns the Inflections field if non-nil, zero value otherwise.

### GetInflectionsOk

`func (o *Sense) GetInflectionsOk() (*[]InflectedForm, bool)`

GetInflectionsOk returns a tuple with the Inflections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflections

`func (o *Sense) SetInflections(v []InflectedForm)`

SetInflections sets Inflections field to given value.

### HasInflections

`func (o *Sense) HasInflections() bool`

HasInflections returns a boolean if a field has been set.

### GetNotes

`func (o *Sense) GetNotes() []map[string]interface{}`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Sense) GetNotesOk() (*[]map[string]interface{}, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Sense) SetNotes(v []map[string]interface{})`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Sense) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPronunciations

`func (o *Sense) GetPronunciations() []map[string]interface{}`

GetPronunciations returns the Pronunciations field if non-nil, zero value otherwise.

### GetPronunciationsOk

`func (o *Sense) GetPronunciationsOk() (*[]map[string]interface{}, bool)`

GetPronunciationsOk returns a tuple with the Pronunciations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronunciations

`func (o *Sense) SetPronunciations(v []map[string]interface{})`

SetPronunciations sets Pronunciations field to given value.

### HasPronunciations

`func (o *Sense) HasPronunciations() bool`

HasPronunciations returns a boolean if a field has been set.

### GetRegions

`func (o *Sense) GetRegions() RegionsList`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Sense) GetRegionsOk() (*RegionsList, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Sense) SetRegions(v RegionsList)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Sense) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRegisters

`func (o *Sense) GetRegisters() RegistersList`

GetRegisters returns the Registers field if non-nil, zero value otherwise.

### GetRegistersOk

`func (o *Sense) GetRegistersOk() (*RegistersList, bool)`

GetRegistersOk returns a tuple with the Registers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisters

`func (o *Sense) SetRegisters(v RegistersList)`

SetRegisters sets Registers field to given value.

### HasRegisters

`func (o *Sense) HasRegisters() bool`

HasRegisters returns a boolean if a field has been set.

### GetSemanticClasses

`func (o *Sense) GetSemanticClasses() SemanticClassesList`

GetSemanticClasses returns the SemanticClasses field if non-nil, zero value otherwise.

### GetSemanticClassesOk

`func (o *Sense) GetSemanticClassesOk() (*SemanticClassesList, bool)`

GetSemanticClassesOk returns a tuple with the SemanticClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticClasses

`func (o *Sense) SetSemanticClasses(v SemanticClassesList)`

SetSemanticClasses sets SemanticClasses field to given value.

### HasSemanticClasses

`func (o *Sense) HasSemanticClasses() bool`

HasSemanticClasses returns a boolean if a field has been set.

### GetShortDefinitions

`func (o *Sense) GetShortDefinitions() []string`

GetShortDefinitions returns the ShortDefinitions field if non-nil, zero value otherwise.

### GetShortDefinitionsOk

`func (o *Sense) GetShortDefinitionsOk() (*[]string, bool)`

GetShortDefinitionsOk returns a tuple with the ShortDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDefinitions

`func (o *Sense) SetShortDefinitions(v []string)`

SetShortDefinitions sets ShortDefinitions field to given value.

### HasShortDefinitions

`func (o *Sense) HasShortDefinitions() bool`

HasShortDefinitions returns a boolean if a field has been set.

### GetSubsenses

`func (o *Sense) GetSubsenses() []Sense`

GetSubsenses returns the Subsenses field if non-nil, zero value otherwise.

### GetSubsensesOk

`func (o *Sense) GetSubsensesOk() (*[]Sense, bool)`

GetSubsensesOk returns a tuple with the Subsenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsenses

`func (o *Sense) SetSubsenses(v []Sense)`

SetSubsenses sets Subsenses field to given value.

### HasSubsenses

`func (o *Sense) HasSubsenses() bool`

HasSubsenses returns a boolean if a field has been set.

### GetSynonyms

`func (o *Sense) GetSynonyms() []map[string]interface{}`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *Sense) GetSynonymsOk() (*[]map[string]interface{}, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *Sense) SetSynonyms(v []map[string]interface{})`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *Sense) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetThesaurusLinks

`func (o *Sense) GetThesaurusLinks() []ThesaurusLink`

GetThesaurusLinks returns the ThesaurusLinks field if non-nil, zero value otherwise.

### GetThesaurusLinksOk

`func (o *Sense) GetThesaurusLinksOk() (*[]ThesaurusLink, bool)`

GetThesaurusLinksOk returns a tuple with the ThesaurusLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThesaurusLinks

`func (o *Sense) SetThesaurusLinks(v []ThesaurusLink)`

SetThesaurusLinks sets ThesaurusLinks field to given value.

### HasThesaurusLinks

`func (o *Sense) HasThesaurusLinks() bool`

HasThesaurusLinks returns a boolean if a field has been set.

### GetVariantForms

`func (o *Sense) GetVariantForms() []map[string]interface{}`

GetVariantForms returns the VariantForms field if non-nil, zero value otherwise.

### GetVariantFormsOk

`func (o *Sense) GetVariantFormsOk() (*[]map[string]interface{}, bool)`

GetVariantFormsOk returns a tuple with the VariantForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantForms

`func (o *Sense) SetVariantForms(v []map[string]interface{})`

SetVariantForms sets VariantForms field to given value.

### HasVariantForms

`func (o *Sense) HasVariantForms() bool`

HasVariantForms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


