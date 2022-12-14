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

// InflectedForm Description of an Inflected form.
type InflectedForm struct {
	Domains *DomainsList `json:"domains,omitempty"`
	// The different forms are correlated with meanings or functions which we text as 'features'
	GrammaticalFeatures *[]map[string]interface{} `json:"grammaticalFeatures,omitempty"`
	// Canonical form of an inflection
	InflectedForm string `json:"inflectedForm"`
	LexicalCategory *LexicalCategory `json:"lexicalCategory,omitempty"`
	// A list of possible pronunciations of a word
	Pronunciations *[]map[string]interface{} `json:"pronunciations,omitempty"`
	Regions *RegionsList `json:"regions,omitempty"`
	Registers *RegistersList `json:"registers,omitempty"`
}

// NewInflectedForm instantiates a new InflectedForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInflectedForm(inflectedForm string) *InflectedForm {
	this := InflectedForm{}
	this.InflectedForm = inflectedForm
	return &this
}

// NewInflectedFormWithDefaults instantiates a new InflectedForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInflectedFormWithDefaults() *InflectedForm {
	this := InflectedForm{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *InflectedForm) GetDomains() DomainsList {
	if o == nil || o.Domains == nil {
		var ret DomainsList
		return ret
	}
	return *o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InflectedForm) GetDomainsOk() (*DomainsList, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *InflectedForm) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given DomainsList and assigns it to the Domains field.
func (o *InflectedForm) SetDomains(v DomainsList) {
	o.Domains = &v
}

// GetGrammaticalFeatures returns the GrammaticalFeatures field value if set, zero value otherwise.
func (o *InflectedForm) GetGrammaticalFeatures() []map[string]interface{} {
	if o == nil || o.GrammaticalFeatures == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.GrammaticalFeatures
}

// GetGrammaticalFeaturesOk returns a tuple with the GrammaticalFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InflectedForm) GetGrammaticalFeaturesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.GrammaticalFeatures == nil {
		return nil, false
	}
	return o.GrammaticalFeatures, true
}

// HasGrammaticalFeatures returns a boolean if a field has been set.
func (o *InflectedForm) HasGrammaticalFeatures() bool {
	if o != nil && o.GrammaticalFeatures != nil {
		return true
	}

	return false
}

// SetGrammaticalFeatures gets a reference to the given []map[string]interface{} and assigns it to the GrammaticalFeatures field.
func (o *InflectedForm) SetGrammaticalFeatures(v []map[string]interface{}) {
	o.GrammaticalFeatures = &v
}

// GetInflectedForm returns the InflectedForm field value
func (o *InflectedForm) GetInflectedForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InflectedForm
}

// GetInflectedFormOk returns a tuple with the InflectedForm field value
// and a boolean to check if the value has been set.
func (o *InflectedForm) GetInflectedFormOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InflectedForm, true
}

// SetInflectedForm sets field value
func (o *InflectedForm) SetInflectedForm(v string) {
	o.InflectedForm = v
}

// GetLexicalCategory returns the LexicalCategory field value if set, zero value otherwise.
func (o *InflectedForm) GetLexicalCategory() LexicalCategory {
	if o == nil || o.LexicalCategory == nil {
		var ret LexicalCategory
		return ret
	}
	return *o.LexicalCategory
}

// GetLexicalCategoryOk returns a tuple with the LexicalCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InflectedForm) GetLexicalCategoryOk() (*LexicalCategory, bool) {
	if o == nil || o.LexicalCategory == nil {
		return nil, false
	}
	return o.LexicalCategory, true
}

// HasLexicalCategory returns a boolean if a field has been set.
func (o *InflectedForm) HasLexicalCategory() bool {
	if o != nil && o.LexicalCategory != nil {
		return true
	}

	return false
}

// SetLexicalCategory gets a reference to the given LexicalCategory and assigns it to the LexicalCategory field.
func (o *InflectedForm) SetLexicalCategory(v LexicalCategory) {
	o.LexicalCategory = &v
}

// GetPronunciations returns the Pronunciations field value if set, zero value otherwise.
func (o *InflectedForm) GetPronunciations() []map[string]interface{} {
	if o == nil || o.Pronunciations == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Pronunciations
}

// GetPronunciationsOk returns a tuple with the Pronunciations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InflectedForm) GetPronunciationsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Pronunciations == nil {
		return nil, false
	}
	return o.Pronunciations, true
}

// HasPronunciations returns a boolean if a field has been set.
func (o *InflectedForm) HasPronunciations() bool {
	if o != nil && o.Pronunciations != nil {
		return true
	}

	return false
}

// SetPronunciations gets a reference to the given []map[string]interface{} and assigns it to the Pronunciations field.
func (o *InflectedForm) SetPronunciations(v []map[string]interface{}) {
	o.Pronunciations = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *InflectedForm) GetRegions() RegionsList {
	if o == nil || o.Regions == nil {
		var ret RegionsList
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InflectedForm) GetRegionsOk() (*RegionsList, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *InflectedForm) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given RegionsList and assigns it to the Regions field.
func (o *InflectedForm) SetRegions(v RegionsList) {
	o.Regions = &v
}

// GetRegisters returns the Registers field value if set, zero value otherwise.
func (o *InflectedForm) GetRegisters() RegistersList {
	if o == nil || o.Registers == nil {
		var ret RegistersList
		return ret
	}
	return *o.Registers
}

// GetRegistersOk returns a tuple with the Registers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InflectedForm) GetRegistersOk() (*RegistersList, bool) {
	if o == nil || o.Registers == nil {
		return nil, false
	}
	return o.Registers, true
}

// HasRegisters returns a boolean if a field has been set.
func (o *InflectedForm) HasRegisters() bool {
	if o != nil && o.Registers != nil {
		return true
	}

	return false
}

// SetRegisters gets a reference to the given RegistersList and assigns it to the Registers field.
func (o *InflectedForm) SetRegisters(v RegistersList) {
	o.Registers = &v
}

func (o InflectedForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.GrammaticalFeatures != nil {
		toSerialize["grammaticalFeatures"] = o.GrammaticalFeatures
	}
	if true {
		toSerialize["inflectedForm"] = o.InflectedForm
	}
	if o.LexicalCategory != nil {
		toSerialize["lexicalCategory"] = o.LexicalCategory
	}
	if o.Pronunciations != nil {
		toSerialize["pronunciations"] = o.Pronunciations
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	if o.Registers != nil {
		toSerialize["registers"] = o.Registers
	}
	return json.Marshal(toSerialize)
}

type NullableInflectedForm struct {
	value *InflectedForm
	isSet bool
}

func (v NullableInflectedForm) Get() *InflectedForm {
	return v.value
}

func (v *NullableInflectedForm) Set(val *InflectedForm) {
	v.value = val
	v.isSet = true
}

func (v NullableInflectedForm) IsSet() bool {
	return v.isSet
}

func (v *NullableInflectedForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInflectedForm(val *InflectedForm) *NullableInflectedForm {
	return &NullableInflectedForm{value: val, isSet: true}
}

func (v NullableInflectedForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInflectedForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


