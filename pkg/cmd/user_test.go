// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestUsersLogin(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users", "login",
		"--email", "quantum.visionary@demobank.com",
		"--password", "YourSecurePassword123",
		"--mfa-code", "123456",
	)
}

func TestUsersRegister(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users", "register",
		"--email", "alice.w@example.com",
		"--name", "Alice Wonderland",
		"--password", "SecureP@ssw0rd2024!",
		"--address", "{city: Anytown, country: USA, state: CA, street: 123 Main St, zip: '90210'}",
		"--date-of-birth", "1990-05-10",
		"--phone", "+1-555-987-6543",
	)
}
