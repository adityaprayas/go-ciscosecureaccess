# Alert

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

### NewAlert

`func NewAlert() *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertId

`func (o *Alert) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *Alert) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *Alert) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *Alert) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Alert) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Alert) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Alert) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Alert) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetDescription

`func (o *Alert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Alert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Alert) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Alert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *Alert) GetStatus() StatusAlert`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Alert) GetStatusOk() (*StatusAlert, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Alert) SetStatus(v StatusAlert)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Alert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRuleId

`func (o *Alert) GetRuleId() int64`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *Alert) GetRuleIdOk() (*int64, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *Alert) SetRuleId(v int64)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *Alert) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Alert) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Alert) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Alert) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Alert) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Alert) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Alert) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Alert) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Alert) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Alert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Alert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Alert) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Alert) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleTypeId

`func (o *Alert) GetRuleTypeId() int64`

GetRuleTypeId returns the RuleTypeId field if non-nil, zero value otherwise.

### GetRuleTypeIdOk

`func (o *Alert) GetRuleTypeIdOk() (*int64, bool)`

GetRuleTypeIdOk returns a tuple with the RuleTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleTypeId

`func (o *Alert) SetRuleTypeId(v int64)`

SetRuleTypeId sets RuleTypeId field to given value.

### HasRuleTypeId

`func (o *Alert) HasRuleTypeId() bool`

HasRuleTypeId returns a boolean if a field has been set.

### GetSeverity

`func (o *Alert) GetSeverity() SeverityAlert`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Alert) GetSeverityOk() (*SeverityAlert, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Alert) SetSeverity(v SeverityAlert)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Alert) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetAccessRuleContext

`func (o *Alert) GetAccessRuleContext() AccessRuleContext`

GetAccessRuleContext returns the AccessRuleContext field if non-nil, zero value otherwise.

### GetAccessRuleContextOk

`func (o *Alert) GetAccessRuleContextOk() (*AccessRuleContext, bool)`

GetAccessRuleContextOk returns a tuple with the AccessRuleContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRuleContext

`func (o *Alert) SetAccessRuleContext(v AccessRuleContext)`

SetAccessRuleContext sets AccessRuleContext field to given value.

### HasAccessRuleContext

`func (o *Alert) HasAccessRuleContext() bool`

HasAccessRuleContext returns a boolean if a field has been set.

### GetBehaviorAnalyticsContext

`func (o *Alert) GetBehaviorAnalyticsContext() BehaviorAnalyticsContext`

GetBehaviorAnalyticsContext returns the BehaviorAnalyticsContext field if non-nil, zero value otherwise.

### GetBehaviorAnalyticsContextOk

`func (o *Alert) GetBehaviorAnalyticsContextOk() (*BehaviorAnalyticsContext, bool)`

GetBehaviorAnalyticsContextOk returns a tuple with the BehaviorAnalyticsContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviorAnalyticsContext

`func (o *Alert) SetBehaviorAnalyticsContext(v BehaviorAnalyticsContext)`

SetBehaviorAnalyticsContext sets BehaviorAnalyticsContext field to given value.

### HasBehaviorAnalyticsContext

`func (o *Alert) HasBehaviorAnalyticsContext() bool`

HasBehaviorAnalyticsContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


