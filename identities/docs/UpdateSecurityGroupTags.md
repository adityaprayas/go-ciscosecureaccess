# UpdateSecurityGroupTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The unique identifier specified in the UUID4 format. The key is used for all identity endpoints or security group tags. | 
**Label** | **string** | The descriptive name assigned to the identity endpoint or security group tag. | 
**Status** | [**Status**](Status.md) |  | 
**TagId** | **int64** | The unique identifier of the security group tag. | 

## Methods

### NewUpdateSecurityGroupTags

`func NewUpdateSecurityGroupTags(key string, label string, status Status, tagId int64, ) *UpdateSecurityGroupTags`

NewUpdateSecurityGroupTags instantiates a new UpdateSecurityGroupTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecurityGroupTagsWithDefaults

`func NewUpdateSecurityGroupTagsWithDefaults() *UpdateSecurityGroupTags`

NewUpdateSecurityGroupTagsWithDefaults instantiates a new UpdateSecurityGroupTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UpdateSecurityGroupTags) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateSecurityGroupTags) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateSecurityGroupTags) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *UpdateSecurityGroupTags) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateSecurityGroupTags) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateSecurityGroupTags) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStatus

`func (o *UpdateSecurityGroupTags) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateSecurityGroupTags) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateSecurityGroupTags) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetTagId

`func (o *UpdateSecurityGroupTags) GetTagId() int64`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *UpdateSecurityGroupTags) GetTagIdOk() (*int64, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *UpdateSecurityGroupTags) SetTagId(v int64)`

SetTagId sets TagId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


