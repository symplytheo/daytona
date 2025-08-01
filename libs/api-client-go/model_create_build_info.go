/*
Daytona

Daytona AI platform API Docs

API version: 1.0
Contact: support@daytona.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateBuildInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBuildInfo{}

// CreateBuildInfo struct for CreateBuildInfo
type CreateBuildInfo struct {
	// The Dockerfile content used for the build
	DockerfileContent string `json:"dockerfileContent"`
	// The context hashes used for the build
	ContextHashes []string `json:"contextHashes,omitempty"`
}

type _CreateBuildInfo CreateBuildInfo

// NewCreateBuildInfo instantiates a new CreateBuildInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBuildInfo(dockerfileContent string) *CreateBuildInfo {
	this := CreateBuildInfo{}
	this.DockerfileContent = dockerfileContent
	return &this
}

// NewCreateBuildInfoWithDefaults instantiates a new CreateBuildInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBuildInfoWithDefaults() *CreateBuildInfo {
	this := CreateBuildInfo{}
	return &this
}

// GetDockerfileContent returns the DockerfileContent field value
func (o *CreateBuildInfo) GetDockerfileContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerfileContent
}

// GetDockerfileContentOk returns a tuple with the DockerfileContent field value
// and a boolean to check if the value has been set.
func (o *CreateBuildInfo) GetDockerfileContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerfileContent, true
}

// SetDockerfileContent sets field value
func (o *CreateBuildInfo) SetDockerfileContent(v string) {
	o.DockerfileContent = v
}

// GetContextHashes returns the ContextHashes field value if set, zero value otherwise.
func (o *CreateBuildInfo) GetContextHashes() []string {
	if o == nil || IsNil(o.ContextHashes) {
		var ret []string
		return ret
	}
	return o.ContextHashes
}

// GetContextHashesOk returns a tuple with the ContextHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuildInfo) GetContextHashesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContextHashes) {
		return nil, false
	}
	return o.ContextHashes, true
}

// HasContextHashes returns a boolean if a field has been set.
func (o *CreateBuildInfo) HasContextHashes() bool {
	if o != nil && !IsNil(o.ContextHashes) {
		return true
	}

	return false
}

// SetContextHashes gets a reference to the given []string and assigns it to the ContextHashes field.
func (o *CreateBuildInfo) SetContextHashes(v []string) {
	o.ContextHashes = v
}

func (o CreateBuildInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBuildInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dockerfileContent"] = o.DockerfileContent
	if !IsNil(o.ContextHashes) {
		toSerialize["contextHashes"] = o.ContextHashes
	}
	return toSerialize, nil
}

func (o *CreateBuildInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dockerfileContent",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateBuildInfo := _CreateBuildInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateBuildInfo)

	if err != nil {
		return err
	}

	*o = CreateBuildInfo(varCreateBuildInfo)

	return err
}

type NullableCreateBuildInfo struct {
	value *CreateBuildInfo
	isSet bool
}

func (v NullableCreateBuildInfo) Get() *CreateBuildInfo {
	return v.value
}

func (v *NullableCreateBuildInfo) Set(val *CreateBuildInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBuildInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBuildInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBuildInfo(val *CreateBuildInfo) *NullableCreateBuildInfo {
	return &NullableCreateBuildInfo{value: val, isSet: true}
}

func (v NullableCreateBuildInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBuildInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
