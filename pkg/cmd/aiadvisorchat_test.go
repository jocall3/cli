// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestAIAdvisorChatRetrieveHistory(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:advisor:chat", "retrieve-history",
		"--limit", "{}",
		"--offset", "{}",
		"--session-id", "session-quantum-xyz-789-alpha",
	)
}

func TestAIAdvisorChatSendMessage(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:advisor:chat", "send-message",
		"--function-response", "{name: send_money, response: {status: success, transactionId: pmt_654321, amountSent: 55.5, recipient: Alex}}",
		"--message", "Can you analyze my recent spending patterns and suggest areas for saving, focusing on my dining expenses?",
		"--session-id", "session-quantum-xyz-789-alpha",
	)
}
