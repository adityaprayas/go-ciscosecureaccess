# CreateAlertRule409Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error message explaining the reason for failure. | [optional] 
**Status** | Pointer to **int64** | The HTTP status code returned in the response. | [optional] 

## Methods

### NewCreateAlertRule409Response

`func NewCreateAlertRule409Response() *CreateAlertRule409Response`

NewCreateAlertRule409Response instantiates a new CreateAlertRule409Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlertRule409ResponseWithDefaults

`func NewCreateAlertRule409ResponseWithDefaults() *CreateAlertRule409Response`

NewCreateAlertRule409ResponseWithDefaults instantiates a new CreateAlertRule409Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CreateAlertRule409Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateAlertRule409Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateAlertRule409Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CreateAlertRule409Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatus

`func (o *CreateAlertRule409Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateAlertRule409Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateAlertRule409Response) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateAlertRule409Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


