# ListAlertRules401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Unauthorized | [optional] 
**Status** | Pointer to **int64** | The HTTP status code returned in the response. | [optional] 

## Methods

### NewListAlertRules401Response

`func NewListAlertRules401Response() *ListAlertRules401Response`

NewListAlertRules401Response instantiates a new ListAlertRules401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAlertRules401ResponseWithDefaults

`func NewListAlertRules401ResponseWithDefaults() *ListAlertRules401Response`

NewListAlertRules401ResponseWithDefaults instantiates a new ListAlertRules401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ListAlertRules401Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListAlertRules401Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListAlertRules401Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListAlertRules401Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatus

`func (o *ListAlertRules401Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAlertRules401Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAlertRules401Response) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListAlertRules401Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


