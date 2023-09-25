# StartTraceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**GroupUuid** | Pointer to **string** |  | [optional] 
**ParentUuid** | Pointer to **string** |  | [optional] 
**Start** | **string** |  | 
**End** | **string** |  | 
**Title** | **string** |  | 
**Component** | **string** |  | 
**Finished** | **bool** |  | 

## Methods

### NewStartTraceResponse

`func NewStartTraceResponse(uuid string, start string, end string, title string, component string, finished bool, ) *StartTraceResponse`

NewStartTraceResponse instantiates a new StartTraceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartTraceResponseWithDefaults

`func NewStartTraceResponseWithDefaults() *StartTraceResponse`

NewStartTraceResponseWithDefaults instantiates a new StartTraceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *StartTraceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StartTraceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StartTraceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetGroupUuid

`func (o *StartTraceResponse) GetGroupUuid() string`

GetGroupUuid returns the GroupUuid field if non-nil, zero value otherwise.

### GetGroupUuidOk

`func (o *StartTraceResponse) GetGroupUuidOk() (*string, bool)`

GetGroupUuidOk returns a tuple with the GroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUuid

`func (o *StartTraceResponse) SetGroupUuid(v string)`

SetGroupUuid sets GroupUuid field to given value.

### HasGroupUuid

`func (o *StartTraceResponse) HasGroupUuid() bool`

HasGroupUuid returns a boolean if a field has been set.

### GetParentUuid

`func (o *StartTraceResponse) GetParentUuid() string`

GetParentUuid returns the ParentUuid field if non-nil, zero value otherwise.

### GetParentUuidOk

`func (o *StartTraceResponse) GetParentUuidOk() (*string, bool)`

GetParentUuidOk returns a tuple with the ParentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUuid

`func (o *StartTraceResponse) SetParentUuid(v string)`

SetParentUuid sets ParentUuid field to given value.

### HasParentUuid

`func (o *StartTraceResponse) HasParentUuid() bool`

HasParentUuid returns a boolean if a field has been set.

### GetStart

`func (o *StartTraceResponse) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *StartTraceResponse) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *StartTraceResponse) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *StartTraceResponse) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *StartTraceResponse) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *StartTraceResponse) SetEnd(v string)`

SetEnd sets End field to given value.


### GetTitle

`func (o *StartTraceResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StartTraceResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StartTraceResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetComponent

`func (o *StartTraceResponse) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *StartTraceResponse) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *StartTraceResponse) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetFinished

`func (o *StartTraceResponse) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *StartTraceResponse) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *StartTraceResponse) SetFinished(v bool)`

SetFinished sets Finished field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


