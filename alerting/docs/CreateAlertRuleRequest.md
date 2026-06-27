# CreateAlertRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of the alert rule. | 
**Description** | Pointer to **string** | The description of the alert rule. | [optional] 
**Severity** | [**SeverityAlert**](SeverityAlert.md) |  | 
**Status** | [**StatusAlertRule**](StatusAlertRule.md) |  | 
**RuleTypeId** | **int64** | The identifier of the rule type.  **Category: Connectivity**  | Rule Type ID | Rule Type Name | | --- | --- | | 1 | Network tunnel group disconnected | | 2 | No traffic through network tunnel group |  **Category: API Anomalies**  | Rule Type ID | Rule Type Name | | --- | --- | | 3 | Too many requests to get an access token | | 4 | High percentage of API errors | | 5 | Large number of rate limit errors | | 6 | More than one client using an API credential | | 7 | Authentication failures | | 8 | Programmatic creation of more than one API credential |  **Category: Data Usage**  | Rule Type ID | Rule Type Name | | --- | --- | | 9 | Data usage approaching limit |  **Category: Access Rule Changes**  | Rule Type ID | Rule Type Name | | --- | --- | | 10 | Rule creation, deletion, changing or reordering |  **Category: Behavior Analytics**  | Rule Type ID | Rule Type Name | | --- | --- | | 11 | Impossible travel detected | | 12 | Bulk upload operations | | 13 | Bulk download operations | | 14 | Bulk delete operations | | 15 | DLP violation spike | | 16 | Upload from high-risk countries | | 17 | Upload to high-risk countries | | 
**NotificationInfo** | Pointer to [**[]NotificationInfoAlertRule**](NotificationInfoAlertRule.md) | The list of notifications for the alert rule. | [optional] 
**Conditions** | Pointer to [**ConditionsAlertRule**](ConditionsAlertRule.md) |  | [optional] 

## Methods

### NewCreateAlertRuleRequest

`func NewCreateAlertRuleRequest(name string, severity SeverityAlert, status StatusAlertRule, ruleTypeId int64, ) *CreateAlertRuleRequest`

NewCreateAlertRuleRequest instantiates a new CreateAlertRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlertRuleRequestWithDefaults

`func NewCreateAlertRuleRequestWithDefaults() *CreateAlertRuleRequest`

NewCreateAlertRuleRequestWithDefaults instantiates a new CreateAlertRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAlertRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAlertRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAlertRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAlertRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAlertRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAlertRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAlertRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *CreateAlertRuleRequest) GetSeverity() SeverityAlert`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CreateAlertRuleRequest) GetSeverityOk() (*SeverityAlert, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CreateAlertRuleRequest) SetSeverity(v SeverityAlert)`

SetSeverity sets Severity field to given value.


### GetStatus

`func (o *CreateAlertRuleRequest) GetStatus() StatusAlertRule`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateAlertRuleRequest) GetStatusOk() (*StatusAlertRule, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateAlertRuleRequest) SetStatus(v StatusAlertRule)`

SetStatus sets Status field to given value.


### GetRuleTypeId

`func (o *CreateAlertRuleRequest) GetRuleTypeId() int64`

GetRuleTypeId returns the RuleTypeId field if non-nil, zero value otherwise.

### GetRuleTypeIdOk

`func (o *CreateAlertRuleRequest) GetRuleTypeIdOk() (*int64, bool)`

GetRuleTypeIdOk returns a tuple with the RuleTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleTypeId

`func (o *CreateAlertRuleRequest) SetRuleTypeId(v int64)`

SetRuleTypeId sets RuleTypeId field to given value.


### GetNotificationInfo

`func (o *CreateAlertRuleRequest) GetNotificationInfo() []NotificationInfoAlertRule`

GetNotificationInfo returns the NotificationInfo field if non-nil, zero value otherwise.

### GetNotificationInfoOk

`func (o *CreateAlertRuleRequest) GetNotificationInfoOk() (*[]NotificationInfoAlertRule, bool)`

GetNotificationInfoOk returns a tuple with the NotificationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationInfo

`func (o *CreateAlertRuleRequest) SetNotificationInfo(v []NotificationInfoAlertRule)`

SetNotificationInfo sets NotificationInfo field to given value.

### HasNotificationInfo

`func (o *CreateAlertRuleRequest) HasNotificationInfo() bool`

HasNotificationInfo returns a boolean if a field has been set.

### GetConditions

`func (o *CreateAlertRuleRequest) GetConditions() ConditionsAlertRule`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateAlertRuleRequest) GetConditionsOk() (*ConditionsAlertRule, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateAlertRuleRequest) SetConditions(v ConditionsAlertRule)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CreateAlertRuleRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


