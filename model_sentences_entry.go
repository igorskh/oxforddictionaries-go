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

// SentencesEntry Group of lexicalEntries related to a specific result in the Sentences endpoint for a given word ID.
type SentencesEntry struct {
	// The identifier of a word
	Id string `json:"id"`
	// IANA language code
	Language string `json:"language"`
	// A grouping of various senses in a specific language, and a lexical category that relates to a word
	LexicalEntries []SentencesLexicalEntry `json:"lexicalEntries"`
	// The json object type. Could be 'headword', 'inflection' or 'phrase'
	Type *string `json:"type,omitempty"`
	// (DEPRECATED) A given written or spoken realisation of an entry, lowercased.
	Word string `json:"word"`
}

// NewSentencesEntry instantiates a new SentencesEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSentencesEntry(id string, language string, lexicalEntries []SentencesLexicalEntry, word string) *SentencesEntry {
	this := SentencesEntry{}
	this.Id = id
	this.Language = language
	this.LexicalEntries = lexicalEntries
	this.Word = word
	return &this
}

// NewSentencesEntryWithDefaults instantiates a new SentencesEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSentencesEntryWithDefaults() *SentencesEntry {
	this := SentencesEntry{}
	return &this
}

// GetId returns the Id field value
func (o *SentencesEntry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SentencesEntry) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SentencesEntry) SetId(v string) {
	o.Id = v
}

// GetLanguage returns the Language field value
func (o *SentencesEntry) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *SentencesEntry) GetLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *SentencesEntry) SetLanguage(v string) {
	o.Language = v
}

// GetLexicalEntries returns the LexicalEntries field value
func (o *SentencesEntry) GetLexicalEntries() []SentencesLexicalEntry {
	if o == nil {
		var ret []SentencesLexicalEntry
		return ret
	}

	return o.LexicalEntries
}

// GetLexicalEntriesOk returns a tuple with the LexicalEntries field value
// and a boolean to check if the value has been set.
func (o *SentencesEntry) GetLexicalEntriesOk() (*[]SentencesLexicalEntry, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LexicalEntries, true
}

// SetLexicalEntries sets field value
func (o *SentencesEntry) SetLexicalEntries(v []SentencesLexicalEntry) {
	o.LexicalEntries = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SentencesEntry) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SentencesEntry) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SentencesEntry) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SentencesEntry) SetType(v string) {
	o.Type = &v
}

// GetWord returns the Word field value
func (o *SentencesEntry) GetWord() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Word
}

// GetWordOk returns a tuple with the Word field value
// and a boolean to check if the value has been set.
func (o *SentencesEntry) GetWordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Word, true
}

// SetWord sets field value
func (o *SentencesEntry) SetWord(v string) {
	o.Word = v
}

func (o SentencesEntry) MarshalJSON() ([]byte, error) {
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

type NullableSentencesEntry struct {
	value *SentencesEntry
	isSet bool
}

func (v NullableSentencesEntry) Get() *SentencesEntry {
	return v.value
}

func (v *NullableSentencesEntry) Set(val *SentencesEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableSentencesEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableSentencesEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSentencesEntry(val *SentencesEntry) *NullableSentencesEntry {
	return &NullableSentencesEntry{value: val, isSet: true}
}

func (v NullableSentencesEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSentencesEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


