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

// HeadwordLemmatron Description of an inflected form of a word
type HeadwordLemmatron struct {
	// The identifier of a word
	Id string `json:"id"`
	// IANA language code
	Language string `json:"language"`
	// A grouping of various senses in a specific language, and a lexical category that relates to a word
	LexicalEntries []LemmatronLexicalEntry `json:"lexicalEntries"`
	// The json object type. Could be 'headword', 'inflection' or 'phrase'
	Type *string `json:"type,omitempty"`
	// (DEPRECATED) A given written or spoken realisation of an entry, lowercased.
	Word string `json:"word"`
}

// NewHeadwordLemmatron instantiates a new HeadwordLemmatron object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeadwordLemmatron(id string, language string, lexicalEntries []LemmatronLexicalEntry, word string) *HeadwordLemmatron {
	this := HeadwordLemmatron{}
	this.Id = id
	this.Language = language
	this.LexicalEntries = lexicalEntries
	this.Word = word
	return &this
}

// NewHeadwordLemmatronWithDefaults instantiates a new HeadwordLemmatron object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeadwordLemmatronWithDefaults() *HeadwordLemmatron {
	this := HeadwordLemmatron{}
	return &this
}

// GetId returns the Id field value
func (o *HeadwordLemmatron) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HeadwordLemmatron) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HeadwordLemmatron) SetId(v string) {
	o.Id = v
}

// GetLanguage returns the Language field value
func (o *HeadwordLemmatron) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *HeadwordLemmatron) GetLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *HeadwordLemmatron) SetLanguage(v string) {
	o.Language = v
}

// GetLexicalEntries returns the LexicalEntries field value
func (o *HeadwordLemmatron) GetLexicalEntries() []LemmatronLexicalEntry {
	if o == nil {
		var ret []LemmatronLexicalEntry
		return ret
	}

	return o.LexicalEntries
}

// GetLexicalEntriesOk returns a tuple with the LexicalEntries field value
// and a boolean to check if the value has been set.
func (o *HeadwordLemmatron) GetLexicalEntriesOk() (*[]LemmatronLexicalEntry, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LexicalEntries, true
}

// SetLexicalEntries sets field value
func (o *HeadwordLemmatron) SetLexicalEntries(v []LemmatronLexicalEntry) {
	o.LexicalEntries = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HeadwordLemmatron) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeadwordLemmatron) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HeadwordLemmatron) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HeadwordLemmatron) SetType(v string) {
	o.Type = &v
}

// GetWord returns the Word field value
func (o *HeadwordLemmatron) GetWord() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Word
}

// GetWordOk returns a tuple with the Word field value
// and a boolean to check if the value has been set.
func (o *HeadwordLemmatron) GetWordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Word, true
}

// SetWord sets field value
func (o *HeadwordLemmatron) SetWord(v string) {
	o.Word = v
}

func (o HeadwordLemmatron) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["lexicalEntries"] = o.LexicalEntries
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["word"] = o.Word
	}
	return json.Marshal(toSerialize)
}

type NullableHeadwordLemmatron struct {
	value *HeadwordLemmatron
	isSet bool
}

func (v NullableHeadwordLemmatron) Get() *HeadwordLemmatron {
	return v.value
}

func (v *NullableHeadwordLemmatron) Set(val *HeadwordLemmatron) {
	v.value = val
	v.isSet = true
}

func (v NullableHeadwordLemmatron) IsSet() bool {
	return v.isSet
}

func (v *NullableHeadwordLemmatron) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeadwordLemmatron(val *HeadwordLemmatron) *NullableHeadwordLemmatron {
	return &NullableHeadwordLemmatron{value: val, isSet: true}
}

func (v NullableHeadwordLemmatron) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeadwordLemmatron) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


