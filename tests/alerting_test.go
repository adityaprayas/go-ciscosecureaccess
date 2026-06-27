//go:build functional

// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

package tests

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/CiscoDevNet/go-ciscosecureaccess/alerting"
	"github.com/stretchr/testify/require"
)

func TestAlertRulesFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetAlertingClient(env.ctx)

	_, _, err := apiClient.AlertRulesAPI.ListAlertRules(env.ctx).Execute()
	require.NoError(t, err)

	notif := alerting.NewNotificationInfoAlertRule()
	notif.SetType(alerting.NOTIFICATIONTYPEALL_EMAIL)
	notif.SetRecipients([]string{"sdk-test@example.com"})

	createReq := alerting.NewCreateAlertRuleRequest(
		uniqueName("go-sdk-alert-rule"),
		alerting.SEVERITYALERT__2,
		alerting.STATUSALERTRULE__1,
		1,
	)
	createReq.SetNotificationInfo([]alerting.NotificationInfoAlertRule{*notif})

	_, createHTTP, err := apiClient.AlertRulesAPI.CreateAlertRule(env.ctx).
		CreateAlertRuleRequest(*createReq).
		Execute()
	require.NoError(t, err)
	require.True(t, isHTTP2xx(createHTTP))

	listResp, _, err := apiClient.AlertRulesAPI.ListAlertRules(env.ctx).Execute()
	require.NoError(t, err)

	var ruleID int64
	for _, r := range listResp {
		if r.GetName() == createReq.GetName() {
			ruleID = r.GetId()
			break
		}
	}
	require.NotZero(t, ruleID, "created alert rule not found in list response")

	t.Cleanup(func() {
		if ruleID == 0 {
			return
		}
		delReq := alerting.NewDeleteAlertRulesRequest([]int64{ruleID})
		_, _, cleanupErr := apiClient.AlertRulesAPI.DeleteAlertRules(env.ctx).
			DeleteAlertRulesRequest(*delReq).
			Execute()
		if cleanupErr != nil {
			t.Logf("cleanup alert rule %d: %v", ruleID, cleanupErr)
		}
	})

	getResp, _, err := apiClient.AlertRulesAPI.GetAlertRuleById(env.ctx, ruleID).Execute()
	require.NoError(t, err)
	require.Equal(t, createReq.GetName(), getResp.GetName())
	require.Equal(t, int64(1), getResp.GetRuleTypeId())

	updatedName := uniqueName("go-sdk-alert-rule-updated")
	updateReq := alerting.NewUpdateAlertRuleRequest()
	updateReq.SetName(updatedName)
	updateReq.SetSeverity(alerting.SEVERITYALERT__1)
	updateReq.SetNotificationInfo([]alerting.NotificationInfoAlertRule{*notif})

	updateResp, _, err := apiClient.AlertRulesAPI.UpdateAlertRule(env.ctx, ruleID).
		UpdateAlertRuleRequest(*updateReq).
		Execute()
	require.NoError(t, err)
	require.NotEmpty(t, updateResp.GetMessage())

	getAfterUpdate, _, err := apiClient.AlertRulesAPI.GetAlertRuleById(env.ctx, ruleID).Execute()
	require.NoError(t, err)
	require.Equal(t, updatedName, getAfterUpdate.GetName())

	disableReq := alerting.NewUpdateAlertRulesStatusRequest(2, []int64{ruleID})
	disableResp, _, err := apiClient.AlertRulesAPI.UpdateAlertRulesStatus(env.ctx).
		UpdateAlertRulesStatusRequest(*disableReq).
		Execute()
	require.NoError(t, err)
	require.True(t, disableResp.GetSuccess())

	enableReq := alerting.NewUpdateAlertRulesStatusRequest(1, []int64{ruleID})
	enableResp, _, err := apiClient.AlertRulesAPI.UpdateAlertRulesStatus(env.ctx).
		UpdateAlertRulesStatusRequest(*enableReq).
		Execute()
	require.NoError(t, err)
	require.True(t, enableResp.GetSuccess())

	delReq := alerting.NewDeleteAlertRulesRequest([]int64{ruleID})
	delResp, _, err := apiClient.AlertRulesAPI.DeleteAlertRules(env.ctx).
		DeleteAlertRulesRequest(*delReq).
		Execute()
	require.NoError(t, err)
	require.True(t, delResp.GetSuccess())
	require.Contains(t, delResp.GetSuccessfulIds(), ruleID)
	ruleID = 0
}

