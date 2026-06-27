# UpdateAlertRulesStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The numeric identifier of the status for the alert rule. Use &#x60;1&#x60; to enable the alert rule or &#x60;2&#x60; to disable the alert rule. | 
**EntityIds** | **[]int64** | The list of identifiers used by the system to update the alert rules. | 

## Methods

### NewUpdateAlertRulesStatusRequest

`func NewUpdateAlertRulesStatusRequest(status int64, entityIds []int64, ) *UpdateAlertRulesStatusRequest`

NewUpdateAlertRulesStatusRequest instantiates a new UpdateAlertRulesStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlertRulesStatusRequestWithDefaults

`func NewUpdateAlertRulesStatusRequestWithDefaults() *UpdateAlertRulesStatusRequest`

NewUpdateAlertRulesStatusRequestWithDefaults instantiates a new UpdateAlertRulesStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateAlertRulesStatusRequest) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAlertRulesStatusRequest) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAlertRulesStatusRequest) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetEntityIds

`func (o *UpdateAlertRulesStatusRequest) GetEntityIds() []int64`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *UpdateAlertRulesStatusRequest) GetEntityIdsOk() (*[]int64, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *UpdateAlertRulesStatusRequest) SetEntityIds(v []int64)`

SetEntityIds sets EntityIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


