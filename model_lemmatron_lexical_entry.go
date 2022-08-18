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

// LemmatronLexicalEntry Description of an entry for a particular part of speech and grammatical features
type LemmatronLexicalEntry struct {
	// The different forms are correlated with meanings or functions which we text as 'features'
	GrammaticalFeatures *[]map[string]interface{} `json:"grammaticalFeatures,omitempty"`
	// A grouping of the modifications of a word to express different grammatical categories
	InflectionOf []map[string]interface{} `json:"inflectionOf"`
	// IANA language code
	Language string `json:"language"`
	LexicalCategory LexicalCategory `json:"lexicalCategory"`
	// A given written or spoken realisation of an entry.
	Text string `json:"text"`
}

// NewLemmatronLexicalEntry instantiates a new LemmatronLexicalEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLemmatronLexicalEntry(inflectionOf []map[string]interface{}, language string, lexicalCategory LexicalCategory, text string) *LemmatronLexicalEntry {
	this := LemmatronLexicalEntry{}
	this.InflectionOf = inflectionOf
	this.Language = language
	this.LexicalCategory = lexicalCategory
	this.Text = text
	return &this
}

// NewLemmatronLexicalEntryWithDefaults instantiates a new LemmatronLexicalEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLemmatronLexicalEntryWithDefaults() *LemmatronLexicalEntry {
	this := LemmatronLexicalEntry{}
	return &this
}

// GetGrammaticalFeatures returns the GrammaticalFeatures field value if set, zero value otherwise.
func (o *LemmatronLexicalEntry) GetGrammaticalFeatures() []map[string]interface{} {
	if o == nil || o.GrammaticalFeatures == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.GrammaticalFeatures
}

// GetGrammaticalFeaturesOk returns a tuple with the GrammaticalFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LemmatronLexicalEntry) GetGrammaticalFeaturesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.GrammaticalFeatures == nil {
		return nil, false
	}
	return o.GrammaticalFeatures, true
}

// HasGrammaticalFeatures returns a boolean if a field has been set.
func (o *LemmatronLexicalEntry) HasGrammaticalFeatures() bool {
	if o != nil && o.GrammaticalFeatures != nil {
		return true
	}

	return false
}

// SetGrammaticalFeatures gets a reference to the given []map[string]interface{} and assigns it to the GrammaticalFeatures field.
func (o *LemmatronLexicalEntry) SetGrammaticalFeatures(v []map[string]interface{}) {
	o.GrammaticalFeatures = &v
}

// GetInflectionOf returns the InflectionOf field value
func (o *LemmatronLexicalEntry) GetInflectionOf() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.InflectionOf
}

// GetInflectionOfOk returns a tuple with the InflectionOf field value
// and a boolean to check if the value has been set.
func (o *LemmatronLexicalEntry) GetInflectionOfOk() (*[]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InflectionOf, true
}

// SetInflectionOf sets field value
func (o *LemmatronLexicalEntry) SetInflectionOf(v []map[string]interface{}) {
	o.InflectionOf = v
}

// GetLanguage returns the Language field value
func (o *LemmatronLexicalEntry) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *LemmatronLexicalEntry) GetLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *LemmatronLexicalEntry) SetLanguage(v string) {
	o.Language = v
}

// GetLexicalCategory returns the LexicalCategory field value
func (o *LemmatronLexicalEntry) GetLexicalCategory() LexicalCategory {
	if o == nil {
		var ret LexicalCategory
		return ret
	}

	return o.LexicalCategory
}

// GetLexicalCategoryOk returns a tuple with the LexicalCategory field value
// and a boolean to check if the value has been set.
func (o *LemmatronLexicalEntry) GetLexicalCategoryOk() (*LexicalCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LexicalCategory, true
}

// SetLexicalCategory sets field value
func (o *LemmatronLexicalEntry) SetLexicalCategory(v LexicalCategory) {
	o.LexicalCategory = v
}

// GetText returns the Text field value
func (o *LemmatronLexicalEntry) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *LemmatronLexicalEntry) GetTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *LemmatronLexicalEntry) SetText(v string) {
	o.Text = v
}

func (o LemmatronLexicalEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GrammaticalFeatures != nil {
		toSerialize["grammaticalFeatures"] = o.GrammaticalFeatures
	}
	if true {
		toSerialize["inflectionOf"] = o.InflectionOf
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["lexicalCategory"] = o.LexicalCategory
	}
	if true {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableLemmatronLexicalEntry struct {
	value *LemmatronLexicalEntry
	isSet bool
}

func (v NullableLemmatronLexicalEntry) Get() *LemmatronLexicalEntry {
	return v.value
}

func (v *NullableLemmatronLexicalEntry) Set(val *LemmatronLexicalEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableLemmatronLexicalEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableLemmatronLexicalEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLemmatronLexicalEntry(val *LemmatronLexicalEntry) *NullableLemmatronLexicalEntry {
	return &NullableLemmatronLexicalEntry{value: val, isSet: true}
}

func (v NullableLemmatronLexicalEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLemmatronLexicalEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

