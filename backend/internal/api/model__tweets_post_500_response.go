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

// checks if the TweetsPost500Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetsPost500Response{}

// TweetsPost500Response struct for TweetsPost500Response
type TweetsPost500Response struct {
	Message *string `json:"message,omitempty"`
}

// NewTweetsPost500Response instantiates a new TweetsPost500Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetsPost500Response() *TweetsPost500Response {
	this := TweetsPost500Response{}
	return &this
}

// NewTweetsPost500ResponseWithDefaults instantiates a new TweetsPost500Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetsPost500ResponseWithDefaults() *TweetsPost500Response {
	this := TweetsPost500Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TweetsPost500Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetsPost500Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TweetsPost500Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TweetsPost500Response) SetMessage(v string) {
	o.Message = &v
}

func (o TweetsPost500Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetsPost500Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableTweetsPost500Response struct {
	value *TweetsPost500Response
	isSet bool
}

func (v NullableTweetsPost500Response) Get() *TweetsPost500Response {
	return v.value
}

func (v *NullableTweetsPost500Response) Set(val *TweetsPost500Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetsPost500Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetsPost500Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetsPost500Response(val *TweetsPost500Response) *NullableTweetsPost500Response {
	return &NullableTweetsPost500Response{value: val, isSet: true}
}

func (v NullableTweetsPost500Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetsPost500Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

