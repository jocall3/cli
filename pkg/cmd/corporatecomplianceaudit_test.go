// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestCorporateComplianceAuditsRequest(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:compliance:audits", "request",
		"--audit-scope", "all_transactions",
		"--end-date", "2024-06-30",
		"--regulatory-framework", "AML",
		"--regulatory-framework", "PCI-DSS",
		"--start-date", "2024-01-01",
		"--additional-context", "{}",
	)
}

func TestCorporateComplianceAuditsRetrieveReport(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate:compliance:audits", "retrieve-report",
		"--audit-id", "audit_corp_xyz789",
	)
}
