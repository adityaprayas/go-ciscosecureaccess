# \IdentitiesAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdentities**](IdentitiesAPI.md#GetIdentities) | **Get** /identities/registrations/{type} | List Identities
[**UpdateIdentities**](IdentitiesAPI.md#UpdateIdentities) | **Put** /identities/registrations/{type} | Update Identities



## GetIdentities

> GetIdentities200Response GetIdentities(ctx, type_).Label(label).Offset(offset).Limit(limit).Execute()

List Identities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/identities"
)

func main() {
	type_ := "type__example" // string | Use `device` to query the collection for the identity endpoints in the organization. Use `securityGroupTag` to query the collection for the security group tags in the organization.
	label := "label_example" // string | The descriptive name of the identity endpoint or security group tag (SGT). (optional)
	offset := int64(1) // int64 | The place to start reading in the collection. The default offset is `0`. (optional) (default to 0)
	limit := int64(25) // int64 | The number of items to return in the page. The default limit is `100`. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.GetIdentities(context.Background(), type_).Label(label).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.GetIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentities`: GetIdentities200Response
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.GetIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Use &#x60;device&#x60; to query the collection for the identity endpoints in the organization. Use &#x60;securityGroupTag&#x60; to query the collection for the security group tags in the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | **string** | The descriptive name of the identity endpoint or security group tag (SGT). | 
 **offset** | **int64** | The place to start reading in the collection. The default offset is &#x60;0&#x60;. | [default to 0]
 **limit** | **int64** | The number of items to return in the page. The default limit is &#x60;100&#x60;. | [default to 100]

### Return type

[**GetIdentities200Response**](GetIdentities200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentities

> UpdateIdentities200Response UpdateIdentities(ctx, type_).UpdateIdentitiesRequestInner(updateIdentitiesRequestInner).Execute()

Update Identities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/identities"
)

func main() {
	type_ := "type__example" // string | Use `device` to query the collection for the identity endpoints in the organization. Use `securityGroupTag` to query the collection for the security group tags in the organization.
	updateIdentitiesRequestInner := []openapiclient.UpdateIdentitiesRequestInner{openapiclient.updateIdentities_request_inner{UpdateIdentityDevices: openapiclient.NewUpdateIdentityDevices("123e4567-e89b-12d3-a456-426614174001", "device one", openapiclient.status("active"), "auth_device_beta")}} // []UpdateIdentitiesRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.UpdateIdentities(context.Background(), type_).UpdateIdentitiesRequestInner(updateIdentitiesRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.UpdateIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdentities`: UpdateIdentities200Response
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.UpdateIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Use &#x60;device&#x60; to query the collection for the identity endpoints in the organization. Use &#x60;securityGroupTag&#x60; to query the collection for the security group tags in the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIdentitiesRequestInner** | [**[]UpdateIdentitiesRequestInner**](UpdateIdentitiesRequestInner.md) |  | 

### Return type

[**UpdateIdentities200Response**](UpdateIdentities200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

