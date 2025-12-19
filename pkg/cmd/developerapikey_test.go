// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestDevelopersAPIKeysCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"developers:api-keys", "create",
		"--name", "My Analytics Service Key",
		"--scope", "read:accounts",
		"--scope", "read:transactions",
		"--expires-in-days", "90",
	)
}

func TestDevelopersAPIKeysList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"developers:api-keys", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestDevelopersAPIKeysRevoke(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"developers:api-keys", "revoke",
		"--key-id", "api_key_dev_app_01",
	)
}
