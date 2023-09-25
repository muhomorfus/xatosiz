# Trace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**GroupUuid** | **string** |  | 
**ParentUuid** | Pointer to **string** |  | [optional] 
**Start** | **string** |  | 
**End** | **string** |  | 
**Title** | **string** |  | 
**Component** | **string** |  | 
**Finished** | **bool** |  | 
**Children** | [**[]Trace**](Trace.md) |  | 
**Events** | [**[]Event**](Event.md) |  | 

## Methods

### NewTrace

`func NewTrace(uuid string, groupUuid string, start string, end string, title string, component string, finished bool, children []Trace, events []Event, ) *Trace`

NewTrace instantiates a new Trace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceWithDefaults

`func NewTraceWithDefaults() *Trace`

NewTraceWithDefaults instantiates a new Trace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Trace) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Trace) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Trace) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetGroupUuid

`func (o *Trace) GetGroupUuid() string`

GetGroupUuid returns the GroupUuid field if non-nil, zero value otherwise.

### GetGroupUuidOk

`func (o *Trace) GetGroupUuidOk() (*string, bool)`

GetGroupUuidOk returns a tuple with the GroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUuid

`func (o *Trace) SetGroupUuid(v string)`

SetGroupUuid sets GroupUuid field to given value.


### GetParentUuid

`func (o *Trace) GetParentUuid() string`

GetParentUuid returns the ParentUuid field if non-nil, zero value otherwise.

### GetParentUuidOk

`func (o *Trace) GetParentUuidOk() (*string, bool)`

GetParentUuidOk returns a tuple with the ParentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUuid

`func (o *Trace) SetParentUuid(v string)`

SetParentUuid sets ParentUuid field to given value.

### HasParentUuid

`func (o *Trace) HasParentUuid() bool`

HasParentUuid returns a boolean if a field has been set.

### GetStart

`func (o *Trace) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Trace) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Trace) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *Trace) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Trace) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Trace) SetEnd(v string)`

SetEnd sets End field to given value.


### GetTitle

`func (o *Trace) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Trace) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Trace) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetComponent

`func (o *Trace) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *Trace) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *Trace) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetFinished

`func (o *Trace) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *Trace) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *Trace) SetFinished(v bool)`

SetFinished sets Finished field to given value.


### GetChildren

`func (o *Trace) GetChildren() []Trace`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Trace) GetChildrenOk() (*[]Trace, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Trace) SetChildren(v []Trace)`

SetChildren sets Children field to given value.


### GetEvents

`func (o *Trace) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Trace) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Trace) SetEvents(v []Event)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


