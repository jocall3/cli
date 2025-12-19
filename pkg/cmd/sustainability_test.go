// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestSustainabilityPurchaseCarbonOffsets(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sustainability", "purchase-carbon-offsets",
		"--amount-kg-co2e", "500",
		"--offset-project", "Verified Carbon Standard Project X",
		"--payment-account-id", "acc_chase_checking_4567",
	)
}

func TestSustainabilityRetrieveCarbonFootprint(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sustainability", "retrieve-carbon-footprint",
	)
}
