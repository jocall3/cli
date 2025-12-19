// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestInvestmentsAssetsSearch(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"investments:assets", "search",
		"--query", "Tesla",
		"--limit", "{}",
		"--min-esg-score", "7",
		"--offset", "{}",
	)
}
