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

var transactionsRecurringCreate = cli.Command{
	Name:  "create",
	Usage: "Defines a new recurring transaction pattern for future tracking and budgeting.",
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "amount",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "category",
			BodyPath: "category",
		},
		&requestflag.Flag[string]{
			Name:     "currency",
			BodyPath: "currency",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "frequency",
			BodyPath: "frequency",
		},
		&requestflag.Flag[requestflag.DateValue]{
			Name:     "start-date",
			BodyPath: "startDate",
		},
		&requestflag.Flag[string]{
			Name:     "linked-account-id",
			BodyPath: "linkedAccountId",
		},
	},
	Action:          handleTransactionsRecurringCreate,
	HideHelpCommand: true,
}

var transactionsRecurringList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of all detected or user-defined recurring transactions, useful\nfor budget tracking and subscription management.",
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
	Action:          handleTransactionsRecurringList,
	HideHelpCommand: true,
}

func handleTransactionsRecurringCreate(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionRecurringNewParams{}

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
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionRecurringListParams{}

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
