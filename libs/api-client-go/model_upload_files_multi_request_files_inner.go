/*
Daytona

Daytona AI platform API Docs

API version: 1.0
Contact: support@daytona.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daytonaapiclient

import (
	"encoding/json"
	"os"
)

// checks if the UploadFilesMultiRequestFilesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadFilesMultiRequestFilesInner{}

// UploadFilesMultiRequestFilesInner struct for UploadFilesMultiRequestFilesInner
type UploadFilesMultiRequestFilesInner struct {
	Path *string   `json:"path,omitempty"`
	File **os.File `json:"file,omitempty"`
}

// NewUploadFilesMultiRequestFilesInner instantiates a new UploadFilesMultiRequestFilesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadFilesMultiRequestFilesInner() *UploadFilesMultiRequestFilesInner {
	this := UploadFilesMultiRequestFilesInner{}
	return &this
}

// NewUploadFilesMultiRequestFilesInnerWithDefaults instantiates a new UploadFilesMultiRequestFilesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadFilesMultiRequestFilesInnerWithDefaults() *UploadFilesMultiRequestFilesInner {
	this := UploadFilesMultiRequestFilesInner{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *UploadFilesMultiRequestFilesInner) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFilesMultiRequestFilesInner) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *UploadFilesMultiRequestFilesInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *UploadFilesMultiRequestFilesInner) SetPath(v string) {
	o.Path = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *UploadFilesMultiRequestFilesInner) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFilesMultiRequestFilesInner) GetFileOk() (**os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *UploadFilesMultiRequestFilesInner) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *UploadFilesMultiRequestFilesInner) SetFile(v *os.File) {
	o.File = &v
}

func (o UploadFilesMultiRequestFilesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadFilesMultiRequestFilesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	return toSerialize, nil
}

type NullableUploadFilesMultiRequestFilesInner struct {
	value *UploadFilesMultiRequestFilesInner
	isSet bool
}

func (v NullableUploadFilesMultiRequestFilesInner) Get() *UploadFilesMultiRequestFilesInner {
	return v.value
}

func (v *NullableUploadFilesMultiRequestFilesInner) Set(val *UploadFilesMultiRequestFilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadFilesMultiRequestFilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadFilesMultiRequestFilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadFilesMultiRequestFilesInner(val *UploadFilesMultiRequestFilesInner) *NullableUploadFilesMultiRequestFilesInner {
	return &NullableUploadFilesMultiRequestFilesInner{value: val, isSet: true}
}

func (v NullableUploadFilesMultiRequestFilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadFilesMultiRequestFilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
