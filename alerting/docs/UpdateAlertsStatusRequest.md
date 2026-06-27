# UpdateAlertsStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The numeric identifier of the alert status. Use &#x60;1&#x60; to enable the alert or &#x60;2&#x60; to disable the alert. | 
**EntityIds** | **[]string** | The list of identifiers used by the system to update the alerts. | 

## Methods

### NewUpdateAlertsStatusRequest

`func NewUpdateAlertsStatusRequest(status int64, entityIds []string, ) *UpdateAlertsStatusRequest`

NewUpdateAlertsStatusRequest instantiates a new UpdateAlertsStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlertsStatusRequestWithDefaults

`func NewUpdateAlertsStatusRequestWithDefaults() *UpdateAlertsStatusRequest`

NewUpdateAlertsStatusRequestWithDefaults instantiates a new UpdateAlertsStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateAlertsStatusRequest) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAlertsStatusRequest) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAlertsStatusRequest) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetEntityIds

`func (o *UpdateAlertsStatusRequest) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *UpdateAlertsStatusRequest) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *UpdateAlertsStatusRequest) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


