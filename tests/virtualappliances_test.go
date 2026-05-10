//go:build functional

package tests

import (
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/virtualappliances"
	"github.com/stretchr/testify/require"
)

func TestVirtualAppliancesFunctional(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetVirtualAppliancesClient(env.ctx)

	vaList, listHTTP, err := apiClient.VirtualAppliancesAPI.ListVirtualAppliances(env.ctx).Page(1).Limit(10).Execute()
	if err != nil && listHTTP != nil && (listHTTP.StatusCode == 401 || listHTTP.StatusCode == 403) {
		t.Skipf("skipping virtual appliances list: insufficient API permissions (%s)", listHTTP.Status)
	}
	require.NoError(t, err)
	require.NotNil(t, vaList)

	if len(vaList) == 0 {
		t.Skip("no virtual appliances found in the organization; skipping get/update/delete tests")
	}

	vaID := vaList[0].GetOriginId()

	getResp, _, err := apiClient.VirtualAppliancesAPI.GetVirtualAppliance(env.ctx, vaID).Execute()
	require.NoError(t, err)
	require.NotNil(t, getResp)
	require.Equal(t, vaID, getResp.GetOriginId())

	if getResp.HasSiteId() {
		siteID := getResp.GetSiteId()
		updateReq := *virtualappliances.NewUpdateVirtualApplianceRequest(siteID)
		updateResp, updateHTTP, updateErr := apiClient.VirtualAppliancesAPI.UpdateVirtualAppliance(env.ctx, vaID).UpdateVirtualApplianceRequest(updateReq).Execute()
		if updateErr != nil && updateHTTP != nil && (updateHTTP.StatusCode == 401 || updateHTTP.StatusCode == 403) {
			t.Logf("skipping virtual appliance update: insufficient API permissions (%s)", updateHTTP.Status)
		} else {
			require.NoError(t, updateErr)
			if updateResp != nil {
				require.Equal(t, vaID, updateResp.GetOriginId())
			}
		}
	}
}

func TestVirtualAppliancesListPagination(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetVirtualAppliancesClient(env.ctx)

	page1, listHTTP, err := apiClient.VirtualAppliancesAPI.ListVirtualAppliances(env.ctx).Page(1).Limit(1).Execute()
	if err != nil && listHTTP != nil && (listHTTP.StatusCode == 401 || listHTTP.StatusCode == 403) {
		t.Skipf("skipping virtual appliances pagination: insufficient API permissions (%s)", listHTTP.Status)
	}
	require.NoError(t, err)
	require.NotNil(t, page1)
}
