/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RegisteredModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisteredModel{}

// RegisteredModel A registered model in model registry. A registered model has ModelVersion children.
type RegisteredModel struct {
	// User provided custom properties which are not defined by its type.
	CustomProperties *map[string]MetadataValue `json:"customProperties,omitempty"`
	// An optional description about the resource.
	Description *string `json:"description,omitempty"`
	// The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance.
	ExternalID *string `json:"externalID,omitempty"`
	// The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set.
	Name *string `json:"name,omitempty"`
	// Output only. The unique server generated id of the resource.
	Id *string `json:"id,omitempty"`
	// Output only. Create time of the resource in millisecond since epoch.
	CreateTimeSinceEpoch *string `json:"createTimeSinceEpoch,omitempty"`
	// Output only. Last update time of the resource since epoch in millisecond since epoch.
	LastUpdateTimeSinceEpoch *string `json:"lastUpdateTimeSinceEpoch,omitempty"`
}

// NewRegisteredModel instantiates a new RegisteredModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisteredModel() *RegisteredModel {
	this := RegisteredModel{}
	return &this
}

// NewRegisteredModelWithDefaults instantiates a new RegisteredModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisteredModelWithDefaults() *RegisteredModel {
	this := RegisteredModel{}
	return &this
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *RegisteredModel) GetCustomProperties() map[string]MetadataValue {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]MetadataValue
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredModel) GetCustomPropertiesOk() (*map[string]MetadataValue, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *RegisteredModel) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]MetadataValue and assigns it to the CustomProperties field.
func (o *RegisteredModel) SetCustomProperties(v map[string]MetadataValue) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RegisteredModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RegisteredModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RegisteredModel) SetDescription(v string) {
	o.Description = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *RegisteredModel) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredModel) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *RegisteredModel) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *RegisteredModel) SetExternalID(v string) {
	o.ExternalID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegisteredModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegisteredModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegisteredModel) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegisteredModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegisteredModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegisteredModel) SetId(v string) {
	o.Id = &v
}

// GetCreateTimeSinceEpoch returns the CreateTimeSinceEpoch field value if set, zero value otherwise.
func (o *RegisteredModel) GetCreateTimeSinceEpoch() string {
	if o == nil || IsNil(o.CreateTimeSinceEpoch) {
		var ret string
		return ret
	}
	return *o.CreateTimeSinceEpoch
}

// GetCreateTimeSinceEpochOk returns a tuple with the CreateTimeSinceEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredModel) GetCreateTimeSinceEpochOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTimeSinceEpoch) {
		return nil, false
	}
	return o.CreateTimeSinceEpoch, true
}

// HasCreateTimeSinceEpoch returns a boolean if a field has been set.
func (o *RegisteredModel) HasCreateTimeSinceEpoch() bool {
	if o != nil && !IsNil(o.CreateTimeSinceEpoch) {
		return true
	}

	return false
}

// SetCreateTimeSinceEpoch gets a reference to the given string and assigns it to the CreateTimeSinceEpoch field.
func (o *RegisteredModel) SetCreateTimeSinceEpoch(v string) {
	o.CreateTimeSinceEpoch = &v
}

// GetLastUpdateTimeSinceEpoch returns the LastUpdateTimeSinceEpoch field value if set, zero value otherwise.
func (o *RegisteredModel) GetLastUpdateTimeSinceEpoch() string {
	if o == nil || IsNil(o.LastUpdateTimeSinceEpoch) {
		var ret string
		return ret
	}
	return *o.LastUpdateTimeSinceEpoch
}

// GetLastUpdateTimeSinceEpochOk returns a tuple with the LastUpdateTimeSinceEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredModel) GetLastUpdateTimeSinceEpochOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdateTimeSinceEpoch) {
		return nil, false
	}
	return o.LastUpdateTimeSinceEpoch, true
}

// HasLastUpdateTimeSinceEpoch returns a boolean if a field has been set.
func (o *RegisteredModel) HasLastUpdateTimeSinceEpoch() bool {
	if o != nil && !IsNil(o.LastUpdateTimeSinceEpoch) {
		return true
	}

	return false
}

// SetLastUpdateTimeSinceEpoch gets a reference to the given string and assigns it to the LastUpdateTimeSinceEpoch field.
func (o *RegisteredModel) SetLastUpdateTimeSinceEpoch(v string) {
	o.LastUpdateTimeSinceEpoch = &v
}

func (o RegisteredModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisteredModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalID) {
		toSerialize["externalID"] = o.ExternalID
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreateTimeSinceEpoch) {
		toSerialize["createTimeSinceEpoch"] = o.CreateTimeSinceEpoch
	}
	if !IsNil(o.LastUpdateTimeSinceEpoch) {
		toSerialize["lastUpdateTimeSinceEpoch"] = o.LastUpdateTimeSinceEpoch
	}
	return toSerialize, nil
}

type NullableRegisteredModel struct {
	value *RegisteredModel
	isSet bool
}

func (v NullableRegisteredModel) Get() *RegisteredModel {
	return v.value
}

func (v *NullableRegisteredModel) Set(val *RegisteredModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisteredModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisteredModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisteredModel(val *RegisteredModel) *NullableRegisteredModel {
	return &NullableRegisteredModel{value: val, isSet: true}
}

func (v NullableRegisteredModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisteredModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
