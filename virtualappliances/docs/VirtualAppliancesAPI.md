# \VirtualAppliancesAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVirtualAppliance**](VirtualAppliancesAPI.md#DeleteVirtualAppliance) | **Delete** /virtualappliances/{virtualApplianceId} | Delete Virtual Appliance
[**GetVirtualAppliance**](VirtualAppliancesAPI.md#GetVirtualAppliance) | **Get** /virtualappliances/{virtualApplianceId} | Get Virtual Appliance
[**ListVirtualAppliances**](VirtualAppliancesAPI.md#ListVirtualAppliances) | **Get** /virtualappliances | List Virtual Appliances
[**UpdateVirtualAppliance**](VirtualAppliancesAPI.md#UpdateVirtualAppliance) | **Put** /virtualappliances/{virtualApplianceId} | Update Virtual Appliance



## DeleteVirtualAppliance

> DeleteVirtualAppliance(ctx, virtualApplianceId).Execute()

Delete Virtual Appliance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/virtualappliances"
)

func main() {
	virtualApplianceId := int64(135678) // int64 | The origin ID (originId) of the virtual appliance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualAppliancesAPI.DeleteVirtualAppliance(context.Background(), virtualApplianceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualAppliancesAPI.DeleteVirtualAppliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualApplianceId** | **int64** | The origin ID (originId) of the virtual appliance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualAppliance

> VirtualApplianceObject GetVirtualAppliance(ctx, virtualApplianceId).Execute()

Get Virtual Appliance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/virtualappliances"
)

func main() {
	virtualApplianceId := int64(135678) // int64 | The origin ID (originId) of the virtual appliance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualAppliancesAPI.GetVirtualAppliance(context.Background(), virtualApplianceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualAppliancesAPI.GetVirtualAppliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualAppliance`: VirtualApplianceObject
	fmt.Fprintf(os.Stdout, "Response from `VirtualAppliancesAPI.GetVirtualAppliance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualApplianceId** | **int64** | The origin ID (originId) of the virtual appliance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualApplianceObject**](VirtualApplianceObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualAppliances

> []VirtualApplianceObject ListVirtualAppliances(ctx).Page(page).Limit(limit).Execute()

List Virtual Appliances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/virtualappliances"
)

func main() {
	page := int64(56) // int64 | The number of a page in the collection. (optional) (default to 1)
	limit := int64(56) // int64 | The number of records in the collection to return on the page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualAppliancesAPI.ListVirtualAppliances(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualAppliancesAPI.ListVirtualAppliances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualAppliances`: []VirtualApplianceObject
	fmt.Fprintf(os.Stdout, "Response from `VirtualAppliancesAPI.ListVirtualAppliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualAppliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | The number of a page in the collection. | [default to 1]
 **limit** | **int64** | The number of records in the collection to return on the page. | [default to 100]

### Return type

[**[]VirtualApplianceObject**](VirtualApplianceObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualAppliance

> VirtualApplianceObject UpdateVirtualAppliance(ctx, virtualApplianceId).UpdateVirtualApplianceRequest(updateVirtualApplianceRequest).Execute()

Update Virtual Appliance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/virtualappliances"
)

func main() {
	virtualApplianceId := int64(135678) // int64 | The origin ID (originId) of the virtual appliance.
	updateVirtualApplianceRequest := *openapiclient.NewUpdateVirtualApplianceRequest(int64(123)) // UpdateVirtualApplianceRequest | Update the virtual appliance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualAppliancesAPI.UpdateVirtualAppliance(context.Background(), virtualApplianceId).UpdateVirtualApplianceRequest(updateVirtualApplianceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualAppliancesAPI.UpdateVirtualAppliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVirtualAppliance`: VirtualApplianceObject
	fmt.Fprintf(os.Stdout, "Response from `VirtualAppliancesAPI.UpdateVirtualAppliance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualApplianceId** | **int64** | The origin ID (originId) of the virtual appliance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVirtualApplianceRequest** | [**UpdateVirtualApplianceRequest**](UpdateVirtualApplianceRequest.md) | Update the virtual appliance. | 

### Return type

[**VirtualApplianceObject**](VirtualApplianceObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

