// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestAIAdsListGenerated(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:ads", "list-generated",
		"--limit", "{}",
		"--offset", "{}",
		"--status", "done",
	)
}

func TestAIAdsRetrieveStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:ads", "retrieve-status",
		"--operation-id", "op-video-gen-12345-abcde",
	)
}
