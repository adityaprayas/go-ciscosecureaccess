# TestEmailNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the notification. | [optional] 
**Message** | Pointer to **string** | The message sent by the system. | [optional] 
**Results** | Pointer to [**[]TestEmailNotificationResponseResultsInner**](TestEmailNotificationResponseResultsInner.md) | The list of detailed results for each type of notification. | [optional] 

## Methods

### NewTestEmailNotificationResponse

`func NewTestEmailNotificationResponse() *TestEmailNotificationResponse`

NewTestEmailNotificationResponse instantiates a new TestEmailNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestEmailNotificationResponseWithDefaults

`func NewTestEmailNotificationResponseWithDefaults() *TestEmailNotificationResponse`

NewTestEmailNotificationResponseWithDefaults instantiates a new TestEmailNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TestEmailNotificationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestEmailNotificationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestEmailNotificationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestEmailNotificationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *TestEmailNotificationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestEmailNotificationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestEmailNotificationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestEmailNotificationResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *TestEmailNotificationResponse) GetResults() []TestEmailNotificationResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TestEmailNotificationResponse) GetResultsOk() (*[]TestEmailNotificationResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TestEmailNotificationResponse) SetResults(v []TestEmailNotificationResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *TestEmailNotificationResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


