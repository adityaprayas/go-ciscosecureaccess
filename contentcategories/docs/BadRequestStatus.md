# BadRequestStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** | HTTP status error code | [optional] 
**Text** | Pointer to **string** | HTTP status error text | [optional] 

## Methods

### NewBadRequestStatus

`func NewBadRequestStatus() *BadRequestStatus`

NewBadRequestStatus instantiates a new BadRequestStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestStatusWithDefaults

`func NewBadRequestStatusWithDefaults() *BadRequestStatus`

NewBadRequestStatusWithDefaults instantiates a new BadRequestStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BadRequestStatus) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BadRequestStatus) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BadRequestStatus) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *BadRequestStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetText

`func (o *BadRequestStatus) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *BadRequestStatus) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *BadRequestStatus) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *BadRequestStatus) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


