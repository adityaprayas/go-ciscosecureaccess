# VirtualApplianceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginId** | **int64** | The origin ID of the virtual appliance. | 
**Name** | **string** | The name of the virtual appliance. | 
**SiteId** | Pointer to **int64** | The site ID of the virtual appliance. | [optional] 
**IsUpgradable** | **bool** | Specifies whether you can upgrade the virtual appliance (VA) to the latest VA version. | 
**State** | Pointer to [**VirtualApplianceObjectState**](VirtualApplianceObjectState.md) |  | [optional] 
**Health** | **string** | A description of the health of the virtual appliance. | 
**Type** | **string** | The type of the virtual appliance. | 
**Settings** | Pointer to [**VirtualApplianceObjectSettings**](VirtualApplianceObjectSettings.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time (ISO8601 timestamp) when the VA was created. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | The date and time (ISO8601 timestamp) when the VA was modified. | [optional] 
**StateUpdatedAt** | **time.Time** | The date and time (ISO8601 timestamp) when the virtual appliance&#39;s state was updated. | 

## Methods

### NewVirtualApplianceObject

`func NewVirtualApplianceObject(originId int64, name string, isUpgradable bool, health string, type_ string, stateUpdatedAt time.Time, ) *VirtualApplianceObject`

NewVirtualApplianceObject instantiates a new VirtualApplianceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualApplianceObjectWithDefaults

`func NewVirtualApplianceObjectWithDefaults() *VirtualApplianceObject`

NewVirtualApplianceObjectWithDefaults instantiates a new VirtualApplianceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginId

`func (o *VirtualApplianceObject) GetOriginId() int64`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *VirtualApplianceObject) GetOriginIdOk() (*int64, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *VirtualApplianceObject) SetOriginId(v int64)`

SetOriginId sets OriginId field to given value.


### GetName

`func (o *VirtualApplianceObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualApplianceObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualApplianceObject) SetName(v string)`

SetName sets Name field to given value.


### GetSiteId

`func (o *VirtualApplianceObject) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VirtualApplianceObject) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VirtualApplianceObject) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VirtualApplianceObject) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetIsUpgradable

`func (o *VirtualApplianceObject) GetIsUpgradable() bool`

GetIsUpgradable returns the IsUpgradable field if non-nil, zero value otherwise.

### GetIsUpgradableOk

`func (o *VirtualApplianceObject) GetIsUpgradableOk() (*bool, bool)`

GetIsUpgradableOk returns a tuple with the IsUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradable

`func (o *VirtualApplianceObject) SetIsUpgradable(v bool)`

SetIsUpgradable sets IsUpgradable field to given value.


### GetState

`func (o *VirtualApplianceObject) GetState() VirtualApplianceObjectState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualApplianceObject) GetStateOk() (*VirtualApplianceObjectState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VirtualApplianceObject) SetState(v VirtualApplianceObjectState)`

SetState sets State field to given value.

### HasState

`func (o *VirtualApplianceObject) HasState() bool`

HasState returns a boolean if a field has been set.

### GetHealth

`func (o *VirtualApplianceObject) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VirtualApplianceObject) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VirtualApplianceObject) SetHealth(v string)`

SetHealth sets Health field to given value.


### GetType

`func (o *VirtualApplianceObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualApplianceObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualApplianceObject) SetType(v string)`

SetType sets Type field to given value.


### GetSettings

`func (o *VirtualApplianceObject) GetSettings() VirtualApplianceObjectSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *VirtualApplianceObject) GetSettingsOk() (*VirtualApplianceObjectSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *VirtualApplianceObject) SetSettings(v VirtualApplianceObjectSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *VirtualApplianceObject) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VirtualApplianceObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VirtualApplianceObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VirtualApplianceObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VirtualApplianceObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *VirtualApplianceObject) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VirtualApplianceObject) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VirtualApplianceObject) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *VirtualApplianceObject) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetStateUpdatedAt

`func (o *VirtualApplianceObject) GetStateUpdatedAt() time.Time`

GetStateUpdatedAt returns the StateUpdatedAt field if non-nil, zero value otherwise.

### GetStateUpdatedAtOk

`func (o *VirtualApplianceObject) GetStateUpdatedAtOk() (*time.Time, bool)`

GetStateUpdatedAtOk returns a tuple with the StateUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateUpdatedAt

`func (o *VirtualApplianceObject) SetStateUpdatedAt(v time.Time)`

SetStateUpdatedAt sets StateUpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


