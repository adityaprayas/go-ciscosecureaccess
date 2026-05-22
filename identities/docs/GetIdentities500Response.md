# GetIdentities500Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Indicates if the request was successful. | [optional] 
**Error** | Pointer to **string** | Error string translation if any. | [optional] 

## Methods

### NewGetIdentities500Response

`func NewGetIdentities500Response() *GetIdentities500Response`

NewGetIdentities500Response instantiates a new GetIdentities500Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdentities500ResponseWithDefaults

`func NewGetIdentities500ResponseWithDefaults() *GetIdentities500Response`

NewGetIdentities500ResponseWithDefaults instantiates a new GetIdentities500Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetIdentities500Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetIdentities500Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetIdentities500Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetIdentities500Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *GetIdentities500Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdentities500Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdentities500Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdentities500Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


