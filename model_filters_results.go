/*
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.5.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxforddictionaries

import (
	"encoding/json"
)

// FiltersResults A mapping of filters available per endpoints.
type FiltersResults struct {
	Entries *[]string `json:"entries,omitempty"`
	Lemmas *[]string `json:"lemmas,omitempty"`
	Sentences *[]string `json:"sentences,omitempty"`
	Thesaurus *[]string `json:"thesaurus,omitempty"`
	Translations *[]string `json:"translations,omitempty"`
}

// NewFiltersResults instantiates a new FiltersResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersResults() *FiltersResults {
	this := FiltersResults{}
	return &this
}

// NewFiltersResultsWithDefaults instantiates a new FiltersResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersResultsWithDefaults() *FiltersResults {
	this := FiltersResults{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *FiltersResults) GetEntries() []string {
	if o == nil || o.Entries == nil {
		var ret []string
		return ret
	}
	return *o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersResults) GetEntriesOk() (*[]string, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *FiltersResults) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []string and assigns it to the Entries field.
func (o *FiltersResults) SetEntries(v []string) {
	o.Entries = &v
}

// GetLemmas returns the Lemmas field value if set, zero value otherwise.
func (o *FiltersResults) GetLemmas() []string {
	if o == nil || o.Lemmas == nil {
		var ret []string
		return ret
	}
	return *o.Lemmas
}

// GetLemmasOk returns a tuple with the Lemmas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersResults) GetLemmasOk() (*[]string, bool) {
	if o == nil || o.Lemmas == nil {
		return nil, false
	}
	return o.Lemmas, true
}

// HasLemmas returns a boolean if a field has been set.
func (o *FiltersResults) HasLemmas() bool {
	if o != nil && o.Lemmas != nil {
		return true
	}

	return false
}

// SetLemmas gets a reference to the given []string and assigns it to the Lemmas field.
func (o *FiltersResults) SetLemmas(v []string) {
	o.Lemmas = &v
}

// GetSentences returns the Sentences field value if set, zero value otherwise.
func (o *FiltersResults) GetSentences() []string {
	if o == nil || o.Sentences == nil {
		var ret []string
		return ret
	}
	return *o.Sentences
}

// GetSentencesOk returns a tuple with the Sentences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersResults) GetSentencesOk() (*[]string, bool) {
	if o == nil || o.Sentences == nil {
		return nil, false
	}
	return o.Sentences, true
}

// HasSentences returns a boolean if a field has been set.
func (o *FiltersResults) HasSentences() bool {
	if o != nil && o.Sentences != nil {
		return true
	}

	return false
}

// SetSentences gets a reference to the given []string and assigns it to the Sentences field.
func (o *FiltersResults) SetSentences(v []string) {
	o.Sentences = &v
}

// GetThesaurus returns the Thesaurus field value if set, zero value otherwise.
func (o *FiltersResults) GetThesaurus() []string {
	if o == nil || o.Thesaurus == nil {
		var ret []string
		return ret
	}
	return *o.Thesaurus
}

// GetThesaurusOk returns a tuple with the Thesaurus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersResults) GetThesaurusOk() (*[]string, bool) {
	if o == nil || o.Thesaurus == nil {
		return nil, false
	}
	return o.Thesaurus, true
}

// HasThesaurus returns a boolean if a field has been set.
func (o *FiltersResults) HasThesaurus() bool {
	if o != nil && o.Thesaurus != nil {
		return true
	}

	return false
}

// SetThesaurus gets a reference to the given []string and assigns it to the Thesaurus field.
func (o *FiltersResults) SetThesaurus(v []string) {
	o.Thesaurus = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *FiltersResults) GetTranslations() []string {
	if o == nil || o.Translations == nil {
		var ret []string
		return ret
	}
	return *o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersResults) GetTranslationsOk() (*[]string, bool) {
	if o == nil || o.Translations == nil {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *FiltersResults) HasTranslations() bool {
	if o != nil && o.Translations != nil {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []string and assigns it to the Translations field.
func (o *FiltersResults) SetTranslations(v []string) {
	o.Translations = &v
}

func (o FiltersResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.Lemmas != nil {
		toSerialize["lemmas"] = o.Lemmas
	}
	if o.Sentences != nil {
		toSerialize["sentences"] = o.Sentences
	}
	if o.Thesaurus != nil {
		toSerialize["thesaurus"] = o.Thesaurus
	}
	if o.Translations != nil {
		toSerialize["translations"] = o.Translations
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersResults struct {
	value *FiltersResults
	isSet bool
}

func (v NullableFiltersResults) Get() *FiltersResults {
	return v.value
}

func (v *NullableFiltersResults) Set(val *FiltersResults) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersResults) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersResults(val *FiltersResults) *NullableFiltersResults {
	return &NullableFiltersResults{value: val, isSet: true}
}

func (v NullableFiltersResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

