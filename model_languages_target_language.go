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

// LanguagesTargetLanguage Translation language of the results
type LanguagesTargetLanguage struct {
	// IANA language code
	Id *string `json:"id,omitempty"`
	// Language label.
	Language *string `json:"language,omitempty"`
}

// NewLanguagesTargetLanguage instantiates a new LanguagesTargetLanguage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguagesTargetLanguage() *LanguagesTargetLanguage {
	this := LanguagesTargetLanguage{}
	return &this
}

// NewLanguagesTargetLanguageWithDefaults instantiates a new LanguagesTargetLanguage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguagesTargetLanguageWithDefaults() *LanguagesTargetLanguage {
	this := LanguagesTargetLanguage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LanguagesTargetLanguage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguagesTargetLanguage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LanguagesTargetLanguage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LanguagesTargetLanguage) SetId(v string) {
	o.Id = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *LanguagesTargetLanguage) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguagesTargetLanguage) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *LanguagesTargetLanguage) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *LanguagesTargetLanguage) SetLanguage(v string) {
	o.Language = &v
}

func (o LanguagesTargetLanguage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	return json.Marshal(toSerialize)
}

type NullableLanguagesTargetLanguage struct {
	value *LanguagesTargetLanguage
	isSet bool
}

func (v NullableLanguagesTargetLanguage) Get() *LanguagesTargetLanguage {
	return v.value
}

func (v *NullableLanguagesTargetLanguage) Set(val *LanguagesTargetLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguagesTargetLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguagesTargetLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguagesTargetLanguage(val *LanguagesTargetLanguage) *NullableLanguagesTargetLanguage {
	return &NullableLanguagesTargetLanguage{value: val, isSet: true}
}

func (v NullableLanguagesTargetLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguagesTargetLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


