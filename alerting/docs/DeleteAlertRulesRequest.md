# DeleteAlertRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleIds** | **[]int64** | The list of rule IDs used by the system to delete the alert rules. | 

## Methods

### NewDeleteAlertRulesRequest

`func NewDeleteAlertRulesRequest(ruleIds []int64, ) *DeleteAlertRulesRequest`

NewDeleteAlertRulesRequest instantiates a new DeleteAlertRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAlertRulesRequestWithDefaults

`func NewDeleteAlertRulesRequestWithDefaults() *DeleteAlertRulesRequest`

NewDeleteAlertRulesRequestWithDefaults instantiates a new DeleteAlertRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleIds

`func (o *DeleteAlertRulesRequest) GetRuleIds() []int64`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *DeleteAlertRulesRequest) GetRuleIdsOk() (*[]int64, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *DeleteAlertRulesRequest) SetRuleIds(v []int64)`

SetRuleIds sets RuleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


