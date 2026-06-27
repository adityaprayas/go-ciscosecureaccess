# UpdateAlertRulesStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Specifies whether the operation succeeded. Returns false if any ID is invalid. | [optional] 
**SuccessfulIds** | Pointer to **[]int64** | The list of IDs for the alert rules that the system updated. | [optional] 
**ErrorIds** | Pointer to **[]int64** | The list of unique identifiers that either do no exist or the system failed to delete or update. | [optional] 

## Methods

### NewUpdateAlertRulesStatus200Response

`func NewUpdateAlertRulesStatus200Response() *UpdateAlertRulesStatus200Response`

NewUpdateAlertRulesStatus200Response instantiates a new UpdateAlertRulesStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlertRulesStatus200ResponseWithDefaults

`func NewUpdateAlertRulesStatus200ResponseWithDefaults() *UpdateAlertRulesStatus200Response`

NewUpdateAlertRulesStatus200ResponseWithDefaults instantiates a new UpdateAlertRulesStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdateAlertRulesStatus200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateAlertRulesStatus200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateAlertRulesStatus200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateAlertRulesStatus200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSuccessfulIds

`func (o *UpdateAlertRulesStatus200Response) GetSuccessfulIds() []int64`

GetSuccessfulIds returns the SuccessfulIds field if non-nil, zero value otherwise.

### GetSuccessfulIdsOk

`func (o *UpdateAlertRulesStatus200Response) GetSuccessfulIdsOk() (*[]int64, bool)`

GetSuccessfulIdsOk returns a tuple with the SuccessfulIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulIds

`func (o *UpdateAlertRulesStatus200Response) SetSuccessfulIds(v []int64)`

SetSuccessfulIds sets SuccessfulIds field to given value.

### HasSuccessfulIds

`func (o *UpdateAlertRulesStatus200Response) HasSuccessfulIds() bool`

HasSuccessfulIds returns a boolean if a field has been set.

### GetErrorIds

`func (o *UpdateAlertRulesStatus200Response) GetErrorIds() []int64`

GetErrorIds returns the ErrorIds field if non-nil, zero value otherwise.

### GetErrorIdsOk

`func (o *UpdateAlertRulesStatus200Response) GetErrorIdsOk() (*[]int64, bool)`

GetErrorIdsOk returns a tuple with the ErrorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorIds

`func (o *UpdateAlertRulesStatus200Response) SetErrorIds(v []int64)`

SetErrorIds sets ErrorIds field to given value.

### HasErrorIds

`func (o *UpdateAlertRulesStatus200Response) HasErrorIds() bool`

HasErrorIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


