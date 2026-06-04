# Forbidden

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ForbiddenStatus**](ForbiddenStatus.md) |  | [optional] 
**Data** | Pointer to [**ForbiddenData**](ForbiddenData.md) |  | [optional] 

## Methods

### NewForbidden

`func NewForbidden() *Forbidden`

NewForbidden instantiates a new Forbidden object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForbiddenWithDefaults

`func NewForbiddenWithDefaults() *Forbidden`

NewForbiddenWithDefaults instantiates a new Forbidden object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Forbidden) GetStatus() ForbiddenStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Forbidden) GetStatusOk() (*ForbiddenStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Forbidden) SetStatus(v ForbiddenStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Forbidden) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *Forbidden) GetData() ForbiddenData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Forbidden) GetDataOk() (*ForbiddenData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Forbidden) SetData(v ForbiddenData)`

SetData sets Data field to given value.

### HasData

`func (o *Forbidden) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


