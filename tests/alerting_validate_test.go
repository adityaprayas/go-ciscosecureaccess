//go:build functional

// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

package tests

import (
	"strings"
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/alerting"
	"github.com/stretchr/testify/require"
)

func TestCreateAlertRuleRequestValidate(t *testing.T) {
	validNotif := func() alerting.NotificationInfoAlertRule {
		n := alerting.NewNotificationInfoAlertRule()
		n.SetType(alerting.NOTIFICATIONTYPEALL_EMAIL)
		n.SetRecipients([]string{"a@b.com"})
		return *n
	}

	t.Run("Valid", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("my-rule", alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 1)
		require.NoError(t, req.Validate())
	})

	t.Run("EmptyName", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("", alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 1)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "name is required")
	})

	t.Run("NameTooLong", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest(strings.Repeat("a", 101), alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 1)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "name must be at most")
	})

	t.Run("NameAtMaxLength", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest(strings.Repeat("a", 100), alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 1)
		require.NoError(t, req.Validate())
	})

	t.Run("DescriptionTooLong", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("ok-name", alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 1)
		req.SetDescription(strings.Repeat("d", 101))
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "description must be at most")
	})

	t.Run("InvalidSeverity", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("ok-name", alerting.SeverityAlert(999), alerting.STATUSALERTRULE__1, 1)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "severity")
	})

	t.Run("InvalidStatus", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("ok-name", alerting.SEVERITYALERT__1, alerting.StatusAlertRule(99), 1)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "status")
	})

	t.Run("RuleTypeIdTooLow", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("ok-name", alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 0)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "rule_type_id")
	})

	t.Run("RuleTypeIdTooHigh", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("ok-name", alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 18)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "rule_type_id")
	})

	t.Run("AllValidRuleTypeIds", func(t *testing.T) {
		for i := int64(1); i <= 17; i++ {
			req := alerting.NewCreateAlertRuleRequest("ok-name", alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, i)
			require.NoError(t, req.Validate(), "rule_type_id=%d", i)
		}
	})

	t.Run("TooManyEmailRecipients", func(t *testing.T) {
		many := make([]string, 51)
		for i := range many {
			many[i] = "a@b.com"
		}
		n := alerting.NewNotificationInfoAlertRule()
		n.SetType(alerting.NOTIFICATIONTYPEALL_EMAIL)
		n.SetRecipients(many)
		req := alerting.NewCreateAlertRuleRequest("ok-name", alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 1)
		req.SetNotificationInfo([]alerting.NotificationInfoAlertRule{*n})
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "recipients must contain at most")
	})

	t.Run("EmailNotifWithNoRecipients", func(t *testing.T) {
		n := alerting.NewNotificationInfoAlertRule()
		n.SetType(alerting.NOTIFICATIONTYPEALL_EMAIL)
		n.SetRecipients([]string{})
		req := alerting.NewCreateAlertRuleRequest("ok-name", alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 1)
		req.SetNotificationInfo([]alerting.NotificationInfoAlertRule{*n})
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "recipients must not be empty")
	})

	t.Run("InvalidNotifType", func(t *testing.T) {
		n := alerting.NewNotificationInfoAlertRule()
		n.SetType(alerting.NotificationTypeAll("sms"))
		req := alerting.NewCreateAlertRuleRequest("ok-name", alerting.SEVERITYALERT__1, alerting.STATUSALERTRULE__1, 1)
		req.SetNotificationInfo([]alerting.NotificationInfoAlertRule{*n})
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "type")
	})

	t.Run("MultipleErrors", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("", alerting.SeverityAlert(0), alerting.StatusAlertRule(0), 99)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "name")
		require.Contains(t, err.Error(), "severity")
		require.Contains(t, err.Error(), "status")
		require.Contains(t, err.Error(), "rule_type_id")
	})

	t.Run("ValidWithNotification", func(t *testing.T) {
		req := alerting.NewCreateAlertRuleRequest("ok-name", alerting.SEVERITYALERT__2, alerting.STATUSALERTRULE__1, 5)
		req.SetNotificationInfo([]alerting.NotificationInfoAlertRule{validNotif()})
		require.NoError(t, req.Validate())
	})
}

