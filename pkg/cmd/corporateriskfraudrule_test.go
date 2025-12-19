// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestCorporateRiskFraudRulesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:risk:fraud:rules", "create",
		"--action", "{details: 'Hold payment, notify sender for additional verification, and escalate to compliance.', type: auto_review, targetTeam: Fraud Prevention Team}",
		"--criteria", "{accountInactivityDays: 90, countryOfOrigin: [US, CA], geographicDistanceKm: 5000, lastLoginDays: 7, noTravelNotification: true, paymentCountMin: 3, recipientCountryRiskLevel: [High, Very High], recipientNew: true, timeframeHours: 24, transactionAmountMin: 5000, transactionType: debit}",
		"--description", "Detects multiple international payments to new beneficiaries in high-risk countries within a short timeframe.",
		"--name", "Suspicious International Payment Pattern",
		"--severity", "Critical",
		"--status", "active",
	)
}

func TestCorporateRiskFraudRulesUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:risk:fraud:rules", "update",
		"--rule-id", "fraud_rule_high_value_inactive",
		"--action", "{details: 'Flag for manual review only, do not block.', type: block, targetTeam: Fraud Prevention Team}",
		"--criteria", "{accountInactivityDays: 60, countryOfOrigin: [US, CA], geographicDistanceKm: 5000, lastLoginDays: 7, noTravelNotification: true, paymentCountMin: 3, recipientCountryRiskLevel: [Low], recipientNew: true, timeframeHours: 24, transactionAmountMin: 7500, transactionType: debit}",
		"--description", "Revised logic for flagging high value transactions from dormant accounts.",
		"--name", "Revised High Value Transaction Rule",
		"--severity", "High",
		"--status", "inactive",
	)
}

func TestCorporateRiskFraudRulesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:risk:fraud:rules", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestCorporateRiskFraudRulesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:risk:fraud:rules", "delete",
		"--rule-id", "fraud_rule_high_value_inactive",
	)
}
