// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jocall3/cli/internal/apiquery"
	"github.com/jocall3/cli/internal/requestflag"
	"github.com/jocall3/go"
	"github.com/jocall3/go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var web3TransactionsInitiateTransfer = cli.Command{
	Name:  "initiate-transfer",
	Usage: "Prepares and initiates a cryptocurrency transfer from a connected wallet to a\nspecified recipient address. Requires user confirmation (e.g., via wallet\nsignature).",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "amount",
			Usage:    "The amount of cryptocurrency to transfer.",
			BodyPath: "amount",
		},
		&requestflag.Flag[any]{
			Name:     "asset-symbol",
			Usage:    "Symbol of the crypto asset to transfer (e.g., ETH, USDC).",
			BodyPath: "assetSymbol",
		},
		&requestflag.Flag[any]{
			Name:     "blockchain-network",
			Usage:    "The blockchain network for the transfer.",
			BodyPath: "blockchainNetwork",
		},
		&requestflag.Flag[any]{
			Name:     "recipient-address",
			Usage:    "The recipient's blockchain address.",
			BodyPath: "recipientAddress",
		},
		&requestflag.Flag[any]{
			Name:     "source-wallet-id",
			Usage:    "ID of the connected wallet from which to send funds.",
			BodyPath: "sourceWalletId",
		},
		&requestflag.Flag[any]{
			Name:     "gas-price-gwei",
			Usage:    "Optional: Gas price in Gwei for Ethereum-based transactions.",
			BodyPath: "gasPriceGwei",
		},
		&requestflag.Flag[any]{
			Name:     "memo",
			Usage:    "Optional: A short memo or note for the transaction.",
			BodyPath: "memo",
		},
	},
	Action:          handleWeb3TransactionsInitiateTransfer,
	HideHelpCommand: true,
}

func handleWeb3TransactionsInitiateTransfer(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.Web3TransactionInitiateTransferParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web3.Transactions.InitiateTransfer(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "web3:transactions initiate-transfer", obj, format, transform)
}
