# TestEmailNotificationResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NotificationTypeEmail**](NotificationTypeEmail.md) |  | [optional] 
**Success** | Pointer to **bool** | Specifies whether the notification was sent by the system successfully. | [optional] 
**Error** | Pointer to **string** | When the system fails to send the test message, the system reports an error message. | [optional] 
**RecipientsCount** | Pointer to **int64** | The number of recipients of the test email notification. | [optional] 

## Methods

### NewTestEmailNotificationResponseResultsInner

`func NewTestEmailNotificationResponseResultsInner() *TestEmailNotificationResponseResultsInner`

NewTestEmailNotificationResponseResultsInner instantiates a new TestEmailNotificationResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestEmailNotificationResponseResultsInnerWithDefaults

`func NewTestEmailNotificationResponseResultsInnerWithDefaults() *TestEmailNotificationResponseResultsInner`

NewTestEmailNotificationResponseResultsInnerWithDefaults instantiates a new TestEmailNotificationResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TestEmailNotificationResponseResultsInner) GetType() NotificationTypeEmail`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestEmailNotificationResponseResultsInner) GetTypeOk() (*NotificationTypeEmail, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestEmailNotificationResponseResultsInner) SetType(v NotificationTypeEmail)`

SetType sets Type field to given value.

### HasType

`func (o *TestEmailNotificationResponseResultsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSuccess

`func (o *TestEmailNotificationResponseResultsInner) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TestEmailNotificationResponseResultsInner) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TestEmailNotificationResponseResultsInner) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TestEmailNotificationResponseResultsInner) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *TestEmailNotificationResponseResultsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TestEmailNotificationResponseResultsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TestEmailNotificationResponseResultsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TestEmailNotificationResponseResultsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRecipientsCount

`func (o *TestEmailNotificationResponseResultsInner) GetRecipientsCount() int64`

GetRecipientsCount returns the RecipientsCount field if non-nil, zero value otherwise.

### GetRecipientsCountOk

`func (o *TestEmailNotificationResponseResultsInner) GetRecipientsCountOk() (*int64, bool)`

GetRecipientsCountOk returns a tuple with the RecipientsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientsCount

`func (o *TestEmailNotificationResponseResultsInner) SetRecipientsCount(v int64)`

SetRecipientsCount sets RecipientsCount field to given value.

### HasRecipientsCount

`func (o *TestEmailNotificationResponseResultsInner) HasRecipientsCount() bool`

HasRecipientsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


