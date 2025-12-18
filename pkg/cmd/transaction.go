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

var transactionsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Retrieves granular information for a single transaction by its unique ID,\nincluding AI categorization confidence, merchant details, and associated carbon\nfootprint.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "transaction-id",
		},
	},
	Action:          handleTransactionsRetrieve,
	HideHelpCommand: true,
}

var transactionsList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a paginated list of the user's transactions, with extensive options\nfor filtering by type, category, date range, amount, and intelligent AI-driven\nsorting and search capabilities.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "category",
			Usage:     "Filter transactions by their AI-assigned or user-defined category.",
			QueryPath: "category",
		},
		&requestflag.Flag[any]{
			Name:      "end-date",
			Usage:     "Retrieve transactions up to this date (inclusive).",
			QueryPath: "endDate",
		},
		&requestflag.Flag[any]{
			Name:      "limit",
			Usage:     "Maximum number of items to return in a single page.",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "max-amount",
			Usage:     "Filter for transactions with an amount less than or equal to this value.",
			QueryPath: "maxAmount",
		},
		&requestflag.Flag[any]{
			Name:      "min-amount",
			Usage:     "Filter for transactions with an amount greater than or equal to this value.",
			QueryPath: "minAmount",
		},
		&requestflag.Flag[any]{
			Name:      "offset",
			Usage:     "Number of items to skip before starting to collect the result set.",
			QueryPath: "offset",
		},
		&requestflag.Flag[any]{
			Name:      "search-query",
			Usage:     "Free-text search across transaction descriptions, merchants, and notes.",
			QueryPath: "searchQuery",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			Usage:     "Retrieve transactions from this date (inclusive).",
			QueryPath: "startDate",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Filter transactions by type (e.g., income, expense, transfer).",
			QueryPath: "type",
		},
	},
	Action:          handleTransactionsList,
	HideHelpCommand: true,
}

var transactionsCategorize = cli.Command{
	Name:  "categorize",
	Usage: "Allows the user to override or refine the AI's categorization for a transaction,\nimproving future AI accuracy and personal financial reporting.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "transaction-id",
		},
		&requestflag.Flag[any]{
			Name:     "category",
			Usage:    "The new category for the transaction. Can be hierarchical.",
			BodyPath: "category",
		},
		&requestflag.Flag[any]{
			Name:     "apply-to-future",
			Usage:    "If true, the AI will learn from this correction and try to apply it to similar future transactions.",
			BodyPath: "applyToFuture",
		},
		&requestflag.Flag[any]{
			Name:     "notes",
			Usage:    "Optional notes to add to the transaction.",
			BodyPath: "notes",
		},
	},
	Action:          handleTransactionsCategorize,
	HideHelpCommand: true,
}

var transactionsDispute = cli.Command{
	Name:  "dispute",
	Usage: "Begins the process of disputing a specific transaction, providing details and\nsupporting documentation for review by our compliance team and AI.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "transaction-id",
		},
		&requestflag.Flag[any]{
			Name:     "details",
			Usage:    "Detailed explanation of the dispute.",
			BodyPath: "details",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "The primary reason for disputing the transaction.",
			BodyPath: "reason",
		},
		&requestflag.Flag[[]any]{
			Name:     "supporting-document",
			Usage:    "URLs to supporting documents (e.g., receipts, communication).",
			BodyPath: "supportingDocuments",
		},
	},
	Action:          handleTransactionsDispute,
	HideHelpCommand: true,
}

var transactionsUpdateNotes = cli.Command{
	Name:  "update-notes",
	Usage: "Allows the user to add or update personal notes for a specific transaction.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "transaction-id",
		},
		&requestflag.Flag[any]{
			Name:     "notes",
			Usage:    "The personal notes to add or update for the transaction.",
			BodyPath: "notes",
		},
	},
	Action:          handleTransactionsUpdateNotes,
	HideHelpCommand: true,
}

func handleTransactionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("transaction-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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
	_, err = client.Transactions.Get(ctx, cmd.Value("transaction-id").(any), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "transactions retrieve", obj, format, transform)
}

func handleTransactionsList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.TransactionListParams{}

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
	_, err = client.Transactions.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "transactions list", obj, format, transform)
}

func handleTransactionsCategorize(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("transaction-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.TransactionCategorizeParams{}

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
	_, err = client.Transactions.Categorize(
		ctx,
		cmd.Value("transaction-id").(any),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "transactions categorize", obj, format, transform)
}

func handleTransactionsDispute(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("transaction-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.TransactionDisputeParams{}

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
	_, err = client.Transactions.Dispute(
		ctx,
		cmd.Value("transaction-id").(any),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "transactions dispute", obj, format, transform)
}

func handleTransactionsUpdateNotes(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("transaction-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.TransactionUpdateNotesParams{}

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
	_, err = client.Transactions.UpdateNotes(
		ctx,
		cmd.Value("transaction-id").(any),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "transactions update-notes", obj, format, transform)
}
