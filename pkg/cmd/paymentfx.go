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

var paymentsFxConvert = cli.Command{
	Name:  "convert",
	Usage: "Executes an instant currency conversion between two currencies, either from a\nbalance or into a specified account.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "source-account-id",
			BodyPath: "sourceAccountId",
		},
		&requestflag.Flag[float64]{
			Name:     "source-amount",
			BodyPath: "sourceAmount",
		},
		&requestflag.Flag[string]{
			Name:     "source-currency",
			BodyPath: "sourceCurrency",
		},
		&requestflag.Flag[string]{
			Name:     "target-currency",
			BodyPath: "targetCurrency",
		},
		&requestflag.Flag[bool]{
			Name:     "fx-rate-lock",
			BodyPath: "fxRateLock",
		},
		&requestflag.Flag[string]{
			Name:     "target-account-id",
			BodyPath: "targetAccountId",
		},
	},
	Action:          handlePaymentsFxConvert,
	HideHelpCommand: true,
}

var paymentsFxRetrieveRates = cli.Command{
	Name:  "retrieve-rates",
	Usage: "Retrieves current and AI-predicted future foreign exchange rates for a specified\ncurrency pair, including bid/ask spreads and historical volatility data for\ninformed decisions.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "base-currency",
			Usage:     "The base currency code (e.g., USD).",
			QueryPath: "baseCurrency",
		},
		&requestflag.Flag[any]{
			Name:      "target-currency",
			Usage:     "The target currency code (e.g., EUR).",
			QueryPath: "targetCurrency",
		},
		&requestflag.Flag[any]{
			Name:      "forecast-days",
			Usage:     "Number of days into the future to provide an AI-driven prediction.",
			Default:   7,
			QueryPath: "forecastDays",
		},
	},
	Action:          handlePaymentsFxRetrieveRates,
	HideHelpCommand: true,
}

func handlePaymentsFxConvert(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.PaymentFxConvertParams{}

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
	_, err = client.Payments.Fx.Convert(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "payments:fx convert", obj, format, transform)
}

func handlePaymentsFxRetrieveRates(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.PaymentFxGetRatesParams{}

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
	_, err = client.Payments.Fx.GetRates(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "payments:fx retrieve-rates", obj, format, transform)
}
