// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestUsersMeBiometricsDeregister(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me:biometrics", "deregister",
	)
}

func TestUsersMeBiometricsEnroll(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me:biometrics", "enroll",
		"--biometric-signature", "base64encoded_facial_template_for_enrollment",
		"--biometric-type", "facial_recognition",
		"--device-id", "dev_mobile_ios_aabbcc",
		"--device-name", "My Primary iPhone",
	)
}

func TestUsersMeBiometricsStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me:biometrics", "status",
	)
}

func TestUsersMeBiometricsVerify(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me:biometrics", "verify",
		"--biometric-signature", "base64encoded_one_time_fingerprint_proof",
		"--biometric-type", "fingerprint",
		"--device-id", "dev_mobile_android_ddeeff",
	)
}
