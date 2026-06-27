# ListAlerts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** | The total number of active alerts. | 
**SeverityCounts** | **map[string]int64** | The number of alerts categorized by the severity. | 
**Alerts** | [**[]AlertWithAdditionalContext**](AlertWithAdditionalContext.md) | The list of the details for the alerts. | 

## Methods

### NewListAlerts200Response

`func NewListAlerts200Response(total int64, severityCounts map[string]int64, alerts []AlertWithAdditionalContext, ) *ListAlerts200Response`

NewListAlerts200Response instantiates a new ListAlerts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAlerts200ResponseWithDefaults

`func NewListAlerts200ResponseWithDefaults() *ListAlerts200Response`

NewListAlerts200ResponseWithDefaults instantiates a new ListAlerts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListAlerts200Response) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListAlerts200Response) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListAlerts200Response) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetSeverityCounts

`func (o *ListAlerts200Response) GetSeverityCounts() map[string]int64`

GetSeverityCounts returns the SeverityCounts field if non-nil, zero value otherwise.

### GetSeverityCountsOk

`func (o *ListAlerts200Response) GetSeverityCountsOk() (*map[string]int64, bool)`

GetSeverityCountsOk returns a tuple with the SeverityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityCounts

`func (o *ListAlerts200Response) SetSeverityCounts(v map[string]int64)`

SetSeverityCounts sets SeverityCounts field to given value.


### GetAlerts

`func (o *ListAlerts200Response) GetAlerts() []AlertWithAdditionalContext`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *ListAlerts200Response) GetAlertsOk() (*[]AlertWithAdditionalContext, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *ListAlerts200Response) SetAlerts(v []AlertWithAdditionalContext)`

SetAlerts sets Alerts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


