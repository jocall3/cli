// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestTransactionsRecurringCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions:recurring", "create",
		"--amount", "55.5",
		"--category", "Health & Fitness",
		"--currency", "USD",
		"--description", "New Gym Membership",
		"--frequency", "monthly",
		"--linked-account-id", "acc_chase_checking_4567",
		"--start-date", "2024-09-01",
	)
}

func TestTransactionsRecurringList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions:recurring", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}
