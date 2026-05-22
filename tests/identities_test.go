//go:build functional

package tests

import (
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/identities"
	"github.com/stretchr/testify/require"
)

func TestIdentitiesListDevices(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetIdentitiesClient(env.ctx)

	resp, httpResp, err := apiClient.IdentitiesAPI.GetIdentities(env.ctx, "device").Limit(10).Offset(0).Execute()
	if err != nil && httpResp != nil && (httpResp.StatusCode == 401 || httpResp.StatusCode == 403) {
		t.Skipf("skipping identities list devices: insufficient API permissions (%s)", httpResp.Status)
	}
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.GreaterOrEqual(t, resp.GetTotal(), int64(0))
}

func TestIdentitiesListSecurityGroupTags(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetIdentitiesClient(env.ctx)

	resp, httpResp, err := apiClient.IdentitiesAPI.GetIdentities(env.ctx, "securityGroupTag").Limit(10).Offset(0).Execute()
	if err != nil && httpResp != nil && (httpResp.StatusCode == 401 || httpResp.StatusCode == 403) {
		t.Skipf("skipping identities list SGTs: insufficient API permissions (%s)", httpResp.Status)
	}
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.GreaterOrEqual(t, resp.GetTotal(), int64(0))
}

func TestIdentitiesListWithLabelFilter(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetIdentitiesClient(env.ctx)

	resp, httpResp, err := apiClient.IdentitiesAPI.GetIdentities(env.ctx, "device").Limit(10).Offset(0).Label("nonexistent-label-filter-test").Execute()
	if err != nil && httpResp != nil && (httpResp.StatusCode == 401 || httpResp.StatusCode == 403) {
		t.Skipf("skipping identities label filter: insufficient API permissions (%s)", httpResp.Status)
	}
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestIdentitiesUpdateDevices(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetIdentitiesClient(env.ctx)

	listResp, httpResp, err := apiClient.IdentitiesAPI.GetIdentities(env.ctx, "device").Limit(1).Offset(0).Execute()
	if err != nil && httpResp != nil && (httpResp.StatusCode == 401 || httpResp.StatusCode == 403) {
		t.Skipf("skipping identities update: insufficient API permissions (%s)", httpResp.Status)
	}
	require.NoError(t, err)

	if listResp == nil || len(listResp.GetData()) == 0 {
		t.Skip("no identity devices found in the organization; skipping update test")
	}

	first := listResp.GetData()[0].IdentityEndpointsGetResponse
	if first == nil {
		t.Skip("first identity is not a device endpoint; skipping update test")
	}

	updateItem := identities.UpdateIdentitiesRequestInner{
		UpdateIdentityDevices: &identities.UpdateIdentityDevices{
			Key:      first.GetKey(),
			Label:    first.GetLabel(),
			Status:   first.GetStatus(),
			AuthName: first.GetAuthName(),
		},
	}

	updateResp, updateHTTP, updateErr := apiClient.IdentitiesAPI.UpdateIdentities(env.ctx, "device").
		UpdateIdentitiesRequestInner([]identities.UpdateIdentitiesRequestInner{updateItem}).Execute()
	if updateErr != nil && updateHTTP != nil && (updateHTTP.StatusCode == 401 || updateHTTP.StatusCode == 403) {
		t.Skipf("skipping identities update write: insufficient API permissions (%s)", updateHTTP.Status)
	}
	require.NoError(t, updateErr)
	require.NotNil(t, updateResp)
	require.True(t, updateResp.GetSuccess())
}
