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

var corporateCardsList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a comprehensive list of all physical and virtual corporate cards\nassociated with the user's organization, including their status, assigned\nholder, and current spending controls.",
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
	Action:          handleCorporateCardsList,
	HideHelpCommand: true,
}

var corporateCardsCreateVirtual = cli.Command{
	Name:  "create-virtual",
	Usage: "Creates and issues a new virtual corporate card with specified spending limits,\nmerchant restrictions, and expiration dates, ideal for secure online purchases\nand temporary projects.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "controls",
			BodyPath: "controls",
		},
		&requestflag.Flag[string]{
			Name:     "holder-name",
			BodyPath: "holderName",
		},
		&requestflag.Flag[string]{
			Name:     "associated-employee-id",
			BodyPath: "associatedEmployeeId",
		},
		&requestflag.Flag[requestflag.DateValue]{
			Name:     "expiration-date",
			BodyPath: "expirationDate",
		},
		&requestflag.Flag[string]{
			Name:     "purpose",
			BodyPath: "purpose",
		},
	},
	Action:          handleCorporateCardsCreateVirtual,
	HideHelpCommand: true,
}

var corporateCardsFreeze = cli.Command{
	Name:  "freeze",
	Usage: "Immediately changes the frozen status of a corporate card, preventing or\nallowing transactions in real-time, critical for security and expense\nmanagement.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "card-id",
		},
		&requestflag.Flag[bool]{
			Name:     "freeze",
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
		&requestflag.Flag[string]{
			Name: "card-id",
		},
		&requestflag.Flag[requestflag.DateValue]{
			Name:      "end-date",
			Usage:     "The end date for the query range (inclusive).",
			QueryPath: "endDate",
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
		&requestflag.Flag[requestflag.DateValue]{
			Name:      "start-date",
			Usage:     "The start date for the query range (inclusive).",
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
		&requestflag.Flag[string]{
			Name: "card-id",
		},
		&requestflag.Flag[bool]{
			Name:     "atm-withdrawals",
			BodyPath: "atmWithdrawals",
		},
		&requestflag.Flag[bool]{
			Name:     "contactless-payments",
			BodyPath: "contactlessPayments",
		},
		&requestflag.Flag[float64]{
			Name:     "daily-limit",
			BodyPath: "dailyLimit",
		},
		&requestflag.Flag[bool]{
			Name:     "international-transactions",
			BodyPath: "internationalTransactions",
		},
		&requestflag.Flag[[]string]{
			Name:     "merchant-category-restriction",
			BodyPath: "merchantCategoryRestrictions",
		},
		&requestflag.Flag[float64]{
			Name:     "monthly-limit",
			BodyPath: "monthlyLimit",
		},
		&requestflag.Flag[bool]{
			Name:     "online-transactions",
			BodyPath: "onlineTransactions",
		},
		&requestflag.Flag[float64]{
			Name:     "single-transaction-limit",
			BodyPath: "singleTransactionLimit",
		},
		&requestflag.Flag[[]string]{
			Name:     "vendor-restriction",
			BodyPath: "vendorRestrictions",
		},
	},
	Action:          handleCorporateCardsUpdateControls,
	HideHelpCommand: true,
}

func handleCorporateCardsList(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardListParams{}

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
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardNewVirtualParams{}

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
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardFreezeParams{}

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
		cmd.Value("card-id").(string),
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
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardListTransactionsParams{}

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
		cmd.Value("card-id").(string),
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
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardUpdateControlsParams{}

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
		cmd.Value("card-id").(string),
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
