// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestWeb3TransactionsInitiateTransfer(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"web3:transactions", "initiate-transfer",
		"--amount", "0.1",
		"--asset-symbol", "ETH",
		"--blockchain-network", "Ethereum",
		"--recipient-address", "0xdef4567890abcdef1234567890abcdef1234567890",
		"--source-wallet-id", "wallet_conn_eth_0xabc123",
		"--gas-price-gwei", "50",
		"--memo", "Payment for services",
	)
}
