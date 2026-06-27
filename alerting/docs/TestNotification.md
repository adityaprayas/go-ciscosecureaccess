# TestNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the alert rule. | 
**Type** | **int64** | The identifier of the rule type.  **Category: Connectivity**  | Rule Type ID | Rule Type Name | | --- | --- | | 1 | Network tunnel group disconnected | | 2 | No traffic through network tunnel group |  **Category: API Anomalies**  | Rule Type ID | Rule Type Name | | --- | --- | | 3 | Too many requests to get an access token | | 4 | High percentage of API errors | | 5 | Large number of rate limit errors | | 6 | More than one client using an API credential | | 7 | Authentication failures | | 8 | Programmatic creation of more than one API credential |  **Category: Data Usage**  | Rule Type ID | Rule Type Name | | --- | --- | | 9 | Data usage approaching limit |  **Category: Access Rule Changes**  | Rule Type ID | Rule Type Name | | --- | --- | | 10 | Rule creation, deletion, changing or reordering |  **Category: Behavior Analytics**  | Rule Type ID | Rule Type Name | | --- | --- | | 11 | Impossible travel detected | | 12 | Bulk upload operations | | 13 | Bulk download operations | | 14 | Bulk delete operations | | 15 | DLP violation spike | | 16 | Upload from high-risk countries | | 17 | Upload to high-risk countries | | 
**Severity** | [**SeverityAlert**](SeverityAlert.md) |  | 
**Conditions** | Pointer to [**ConditionsAlertRule**](ConditionsAlertRule.md) |  | [optional] 
**NotificationInfo** | [**[]NotificationInfoAlertRule**](NotificationInfoAlertRule.md) | The list of notification details. | 
**EventTime** | Pointer to **time.Time** | The date and time (ISO 8601 timestamp) when the system recorded the event. | [optional] 
**RedirectLink** | Pointer to **string** | The URL of the alert page in Secure Access. | [optional] 

## Methods

### NewTestNotification

`func NewTestNotification(name string, type_ int64, severity SeverityAlert, notificationInfo []NotificationInfoAlertRule, ) *TestNotification`

NewTestNotification instantiates a new TestNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestNotificationWithDefaults

`func NewTestNotificationWithDefaults() *TestNotification`

NewTestNotificationWithDefaults instantiates a new TestNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TestNotification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestNotification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestNotification) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestNotification) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestNotification) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestNotification) SetType(v int64)`

SetType sets Type field to given value.


### GetSeverity

`func (o *TestNotification) GetSeverity() SeverityAlert`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *TestNotification) GetSeverityOk() (*SeverityAlert, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *TestNotification) SetSeverity(v SeverityAlert)`

SetSeverity sets Severity field to given value.


### GetConditions

`func (o *TestNotification) GetConditions() ConditionsAlertRule`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TestNotification) GetConditionsOk() (*ConditionsAlertRule, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TestNotification) SetConditions(v ConditionsAlertRule)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *TestNotification) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetNotificationInfo

`func (o *TestNotification) GetNotificationInfo() []NotificationInfoAlertRule`

GetNotificationInfo returns the NotificationInfo field if non-nil, zero value otherwise.

### GetNotificationInfoOk

`func (o *TestNotification) GetNotificationInfoOk() (*[]NotificationInfoAlertRule, bool)`

GetNotificationInfoOk returns a tuple with the NotificationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationInfo

`func (o *TestNotification) SetNotificationInfo(v []NotificationInfoAlertRule)`

SetNotificationInfo sets NotificationInfo field to given value.


### GetEventTime

`func (o *TestNotification) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *TestNotification) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *TestNotification) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *TestNotification) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetRedirectLink

`func (o *TestNotification) GetRedirectLink() string`

GetRedirectLink returns the RedirectLink field if non-nil, zero value otherwise.

### GetRedirectLinkOk

`func (o *TestNotification) GetRedirectLinkOk() (*string, bool)`

GetRedirectLinkOk returns a tuple with the RedirectLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectLink

`func (o *TestNotification) SetRedirectLink(v string)`

SetRedirectLink sets RedirectLink field to given value.

### HasRedirectLink

`func (o *TestNotification) HasRedirectLink() bool`

HasRedirectLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


