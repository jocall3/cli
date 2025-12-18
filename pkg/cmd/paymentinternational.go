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

var paymentsInternationalInitiate = cli.Command{
	Name:  "initiate",
	Usage: "Facilitates the secure initiation of an international wire transfer to a\nbeneficiary in another country and currency, leveraging optimal FX rates and\ntracking capabilities.",
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "amount",
			BodyPath: "amount",
		},
		&requestflag.Flag[any]{
			Name:     "beneficiary",
			BodyPath: "beneficiary",
		},
		&requestflag.Flag[string]{
			Name:     "source-account-id",
			BodyPath: "sourceAccountId",
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
			Name:     "fx-rate-provider",
			BodyPath: "fxRateProvider",
		},
		&requestflag.Flag[string]{
			Name:     "purpose",
			BodyPath: "purpose",
		},
	},
	Action:          handlePaymentsInternationalInitiate,
	HideHelpCommand: true,
}

var paymentsInternationalRetrieveStatus = cli.Command{
	Name:  "retrieve-status",
	Usage: "Retrieves the current processing status and details of an initiated\ninternational payment.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "payment-id",
		},
	},
	Action:          handlePaymentsInternationalRetrieveStatus,
	HideHelpCommand: true,
}

func handlePaymentsInternationalInitiate(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.PaymentInternationalInitiateParams{}

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
	_, err = client.Payments.International.Initiate(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "payments:international initiate", obj, format, transform)
}

func handlePaymentsInternationalRetrieveStatus(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("payment-id") && len(unusedArgs) > 0 {
		cmd.Set("payment-id", unusedArgs[0])
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
	_, err = client.Payments.International.GetStatus(ctx, cmd.Value("payment-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "payments:international retrieve-status", obj, format, transform)
}
