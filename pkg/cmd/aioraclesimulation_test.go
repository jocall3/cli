// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestAIOracleSimulationsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:oracle:simulations", "retrieve",
		"--simulation-id", "sim_oracle-growth-2024-xyz",
	)
}

func TestAIOracleSimulationsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:oracle:simulations", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestAIOracleSimulationsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:oracle:simulations", "delete",
		"--simulation-id", "sim_oracle-growth-2024-xyz",
	)
}
