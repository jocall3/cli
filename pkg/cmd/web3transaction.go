// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/1231-cli/internal/apiquery"
	"github.com/stainless-sdks/1231-cli/internal/requestflag"
	"github.com/stainless-sdks/1231-go"
	"github.com/stainless-sdks/1231-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var web3TransactionsInitiateTransfer = cli.Command{
	Name:  "initiate-transfer",
	Usage: "Prepares and initiates a cryptocurrency transfer from a connected wallet to a\nspecified recipient address. Requires user confirmation (e.g., via wallet\nsignature).",
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "amount",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "asset-symbol",
			BodyPath: "assetSymbol",
		},
		&requestflag.Flag[string]{
			Name:     "blockchain-network",
			BodyPath: "blockchainNetwork",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-address",
			BodyPath: "recipientAddress",
		},
		&requestflag.Flag[string]{
			Name:     "source-wallet-id",
			BodyPath: "sourceWalletId",
		},
		&requestflag.Flag[float64]{
			Name:     "gas-price-gwei",
			BodyPath: "gasPriceGwei",
		},
		&requestflag.Flag[string]{
			Name:     "memo",
			BodyPath: "memo",
		},
	},
	Action:          handleWeb3TransactionsInitiateTransfer,
	HideHelpCommand: true,
}

func handleWeb3TransactionsInitiateTransfer(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.Web3TransactionInitiateTransferParams{}

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
