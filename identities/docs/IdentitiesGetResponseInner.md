# IdentitiesGetResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The unique identifier specified in the UUID4 format. The key is used for all identity endpoints or security group tags. | 
**Label** | **string** | The descriptive name assigned to the identity endpoint or security group tag. | 
**Status** | [**Status**](Status.md) |  | 
**AuthName** | **string** | The attribute used to authenticate the identity endpoint. | 
**CreatedAt** | **int64** | The date and time in nanoseconds since the Unix Epoch when the system created the identity. | 
**ModifiedAt** | **int64** | The date and time in nanoseconds since the Unix Epoch when the system updated the identity. | 
**TagId** | **int64** | The unique identifier of the security group tag. | 

## Methods

### NewIdentitiesGetResponseInner

`func NewIdentitiesGetResponseInner(key string, label string, status Status, authName string, createdAt int64, modifiedAt int64, tagId int64, ) *IdentitiesGetResponseInner`

NewIdentitiesGetResponseInner instantiates a new IdentitiesGetResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitiesGetResponseInnerWithDefaults

`func NewIdentitiesGetResponseInnerWithDefaults() *IdentitiesGetResponseInner`

NewIdentitiesGetResponseInnerWithDefaults instantiates a new IdentitiesGetResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *IdentitiesGetResponseInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IdentitiesGetResponseInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IdentitiesGetResponseInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *IdentitiesGetResponseInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IdentitiesGetResponseInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IdentitiesGetResponseInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStatus

`func (o *IdentitiesGetResponseInner) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentitiesGetResponseInner) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentitiesGetResponseInner) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetAuthName

`func (o *IdentitiesGetResponseInner) GetAuthName() string`

GetAuthName returns the AuthName field if non-nil, zero value otherwise.

### GetAuthNameOk

`func (o *IdentitiesGetResponseInner) GetAuthNameOk() (*string, bool)`

GetAuthNameOk returns a tuple with the AuthName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthName

`func (o *IdentitiesGetResponseInner) SetAuthName(v string)`

SetAuthName sets AuthName field to given value.


### GetCreatedAt

`func (o *IdentitiesGetResponseInner) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentitiesGetResponseInner) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentitiesGetResponseInner) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *IdentitiesGetResponseInner) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *IdentitiesGetResponseInner) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *IdentitiesGetResponseInner) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.


### GetTagId

`func (o *IdentitiesGetResponseInner) GetTagId() int64`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *IdentitiesGetResponseInner) GetTagIdOk() (*int64, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *IdentitiesGetResponseInner) SetTagId(v int64)`

SetTagId sets TagId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


