// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestCorporatePerformSanctionScreening(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate", "perform-sanction-screening",
		"--country", "US",
		"--entity-type", "individual",
		"--name", "John Doe",
		"--address", "{city: Anytown, country: USA, state: CA, street: 123 Main St, zip: '90210'}",
		"--date-of-birth", "1970-01-01",
		"--identification-number", "{}",
	)
}
