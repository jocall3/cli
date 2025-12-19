// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestUsersMeDevicesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me:devices", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestUsersMeDevicesDeregister(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me:devices", "deregister",
		"--device-id", "dev_mobile_ios_aabbcc",
	)
}

func TestUsersMeDevicesRegister(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me:devices", "register",
		"--device-type", "mobile",
		"--model", "Samsung Galaxy S24 Ultra",
		"--os", "Android 14",
		"--biometric-signature", "base64encoded_android_biometric_proof",
		"--device-name", "Alex's Work Phone",
		"--push-token", "ExponentPushToken[yzzzyzzzyzzzyzzz]",
	)
}
