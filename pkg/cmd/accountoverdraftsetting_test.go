// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestAccountsOverdraftSettingsRetrieveOverdraftSettings(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts:overdraft-settings", "retrieve-overdraft-settings",
		"--account-id", "acc_chase_checking_4567",
	)
}

func TestAccountsOverdraftSettingsUpdateOverdraftSettings(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts:overdraft-settings", "update-overdraft-settings",
		"--account-id", "acc_chase_checking_4567",
		"--enabled",
		"--fee-preference", "decline_if_over_limit",
		"--linked-savings-account-id", "acc_new_savings_5678",
		"--link-to-savings",
		"--protection-limit", "750",
	)
}
