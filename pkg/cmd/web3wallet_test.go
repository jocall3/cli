// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestWeb3WalletsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"web3:wallets", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestWeb3WalletsConnect(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"web3:wallets", "connect",
		"--blockchain-network", "Ethereum",
		"--signed-message", "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
		"--wallet-address", "0x123abc456def7890...",
		"--wallet-provider", "MetaMask",
		"--request-write-access",
	)
}

func TestWeb3WalletsRetrieveBalances(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"web3:wallets", "retrieve-balances",
		"--wallet-id", "wallet_conn_eth_0xabc123",
		"--limit", "{}",
		"--offset", "{}",
	)
}
