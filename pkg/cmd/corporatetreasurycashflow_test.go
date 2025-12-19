// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestCorporateTreasuryCashFlowGetForecast(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:treasury:cash-flow", "get-forecast",
		"--forecast-horizon-days", "{}",
		"--include-scenario-analysis", "{}",
	)
}
