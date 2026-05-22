# GetIdentities200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** | The total number of identities provisioned in the organization. | 
**Limit** | **int64** | The number of items returned in the page. The default value of the &#x60;limit&#x60; is &#x60;100&#x60;. | [default to 100]
**Offset** | **int64** | The place in the collection where the system started to read the resources. The default value of the &#x60;offset&#x60; is &#x60;0&#x60;. | [default to 0]
**Data** | [**[]IdentitiesGetResponseInner**](IdentitiesGetResponseInner.md) | The list of identities in the collection returned for either the &#x60;device&#x60; or &#x60;securityGroupTag&#x60; identity type. | 

## Methods

### NewGetIdentities200Response

`func NewGetIdentities200Response(total int64, limit int64, offset int64, data []IdentitiesGetResponseInner, ) *GetIdentities200Response`

NewGetIdentities200Response instantiates a new GetIdentities200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdentities200ResponseWithDefaults

`func NewGetIdentities200ResponseWithDefaults() *GetIdentities200Response`

NewGetIdentities200ResponseWithDefaults instantiates a new GetIdentities200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetIdentities200Response) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetIdentities200Response) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetIdentities200Response) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetLimit

`func (o *GetIdentities200Response) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetIdentities200Response) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetIdentities200Response) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *GetIdentities200Response) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetIdentities200Response) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetIdentities200Response) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetData

`func (o *GetIdentities200Response) GetData() []IdentitiesGetResponseInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetIdentities200Response) GetDataOk() (*[]IdentitiesGetResponseInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetIdentities200Response) SetData(v []IdentitiesGetResponseInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