func TestUpdateAlertRuleRequestValidate(t *testing.T) {
	t.Run("AllFieldsValid", func(t *testing.T) {
		req := alerting.NewUpdateAlertRuleRequest()
		req.SetName("updated-name")
		req.SetSeverity(alerting.SEVERITYALERT__3)
		req.SetStatus(alerting.STATUSALERTRULE__2)
		req.SetRuleTypeId(7)
		require.NoError(t, req.Validate())
	})

	t.Run("EmptyRequest", func(t *testing.T) {
		req := alerting.NewUpdateAlertRuleRequest()
		require.NoError(t, req.Validate())
	})

	t.Run("EmptyNameWhenSet", func(t *testing.T) {
		req := alerting.NewUpdateAlertRuleRequest()
		req.SetName("")
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "name must not be empty")
	})

	t.Run("NameTooLong", func(t *testing.T) {
		req := alerting.NewUpdateAlertRuleRequest()
		req.SetName(strings.Repeat("b", 101))
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "name must be at most")
	})

	t.Run("InvalidSeverity", func(t *testing.T) {
		req := alerting.NewUpdateAlertRuleRequest()
		req.SetSeverity(alerting.SeverityAlert(5))
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "severity")
	})

	t.Run("InvalidRuleTypeId", func(t *testing.T) {
		req := alerting.NewUpdateAlertRuleRequest()
		req.SetRuleTypeId(18)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "rule_type_id")
	})
}

func TestDeleteAlertRulesRequestValidate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		req := alerting.NewDeleteAlertRulesRequest([]int64{1, 2, 3})
		require.NoError(t, req.Validate())
	})

	t.Run("EmptyIds", func(t *testing.T) {
		req := alerting.NewDeleteAlertRulesRequest([]int64{})
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "ruleIds must not be empty")
	})

	t.Run("TooManyIds", func(t *testing.T) {
		ids := make([]int64, 101)
		for i := range ids {
			ids[i] = int64(i + 1)
		}
		req := alerting.NewDeleteAlertRulesRequest(ids)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "at most 100")
	})

	t.Run("ExactlyMaxIds", func(t *testing.T) {
		ids := make([]int64, 100)
		for i := range ids {
			ids[i] = int64(i + 1)
		}
		req := alerting.NewDeleteAlertRulesRequest(ids)
		require.NoError(t, req.Validate())
	})
}

func TestUpdateAlertRulesStatusRequestValidate(t *testing.T) {
	t.Run("ValidEnable", func(t *testing.T) {
		req := alerting.NewUpdateAlertRulesStatusRequest(1, []int64{1, 2})
		require.NoError(t, req.Validate())
	})

	t.Run("ValidDisable", func(t *testing.T) {
		req := alerting.NewUpdateAlertRulesStatusRequest(2, []int64{1})
		require.NoError(t, req.Validate())
	})

	t.Run("InvalidStatus", func(t *testing.T) {
		req := alerting.NewUpdateAlertRulesStatusRequest(3, []int64{1})
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "status")
	})

	t.Run("EmptyEntityIds", func(t *testing.T) {
		req := alerting.NewUpdateAlertRulesStatusRequest(1, []int64{})
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "entity_ids must not be empty")
	})

	t.Run("TooManyEntityIds", func(t *testing.T) {
		ids := make([]int64, 101)
		req := alerting.NewUpdateAlertRulesStatusRequest(1, ids)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "at most 100")
	})
}

func TestUpdateAlertsStatusRequestValidate(t *testing.T) {
	t.Run("ValidDismiss", func(t *testing.T) {
		req := alerting.NewUpdateAlertsStatusRequest(2, []string{"AL-abc", "AL-def"})
		require.NoError(t, req.Validate())
	})

	t.Run("InvalidStatus", func(t *testing.T) {
		req := alerting.NewUpdateAlertsStatusRequest(5, []string{"AL-abc"})
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "status")
	})

	t.Run("EmptyEntityIds", func(t *testing.T) {
		req := alerting.NewUpdateAlertsStatusRequest(1, []string{})
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "entity_ids must not be empty")
	})

	t.Run("TooManyEntityIds", func(t *testing.T) {
		ids := make([]string, 101)
		req := alerting.NewUpdateAlertsStatusRequest(1, ids)
		err := req.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "at most 100")
	})
}
