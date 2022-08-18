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

// ErrorSchema Schema for Errors.
type ErrorSchema struct {
	// A human-readable explanation of the problem occurrence.
	Error string `json:"error"`
}

// NewErrorSchema instantiates a new ErrorSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorSchema(error_ string) *ErrorSchema {
	this := ErrorSchema{}
	this.Error = error_
	return &this
}

// NewErrorSchemaWithDefaults instantiates a new ErrorSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorSchemaWithDefaults() *ErrorSchema {
	this := ErrorSchema{}
	return &this
}

// GetError returns the Error field value
func (o *ErrorSchema) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ErrorSchema) GetErrorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ErrorSchema) SetError(v string) {
	o.Error = v
}

func (o ErrorSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableErrorSchema struct {
	value *ErrorSchema
	isSet bool
}

func (v NullableErrorSchema) Get() *ErrorSchema {
	return v.value
}

func (v *NullableErrorSchema) Set(val *ErrorSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSchema(val *ErrorSchema) *NullableErrorSchema {
	return &NullableErrorSchema{value: val, isSet: true}
}

func (v NullableErrorSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


