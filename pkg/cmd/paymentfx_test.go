// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestPaymentsFxConvert(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments:fx", "convert",
		"--source-account-id", "acc_chase_checking_4567",
		"--source-amount", "1000",
		"--source-currency", "USD",
		"--target-currency", "EUR",
		"--fx-rate-lock",
		"--target-account-id", "acc_euro_savings_9876",
	)
}

func TestPaymentsFxRetrieveRates(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments:fx", "retrieve-rates",
		"--base-currency", "USD",
		"--target-currency", "EUR",
		"--forecast-days", "7",
	)
}
