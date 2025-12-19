// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestNotificationsSettingsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:settings", "retrieve",
	)
}

func TestNotificationsSettingsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:settings", "update",
		"--channel-preferences", "{email: true, inApp: true, push: true, sms: true}",
		"--event-preferences", "{aiInsights: true, budgetAlerts: true, promotionalOffers: false, securityAlerts: true, transactionAlerts: true}",
		"--quiet-hours", "{enabled: true, endTime: '08:00', startTime: '22:00'}",
	)
}
