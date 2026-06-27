// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

package alerting

import (
	"errors"
	"fmt"
	"strings"
)

const (
	maxAlertRuleNameLen        = 100
	maxAlertRuleDescriptionLen = 100
	maxRuleTypeId              = 17
	minRuleTypeId              = 1
	maxBulkIds                 = 100
	maxEmailRecipients         = 50
)

// Validate checks all constraints on CreateAlertRuleRequest before sending to the API.
// Returns a non-nil error describing every violation found.
func (r *CreateAlertRuleRequest) Validate() error {
	var errs []string

	if r.GetName() == "" {
		errs = append(errs, "name is required")
	} else if len(r.GetName()) > maxAlertRuleNameLen {
		errs = append(errs, fmt.Sprintf("name must be at most %d characters", maxAlertRuleNameLen))
	}

	if desc, ok := r.GetDescriptionOk(); ok && desc != nil && len(*desc) > maxAlertRuleDescriptionLen {
		errs = append(errs, fmt.Sprintf("description must be at most %d characters", maxAlertRuleDescriptionLen))
	}

	if !r.GetSeverity().IsValid() {
		errs = append(errs, fmt.Sprintf("severity %d is not valid; allowed: 1 (High), 2 (Medium), 3 (Low), 4 (Info)", r.GetSeverity()))
	}

	if !r.GetStatus().IsValid() {
		errs = append(errs, fmt.Sprintf("status %d is not valid; allowed: 1 (Enabled), 2 (Disabled)", r.GetStatus()))
	}

	if r.GetRuleTypeId() < minRuleTypeId || r.GetRuleTypeId() > maxRuleTypeId {
		errs = append(errs, fmt.Sprintf("rule_type_id %d is not valid; allowed: %d–%d", r.GetRuleTypeId(), minRuleTypeId, maxRuleTypeId))
	}

	errs = append(errs, validateNotificationInfo(r.GetNotificationInfo())...)

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}
	return nil
}

// Validate checks all constraints on UpdateAlertRuleRequest before sending to the API.
func (r *UpdateAlertRuleRequest) Validate() error {
	var errs []string

	if name, ok := r.GetNameOk(); ok && name != nil {
		if len(*name) == 0 {
			errs = append(errs, "name must not be empty")
		} else if len(*name) > maxAlertRuleNameLen {
			errs = append(errs, fmt.Sprintf("name must be at most %d characters", maxAlertRuleNameLen))
		}
	}

	if desc, ok := r.GetDescriptionOk(); ok && desc != nil && len(*desc) > maxAlertRuleDescriptionLen {
		errs = append(errs, fmt.Sprintf("description must be at most %d characters", maxAlertRuleDescriptionLen))
	}

	if sev, ok := r.GetSeverityOk(); ok && sev != nil && !sev.IsValid() {
		errs = append(errs, fmt.Sprintf("severity %d is not valid; allowed: 1 (High), 2 (Medium), 3 (Low), 4 (Info)", *sev))
	}

	if st, ok := r.GetStatusOk(); ok && st != nil && !st.IsValid() {
		errs = append(errs, fmt.Sprintf("status %d is not valid; allowed: 1 (Enabled), 2 (Disabled)", *st))
	}

	if tid, ok := r.GetRuleTypeIdOk(); ok && tid != nil {
		if *tid < minRuleTypeId || *tid > maxRuleTypeId {
			errs = append(errs, fmt.Sprintf("rule_type_id %d is not valid; allowed: %d–%d", *tid, minRuleTypeId, maxRuleTypeId))
		}
	}

	errs = append(errs, validateNotificationInfo(r.GetNotificationInfo())...)

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}
	return nil
}

// Validate checks all constraints on DeleteAlertRulesRequest before sending to the API.
func (r *DeleteAlertRulesRequest) Validate() error {
	ids := r.GetRuleIds()
	if len(ids) == 0 {
		return errors.New("ruleIds must not be empty")
	}
	if len(ids) > maxBulkIds {
		return fmt.Errorf("ruleIds must contain at most %d items, got %d", maxBulkIds, len(ids))
	}
	return nil
}

// Validate checks all constraints on UpdateAlertRulesStatusRequest before sending to the API.
func (r *UpdateAlertRulesStatusRequest) Validate() error {
	var errs []string

	if r.GetStatus() != 1 && r.GetStatus() != 2 {
		errs = append(errs, fmt.Sprintf("status %d is not valid; allowed: 1 (Enabled), 2 (Disabled)", r.GetStatus()))
	}

	ids := r.GetEntityIds()
	if len(ids) == 0 {
		errs = append(errs, "entity_ids must not be empty")
	} else if len(ids) > maxBulkIds {
		errs = append(errs, fmt.Sprintf("entity_ids must contain at most %d items, got %d", maxBulkIds, len(ids)))
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}
	return nil
}

// Validate checks all constraints on UpdateAlertsStatusRequest before sending to the API.
func (r *UpdateAlertsStatusRequest) Validate() error {
	var errs []string

	if r.GetStatus() != 1 && r.GetStatus() != 2 {
		errs = append(errs, fmt.Sprintf("status %d is not valid; allowed: 1 (Active/Dismissed)", r.GetStatus()))
	}

	ids := r.GetEntityIds()
	if len(ids) == 0 {
		errs = append(errs, "entity_ids must not be empty")
	} else if len(ids) > maxBulkIds {
		errs = append(errs, fmt.Sprintf("entity_ids must contain at most %d items, got %d", maxBulkIds, len(ids)))
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}
	return nil
}

func validateNotificationInfo(notifs []NotificationInfoAlertRule) []string {
	var errs []string
	for i, n := range notifs {
		if !n.GetType().IsValid() {
			errs = append(errs, fmt.Sprintf("notification_info[%d].type %q is not valid; allowed: email, webhook", i, n.GetType()))
		}
		if n.GetType() == NOTIFICATIONTYPEALL_EMAIL {
			if len(n.GetRecipients()) == 0 {
				errs = append(errs, fmt.Sprintf("notification_info[%d].recipients must not be empty for email notifications", i))
			} else if len(n.GetRecipients()) > maxEmailRecipients {
				errs = append(errs, fmt.Sprintf("notification_info[%d].recipients must contain at most %d items, got %d", i, maxEmailRecipients, len(n.GetRecipients())))
			}
		}
	}
	return errs
}
