# StartTraceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupUuid** | Pointer to **string** |  | [optional] 
**ParentUuid** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Component** | **string** |  | 

## Methods

### NewStartTraceRequest

`func NewStartTraceRequest(title string, component string, ) *StartTraceRequest`

NewStartTraceRequest instantiates a new StartTraceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartTraceRequestWithDefaults

`func NewStartTraceRequestWithDefaults() *StartTraceRequest`

NewStartTraceRequestWithDefaults instantiates a new StartTraceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupUuid

`func (o *StartTraceRequest) GetGroupUuid() string`

GetGroupUuid returns the GroupUuid field if non-nil, zero value otherwise.

### GetGroupUuidOk

`func (o *StartTraceRequest) GetGroupUuidOk() (*string, bool)`

GetGroupUuidOk returns a tuple with the GroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUuid

`func (o *StartTraceRequest) SetGroupUuid(v string)`

SetGroupUuid sets GroupUuid field to given value.

### HasGroupUuid

`func (o *StartTraceRequest) HasGroupUuid() bool`

HasGroupUuid returns a boolean if a field has been set.

### GetParentUuid

`func (o *StartTraceRequest) GetParentUuid() string`

GetParentUuid returns the ParentUuid field if non-nil, zero value otherwise.

### GetParentUuidOk

`func (o *StartTraceRequest) GetParentUuidOk() (*string, bool)`

GetParentUuidOk returns a tuple with the ParentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUuid

`func (o *StartTraceRequest) SetParentUuid(v string)`

SetParentUuid sets ParentUuid field to given value.

### HasParentUuid

`func (o *StartTraceRequest) HasParentUuid() bool`

HasParentUuid returns a boolean if a field has been set.

### GetTitle

`func (o *StartTraceRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StartTraceRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StartTraceRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetComponent

`func (o *StartTraceRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *StartTraceRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *StartTraceRequest) SetComponent(v string)`

SetComponent sets Component field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


