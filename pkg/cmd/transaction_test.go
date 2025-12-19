// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestTransactionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions", "retrieve",
		"--transaction-id", "txn_quantum-2024-07-21-A7B8C9",
	)
}

func TestTransactionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions", "list",
		"--category", "Groceries",
		"--end-date", "2024-12-31",
		"--limit", "{}",
		"--max-amount", "100",
		"--min-amount", "20",
		"--offset", "{}",
		"--search-query", "Starbucks",
		"--start-date", "2024-01-01",
		"--type", "expense",
	)
}

func TestTransactionsCategorize(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions", "categorize",
		"--transaction-id", "txn_quantum-2024-07-21-A7B8C9",
		"--category", "Home > Groceries",
		"--apply-to-future",
		"--notes", "Bulk purchase for party",
	)
}

func TestTransactionsDispute(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions", "dispute",
		"--transaction-id", "txn_quantum-2024-07-21-A7B8C9",
		"--details", "I did not authorize this purchase. My card may have been compromised and I was traveling internationally on this date.",
		"--reason", "unauthorized",
		"--supporting-document", "https://demobank.com/uploads/flight_ticket.png",
	)
}

func TestTransactionsUpdateNotes(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions", "update-notes",
		"--transaction-id", "txn_quantum-2024-07-21-A7B8C9",
		"--notes", "This was a special coffee for a client meeting.",
	)
}
