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

var transactionsRecurringCreate = cli.Command{
	Name:  "create",
	Usage: "Defines a new recurring transaction pattern for future tracking and budgeting.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "amount",
			Usage:    "Amount of the recurring transaction.",
			BodyPath: "amount",
		},
		&requestflag.Flag[any]{
			Name:     "category",
			Usage:    "Category of the recurring transaction.",
			BodyPath: "category",
		},
		&requestflag.Flag[any]{
			Name:     "currency",
			Usage:    "ISO 4217 currency code.",
			BodyPath: "currency",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Description of the recurring transaction.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "frequency",
			Usage:    "Frequency of the recurring transaction.",
			BodyPath: "frequency",
		},
		&requestflag.Flag[any]{
			Name:     "linked-account-id",
			Usage:    "ID of the account to associate with this recurring transaction.",
			BodyPath: "linkedAccountId",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			Usage:    "The date when this recurring transaction is expected to start.",
			BodyPath: "startDate",
		},
	},
	Action:          handleTransactionsRecurringCreate,
	HideHelpCommand: true,
}

var transactionsRecurringList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of all detected or user-defined recurring transactions, useful\nfor budget tracking and subscription management.",
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
	Action:          handleTransactionsRecurringList,
	HideHelpCommand: true,
}

func handleTransactionsRecurringCreate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.TransactionRecurringNewParams{}

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
	_, err = client.Transactions.Recurring.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "transactions:recurring create", obj, format, transform)
}

func handleTransactionsRecurringList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.TransactionRecurringListParams{}

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
	_, err = client.Transactions.Recurring.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "transactions:recurring list", obj, format, transform)
}
