// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestCorporateAnomaliesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:anomalies", "list",
		"--end-date", "2024-12-31",
		"--entity-type", "Transaction",
		"--limit", "{}",
		"--offset", "{}",
		"--severity", "Critical",
		"--start-date", "2024-01-01",
		"--status", "New",
	)
}

func TestCorporateAnomaliesUpdateStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:anomalies", "update-status",
		"--anomaly-id", "anom_risk-2024-07-21-D1E2F3",
		"--status", "Resolved",
		"--resolution-notes", "Confirmed legitimate transaction after contacting vendor. Marked as resolved.",
	)
}
