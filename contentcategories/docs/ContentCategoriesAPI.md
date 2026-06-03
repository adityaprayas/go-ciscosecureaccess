# \ContentCategoriesAPI

All URIs are relative to *https://api.sse.cisco.com/policies/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCategorySettings**](ContentCategoriesAPI.md#GetCategorySettings) | **Get** /categorySettings | List Content Categories



## GetCategorySettings

> []ContentCategorySetting GetCategorySettings(ctx).Page(page).Limit(limit).Execute()

List Content Categories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/contentcategories"
)

func main() {
	page := int64(56) // int64 | The number of a page in the collection. (optional) (default to 1)
	limit := int64(20) // int64 | The number of items on a page. The maximum items that are allowed on a page in the response are 100. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentCategoriesAPI.GetCategorySettings(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentCategoriesAPI.GetCategorySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategorySettings`: []ContentCategorySetting
	fmt.Fprintf(os.Stdout, "Response from `ContentCategoriesAPI.GetCategorySettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategorySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | The number of a page in the collection. | [default to 1]
 **limit** | **int64** | The number of items on a page. The maximum items that are allowed on a page in the response are 100. | [default to 10]

### Return type

[**[]ContentCategorySetting**](ContentCategorySetting.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

