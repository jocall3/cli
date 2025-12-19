// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestNotificationsListUserNotifications(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications", "list-user-notifications",
		"--limit", "{}",
		"--offset", "{}",
		"--severity", "high",
		"--status", "unread",
	)
}

func TestNotificationsMarkAsRead(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications", "mark-as-read",
		"--notification-id", "notif_budget_alert_002",
	)
}