func TestAlertsFunctionalReadOnly(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetAlertingClient(env.ctx)

	listResp, _, err := apiClient.AlertsAPI.ListAlerts(env.ctx).
		Limit(10).
		Offset(0).
		Execute()
	require.NoError(t, err)

	if listResp == nil {
		return
	}

	var firstAlertID string
	if fullList := listResp.ListAlertsResponse; fullList != nil && len(fullList.GetAlerts()) > 0 {
		if a := fullList.GetAlerts()[0].Alert; a != nil {
			firstAlertID = a.GetAlertId()
		}
	}

	if firstAlertID == "" {
		t.Log("no existing alerts found; skipping get-by-ID and status update sub-tests")
		return
	}

	getResp, getHTTP, err := apiClient.AlertsAPI.GetAlertById(env.ctx, firstAlertID).Execute()
	if err != nil {
		require.NotNil(t, getHTTP)
		require.Contains(t, []int{http.StatusOK, http.StatusNotFound}, getHTTP.StatusCode)
	} else {
		require.Equal(t, firstAlertID, getResp.GetAlertId())
	}

	statusReq := alerting.NewUpdateAlertsStatusRequest(2, []string{firstAlertID})
	statusResp, _, err := apiClient.AlertsAPI.UpdateAlertsStatus(env.ctx).
		UpdateAlertsStatusRequest(*statusReq).
		Execute()
	require.NoError(t, err)
	_ = statusResp.GetSuccess()
}

func TestAlertsCountOnly(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetAlertingClient(env.ctx)

	filters := alerting.NewFiltersAlertObject()
	filters.SetOnlyActiveAlertsCount(true)

	resp, _, err := apiClient.AlertsAPI.ListAlerts(env.ctx).
		Filters(*filters).
		Execute()
	require.NoError(t, err)

	if resp != nil && resp.TotalCountAlertsResponse != nil {
		require.GreaterOrEqual(t, resp.TotalCountAlertsResponse.GetTotal(), int64(0))
	}
}

func TestSendAlertNotificationsFunctional(t *testing.T) {
	env := newFunctionalEnv(t)

	testEmail := strings.TrimSpace(os.Getenv("GO_CISCOSECUREACCESS_ALERT_TEST_EMAIL"))
	if testEmail == "" {
		t.Skip("set GO_CISCOSECUREACCESS_ALERT_TEST_EMAIL to run test notification tests")
	}

	apiClient := env.factory.GetAlertingClient(env.ctx)

	notif := alerting.NewNotificationInfoAlertRule()
	notif.SetType(alerting.NOTIFICATIONTYPEALL_EMAIL)
	notif.SetRecipients([]string{testEmail})

	testNotif := alerting.NewTestNotification(
		"go-sdk-test-notification",
		1,
		alerting.SEVERITYALERT__2,
		[]alerting.NotificationInfoAlertRule{*notif},
	)

	resp, httpResp, err := apiClient.AlertsAPI.SendAlertNotifications(env.ctx).
		TestNotification(*testNotif).
		Execute()
	require.NoError(t, err)
	require.True(t, isHTTP2xx(httpResp))
	require.NotEmpty(t, resp.GetStatus())
}

