# FiltersAlertObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**StatusAlert**](StatusAlert.md) |  | [optional] 
**Severity** | Pointer to [**SeverityAlert**](SeverityAlert.md) |  | [optional] 
**CreatedAfter** | Pointer to **time.Time** | Filter for the alerts in the collection that the system created after the timestamp. Provide a date and time (ISO 8601) using the YYYY-MM-DD HH:MM:SS format. **Note:** You cannot use the &#x60;created_after&#x60; query parameter with the &#x60;time_range&#x60; filter. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | The time and date (ISO 8601 timestamp) when the system last modified the alert. | [optional] 
**AlertName** | Pointer to **string** | The name of the alert. | [optional] 
**PatternSearch** | Pointer to **string** | Provide a search pattern to query for by the alert name, alert rule name, or alert rule category name. | [optional] 
**OnlyActiveAlertsCount** | Pointer to **bool** | Specify whether to query for the count of the active alerts only. | [optional] [default to false]
**TimeRange** | Pointer to [**TimeRange**](TimeRange.md) |  | [optional] 
**IncludeContext** | Pointer to **bool** | Include the context field in each alert response. The context contains additional metadata about the alert. | [optional] [default to false]

## Methods

### NewFiltersAlertObject

`func NewFiltersAlertObject() *FiltersAlertObject`

NewFiltersAlertObject instantiates a new FiltersAlertObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersAlertObjectWithDefaults

`func NewFiltersAlertObjectWithDefaults() *FiltersAlertObject`

NewFiltersAlertObjectWithDefaults instantiates a new FiltersAlertObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FiltersAlertObject) GetStatus() StatusAlert`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FiltersAlertObject) GetStatusOk() (*StatusAlert, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FiltersAlertObject) SetStatus(v StatusAlert)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FiltersAlertObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeverity

`func (o *FiltersAlertObject) GetSeverity() SeverityAlert`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FiltersAlertObject) GetSeverityOk() (*SeverityAlert, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FiltersAlertObject) SetSeverity(v SeverityAlert)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *FiltersAlertObject) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreatedAfter

`func (o *FiltersAlertObject) GetCreatedAfter() time.Time`

GetCreatedAfter returns the CreatedAfter field if non-nil, zero value otherwise.

### GetCreatedAfterOk

`func (o *FiltersAlertObject) GetCreatedAfterOk() (*time.Time, bool)`

GetCreatedAfterOk returns a tuple with the CreatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfter

`func (o *FiltersAlertObject) SetCreatedAfter(v time.Time)`

SetCreatedAfter sets CreatedAfter field to given value.

### HasCreatedAfter

`func (o *FiltersAlertObject) HasCreatedAfter() bool`

HasCreatedAfter returns a boolean if a field has been set.

### GetModifiedAt

`func (o *FiltersAlertObject) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *FiltersAlertObject) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *FiltersAlertObject) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *FiltersAlertObject) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetAlertName

`func (o *FiltersAlertObject) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *FiltersAlertObject) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *FiltersAlertObject) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *FiltersAlertObject) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetPatternSearch

`func (o *FiltersAlertObject) GetPatternSearch() string`

GetPatternSearch returns the PatternSearch field if non-nil, zero value otherwise.

### GetPatternSearchOk

`func (o *FiltersAlertObject) GetPatternSearchOk() (*string, bool)`

GetPatternSearchOk returns a tuple with the PatternSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternSearch

`func (o *FiltersAlertObject) SetPatternSearch(v string)`

SetPatternSearch sets PatternSearch field to given value.

### HasPatternSearch

`func (o *FiltersAlertObject) HasPatternSearch() bool`

HasPatternSearch returns a boolean if a field has been set.

### GetOnlyActiveAlertsCount

`func (o *FiltersAlertObject) GetOnlyActiveAlertsCount() bool`

GetOnlyActiveAlertsCount returns the OnlyActiveAlertsCount field if non-nil, zero value otherwise.

### GetOnlyActiveAlertsCountOk

`func (o *FiltersAlertObject) GetOnlyActiveAlertsCountOk() (*bool, bool)`

GetOnlyActiveAlertsCountOk returns a tuple with the OnlyActiveAlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyActiveAlertsCount

`func (o *FiltersAlertObject) SetOnlyActiveAlertsCount(v bool)`

SetOnlyActiveAlertsCount sets OnlyActiveAlertsCount field to given value.

### HasOnlyActiveAlertsCount

`func (o *FiltersAlertObject) HasOnlyActiveAlertsCount() bool`

HasOnlyActiveAlertsCount returns a boolean if a field has been set.

### GetTimeRange

`func (o *FiltersAlertObject) GetTimeRange() TimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *FiltersAlertObject) GetTimeRangeOk() (*TimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *FiltersAlertObject) SetTimeRange(v TimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *FiltersAlertObject) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetIncludeContext

`func (o *FiltersAlertObject) GetIncludeContext() bool`

GetIncludeContext returns the IncludeContext field if non-nil, zero value otherwise.

### GetIncludeContextOk

`func (o *FiltersAlertObject) GetIncludeContextOk() (*bool, bool)`

GetIncludeContextOk returns a tuple with the IncludeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeContext

`func (o *FiltersAlertObject) SetIncludeContext(v bool)`

SetIncludeContext sets IncludeContext field to given value.

### HasIncludeContext

`func (o *FiltersAlertObject) HasIncludeContext() bool`

HasIncludeContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


