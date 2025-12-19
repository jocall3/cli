// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestUsersPasswordResetConfirm(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:password-reset", "confirm",
		"--identifier", "reset.user@example.com",
		"--new-password", "MyNewStrongPassword@789",
		"--verification-code", "654321",
	)
}

func TestUsersPasswordResetInitiate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:password-reset", "initiate",
		"--identifier", "reset.user@example.com",
	)
}
