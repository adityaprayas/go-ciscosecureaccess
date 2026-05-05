// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/CiscoDevNet/go-ciscosecureaccess/destinationlists"
	"github.com/CiscoDevNet/go-ciscosecureaccess/internaldomains"
	"github.com/CiscoDevNet/go-ciscosecureaccess/internalnetworks"
	"github.com/CiscoDevNet/go-ciscosecureaccess/networks"
	"github.com/CiscoDevNet/go-ciscosecureaccess/ntg"
	"github.com/CiscoDevNet/go-ciscosecureaccess/privateapps"
	"github.com/CiscoDevNet/go-ciscosecureaccess/reports"
	"github.com/CiscoDevNet/go-ciscosecureaccess/resconn"
	"github.com/CiscoDevNet/go-ciscosecureaccess/roaming"
	"github.com/CiscoDevNet/go-ciscosecureaccess/rules"
	"github.com/CiscoDevNet/go-ciscosecureaccess/sites"
	"github.com/CiscoDevNet/go-ciscosecureaccess/swg"
	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const ciscoSseApiEndpoint = "api.sse.cisco.com"

//const ciscoSseTokenUri = fmt.Sprintf("https://%s/auth/v2/token", ciscoSseApiEndpoint)

type SSEClientFactory struct {
	KeyId         string
	KeySecret     string
	ApiEndpoint   string
	SSEHttpClient *http.Client
	mu            sync.Mutex
}

func noRetryOnAuthFailure(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if resp != nil && (resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden) {
		return false, nil
	}

	if err != nil {
		var retrieveErr *oauth2.RetrieveError
		if errors.As(err, &retrieveErr) &&
			(retrieveErr.Response.StatusCode == http.StatusUnauthorized ||
				retrieveErr.Response.StatusCode == http.StatusForbidden) {
			return false, nil
		}
	}

	return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
}

func (c *SSEClientFactory) GetHttpClient(ctx context.Context) *http.Client {

	c.mu.Lock()
	defer c.mu.Unlock()
	if c.ApiEndpoint == "" {
		c.ApiEndpoint = ciscoSseApiEndpoint
	}
	if c.SSEHttpClient == nil {
		// Create a retryable HTTP client for token requests
		tokenRetryClient := retryablehttp.NewClient()
		tokenRetryClient.RetryMax = 10
		tokenRetryClient.CheckRetry = noRetryOnAuthFailure

		sSEAuthconfig := &clientcredentials.Config{
			ClientID:     c.KeyId,
			ClientSecret: c.KeySecret,
			TokenURL:     fmt.Sprintf("https://%s/auth/v2/token", c.ApiEndpoint)}

		// Use context.Background() so the cached HTTP client's token source is not
		// tied to the short-lived context passed by the caller (e.g. Terraform's
		// configure phase context, which is cancelled after Configure returns).
		bgCtx := context.WithValue(context.Background(), oauth2.HTTPClient, tokenRetryClient.StandardClient())

		oauthHttpClient := oauth2.NewClient(bgCtx, sSEAuthconfig.TokenSource(bgCtx))

		// Create another retryable client for API requests
		retryHttpClient := retryablehttp.NewClient()
		retryHttpClient.HTTPClient = oauthHttpClient
		retryHttpClient.RetryMax = 10
		retryHttpClient.CheckRetry = noRetryOnAuthFailure

		c.SSEHttpClient = retryHttpClient.StandardClient()
	}
	return c.SSEHttpClient
}
func (c *SSEClientFactory) GetURLString(suffix string) string {
	var hostname string
	if c.ApiEndpoint == "" {
		hostname = ciscoSseApiEndpoint
	} else {
		hostname = c.ApiEndpoint
	}

	return fmt.Sprintf("https://%s/%s", hostname, suffix)

}

func (c *SSEClientFactory) GetDestinationListsClient(ctx context.Context) *destinationlists.APIClient {
	configuration := destinationlists.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return destinationlists.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetNtgClient(ctx context.Context) *ntg.APIClient {
	configuration := ntg.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return ntg.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetResConnClient(ctx context.Context) *resconn.APIClient {
	configuration := resconn.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return resconn.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetRulesClient(ctx context.Context) *rules.APIClient {
	configuration := rules.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return rules.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetReportsClient(ctx context.Context) *reports.APIClient {
	configuration := reports.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("")
	return reports.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetPrivateAppsClient(ctx context.Context) *privateapps.APIClient {
	configuration := privateapps.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return privateapps.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetNetworksClient(ctx context.Context) *networks.APIClient {
	configuration := networks.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return networks.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetInternalDomainsClient(ctx context.Context) *internaldomains.APIClient {
	configuration := internaldomains.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return internaldomains.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetInternalNetworksClient(ctx context.Context) *internalnetworks.APIClient {
	configuration := internalnetworks.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return internalnetworks.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetRoamingClient(ctx context.Context) *roaming.APIClient {
	configuration := roaming.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return roaming.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetSitesClient(ctx context.Context) *sites.APIClient {
	configuration := sites.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return sites.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetSwgClient(ctx context.Context) *swg.APIClient {
	configuration := swg.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return swg.NewAPIClient(configuration)
}
