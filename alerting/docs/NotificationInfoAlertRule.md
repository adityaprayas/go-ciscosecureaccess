# NotificationInfoAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NotificationTypeAll**](NotificationTypeAll.md) |  | [optional] 
**Recipients** | Pointer to **[]string** | The list of email recipients configured in the alert rule to receive the notification. | [optional] 
**WebhookIds** | Pointer to **[]string** | The list of unique identifiers for the Webhooks configured on the alert rule. | [optional] 

## Methods

### NewNotificationInfoAlertRule

`func NewNotificationInfoAlertRule() *NotificationInfoAlertRule`

NewNotificationInfoAlertRule instantiates a new NotificationInfoAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationInfoAlertRuleWithDefaults

`func NewNotificationInfoAlertRuleWithDefaults() *NotificationInfoAlertRule`

NewNotificationInfoAlertRuleWithDefaults instantiates a new NotificationInfoAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationInfoAlertRule) GetType() NotificationTypeAll`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationInfoAlertRule) GetTypeOk() (*NotificationTypeAll, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationInfoAlertRule) SetType(v NotificationTypeAll)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationInfoAlertRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRecipients

`func (o *NotificationInfoAlertRule) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *NotificationInfoAlertRule) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *NotificationInfoAlertRule) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *NotificationInfoAlertRule) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetWebhookIds

`func (o *NotificationInfoAlertRule) GetWebhookIds() []string`

GetWebhookIds returns the WebhookIds field if non-nil, zero value otherwise.

### GetWebhookIdsOk

`func (o *NotificationInfoAlertRule) GetWebhookIdsOk() (*[]string, bool)`

GetWebhookIdsOk returns a tuple with the WebhookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookIds

`func (o *NotificationInfoAlertRule) SetWebhookIds(v []string)`

SetWebhookIds sets WebhookIds field to given value.

### HasWebhookIds

`func (o *NotificationInfoAlertRule) HasWebhookIds() bool`

HasWebhookIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


