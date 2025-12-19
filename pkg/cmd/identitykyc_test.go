// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestIdentityKYCRetrieveStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"identity:kyc", "retrieve-status",
	)
}

func TestIdentityKYCSubmit(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"identity:kyc", "submit",
		"--country-of-issue", "US",
		"--document-number", "ABC12345",
		"--document-type", "drivers_license",
		"--expiration-date", "2030-01-01",
		"--issue-date", "2020-01-01",
		"--additional-document", "{}\n",
		"--document-back-image", "base64encoded_image_of_drivers_license_back",
		"--document-front-image", "base64encoded_image_of_drivers_license_front",
	)
}
