// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestAIIncubatorPitchRetrieveDetails(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:incubator:pitch", "retrieve-details",
		"--pitch-id", "pitch_qw_synergychain-xyz",
	)
}

func TestAIIncubatorPitchSubmit(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:incubator:pitch", "submit",
		"--business-plan", "Quantum-AI powered financial advisor platform leveraging neural networks for predictive analytics and hyper-personalized advice...",
		"--financial-projections", "{profitabilityEstimate: Achieve profitability within 18 months., projectionYears: 3, revenueForecast: [500000, 2000000, 6000000], seedRoundAmount: 2500000, valuationPreMoney: 10000000}",
		"--founding-team", "{experience: '15+ years in AI/ML, PhD in Quantum Computing, ex-Google Brain', name: Dr. Eleanor Vance, role: CEO & Lead AI Scientist}",
		"--founding-team", "{experience: '20+ years in Fintech, ex-Goldman Sachs', name: Marcus Thorne, role: COO & Finance Expert}",
		"--market-opportunity", "The booming digital finance market coupled with demand for truly personalized, AI-driven financial guidance presents a multi-billion dollar opportunity. Our unique quantum-AI approach provides unparalleled accuracy and foresight.",
	)
}

func TestAIIncubatorPitchSubmitFeedback(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:incubator:pitch", "submit-feedback",
		"--pitch-id", "pitch_qw_synergychain-xyz",
		"--answer", "{answer: Our mitigation strategy includes dedicated R&D and new hires with specific expertise., questionId: q_qa-team-001}",
		"--answer", "{answer: Our CAC projections are based on pilot program results showing $500 per enterprise client with a conversion rate of 10% from trials., questionId: q_qa-market-002}",
		"--feedback", "Regarding the technical challenges, our team has allocated 3 months for R&D on quantum-resistant cryptography, mitigating the risk. We've also brought in Dr. Elena Petrova, a leading expert in secure multi-party computation.",
	)
}
