/*
Twitter API

This is a sample server Twitter API server.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TweetsGet400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetsGet400Response{}

// TweetsGet400Response struct for TweetsGet400Response
type TweetsGet400Response struct {
	Message *string `json:"message,omitempty"`
}

// NewTweetsGet400Response instantiates a new TweetsGet400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetsGet400Response() *TweetsGet400Response {
	this := TweetsGet400Response{}
	return &this
}

// NewTweetsGet400ResponseWithDefaults instantiates a new TweetsGet400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetsGet400ResponseWithDefaults() *TweetsGet400Response {
	this := TweetsGet400Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TweetsGet400Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetsGet400Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TweetsGet400Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TweetsGet400Response) SetMessage(v string) {
	o.Message = &v
}

func (o TweetsGet400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetsGet400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableTweetsGet400Response struct {
	value *TweetsGet400Response
	isSet bool
}

func (v NullableTweetsGet400Response) Get() *TweetsGet400Response {
	return v.value
}

func (v *NullableTweetsGet400Response) Set(val *TweetsGet400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetsGet400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetsGet400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetsGet400Response(val *TweetsGet400Response) *NullableTweetsGet400Response {
	return &NullableTweetsGet400Response{value: val, isSet: true}
}

func (v NullableTweetsGet400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetsGet400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


