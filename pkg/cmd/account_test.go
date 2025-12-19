// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestAccountsLink(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts", "link",
		"--country-code", "US",
		"--institution-name", "Bank of America",
		"--provider-identifier", "{}",
		"--redirect-uri", "{}",
	)
}

func TestAccountsRetrieveDetails(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts", "retrieve-details",
		"--account-id", "acc_chase_checking_4567",
	)
}

func TestAccountsRetrieveMe(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts", "retrieve-me",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestAccountsRetrieveStatements(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts", "retrieve-statements",
		"--account-id", "acc_chase_checking_4567",
		"--month", "7",
		"--year", "2024",
		"--format", "pdf",
	)
}
