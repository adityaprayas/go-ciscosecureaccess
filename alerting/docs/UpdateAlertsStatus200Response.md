# UpdateAlertsStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | If all operations succeeded, the system returns &#x60;true&#x60;. If any ID is invalid, the system returns &#x60;false&#x60;. | [optional] 
**SuccessfulIds** | Pointer to **[]string** | The list of IDs for the alert rules that the system updated. | [optional] 
**ErrorIds** | Pointer to **[]string** | The list of unique identifiers that either do no exist or the system failed to delete or update. | [optional] 

## Methods

### NewUpdateAlertsStatus200Response

`func NewUpdateAlertsStatus200Response() *UpdateAlertsStatus200Response`

NewUpdateAlertsStatus200Response instantiates a new UpdateAlertsStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlertsStatus200ResponseWithDefaults

`func NewUpdateAlertsStatus200ResponseWithDefaults() *UpdateAlertsStatus200Response`

NewUpdateAlertsStatus200ResponseWithDefaults instantiates a new UpdateAlertsStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdateAlertsStatus200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateAlertsStatus200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateAlertsStatus200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateAlertsStatus200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSuccessfulIds

`func (o *UpdateAlertsStatus200Response) GetSuccessfulIds() []string`

GetSuccessfulIds returns the SuccessfulIds field if non-nil, zero value otherwise.

### GetSuccessfulIdsOk

`func (o *UpdateAlertsStatus200Response) GetSuccessfulIdsOk() (*[]string, bool)`

GetSuccessfulIdsOk returns a tuple with the SuccessfulIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulIds

`func (o *UpdateAlertsStatus200Response) SetSuccessfulIds(v []string)`

SetSuccessfulIds sets SuccessfulIds field to given value.

### HasSuccessfulIds

`func (o *UpdateAlertsStatus200Response) HasSuccessfulIds() bool`

HasSuccessfulIds returns a boolean if a field has been set.

### GetErrorIds

`func (o *UpdateAlertsStatus200Response) GetErrorIds() []string`

GetErrorIds returns the ErrorIds field if non-nil, zero value otherwise.

### GetErrorIdsOk

`func (o *UpdateAlertsStatus200Response) GetErrorIdsOk() (*[]string, bool)`

GetErrorIdsOk returns a tuple with the ErrorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorIds

`func (o *UpdateAlertsStatus200Response) SetErrorIds(v []string)`

SetErrorIds sets ErrorIds field to given value.

### HasErrorIds

`func (o *UpdateAlertsStatus200Response) HasErrorIds() bool`

HasErrorIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


