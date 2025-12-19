// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestPaymentsInternationalInitiate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments:international", "initiate",
		"--amount", "5000",
		"--beneficiary", "{address: 'Hauptstrasse 1, 10115 Berlin, Germany', bankName: Deutsche Bank, name: Maria Schmidt, accountNumber: {}, iban: DE89370400440532013000, routingNumber: {}, swiftBic: DEUTDEFF}",
		"--purpose", "Vendor payment for Q2 services.",
		"--source-account-id", "acc_chase_checking_4567",
		"--source-currency", "USD",
		"--target-currency", "EUR",
		"--fx-rate-lock",
		"--fx-rate-provider", "proprietary_ai",
		"--reference", "{}",
	)
}

func TestPaymentsInternationalRetrieveStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments:international", "retrieve-status",
		"--payment-id", "int_pmt_xyz7890",
	)
}
