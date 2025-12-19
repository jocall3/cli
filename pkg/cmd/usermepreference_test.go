// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestUsersMePreferencesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me:preferences", "retrieve",
	)
}

func TestUsersMePreferencesUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me:preferences", "update",
		"--ai-interaction-mode", "proactive",
		"--data-sharing-consent",
		"--notification-channels", "{email: true, inApp: true, push: true, sms: false}",
		"--preferred-language", "en-US",
		"--theme", "Dark-Quantum",
		"--transaction-grouping", "category",
	)
}
