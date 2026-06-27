# DeleteAlertRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Specifies whether the system completed all operations successfully. If any ID is invalid, the system is unable to complete the operation. | [optional] 
**SuccessfulIds** | Pointer to **[]int64** | The list of IDs for the alert rules that the system deleted. | [optional] 
**ErrorIds** | Pointer to **[]int64** | The list of unique identifiers that either do no exist or the system failed to delete or update. | [optional] 

## Methods

### NewDeleteAlertRules200Response

`func NewDeleteAlertRules200Response() *DeleteAlertRules200Response`

NewDeleteAlertRules200Response instantiates a new DeleteAlertRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAlertRules200ResponseWithDefaults

`func NewDeleteAlertRules200ResponseWithDefaults() *DeleteAlertRules200Response`

NewDeleteAlertRules200ResponseWithDefaults instantiates a new DeleteAlertRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DeleteAlertRules200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DeleteAlertRules200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DeleteAlertRules200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DeleteAlertRules200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSuccessfulIds

`func (o *DeleteAlertRules200Response) GetSuccessfulIds() []int64`

GetSuccessfulIds returns the SuccessfulIds field if non-nil, zero value otherwise.

### GetSuccessfulIdsOk

`func (o *DeleteAlertRules200Response) GetSuccessfulIdsOk() (*[]int64, bool)`

GetSuccessfulIdsOk returns a tuple with the SuccessfulIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulIds

`func (o *DeleteAlertRules200Response) SetSuccessfulIds(v []int64)`

SetSuccessfulIds sets SuccessfulIds field to given value.

### HasSuccessfulIds

`func (o *DeleteAlertRules200Response) HasSuccessfulIds() bool`

HasSuccessfulIds returns a boolean if a field has been set.

### GetErrorIds

`func (o *DeleteAlertRules200Response) GetErrorIds() []int64`

GetErrorIds returns the ErrorIds field if non-nil, zero value otherwise.

### GetErrorIdsOk

`func (o *DeleteAlertRules200Response) GetErrorIdsOk() (*[]int64, bool)`

GetErrorIdsOk returns a tuple with the ErrorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorIds

`func (o *DeleteAlertRules200Response) SetErrorIds(v []int64)`

SetErrorIds sets ErrorIds field to given value.

### HasErrorIds

`func (o *DeleteAlertRules200Response) HasErrorIds() bool`

HasErrorIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


