# \AlertRulesAPI

All URIs are relative to *https://api.sse.cisco.com/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertRule**](AlertRulesAPI.md#CreateAlertRule) | **Post** /alerting/rules | Create Alert Rule
[**DeleteAlertRules**](AlertRulesAPI.md#DeleteAlertRules) | **Delete** /alerting/rules | Delete Alert Rules
[**GetAlertRuleById**](AlertRulesAPI.md#GetAlertRuleById) | **Get** /alerting/rules/{ruleId} | Get Alert Rule
[**ListAlertRules**](AlertRulesAPI.md#ListAlertRules) | **Get** /alerting/rules | List Alert Rules
[**UpdateAlertRule**](AlertRulesAPI.md#UpdateAlertRule) | **Put** /alerting/rules/{ruleId} | Update Alert Rule
[**UpdateAlertRulesStatus**](AlertRulesAPI.md#UpdateAlertRulesStatus) | **Put** /alerting/rules/status | Update Status of Alert Rules



## CreateAlertRule

> CreateAlertRule201Response CreateAlertRule(ctx).CreateAlertRuleRequest(createAlertRuleRequest).Execute()

Create Alert Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/alerting"
)

func main() {
	createAlertRuleRequest := *openapiclient.NewCreateAlertRuleRequest("Production VPN Tunnel Alert", openapiclient.severityAlert(1), openapiclient.statusAlertRule(1), int64(1)) // CreateAlertRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.CreateAlertRule(context.Background()).CreateAlertRuleRequest(createAlertRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.CreateAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertRule`: CreateAlertRule201Response
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.CreateAlertRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAlertRuleRequest** | [**CreateAlertRuleRequest**](CreateAlertRuleRequest.md) |  | 

### Return type

[**CreateAlertRule201Response**](CreateAlertRule201Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertRules

> DeleteAlertRules200Response DeleteAlertRules(ctx).DeleteAlertRulesRequest(deleteAlertRulesRequest).Execute()

Delete Alert Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/alerting"
)

func main() {
	deleteAlertRulesRequest := *openapiclient.NewDeleteAlertRulesRequest([]int64{int64(123)}) // DeleteAlertRulesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.DeleteAlertRules(context.Background()).DeleteAlertRulesRequest(deleteAlertRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.DeleteAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertRules`: DeleteAlertRules200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.DeleteAlertRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAlertRulesRequest** | [**DeleteAlertRulesRequest**](DeleteAlertRulesRequest.md) |  | 

### Return type

[**DeleteAlertRules200Response**](DeleteAlertRules200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertRuleById

> AlertRule GetAlertRuleById(ctx, ruleId).Execute()

Get Alert Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/alerting"
)

func main() {
	ruleId := int64(42) // int64 | The unique identifier of the alert rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.GetAlertRuleById(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.GetAlertRuleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertRuleById`: AlertRule
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.GetAlertRuleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | The unique identifier of the alert rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertRule**](AlertRule.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertRules

> []AlertRule ListAlertRules(ctx).Execute()

List Alert Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/alerting"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.ListAlertRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.ListAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertRules`: []AlertRule
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.ListAlertRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertRulesRequest struct via the builder pattern


### Return type

[**[]AlertRule**](AlertRule.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertRule

> UpdateAlertRule200Response UpdateAlertRule(ctx, ruleId).UpdateAlertRuleRequest(updateAlertRuleRequest).Execute()

Update Alert Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/alerting"
)

func main() {
	ruleId := int64(42) // int64 | The unique identifier of the alert rule.
	updateAlertRuleRequest := *openapiclient.NewUpdateAlertRuleRequest() // UpdateAlertRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.UpdateAlertRule(context.Background(), ruleId).UpdateAlertRuleRequest(updateAlertRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.UpdateAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertRule`: UpdateAlertRule200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.UpdateAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | The unique identifier of the alert rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAlertRuleRequest** | [**UpdateAlertRuleRequest**](UpdateAlertRuleRequest.md) |  | 

### Return type

[**UpdateAlertRule200Response**](UpdateAlertRule200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertRulesStatus

> UpdateAlertRulesStatus200Response UpdateAlertRulesStatus(ctx).UpdateAlertRulesStatusRequest(updateAlertRulesStatusRequest).Execute()

Update Status of Alert Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/alerting"
)

func main() {
	updateAlertRulesStatusRequest := *openapiclient.NewUpdateAlertRulesStatusRequest(int64(1), []int64{int64(1)}) // UpdateAlertRulesStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.UpdateAlertRulesStatus(context.Background()).UpdateAlertRulesStatusRequest(updateAlertRulesStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.UpdateAlertRulesStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertRulesStatus`: UpdateAlertRulesStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.UpdateAlertRulesStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertRulesStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAlertRulesStatusRequest** | [**UpdateAlertRulesStatusRequest**](UpdateAlertRulesStatusRequest.md) |  | 

### Return type

[**UpdateAlertRulesStatus200Response**](UpdateAlertRulesStatus200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

