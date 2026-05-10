# UpdateIdentityDevices

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

## Methods

### NewUpdateIdentityDevices

`func NewUpdateIdentityDevices(key string, label string, status Status, authName string, ) *UpdateIdentityDevices`

NewUpdateIdentityDevices instantiates a new UpdateIdentityDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdentityDevicesWithDefaults

`func NewUpdateIdentityDevicesWithDefaults() *UpdateIdentityDevices`

NewUpdateIdentityDevicesWithDefaults instantiates a new UpdateIdentityDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UpdateIdentityDevices) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateIdentityDevices) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateIdentityDevices) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *UpdateIdentityDevices) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateIdentityDevices) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateIdentityDevices) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStatus

`func (o *UpdateIdentityDevices) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateIdentityDevices) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateIdentityDevices) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetAuthName

`func (o *UpdateIdentityDevices) GetAuthName() string`

GetAuthName returns the AuthName field if non-nil, zero value otherwise.

### GetAuthNameOk

`func (o *UpdateIdentityDevices) GetAuthNameOk() (*string, bool)`

GetAuthNameOk returns a tuple with the AuthName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthName

`func (o *UpdateIdentityDevices) SetAuthName(v string)`

SetAuthName sets AuthName field to given value.


### GetGuidHash

`func (o *UpdateIdentityDevices) GetGuidHash() string`

GetGuidHash returns the GuidHash field if non-nil, zero value otherwise.

### GetGuidHashOk

`func (o *UpdateIdentityDevices) GetGuidHashOk() (*string, bool)`

GetGuidHashOk returns a tuple with the GuidHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidHash

`func (o *UpdateIdentityDevices) SetGuidHash(v string)`

SetGuidHash sets GuidHash field to given value.

### HasGuidHash

`func (o *UpdateIdentityDevices) HasGuidHash() bool`

HasGuidHash returns a boolean if a field has been set.

### GetDomainName

`func (o *UpdateIdentityDevices) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *UpdateIdentityDevices) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *UpdateIdentityDevices) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *UpdateIdentityDevices) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetSamAccountName

`func (o *UpdateIdentityDevices) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *UpdateIdentityDevices) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *UpdateIdentityDevices) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.

### HasSamAccountName

`func (o *UpdateIdentityDevices) HasSamAccountName() bool`

HasSamAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


