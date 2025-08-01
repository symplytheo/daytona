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

// checks if the GitDeleteBranchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitDeleteBranchRequest{}

// GitDeleteBranchRequest struct for GitDeleteBranchRequest
type GitDeleteBranchRequest struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type _GitDeleteBranchRequest GitDeleteBranchRequest

// NewGitDeleteBranchRequest instantiates a new GitDeleteBranchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitDeleteBranchRequest(path string, name string) *GitDeleteBranchRequest {
	this := GitDeleteBranchRequest{}
	this.Path = path
	this.Name = name
	return &this
}

// NewGitDeleteBranchRequestWithDefaults instantiates a new GitDeleteBranchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitDeleteBranchRequestWithDefaults() *GitDeleteBranchRequest {
	this := GitDeleteBranchRequest{}
	return &this
}

// GetPath returns the Path field value
func (o *GitDeleteBranchRequest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *GitDeleteBranchRequest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *GitDeleteBranchRequest) SetPath(v string) {
	o.Path = v
}

// GetName returns the Name field value
func (o *GitDeleteBranchRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitDeleteBranchRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GitDeleteBranchRequest) SetName(v string) {
	o.Name = v
}

func (o GitDeleteBranchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitDeleteBranchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GitDeleteBranchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
		"name",
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

	varGitDeleteBranchRequest := _GitDeleteBranchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGitDeleteBranchRequest)

	if err != nil {
		return err
	}

	*o = GitDeleteBranchRequest(varGitDeleteBranchRequest)

	return err
}

type NullableGitDeleteBranchRequest struct {
	value *GitDeleteBranchRequest
	isSet bool
}

func (v NullableGitDeleteBranchRequest) Get() *GitDeleteBranchRequest {
	return v.value
}

func (v *NullableGitDeleteBranchRequest) Set(val *GitDeleteBranchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGitDeleteBranchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGitDeleteBranchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitDeleteBranchRequest(val *GitDeleteBranchRequest) *NullableGitDeleteBranchRequest {
	return &NullableGitDeleteBranchRequest{value: val, isSet: true}
}

func (v NullableGitDeleteBranchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitDeleteBranchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
