// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestAIOracleSimulateRunAdvanced(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:oracle:simulate", "run-advanced",
		"--prompt", "Evaluate the long-term impact of a sudden job loss combined with a variable market downturn, analyzing worst-case and best-case recovery scenarios over a decade.",
		"--scenario", "{durationYears: 10, events: [{details: {durationMonths: 6, severanceAmount: 10000, unemploymentBenefits: 2000}, type: job_loss}, {details: {impactPercentage: 0.15, recoveryYears: 3}, type: market_downturn}], name: Job Loss & Mild Market Recovery, sensitivityAnalysisParams: [{max: 0.07, min: 0.03, paramName: marketRecoveryRate, step: 0.01}]}",
		"--global-economic-factors", "{inflationRate: 0.03, interestRateBaseline: 0.05}",
		"--personal-assumptions", "{annualSavingsRate: 0.15, riskTolerance: aggressive}",
	)
}

func TestAIOracleSimulateRunStandard(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:oracle:simulate", "run-standard",
		"--prompt", "What if I invest an additional $1,000 per month into my aggressive growth portfolio for the next 5 years?",
		"--parameters", "{durationYears: 5, monthlyInvestmentAmount: 1000, riskTolerance: aggressive}",
	)
}
