# ListAlertsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** | The total number of active alerts. | 
**SeverityCounts** | **map[string]int64** | The number of alerts categorized by the severity. | 
**Alerts** | [**[]AlertWithAdditionalContext**](AlertWithAdditionalContext.md) | The list of the details for the alerts. | 

## Methods

### NewListAlertsResponse

`func NewListAlertsResponse(total int64, severityCounts map[string]int64, alerts []AlertWithAdditionalContext, ) *ListAlertsResponse`

NewListAlertsResponse instantiates a new ListAlertsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAlertsResponseWithDefaults

`func NewListAlertsResponseWithDefaults() *ListAlertsResponse`

NewListAlertsResponseWithDefaults instantiates a new ListAlertsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListAlertsResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListAlertsResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListAlertsResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetSeverityCounts

`func (o *ListAlertsResponse) GetSeverityCounts() map[string]int64`

GetSeverityCounts returns the SeverityCounts field if non-nil, zero value otherwise.

### GetSeverityCountsOk

`func (o *ListAlertsResponse) GetSeverityCountsOk() (*map[string]int64, bool)`

GetSeverityCountsOk returns a tuple with the SeverityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityCounts

`func (o *ListAlertsResponse) SetSeverityCounts(v map[string]int64)`

SetSeverityCounts sets SeverityCounts field to given value.


### GetAlerts

`func (o *ListAlertsResponse) GetAlerts() []AlertWithAdditionalContext`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *ListAlertsResponse) GetAlertsOk() (*[]AlertWithAdditionalContext, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *ListAlertsResponse) SetAlerts(v []AlertWithAdditionalContext)`

SetAlerts sets Alerts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


