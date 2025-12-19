// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestAIIncubatorListPitches(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:incubator", "list-pitches",
		"--limit", "{}",
		"--offset", "{}",
		"--status", "feedback_required",
	)
}
