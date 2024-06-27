/*
Twitter API

This is a sample server Twitter API server.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TweetsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetsPostRequest{}

// TweetsPostRequest struct for TweetsPostRequest
type TweetsPostRequest struct {
	// Reference to User ID
	UserId int `json:"user_id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

type _TweetsPostRequest TweetsPostRequest

// NewTweetsPostRequest instantiates a new TweetsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetsPostRequest(userId int, title string, body string) *TweetsPostRequest {
	this := TweetsPostRequest{}
	this.UserId = userId
	this.Title = title
	this.Body = body
	return &this
}

// NewTweetsPostRequestWithDefaults instantiates a new TweetsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetsPostRequestWithDefaults() *TweetsPostRequest {
	this := TweetsPostRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *TweetsPostRequest) GetUserId() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *TweetsPostRequest) GetUserIdOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *TweetsPostRequest) SetUserId(v int) {
	o.UserId = v
}

// GetTitle returns the Title field value
func (o *TweetsPostRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TweetsPostRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TweetsPostRequest) SetTitle(v string) {
	o.Title = v
}

// GetBody returns the Body field value
func (o *TweetsPostRequest) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *TweetsPostRequest) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *TweetsPostRequest) SetBody(v string) {
	o.Body = v
}

func (o TweetsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["title"] = o.Title
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

func (o *TweetsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
		"title",
		"body",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTweetsPostRequest := _TweetsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetsPostRequest)

	if err != nil {
		return err
	}

	*o = TweetsPostRequest(varTweetsPostRequest)

	return err
}

type NullableTweetsPostRequest struct {
	value *TweetsPostRequest
	isSet bool
}

func (v NullableTweetsPostRequest) Get() *TweetsPostRequest {
	return v.value
}

func (v *NullableTweetsPostRequest) Set(val *TweetsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetsPostRequest(val *TweetsPostRequest) *NullableTweetsPostRequest {
	return &NullableTweetsPostRequest{value: val, isSet: true}
}

func (v NullableTweetsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


