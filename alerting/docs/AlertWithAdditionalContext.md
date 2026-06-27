# AlertWithAdditionalContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertId** | Pointer to **string** | The unique identifier of the alert. | [optional] 
**OrganizationId** | Pointer to **int64** | The unique identifier of the organization. | [optional] 
**Description** | Pointer to **string** | The description of the alert. | [optional] 
**Status** | Pointer to [**StatusAlert**](StatusAlert.md) |  | [optional] 
**RuleId** | Pointer to **int64** | The unique identifier of the associated rule. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time and date (ISO 8601 timestamp) when the system created the alert. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | The time and date (ISO 8601 timestamp) when the system last modified the alert. | [optional] 
**Name** | Pointer to **string** | The name of the alert. | [optional] 
**RuleTypeId** | Pointer to **int64** | The identifier of the rule type.  **Category: Connectivity**  | Rule Type ID | Rule Type Name | | --- | --- | | 1 | Network tunnel group disconnected | | 2 | No traffic through network tunnel group |  **Category: API Anomalies**  | Rule Type ID | Rule Type Name | | --- | --- | | 3 | Too many requests to get an access token | | 4 | High percentage of API errors | | 5 | Large number of rate limit errors | | 6 | More than one client using an API credential | | 7 | Authentication failures | | 8 | Programmatic creation of more than one API credential |  **Category: Data Usage**  | Rule Type ID | Rule Type Name | | --- | --- | | 9 | Data usage approaching limit |  **Category: Access Rule Changes**  | Rule Type ID | Rule Type Name | | --- | --- | | 10 | Rule creation, deletion, changing or reordering |  **Category: Behavior Analytics**  | Rule Type ID | Rule Type Name | | --- | --- | | 11 | Impossible travel detected | | 12 | Bulk upload operations | | 13 | Bulk download operations | | 14 | Bulk delete operations | | 15 | DLP violation spike | | 16 | Upload from high-risk countries | | 17 | Upload to high-risk countries | | [optional] 
**Severity** | Pointer to [**SeverityAlert**](SeverityAlert.md) |  | [optional] 
**AccessRuleContext** | Pointer to [**AccessRuleContext**](AccessRuleContext.md) |  | [optional] 
**BehaviorAnalyticsContext** | Pointer to [**BehaviorAnalyticsContext**](BehaviorAnalyticsContext.md) |  | [optional] 

## Methods

### NewAlertWithAdditionalContext

`func NewAlertWithAdditionalContext() *AlertWithAdditionalContext`

NewAlertWithAdditionalContext instantiates a new AlertWithAdditionalContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithAdditionalContextWithDefaults

`func NewAlertWithAdditionalContextWithDefaults() *AlertWithAdditionalContext`

NewAlertWithAdditionalContextWithDefaults instantiates a new AlertWithAdditionalContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertId

`func (o *AlertWithAdditionalContext) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *AlertWithAdditionalContext) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *AlertWithAdditionalContext) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *AlertWithAdditionalContext) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AlertWithAdditionalContext) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AlertWithAdditionalContext) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AlertWithAdditionalContext) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AlertWithAdditionalContext) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetDescription

`func (o *AlertWithAdditionalContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertWithAdditionalContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertWithAdditionalContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertWithAdditionalContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *AlertWithAdditionalContext) GetStatus() StatusAlert`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertWithAdditionalContext) GetStatusOk() (*StatusAlert, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertWithAdditionalContext) SetStatus(v StatusAlert)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertWithAdditionalContext) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRuleId

`func (o *AlertWithAdditionalContext) GetRuleId() int64`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *AlertWithAdditionalContext) GetRuleIdOk() (*int64, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *AlertWithAdditionalContext) SetRuleId(v int64)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *AlertWithAdditionalContext) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AlertWithAdditionalContext) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertWithAdditionalContext) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertWithAdditionalContext) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AlertWithAdditionalContext) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *AlertWithAdditionalContext) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AlertWithAdditionalContext) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AlertWithAdditionalContext) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AlertWithAdditionalContext) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *AlertWithAdditionalContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertWithAdditionalContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertWithAdditionalContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertWithAdditionalContext) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleTypeId

`func (o *AlertWithAdditionalContext) GetRuleTypeId() int64`

GetRuleTypeId returns the RuleTypeId field if non-nil, zero value otherwise.

### GetRuleTypeIdOk

`func (o *AlertWithAdditionalContext) GetRuleTypeIdOk() (*int64, bool)`

GetRuleTypeIdOk returns a tuple with the RuleTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleTypeId

`func (o *AlertWithAdditionalContext) SetRuleTypeId(v int64)`

SetRuleTypeId sets RuleTypeId field to given value.

### HasRuleTypeId

`func (o *AlertWithAdditionalContext) HasRuleTypeId() bool`

HasRuleTypeId returns a boolean if a field has been set.

### GetSeverity

`func (o *AlertWithAdditionalContext) GetSeverity() SeverityAlert`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertWithAdditionalContext) GetSeverityOk() (*SeverityAlert, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertWithAdditionalContext) SetSeverity(v SeverityAlert)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertWithAdditionalContext) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetAccessRuleContext

`func (o *AlertWithAdditionalContext) GetAccessRuleContext() AccessRuleContext`

GetAccessRuleContext returns the AccessRuleContext field if non-nil, zero value otherwise.

### GetAccessRuleContextOk

`func (o *AlertWithAdditionalContext) GetAccessRuleContextOk() (*AccessRuleContext, bool)`

GetAccessRuleContextOk returns a tuple with the AccessRuleContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRuleContext

`func (o *AlertWithAdditionalContext) SetAccessRuleContext(v AccessRuleContext)`

SetAccessRuleContext sets AccessRuleContext field to given value.

### HasAccessRuleContext

`func (o *AlertWithAdditionalContext) HasAccessRuleContext() bool`

HasAccessRuleContext returns a boolean if a field has been set.

### GetBehaviorAnalyticsContext

`func (o *AlertWithAdditionalContext) GetBehaviorAnalyticsContext() BehaviorAnalyticsContext`

GetBehaviorAnalyticsContext returns the BehaviorAnalyticsContext field if non-nil, zero value otherwise.

### GetBehaviorAnalyticsContextOk

`func (o *AlertWithAdditionalContext) GetBehaviorAnalyticsContextOk() (*BehaviorAnalyticsContext, bool)`

GetBehaviorAnalyticsContextOk returns a tuple with the BehaviorAnalyticsContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviorAnalyticsContext

`func (o *AlertWithAdditionalContext) SetBehaviorAnalyticsContext(v BehaviorAnalyticsContext)`

SetBehaviorAnalyticsContext sets BehaviorAnalyticsContext field to given value.

### HasBehaviorAnalyticsContext

`func (o *AlertWithAdditionalContext) HasBehaviorAnalyticsContext() bool`

HasBehaviorAnalyticsContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


