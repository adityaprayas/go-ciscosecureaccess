# Unauthorized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**UnauthorizedStatus**](UnauthorizedStatus.md) |  | [optional] 
**Data** | Pointer to [**UnauthorizedData**](UnauthorizedData.md) |  | [optional] 

## Methods

### NewUnauthorized

`func NewUnauthorized() *Unauthorized`

NewUnauthorized instantiates a new Unauthorized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedWithDefaults

`func NewUnauthorizedWithDefaults() *Unauthorized`

NewUnauthorizedWithDefaults instantiates a new Unauthorized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Unauthorized) GetStatus() UnauthorizedStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Unauthorized) GetStatusOk() (*UnauthorizedStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Unauthorized) SetStatus(v UnauthorizedStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Unauthorized) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *Unauthorized) GetData() UnauthorizedData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Unauthorized) GetDataOk() (*UnauthorizedData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Unauthorized) SetData(v UnauthorizedData)`

SetData sets Data field to given value.

### HasData

`func (o *Unauthorized) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


