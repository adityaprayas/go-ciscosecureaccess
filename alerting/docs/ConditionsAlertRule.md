# ConditionsAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchType** | Pointer to **string** | Choose from one of the accepted conditions: &#x60;all&#x60; (AND), &#x60;any&#x60; (OR), or &#39;&#39;. | [optional] 
**Rows** | Pointer to [**[]ConditionsAlertRuleRowsInner**](ConditionsAlertRuleRowsInner.md) | The list of conditions that the system uses to evaluate the alert rule. | [optional] 

## Methods

### NewConditionsAlertRule

`func NewConditionsAlertRule() *ConditionsAlertRule`

NewConditionsAlertRule instantiates a new ConditionsAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionsAlertRuleWithDefaults

`func NewConditionsAlertRuleWithDefaults() *ConditionsAlertRule`

NewConditionsAlertRuleWithDefaults instantiates a new ConditionsAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchType

`func (o *ConditionsAlertRule) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *ConditionsAlertRule) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *ConditionsAlertRule) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *ConditionsAlertRule) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetRows

`func (o *ConditionsAlertRule) GetRows() []ConditionsAlertRuleRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ConditionsAlertRule) GetRowsOk() (*[]ConditionsAlertRuleRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ConditionsAlertRule) SetRows(v []ConditionsAlertRuleRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ConditionsAlertRule) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


