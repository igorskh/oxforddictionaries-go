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

// RetrieveEntry Schema for the 'entries' endpoints
type RetrieveEntry struct {
	// Additional Information provided by OUP
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// A list of entries and all the data related to them
	Results *[]HeadwordEntry `json:"results,omitempty"`
}

// NewRetrieveEntry instantiates a new RetrieveEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrieveEntry() *RetrieveEntry {
	this := RetrieveEntry{}
	return &this
}

// NewRetrieveEntryWithDefaults instantiates a new RetrieveEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrieveEntryWithDefaults() *RetrieveEntry {
	this := RetrieveEntry{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RetrieveEntry) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveEntry) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RetrieveEntry) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RetrieveEntry) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *RetrieveEntry) GetResults() []HeadwordEntry {
	if o == nil || o.Results == nil {
		var ret []HeadwordEntry
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveEntry) GetResultsOk() (*[]HeadwordEntry, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *RetrieveEntry) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []HeadwordEntry and assigns it to the Results field.
func (o *RetrieveEntry) SetResults(v []HeadwordEntry) {
	o.Results = &v
}

func (o RetrieveEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableRetrieveEntry struct {
	value *RetrieveEntry
	isSet bool
}

func (v NullableRetrieveEntry) Get() *RetrieveEntry {
	return v.value
}

func (v *NullableRetrieveEntry) Set(val *RetrieveEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveEntry(val *RetrieveEntry) *NullableRetrieveEntry {
	return &NullableRetrieveEntry{value: val, isSet: true}
}

func (v NullableRetrieveEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetrieveEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

