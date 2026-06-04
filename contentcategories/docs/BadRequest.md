# BadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**BadRequestStatus**](BadRequestStatus.md) |  | [optional] 
**Data** | Pointer to [**BadRequestData**](BadRequestData.md) |  | [optional] 

## Methods

### NewBadRequest

`func NewBadRequest() *BadRequest`

NewBadRequest instantiates a new BadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestWithDefaults

`func NewBadRequestWithDefaults() *BadRequest`

NewBadRequestWithDefaults instantiates a new BadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BadRequest) GetStatus() BadRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BadRequest) GetStatusOk() (*BadRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BadRequest) SetStatus(v BadRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BadRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *BadRequest) GetData() BadRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BadRequest) GetDataOk() (*BadRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BadRequest) SetData(v BadRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *BadRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


