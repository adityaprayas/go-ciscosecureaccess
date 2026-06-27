# UpdateAlertRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name of the alert rule. | [optional] 
**Description** | Pointer to **string** | The description of the alert rule. | [optional] 
**Severity** | Pointer to [**SeverityAlert**](SeverityAlert.md) |  | [optional] 
**Status** | Pointer to [**StatusAlertRule**](StatusAlertRule.md) |  | [optional] 
**RuleTypeId** | Pointer to **int64** | The identifier of the rule type.  **Category: Connectivity**  | Rule Type ID | Rule Type Name | | --- | --- | | 1 | Network tunnel group disconnected | | 2 | No traffic through network tunnel group |  **Category: API Anomalies**  | Rule Type ID | Rule Type Name | | --- | --- | | 3 | Too many requests to get an access token | | 4 | High percentage of API errors | | 5 | Large number of rate limit errors | | 6 | More than one client using an API credential | | 7 | Authentication failures | | 8 | Programmatic creation of more than one API credential |  **Category: Data Usage**  | Rule Type ID | Rule Type Name | | --- | --- | | 9 | Data usage approaching limit |  **Category: Access Rule Changes**  | Rule Type ID | Rule Type Name | | --- | --- | | 10 | Rule creation, deletion, changing or reordering |  **Category: Behavior Analytics**  | Rule Type ID | Rule Type Name | | --- | --- | | 11 | Impossible travel detected | | 12 | Bulk upload operations | | 13 | Bulk download operations | | 14 | Bulk delete operations | | 15 | DLP violation spike | | 16 | Upload from high-risk countries | | 17 | Upload to high-risk countries | | [optional] 
**NotificationInfo** | Pointer to [**[]NotificationInfoAlertRule**](NotificationInfoAlertRule.md) | The list of notifications for the alert rule. | [optional] 
**Conditions** | Pointer to [**ConditionsAlertRule**](ConditionsAlertRule.md) |  | [optional] 

## Methods

### NewUpdateAlertRuleRequest

`func NewUpdateAlertRuleRequest() *UpdateAlertRuleRequest`

NewUpdateAlertRuleRequest instantiates a new UpdateAlertRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlertRuleRequestWithDefaults

`func NewUpdateAlertRuleRequestWithDefaults() *UpdateAlertRuleRequest`

NewUpdateAlertRuleRequestWithDefaults instantiates a new UpdateAlertRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAlertRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAlertRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAlertRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAlertRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAlertRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAlertRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAlertRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAlertRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *UpdateAlertRuleRequest) GetSeverity() SeverityAlert`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UpdateAlertRuleRequest) GetSeverityOk() (*SeverityAlert, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UpdateAlertRuleRequest) SetSeverity(v SeverityAlert)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *UpdateAlertRuleRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateAlertRuleRequest) GetStatus() StatusAlertRule`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAlertRuleRequest) GetStatusOk() (*StatusAlertRule, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAlertRuleRequest) SetStatus(v StatusAlertRule)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateAlertRuleRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRuleTypeId

`func (o *UpdateAlertRuleRequest) GetRuleTypeId() int64`

GetRuleTypeId returns the RuleTypeId field if non-nil, zero value otherwise.

### GetRuleTypeIdOk

`func (o *UpdateAlertRuleRequest) GetRuleTypeIdOk() (*int64, bool)`

GetRuleTypeIdOk returns a tuple with the RuleTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleTypeId

`func (o *UpdateAlertRuleRequest) SetRuleTypeId(v int64)`

SetRuleTypeId sets RuleTypeId field to given value.

### HasRuleTypeId

`func (o *UpdateAlertRuleRequest) HasRuleTypeId() bool`

HasRuleTypeId returns a boolean if a field has been set.

### GetNotificationInfo

`func (o *UpdateAlertRuleRequest) GetNotificationInfo() []NotificationInfoAlertRule`

GetNotificationInfo returns the NotificationInfo field if non-nil, zero value otherwise.

### GetNotificationInfoOk

`func (o *UpdateAlertRuleRequest) GetNotificationInfoOk() (*[]NotificationInfoAlertRule, bool)`

GetNotificationInfoOk returns a tuple with the NotificationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationInfo

`func (o *UpdateAlertRuleRequest) SetNotificationInfo(v []NotificationInfoAlertRule)`

SetNotificationInfo sets NotificationInfo field to given value.

### HasNotificationInfo

`func (o *UpdateAlertRuleRequest) HasNotificationInfo() bool`

HasNotificationInfo returns a boolean if a field has been set.

### GetConditions

`func (o *UpdateAlertRuleRequest) GetConditions() ConditionsAlertRule`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UpdateAlertRuleRequest) GetConditionsOk() (*ConditionsAlertRule, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UpdateAlertRuleRequest) SetConditions(v ConditionsAlertRule)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *UpdateAlertRuleRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


