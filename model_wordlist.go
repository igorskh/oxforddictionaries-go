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

// Wordlist Schema for wordlist endpoint.
type Wordlist struct {
	Metadata *WordlistMetadata `json:"metadata,omitempty"`
	// A list of found words
	Results *[]WordlistResults `json:"results,omitempty"`
}

// NewWordlist instantiates a new Wordlist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWordlist() *Wordlist {
	this := Wordlist{}
	return &this
}

// NewWordlistWithDefaults instantiates a new Wordlist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWordlistWithDefaults() *Wordlist {
	this := Wordlist{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Wordlist) GetMetadata() WordlistMetadata {
	if o == nil || o.Metadata == nil {
		var ret WordlistMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wordlist) GetMetadataOk() (*WordlistMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Wordlist) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given WordlistMetadata and assigns it to the Metadata field.
func (o *Wordlist) SetMetadata(v WordlistMetadata) {
	o.Metadata = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *Wordlist) GetResults() []WordlistResults {
	if o == nil || o.Results == nil {
		var ret []WordlistResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wordlist) GetResultsOk() (*[]WordlistResults, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *Wordlist) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []WordlistResults and assigns it to the Results field.
func (o *Wordlist) SetResults(v []WordlistResults) {
	o.Results = &v
}

func (o Wordlist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableWordlist struct {
	value *Wordlist
	isSet bool
}

func (v NullableWordlist) Get() *Wordlist {
	return v.value
}

func (v *NullableWordlist) Set(val *Wordlist) {
	v.value = val
	v.isSet = true
}

func (v NullableWordlist) IsSet() bool {
	return v.isSet
}

func (v *NullableWordlist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWordlist(val *Wordlist) *NullableWordlist {
	return &NullableWordlist{value: val, isSet: true}
}

func (v NullableWordlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWordlist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


