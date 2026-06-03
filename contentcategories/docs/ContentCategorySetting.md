# ContentCategorySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the Content Category setting. | [optional] 
**OrganizationId** | Pointer to **int64** | The ID of the organization. | [optional] 
**IsDefault** | Pointer to **bool** | Specifies whether the setting is the default setting. | [optional] 
**Name** | Pointer to **string** | The name of the Content Category setting. | [optional] 
**CategoryBits** | Pointer to **string** | The number of bits in the Content Category. | [optional] 
**Type** | Pointer to **string** | The type of the Content Category setting. | [optional] 
**CreatedAt** | Pointer to **int64** | The date and time (in seconds since the Unix Epoch) when the system created the Content Category setting. | [optional] 
**ModifiedAt** | Pointer to **int64** | The date and time (in seconds since the Unix Epoch) when the system modified the Content Category setting. | [optional] 
**MarkedForDeletion** | Pointer to **bool** | Specifies whether the system marks the setting for deletion. | [optional] 
**BundleTypeId** | Pointer to [**BundleTypeId**](BundleTypeId.md) |  | [optional] 
**IsSwgDefault** | Pointer to **bool** | Specifies whether the Security profile for the secure web gateway (SWG) is the default Security profile. | [optional] 
**WarnCategoryBits** | Pointer to **string** | The number of bits in Content Category for a Warn action. | [optional] 

## Methods

### NewContentCategorySetting

`func NewContentCategorySetting() *ContentCategorySetting`

NewContentCategorySetting instantiates a new ContentCategorySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentCategorySettingWithDefaults

`func NewContentCategorySettingWithDefaults() *ContentCategorySetting`

NewContentCategorySettingWithDefaults instantiates a new ContentCategorySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContentCategorySetting) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentCategorySetting) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentCategorySetting) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ContentCategorySetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ContentCategorySetting) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ContentCategorySetting) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ContentCategorySetting) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ContentCategorySetting) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetIsDefault

`func (o *ContentCategorySetting) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ContentCategorySetting) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ContentCategorySetting) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ContentCategorySetting) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *ContentCategorySetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentCategorySetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentCategorySetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentCategorySetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategoryBits

`func (o *ContentCategorySetting) GetCategoryBits() string`

GetCategoryBits returns the CategoryBits field if non-nil, zero value otherwise.

### GetCategoryBitsOk

`func (o *ContentCategorySetting) GetCategoryBitsOk() (*string, bool)`

GetCategoryBitsOk returns a tuple with the CategoryBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryBits

`func (o *ContentCategorySetting) SetCategoryBits(v string)`

SetCategoryBits sets CategoryBits field to given value.

### HasCategoryBits

`func (o *ContentCategorySetting) HasCategoryBits() bool`

HasCategoryBits returns a boolean if a field has been set.

### GetType

`func (o *ContentCategorySetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentCategorySetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentCategorySetting) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentCategorySetting) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContentCategorySetting) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContentCategorySetting) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContentCategorySetting) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContentCategorySetting) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ContentCategorySetting) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ContentCategorySetting) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ContentCategorySetting) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ContentCategorySetting) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetMarkedForDeletion

`func (o *ContentCategorySetting) GetMarkedForDeletion() bool`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *ContentCategorySetting) GetMarkedForDeletionOk() (*bool, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *ContentCategorySetting) SetMarkedForDeletion(v bool)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.

### HasMarkedForDeletion

`func (o *ContentCategorySetting) HasMarkedForDeletion() bool`

HasMarkedForDeletion returns a boolean if a field has been set.

### GetBundleTypeId

`func (o *ContentCategorySetting) GetBundleTypeId() BundleTypeId`

GetBundleTypeId returns the BundleTypeId field if non-nil, zero value otherwise.

### GetBundleTypeIdOk

`func (o *ContentCategorySetting) GetBundleTypeIdOk() (*BundleTypeId, bool)`

GetBundleTypeIdOk returns a tuple with the BundleTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleTypeId

`func (o *ContentCategorySetting) SetBundleTypeId(v BundleTypeId)`

SetBundleTypeId sets BundleTypeId field to given value.

### HasBundleTypeId

`func (o *ContentCategorySetting) HasBundleTypeId() bool`

HasBundleTypeId returns a boolean if a field has been set.

### GetIsSwgDefault

`func (o *ContentCategorySetting) GetIsSwgDefault() bool`

GetIsSwgDefault returns the IsSwgDefault field if non-nil, zero value otherwise.

### GetIsSwgDefaultOk

`func (o *ContentCategorySetting) GetIsSwgDefaultOk() (*bool, bool)`

GetIsSwgDefaultOk returns a tuple with the IsSwgDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSwgDefault

`func (o *ContentCategorySetting) SetIsSwgDefault(v bool)`

SetIsSwgDefault sets IsSwgDefault field to given value.

### HasIsSwgDefault

`func (o *ContentCategorySetting) HasIsSwgDefault() bool`

HasIsSwgDefault returns a boolean if a field has been set.

### GetWarnCategoryBits

`func (o *ContentCategorySetting) GetWarnCategoryBits() string`

GetWarnCategoryBits returns the WarnCategoryBits field if non-nil, zero value otherwise.

### GetWarnCategoryBitsOk

`func (o *ContentCategorySetting) GetWarnCategoryBitsOk() (*string, bool)`

GetWarnCategoryBitsOk returns a tuple with the WarnCategoryBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnCategoryBits

`func (o *ContentCategorySetting) SetWarnCategoryBits(v string)`

SetWarnCategoryBits sets WarnCategoryBits field to given value.

### HasWarnCategoryBits

`func (o *ContentCategorySetting) HasWarnCategoryBits() bool`

HasWarnCategoryBits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


