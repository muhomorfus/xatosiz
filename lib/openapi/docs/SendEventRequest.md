# SendEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupUuid** | Pointer to **string** |  | [optional] 
**TraceUuid** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**Component** | **string** |  | 
**Priority** | **string** |  | 
**Payload** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSendEventRequest

`func NewSendEventRequest(message string, component string, priority string, ) *SendEventRequest`

NewSendEventRequest instantiates a new SendEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEventRequestWithDefaults

`func NewSendEventRequestWithDefaults() *SendEventRequest`

NewSendEventRequestWithDefaults instantiates a new SendEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupUuid

`func (o *SendEventRequest) GetGroupUuid() string`

GetGroupUuid returns the GroupUuid field if non-nil, zero value otherwise.

### GetGroupUuidOk

`func (o *SendEventRequest) GetGroupUuidOk() (*string, bool)`

GetGroupUuidOk returns a tuple with the GroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUuid

`func (o *SendEventRequest) SetGroupUuid(v string)`

SetGroupUuid sets GroupUuid field to given value.

### HasGroupUuid

`func (o *SendEventRequest) HasGroupUuid() bool`

HasGroupUuid returns a boolean if a field has been set.

### GetTraceUuid

`func (o *SendEventRequest) GetTraceUuid() string`

GetTraceUuid returns the TraceUuid field if non-nil, zero value otherwise.

### GetTraceUuidOk

`func (o *SendEventRequest) GetTraceUuidOk() (*string, bool)`

GetTraceUuidOk returns a tuple with the TraceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceUuid

`func (o *SendEventRequest) SetTraceUuid(v string)`

SetTraceUuid sets TraceUuid field to given value.

### HasTraceUuid

`func (o *SendEventRequest) HasTraceUuid() bool`

HasTraceUuid returns a boolean if a field has been set.

### GetMessage

`func (o *SendEventRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendEventRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendEventRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetComponent

`func (o *SendEventRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *SendEventRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *SendEventRequest) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetPriority

`func (o *SendEventRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SendEventRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SendEventRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetPayload

`func (o *SendEventRequest) GetPayload() map[string]string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SendEventRequest) GetPayloadOk() (*map[string]string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SendEventRequest) SetPayload(v map[string]string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SendEventRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


