# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**TraceUuid** | **string** |  | 
**Message** | **string** |  | 
**Priority** | **string** |  | 
**Fixed** | **bool** |  | 
**Time** | **string** |  | 
**Payload** | **map[string]string** |  | 

## Methods

### NewEvent

`func NewEvent(uuid string, traceUuid string, message string, priority string, fixed bool, time string, payload map[string]string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Event) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Event) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Event) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetTraceUuid

`func (o *Event) GetTraceUuid() string`

GetTraceUuid returns the TraceUuid field if non-nil, zero value otherwise.

### GetTraceUuidOk

`func (o *Event) GetTraceUuidOk() (*string, bool)`

GetTraceUuidOk returns a tuple with the TraceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceUuid

`func (o *Event) SetTraceUuid(v string)`

SetTraceUuid sets TraceUuid field to given value.


### GetMessage

`func (o *Event) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Event) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Event) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPriority

`func (o *Event) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Event) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Event) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetFixed

`func (o *Event) GetFixed() bool`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *Event) GetFixedOk() (*bool, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *Event) SetFixed(v bool)`

SetFixed sets Fixed field to given value.


### GetTime

`func (o *Event) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Event) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Event) SetTime(v string)`

SetTime sets Time field to given value.


### GetPayload

`func (o *Event) GetPayload() map[string]string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Event) GetPayloadOk() (*map[string]string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Event) SetPayload(v map[string]string)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


