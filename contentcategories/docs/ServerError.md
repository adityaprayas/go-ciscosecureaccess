# ServerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ServerErrorStatus**](ServerErrorStatus.md) |  | [optional] 
**Data** | Pointer to [**ServerErrorData**](ServerErrorData.md) |  | [optional] 

## Methods

### NewServerError

`func NewServerError() *ServerError`

NewServerError instantiates a new ServerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerErrorWithDefaults

`func NewServerErrorWithDefaults() *ServerError`

NewServerErrorWithDefaults instantiates a new ServerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ServerError) GetStatus() ServerErrorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerError) GetStatusOk() (*ServerErrorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerError) SetStatus(v ServerErrorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *ServerError) GetData() ServerErrorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServerError) GetDataOk() (*ServerErrorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServerError) SetData(v ServerErrorData)`

SetData sets Data field to given value.

### HasData

`func (o *ServerError) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


