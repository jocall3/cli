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

var paymentsInternationalInitiate = cli.Command{
	Name:  "initiate",
	Usage: "Facilitates the secure initiation of an international wire transfer to a\nbeneficiary in another country and currency, leveraging optimal FX rates and\ntracking capabilities.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "amount",
			Usage:    "The amount to send in the source currency.",
			BodyPath: "amount",
		},
		&requestflag.Flag[any]{
			Name:     "beneficiary",
			Usage:    "Details of the payment beneficiary.",
			BodyPath: "beneficiary",
		},
		&requestflag.Flag[any]{
			Name:     "purpose",
			Usage:    "Purpose of the payment.",
			BodyPath: "purpose",
		},
		&requestflag.Flag[any]{
			Name:     "source-account-id",
			Usage:    "The ID of the user's source account for the payment.",
			BodyPath: "sourceAccountId",
		},
		&requestflag.Flag[any]{
			Name:     "source-currency",
			Usage:    "The ISO 4217 currency code of the source funds.",
			BodyPath: "sourceCurrency",
		},
		&requestflag.Flag[any]{
			Name:     "target-currency",
			Usage:    "The ISO 4217 currency code for the beneficiary's currency.",
			BodyPath: "targetCurrency",
		},
		&requestflag.Flag[any]{
			Name:     "fx-rate-lock",
			Usage:    "If true, attempts to lock the quoted FX rate for a short period.",
			BodyPath: "fxRateLock",
		},
		&requestflag.Flag[string]{
			Name:     "fx-rate-provider",
			Usage:    "Indicates whether to use AI-optimized FX rates or standard market rates.",
			Default:  "proprietary_ai",
			BodyPath: "fxRateProvider",
		},
		&requestflag.Flag[any]{
			Name:     "reference",
			Usage:    "Optional: Your internal reference for this payment.",
			BodyPath: "reference",
		},
	},
	Action:          handlePaymentsInternationalInitiate,
	HideHelpCommand: true,
}

var paymentsInternationalRetrieveStatus = cli.Command{
	Name:  "retrieve-status",
	Usage: "Retrieves the current processing status and details of an initiated\ninternational payment.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "payment-id",
		},
	},
	Action:          handlePaymentsInternationalRetrieveStatus,
	HideHelpCommand: true,
}

func handlePaymentsInternationalInitiate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.PaymentInternationalInitiateParams{}

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
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Payments.International.GetStatus(ctx, cmd.Value("payment-id").(any), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "payments:international retrieve-status", obj, format, transform)
}
