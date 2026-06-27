# AccessRuleContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectEntityUrl** | Pointer to **string** | The URL for the entity in Secure Access. | [optional] 
**RedirectEntityUrlLabel** | Pointer to **string** | The label for the entity&#39;s redirect URL. | [optional] 
**ChangeType** | Pointer to **string** | The type of change made on the access rule. | [optional] 
**ChangesMade** | Pointer to **string** | The description of the changes made by the system on the access rule. | [optional] 

## Methods

### NewAccessRuleContext

`func NewAccessRuleContext() *AccessRuleContext`

NewAccessRuleContext instantiates a new AccessRuleContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRuleContextWithDefaults

`func NewAccessRuleContextWithDefaults() *AccessRuleContext`

NewAccessRuleContextWithDefaults instantiates a new AccessRuleContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectEntityUrl

`func (o *AccessRuleContext) GetRedirectEntityUrl() string`

GetRedirectEntityUrl returns the RedirectEntityUrl field if non-nil, zero value otherwise.

### GetRedirectEntityUrlOk

`func (o *AccessRuleContext) GetRedirectEntityUrlOk() (*string, bool)`

GetRedirectEntityUrlOk returns a tuple with the RedirectEntityUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectEntityUrl

`func (o *AccessRuleContext) SetRedirectEntityUrl(v string)`

SetRedirectEntityUrl sets RedirectEntityUrl field to given value.

### HasRedirectEntityUrl

`func (o *AccessRuleContext) HasRedirectEntityUrl() bool`

HasRedirectEntityUrl returns a boolean if a field has been set.

### GetRedirectEntityUrlLabel

`func (o *AccessRuleContext) GetRedirectEntityUrlLabel() string`

GetRedirectEntityUrlLabel returns the RedirectEntityUrlLabel field if non-nil, zero value otherwise.

### GetRedirectEntityUrlLabelOk

`func (o *AccessRuleContext) GetRedirectEntityUrlLabelOk() (*string, bool)`

GetRedirectEntityUrlLabelOk returns a tuple with the RedirectEntityUrlLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectEntityUrlLabel

`func (o *AccessRuleContext) SetRedirectEntityUrlLabel(v string)`

SetRedirectEntityUrlLabel sets RedirectEntityUrlLabel field to given value.

### HasRedirectEntityUrlLabel

`func (o *AccessRuleContext) HasRedirectEntityUrlLabel() bool`

HasRedirectEntityUrlLabel returns a boolean if a field has been set.

### GetChangeType

`func (o *AccessRuleContext) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *AccessRuleContext) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *AccessRuleContext) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *AccessRuleContext) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetChangesMade

`func (o *AccessRuleContext) GetChangesMade() string`

GetChangesMade returns the ChangesMade field if non-nil, zero value otherwise.

### GetChangesMadeOk

`func (o *AccessRuleContext) GetChangesMadeOk() (*string, bool)`

GetChangesMadeOk returns a tuple with the ChangesMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesMade

`func (o *AccessRuleContext) SetChangesMade(v string)`

SetChangesMade sets ChangesMade field to given value.

### HasChangesMade

`func (o *AccessRuleContext) HasChangesMade() bool`

HasChangesMade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


