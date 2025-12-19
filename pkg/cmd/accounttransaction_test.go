// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestAccountsTransactionsRetrievePending(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts:transactions", "retrieve-pending",
		"--account-id", "acc_chase_checking_4567",
		"--limit", "{}",
		"--offset", "{}",
	)
}
