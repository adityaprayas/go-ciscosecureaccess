# AlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of the alert rule. | [optional] 
**Name** | Pointer to **string** | The unique name of the alert rule. | [optional] 
**Description** | Pointer to **string** | The description of the alert rule. | [optional] 
**Severity** | Pointer to [**SeverityAlert**](SeverityAlert.md) |  | [optional] 
**Status** | Pointer to [**StatusAlertRule**](StatusAlertRule.md) |  | [optional] 
**RuleTypeId** | Pointer to **int64** | The identifier of the rule type.  **Category: Connectivity**  | Rule Type ID | Rule Type Name | | --- | --- | | 1 | Network tunnel group disconnected | | 2 | No traffic through network tunnel group |  **Category: API Anomalies**  | Rule Type ID | Rule Type Name | | --- | --- | | 3 | Too many requests to get an access token | | 4 | High percentage of API errors | | 5 | Large number of rate limit errors | | 6 | More than one client using an API credential | | 7 | Authentication failures | | 8 | Programmatic creation of more than one API credential |  **Category: Data Usage**  | Rule Type ID | Rule Type Name | | --- | --- | | 9 | Data usage approaching limit |  **Category: Access Rule Changes**  | Rule Type ID | Rule Type Name | | --- | --- | | 10 | Rule creation, deletion, changing or reordering |  **Category: Behavior Analytics**  | Rule Type ID | Rule Type Name | | --- | --- | | 11 | Impossible travel detected | | 12 | Bulk upload operations | | 13 | Bulk download operations | | 14 | Bulk delete operations | | 15 | DLP violation spike | | 16 | Upload from high-risk countries | | 17 | Upload to high-risk countries | | [optional] 
**NotificationInfo** | Pointer to [**[]NotificationInfoAlertRule**](NotificationInfoAlertRule.md) | List of notification configurations for this rule | [optional] 
**OrganizationId** | Pointer to **int64** | The unique identifier of the organization. | [optional] 
**Conditions** | Pointer to [**ConditionsAlertRule**](ConditionsAlertRule.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time and date (ISO 8601 timestamp) when the system created the alert rule. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | The time and date (ISO 8601 timestamp) when the system last modified the alert rule. | [optional] 

## Methods

### NewAlertRule

`func NewAlertRule() *AlertRule`

NewAlertRule instantiates a new AlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleWithDefaults

`func NewAlertRuleWithDefaults() *AlertRule`

NewAlertRuleWithDefaults instantiates a new AlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertRule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AlertRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AlertRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AlertRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *AlertRule) GetSeverity() SeverityAlert`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertRule) GetSeverityOk() (*SeverityAlert, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertRule) SetSeverity(v SeverityAlert)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertRule) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *AlertRule) GetStatus() StatusAlertRule`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertRule) GetStatusOk() (*StatusAlertRule, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertRule) SetStatus(v StatusAlertRule)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRuleTypeId

`func (o *AlertRule) GetRuleTypeId() int64`

GetRuleTypeId returns the RuleTypeId field if non-nil, zero value otherwise.

### GetRuleTypeIdOk

`func (o *AlertRule) GetRuleTypeIdOk() (*int64, bool)`

GetRuleTypeIdOk returns a tuple with the RuleTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleTypeId

`func (o *AlertRule) SetRuleTypeId(v int64)`

SetRuleTypeId sets RuleTypeId field to given value.

### HasRuleTypeId

`func (o *AlertRule) HasRuleTypeId() bool`

HasRuleTypeId returns a boolean if a field has been set.

### GetNotificationInfo

`func (o *AlertRule) GetNotificationInfo() []NotificationInfoAlertRule`

GetNotificationInfo returns the NotificationInfo field if non-nil, zero value otherwise.

### GetNotificationInfoOk

`func (o *AlertRule) GetNotificationInfoOk() (*[]NotificationInfoAlertRule, bool)`

GetNotificationInfoOk returns a tuple with the NotificationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationInfo

`func (o *AlertRule) SetNotificationInfo(v []NotificationInfoAlertRule)`

SetNotificationInfo sets NotificationInfo field to given value.

### HasNotificationInfo

`func (o *AlertRule) HasNotificationInfo() bool`

HasNotificationInfo returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AlertRule) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AlertRule) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AlertRule) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AlertRule) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetConditions

`func (o *AlertRule) GetConditions() ConditionsAlertRule`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AlertRule) GetConditionsOk() (*ConditionsAlertRule, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AlertRule) SetConditions(v ConditionsAlertRule)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AlertRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AlertRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AlertRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *AlertRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AlertRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AlertRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AlertRule) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


