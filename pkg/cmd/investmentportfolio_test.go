// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestInvestmentsPortfoliosCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"investments:portfolios", "create",
		"--currency", "USD",
		"--initial-investment", "10000",
		"--name", "My First Growth Portfolio",
		"--risk-tolerance", "conservative",
		"--type", "diversified",
		"--ai-auto-allocate",
		"--linked-account-id", "acc_chase_checking_4567",
	)
}

func TestInvestmentsPortfoliosRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"investments:portfolios", "retrieve",
		"--portfolio-id", "portfolio_equity_growth",
	)
}

func TestInvestmentsPortfoliosUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"investments:portfolios", "update",
		"--portfolio-id", "portfolio_equity_growth",
		"--ai-rebalancing-frequency", "quarterly",
		"--name", "Revised Growth Portfolio",
		"--risk-tolerance", "conservative",
	)
}

func TestInvestmentsPortfoliosList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"investments:portfolios", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestInvestmentsPortfoliosRebalance(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"investments:portfolios", "rebalance",
		"--portfolio-id", "portfolio_equity_growth",
		"--target-risk-tolerance", "conservative",
		"--confirmation-required",
		"--dry-run",
	)
}
