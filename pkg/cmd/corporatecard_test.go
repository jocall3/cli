// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestCorporateCardsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:cards", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestCorporateCardsCreateVirtual(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:cards", "create-virtual",
		"--controls", "{atmWithdrawals: false, contactlessPayments: false, dailyLimit: 500, internationalTransactions: false, merchantCategoryRestrictions: [Advertising], monthlyLimit: 1000, onlineTransactions: true, singleTransactionLimit: 200, vendorRestrictions: [Facebook Ads, Google Ads]}",
		"--expiration-date", "2025-12-31",
		"--holder-name", "Marketing Campaign Q4",
		"--purpose", "Online advertising for Q4 campaigns",
		"--associated-employee-id", "emp_marketing_01",
		"--spending-policy-id", "policy_marketing_fixed",
	)
}

func TestCorporateCardsFreeze(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:cards", "freeze",
		"--card-id", "corp_card_xyz987654",
		"--freeze",
	)
}

func TestCorporateCardsListTransactions(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:cards", "list-transactions",
		"--card-id", "corp_card_xyz987654",
		"--end-date", "2024-12-31",
		"--limit", "{}",
		"--offset", "{}",
		"--start-date", "2024-01-01",
	)
}

func TestCorporateCardsUpdateControls(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:cards", "update-controls",
		"--card-id", "corp_card_xyz987654",
		"--atm-withdrawals",
		"--contactless-payments",
		"--daily-limit", "750",
		"--international-transactions",
		"--merchant-category-restriction", "Software Subscriptions",
		"--merchant-category-restriction", "Conferences",
		"--monthly-limit", "3000",
		"--online-transactions",
		"--single-transaction-limit", "1000",
		"--vendor-restriction", "Amazon",
		"--vendor-restriction", "Uber",
	)
}
