// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestMarketplaceProductsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"marketplace:products", "list",
		"--ai-personalization-level", "high",
		"--category", "insurance",
		"--limit", "{}",
		"--min-rating", "4",
		"--offset", "{}",
	)
}

func TestMarketplaceProductsSimulateImpact(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"marketplace:products", "simulate-impact",
		"--product-id", "prod_home_insurance_quantum",
		"--simulation-parameters", "{loanAmount: 20000, repaymentTermMonths: 48}",
	)
}
