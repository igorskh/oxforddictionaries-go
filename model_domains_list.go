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

// DomainsList struct for DomainsList
type DomainsList struct {
}

// NewDomainsList instantiates a new DomainsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainsList() *DomainsList {
	this := DomainsList{}
	return &this
}

// NewDomainsListWithDefaults instantiates a new DomainsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainsListWithDefaults() *DomainsList {
	this := DomainsList{}
	return &this
}

func (o DomainsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableDomainsList struct {
	value *DomainsList
	isSet bool
}

func (v NullableDomainsList) Get() *DomainsList {
	return v.value
}

func (v *NullableDomainsList) Set(val *DomainsList) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainsList) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainsList(val *DomainsList) *NullableDomainsList {
	return &NullableDomainsList{value: val, isSet: true}
}

func (v NullableDomainsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


