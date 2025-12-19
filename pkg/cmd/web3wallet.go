// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jocall3/1231-cli/internal/apiquery"
	"github.com/jocall3/1231-cli/internal/requestflag"
	"github.com/jocall3/go"
	"github.com/jocall3/go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var web3WalletsList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of all securely linked cryptocurrency wallets (e.g., MetaMask,\nLedger integration), showing their addresses, associated networks, and\nverification status.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "limit",
			Usage:     "Maximum number of items to return in a single page.",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "offset",
			Usage:     "Number of items to skip before starting to collect the result set.",
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
		&requestflag.Flag[any]{
			Name:     "blockchain-network",
			Usage:    "The blockchain network for this wallet (e.g., Ethereum, Solana).",
			BodyPath: "blockchainNetwork",
		},
		&requestflag.Flag[any]{
			Name:     "signed-message",
			Usage:    "A message cryptographically signed by the wallet owner to prove ownership/intent.",
			BodyPath: "signedMessage",
		},
		&requestflag.Flag[any]{
			Name:     "wallet-address",
			Usage:    "The public address of the cryptocurrency wallet.",
			BodyPath: "walletAddress",
		},
		&requestflag.Flag[any]{
			Name:     "wallet-provider",
			Usage:    "The name of the wallet provider (e.g., MetaMask, Phantom).",
			BodyPath: "walletProvider",
		},
		&requestflag.Flag[any]{
			Name:     "request-write-access",
			Usage:    "If true, requests write access to initiate transactions from this wallet.",
			BodyPath: "requestWriteAccess",
		},
	},
	Action:          handleWeb3WalletsConnect,
	HideHelpCommand: true,
}

var web3WalletsRetrieveBalances = cli.Command{
	Name:  "retrieve-balances",
	Usage: "Retrieves the current balances of all recognized crypto assets within a specific\nconnected wallet.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "wallet-id",
		},
		&requestflag.Flag[any]{
			Name:      "limit",
			Usage:     "Maximum number of items to return in a single page.",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "offset",
			Usage:     "Number of items to skip before starting to collect the result set.",
			QueryPath: "offset",
		},
	},
	Action:          handleWeb3WalletsRetrieveBalances,
	HideHelpCommand: true,
}

func handleWeb3WalletsList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.Web3WalletListParams{}

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
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.Web3WalletConnectParams{}

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
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("wallet-id") && len(unusedArgs) > 0 {
		cmd.Set("wallet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.Web3WalletGetBalancesParams{}

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
		interface{}(cmd.Value("wallet-id").(any)),
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
