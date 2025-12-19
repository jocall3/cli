// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestGoalsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"goals", "create",
		"--name", "Dream Vacation Fund",
		"--target-amount", "15000",
		"--target-date", "2026-06-30",
		"--type", "large_purchase",
		"--contributing-account", "{}",
		"--generate-ai-plan",
		"--initial-contribution", "1000",
		"--risk-tolerance", "conservative",
	)
}

func TestGoalsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"goals", "retrieve",
		"--goal-id", "goal_retirement_2050",
	)
}

func TestGoalsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"goals", "update",
		"--goal-id", "goal_retirement_2050",
		"--contributing-account", "{}",
		"--generate-ai-plan",
		"--name", "Revised Retirement Fund Goal",
		"--risk-tolerance", "conservative",
		"--status", "paused",
		"--target-amount", "1200000",
		"--target-date", "2029-12-31",
	)
}

func TestGoalsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"goals", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestGoalsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"goals", "delete",
		"--goal-id", "goal_retirement_2050",
	)
}
