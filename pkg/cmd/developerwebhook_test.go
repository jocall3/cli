// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestDevelopersWebhooksCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"developers:webhooks", "create",
		"--callback-url", "https://my-analytics-app.com/webhooks/transactions",
		"--event", "transaction.created",
		"--event", "transaction.updated",
		"--secret", "my_custom_webhook_secret_123",
	)
}

func TestDevelopersWebhooksUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"developers:webhooks", "update",
		"--subscription-id", "whsub_devtool_finance_events",
		"--callback-url", "https://my-new-app.com/webhooks/demobank-events",
		"--event", "transaction.created",
		"--event", "user.login_failed",
		"--status", "paused",
	)
}

func TestDevelopersWebhooksList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"developers:webhooks", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestDevelopersWebhooksDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"developers:webhooks", "delete",
		"--subscription-id", "whsub_devtool_finance_events",
	)
}
