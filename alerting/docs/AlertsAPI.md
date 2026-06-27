# \AlertsAPI

All URIs are relative to *https://api.sse.cisco.com/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlertById**](AlertsAPI.md#GetAlertById) | **Get** /alerting/alerts/{alertId} | Get Alert
[**ListAlerts**](AlertsAPI.md#ListAlerts) | **Get** /alerting/alerts | List Alerts
[**SendAlertNotifications**](AlertsAPI.md#SendAlertNotifications) | **Post** /alerting/alerts/testNotifications/emails | Create Test Alert by Email
[**UpdateAlertsStatus**](AlertsAPI.md#UpdateAlertsStatus) | **Put** /alerting/alerts/status | Update Status of Alerts



## GetAlertById

> Alert GetAlertById(ctx, alertId).Execute()

Get Alert



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
	alertId := "AL-2048-833125-1764567890123-3f9a1c4b2d7e8f01" // string | The unique identifier of the alert.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertById(context.Background(), alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertById`: Alert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The unique identifier of the alert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Alert**](Alert.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlerts

> ListAlerts200Response ListAlerts(ctx).Filters(filters).Limit(limit).Offset(offset).Execute()

List Alerts



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
	filters := *openapiclient.NewFiltersAlertObject() // FiltersAlertObject | Filter the alerts by one or more properties. Specify the properties of the filters query parameter in the JSON format.  Example:  ``` {     \"alert_name\": \"alert number 50\",     \"created_after\": '2025-01-01T00:00:00Z' } ``` (optional)
	limit := int64(5) // int64 | The maximum number of items to return from the collection in the response. (optional) (default to 10)
	offset := int64(0) // int64 | The place to start reading in the collection. The default offset is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListAlerts(context.Background()).Filters(filters).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlerts`: ListAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**FiltersAlertObject**](FiltersAlertObject.md) | Filter the alerts by one or more properties. Specify the properties of the filters query parameter in the JSON format.  Example:  &#x60;&#x60;&#x60; {     \&quot;alert_name\&quot;: \&quot;alert number 50\&quot;,     \&quot;created_after\&quot;: &#39;2025-01-01T00:00:00Z&#39; } &#x60;&#x60;&#x60; | 
 **limit** | **int64** | The maximum number of items to return from the collection in the response. | [default to 10]
 **offset** | **int64** | The place to start reading in the collection. The default offset is 0. | [default to 0]

### Return type

[**ListAlerts200Response**](ListAlerts200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendAlertNotifications

> TestEmailNotificationResponse SendAlertNotifications(ctx).TestNotification(testNotification).Execute()

Create Test Alert by Email



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
	testNotification := *openapiclient.NewTestNotification("Production VPN Tunnel Disconnected", int64(1), openapiclient.severityAlert(1), []openapiclient.NotificationInfoAlertRule{*openapiclient.NewNotificationInfoAlertRule()}) // TestNotification | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.SendAlertNotifications(context.Background()).TestNotification(testNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.SendAlertNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendAlertNotifications`: TestEmailNotificationResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.SendAlertNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendAlertNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testNotification** | [**TestNotification**](TestNotification.md) |  | 

### Return type

[**TestEmailNotificationResponse**](TestEmailNotificationResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertsStatus

> UpdateAlertsStatus200Response UpdateAlertsStatus(ctx).UpdateAlertsStatusRequest(updateAlertsStatusRequest).Execute()

Update Status of Alerts



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
	updateAlertsStatusRequest := *openapiclient.NewUpdateAlertsStatusRequest(int64(1), []string{"1abc"}) // UpdateAlertsStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateAlertsStatus(context.Background()).UpdateAlertsStatusRequest(updateAlertsStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateAlertsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertsStatus`: UpdateAlertsStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateAlertsStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAlertsStatusRequest** | [**UpdateAlertsStatusRequest**](UpdateAlertsStatusRequest.md) |  | 

### Return type

[**UpdateAlertsStatus200Response**](UpdateAlertsStatus200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

