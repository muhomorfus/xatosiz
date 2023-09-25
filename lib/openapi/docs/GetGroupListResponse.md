# GetGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | [**[]Group**](Group.md) |  | 
**Fixed** | [**[]Group**](Group.md) |  | 

## Methods

### NewGetGroupListResponse

`func NewGetGroupListResponse(active []Group, fixed []Group, ) *GetGroupListResponse`

NewGetGroupListResponse instantiates a new GetGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupListResponseWithDefaults

`func NewGetGroupListResponseWithDefaults() *GetGroupListResponse`

NewGetGroupListResponseWithDefaults instantiates a new GetGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GetGroupListResponse) GetActive() []Group`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetGroupListResponse) GetActiveOk() (*[]Group, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetGroupListResponse) SetActive(v []Group)`

SetActive sets Active field to given value.


### GetFixed

`func (o *GetGroupListResponse) GetFixed() []Group`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *GetGroupListResponse) GetFixedOk() (*[]Group, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *GetGroupListResponse) SetFixed(v []Group)`

SetFixed sets Fixed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