func TestAlertRulesNegative(t *testing.T) {
	env := newFunctionalEnv(t)
	api := env.factory.GetAlertingClient(env.ctx)

	t.Run("GetRuleNotFound", func(t *testing.T) {
		_, resp, err := api.AlertRulesAPI.GetAlertRuleById(env.ctx, 999999999).Execute()
		require.Error(t, err)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("UpdateRuleNotFound", func(t *testing.T) {
		upd := alerting.NewUpdateAlertRuleRequest()
		upd.SetName("no-such-rule")
		_, resp, err := api.AlertRulesAPI.UpdateAlertRule(env.ctx, 999999999).UpdateAlertRuleRequest(*upd).Execute()
		require.Error(t, err)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("CreateEmptyName", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("", alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 1)
		_, resp, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.Error(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("CreateInvalidRuleTypeId", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("go-sdk-bad-type", alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 9999)
		_, resp, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.Error(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("CreateDuplicateName", func(t *testing.T) {
		name := uniqueName("go-sdk-dup")
		req := alerting.NewCreateAlertRuleRequest(name, alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 1)
		_, _, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.NoError(t, err)

		t.Cleanup(func() {
			rules, _, _ := api.AlertRulesAPI.ListAlertRules(env.ctx).Execute()
			for _, r := range rules {
				if r.GetName() == name {
					dr := alerting.NewDeleteAlertRulesRequest([]int64{r.GetId()})
					api.AlertRulesAPI.DeleteAlertRules(env.ctx).DeleteAlertRulesRequest(*dr).Execute()
				}
			}
		})

		_, resp, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.Error(t, err)
		require.Equal(t, 409, resp.StatusCode)
	})

	t.Run("CreateTooManyRecipients", func(t *testing.T) {
		recipients := make([]string, 51)
		for i := range recipients {
			recipients[i] = fmt.Sprintf("u%d@example.com", i)
		}
		notif := alerting.NewNotificationInfoAlertRule()
		notif.SetType(alerting.NOTIFICATIONTYPEALL_EMAIL)
		notif.SetRecipients(recipients)

		req := alerting.NewCreateAlertRuleRequest("go-sdk-too-many-recip", alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 1)
		req.SetNotificationInfo([]alerting.NotificationInfoAlertRule{*notif})

		_, resp, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.Error(t, err)
		require.Equal(t, 422, resp.StatusCode)
	})

	t.Run("DeleteNonexistentIds", func(t *testing.T) {
		req := alerting.NewDeleteAlertRulesRequest([]int64{999999991, 999999992})
		resp, httpResp, err := api.AlertRulesAPI.DeleteAlertRules(env.ctx).DeleteAlertRulesRequest(*req).Execute()
		require.NoError(t, err)
		require.Equal(t, 207, httpResp.StatusCode)
		require.False(t, resp.GetSuccess())
		require.ElementsMatch(t, []int64{999999991, 999999992}, resp.GetErrorIds())
		require.Empty(t, resp.GetSuccessfulIds())
	})

	t.Run("BulkStatusUpdateNonexistentIds", func(t *testing.T) {
		req := alerting.NewUpdateAlertRulesStatusRequest(1, []int64{999999991, 999999992})
		resp, httpResp, err := api.AlertRulesAPI.UpdateAlertRulesStatus(env.ctx).UpdateAlertRulesStatusRequest(*req).Execute()
		require.NoError(t, err)
		require.Equal(t, 207, httpResp.StatusCode)
		require.False(t, resp.GetSuccess())
		require.ElementsMatch(t, []int64{999999991, 999999992}, resp.GetErrorIds())
	})
}

func TestAlertsNegative(t *testing.T) {
	env := newFunctionalEnv(t)
	api := env.factory.GetAlertingClient(env.ctx)

	t.Run("GetAlertNotFound", func(t *testing.T) {
		_, resp, err := api.AlertsAPI.GetAlertById(env.ctx, "AL-DOES-NOT-EXIST").Execute()
		require.Error(t, err)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("BulkAlertStatusUpdateNonexistentId", func(t *testing.T) {
		req := alerting.NewUpdateAlertsStatusRequest(2, []string{"AL-FAKE-000"})
		resp, httpResp, err := api.AlertsAPI.UpdateAlertsStatus(env.ctx).UpdateAlertsStatusRequest(*req).Execute()
		require.NoError(t, err)
		require.Equal(t, 207, httpResp.StatusCode)
		require.False(t, resp.GetSuccess())
		require.Contains(t, resp.GetErrorIds(), "AL-FAKE-000")
	})

	t.Run("DeleteEmptyRuleIds", func(t *testing.T) {
		req := alerting.NewDeleteAlertRulesRequest([]int64{})
		_, resp, err := api.AlertRulesAPI.DeleteAlertRules(env.ctx).DeleteAlertRulesRequest(*req).Execute()
		require.Error(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("BulkRuleStatusEmptyEntityIds", func(t *testing.T) {
		req := alerting.NewUpdateAlertRulesStatusRequest(1, []int64{})
		_, resp, err := api.AlertRulesAPI.UpdateAlertRulesStatus(env.ctx).UpdateAlertRulesStatusRequest(*req).Execute()
		require.Error(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("BulkAlertStatusEmptyEntityIds", func(t *testing.T) {
		req := alerting.NewUpdateAlertsStatusRequest(2, []string{})
		_, resp, err := api.AlertsAPI.UpdateAlertsStatus(env.ctx).UpdateAlertsStatusRequest(*req).Execute()
		require.Error(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("CreateNameExceedsMaxLength", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest(strings.Repeat("a", 101), alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 1)
		_, resp, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.Error(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("CreateInvalidSeverityBypass", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("go-sdk-bad-severity", alerting.SeverityAlert(999), alerting.STATUSALERTRULE__1, 1)
		_, resp, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.Error(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})
}

func TestAlertRulesConditionsAndWebhook(t *testing.T) {
	env := newFunctionalEnv(t)
	api := env.factory.GetAlertingClient(env.ctx)

	t.Run("CreateWithConditions", func(t *testing.T) {
		cond := alerting.NewConditionsAlertRule()
		cond.SetMatchType("all")
		cond.SetRows([]alerting.ConditionsAlertRuleRowsInner{
			*alerting.NewConditionsAlertRuleRowsInner("region", "US West-1"),
		})
		name := uniqueName("go-sdk-cond")
		req := alerting.NewCreateAlertRuleRequest(name, alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 1)
		req.SetConditions(*cond)

		_, createHTTP, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.NoError(t, err)
		require.Equal(t, 201, createHTTP.StatusCode)

		var ruleID int64
		t.Cleanup(func() {
			if ruleID == 0 {
				return
			}
			dr := alerting.NewDeleteAlertRulesRequest([]int64{ruleID})
			api.AlertRulesAPI.DeleteAlertRules(env.ctx).DeleteAlertRulesRequest(*dr).Execute()
		})

		rules, _, err := api.AlertRulesAPI.ListAlertRules(env.ctx).Execute()
		require.NoError(t, err)
		for _, r := range rules {
			if r.GetName() == name {
				ruleID = r.GetId()
				break
			}
		}
		require.NotZero(t, ruleID)

		got, _, err := api.AlertRulesAPI.GetAlertRuleById(env.ctx, ruleID).Execute()
		require.NoError(t, err)
		gotCond, hasCond := got.GetConditionsOk()
		require.True(t, hasCond)
		require.Equal(t, "all", gotCond.GetMatchType())
		require.Len(t, gotCond.GetRows(), 1)
		require.Equal(t, "region", gotCond.GetRows()[0].GetField())
		require.Equal(t, "US West-1", gotCond.GetRows()[0].GetValue())

		updCond := alerting.NewConditionsAlertRule()
		updCond.SetMatchType("any")
		updCond.SetRows([]alerting.ConditionsAlertRuleRowsInner{
			*alerting.NewConditionsAlertRuleRowsInner("region", "US East-1"),
			*alerting.NewConditionsAlertRuleRowsInner("region", "US West-2"),
		})
		upd := alerting.NewUpdateAlertRuleRequest()
		upd.SetConditions(*updCond)
		_, _, err = api.AlertRulesAPI.UpdateAlertRule(env.ctx, ruleID).UpdateAlertRuleRequest(*upd).Execute()
		require.NoError(t, err)

		got2, _, err := api.AlertRulesAPI.GetAlertRuleById(env.ctx, ruleID).Execute()
		require.NoError(t, err)
		gotCond2, hasCond2 := got2.GetConditionsOk()
		require.True(t, hasCond2)
		require.Equal(t, "any", gotCond2.GetMatchType())
		require.Len(t, gotCond2.GetRows(), 2)
	})

	t.Run("CreateWithWebhookNotification", func(t *testing.T) {
		notif := alerting.NewNotificationInfoAlertRule()
		notif.SetType(alerting.NOTIFICATIONTYPEALL_WEBHOOK)
		notif.SetWebhookIds([]string{"fake-webhook-id"})
		name := uniqueName("go-sdk-webhook")
		req := alerting.NewCreateAlertRuleRequest(name, alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 1)
		req.SetNotificationInfo([]alerting.NotificationInfoAlertRule{*notif})

		_, createHTTP, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.NoError(t, err)
		require.Equal(t, 201, createHTTP.StatusCode)

		t.Cleanup(func() {
			rules, _, _ := api.AlertRulesAPI.ListAlertRules(env.ctx).Execute()
			for _, r := range rules {
				if r.GetName() == name {
					dr := alerting.NewDeleteAlertRulesRequest([]int64{r.GetId()})
					api.AlertRulesAPI.DeleteAlertRules(env.ctx).DeleteAlertRulesRequest(*dr).Execute()
				}
			}
		})
	})
}

func TestAlertRulesAllRuleTypeIds(t *testing.T) {
	env := newFunctionalEnv(t)
	api := env.factory.GetAlertingClient(env.ctx)

	for _, typeID := range []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17} {
		typeID := typeID
		t.Run(fmt.Sprintf("RuleTypeId_%d", typeID), func(t *testing.T) {
			name := uniqueName(fmt.Sprintf("go-sdk-type-%d", typeID))
			req := alerting.NewCreateAlertRuleRequest(name, alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, typeID)

			_, createHTTP, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
			require.NoError(t, err)
			require.Equal(t, 201, createHTTP.StatusCode)

			t.Cleanup(func() {
				rules, _, _ := api.AlertRulesAPI.ListAlertRules(env.ctx).Execute()
				for _, r := range rules {
					if r.GetName() == name {
						dr := alerting.NewDeleteAlertRulesRequest([]int64{r.GetId()})
						api.AlertRulesAPI.DeleteAlertRules(env.ctx).DeleteAlertRulesRequest(*dr).Execute()
					}
				}
			})
		})
	}
}

func TestAlertRulesBoundaryValues(t *testing.T) {
	env := newFunctionalEnv(t)
	api := env.factory.GetAlertingClient(env.ctx)

	t.Run("NameAtMaxLength", func(t *testing.T) {
		name := strings.Repeat("a", 100)
		req := alerting.NewCreateAlertRuleRequest(name, alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 1)
		_, createHTTP, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.NoError(t, err)
		require.Equal(t, 201, createHTTP.StatusCode)
		t.Cleanup(func() {
			rules, _, _ := api.AlertRulesAPI.ListAlertRules(env.ctx).Execute()
			for _, r := range rules {
				if r.GetName() == name {
					dr := alerting.NewDeleteAlertRulesRequest([]int64{r.GetId()})
					api.AlertRulesAPI.DeleteAlertRules(env.ctx).DeleteAlertRulesRequest(*dr).Execute()
				}
			}
		})
	})

	t.Run("NameOverMaxLength", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest(strings.Repeat("a", 101), alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 1)
		_, resp, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
		require.Error(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("AllSeverityLevels", func(t *testing.T) {
		for _, sev := range []alerting.SeverityAlert{
			alerting.SEVERITYALERT__1,
			alerting.SEVERITYALERT__2,
			alerting.SEVERITYALERT__3,
			alerting.SEVERITYALERT__4,
		} {
			sev := sev
			name := uniqueName(fmt.Sprintf("go-sdk-sev-%d", sev))
			req := alerting.NewCreateAlertRuleRequest(name, sev, alerting.STATUSALERTRULE__1, 1)
			_, createHTTP, err := api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()
			require.NoError(t, err)
			require.Equal(t, 201, createHTTP.StatusCode)
			t.Cleanup(func() {
				rules, _, _ := api.AlertRulesAPI.ListAlertRules(env.ctx).Execute()
				for _, r := range rules {
					if r.GetName() == name {
						dr := alerting.NewDeleteAlertRulesRequest([]int64{r.GetId()})
						api.AlertRulesAPI.DeleteAlertRules(env.ctx).DeleteAlertRulesRequest(*dr).Execute()
					}
				}
			})
		}
	})

	t.Run("MixedValidInvalidDeleteIds", func(t *testing.T) {
		name := uniqueName("go-sdk-mixed-del")
		req := alerting.NewCreateAlertRuleRequest(name, alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 1)
		api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()

		rules, _, _ := api.AlertRulesAPI.ListAlertRules(env.ctx).Execute()
		var validID int64
		for _, r := range rules {
			if r.GetName() == name {
				validID = r.GetId()
				break
			}
		}
		require.NotZero(t, validID)

		delReq := alerting.NewDeleteAlertRulesRequest([]int64{validID, 999999999})
		resp, httpResp, err := api.AlertRulesAPI.DeleteAlertRules(env.ctx).DeleteAlertRulesRequest(*delReq).Execute()
		require.NoError(t, err)
		require.Equal(t, 207, httpResp.StatusCode)
		require.False(t, resp.GetSuccess())
		require.Contains(t, resp.GetSuccessfulIds(), validID)
		require.Contains(t, resp.GetErrorIds(), int64(999999999))
	})

	t.Run("MixedValidInvalidBulkStatusIds", func(t *testing.T) {
		name := uniqueName("go-sdk-mixed-status")
		req := alerting.NewCreateAlertRuleRequest(name, alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 1)
		api.AlertRulesAPI.CreateAlertRule(env.ctx).CreateAlertRuleRequest(*req).Execute()

		rules, _, _ := api.AlertRulesAPI.ListAlertRules(env.ctx).Execute()
		var validID int64
		for _, r := range rules {
			if r.GetName() == name {
				validID = r.GetId()
				break
			}
		}
		require.NotZero(t, validID)

		t.Cleanup(func() {
			dr := alerting.NewDeleteAlertRulesRequest([]int64{validID})
			api.AlertRulesAPI.DeleteAlertRules(env.ctx).DeleteAlertRulesRequest(*dr).Execute()
		})

		stReq := alerting.NewUpdateAlertRulesStatusRequest(2, []int64{validID, 999999999})
		resp, httpResp, err := api.AlertRulesAPI.UpdateAlertRulesStatus(env.ctx).UpdateAlertRulesStatusRequest(*stReq).Execute()
		require.NoError(t, err)
		require.Equal(t, 207, httpResp.StatusCode)
		require.False(t, resp.GetSuccess())
		require.Contains(t, resp.GetSuccessfulIds(), validID)
		require.Contains(t, resp.GetErrorIds(), int64(999999999))
	})
}

func TestListAlertsFilters(t *testing.T) {
	env := newFunctionalEnv(t)
	api := env.factory.GetAlertingClient(env.ctx)

	t.Run("FilterBySeverity", func(t *testing.T) {
		f := alerting.NewFiltersAlertObject()
		f.SetSeverity(alerting.SEVERITYALERT__1)
		_, resp, err := api.AlertsAPI.ListAlerts(env.ctx).Filters(*f).Limit(5).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("FilterByStatus", func(t *testing.T) {
		f := alerting.NewFiltersAlertObject()
		f.SetStatus(alerting.STATUSALERT__1)
		_, resp, err := api.AlertsAPI.ListAlerts(env.ctx).Filters(*f).Limit(5).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("FilterByAlertName", func(t *testing.T) {
		f := alerting.NewFiltersAlertObject()
		f.SetAlertName("test")
		_, resp, err := api.AlertsAPI.ListAlerts(env.ctx).Filters(*f).Limit(5).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("FilterByPatternSearch", func(t *testing.T) {
		f := alerting.NewFiltersAlertObject()
		f.SetPatternSearch("Network")
		_, resp, err := api.AlertsAPI.ListAlerts(env.ctx).Filters(*f).Limit(5).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("FilterByCreatedAfter", func(t *testing.T) {
		f := alerting.NewFiltersAlertObject()
		f.SetCreatedAfter(time.Now().AddDate(0, -1, 0))
		_, resp, err := api.AlertsAPI.ListAlerts(env.ctx).Filters(*f).Limit(5).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("FilterByTimeRange", func(t *testing.T) {
		tr := alerting.NewTimeRange()
		tr.SetStartTime(time.Now().AddDate(0, -1, 0))
		tr.SetEndTime(time.Now())
		f := alerting.NewFiltersAlertObject()
		f.SetTimeRange(*tr)
		_, resp, err := api.AlertsAPI.ListAlerts(env.ctx).Filters(*f).Limit(5).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("IncludeContext", func(t *testing.T) {
		f := alerting.NewFiltersAlertObject()
		f.SetIncludeContext(true)
		resp, httpResp, err := api.AlertsAPI.ListAlerts(env.ctx).Filters(*f).Limit(5).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		_ = resp
	})

	t.Run("OffsetBeyondTotal", func(t *testing.T) {
		_, resp, err := api.AlertsAPI.ListAlerts(env.ctx).Limit(10).Offset(999999).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	})
}
