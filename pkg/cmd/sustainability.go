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

var sustainabilityPurchaseCarbonOffsets = cli.Command{
	Name:  "purchase-carbon-offsets",
	Usage: "Allows users to purchase carbon offsets to neutralize their estimated carbon\nfootprint, supporting environmental initiatives.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "amount-kg-co2e",
			Usage:    "The amount of carbon dioxide equivalent to offset in kilograms.",
			BodyPath: "amountKgCO2e",
		},
		&requestflag.Flag[any]{
			Name:     "offset-project",
			Usage:    "Optional: The specific carbon offset project to support.",
			BodyPath: "offsetProject",
		},
		&requestflag.Flag[any]{
			Name:     "payment-account-id",
			Usage:    "The ID of the user's account to use for payment.",
			BodyPath: "paymentAccountId",
		},
	},
	Action:          handleSustainabilityPurchaseCarbonOffsets,
	HideHelpCommand: true,
}

var sustainabilityRetrieveCarbonFootprint = cli.Command{
	Name:            "retrieve-carbon-footprint",
	Usage:           "Generates a detailed report of the user's estimated carbon footprint based on\ntransaction data, lifestyle choices, and AI-driven impact assessments, offering\ninsights and reduction strategies.",
	Flags:           []cli.Flag{},
	Action:          handleSustainabilityRetrieveCarbonFootprint,
	HideHelpCommand: true,
}

func handleSustainabilityPurchaseCarbonOffsets(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.SustainabilityPurchaseCarbonOffsetsParams{}

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
	_, err = client.Sustainability.PurchaseCarbonOffsets(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sustainability purchase-carbon-offsets", obj, format, transform)
}

func handleSustainabilityRetrieveCarbonFootprint(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.Sustainability.GetCarbonFootprint(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sustainability retrieve-carbon-footprint", obj, format, transform)
}
