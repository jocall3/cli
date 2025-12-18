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

var budgetsCreate = cli.Command{
	Name:  "create",
	Usage: "Creates a new financial budget for the user, with optional AI auto-population of\ncategories and amounts.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "end-date",
			Usage:    "End date of the budget period.",
			BodyPath: "endDate",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Name of the new budget.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "period",
			Usage:    "The frequency or period of the budget.",
			BodyPath: "period",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			Usage:    "Start date of the budget period.",
			BodyPath: "startDate",
		},
		&requestflag.Flag[any]{
			Name:     "total-amount",
			Usage:    "Total amount allocated for the entire budget.",
			BodyPath: "totalAmount",
		},
		&requestflag.Flag[any]{
			Name:     "ai-auto-populate",
			Usage:    "If true, AI will automatically populate categories and amounts based on historical spending.",
			BodyPath: "aiAutoPopulate",
		},
		&requestflag.Flag[any]{
			Name:     "alert-threshold",
			Usage:    "Percentage threshold at which an alert is triggered.",
			BodyPath: "alertThreshold",
		},
		&requestflag.Flag[[]any]{
			Name:     "category",
			Usage:    "Initial breakdown of the budget by categories.",
			BodyPath: "categories",
		},
	},
	Action:          handleBudgetsCreate,
	HideHelpCommand: true,
}

var budgetsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Retrieves detailed information for a specific budget, including current\nspending, remaining amounts, and AI recommendations.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "budget-id",
		},
	},
	Action:          handleBudgetsRetrieve,
	HideHelpCommand: true,
}

var budgetsUpdate = cli.Command{
	Name:  "update",
	Usage: "Updates the parameters of an existing budget, such as total amount, dates, or\ncategories.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "budget-id",
		},
		&requestflag.Flag[any]{
			Name:     "alert-threshold",
			Usage:    "Updated percentage threshold for alerts.",
			BodyPath: "alertThreshold",
		},
		&requestflag.Flag[[]any]{
			Name:     "category",
			Usage:    "Updated breakdown of the budget by categories. Existing categories will be updated, new ones added.",
			BodyPath: "categories",
		},
		&requestflag.Flag[any]{
			Name:     "end-date",
			Usage:    "Updated end date of the budget period.",
			BodyPath: "endDate",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Updated name of the budget.",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			Usage:    "Updated start date of the budget period.",
			BodyPath: "startDate",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Updated status of the budget.",
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "total-amount",
			Usage:    "Updated total amount for the entire budget.",
			BodyPath: "totalAmount",
		},
	},
	Action:          handleBudgetsUpdate,
	HideHelpCommand: true,
}

var budgetsList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of all active and historical budgets for the authenticated\nuser.",
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
	Action:          handleBudgetsList,
	HideHelpCommand: true,
}

var budgetsDelete = cli.Command{
	Name:  "delete",
	Usage: "Deletes a specific budget from the user's profile.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "budget-id",
		},
	},
	Action:          handleBudgetsDelete,
	HideHelpCommand: true,
}

func handleBudgetsCreate(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetNewParams{}

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
	_, err = client.Budgets.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "budgets create", obj, format, transform)
}

func handleBudgetsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("budget-id") && len(unusedArgs) > 0 {
		cmd.Set("budget-id", unusedArgs[0])
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
	_, err = client.Budgets.Get(ctx, cmd.Value("budget-id").(any), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "budgets retrieve", obj, format, transform)
}

func handleBudgetsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("budget-id") && len(unusedArgs) > 0 {
		cmd.Set("budget-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetUpdateParams{}

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
	_, err = client.Budgets.Update(
		ctx,
		cmd.Value("budget-id").(any),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "budgets update", obj, format, transform)
}

func handleBudgetsList(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetListParams{}

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
	_, err = client.Budgets.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "budgets list", obj, format, transform)
}

func handleBudgetsDelete(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("budget-id") && len(unusedArgs) > 0 {
		cmd.Set("budget-id", unusedArgs[0])
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

	return client.Budgets.Delete(ctx, cmd.Value("budget-id").(any), options...)
}
