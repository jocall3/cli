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

var web3WalletsList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of all securely linked cryptocurrency wallets (e.g., MetaMask,\nLedger integration), showing their addresses, associated networks, and\nverification status.",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The maximum number of items to return.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "The number of items to skip before starting to collect the result set.",
			QueryPath: "offset",
		},
	},
	Action:          handleWeb3WalletsList,
	HideHelpCommand: true,
}

var web3WalletsConnect = cli.Command{
	Name:  "connect",
	Usage: "Initiates the process to securely connect a new cryptocurrency wallet to the\nuser's profile, typically involving a signed message or OAuth flow from the\nwallet provider.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "blockchain-network",
			BodyPath: "blockchainNetwork",
		},
		&requestflag.Flag[string]{
			Name:     "signed-message",
			Usage:    "A message signed by the wallet's private key to prove ownership.",
			BodyPath: "signedMessage",
		},
		&requestflag.Flag[string]{
			Name:     "wallet-address",
			BodyPath: "walletAddress",
		},
		&requestflag.Flag[string]{
			Name:     "wallet-provider",
			BodyPath: "walletProvider",
		},
	},
	Action:          handleWeb3WalletsConnect,
	HideHelpCommand: true,
}

var web3WalletsRetrieveBalances = cli.Command{
	Name:  "retrieve-balances",
	Usage: "Retrieves the current balances of all recognized crypto assets within a specific\nconnected wallet.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "wallet-id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The maximum number of items to return.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "The number of items to skip before starting to collect the result set.",
			QueryPath: "offset",
		},
	},
	Action:          handleWeb3WalletsRetrieveBalances,
	HideHelpCommand: true,
}

func handleWeb3WalletsList(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.Web3WalletListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web3.Wallets.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "web3:wallets list", obj, format, transform)
}

func handleWeb3WalletsConnect(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.Web3WalletConnectParams{}

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
	_, err = client.Web3.Wallets.Connect(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "web3:wallets connect", obj, format, transform)
}

func handleWeb3WalletsRetrieveBalances(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("wallet-id") && len(unusedArgs) > 0 {
		cmd.Set("wallet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.Web3WalletGetBalancesParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web3.Wallets.GetBalances(
		ctx,
		cmd.Value("wallet-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "web3:wallets retrieve-balances", obj, format, transform)
}
