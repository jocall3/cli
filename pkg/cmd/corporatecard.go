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

var corporateCardsList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a comprehensive list of all physical and virtual corporate cards\nassociated with the user's organization, including their status, assigned\nholder, and current spending controls.",
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
	Action:          handleCorporateCardsList,
	HideHelpCommand: true,
}

var corporateCardsCreateVirtual = cli.Command{
	Name:  "create-virtual",
	Usage: "Creates and issues a new virtual corporate card with specified spending limits,\nmerchant restrictions, and expiration dates, ideal for secure online purchases\nand temporary projects.",
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "controls",
			Usage:    "Granular spending controls for a corporate card.",
			BodyPath: "controls",
		},
		&requestflag.Flag[any]{
			Name:     "expiration-date",
			Usage:    "Expiration date for the virtual card (YYYY-MM-DD).",
			BodyPath: "expirationDate",
		},
		&requestflag.Flag[any]{
			Name:     "holder-name",
			Usage:    "Name to appear on the virtual card.",
			BodyPath: "holderName",
		},
		&requestflag.Flag[any]{
			Name:     "purpose",
			Usage:    "Brief description of the virtual card's purpose.",
			BodyPath: "purpose",
		},
		&requestflag.Flag[any]{
			Name:     "associated-employee-id",
			Usage:    "Optional: ID of the employee or department this card is for.",
			BodyPath: "associatedEmployeeId",
		},
		&requestflag.Flag[any]{
			Name:     "spending-policy-id",
			Usage:    "Optional: ID of a spending policy to link with this virtual card.",
			BodyPath: "spendingPolicyId",
		},
	},
	Action:          handleCorporateCardsCreateVirtual,
	HideHelpCommand: true,
}

var corporateCardsFreeze = cli.Command{
	Name:  "freeze",
	Usage: "Immediately changes the frozen status of a corporate card, preventing or\nallowing transactions in real-time, critical for security and expense\nmanagement.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "card-id",
		},
		&requestflag.Flag[any]{
			Name:     "freeze",
			Usage:    "Set to `true` to freeze the card, `false` to unfreeze.",
			BodyPath: "freeze",
		},
	},
	Action:          handleCorporateCardsFreeze,
	HideHelpCommand: true,
}

var corporateCardsListTransactions = cli.Command{
	Name:  "list-transactions",
	Usage: "Retrieves a paginated list of transactions made with a specific corporate card,\nincluding AI categorization and compliance flags.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "card-id",
		},
		&requestflag.Flag[any]{
			Name:      "end-date",
			Usage:     "End date for filtering results (inclusive, YYYY-MM-DD).",
			QueryPath: "endDate",
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
		&requestflag.Flag[any]{
			Name:      "start-date",
			Usage:     "Start date for filtering results (inclusive, YYYY-MM-DD).",
			QueryPath: "startDate",
		},
	},
	Action:          handleCorporateCardsListTransactions,
	HideHelpCommand: true,
}

var corporateCardsUpdateControls = cli.Command{
	Name:  "update-controls",
	Usage: "Updates the sophisticated spending controls, limits, and policy overrides for a\nspecific corporate card, enabling real-time adjustments for security and budget\nadherence.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "card-id",
		},
		&requestflag.Flag[any]{
			Name:     "atm-withdrawals",
			Usage:    "If true, ATM cash withdrawals are allowed.",
			BodyPath: "atmWithdrawals",
		},
		&requestflag.Flag[any]{
			Name:     "contactless-payments",
			Usage:    "If true, contactless payments are allowed.",
			BodyPath: "contactlessPayments",
		},
		&requestflag.Flag[any]{
			Name:     "daily-limit",
			Usage:    "Maximum spending limit per day (null for no limit).",
			BodyPath: "dailyLimit",
		},
		&requestflag.Flag[any]{
			Name:     "international-transactions",
			Usage:    "If true, international transactions are allowed.",
			BodyPath: "internationalTransactions",
		},
		&requestflag.Flag[[]any]{
			Name:     "merchant-category-restriction",
			Usage:    "List of allowed merchant categories. If empty, all are allowed unless explicitly denied.",
			BodyPath: "merchantCategoryRestrictions",
		},
		&requestflag.Flag[any]{
			Name:     "monthly-limit",
			Usage:    "Maximum spending limit per month (null for no limit).",
			BodyPath: "monthlyLimit",
		},
		&requestflag.Flag[any]{
			Name:     "online-transactions",
			Usage:    "If true, online transactions are allowed.",
			BodyPath: "onlineTransactions",
		},
		&requestflag.Flag[any]{
			Name:     "single-transaction-limit",
			Usage:    "Maximum amount for a single transaction (null for no limit).",
			BodyPath: "singleTransactionLimit",
		},
		&requestflag.Flag[[]any]{
			Name:     "vendor-restriction",
			Usage:    "List of allowed vendors/merchants by name.",
			BodyPath: "vendorRestrictions",
		},
	},
	Action:          handleCorporateCardsUpdateControls,
	HideHelpCommand: true,
}

func handleCorporateCardsList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.CorporateCardListParams{}

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
	_, err = client.Corporate.Cards.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:cards list", obj, format, transform)
}

func handleCorporateCardsCreateVirtual(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.CorporateCardNewVirtualParams{}

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
	_, err = client.Corporate.Cards.NewVirtual(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:cards create-virtual", obj, format, transform)
}

func handleCorporateCardsFreeze(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.CorporateCardFreezeParams{}

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
	_, err = client.Corporate.Cards.Freeze(
		ctx,
		interface{}(cmd.Value("card-id").(any)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:cards freeze", obj, format, transform)
}

func handleCorporateCardsListTransactions(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.CorporateCardListTransactionsParams{}

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
	_, err = client.Corporate.Cards.ListTransactions(
		ctx,
		interface{}(cmd.Value("card-id").(any)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:cards list-transactions", obj, format, transform)
}

func handleCorporateCardsUpdateControls(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.CorporateCardUpdateControlsParams{}

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
	_, err = client.Corporate.Cards.UpdateControls(
		ctx,
		interface{}(cmd.Value("card-id").(any)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:cards update-controls", obj, format, transform)
}
