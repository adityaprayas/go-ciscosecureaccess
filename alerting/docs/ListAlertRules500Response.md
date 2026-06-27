# ListAlertRules500Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error message explaining the reason for failure. | [optional] 
**Status** | Pointer to **int64** | The HTTP status code returned in the response. | [optional] 

## Methods

### NewListAlertRules500Response

`func NewListAlertRules500Response() *ListAlertRules500Response`

NewListAlertRules500Response instantiates a new ListAlertRules500Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAlertRules500ResponseWithDefaults

`func NewListAlertRules500ResponseWithDefaults() *ListAlertRules500Response`

NewListAlertRules500ResponseWithDefaults instantiates a new ListAlertRules500Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ListAlertRules500Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListAlertRules500Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListAlertRules500Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListAlertRules500Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatus

`func (o *ListAlertRules500Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAlertRules500Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAlertRules500Response) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListAlertRules500Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


