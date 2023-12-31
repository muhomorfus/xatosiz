/*
API

API for traces

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Trace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Trace{}

// Trace struct for Trace
type Trace struct {
	Uuid       string  `json:"uuid"`
	GroupUuid  string  `json:"group_uuid"`
	ParentUuid *string `json:"parent_uuid,omitempty"`
	Start      string  `json:"start"`
	End        string  `json:"end"`
	Title      string  `json:"title"`
	Component  string  `json:"component"`
	Finished   bool    `json:"finished"`
	Children   []Trace `json:"children"`
	Events     []Event `json:"events"`
}

// NewTrace instantiates a new Trace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrace(uuid string, groupUuid string, start string, end string, title string, component string, finished bool, children []Trace, events []Event) *Trace {
	this := Trace{}
	this.Uuid = uuid
	this.GroupUuid = groupUuid
	this.Start = start
	this.End = end
	this.Title = title
	this.Component = component
	this.Finished = finished
	this.Children = children
	this.Events = events
	return &this
}

// NewTraceWithDefaults instantiates a new Trace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceWithDefaults() *Trace {
	this := Trace{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *Trace) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Trace) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Trace) SetUuid(v string) {
	o.Uuid = v
}

// GetGroupUuid returns the GroupUuid field value
func (o *Trace) GetGroupUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupUuid
}

// GetGroupUuidOk returns a tuple with the GroupUuid field value
// and a boolean to check if the value has been set.
func (o *Trace) GetGroupUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupUuid, true
}

// SetGroupUuid sets field value
func (o *Trace) SetGroupUuid(v string) {
	o.GroupUuid = v
}

// GetParentUuid returns the ParentUuid field value if set, zero value otherwise.
func (o *Trace) GetParentUuid() string {
	if o == nil || IsNil(o.ParentUuid) {
		var ret string
		return ret
	}
	return *o.ParentUuid
}

// GetParentUuidOk returns a tuple with the ParentUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trace) GetParentUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ParentUuid) {
		return nil, false
	}
	return o.ParentUuid, true
}

// HasParentUuid returns a boolean if a field has been set.
func (o *Trace) HasParentUuid() bool {
	if o != nil && !IsNil(o.ParentUuid) {
		return true
	}

	return false
}

// SetParentUuid gets a reference to the given string and assigns it to the ParentUuid field.
func (o *Trace) SetParentUuid(v string) {
	o.ParentUuid = &v
}

// GetStart returns the Start field value
func (o *Trace) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *Trace) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *Trace) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *Trace) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *Trace) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *Trace) SetEnd(v string) {
	o.End = v
}

// GetTitle returns the Title field value
func (o *Trace) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Trace) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Trace) SetTitle(v string) {
	o.Title = v
}

// GetComponent returns the Component field value
func (o *Trace) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *Trace) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *Trace) SetComponent(v string) {
	o.Component = v
}

// GetFinished returns the Finished field value
func (o *Trace) GetFinished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value
// and a boolean to check if the value has been set.
func (o *Trace) GetFinishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Finished, true
}

// SetFinished sets field value
func (o *Trace) SetFinished(v bool) {
	o.Finished = v
}

// GetChildren returns the Children field value
func (o *Trace) GetChildren() []Trace {
	if o == nil {
		var ret []Trace
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *Trace) GetChildrenOk() ([]Trace, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *Trace) SetChildren(v []Trace) {
	o.Children = v
}

// GetEvents returns the Events field value
func (o *Trace) GetEvents() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *Trace) GetEventsOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *Trace) SetEvents(v []Event) {
	o.Events = v
}

func (o Trace) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Trace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["group_uuid"] = o.GroupUuid
	if !IsNil(o.ParentUuid) {
		toSerialize["parent_uuid"] = o.ParentUuid
	}
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	toSerialize["title"] = o.Title
	toSerialize["component"] = o.Component
	toSerialize["finished"] = o.Finished
	toSerialize["children"] = o.Children
	toSerialize["events"] = o.Events
	return toSerialize, nil
}

type NullableTrace struct {
	value *Trace
	isSet bool
}

func (v NullableTrace) Get() *Trace {
	return v.value
}

func (v *NullableTrace) Set(val *Trace) {
	v.value = val
	v.isSet = true
}

func (v NullableTrace) IsSet() bool {
	return v.isSet
}

func (v *NullableTrace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrace(val *Trace) *NullableTrace {
	return &NullableTrace{value: val, isSet: true}
}

func (v NullableTrace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
