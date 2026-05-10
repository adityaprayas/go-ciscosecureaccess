# UpdateIdentitiesRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The unique identifier specified in the UUID4 format. The key is used for all identity endpoints or security group tags. | 
**Label** | **string** | The descriptive name assigned to the identity endpoint or security group tag. | 
**Status** | [**Status**](Status.md) |  | 
**AuthName** | **string** | The attribute used to authenticate the identity endpoint. | 
**GuidHash** | Pointer to **string** | The globally unique identifier for the Active Directory integration. | [optional] 
**DomainName** | Pointer to **string** | The configured domain name for the identity device. | [optional] 
**SamAccountName** | Pointer to **string** | The configured SAM account name for the identity device. | [optional] 
**TagId** | **int64** | The unique identifier of the security group tag. | 

## Methods

### NewUpdateIdentitiesRequestInner

`func NewUpdateIdentitiesRequestInner(key string, label string, status Status, authName string, tagId int64, ) *UpdateIdentitiesRequestInner`

NewUpdateIdentitiesRequestInner instantiates a new UpdateIdentitiesRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdentitiesRequestInnerWithDefaults

`func NewUpdateIdentitiesRequestInnerWithDefaults() *UpdateIdentitiesRequestInner`

NewUpdateIdentitiesRequestInnerWithDefaults instantiates a new UpdateIdentitiesRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UpdateIdentitiesRequestInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateIdentitiesRequestInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateIdentitiesRequestInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *UpdateIdentitiesRequestInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateIdentitiesRequestInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateIdentitiesRequestInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStatus

`func (o *UpdateIdentitiesRequestInner) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateIdentitiesRequestInner) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateIdentitiesRequestInner) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetAuthName

`func (o *UpdateIdentitiesRequestInner) GetAuthName() string`

GetAuthName returns the AuthName field if non-nil, zero value otherwise.

### GetAuthNameOk

`func (o *UpdateIdentitiesRequestInner) GetAuthNameOk() (*string, bool)`

GetAuthNameOk returns a tuple with the AuthName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthName

`func (o *UpdateIdentitiesRequestInner) SetAuthName(v string)`

SetAuthName sets AuthName field to given value.


### GetGuidHash

`func (o *UpdateIdentitiesRequestInner) GetGuidHash() string`

GetGuidHash returns the GuidHash field if non-nil, zero value otherwise.

### GetGuidHashOk

`func (o *UpdateIdentitiesRequestInner) GetGuidHashOk() (*string, bool)`

GetGuidHashOk returns a tuple with the GuidHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidHash

`func (o *UpdateIdentitiesRequestInner) SetGuidHash(v string)`

SetGuidHash sets GuidHash field to given value.

### HasGuidHash

`func (o *UpdateIdentitiesRequestInner) HasGuidHash() bool`

HasGuidHash returns a boolean if a field has been set.

### GetDomainName

`func (o *UpdateIdentitiesRequestInner) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *UpdateIdentitiesRequestInner) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *UpdateIdentitiesRequestInner) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *UpdateIdentitiesRequestInner) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetSamAccountName

`func (o *UpdateIdentitiesRequestInner) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *UpdateIdentitiesRequestInner) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *UpdateIdentitiesRequestInner) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.

### HasSamAccountName

`func (o *UpdateIdentitiesRequestInner) HasSamAccountName() bool`

HasSamAccountName returns a boolean if a field has been set.

### GetTagId

`func (o *UpdateIdentitiesRequestInner) GetTagId() int64`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *UpdateIdentitiesRequestInner) GetTagIdOk() (*int64, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *UpdateIdentitiesRequestInner) SetTagId(v int64)`

SetTagId sets TagId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


