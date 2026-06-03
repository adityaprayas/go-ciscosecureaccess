# NotFound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NotFoundStatus**](NotFoundStatus.md) |  | [optional] 
**Data** | Pointer to [**NotFoundData**](NotFoundData.md) |  | [optional] 

## Methods

### NewNotFound

`func NewNotFound() *NotFound`

NewNotFound instantiates a new NotFound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotFoundWithDefaults

`func NewNotFoundWithDefaults() *NotFound`

NewNotFoundWithDefaults instantiates a new NotFound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NotFound) GetStatus() NotFoundStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotFound) GetStatusOk() (*NotFoundStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotFound) SetStatus(v NotFoundStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotFound) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *NotFound) GetData() NotFoundData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotFound) GetDataOk() (*NotFoundData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotFound) SetData(v NotFoundData)`

SetData sets Data field to given value.

### HasData

`func (o *NotFound) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


