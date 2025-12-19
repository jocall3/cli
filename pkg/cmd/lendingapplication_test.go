// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestLendingApplicationsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"lending:applications", "retrieve",
		"--application-id", "loan_app_creditflow-123",
	)
}

func TestLendingApplicationsSubmit(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"lending:applications", "submit",
		"--loan-amount", "10000",
		"--loan-purpose", "home_improvement",
		"--repayment-term-months", "36",
		"--additional-notes", "Funds needed to replace a broken HVAC system.",
		"--co-applicant", "{email: jane.doe@example.com, income: 75000, name: Jane Doe}",
	)
}
